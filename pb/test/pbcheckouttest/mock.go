// Code generated by MockGen. DO NOT EDIT.
// Source: ../pbcheckout/checkout.pb.go

// Package pbcheckouttest is a generated GoMock package.
package pbcheckouttest

import (
	context "context"
	pbcheckout "github.com/athomecomar/athome/pb/pbcheckout"
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockReservationsClient is a mock of ReservationsClient interface
type MockReservationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockReservationsClientMockRecorder
}

// MockReservationsClientMockRecorder is the mock recorder for MockReservationsClient
type MockReservationsClientMockRecorder struct {
	mock *MockReservationsClient
}

// NewMockReservationsClient creates a new mock instance
func NewMockReservationsClient(ctrl *gomock.Controller) *MockReservationsClient {
	mock := &MockReservationsClient{ctrl: ctrl}
	mock.recorder = &MockReservationsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReservationsClient) EXPECT() *MockReservationsClientMockRecorder {
	return m.recorder
}

// CreateReservation mocks base method
func (m *MockReservationsClient) CreateReservation(ctx context.Context, in *pbcheckout.CreateReservationRequest, opts ...grpc.CallOption) (*pbcheckout.CreateReservationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateReservation", varargs...)
	ret0, _ := ret[0].(*pbcheckout.CreateReservationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReservation indicates an expected call of CreateReservation
func (mr *MockReservationsClientMockRecorder) CreateReservation(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReservation", reflect.TypeOf((*MockReservationsClient)(nil).CreateReservation), varargs...)
}

// RetrieveCurrent mocks base method
func (m *MockReservationsClient) RetrieveCurrent(ctx context.Context, in *pbcheckout.RetrieveCurrentRequest, opts ...grpc.CallOption) (*pbcheckout.RetrieveReservationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveCurrent", varargs...)
	ret0, _ := ret[0].(*pbcheckout.RetrieveReservationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveCurrent indicates an expected call of RetrieveCurrent
func (mr *MockReservationsClientMockRecorder) RetrieveCurrent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveCurrent", reflect.TypeOf((*MockReservationsClient)(nil).RetrieveCurrent), varargs...)
}

// StateMachine mocks base method
func (m *MockReservationsClient) StateMachine(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*pbcheckout.StateMachineResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StateMachine", varargs...)
	ret0, _ := ret[0].(*pbcheckout.StateMachineResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMachine indicates an expected call of StateMachine
func (mr *MockReservationsClientMockRecorder) StateMachine(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMachine", reflect.TypeOf((*MockReservationsClient)(nil).StateMachine), varargs...)
}

// Prev mocks base method
func (m *MockReservationsClient) Prev(ctx context.Context, in *pbcheckout.UpdateOrderStateRequest, opts ...grpc.CallOption) (*pbcheckout.RetrieveReservationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Prev", varargs...)
	ret0, _ := ret[0].(*pbcheckout.RetrieveReservationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prev indicates an expected call of Prev
func (mr *MockReservationsClientMockRecorder) Prev(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prev", reflect.TypeOf((*MockReservationsClient)(nil).Prev), varargs...)
}

// Next mocks base method
func (m *MockReservationsClient) Next(ctx context.Context, in *pbcheckout.UpdateOrderStateRequest, opts ...grpc.CallOption) (*pbcheckout.RetrieveReservationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Next", varargs...)
	ret0, _ := ret[0].(*pbcheckout.RetrieveReservationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockReservationsClientMockRecorder) Next(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockReservationsClient)(nil).Next), varargs...)
}

// Cancel mocks base method
func (m *MockReservationsClient) Cancel(ctx context.Context, in *pbcheckout.UpdateOrderStateRequest, opts ...grpc.CallOption) (*pbcheckout.RetrieveReservationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Cancel", varargs...)
	ret0, _ := ret[0].(*pbcheckout.RetrieveReservationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cancel indicates an expected call of Cancel
func (mr *MockReservationsClientMockRecorder) Cancel(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockReservationsClient)(nil).Cancel), varargs...)
}

// MockReservationsServer is a mock of ReservationsServer interface
type MockReservationsServer struct {
	ctrl     *gomock.Controller
	recorder *MockReservationsServerMockRecorder
}

// MockReservationsServerMockRecorder is the mock recorder for MockReservationsServer
type MockReservationsServerMockRecorder struct {
	mock *MockReservationsServer
}

// NewMockReservationsServer creates a new mock instance
func NewMockReservationsServer(ctrl *gomock.Controller) *MockReservationsServer {
	mock := &MockReservationsServer{ctrl: ctrl}
	mock.recorder = &MockReservationsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReservationsServer) EXPECT() *MockReservationsServerMockRecorder {
	return m.recorder
}

// CreateReservation mocks base method
func (m *MockReservationsServer) CreateReservation(arg0 context.Context, arg1 *pbcheckout.CreateReservationRequest) (*pbcheckout.CreateReservationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReservation", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.CreateReservationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReservation indicates an expected call of CreateReservation
func (mr *MockReservationsServerMockRecorder) CreateReservation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReservation", reflect.TypeOf((*MockReservationsServer)(nil).CreateReservation), arg0, arg1)
}

// RetrieveCurrent mocks base method
func (m *MockReservationsServer) RetrieveCurrent(arg0 context.Context, arg1 *pbcheckout.RetrieveCurrentRequest) (*pbcheckout.RetrieveReservationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveCurrent", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.RetrieveReservationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveCurrent indicates an expected call of RetrieveCurrent
func (mr *MockReservationsServerMockRecorder) RetrieveCurrent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveCurrent", reflect.TypeOf((*MockReservationsServer)(nil).RetrieveCurrent), arg0, arg1)
}

// StateMachine mocks base method
func (m *MockReservationsServer) StateMachine(arg0 context.Context, arg1 *empty.Empty) (*pbcheckout.StateMachineResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMachine", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.StateMachineResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMachine indicates an expected call of StateMachine
func (mr *MockReservationsServerMockRecorder) StateMachine(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMachine", reflect.TypeOf((*MockReservationsServer)(nil).StateMachine), arg0, arg1)
}

// Prev mocks base method
func (m *MockReservationsServer) Prev(arg0 context.Context, arg1 *pbcheckout.UpdateOrderStateRequest) (*pbcheckout.RetrieveReservationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prev", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.RetrieveReservationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prev indicates an expected call of Prev
func (mr *MockReservationsServerMockRecorder) Prev(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prev", reflect.TypeOf((*MockReservationsServer)(nil).Prev), arg0, arg1)
}

// Next mocks base method
func (m *MockReservationsServer) Next(arg0 context.Context, arg1 *pbcheckout.UpdateOrderStateRequest) (*pbcheckout.RetrieveReservationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.RetrieveReservationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockReservationsServerMockRecorder) Next(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockReservationsServer)(nil).Next), arg0, arg1)
}

// Cancel mocks base method
func (m *MockReservationsServer) Cancel(arg0 context.Context, arg1 *pbcheckout.UpdateOrderStateRequest) (*pbcheckout.RetrieveReservationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.RetrieveReservationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cancel indicates an expected call of Cancel
func (mr *MockReservationsServerMockRecorder) Cancel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockReservationsServer)(nil).Cancel), arg0, arg1)
}

// MockBookingsClient is a mock of BookingsClient interface
type MockBookingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockBookingsClientMockRecorder
}

// MockBookingsClientMockRecorder is the mock recorder for MockBookingsClient
type MockBookingsClientMockRecorder struct {
	mock *MockBookingsClient
}

// NewMockBookingsClient creates a new mock instance
func NewMockBookingsClient(ctrl *gomock.Controller) *MockBookingsClient {
	mock := &MockBookingsClient{ctrl: ctrl}
	mock.recorder = &MockBookingsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBookingsClient) EXPECT() *MockBookingsClientMockRecorder {
	return m.recorder
}

// CreateBooking mocks base method
func (m *MockBookingsClient) CreateBooking(ctx context.Context, in *pbcheckout.CreateBookingRequest, opts ...grpc.CallOption) (*pbcheckout.CreateBookingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateBooking", varargs...)
	ret0, _ := ret[0].(*pbcheckout.CreateBookingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBooking indicates an expected call of CreateBooking
func (mr *MockBookingsClientMockRecorder) CreateBooking(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBooking", reflect.TypeOf((*MockBookingsClient)(nil).CreateBooking), varargs...)
}

// RetrieveCurrent mocks base method
func (m *MockBookingsClient) RetrieveCurrent(ctx context.Context, in *pbcheckout.RetrieveCurrentRequest, opts ...grpc.CallOption) (*pbcheckout.RetrieveBookingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveCurrent", varargs...)
	ret0, _ := ret[0].(*pbcheckout.RetrieveBookingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveCurrent indicates an expected call of RetrieveCurrent
func (mr *MockBookingsClientMockRecorder) RetrieveCurrent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveCurrent", reflect.TypeOf((*MockBookingsClient)(nil).RetrieveCurrent), varargs...)
}

// StateMachine mocks base method
func (m *MockBookingsClient) StateMachine(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*pbcheckout.StateMachineResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StateMachine", varargs...)
	ret0, _ := ret[0].(*pbcheckout.StateMachineResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMachine indicates an expected call of StateMachine
func (mr *MockBookingsClientMockRecorder) StateMachine(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMachine", reflect.TypeOf((*MockBookingsClient)(nil).StateMachine), varargs...)
}

// Prev mocks base method
func (m *MockBookingsClient) Prev(ctx context.Context, in *pbcheckout.UpdateOrderStateRequest, opts ...grpc.CallOption) (*pbcheckout.RetrieveBookingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Prev", varargs...)
	ret0, _ := ret[0].(*pbcheckout.RetrieveBookingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prev indicates an expected call of Prev
func (mr *MockBookingsClientMockRecorder) Prev(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prev", reflect.TypeOf((*MockBookingsClient)(nil).Prev), varargs...)
}

// Next mocks base method
func (m *MockBookingsClient) Next(ctx context.Context, in *pbcheckout.UpdateOrderStateRequest, opts ...grpc.CallOption) (*pbcheckout.RetrieveBookingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Next", varargs...)
	ret0, _ := ret[0].(*pbcheckout.RetrieveBookingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockBookingsClientMockRecorder) Next(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockBookingsClient)(nil).Next), varargs...)
}

// Cancel mocks base method
func (m *MockBookingsClient) Cancel(ctx context.Context, in *pbcheckout.UpdateOrderStateRequest, opts ...grpc.CallOption) (*pbcheckout.RetrieveBookingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Cancel", varargs...)
	ret0, _ := ret[0].(*pbcheckout.RetrieveBookingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cancel indicates an expected call of Cancel
func (mr *MockBookingsClientMockRecorder) Cancel(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockBookingsClient)(nil).Cancel), varargs...)
}

// MockBookingsServer is a mock of BookingsServer interface
type MockBookingsServer struct {
	ctrl     *gomock.Controller
	recorder *MockBookingsServerMockRecorder
}

// MockBookingsServerMockRecorder is the mock recorder for MockBookingsServer
type MockBookingsServerMockRecorder struct {
	mock *MockBookingsServer
}

// NewMockBookingsServer creates a new mock instance
func NewMockBookingsServer(ctrl *gomock.Controller) *MockBookingsServer {
	mock := &MockBookingsServer{ctrl: ctrl}
	mock.recorder = &MockBookingsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBookingsServer) EXPECT() *MockBookingsServerMockRecorder {
	return m.recorder
}

// CreateBooking mocks base method
func (m *MockBookingsServer) CreateBooking(arg0 context.Context, arg1 *pbcheckout.CreateBookingRequest) (*pbcheckout.CreateBookingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBooking", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.CreateBookingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBooking indicates an expected call of CreateBooking
func (mr *MockBookingsServerMockRecorder) CreateBooking(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBooking", reflect.TypeOf((*MockBookingsServer)(nil).CreateBooking), arg0, arg1)
}

// RetrieveCurrent mocks base method
func (m *MockBookingsServer) RetrieveCurrent(arg0 context.Context, arg1 *pbcheckout.RetrieveCurrentRequest) (*pbcheckout.RetrieveBookingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveCurrent", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.RetrieveBookingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveCurrent indicates an expected call of RetrieveCurrent
func (mr *MockBookingsServerMockRecorder) RetrieveCurrent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveCurrent", reflect.TypeOf((*MockBookingsServer)(nil).RetrieveCurrent), arg0, arg1)
}

// StateMachine mocks base method
func (m *MockBookingsServer) StateMachine(arg0 context.Context, arg1 *empty.Empty) (*pbcheckout.StateMachineResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMachine", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.StateMachineResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMachine indicates an expected call of StateMachine
func (mr *MockBookingsServerMockRecorder) StateMachine(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMachine", reflect.TypeOf((*MockBookingsServer)(nil).StateMachine), arg0, arg1)
}

// Prev mocks base method
func (m *MockBookingsServer) Prev(arg0 context.Context, arg1 *pbcheckout.UpdateOrderStateRequest) (*pbcheckout.RetrieveBookingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prev", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.RetrieveBookingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prev indicates an expected call of Prev
func (mr *MockBookingsServerMockRecorder) Prev(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prev", reflect.TypeOf((*MockBookingsServer)(nil).Prev), arg0, arg1)
}

// Next mocks base method
func (m *MockBookingsServer) Next(arg0 context.Context, arg1 *pbcheckout.UpdateOrderStateRequest) (*pbcheckout.RetrieveBookingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.RetrieveBookingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockBookingsServerMockRecorder) Next(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockBookingsServer)(nil).Next), arg0, arg1)
}

// Cancel mocks base method
func (m *MockBookingsServer) Cancel(arg0 context.Context, arg1 *pbcheckout.UpdateOrderStateRequest) (*pbcheckout.RetrieveBookingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.RetrieveBookingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cancel indicates an expected call of Cancel
func (mr *MockBookingsServerMockRecorder) Cancel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockBookingsServer)(nil).Cancel), arg0, arg1)
}

// MockPurchasesClient is a mock of PurchasesClient interface
type MockPurchasesClient struct {
	ctrl     *gomock.Controller
	recorder *MockPurchasesClientMockRecorder
}

// MockPurchasesClientMockRecorder is the mock recorder for MockPurchasesClient
type MockPurchasesClientMockRecorder struct {
	mock *MockPurchasesClient
}

// NewMockPurchasesClient creates a new mock instance
func NewMockPurchasesClient(ctrl *gomock.Controller) *MockPurchasesClient {
	mock := &MockPurchasesClient{ctrl: ctrl}
	mock.recorder = &MockPurchasesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPurchasesClient) EXPECT() *MockPurchasesClientMockRecorder {
	return m.recorder
}

// CreatePurchase mocks base method
func (m *MockPurchasesClient) CreatePurchase(ctx context.Context, in *pbcheckout.CreatePurchaseRequest, opts ...grpc.CallOption) (*pbcheckout.CreatePurchaseResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreatePurchase", varargs...)
	ret0, _ := ret[0].(*pbcheckout.CreatePurchaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePurchase indicates an expected call of CreatePurchase
func (mr *MockPurchasesClientMockRecorder) CreatePurchase(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePurchase", reflect.TypeOf((*MockPurchasesClient)(nil).CreatePurchase), varargs...)
}

// AssignAddress mocks base method
func (m *MockPurchasesClient) AssignAddress(ctx context.Context, in *pbcheckout.AssignAddressRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssignAddress", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignAddress indicates an expected call of AssignAddress
func (mr *MockPurchasesClientMockRecorder) AssignAddress(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignAddress", reflect.TypeOf((*MockPurchasesClient)(nil).AssignAddress), varargs...)
}

// RetrieveCurrent mocks base method
func (m *MockPurchasesClient) RetrieveCurrent(ctx context.Context, in *pbcheckout.RetrieveCurrentRequest, opts ...grpc.CallOption) (*pbcheckout.RetrievePurchaseResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveCurrent", varargs...)
	ret0, _ := ret[0].(*pbcheckout.RetrievePurchaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveCurrent indicates an expected call of RetrieveCurrent
func (mr *MockPurchasesClientMockRecorder) RetrieveCurrent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveCurrent", reflect.TypeOf((*MockPurchasesClient)(nil).RetrieveCurrent), varargs...)
}

// StateMachine mocks base method
func (m *MockPurchasesClient) StateMachine(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*pbcheckout.StateMachineResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StateMachine", varargs...)
	ret0, _ := ret[0].(*pbcheckout.StateMachineResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMachine indicates an expected call of StateMachine
func (mr *MockPurchasesClientMockRecorder) StateMachine(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMachine", reflect.TypeOf((*MockPurchasesClient)(nil).StateMachine), varargs...)
}

// RetrieveShippingMethods mocks base method
func (m *MockPurchasesClient) RetrieveShippingMethods(ctx context.Context, in *pbcheckout.RetrieveShippingMethodsRequest, opts ...grpc.CallOption) (*pbcheckout.RetrieveShippingMethodsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveShippingMethods", varargs...)
	ret0, _ := ret[0].(*pbcheckout.RetrieveShippingMethodsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveShippingMethods indicates an expected call of RetrieveShippingMethods
func (mr *MockPurchasesClientMockRecorder) RetrieveShippingMethods(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveShippingMethods", reflect.TypeOf((*MockPurchasesClient)(nil).RetrieveShippingMethods), varargs...)
}

// Prev mocks base method
func (m *MockPurchasesClient) Prev(ctx context.Context, in *pbcheckout.UpdateOrderStateRequest, opts ...grpc.CallOption) (*pbcheckout.RetrievePurchaseResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Prev", varargs...)
	ret0, _ := ret[0].(*pbcheckout.RetrievePurchaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prev indicates an expected call of Prev
func (mr *MockPurchasesClientMockRecorder) Prev(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prev", reflect.TypeOf((*MockPurchasesClient)(nil).Prev), varargs...)
}

// Next mocks base method
func (m *MockPurchasesClient) Next(ctx context.Context, in *pbcheckout.UpdateOrderStateRequest, opts ...grpc.CallOption) (*pbcheckout.RetrievePurchaseResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Next", varargs...)
	ret0, _ := ret[0].(*pbcheckout.RetrievePurchaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockPurchasesClientMockRecorder) Next(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockPurchasesClient)(nil).Next), varargs...)
}

// Cancel mocks base method
func (m *MockPurchasesClient) Cancel(ctx context.Context, in *pbcheckout.UpdateOrderStateRequest, opts ...grpc.CallOption) (*pbcheckout.RetrievePurchaseResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Cancel", varargs...)
	ret0, _ := ret[0].(*pbcheckout.RetrievePurchaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cancel indicates an expected call of Cancel
func (mr *MockPurchasesClientMockRecorder) Cancel(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockPurchasesClient)(nil).Cancel), varargs...)
}

// MockPurchasesServer is a mock of PurchasesServer interface
type MockPurchasesServer struct {
	ctrl     *gomock.Controller
	recorder *MockPurchasesServerMockRecorder
}

// MockPurchasesServerMockRecorder is the mock recorder for MockPurchasesServer
type MockPurchasesServerMockRecorder struct {
	mock *MockPurchasesServer
}

// NewMockPurchasesServer creates a new mock instance
func NewMockPurchasesServer(ctrl *gomock.Controller) *MockPurchasesServer {
	mock := &MockPurchasesServer{ctrl: ctrl}
	mock.recorder = &MockPurchasesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPurchasesServer) EXPECT() *MockPurchasesServerMockRecorder {
	return m.recorder
}

// CreatePurchase mocks base method
func (m *MockPurchasesServer) CreatePurchase(arg0 context.Context, arg1 *pbcheckout.CreatePurchaseRequest) (*pbcheckout.CreatePurchaseResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePurchase", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.CreatePurchaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePurchase indicates an expected call of CreatePurchase
func (mr *MockPurchasesServerMockRecorder) CreatePurchase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePurchase", reflect.TypeOf((*MockPurchasesServer)(nil).CreatePurchase), arg0, arg1)
}

// AssignAddress mocks base method
func (m *MockPurchasesServer) AssignAddress(arg0 context.Context, arg1 *pbcheckout.AssignAddressRequest) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignAddress", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignAddress indicates an expected call of AssignAddress
func (mr *MockPurchasesServerMockRecorder) AssignAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignAddress", reflect.TypeOf((*MockPurchasesServer)(nil).AssignAddress), arg0, arg1)
}

// RetrieveCurrent mocks base method
func (m *MockPurchasesServer) RetrieveCurrent(arg0 context.Context, arg1 *pbcheckout.RetrieveCurrentRequest) (*pbcheckout.RetrievePurchaseResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveCurrent", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.RetrievePurchaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveCurrent indicates an expected call of RetrieveCurrent
func (mr *MockPurchasesServerMockRecorder) RetrieveCurrent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveCurrent", reflect.TypeOf((*MockPurchasesServer)(nil).RetrieveCurrent), arg0, arg1)
}

// StateMachine mocks base method
func (m *MockPurchasesServer) StateMachine(arg0 context.Context, arg1 *empty.Empty) (*pbcheckout.StateMachineResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMachine", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.StateMachineResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMachine indicates an expected call of StateMachine
func (mr *MockPurchasesServerMockRecorder) StateMachine(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMachine", reflect.TypeOf((*MockPurchasesServer)(nil).StateMachine), arg0, arg1)
}

// RetrieveShippingMethods mocks base method
func (m *MockPurchasesServer) RetrieveShippingMethods(arg0 context.Context, arg1 *pbcheckout.RetrieveShippingMethodsRequest) (*pbcheckout.RetrieveShippingMethodsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveShippingMethods", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.RetrieveShippingMethodsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveShippingMethods indicates an expected call of RetrieveShippingMethods
func (mr *MockPurchasesServerMockRecorder) RetrieveShippingMethods(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveShippingMethods", reflect.TypeOf((*MockPurchasesServer)(nil).RetrieveShippingMethods), arg0, arg1)
}

// Prev mocks base method
func (m *MockPurchasesServer) Prev(arg0 context.Context, arg1 *pbcheckout.UpdateOrderStateRequest) (*pbcheckout.RetrievePurchaseResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prev", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.RetrievePurchaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prev indicates an expected call of Prev
func (mr *MockPurchasesServerMockRecorder) Prev(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prev", reflect.TypeOf((*MockPurchasesServer)(nil).Prev), arg0, arg1)
}

// Next mocks base method
func (m *MockPurchasesServer) Next(arg0 context.Context, arg1 *pbcheckout.UpdateOrderStateRequest) (*pbcheckout.RetrievePurchaseResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.RetrievePurchaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockPurchasesServerMockRecorder) Next(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockPurchasesServer)(nil).Next), arg0, arg1)
}

// Cancel mocks base method
func (m *MockPurchasesServer) Cancel(arg0 context.Context, arg1 *pbcheckout.UpdateOrderStateRequest) (*pbcheckout.RetrievePurchaseResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", arg0, arg1)
	ret0, _ := ret[0].(*pbcheckout.RetrievePurchaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cancel indicates an expected call of Cancel
func (mr *MockPurchasesServerMockRecorder) Cancel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockPurchasesServer)(nil).Cancel), arg0, arg1)
}
