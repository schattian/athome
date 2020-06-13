// Code generated by MockGen. DO NOT EDIT.
// Source: ../pbaddress/address.pb.go

// Package pbaddresstest is a generated GoMock package.
package pbaddresstest

import (
	context "context"
	pbaddress "github.com/athomecomar/athome/pb/pbaddress"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockAddressesClient is a mock of AddressesClient interface
type MockAddressesClient struct {
	ctrl     *gomock.Controller
	recorder *MockAddressesClientMockRecorder
}

// MockAddressesClientMockRecorder is the mock recorder for MockAddressesClient
type MockAddressesClientMockRecorder struct {
	mock *MockAddressesClient
}

// NewMockAddressesClient creates a new mock instance
func NewMockAddressesClient(ctrl *gomock.Controller) *MockAddressesClient {
	mock := &MockAddressesClient{ctrl: ctrl}
	mock.recorder = &MockAddressesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAddressesClient) EXPECT() *MockAddressesClientMockRecorder {
	return m.recorder
}

// CreateAddress mocks base method
func (m *MockAddressesClient) CreateAddress(ctx context.Context, in *pbaddress.CreateAddressRequest, opts ...grpc.CallOption) (*pbaddress.CreateAddressResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAddress", varargs...)
	ret0, _ := ret[0].(*pbaddress.CreateAddressResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAddress indicates an expected call of CreateAddress
func (mr *MockAddressesClientMockRecorder) CreateAddress(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAddress", reflect.TypeOf((*MockAddressesClient)(nil).CreateAddress), varargs...)
}

// RetrieveAddress mocks base method
func (m *MockAddressesClient) RetrieveAddress(ctx context.Context, in *pbaddress.RetrieveAddressRequest, opts ...grpc.CallOption) (*pbaddress.Address, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveAddress", varargs...)
	ret0, _ := ret[0].(*pbaddress.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveAddress indicates an expected call of RetrieveAddress
func (mr *MockAddressesClientMockRecorder) RetrieveAddress(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveAddress", reflect.TypeOf((*MockAddressesClient)(nil).RetrieveAddress), varargs...)
}

// RetrieveMyAddresses mocks base method
func (m *MockAddressesClient) RetrieveMyAddresses(ctx context.Context, in *pbaddress.RetrieveMyAddressesRequest, opts ...grpc.CallOption) (*pbaddress.RetrieveMyAddressesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveMyAddresses", varargs...)
	ret0, _ := ret[0].(*pbaddress.RetrieveMyAddressesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveMyAddresses indicates an expected call of RetrieveMyAddresses
func (mr *MockAddressesClientMockRecorder) RetrieveMyAddresses(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveMyAddresses", reflect.TypeOf((*MockAddressesClient)(nil).RetrieveMyAddresses), varargs...)
}

// MockAddressesServer is a mock of AddressesServer interface
type MockAddressesServer struct {
	ctrl     *gomock.Controller
	recorder *MockAddressesServerMockRecorder
}

// MockAddressesServerMockRecorder is the mock recorder for MockAddressesServer
type MockAddressesServerMockRecorder struct {
	mock *MockAddressesServer
}

// NewMockAddressesServer creates a new mock instance
func NewMockAddressesServer(ctrl *gomock.Controller) *MockAddressesServer {
	mock := &MockAddressesServer{ctrl: ctrl}
	mock.recorder = &MockAddressesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAddressesServer) EXPECT() *MockAddressesServerMockRecorder {
	return m.recorder
}

// CreateAddress mocks base method
func (m *MockAddressesServer) CreateAddress(arg0 context.Context, arg1 *pbaddress.CreateAddressRequest) (*pbaddress.CreateAddressResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAddress", arg0, arg1)
	ret0, _ := ret[0].(*pbaddress.CreateAddressResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAddress indicates an expected call of CreateAddress
func (mr *MockAddressesServerMockRecorder) CreateAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAddress", reflect.TypeOf((*MockAddressesServer)(nil).CreateAddress), arg0, arg1)
}

// RetrieveAddress mocks base method
func (m *MockAddressesServer) RetrieveAddress(arg0 context.Context, arg1 *pbaddress.RetrieveAddressRequest) (*pbaddress.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveAddress", arg0, arg1)
	ret0, _ := ret[0].(*pbaddress.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveAddress indicates an expected call of RetrieveAddress
func (mr *MockAddressesServerMockRecorder) RetrieveAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveAddress", reflect.TypeOf((*MockAddressesServer)(nil).RetrieveAddress), arg0, arg1)
}

// RetrieveMyAddresses mocks base method
func (m *MockAddressesServer) RetrieveMyAddresses(arg0 context.Context, arg1 *pbaddress.RetrieveMyAddressesRequest) (*pbaddress.RetrieveMyAddressesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveMyAddresses", arg0, arg1)
	ret0, _ := ret[0].(*pbaddress.RetrieveMyAddressesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveMyAddresses indicates an expected call of RetrieveMyAddresses
func (mr *MockAddressesServerMockRecorder) RetrieveMyAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveMyAddresses", reflect.TypeOf((*MockAddressesServer)(nil).RetrieveMyAddresses), arg0, arg1)
}