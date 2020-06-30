// Code generated by MockGen. DO NOT EDIT.
// Source: ../pbagreement/agreement.pb.go

// Package pbagreementtest is a generated GoMock package.
package pbagreementtest

import (
	context "context"
	pbagreement "github.com/athomecomar/athome/pb/pbagreement"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockAgreementClient is a mock of AgreementClient interface
type MockAgreementClient struct {
	ctrl     *gomock.Controller
	recorder *MockAgreementClientMockRecorder
}

// MockAgreementClientMockRecorder is the mock recorder for MockAgreementClient
type MockAgreementClientMockRecorder struct {
	mock *MockAgreementClient
}

// NewMockAgreementClient creates a new mock instance
func NewMockAgreementClient(ctrl *gomock.Controller) *MockAgreementClient {
	mock := &MockAgreementClient{ctrl: ctrl}
	mock.recorder = &MockAgreementClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAgreementClient) EXPECT() *MockAgreementClientMockRecorder {
	return m.recorder
}

// Retrieve mocks base method
func (m *MockAgreementClient) Retrieve(ctx context.Context, in *pbagreement.RetrieveRequest, opts ...grpc.CallOption) (*pbagreement.RetrieveResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Retrieve", varargs...)
	ret0, _ := ret[0].(*pbagreement.RetrieveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Retrieve indicates an expected call of Retrieve
func (mr *MockAgreementClientMockRecorder) Retrieve(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retrieve", reflect.TypeOf((*MockAgreementClient)(nil).Retrieve), varargs...)
}

// MockAgreementServer is a mock of AgreementServer interface
type MockAgreementServer struct {
	ctrl     *gomock.Controller
	recorder *MockAgreementServerMockRecorder
}

// MockAgreementServerMockRecorder is the mock recorder for MockAgreementServer
type MockAgreementServerMockRecorder struct {
	mock *MockAgreementServer
}

// NewMockAgreementServer creates a new mock instance
func NewMockAgreementServer(ctrl *gomock.Controller) *MockAgreementServer {
	mock := &MockAgreementServer{ctrl: ctrl}
	mock.recorder = &MockAgreementServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAgreementServer) EXPECT() *MockAgreementServerMockRecorder {
	return m.recorder
}

// Retrieve mocks base method
func (m *MockAgreementServer) Retrieve(arg0 context.Context, arg1 *pbagreement.RetrieveRequest) (*pbagreement.RetrieveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retrieve", arg0, arg1)
	ret0, _ := ret[0].(*pbagreement.RetrieveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Retrieve indicates an expected call of Retrieve
func (mr *MockAgreementServerMockRecorder) Retrieve(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retrieve", reflect.TypeOf((*MockAgreementServer)(nil).Retrieve), arg0, arg1)
}
