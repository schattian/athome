// Code generated by protoc-gen-go. DO NOT EDIT.
// source: semantic.proto

package pbsemantic

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Category struct {
	Id                   uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ParentId             uint64      `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Childs               []*Category `protobuf:"bytes,4,rep,name=childs,proto3" json:"childs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_32e5db2b9bc8653e, []int{0}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetParentId() uint64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *Category) GetChilds() []*Category {
	if m != nil {
		return m.Childs
	}
	return nil
}

type Attribute struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CategoryId           uint64   `protobuf:"varint,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ValueType            string   `protobuf:"bytes,4,opt,name=value_type,json=valueType,proto3" json:"value_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attribute) Reset()         { *m = Attribute{} }
func (m *Attribute) String() string { return proto.CompactTextString(m) }
func (*Attribute) ProtoMessage()    {}
func (*Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_32e5db2b9bc8653e, []int{1}
}

func (m *Attribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attribute.Unmarshal(m, b)
}
func (m *Attribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attribute.Marshal(b, m, deterministic)
}
func (m *Attribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attribute.Merge(m, src)
}
func (m *Attribute) XXX_Size() int {
	return xxx_messageInfo_Attribute.Size(m)
}
func (m *Attribute) XXX_DiscardUnknown() {
	xxx_messageInfo_Attribute.DiscardUnknown(m)
}

var xxx_messageInfo_Attribute proto.InternalMessageInfo

func (m *Attribute) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Attribute) GetCategoryId() uint64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *Attribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Attribute) GetValueType() string {
	if m != nil {
		return m.ValueType
	}
	return ""
}

type GetCategoriesResponse struct {
	Categories           []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetCategoriesResponse) Reset()         { *m = GetCategoriesResponse{} }
func (m *GetCategoriesResponse) String() string { return proto.CompactTextString(m) }
func (*GetCategoriesResponse) ProtoMessage()    {}
func (*GetCategoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32e5db2b9bc8653e, []int{2}
}

func (m *GetCategoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCategoriesResponse.Unmarshal(m, b)
}
func (m *GetCategoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCategoriesResponse.Marshal(b, m, deterministic)
}
func (m *GetCategoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCategoriesResponse.Merge(m, src)
}
func (m *GetCategoriesResponse) XXX_Size() int {
	return xxx_messageInfo_GetCategoriesResponse.Size(m)
}
func (m *GetCategoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCategoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCategoriesResponse proto.InternalMessageInfo

func (m *GetCategoriesResponse) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type GetAttributesRequest struct {
	CategoryId           uint64   `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAttributesRequest) Reset()         { *m = GetAttributesRequest{} }
func (m *GetAttributesRequest) String() string { return proto.CompactTextString(m) }
func (*GetAttributesRequest) ProtoMessage()    {}
func (*GetAttributesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32e5db2b9bc8653e, []int{3}
}

func (m *GetAttributesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAttributesRequest.Unmarshal(m, b)
}
func (m *GetAttributesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAttributesRequest.Marshal(b, m, deterministic)
}
func (m *GetAttributesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAttributesRequest.Merge(m, src)
}
func (m *GetAttributesRequest) XXX_Size() int {
	return xxx_messageInfo_GetAttributesRequest.Size(m)
}
func (m *GetAttributesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAttributesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAttributesRequest proto.InternalMessageInfo

func (m *GetAttributesRequest) GetCategoryId() uint64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type GetAttributesResponse struct {
	Attributes           []*Attribute `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetAttributesResponse) Reset()         { *m = GetAttributesResponse{} }
func (m *GetAttributesResponse) String() string { return proto.CompactTextString(m) }
func (*GetAttributesResponse) ProtoMessage()    {}
func (*GetAttributesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32e5db2b9bc8653e, []int{4}
}

func (m *GetAttributesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAttributesResponse.Unmarshal(m, b)
}
func (m *GetAttributesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAttributesResponse.Marshal(b, m, deterministic)
}
func (m *GetAttributesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAttributesResponse.Merge(m, src)
}
func (m *GetAttributesResponse) XXX_Size() int {
	return xxx_messageInfo_GetAttributesResponse.Size(m)
}
func (m *GetAttributesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAttributesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAttributesResponse proto.InternalMessageInfo

func (m *GetAttributesResponse) GetAttributes() []*Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func init() {
	proto.RegisterType((*Category)(nil), "semantic.Category")
	proto.RegisterType((*Attribute)(nil), "semantic.Attribute")
	proto.RegisterType((*GetCategoriesResponse)(nil), "semantic.GetCategoriesResponse")
	proto.RegisterType((*GetAttributesRequest)(nil), "semantic.GetAttributesRequest")
	proto.RegisterType((*GetAttributesResponse)(nil), "semantic.GetAttributesResponse")
}

func init() {
	proto.RegisterFile("semantic.proto", fileDescriptor_32e5db2b9bc8653e)
}

var fileDescriptor_32e5db2b9bc8653e = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0xbb, 0x49, 0x54, 0xc5, 0x53, 0x51, 0xa1, 0xe5, 0xcb, 0x4a, 0x05, 0x8d, 0x7c, 0x8a,
	0x38, 0xd8, 0x92, 0x8b, 0xc4, 0x19, 0x10, 0xaa, 0xca, 0x87, 0x54, 0x19, 0x4e, 0x5c, 0xaa, 0xf5,
	0xee, 0x90, 0xac, 0x88, 0xbd, 0x66, 0x77, 0x1c, 0x94, 0xbf, 0x04, 0x37, 0x7e, 0x21, 0xea, 0xc6,
	0x9b, 0xa4, 0x09, 0xbd, 0xe7, 0x36, 0x3b, 0xf3, 0xce, 0x3b, 0xcf, 0xac, 0x06, 0x4e, 0x1d, 0x56,
	0xa2, 0x26, 0x2d, 0xd3, 0xc6, 0x1a, 0x32, 0x7c, 0x18, 0xde, 0xa3, 0x67, 0x0b, 0x31, 0xd7, 0x4a,
	0x10, 0x66, 0x21, 0x58, 0x49, 0x46, 0x67, 0x53, 0x63, 0xa6, 0x73, 0xcc, 0xfc, 0xab, 0x6c, 0xbf,
	0x67, 0x58, 0x35, 0xb4, 0x5c, 0x15, 0x93, 0x5f, 0x30, 0x7c, 0x27, 0x08, 0xa7, 0xc6, 0x2e, 0xf9,
	0x29, 0xf4, 0xb4, 0x8a, 0xd9, 0x98, 0x4d, 0x06, 0x45, 0x4f, 0x2b, 0xce, 0x61, 0x50, 0x8b, 0x0a,
	0xe3, 0xde, 0x98, 0x4d, 0xa2, 0xc2, 0xc7, 0xfc, 0x0c, 0xa2, 0x46, 0x58, 0xac, 0xe9, 0x46, 0xab,
	0xb8, 0xef, 0xa5, 0xc3, 0x55, 0xe2, 0x4a, 0xf1, 0x97, 0x70, 0x2c, 0x67, 0x7a, 0xae, 0x5c, 0x3c,
	0x18, 0xf7, 0x27, 0x27, 0x39, 0x4f, 0xd7, 0xb4, 0x61, 0x48, 0xd1, 0x29, 0x12, 0x03, 0xd1, 0x1b,
	0x22, 0xab, 0xcb, 0x96, 0x70, 0x6f, 0xf2, 0x39, 0x9c, 0xc8, 0xae, 0xe1, 0x76, 0x4e, 0xcf, 0x17,
	0x20, 0xa4, 0xae, 0x36, 0x68, 0xfd, 0x2d, 0xb4, 0xe7, 0x00, 0x0b, 0x31, 0x6f, 0xf1, 0x86, 0x96,
	0x0d, 0xc6, 0x03, 0x5f, 0x89, 0x7c, 0xe6, 0xeb, 0xb2, 0xc1, 0xe4, 0x23, 0x3c, 0xb9, 0x44, 0xea,
	0x38, 0x34, 0xba, 0x02, 0x5d, 0x63, 0x6a, 0x87, 0x3c, 0x87, 0xe0, 0xac, 0xd1, 0xc5, 0xec, 0x5e,
	0xf2, 0x2d, 0x55, 0xf2, 0x1a, 0x1e, 0x5f, 0x22, 0xad, 0x17, 0x70, 0x05, 0xfe, 0x6c, 0xd1, 0xd1,
	0x2e, 0x38, 0xdb, 0x05, 0x4f, 0x3e, 0x79, 0x8a, 0xed, 0xc6, 0x8e, 0xe2, 0x02, 0x40, 0xac, 0xb3,
	0x1d, 0xc5, 0xa3, 0x0d, 0xc5, 0xba, 0xa3, 0xd8, 0x92, 0xe5, 0x7f, 0x19, 0x3c, 0xfc, 0x82, 0x76,
	0xa1, 0x25, 0x5e, 0x5b, 0xb3, 0xd0, 0x0a, 0xad, 0xe3, 0x1f, 0xe0, 0xc1, 0x9d, 0x45, 0xf9, 0xd3,
	0x74, 0x75, 0x01, 0x69, 0xb8, 0x80, 0xf4, 0xfd, 0xed, 0x05, 0x8c, 0xce, 0x37, 0xf6, 0xff, 0xfd,
	0x99, 0xe4, 0x88, 0x17, 0xde, 0x6b, 0x83, 0xcb, 0x5f, 0xdc, 0xe9, 0xd9, 0xfb, 0x80, 0x1d, 0xcf,
	0xfd, 0x3d, 0x93, 0xa3, 0xfc, 0x0f, 0x83, 0xe8, 0x33, 0x5a, 0x39, 0x13, 0x35, 0x1d, 0x3e, 0xed,
	0x6f, 0x06, 0xc3, 0x6b, 0x6b, 0x54, 0x2b, 0x0f, 0x1f, 0xf6, 0xed, 0xab, 0x6f, 0xf9, 0x54, 0xd3,
	0xac, 0x2d, 0x53, 0x69, 0xaa, 0x4c, 0xd0, 0xcc, 0x54, 0x28, 0x4d, 0x25, 0x6c, 0x56, 0x0a, 0xf9,
	0x03, 0x6b, 0x95, 0x05, 0x8b, 0xac, 0x29, 0x43, 0x58, 0x1e, 0x7b, 0xf8, 0x8b, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x61, 0xf3, 0x1d, 0x98, 0x5c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceProvidersClient is the client API for ServiceProviders service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceProvidersClient interface {
	GetCategories(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
	GetAttributes(ctx context.Context, in *GetAttributesRequest, opts ...grpc.CallOption) (*GetAttributesResponse, error)
}

type serviceProvidersClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceProvidersClient(cc grpc.ClientConnInterface) ServiceProvidersClient {
	return &serviceProvidersClient{cc}
}

func (c *serviceProvidersClient) GetCategories(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, "/semantic.ServiceProviders/GetCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProvidersClient) GetAttributes(ctx context.Context, in *GetAttributesRequest, opts ...grpc.CallOption) (*GetAttributesResponse, error) {
	out := new(GetAttributesResponse)
	err := c.cc.Invoke(ctx, "/semantic.ServiceProviders/GetAttributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceProvidersServer is the server API for ServiceProviders service.
type ServiceProvidersServer interface {
	GetCategories(context.Context, *empty.Empty) (*GetCategoriesResponse, error)
	GetAttributes(context.Context, *GetAttributesRequest) (*GetAttributesResponse, error)
}

// UnimplementedServiceProvidersServer can be embedded to have forward compatible implementations.
type UnimplementedServiceProvidersServer struct {
}

func (*UnimplementedServiceProvidersServer) GetCategories(ctx context.Context, req *empty.Empty) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (*UnimplementedServiceProvidersServer) GetAttributes(ctx context.Context, req *GetAttributesRequest) (*GetAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttributes not implemented")
}

func RegisterServiceProvidersServer(s *grpc.Server, srv ServiceProvidersServer) {
	s.RegisterService(&_ServiceProviders_serviceDesc, srv)
}

func _ServiceProviders_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProvidersServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/semantic.ServiceProviders/GetCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProvidersServer).GetCategories(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviders_GetAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProvidersServer).GetAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/semantic.ServiceProviders/GetAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProvidersServer).GetAttributes(ctx, req.(*GetAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceProviders_serviceDesc = grpc.ServiceDesc{
	ServiceName: "semantic.ServiceProviders",
	HandlerType: (*ServiceProvidersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategories",
			Handler:    _ServiceProviders_GetCategories_Handler,
		},
		{
			MethodName: "GetAttributes",
			Handler:    _ServiceProviders_GetAttributes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "semantic.proto",
}

// MerchantsClient is the client API for Merchants service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MerchantsClient interface {
	GetCategories(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
	GetAttributes(ctx context.Context, in *GetAttributesRequest, opts ...grpc.CallOption) (*GetAttributesResponse, error)
}

type merchantsClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantsClient(cc grpc.ClientConnInterface) MerchantsClient {
	return &merchantsClient{cc}
}

func (c *merchantsClient) GetCategories(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, "/semantic.Merchants/GetCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantsClient) GetAttributes(ctx context.Context, in *GetAttributesRequest, opts ...grpc.CallOption) (*GetAttributesResponse, error) {
	out := new(GetAttributesResponse)
	err := c.cc.Invoke(ctx, "/semantic.Merchants/GetAttributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantsServer is the server API for Merchants service.
type MerchantsServer interface {
	GetCategories(context.Context, *empty.Empty) (*GetCategoriesResponse, error)
	GetAttributes(context.Context, *GetAttributesRequest) (*GetAttributesResponse, error)
}

// UnimplementedMerchantsServer can be embedded to have forward compatible implementations.
type UnimplementedMerchantsServer struct {
}

func (*UnimplementedMerchantsServer) GetCategories(ctx context.Context, req *empty.Empty) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (*UnimplementedMerchantsServer) GetAttributes(ctx context.Context, req *GetAttributesRequest) (*GetAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttributes not implemented")
}

func RegisterMerchantsServer(s *grpc.Server, srv MerchantsServer) {
	s.RegisterService(&_Merchants_serviceDesc, srv)
}

func _Merchants_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantsServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/semantic.Merchants/GetCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantsServer).GetCategories(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merchants_GetAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantsServer).GetAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/semantic.Merchants/GetAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantsServer).GetAttributes(ctx, req.(*GetAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Merchants_serviceDesc = grpc.ServiceDesc{
	ServiceName: "semantic.Merchants",
	HandlerType: (*MerchantsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategories",
			Handler:    _Merchants_GetCategories_Handler,
		},
		{
			MethodName: "GetAttributes",
			Handler:    _Merchants_GetAttributes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "semantic.proto",
}

// ProductsClient is the client API for Products service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductsClient interface {
	GetCategories(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
	GetAttributes(ctx context.Context, in *GetAttributesRequest, opts ...grpc.CallOption) (*GetAttributesResponse, error)
}

type productsClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsClient(cc grpc.ClientConnInterface) ProductsClient {
	return &productsClient{cc}
}

func (c *productsClient) GetCategories(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, "/semantic.Products/GetCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetAttributes(ctx context.Context, in *GetAttributesRequest, opts ...grpc.CallOption) (*GetAttributesResponse, error) {
	out := new(GetAttributesResponse)
	err := c.cc.Invoke(ctx, "/semantic.Products/GetAttributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsServer is the server API for Products service.
type ProductsServer interface {
	GetCategories(context.Context, *empty.Empty) (*GetCategoriesResponse, error)
	GetAttributes(context.Context, *GetAttributesRequest) (*GetAttributesResponse, error)
}

// UnimplementedProductsServer can be embedded to have forward compatible implementations.
type UnimplementedProductsServer struct {
}

func (*UnimplementedProductsServer) GetCategories(ctx context.Context, req *empty.Empty) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (*UnimplementedProductsServer) GetAttributes(ctx context.Context, req *GetAttributesRequest) (*GetAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttributes not implemented")
}

func RegisterProductsServer(s *grpc.Server, srv ProductsServer) {
	s.RegisterService(&_Products_serviceDesc, srv)
}

func _Products_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/semantic.Products/GetCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetCategories(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/semantic.Products/GetAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetAttributes(ctx, req.(*GetAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Products_serviceDesc = grpc.ServiceDesc{
	ServiceName: "semantic.Products",
	HandlerType: (*ProductsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategories",
			Handler:    _Products_GetCategories_Handler,
		},
		{
			MethodName: "GetAttributes",
			Handler:    _Products_GetAttributes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "semantic.proto",
}
