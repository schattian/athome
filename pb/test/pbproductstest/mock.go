// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/athomecomar/athome/pb/pbproducts (interfaces: CreatorClient,Creator_SecondClient,Creator_ThirdClient,ViewerClient)

// Package pbproductstest is a generated GoMock package.
package pbproductstest

import (
	context "context"
	pbproducts "github.com/athomecomar/athome/pb/pbproducts"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

// MockCreatorClient is a mock of CreatorClient interface
type MockCreatorClient struct {
	ctrl     *gomock.Controller
	recorder *MockCreatorClientMockRecorder
}

// MockCreatorClientMockRecorder is the mock recorder for MockCreatorClient
type MockCreatorClientMockRecorder struct {
	mock *MockCreatorClient
}

// NewMockCreatorClient creates a new mock instance
func NewMockCreatorClient(ctrl *gomock.Controller) *MockCreatorClient {
	mock := &MockCreatorClient{ctrl: ctrl}
	mock.recorder = &MockCreatorClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCreatorClient) EXPECT() *MockCreatorClientMockRecorder {
	return m.recorder
}

// CloneDraftLine mocks base method
func (m *MockCreatorClient) CloneDraftLine(arg0 context.Context, arg1 *pbproducts.CloneDraftLineRequest, arg2 ...grpc.CallOption) (*pbproducts.CloneDraftLineResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CloneDraftLine", varargs...)
	ret0, _ := ret[0].(*pbproducts.CloneDraftLineResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloneDraftLine indicates an expected call of CloneDraftLine
func (mr *MockCreatorClientMockRecorder) CloneDraftLine(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloneDraftLine", reflect.TypeOf((*MockCreatorClient)(nil).CloneDraftLine), varargs...)
}

// DeleteDraft mocks base method
func (m *MockCreatorClient) DeleteDraft(arg0 context.Context, arg1 *pbproducts.DeleteDraftRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteDraft", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDraft indicates an expected call of DeleteDraft
func (mr *MockCreatorClientMockRecorder) DeleteDraft(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDraft", reflect.TypeOf((*MockCreatorClient)(nil).DeleteDraft), varargs...)
}

// DeleteDraftLine mocks base method
func (m *MockCreatorClient) DeleteDraftLine(arg0 context.Context, arg1 *pbproducts.DeleteDraftLineRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteDraftLine", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDraftLine indicates an expected call of DeleteDraftLine
func (mr *MockCreatorClientMockRecorder) DeleteDraftLine(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDraftLine", reflect.TypeOf((*MockCreatorClient)(nil).DeleteDraftLine), varargs...)
}

// First mocks base method
func (m *MockCreatorClient) First(arg0 context.Context, arg1 ...grpc.CallOption) (pbproducts.Creator_FirstClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(pbproducts.Creator_FirstClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// First indicates an expected call of First
func (mr *MockCreatorClientMockRecorder) First(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockCreatorClient)(nil).First), varargs...)
}

// Next mocks base method
func (m *MockCreatorClient) Next(arg0 context.Context, arg1 *pbproducts.StageChangeRequest, arg2 ...grpc.CallOption) (*pbproducts.StageChangeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Next", varargs...)
	ret0, _ := ret[0].(*pbproducts.StageChangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockCreatorClientMockRecorder) Next(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockCreatorClient)(nil).Next), varargs...)
}

// Prev mocks base method
func (m *MockCreatorClient) Prev(arg0 context.Context, arg1 *pbproducts.StageChangeRequest, arg2 ...grpc.CallOption) (*pbproducts.StageChangeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Prev", varargs...)
	ret0, _ := ret[0].(*pbproducts.StageChangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prev indicates an expected call of Prev
func (mr *MockCreatorClientMockRecorder) Prev(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prev", reflect.TypeOf((*MockCreatorClient)(nil).Prev), varargs...)
}

// RetrieveDraft mocks base method
func (m *MockCreatorClient) RetrieveDraft(arg0 context.Context, arg1 *pbproducts.RetrieveDraftRequest, arg2 ...grpc.CallOption) (*pbproducts.Draft, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveDraft", varargs...)
	ret0, _ := ret[0].(*pbproducts.Draft)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveDraft indicates an expected call of RetrieveDraft
func (mr *MockCreatorClientMockRecorder) RetrieveDraft(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveDraft", reflect.TypeOf((*MockCreatorClient)(nil).RetrieveDraft), varargs...)
}

// Second mocks base method
func (m *MockCreatorClient) Second(arg0 context.Context, arg1 ...grpc.CallOption) (pbproducts.Creator_SecondClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Second", varargs...)
	ret0, _ := ret[0].(pbproducts.Creator_SecondClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Second indicates an expected call of Second
func (mr *MockCreatorClientMockRecorder) Second(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Second", reflect.TypeOf((*MockCreatorClient)(nil).Second), varargs...)
}

// Third mocks base method
func (m *MockCreatorClient) Third(arg0 context.Context, arg1 ...grpc.CallOption) (pbproducts.Creator_ThirdClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Third", varargs...)
	ret0, _ := ret[0].(pbproducts.Creator_ThirdClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Third indicates an expected call of Third
func (mr *MockCreatorClientMockRecorder) Third(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Third", reflect.TypeOf((*MockCreatorClient)(nil).Third), varargs...)
}

// MockCreator_SecondClient is a mock of Creator_SecondClient interface
type MockCreator_SecondClient struct {
	ctrl     *gomock.Controller
	recorder *MockCreator_SecondClientMockRecorder
}

// MockCreator_SecondClientMockRecorder is the mock recorder for MockCreator_SecondClient
type MockCreator_SecondClientMockRecorder struct {
	mock *MockCreator_SecondClient
}

// NewMockCreator_SecondClient creates a new mock instance
func NewMockCreator_SecondClient(ctrl *gomock.Controller) *MockCreator_SecondClient {
	mock := &MockCreator_SecondClient{ctrl: ctrl}
	mock.recorder = &MockCreator_SecondClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCreator_SecondClient) EXPECT() *MockCreator_SecondClientMockRecorder {
	return m.recorder
}

// CloseAndRecv mocks base method
func (m *MockCreator_SecondClient) CloseAndRecv() (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAndRecv")
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAndRecv indicates an expected call of CloseAndRecv
func (mr *MockCreator_SecondClientMockRecorder) CloseAndRecv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAndRecv", reflect.TypeOf((*MockCreator_SecondClient)(nil).CloseAndRecv))
}

// CloseSend mocks base method
func (m *MockCreator_SecondClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockCreator_SecondClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockCreator_SecondClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockCreator_SecondClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockCreator_SecondClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockCreator_SecondClient)(nil).Context))
}

// Header mocks base method
func (m *MockCreator_SecondClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockCreator_SecondClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockCreator_SecondClient)(nil).Header))
}

// RecvMsg mocks base method
func (m *MockCreator_SecondClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockCreator_SecondClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockCreator_SecondClient)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockCreator_SecondClient) Send(arg0 *pbproducts.SecondRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockCreator_SecondClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockCreator_SecondClient)(nil).Send), arg0)
}

// SendMsg mocks base method
func (m *MockCreator_SecondClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockCreator_SecondClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockCreator_SecondClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockCreator_SecondClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockCreator_SecondClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockCreator_SecondClient)(nil).Trailer))
}

// MockCreator_ThirdClient is a mock of Creator_ThirdClient interface
type MockCreator_ThirdClient struct {
	ctrl     *gomock.Controller
	recorder *MockCreator_ThirdClientMockRecorder
}

// MockCreator_ThirdClientMockRecorder is the mock recorder for MockCreator_ThirdClient
type MockCreator_ThirdClientMockRecorder struct {
	mock *MockCreator_ThirdClient
}

// NewMockCreator_ThirdClient creates a new mock instance
func NewMockCreator_ThirdClient(ctrl *gomock.Controller) *MockCreator_ThirdClient {
	mock := &MockCreator_ThirdClient{ctrl: ctrl}
	mock.recorder = &MockCreator_ThirdClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCreator_ThirdClient) EXPECT() *MockCreator_ThirdClientMockRecorder {
	return m.recorder
}

// CloseAndRecv mocks base method
func (m *MockCreator_ThirdClient) CloseAndRecv() (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAndRecv")
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAndRecv indicates an expected call of CloseAndRecv
func (mr *MockCreator_ThirdClientMockRecorder) CloseAndRecv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAndRecv", reflect.TypeOf((*MockCreator_ThirdClient)(nil).CloseAndRecv))
}

// CloseSend mocks base method
func (m *MockCreator_ThirdClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockCreator_ThirdClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockCreator_ThirdClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockCreator_ThirdClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockCreator_ThirdClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockCreator_ThirdClient)(nil).Context))
}

// Header mocks base method
func (m *MockCreator_ThirdClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockCreator_ThirdClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockCreator_ThirdClient)(nil).Header))
}

// RecvMsg mocks base method
func (m *MockCreator_ThirdClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockCreator_ThirdClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockCreator_ThirdClient)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockCreator_ThirdClient) Send(arg0 *pbproducts.ThirdRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockCreator_ThirdClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockCreator_ThirdClient)(nil).Send), arg0)
}

// SendMsg mocks base method
func (m *MockCreator_ThirdClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockCreator_ThirdClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockCreator_ThirdClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockCreator_ThirdClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockCreator_ThirdClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockCreator_ThirdClient)(nil).Trailer))
}

// MockViewerClient is a mock of ViewerClient interface
type MockViewerClient struct {
	ctrl     *gomock.Controller
	recorder *MockViewerClientMockRecorder
}

// MockViewerClientMockRecorder is the mock recorder for MockViewerClient
type MockViewerClientMockRecorder struct {
	mock *MockViewerClient
}

// NewMockViewerClient creates a new mock instance
func NewMockViewerClient(ctrl *gomock.Controller) *MockViewerClient {
	mock := &MockViewerClient{ctrl: ctrl}
	mock.recorder = &MockViewerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockViewerClient) EXPECT() *MockViewerClientMockRecorder {
	return m.recorder
}

// RetrieveProduct mocks base method
func (m *MockViewerClient) RetrieveProduct(arg0 context.Context, arg1 *pbproducts.RetrieveProductRequest, arg2 ...grpc.CallOption) (*pbproducts.Product, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveProduct", varargs...)
	ret0, _ := ret[0].(*pbproducts.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveProduct indicates an expected call of RetrieveProduct
func (mr *MockViewerClientMockRecorder) RetrieveProduct(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveProduct", reflect.TypeOf((*MockViewerClient)(nil).RetrieveProduct), varargs...)
}

// RetrieveProducts mocks base method
func (m *MockViewerClient) RetrieveProducts(arg0 context.Context, arg1 *pbproducts.RetrieveProductsRequest, arg2 ...grpc.CallOption) (*pbproducts.RetrieveProductsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveProducts", varargs...)
	ret0, _ := ret[0].(*pbproducts.RetrieveProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveProducts indicates an expected call of RetrieveProducts
func (mr *MockViewerClientMockRecorder) RetrieveProducts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveProducts", reflect.TypeOf((*MockViewerClient)(nil).RetrieveProducts), varargs...)
}

// RetrieveProductsMaxVolWeight mocks base method
func (m *MockViewerClient) RetrieveProductsMaxVolWeight(arg0 context.Context, arg1 *pbproducts.RetrieveProductsRequest, arg2 ...grpc.CallOption) (*pbproducts.RetrieveProductsMalVolWeightResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveProductsMaxVolWeight", varargs...)
	ret0, _ := ret[0].(*pbproducts.RetrieveProductsMalVolWeightResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveProductsMaxVolWeight indicates an expected call of RetrieveProductsMaxVolWeight
func (mr *MockViewerClientMockRecorder) RetrieveProductsMaxVolWeight(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveProductsMaxVolWeight", reflect.TypeOf((*MockViewerClient)(nil).RetrieveProductsMaxVolWeight), varargs...)
}

// SearchProducts mocks base method
func (m *MockViewerClient) SearchProducts(arg0 context.Context, arg1 *pbproducts.SearchProductsRequest, arg2 ...grpc.CallOption) (*pbproducts.SearchProductsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchProducts", varargs...)
	ret0, _ := ret[0].(*pbproducts.SearchProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchProducts indicates an expected call of SearchProducts
func (mr *MockViewerClientMockRecorder) SearchProducts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProducts", reflect.TypeOf((*MockViewerClient)(nil).SearchProducts), varargs...)
}
