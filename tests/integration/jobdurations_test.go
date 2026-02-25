package tests

import (
	"context"
	//"encoding/json"
	"io"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rancher/steve/pkg/server"
	"github.com/rancher/steve/pkg/sqlcache/informer/factory"
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
)

var testdataJobDurationsDir = filepath.Join("testdata", "jobdurations")

// JobDurationTestConfig defines the structure for list test YAML files
type JobDurationTestConfig struct {
	SchemaID string `yaml:"schemaID"`
	Tests    []struct {
		Description    string              `yaml:"description"`
		Namespace      string              `yaml:"namespace"`
		Query          string              `yaml:"query"`
		Expect         []map[string]string `yaml:"expect"`
	} `yaml:"tests"`
}

func (i *IntegrationSuite) TestJobDurations() {
	ctx := i.T().Context()

	manifestsFile := filepath.Join(testdataJobDurationsDir, "jobdurations.manifests.yaml")
	gvrs := make(map[k8sschema.GroupVersionResource]struct{})
	i.doManifest(ctx, manifestsFile, func(ctx context.Context, obj *unstructured.Unstructured, gvr k8sschema.GroupVersionResource) error {
		gvrs[gvr] = struct{}{}
		return i.doApply(ctx, obj, gvr)
	})
	defer i.doManifestReversed(ctx, manifestsFile, i.doDelete)

	// Wait for pods to stabilize - restarting-pod needs time to restart once
	i.T().Log("Waiting for jobs to run (allow 15 secs)")
	time.Sleep(15 * time.Second)
	i.T().Log("Done waiting")

	// Run SQL mode only - these tests are specifically for SQL cache with multi-value field support
	i.runJobDurationsTest(ctx, true, gvrs)
}

func (i *IntegrationSuite) runJobDurationsTest(ctx context.Context, sqlCache bool, gvrs map[k8sschema.GroupVersionResource]struct{}) {
	var steveHandler http.Handler
	var err error
	steveHandler, err = server.New(ctx, i.restCfg, &server.Options{
		SQLCache: true,
		SQLCacheFactoryOptions: factory.CacheFactoryOptions{
			GCInterval:  15 * time.Minute,
			GCKeepCount: 1000,
		},
	})
	i.Require().NoError(err)

	steveServer := httptest.NewServer(steveHandler)
	defer steveServer.Close()

	// Wait for cache to be populated
	if sqlCache {
		time.Sleep(2 * time.Second)
	}

	matches, err := filepath.Glob(filepath.Join(testdataJobDurationsDir, "*.test.yaml"))
	i.Require().NoError(err)

	for _, testFile := range matches {
		name := filepath.Base(testFile)
		name = strings.TrimSuffix(name, ".test.yaml")

		i.Run(name, func() {
			var config JobDurationTestConfig
			data, err := os.ReadFile(testFile)
			i.Require().NoError(err)
			err = yaml.Unmarshal(data, &config)
			i.Require().NoError(err)

			for _, test := range config.Tests {
				i.Run(fmt.Sprintf("%s", test.Description), func() {
					url := buildURLRaw(steveServer.URL, config.SchemaID, test.Namespace, test.Query)
					i.T().Logf("Testing: %s", url)

					req, err := http.NewRequest("GET", url, nil)
					i.Require().NoError(err)

					resp, err := http.DefaultClient.Do(req)
					i.Require().NoError(err)
					defer resp.Body.Close()

					i.Assert().Equal(http.StatusOK, resp.StatusCode, "request should succeed")
					/*  */
					body, err := io.ReadAll(resp.Body)
					if err != nil {
						return
					}
/* */
					fmt.Println(string(body))
/*

					type DataItem struct {
						Metadata struct {
							Name string `json:"name"`
							Namespace string `json:"namespace"`
							Fields []string `json:"fields"`
						} `json:"metadata"`
					}
					type Response struct {
						Data []DataItem `json:"data"`
					}

					var parsed Response
					err = json.NewDecoder(resp.Body).Decode(&parsed)
					i.Require().NoError(err)

					i.Assert().Equal(len(test.Expect), len(parsed.Data))
					for index, job := range parsed.Data {
						fmt.Printf("QQQ: Reading in job output %+v\n", job)
						i.Assert().Equal(test.Expect[index]["name"], job.Metadata.Name)
						i.Assert().Equal(test.Expect[index]["duration"], job.Metadata.Fields[3])
					}
				/* */
				})
			}
		})
	}
}

