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
	IsDeletion           bool     `protobuf:"varint,4,opt,name=is_deletion,json=isDeletion,proto3" json:"is_deletion,omitempty"`
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

func (m *FirstRequest) GetIsDeletion() bool {
	if m != nil {
		return m.IsDeletion
	}
	return false
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
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x14, 0x84, 0x5d, 0x4d, 0xdb, 0xf8, 0x1a, 0x7b, 0x58, 0xa4, 0x8d, 0xbd, 0x18, 0x73, 0x0a, 0x1e,
	0x52, 0xa9, 0x77, 0x0f, 0x2a, 0x42, 0x6e, 0x12, 0x3c, 0x79, 0x09, 0xdb, 0xec, 0xb3, 0x2c, 0x86,
	0x6c, 0xdc, 0x7d, 0x11, 0xf4, 0x1f, 0xf8, 0xaf, 0x25, 0x9b, 0x44, 0x84, 0xde, 0xde, 0x7c, 0x3b,
	0xcc, 0x0e, 0x03, 0x8b, 0xc6, 0x68, 0xd9, 0x96, 0x64, 0xd3, 0xc6, 0x68, 0xd2, 0xdc, 0x1f, 0xf5,
	0x7a, 0xf5, 0x29, 0x2a, 0x25, 0x05, 0xe1, 0x66, 0x3c, 0x7a, 0x4b, 0xfc, 0xc3, 0x20, 0x78, 0x52,
	0xc6, 0x52, 0x8e, 0x1f, 0x2d, 0x5a, 0xe2, 0x57, 0x10, 0x88, 0xb2, 0x44, 0x6b, 0x0b, 0xd2, 0xef,
	0x58, 0x87, 0x2c, 0x62, 0xc9, 0x69, 0x3e, 0xef, 0xd9, 0x4b, 0x87, 0xf8, 0x39, 0x4c, 0x48, 0x51,
	0x85, 0xe1, 0xb1, 0x7b, 0xeb, 0x05, 0xbf, 0x84, 0x79, 0x29, 0x08, 0xf7, 0xda, 0x7c, 0x15, 0x4a,
	0x86, 0x27, 0x11, 0x4b, 0xbc, 0x1c, 0x46, 0x94, 0xc9, 0xce, 0xa0, 0x6c, 0x21, 0xb1, 0x42, 0x52,
	0xba, 0x0e, 0xbd, 0x88, 0x25, 0x7e, 0x0e, 0xca, 0x3e, 0x0e, 0x24, 0xbe, 0x86, 0xb3, 0xa1, 0x8a,
	0x6d, 0x74, 0x6d, 0x91, 0x5f, 0x80, 0x2f, 0x8d, 0x78, 0xa3, 0x2e, 0x8f, 0xb9, 0xbc, 0x99, 0xd3,
	0x99, 0x8c, 0xb7, 0xe0, 0x3d, 0x8b, 0x3d, 0xf2, 0x25, 0x4c, 0xcb, 0xd6, 0x58, 0x6d, 0x06, 0xc3,
	0xa0, 0x38, 0x07, 0xcf, 0xaa, 0xef, 0xbe, 0xa2, 0x97, 0xbb, 0x7b, 0x9b, 0xc1, 0xec, 0xc1, 0xa0,
	0x20, 0x6d, 0xf8, 0x1d, 0x4c, 0xdc, 0x57, 0x7c, 0x99, 0xfe, 0x6d, 0xf6, 0x7f, 0x86, 0xf5, 0xea,
	0x80, 0xf7, 0x9d, 0xe2, 0xa3, 0x84, 0xdd, 0xb0, 0xfb, 0xc5, 0x6b, 0x90, 0x6e, 0x9a, 0xdd, 0xe8,
	0xd9, 0x4d, 0xdd, 0x9a, 0xb7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x30, 0x90, 0x31, 0xbe, 0x82,
	0x01, 0x00, 0x00,
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
