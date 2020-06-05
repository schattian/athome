// Code generated by protoc-gen-go. DO NOT EDIT.
// source: products.proto

package pbproducts

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type FirstRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// replace the another for that
	CategoryId           uint64   `protobuf:"varint,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FirstRequest) Reset()         { *m = FirstRequest{} }
func (m *FirstRequest) String() string { return proto.CompactTextString(m) }
func (*FirstRequest) ProtoMessage()    {}
func (*FirstRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6e54f42122eb82, []int{0}
}

func (m *FirstRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FirstRequest.Unmarshal(m, b)
}
func (m *FirstRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FirstRequest.Marshal(b, m, deterministic)
}
func (m *FirstRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FirstRequest.Merge(m, src)
}
func (m *FirstRequest) XXX_Size() int {
	return xxx_messageInfo_FirstRequest.Size(m)
}
func (m *FirstRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FirstRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FirstRequest proto.InternalMessageInfo

func (m *FirstRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *FirstRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *FirstRequest) GetCategoryId() uint64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type FirstResponse struct {
	DraftId              uint64   `protobuf:"varint,1,opt,name=draft_id,json=draftId,proto3" json:"draft_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FirstResponse) Reset()         { *m = FirstResponse{} }
func (m *FirstResponse) String() string { return proto.CompactTextString(m) }
func (*FirstResponse) ProtoMessage()    {}
func (*FirstResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6e54f42122eb82, []int{1}
}

func (m *FirstResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FirstResponse.Unmarshal(m, b)
}
func (m *FirstResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FirstResponse.Marshal(b, m, deterministic)
}
func (m *FirstResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FirstResponse.Merge(m, src)
}
func (m *FirstResponse) XXX_Size() int {
	return xxx_messageInfo_FirstResponse.Size(m)
}
func (m *FirstResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FirstResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FirstResponse proto.InternalMessageInfo

func (m *FirstResponse) GetDraftId() uint64 {
	if m != nil {
		return m.DraftId
	}
	return 0
}

type Page struct {
	Cursor               uint64   `protobuf:"varint,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Size                 uint64   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6e54f42122eb82, []int{2}
}

func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (m *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(m, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetCursor() uint64 {
	if m != nil {
		return m.Cursor
	}
	return 0
}

func (m *Page) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*FirstRequest)(nil), "products.FirstRequest")
	proto.RegisterType((*FirstResponse)(nil), "products.FirstResponse")
	proto.RegisterType((*Page)(nil), "products.Page")
}

func init() {
	proto.RegisterFile("products.proto", fileDescriptor_8c6e54f42122eb82)
}

var fileDescriptor_8c6e54f42122eb82 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x8d, 0x6e, 0x7f, 0x38, 0xad, 0x3d, 0x0c, 0xd2, 0xae, 0xbd, 0x58, 0xf7, 0xb4, 0x78,
	0xd8, 0x4a, 0xbd, 0x7b, 0x50, 0x10, 0xf6, 0x26, 0x8b, 0x27, 0x2f, 0x25, 0x4d, 0xc6, 0x1a, 0x2c,
	0xcd, 0x9a, 0xcc, 0x0a, 0xfa, 0xd7, 0x4b, 0x93, 0x5d, 0x11, 0xbc, 0xcd, 0xfb, 0xf2, 0xc1, 0x4c,
	0x1e, 0x4c, 0x6a, 0x67, 0x75, 0xa3, 0xd8, 0x17, 0xb5, 0xb3, 0x6c, 0x71, 0xd8, 0xe5, 0xf9, 0xec,
	0x53, 0xee, 0x8c, 0x96, 0x4c, 0xcb, 0x6e, 0x88, 0x4a, 0xf6, 0x06, 0xe3, 0x47, 0xe3, 0x3c, 0x57,
	0xf4, 0xd1, 0x90, 0x67, 0xbc, 0x82, 0xb1, 0x54, 0x8a, 0xbc, 0x5f, 0xb3, 0x7d, 0xa7, 0x7d, 0x2a,
	0x16, 0x22, 0x3f, 0xad, 0x46, 0x91, 0x3d, 0x1f, 0x10, 0x9e, 0x43, 0x8f, 0x0d, 0xef, 0x28, 0x3d,
	0x0e, 0x6f, 0x31, 0xe0, 0x25, 0x8c, 0x94, 0x64, 0xda, 0x5a, 0xf7, 0xb5, 0x36, 0x3a, 0x3d, 0x59,
	0x88, 0x3c, 0xa9, 0xa0, 0x43, 0xa5, 0xce, 0xae, 0xe1, 0xac, 0xdd, 0xe4, 0x6b, 0xbb, 0xf7, 0x84,
	0x17, 0x30, 0xd4, 0x4e, 0xbe, 0xf2, 0x41, 0x17, 0x41, 0x1f, 0x84, 0x5c, 0xea, 0x6c, 0x05, 0xc9,
	0x93, 0xdc, 0x12, 0x4e, 0xa1, 0xaf, 0x1a, 0xe7, 0xad, 0x6b, 0x85, 0x36, 0x21, 0x42, 0xe2, 0xcd,
	0x77, 0xbc, 0x20, 0xa9, 0xc2, 0xbc, 0x2a, 0x61, 0xf0, 0xe0, 0x48, 0xb2, 0x75, 0x78, 0x07, 0xbd,
	0xb0, 0x0a, 0xa7, 0xc5, 0x6f, 0x23, 0x7f, 0x7f, 0x39, 0x9f, 0xfd, 0xe3, 0xf1, 0xa6, 0xec, 0x28,
	0x17, 0x37, 0xe2, 0x7e, 0xf2, 0x32, 0x2e, 0x96, 0xf5, 0xa6, 0x73, 0x36, 0xfd, 0xd0, 0xd5, 0xed,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xcd, 0xad, 0x1d, 0x60, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CreatorClient is the client API for Creator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CreatorClient interface {
	First(ctx context.Context, opts ...grpc.CallOption) (Creator_FirstClient, error)
}

type creatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCreatorClient(cc grpc.ClientConnInterface) CreatorClient {
	return &creatorClient{cc}
}

func (c *creatorClient) First(ctx context.Context, opts ...grpc.CallOption) (Creator_FirstClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Creator_serviceDesc.Streams[0], "/products.Creator/First", opts...)
	if err != nil {
		return nil, err
	}
	x := &creatorFirstClient{stream}
	return x, nil
}

type Creator_FirstClient interface {
	Send(*FirstRequest) error
	Recv() (*FirstResponse, error)
	grpc.ClientStream
}

type creatorFirstClient struct {
	grpc.ClientStream
}

func (x *creatorFirstClient) Send(m *FirstRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *creatorFirstClient) Recv() (*FirstResponse, error) {
	m := new(FirstResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CreatorServer is the server API for Creator service.
type CreatorServer interface {
	First(Creator_FirstServer) error
}

// UnimplementedCreatorServer can be embedded to have forward compatible implementations.
type UnimplementedCreatorServer struct {
}

func (*UnimplementedCreatorServer) First(srv Creator_FirstServer) error {
	return status.Errorf(codes.Unimplemented, "method First not implemented")
}

func RegisterCreatorServer(s *grpc.Server, srv CreatorServer) {
	s.RegisterService(&_Creator_serviceDesc, srv)
}

func _Creator_First_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CreatorServer).First(&creatorFirstServer{stream})
}

type Creator_FirstServer interface {
	Send(*FirstResponse) error
	Recv() (*FirstRequest, error)
	grpc.ServerStream
}

type creatorFirstServer struct {
	grpc.ServerStream
}

func (x *creatorFirstServer) Send(m *FirstResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *creatorFirstServer) Recv() (*FirstRequest, error) {
	m := new(FirstRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Creator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "products.Creator",
	HandlerType: (*CreatorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "First",
			Handler:       _Creator_First_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "products.proto",
}