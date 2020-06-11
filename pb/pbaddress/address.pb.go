// Code generated by protoc-gen-go. DO NOT EDIT.
// source: address.proto

package pbaddress

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

type Address struct {
	Country              string   `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	Province             string   `protobuf:"bytes,2,opt,name=province,proto3" json:"province,omitempty"`
	Zipcode              string   `protobuf:"bytes,3,opt,name=zipcode,proto3" json:"zipcode,omitempty"`
	Street               string   `protobuf:"bytes,4,opt,name=street,proto3" json:"street,omitempty"`
	Number               uint64   `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	Floor                uint64   `protobuf:"varint,6,opt,name=floor,proto3" json:"floor,omitempty"`
	Department           uint64   `protobuf:"varint,7,opt,name=department,proto3" json:"department,omitempty"`
	Latitude             uint64   `protobuf:"varint,8,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            uint64   `protobuf:"varint,9,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Alias                string   `protobuf:"bytes,10,opt,name=alias,proto3" json:"alias,omitempty"`
	UserId               uint64   `protobuf:"varint,11,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{0}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Address) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Address) GetZipcode() string {
	if m != nil {
		return m.Zipcode
	}
	return ""
}

func (m *Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *Address) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Address) GetFloor() uint64 {
	if m != nil {
		return m.Floor
	}
	return 0
}

func (m *Address) GetDepartment() uint64 {
	if m != nil {
		return m.Department
	}
	return 0
}

func (m *Address) GetLatitude() uint64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Address) GetLongitude() uint64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Address) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *Address) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateAddressRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Body                 *Address `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAddressRequest) Reset()         { *m = CreateAddressRequest{} }
func (m *CreateAddressRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAddressRequest) ProtoMessage()    {}
func (*CreateAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{1}
}

func (m *CreateAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAddressRequest.Unmarshal(m, b)
}
func (m *CreateAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAddressRequest.Marshal(b, m, deterministic)
}
func (m *CreateAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAddressRequest.Merge(m, src)
}
func (m *CreateAddressRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAddressRequest.Size(m)
}
func (m *CreateAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAddressRequest proto.InternalMessageInfo

func (m *CreateAddressRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *CreateAddressRequest) GetBody() *Address {
	if m != nil {
		return m.Body
	}
	return nil
}

type CreateAddressResponse struct {
	AddressId            uint64   `protobuf:"varint,1,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	Metadata             *Address `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAddressResponse) Reset()         { *m = CreateAddressResponse{} }
func (m *CreateAddressResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAddressResponse) ProtoMessage()    {}
func (*CreateAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{2}
}

func (m *CreateAddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAddressResponse.Unmarshal(m, b)
}
func (m *CreateAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAddressResponse.Marshal(b, m, deterministic)
}
func (m *CreateAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAddressResponse.Merge(m, src)
}
func (m *CreateAddressResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAddressResponse.Size(m)
}
func (m *CreateAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAddressResponse proto.InternalMessageInfo

func (m *CreateAddressResponse) GetAddressId() uint64 {
	if m != nil {
		return m.AddressId
	}
	return 0
}

func (m *CreateAddressResponse) GetMetadata() *Address {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type RetrieveAddressRequest struct {
	AddressId            uint64   `protobuf:"varint,1,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveAddressRequest) Reset()         { *m = RetrieveAddressRequest{} }
func (m *RetrieveAddressRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveAddressRequest) ProtoMessage()    {}
func (*RetrieveAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{3}
}

func (m *RetrieveAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveAddressRequest.Unmarshal(m, b)
}
func (m *RetrieveAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveAddressRequest.Marshal(b, m, deterministic)
}
func (m *RetrieveAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveAddressRequest.Merge(m, src)
}
func (m *RetrieveAddressRequest) XXX_Size() int {
	return xxx_messageInfo_RetrieveAddressRequest.Size(m)
}
func (m *RetrieveAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveAddressRequest proto.InternalMessageInfo

func (m *RetrieveAddressRequest) GetAddressId() uint64 {
	if m != nil {
		return m.AddressId
	}
	return 0
}

type RetrieveMyAddressesRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveMyAddressesRequest) Reset()         { *m = RetrieveMyAddressesRequest{} }
func (m *RetrieveMyAddressesRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveMyAddressesRequest) ProtoMessage()    {}
func (*RetrieveMyAddressesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{4}
}

func (m *RetrieveMyAddressesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveMyAddressesRequest.Unmarshal(m, b)
}
func (m *RetrieveMyAddressesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveMyAddressesRequest.Marshal(b, m, deterministic)
}
func (m *RetrieveMyAddressesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveMyAddressesRequest.Merge(m, src)
}
func (m *RetrieveMyAddressesRequest) XXX_Size() int {
	return xxx_messageInfo_RetrieveMyAddressesRequest.Size(m)
}
func (m *RetrieveMyAddressesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveMyAddressesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveMyAddressesRequest proto.InternalMessageInfo

func (m *RetrieveMyAddressesRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type RetrieveMyAddressesResponse struct {
	Addresses            map[uint64]*Address `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RetrieveMyAddressesResponse) Reset()         { *m = RetrieveMyAddressesResponse{} }
func (m *RetrieveMyAddressesResponse) String() string { return proto.CompactTextString(m) }
func (*RetrieveMyAddressesResponse) ProtoMessage()    {}
func (*RetrieveMyAddressesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{5}
}

func (m *RetrieveMyAddressesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveMyAddressesResponse.Unmarshal(m, b)
}
func (m *RetrieveMyAddressesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveMyAddressesResponse.Marshal(b, m, deterministic)
}
func (m *RetrieveMyAddressesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveMyAddressesResponse.Merge(m, src)
}
func (m *RetrieveMyAddressesResponse) XXX_Size() int {
	return xxx_messageInfo_RetrieveMyAddressesResponse.Size(m)
}
func (m *RetrieveMyAddressesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveMyAddressesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveMyAddressesResponse proto.InternalMessageInfo

func (m *RetrieveMyAddressesResponse) GetAddresses() map[uint64]*Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func init() {
	proto.RegisterType((*Address)(nil), "address.Address")
	proto.RegisterType((*CreateAddressRequest)(nil), "address.CreateAddressRequest")
	proto.RegisterType((*CreateAddressResponse)(nil), "address.CreateAddressResponse")
	proto.RegisterType((*RetrieveAddressRequest)(nil), "address.RetrieveAddressRequest")
	proto.RegisterType((*RetrieveMyAddressesRequest)(nil), "address.RetrieveMyAddressesRequest")
	proto.RegisterType((*RetrieveMyAddressesResponse)(nil), "address.RetrieveMyAddressesResponse")
	proto.RegisterMapType((map[uint64]*Address)(nil), "address.RetrieveMyAddressesResponse.AddressesEntry")
}

func init() {
	proto.RegisterFile("address.proto", fileDescriptor_982c640dad8fe78e)
}

var fileDescriptor_982c640dad8fe78e = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x5d, 0xfa, 0xdf, 0xaf, 0x14, 0xa6, 0x8f, 0xb1, 0x59, 0x85, 0x8d, 0x12, 0x26, 0xd4, 0x0b,
	0xd4, 0x49, 0xdd, 0x05, 0x88, 0x1b, 0x04, 0x08, 0xa4, 0x5d, 0x80, 0xa0, 0xe2, 0x8a, 0x9b, 0xca,
	0x8d, 0x3f, 0x50, 0xb4, 0x34, 0x0e, 0xb6, 0x53, 0xa9, 0xbc, 0x01, 0x6f, 0xc2, 0x63, 0xf0, 0x68,
	0x28, 0xb6, 0x93, 0xb1, 0xae, 0xad, 0xe0, 0x2e, 0xe7, 0x1c, 0x9f, 0xef, 0xe7, 0xd8, 0x0a, 0xf4,
	0xb9, 0x10, 0x8a, 0xb4, 0x1e, 0x67, 0x4a, 0x1a, 0x89, 0x6d, 0x0f, 0x07, 0x47, 0x4b, 0x9e, 0xc4,
	0x82, 0x1b, 0x3a, 0x2b, 0x3f, 0xdc, 0x89, 0xf0, 0x57, 0x0d, 0xda, 0xaf, 0xdc, 0x21, 0x64, 0xd0,
	0x8e, 0x64, 0x9e, 0x1a, 0xb5, 0x62, 0xc1, 0x30, 0x18, 0x75, 0xa7, 0x25, 0xc4, 0x01, 0x74, 0x32,
	0x25, 0x97, 0x71, 0x1a, 0x11, 0xab, 0x59, 0xa9, 0xc2, 0x85, 0xeb, 0x47, 0x9c, 0x45, 0x52, 0x10,
	0xab, 0x3b, 0x97, 0x87, 0x78, 0x08, 0x2d, 0x6d, 0x14, 0x91, 0x61, 0x0d, 0x2b, 0x78, 0x54, 0xf0,
	0x69, 0xbe, 0x98, 0x93, 0x62, 0xcd, 0x61, 0x30, 0x6a, 0x4c, 0x3d, 0xc2, 0x03, 0x68, 0x7e, 0x4d,
	0xa4, 0x54, 0xac, 0x65, 0x69, 0x07, 0xf0, 0x04, 0x40, 0x50, 0xc6, 0x95, 0x59, 0x50, 0x6a, 0x58,
	0xdb, 0x4a, 0x7f, 0x31, 0xc5, 0x6c, 0x09, 0x37, 0xb1, 0xc9, 0x05, 0xb1, 0x8e, 0x55, 0x2b, 0x8c,
	0x0f, 0xa0, 0x9b, 0xc8, 0xf4, 0x9b, 0x13, 0xbb, 0x56, 0xbc, 0x22, 0x8a, 0x7e, 0x3c, 0x89, 0xb9,
	0x66, 0x60, 0xc7, 0x73, 0x00, 0x8f, 0xa0, 0x9d, 0x6b, 0x52, 0xb3, 0x58, 0xb0, 0x9e, 0x1b, 0xaf,
	0x80, 0x17, 0x22, 0x9c, 0xc1, 0xc1, 0x1b, 0x45, 0xdc, 0x90, 0xcf, 0x6b, 0x4a, 0xdf, 0x73, 0xd2,
	0x06, 0x1f, 0xc1, 0x2d, 0x1e, 0x45, 0xa4, 0xf5, 0xcc, 0xc8, 0x4b, 0x4a, 0x7d, 0x76, 0x3d, 0xc7,
	0x7d, 0x2e, 0x28, 0x3c, 0x85, 0xc6, 0x5c, 0x8a, 0x95, 0xcd, 0xae, 0x37, 0xd9, 0x1f, 0x97, 0xb7,
	0x54, 0x56, 0xb2, 0x6a, 0x28, 0xe0, 0xde, 0x5a, 0x03, 0x9d, 0xc9, 0x54, 0x13, 0x1e, 0x03, 0x78,
	0x47, 0x31, 0x55, 0xe0, 0xf6, 0xf0, 0xcc, 0x85, 0xc0, 0xa7, 0xd0, 0x59, 0x90, 0xe1, 0x82, 0x1b,
	0xbe, 0xb5, 0x43, 0x75, 0x22, 0x7c, 0x06, 0x87, 0x53, 0x32, 0x2a, 0xa6, 0xe5, 0xfa, 0x22, 0xbb,
	0xdb, 0x84, 0x2f, 0x61, 0x50, 0x1a, 0xdf, 0xaf, 0xbc, 0x95, 0xfe, 0x23, 0x85, 0xf0, 0x77, 0x00,
	0xf7, 0x37, 0x56, 0xf0, 0x6b, 0x7e, 0x82, 0xb2, 0x1b, 0x69, 0x16, 0x0c, 0xeb, 0xa3, 0xde, 0xe4,
	0xbc, 0x5a, 0x64, 0x87, 0x71, 0x5c, 0x31, 0x6f, 0x8b, 0xd7, 0x3a, 0xbd, 0xaa, 0x32, 0xf8, 0x00,
	0xb7, 0xaf, 0x8b, 0xb8, 0x0f, 0xf5, 0x4b, 0x5a, 0xf9, 0xed, 0x8a, 0x4f, 0x7c, 0x02, 0xcd, 0x25,
	0x4f, 0x72, 0xda, 0x9a, 0x9d, 0x93, 0x5f, 0xd4, 0x9e, 0x07, 0x93, 0x9f, 0x35, 0xe8, 0x56, 0x05,
	0xf1, 0x23, 0xf4, 0xaf, 0x5d, 0x18, 0x1e, 0x57, 0xde, 0x4d, 0x2f, 0x65, 0x70, 0xb2, 0x4d, 0x76,
	0x7b, 0x84, 0x7b, 0xf8, 0x0e, 0xee, 0xac, 0x5d, 0x0e, 0x3e, 0xbc, 0x11, 0xc1, 0x5a, 0xd5, 0x1b,
	0x03, 0x87, 0x7b, 0x38, 0x87, 0xbb, 0x1b, 0x02, 0xc3, 0xc7, 0xbb, 0xe3, 0x74, 0xf5, 0x4e, 0xff,
	0x25, 0xf3, 0x70, 0xef, 0x75, 0xff, 0x4b, 0x6f, 0x7c, 0x96, 0xcd, 0xfd, 0xe1, 0x79, 0xcb, 0xfe,
	0x50, 0xce, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x4e, 0xb8, 0x49, 0x83, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AddressesClient is the client API for Addresses service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddressesClient interface {
	CreateAddress(ctx context.Context, in *CreateAddressRequest, opts ...grpc.CallOption) (*CreateAddressResponse, error)
	RetrieveAddress(ctx context.Context, in *RetrieveAddressRequest, opts ...grpc.CallOption) (*Address, error)
	RetrieveMyAddresses(ctx context.Context, in *RetrieveMyAddressesRequest, opts ...grpc.CallOption) (*RetrieveMyAddressesResponse, error)
}

type addressesClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressesClient(cc grpc.ClientConnInterface) AddressesClient {
	return &addressesClient{cc}
}

func (c *addressesClient) CreateAddress(ctx context.Context, in *CreateAddressRequest, opts ...grpc.CallOption) (*CreateAddressResponse, error) {
	out := new(CreateAddressResponse)
	err := c.cc.Invoke(ctx, "/address.Addresses/CreateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressesClient) RetrieveAddress(ctx context.Context, in *RetrieveAddressRequest, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/address.Addresses/RetrieveAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressesClient) RetrieveMyAddresses(ctx context.Context, in *RetrieveMyAddressesRequest, opts ...grpc.CallOption) (*RetrieveMyAddressesResponse, error) {
	out := new(RetrieveMyAddressesResponse)
	err := c.cc.Invoke(ctx, "/address.Addresses/RetrieveMyAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressesServer is the server API for Addresses service.
type AddressesServer interface {
	CreateAddress(context.Context, *CreateAddressRequest) (*CreateAddressResponse, error)
	RetrieveAddress(context.Context, *RetrieveAddressRequest) (*Address, error)
	RetrieveMyAddresses(context.Context, *RetrieveMyAddressesRequest) (*RetrieveMyAddressesResponse, error)
}

// UnimplementedAddressesServer can be embedded to have forward compatible implementations.
type UnimplementedAddressesServer struct {
}

func (*UnimplementedAddressesServer) CreateAddress(ctx context.Context, req *CreateAddressRequest) (*CreateAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddress not implemented")
}
func (*UnimplementedAddressesServer) RetrieveAddress(ctx context.Context, req *RetrieveAddressRequest) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveAddress not implemented")
}
func (*UnimplementedAddressesServer) RetrieveMyAddresses(ctx context.Context, req *RetrieveMyAddressesRequest) (*RetrieveMyAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveMyAddresses not implemented")
}

func RegisterAddressesServer(s *grpc.Server, srv AddressesServer) {
	s.RegisterService(&_Addresses_serviceDesc, srv)
}

func _Addresses_CreateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressesServer).CreateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/address.Addresses/CreateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressesServer).CreateAddress(ctx, req.(*CreateAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Addresses_RetrieveAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressesServer).RetrieveAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/address.Addresses/RetrieveAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressesServer).RetrieveAddress(ctx, req.(*RetrieveAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Addresses_RetrieveMyAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveMyAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressesServer).RetrieveMyAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/address.Addresses/RetrieveMyAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressesServer).RetrieveMyAddresses(ctx, req.(*RetrieveMyAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Addresses_serviceDesc = grpc.ServiceDesc{
	ServiceName: "address.Addresses",
	HandlerType: (*AddressesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAddress",
			Handler:    _Addresses_CreateAddress_Handler,
		},
		{
			MethodName: "RetrieveAddress",
			Handler:    _Addresses_RetrieveAddress_Handler,
		},
		{
			MethodName: "RetrieveMyAddresses",
			Handler:    _Addresses_RetrieveMyAddresses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address.proto",
}
