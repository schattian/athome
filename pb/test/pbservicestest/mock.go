// Code generated by MockGen. DO NOT EDIT.
// Source: ../pbservices/services.pb.go

// Package pbservicestest is a generated GoMock package.
package pbservicestest

import (
	context "context"
	pbservices "github.com/athomecomar/athome/pb/pbservices"
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

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

// SearchServices mocks base method
func (m *MockViewerClient) SearchServices(ctx context.Context, in *pbservices.SearchServicesRequest, opts ...grpc.CallOption) (*pbservices.SearchServicesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchServices", varargs...)
	ret0, _ := ret[0].(*pbservices.SearchServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchServices indicates an expected call of SearchServices
func (mr *MockViewerClientMockRecorder) SearchServices(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchServices", reflect.TypeOf((*MockViewerClient)(nil).SearchServices), varargs...)
}

// SearchAvailableShippings mocks base method
func (m *MockViewerClient) SearchAvailableShippings(ctx context.Context, in *pbservices.SearchAvailableShippingsRequest, opts ...grpc.CallOption) (*pbservices.SearchAvailableShippingsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchAvailableShippings", varargs...)
	ret0, _ := ret[0].(*pbservices.SearchAvailableShippingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAvailableShippings indicates an expected call of SearchAvailableShippings
func (mr *MockViewerClientMockRecorder) SearchAvailableShippings(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAvailableShippings", reflect.TypeOf((*MockViewerClient)(nil).SearchAvailableShippings), varargs...)
}

// RetrieveServiceDetail mocks base method
func (m *MockViewerClient) RetrieveServiceDetail(ctx context.Context, in *pbservices.RetrieveServiceDetailRequest, opts ...grpc.CallOption) (*pbservices.ServiceDetail, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveServiceDetail", varargs...)
	ret0, _ := ret[0].(*pbservices.ServiceDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveServiceDetail indicates an expected call of RetrieveServiceDetail
func (mr *MockViewerClientMockRecorder) RetrieveServiceDetail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveServiceDetail", reflect.TypeOf((*MockViewerClient)(nil).RetrieveServiceDetail), varargs...)
}

// MockViewerServer is a mock of ViewerServer interface
type MockViewerServer struct {
	ctrl     *gomock.Controller
	recorder *MockViewerServerMockRecorder
}

// MockViewerServerMockRecorder is the mock recorder for MockViewerServer
type MockViewerServerMockRecorder struct {
	mock *MockViewerServer
}

// NewMockViewerServer creates a new mock instance
func NewMockViewerServer(ctrl *gomock.Controller) *MockViewerServer {
	mock := &MockViewerServer{ctrl: ctrl}
	mock.recorder = &MockViewerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockViewerServer) EXPECT() *MockViewerServerMockRecorder {
	return m.recorder
}

// SearchServices mocks base method
func (m *MockViewerServer) SearchServices(arg0 context.Context, arg1 *pbservices.SearchServicesRequest) (*pbservices.SearchServicesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchServices", arg0, arg1)
	ret0, _ := ret[0].(*pbservices.SearchServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchServices indicates an expected call of SearchServices
func (mr *MockViewerServerMockRecorder) SearchServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchServices", reflect.TypeOf((*MockViewerServer)(nil).SearchServices), arg0, arg1)
}

// SearchAvailableShippings mocks base method
func (m *MockViewerServer) SearchAvailableShippings(arg0 context.Context, arg1 *pbservices.SearchAvailableShippingsRequest) (*pbservices.SearchAvailableShippingsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchAvailableShippings", arg0, arg1)
	ret0, _ := ret[0].(*pbservices.SearchAvailableShippingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAvailableShippings indicates an expected call of SearchAvailableShippings
func (mr *MockViewerServerMockRecorder) SearchAvailableShippings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAvailableShippings", reflect.TypeOf((*MockViewerServer)(nil).SearchAvailableShippings), arg0, arg1)
}

// RetrieveServiceDetail mocks base method
func (m *MockViewerServer) RetrieveServiceDetail(arg0 context.Context, arg1 *pbservices.RetrieveServiceDetailRequest) (*pbservices.ServiceDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveServiceDetail", arg0, arg1)
	ret0, _ := ret[0].(*pbservices.ServiceDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveServiceDetail indicates an expected call of RetrieveServiceDetail
func (mr *MockViewerServerMockRecorder) RetrieveServiceDetail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveServiceDetail", reflect.TypeOf((*MockViewerServer)(nil).RetrieveServiceDetail), arg0, arg1)
}

// MockCalendarsClient is a mock of CalendarsClient interface
type MockCalendarsClient struct {
	ctrl     *gomock.Controller
	recorder *MockCalendarsClientMockRecorder
}

// MockCalendarsClientMockRecorder is the mock recorder for MockCalendarsClient
type MockCalendarsClientMockRecorder struct {
	mock *MockCalendarsClient
}

// NewMockCalendarsClient creates a new mock instance
func NewMockCalendarsClient(ctrl *gomock.Controller) *MockCalendarsClient {
	mock := &MockCalendarsClient{ctrl: ctrl}
	mock.recorder = &MockCalendarsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCalendarsClient) EXPECT() *MockCalendarsClientMockRecorder {
	return m.recorder
}

// CreateCalendar mocks base method
func (m *MockCalendarsClient) CreateCalendar(ctx context.Context, in *pbservices.CreateCalendarRequest, opts ...grpc.CallOption) (*pbservices.CreateCalendarResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCalendar", varargs...)
	ret0, _ := ret[0].(*pbservices.CreateCalendarResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCalendar indicates an expected call of CreateCalendar
func (mr *MockCalendarsClientMockRecorder) CreateCalendar(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCalendar", reflect.TypeOf((*MockCalendarsClient)(nil).CreateCalendar), varargs...)
}

// RetrieveCalendar mocks base method
func (m *MockCalendarsClient) RetrieveCalendar(ctx context.Context, in *pbservices.RetrieveCalendarRequest, opts ...grpc.CallOption) (*pbservices.CalendarDetail, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveCalendar", varargs...)
	ret0, _ := ret[0].(*pbservices.CalendarDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveCalendar indicates an expected call of RetrieveCalendar
func (mr *MockCalendarsClientMockRecorder) RetrieveCalendar(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveCalendar", reflect.TypeOf((*MockCalendarsClient)(nil).RetrieveCalendar), varargs...)
}

// RetrieveMyCalendars mocks base method
func (m *MockCalendarsClient) RetrieveMyCalendars(ctx context.Context, in *pbservices.RetrieveMyCalendarsRequest, opts ...grpc.CallOption) (*pbservices.RetrieveMyCalendarsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveMyCalendars", varargs...)
	ret0, _ := ret[0].(*pbservices.RetrieveMyCalendarsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveMyCalendars indicates an expected call of RetrieveMyCalendars
func (mr *MockCalendarsClientMockRecorder) RetrieveMyCalendars(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveMyCalendars", reflect.TypeOf((*MockCalendarsClient)(nil).RetrieveMyCalendars), varargs...)
}

// CreateEvent mocks base method
func (m *MockCalendarsClient) CreateEvent(ctx context.Context, in *pbservices.CreateEventRequest, opts ...grpc.CallOption) (*pbservices.CreateEventResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateEvent", varargs...)
	ret0, _ := ret[0].(*pbservices.CreateEventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent
func (mr *MockCalendarsClientMockRecorder) CreateEvent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockCalendarsClient)(nil).CreateEvent), varargs...)
}

// ConfirmEvent mocks base method
func (m *MockCalendarsClient) ConfirmEvent(ctx context.Context, in *pbservices.ConfirmEventRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConfirmEvent", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfirmEvent indicates an expected call of ConfirmEvent
func (mr *MockCalendarsClientMockRecorder) ConfirmEvent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmEvent", reflect.TypeOf((*MockCalendarsClient)(nil).ConfirmEvent), varargs...)
}

// MockCalendarsServer is a mock of CalendarsServer interface
type MockCalendarsServer struct {
	ctrl     *gomock.Controller
	recorder *MockCalendarsServerMockRecorder
}

// MockCalendarsServerMockRecorder is the mock recorder for MockCalendarsServer
type MockCalendarsServerMockRecorder struct {
	mock *MockCalendarsServer
}

// NewMockCalendarsServer creates a new mock instance
func NewMockCalendarsServer(ctrl *gomock.Controller) *MockCalendarsServer {
	mock := &MockCalendarsServer{ctrl: ctrl}
	mock.recorder = &MockCalendarsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCalendarsServer) EXPECT() *MockCalendarsServerMockRecorder {
	return m.recorder
}

// CreateCalendar mocks base method
func (m *MockCalendarsServer) CreateCalendar(arg0 context.Context, arg1 *pbservices.CreateCalendarRequest) (*pbservices.CreateCalendarResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCalendar", arg0, arg1)
	ret0, _ := ret[0].(*pbservices.CreateCalendarResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCalendar indicates an expected call of CreateCalendar
func (mr *MockCalendarsServerMockRecorder) CreateCalendar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCalendar", reflect.TypeOf((*MockCalendarsServer)(nil).CreateCalendar), arg0, arg1)
}

// RetrieveCalendar mocks base method
func (m *MockCalendarsServer) RetrieveCalendar(arg0 context.Context, arg1 *pbservices.RetrieveCalendarRequest) (*pbservices.CalendarDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveCalendar", arg0, arg1)
	ret0, _ := ret[0].(*pbservices.CalendarDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveCalendar indicates an expected call of RetrieveCalendar
func (mr *MockCalendarsServerMockRecorder) RetrieveCalendar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveCalendar", reflect.TypeOf((*MockCalendarsServer)(nil).RetrieveCalendar), arg0, arg1)
}

// RetrieveMyCalendars mocks base method
func (m *MockCalendarsServer) RetrieveMyCalendars(arg0 context.Context, arg1 *pbservices.RetrieveMyCalendarsRequest) (*pbservices.RetrieveMyCalendarsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveMyCalendars", arg0, arg1)
	ret0, _ := ret[0].(*pbservices.RetrieveMyCalendarsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveMyCalendars indicates an expected call of RetrieveMyCalendars
func (mr *MockCalendarsServerMockRecorder) RetrieveMyCalendars(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveMyCalendars", reflect.TypeOf((*MockCalendarsServer)(nil).RetrieveMyCalendars), arg0, arg1)
}

// CreateEvent mocks base method
func (m *MockCalendarsServer) CreateEvent(arg0 context.Context, arg1 *pbservices.CreateEventRequest) (*pbservices.CreateEventResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", arg0, arg1)
	ret0, _ := ret[0].(*pbservices.CreateEventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent
func (mr *MockCalendarsServerMockRecorder) CreateEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockCalendarsServer)(nil).CreateEvent), arg0, arg1)
}

// ConfirmEvent mocks base method
func (m *MockCalendarsServer) ConfirmEvent(arg0 context.Context, arg1 *pbservices.ConfirmEventRequest) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmEvent", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfirmEvent indicates an expected call of ConfirmEvent
func (mr *MockCalendarsServerMockRecorder) ConfirmEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmEvent", reflect.TypeOf((*MockCalendarsServer)(nil).ConfirmEvent), arg0, arg1)
}

// MockRegisterClient is a mock of RegisterClient interface
type MockRegisterClient struct {
	ctrl     *gomock.Controller
	recorder *MockRegisterClientMockRecorder
}

// MockRegisterClientMockRecorder is the mock recorder for MockRegisterClient
type MockRegisterClientMockRecorder struct {
	mock *MockRegisterClient
}

// NewMockRegisterClient creates a new mock instance
func NewMockRegisterClient(ctrl *gomock.Controller) *MockRegisterClient {
	mock := &MockRegisterClient{ctrl: ctrl}
	mock.recorder = &MockRegisterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegisterClient) EXPECT() *MockRegisterClientMockRecorder {
	return m.recorder
}

// RetrieveRegistry mocks base method
func (m *MockRegisterClient) RetrieveRegistry(ctx context.Context, in *pbservices.RetrieveRegistryRequest, opts ...grpc.CallOption) (*pbservices.RegistryDetail, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveRegistry", varargs...)
	ret0, _ := ret[0].(*pbservices.RegistryDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveRegistry indicates an expected call of RetrieveRegistry
func (mr *MockRegisterClientMockRecorder) RetrieveRegistry(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveRegistry", reflect.TypeOf((*MockRegisterClient)(nil).RetrieveRegistry), varargs...)
}

// DeleteRegistry mocks base method
func (m *MockRegisterClient) DeleteRegistry(ctx context.Context, in *pbservices.DeleteRegistryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteRegistry", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRegistry indicates an expected call of DeleteRegistry
func (mr *MockRegisterClientMockRecorder) DeleteRegistry(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRegistry", reflect.TypeOf((*MockRegisterClient)(nil).DeleteRegistry), varargs...)
}

// Prev mocks base method
func (m *MockRegisterClient) Prev(ctx context.Context, in *pbservices.PrevRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Prev", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prev indicates an expected call of Prev
func (mr *MockRegisterClientMockRecorder) Prev(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prev", reflect.TypeOf((*MockRegisterClient)(nil).Prev), varargs...)
}

// First mocks base method
func (m *MockRegisterClient) First(ctx context.Context, in *pbservices.FirstRequest, opts ...grpc.CallOption) (*pbservices.FirstResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(*pbservices.FirstResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// First indicates an expected call of First
func (mr *MockRegisterClientMockRecorder) First(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockRegisterClient)(nil).First), varargs...)
}

// Second mocks base method
func (m *MockRegisterClient) Second(ctx context.Context, in *pbservices.SecondRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Second", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Second indicates an expected call of Second
func (mr *MockRegisterClientMockRecorder) Second(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Second", reflect.TypeOf((*MockRegisterClient)(nil).Second), varargs...)
}

// Third mocks base method
func (m *MockRegisterClient) Third(ctx context.Context, in *pbservices.ThirdRequest, opts ...grpc.CallOption) (*pbservices.ThirdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Third", varargs...)
	ret0, _ := ret[0].(*pbservices.ThirdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Third indicates an expected call of Third
func (mr *MockRegisterClientMockRecorder) Third(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Third", reflect.TypeOf((*MockRegisterClient)(nil).Third), varargs...)
}

// MockRegisterServer is a mock of RegisterServer interface
type MockRegisterServer struct {
	ctrl     *gomock.Controller
	recorder *MockRegisterServerMockRecorder
}

// MockRegisterServerMockRecorder is the mock recorder for MockRegisterServer
type MockRegisterServerMockRecorder struct {
	mock *MockRegisterServer
}

// NewMockRegisterServer creates a new mock instance
func NewMockRegisterServer(ctrl *gomock.Controller) *MockRegisterServer {
	mock := &MockRegisterServer{ctrl: ctrl}
	mock.recorder = &MockRegisterServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegisterServer) EXPECT() *MockRegisterServerMockRecorder {
	return m.recorder
}

// RetrieveRegistry mocks base method
func (m *MockRegisterServer) RetrieveRegistry(arg0 context.Context, arg1 *pbservices.RetrieveRegistryRequest) (*pbservices.RegistryDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveRegistry", arg0, arg1)
	ret0, _ := ret[0].(*pbservices.RegistryDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveRegistry indicates an expected call of RetrieveRegistry
func (mr *MockRegisterServerMockRecorder) RetrieveRegistry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveRegistry", reflect.TypeOf((*MockRegisterServer)(nil).RetrieveRegistry), arg0, arg1)
}

// DeleteRegistry mocks base method
func (m *MockRegisterServer) DeleteRegistry(arg0 context.Context, arg1 *pbservices.DeleteRegistryRequest) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRegistry", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRegistry indicates an expected call of DeleteRegistry
func (mr *MockRegisterServerMockRecorder) DeleteRegistry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRegistry", reflect.TypeOf((*MockRegisterServer)(nil).DeleteRegistry), arg0, arg1)
}

// Prev mocks base method
func (m *MockRegisterServer) Prev(arg0 context.Context, arg1 *pbservices.PrevRequest) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prev", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prev indicates an expected call of Prev
func (mr *MockRegisterServerMockRecorder) Prev(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prev", reflect.TypeOf((*MockRegisterServer)(nil).Prev), arg0, arg1)
}

// First mocks base method
func (m *MockRegisterServer) First(arg0 context.Context, arg1 *pbservices.FirstRequest) (*pbservices.FirstResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "First", arg0, arg1)
	ret0, _ := ret[0].(*pbservices.FirstResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// First indicates an expected call of First
func (mr *MockRegisterServerMockRecorder) First(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockRegisterServer)(nil).First), arg0, arg1)
}

// Second mocks base method
func (m *MockRegisterServer) Second(arg0 context.Context, arg1 *pbservices.SecondRequest) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Second", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Second indicates an expected call of Second
func (mr *MockRegisterServerMockRecorder) Second(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Second", reflect.TypeOf((*MockRegisterServer)(nil).Second), arg0, arg1)
}

// Third mocks base method
func (m *MockRegisterServer) Third(arg0 context.Context, arg1 *pbservices.ThirdRequest) (*pbservices.ThirdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Third", arg0, arg1)
	ret0, _ := ret[0].(*pbservices.ThirdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Third indicates an expected call of Third
func (mr *MockRegisterServerMockRecorder) Third(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Third", reflect.TypeOf((*MockRegisterServer)(nil).Third), arg0, arg1)
}
