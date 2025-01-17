// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rancher/steve/pkg/sqlcache/db/transaction (interfaces: Stmt)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -package store -destination ./tx_mocks_test.go github.com/rancher/steve/pkg/sqlcache/db/transaction Stmt
//

// Package store is a generated GoMock package.
package store

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStmt is a mock of Stmt interface.
type MockStmt struct {
	ctrl     *gomock.Controller
	recorder *MockStmtMockRecorder
}

// MockStmtMockRecorder is the mock recorder for MockStmt.
type MockStmtMockRecorder struct {
	mock *MockStmt
}

// NewMockStmt creates a new mock instance.
func NewMockStmt(ctrl *gomock.Controller) *MockStmt {
	mock := &MockStmt{ctrl: ctrl}
	mock.recorder = &MockStmtMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStmt) EXPECT() *MockStmtMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockStmt) Exec(arg0 ...any) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockStmtMockRecorder) Exec(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockStmt)(nil).Exec), arg0...)
}

// Query mocks base method.
func (m *MockStmt) Query(arg0 ...any) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockStmtMockRecorder) Query(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockStmt)(nil).Query), arg0...)
}

// QueryContext mocks base method.
func (m *MockStmt) QueryContext(arg0 context.Context, arg1 ...any) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryContext", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContext indicates an expected call of QueryContext.
func (mr *MockStmtMockRecorder) QueryContext(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContext", reflect.TypeOf((*MockStmt)(nil).QueryContext), varargs...)
}
