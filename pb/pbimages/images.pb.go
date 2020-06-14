// Code generated by protoc-gen-go. DO NOT EDIT.
// source: images.proto

package pbimages

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

type CloneImagesResponse struct {
	Images               map[string]*Image `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CloneImagesResponse) Reset()         { *m = CloneImagesResponse{} }
func (m *CloneImagesResponse) String() string { return proto.CompactTextString(m) }
func (*CloneImagesResponse) ProtoMessage()    {}
func (*CloneImagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77d46971e2a21737, []int{0}
}

func (m *CloneImagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloneImagesResponse.Unmarshal(m, b)
}
func (m *CloneImagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloneImagesResponse.Marshal(b, m, deterministic)
}
func (m *CloneImagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloneImagesResponse.Merge(m, src)
}
func (m *CloneImagesResponse) XXX_Size() int {
	return xxx_messageInfo_CloneImagesResponse.Size(m)
}
func (m *CloneImagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloneImagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloneImagesResponse proto.InternalMessageInfo

func (m *CloneImagesResponse) GetImages() map[string]*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

type CloneImagesRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	FromEntityId         uint64   `protobuf:"varint,2,opt,name=from_entity_id,json=fromEntityId,proto3" json:"from_entity_id,omitempty"`
	DestEntityId         uint64   `protobuf:"varint,3,opt,name=dest_entity_id,json=destEntityId,proto3" json:"dest_entity_id,omitempty"`
	EntityTable          string   `protobuf:"bytes,4,opt,name=entity_table,json=entityTable,proto3" json:"entity_table,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloneImagesRequest) Reset()         { *m = CloneImagesRequest{} }
func (m *CloneImagesRequest) String() string { return proto.CompactTextString(m) }
func (*CloneImagesRequest) ProtoMessage()    {}
func (*CloneImagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77d46971e2a21737, []int{1}
}

func (m *CloneImagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloneImagesRequest.Unmarshal(m, b)
}
func (m *CloneImagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloneImagesRequest.Marshal(b, m, deterministic)
}
func (m *CloneImagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloneImagesRequest.Merge(m, src)
}
func (m *CloneImagesRequest) XXX_Size() int {
	return xxx_messageInfo_CloneImagesRequest.Size(m)
}
func (m *CloneImagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloneImagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloneImagesRequest proto.InternalMessageInfo

func (m *CloneImagesRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *CloneImagesRequest) GetFromEntityId() uint64 {
	if m != nil {
		return m.FromEntityId
	}
	return 0
}

func (m *CloneImagesRequest) GetDestEntityId() uint64 {
	if m != nil {
		return m.DestEntityId
	}
	return 0
}

func (m *CloneImagesRequest) GetEntityTable() string {
	if m != nil {
		return m.EntityTable
	}
	return ""
}

type ChangeEntityImagesRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	From                 *Entity  `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Dest                 *Entity  `protobuf:"bytes,3,opt,name=dest,proto3" json:"dest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeEntityImagesRequest) Reset()         { *m = ChangeEntityImagesRequest{} }
func (m *ChangeEntityImagesRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeEntityImagesRequest) ProtoMessage()    {}
func (*ChangeEntityImagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77d46971e2a21737, []int{2}
}

func (m *ChangeEntityImagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeEntityImagesRequest.Unmarshal(m, b)
}
func (m *ChangeEntityImagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeEntityImagesRequest.Marshal(b, m, deterministic)
}
func (m *ChangeEntityImagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeEntityImagesRequest.Merge(m, src)
}
func (m *ChangeEntityImagesRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeEntityImagesRequest.Size(m)
}
func (m *ChangeEntityImagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeEntityImagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeEntityImagesRequest proto.InternalMessageInfo

func (m *ChangeEntityImagesRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *ChangeEntityImagesRequest) GetFrom() *Entity {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *ChangeEntityImagesRequest) GetDest() *Entity {
	if m != nil {
		return m.Dest
	}
	return nil
}

type Entity struct {
	EntityId             uint64   `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	EntityTable          string   `protobuf:"bytes,2,opt,name=entity_table,json=entityTable,proto3" json:"entity_table,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_77d46971e2a21737, []int{3}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetEntityId() uint64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

func (m *Entity) GetEntityTable() string {
	if m != nil {
		return m.EntityTable
	}
	return ""
}

type CreateImageRequest struct {
	// Types that are valid to be assigned to Corpus:
	//	*CreateImageRequest_Metadata_
	//	*CreateImageRequest_Chunk
	Corpus               isCreateImageRequest_Corpus `protobuf_oneof:"corpus"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *CreateImageRequest) Reset()         { *m = CreateImageRequest{} }
func (m *CreateImageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateImageRequest) ProtoMessage()    {}
func (*CreateImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77d46971e2a21737, []int{4}
}

func (m *CreateImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImageRequest.Unmarshal(m, b)
}
func (m *CreateImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImageRequest.Marshal(b, m, deterministic)
}
func (m *CreateImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImageRequest.Merge(m, src)
}
func (m *CreateImageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateImageRequest.Size(m)
}
func (m *CreateImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImageRequest proto.InternalMessageInfo

type isCreateImageRequest_Corpus interface {
	isCreateImageRequest_Corpus()
}

type CreateImageRequest_Metadata_ struct {
	Metadata *CreateImageRequest_Metadata `protobuf:"bytes,1,opt,name=metadata,proto3,oneof"`
}

type CreateImageRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*CreateImageRequest_Metadata_) isCreateImageRequest_Corpus() {}

func (*CreateImageRequest_Chunk) isCreateImageRequest_Corpus() {}

func (m *CreateImageRequest) GetCorpus() isCreateImageRequest_Corpus {
	if m != nil {
		return m.Corpus
	}
	return nil
}

func (m *CreateImageRequest) GetMetadata() *CreateImageRequest_Metadata {
	if x, ok := m.GetCorpus().(*CreateImageRequest_Metadata_); ok {
		return x.Metadata
	}
	return nil
}

func (m *CreateImageRequest) GetChunk() []byte {
	if x, ok := m.GetCorpus().(*CreateImageRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CreateImageRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CreateImageRequest_Metadata_)(nil),
		(*CreateImageRequest_Chunk)(nil),
	}
}

type CreateImageRequest_Metadata struct {
	Ext                  string   `protobuf:"bytes,1,opt,name=ext,proto3" json:"ext,omitempty"`
	AccessToken          string   `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Entity               *Entity  `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateImageRequest_Metadata) Reset()         { *m = CreateImageRequest_Metadata{} }
func (m *CreateImageRequest_Metadata) String() string { return proto.CompactTextString(m) }
func (*CreateImageRequest_Metadata) ProtoMessage()    {}
func (*CreateImageRequest_Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_77d46971e2a21737, []int{4, 0}
}

func (m *CreateImageRequest_Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImageRequest_Metadata.Unmarshal(m, b)
}
func (m *CreateImageRequest_Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImageRequest_Metadata.Marshal(b, m, deterministic)
}
func (m *CreateImageRequest_Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImageRequest_Metadata.Merge(m, src)
}
func (m *CreateImageRequest_Metadata) XXX_Size() int {
	return xxx_messageInfo_CreateImageRequest_Metadata.Size(m)
}
func (m *CreateImageRequest_Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImageRequest_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImageRequest_Metadata proto.InternalMessageInfo

func (m *CreateImageRequest_Metadata) GetExt() string {
	if m != nil {
		return m.Ext
	}
	return ""
}

func (m *CreateImageRequest_Metadata) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *CreateImageRequest_Metadata) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type Image struct {
	Size                 int64    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Uri                  string   `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	Entity               *Entity  `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_77d46971e2a21737, []int{5}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Image) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Image) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type CreateImageResponse struct {
	ImageId              string   `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	Image                *Image   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateImageResponse) Reset()         { *m = CreateImageResponse{} }
func (m *CreateImageResponse) String() string { return proto.CompactTextString(m) }
func (*CreateImageResponse) ProtoMessage()    {}
func (*CreateImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77d46971e2a21737, []int{6}
}

func (m *CreateImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImageResponse.Unmarshal(m, b)
}
func (m *CreateImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImageResponse.Marshal(b, m, deterministic)
}
func (m *CreateImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImageResponse.Merge(m, src)
}
func (m *CreateImageResponse) XXX_Size() int {
	return xxx_messageInfo_CreateImageResponse.Size(m)
}
func (m *CreateImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImageResponse proto.InternalMessageInfo

func (m *CreateImageResponse) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *CreateImageResponse) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type RetrieveImagesRequest struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveImagesRequest) Reset()         { *m = RetrieveImagesRequest{} }
func (m *RetrieveImagesRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveImagesRequest) ProtoMessage()    {}
func (*RetrieveImagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77d46971e2a21737, []int{7}
}

func (m *RetrieveImagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveImagesRequest.Unmarshal(m, b)
}
func (m *RetrieveImagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveImagesRequest.Marshal(b, m, deterministic)
}
func (m *RetrieveImagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveImagesRequest.Merge(m, src)
}
func (m *RetrieveImagesRequest) XXX_Size() int {
	return xxx_messageInfo_RetrieveImagesRequest.Size(m)
}
func (m *RetrieveImagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveImagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveImagesRequest proto.InternalMessageInfo

func (m *RetrieveImagesRequest) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type RetrieveImagesResponse struct {
	Images               map[string]*Image `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RetrieveImagesResponse) Reset()         { *m = RetrieveImagesResponse{} }
func (m *RetrieveImagesResponse) String() string { return proto.CompactTextString(m) }
func (*RetrieveImagesResponse) ProtoMessage()    {}
func (*RetrieveImagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77d46971e2a21737, []int{8}
}

func (m *RetrieveImagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveImagesResponse.Unmarshal(m, b)
}
func (m *RetrieveImagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveImagesResponse.Marshal(b, m, deterministic)
}
func (m *RetrieveImagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveImagesResponse.Merge(m, src)
}
func (m *RetrieveImagesResponse) XXX_Size() int {
	return xxx_messageInfo_RetrieveImagesResponse.Size(m)
}
func (m *RetrieveImagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveImagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveImagesResponse proto.InternalMessageInfo

func (m *RetrieveImagesResponse) GetImages() map[string]*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

type DeleteImagesRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Entity               *Entity  `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteImagesRequest) Reset()         { *m = DeleteImagesRequest{} }
func (m *DeleteImagesRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteImagesRequest) ProtoMessage()    {}
func (*DeleteImagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77d46971e2a21737, []int{9}
}

func (m *DeleteImagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteImagesRequest.Unmarshal(m, b)
}
func (m *DeleteImagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteImagesRequest.Marshal(b, m, deterministic)
}
func (m *DeleteImagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteImagesRequest.Merge(m, src)
}
func (m *DeleteImagesRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteImagesRequest.Size(m)
}
func (m *DeleteImagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteImagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteImagesRequest proto.InternalMessageInfo

func (m *DeleteImagesRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *DeleteImagesRequest) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*CloneImagesResponse)(nil), "images.CloneImagesResponse")
	proto.RegisterMapType((map[string]*Image)(nil), "images.CloneImagesResponse.ImagesEntry")
	proto.RegisterType((*CloneImagesRequest)(nil), "images.CloneImagesRequest")
	proto.RegisterType((*ChangeEntityImagesRequest)(nil), "images.ChangeEntityImagesRequest")
	proto.RegisterType((*Entity)(nil), "images.Entity")
	proto.RegisterType((*CreateImageRequest)(nil), "images.CreateImageRequest")
	proto.RegisterType((*CreateImageRequest_Metadata)(nil), "images.CreateImageRequest.Metadata")
	proto.RegisterType((*Image)(nil), "images.Image")
	proto.RegisterType((*CreateImageResponse)(nil), "images.CreateImageResponse")
	proto.RegisterType((*RetrieveImagesRequest)(nil), "images.RetrieveImagesRequest")
	proto.RegisterType((*RetrieveImagesResponse)(nil), "images.RetrieveImagesResponse")
	proto.RegisterMapType((map[string]*Image)(nil), "images.RetrieveImagesResponse.ImagesEntry")
	proto.RegisterType((*DeleteImagesRequest)(nil), "images.DeleteImagesRequest")
}

func init() {
	proto.RegisterFile("images.proto", fileDescriptor_77d46971e2a21737)
}

var fileDescriptor_77d46971e2a21737 = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xcd, 0xc6, 0x69, 0xbe, 0x74, 0xec, 0xaf, 0x42, 0x5b, 0xb5, 0x4d, 0x5d, 0x81, 0xda, 0x2d,
	0x82, 0x08, 0x21, 0x47, 0x0a, 0x17, 0xc4, 0xa5, 0x22, 0x25, 0x52, 0x8a, 0xc4, 0xa1, 0xab, 0xf6,
	0xc2, 0xa5, 0x38, 0xc9, 0xd4, 0x75, 0x9b, 0xd8, 0xc6, 0xde, 0x44, 0x84, 0x2b, 0x12, 0x17, 0xfe,
	0x04, 0xa7, 0xfe, 0x48, 0x2e, 0xa0, 0xdd, 0xb5, 0x53, 0xb7, 0x8e, 0x2b, 0x7a, 0xe0, 0xb6, 0x33,
	0xf3, 0x66, 0xe6, 0xcd, 0xb3, 0x77, 0x16, 0x2c, 0x7f, 0xe2, 0x7a, 0x98, 0x38, 0x51, 0x1c, 0x8a,
	0x90, 0xd6, 0xb5, 0x65, 0x6f, 0xcd, 0xdc, 0xb1, 0x3f, 0x72, 0x05, 0xb6, 0xb3, 0x83, 0x06, 0xd8,
	0x3b, 0x5e, 0x18, 0x7a, 0x63, 0x6c, 0x2b, 0x6b, 0x30, 0x3d, 0x6f, 0xe3, 0x24, 0x12, 0x73, 0x1d,
	0x64, 0x3f, 0x09, 0xac, 0x1f, 0x8e, 0xc3, 0x00, 0x8f, 0x54, 0x15, 0x8e, 0x49, 0x14, 0x06, 0x09,
	0xd2, 0x03, 0x48, 0xeb, 0x36, 0xc9, 0xae, 0xd1, 0x32, 0x3b, 0xcf, 0x9d, 0xb4, 0xe9, 0x12, 0xb0,
	0xa3, 0xcd, 0x5e, 0x20, 0xe2, 0x39, 0xcf, 0xe8, 0xf4, 0xc1, 0xcc, 0xb9, 0xe9, 0x23, 0x30, 0xae,
	0x70, 0xde, 0x24, 0xbb, 0xa4, 0xb5, 0xca, 0xe5, 0x91, 0xee, 0xc3, 0xca, 0xcc, 0x1d, 0x4f, 0xb1,
	0x59, 0xdd, 0x25, 0x2d, 0xb3, 0xf3, 0x7f, 0xd6, 0x40, 0x65, 0x71, 0x1d, 0x7b, 0x53, 0x7d, 0x4d,
	0xd8, 0x35, 0x01, 0x7a, 0xab, 0xeb, 0xe7, 0x29, 0x26, 0x82, 0xee, 0x81, 0xe5, 0x0e, 0x87, 0x98,
	0x24, 0x67, 0x22, 0xbc, 0xc2, 0x20, 0x2d, 0x6d, 0x6a, 0xdf, 0x89, 0x74, 0xd1, 0xa7, 0xb0, 0x76,
	0x1e, 0x87, 0x93, 0x33, 0x0c, 0x84, 0x2f, 0xe6, 0x67, 0xfe, 0x48, 0xf5, 0xaa, 0x71, 0x4b, 0x7a,
	0x7b, 0xca, 0x79, 0x34, 0x92, 0xa8, 0x11, 0x26, 0x22, 0x87, 0x32, 0x34, 0x4a, 0x7a, 0x17, 0xa8,
	0x3d, 0xb0, 0x52, 0x80, 0x70, 0x07, 0x63, 0x6c, 0xd6, 0x74, 0x3b, 0xed, 0x3b, 0x91, 0x2e, 0xf6,
	0x9d, 0xc0, 0xf6, 0xe1, 0x85, 0x1b, 0x78, 0x98, 0x66, 0x3d, 0x94, 0x2f, 0x83, 0x9a, 0x64, 0x96,
	0x2a, 0xb2, 0x96, 0x29, 0xa2, 0xab, 0x71, 0x15, 0x93, 0x18, 0xc9, 0x4b, 0x71, 0x5c, 0x82, 0x91,
	0x31, 0xd6, 0x87, 0xba, 0xb6, 0xe9, 0x0e, 0xac, 0xde, 0x8c, 0x45, 0xd4, 0x58, 0x0d, 0x2c, 0x1b,
	0xa9, 0x5a, 0x1c, 0xe9, 0xb7, 0xd4, 0x3e, 0x46, 0x57, 0x68, 0xf1, 0xb3, 0x59, 0xde, 0x42, 0x63,
	0x82, 0xc2, 0x1d, 0xb9, 0xc2, 0x55, 0x55, 0xcd, 0xce, 0xfe, 0xe2, 0xff, 0x28, 0xa0, 0x9d, 0x0f,
	0x29, 0xb4, 0x5f, 0xe1, 0x8b, 0x34, 0xba, 0x09, 0x2b, 0xc3, 0x8b, 0x69, 0x70, 0xa5, 0xba, 0x5a,
	0xfd, 0x0a, 0xd7, 0xa6, 0xfd, 0x8d, 0x40, 0x23, 0x4b, 0xa0, 0x2f, 0xc1, 0xc0, 0x2f, 0x42, 0x4b,
	0xd5, 0xb5, 0x7f, 0x75, 0xb7, 0xe2, 0x0d, 0x6e, 0x24, 0x33, 0x8f, 0x1b, 0x97, 0x91, 0xc7, 0x6b,
	0x97, 0x11, 0x7a, 0xdc, 0x88, 0x02, 0x8f, 0x4b, 0x58, 0x41, 0xe1, 0x6a, 0x51, 0xe1, 0x67, 0x50,
	0xd7, 0xe3, 0x95, 0xe8, 0x97, 0x46, 0xbb, 0x0d, 0xa8, 0x0f, 0xc3, 0x38, 0x9a, 0x26, 0xec, 0x14,
	0x56, 0xd4, 0x30, 0x94, 0x42, 0x2d, 0xf1, 0xbf, 0xa2, 0x22, 0x63, 0x70, 0x75, 0x96, 0x7f, 0xf5,
	0x34, 0xf6, 0xd3, 0x46, 0xf2, 0xf8, 0xb7, 0x0d, 0xd8, 0x29, 0xac, 0xdf, 0x52, 0x2a, 0xbd, 0x76,
	0xdb, 0xd0, 0x50, 0xf8, 0xec, 0x73, 0xad, 0xf2, 0xff, 0x94, 0x7d, 0x34, 0x92, 0xf7, 0x45, 0x1d,
	0x4b, 0xee, 0x8b, 0xb2, 0xd8, 0x01, 0x6c, 0x70, 0x14, 0xb1, 0x8f, 0xb3, 0x3b, 0xb7, 0xe5, 0x86,
	0x17, 0xb9, 0x97, 0xd7, 0x35, 0x81, 0xcd, 0xbb, 0x15, 0x52, 0x6e, 0xdd, 0x3b, 0x2b, 0xe1, 0x45,
	0x56, 0x62, 0x39, 0xfe, 0x1f, 0x6f, 0x85, 0x4f, 0xb0, 0xfe, 0x0e, 0xc7, 0x28, 0x1e, 0xbe, 0x15,
	0x6e, 0xa4, 0xa8, 0xde, 0x27, 0x45, 0xe7, 0x87, 0x01, 0x75, 0x5d, 0x9c, 0xbe, 0x07, 0x33, 0xf7,
	0xb5, 0xa8, 0x5d, 0xfe, 0xb3, 0xdb, 0x3b, 0x4b, 0x63, 0x5a, 0x12, 0x56, 0x69, 0x11, 0x7a, 0x0c,
	0x6b, 0xb7, 0x05, 0xa3, 0x8f, 0xcb, 0x84, 0xd4, 0x15, 0x9f, 0xdc, 0xaf, 0x33, 0xab, 0xd0, 0x1e,
	0x58, 0x79, 0x2d, 0xe8, 0x82, 0xc3, 0x12, 0x85, 0xec, 0x4d, 0x47, 0xbf, 0x07, 0x4e, 0xf6, 0x1e,
	0x38, 0x3d, 0xf9, 0x1e, 0xb0, 0x0a, 0xed, 0x83, 0x99, 0xdb, 0xb3, 0xb9, 0x29, 0x0b, 0xcb, 0x37,
	0x37, 0x65, 0xf1, 0x39, 0x60, 0x15, 0x7a, 0x0c, 0xb4, 0xb8, 0x08, 0xe9, 0xde, 0x22, 0xa9, 0x6c,
	0x49, 0x96, 0x93, 0xeb, 0x5a, 0x1f, 0xc1, 0x69, 0x47, 0x03, 0x5d, 0x61, 0x50, 0x57, 0xf1, 0x57,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xe2, 0x1c, 0xe5, 0x0b, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ImagesClient is the client API for Images service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImagesClient interface {
	CreateImage(ctx context.Context, opts ...grpc.CallOption) (Images_CreateImageClient, error)
	RetrieveImages(ctx context.Context, in *RetrieveImagesRequest, opts ...grpc.CallOption) (*RetrieveImagesResponse, error)
	DeleteImages(ctx context.Context, in *DeleteImagesRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CloneImages(ctx context.Context, in *CloneImagesRequest, opts ...grpc.CallOption) (*CloneImagesResponse, error)
	ChangeEntityImages(ctx context.Context, in *ChangeEntityImagesRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type imagesClient struct {
	cc grpc.ClientConnInterface
}

func NewImagesClient(cc grpc.ClientConnInterface) ImagesClient {
	return &imagesClient{cc}
}

func (c *imagesClient) CreateImage(ctx context.Context, opts ...grpc.CallOption) (Images_CreateImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Images_serviceDesc.Streams[0], "/images.Images/CreateImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &imagesCreateImageClient{stream}
	return x, nil
}

type Images_CreateImageClient interface {
	Send(*CreateImageRequest) error
	CloseAndRecv() (*CreateImageResponse, error)
	grpc.ClientStream
}

type imagesCreateImageClient struct {
	grpc.ClientStream
}

func (x *imagesCreateImageClient) Send(m *CreateImageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imagesCreateImageClient) CloseAndRecv() (*CreateImageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreateImageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imagesClient) RetrieveImages(ctx context.Context, in *RetrieveImagesRequest, opts ...grpc.CallOption) (*RetrieveImagesResponse, error) {
	out := new(RetrieveImagesResponse)
	err := c.cc.Invoke(ctx, "/images.Images/RetrieveImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) DeleteImages(ctx context.Context, in *DeleteImagesRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/images.Images/DeleteImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) CloneImages(ctx context.Context, in *CloneImagesRequest, opts ...grpc.CallOption) (*CloneImagesResponse, error) {
	out := new(CloneImagesResponse)
	err := c.cc.Invoke(ctx, "/images.Images/CloneImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) ChangeEntityImages(ctx context.Context, in *ChangeEntityImagesRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/images.Images/ChangeEntityImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImagesServer is the server API for Images service.
type ImagesServer interface {
	CreateImage(Images_CreateImageServer) error
	RetrieveImages(context.Context, *RetrieveImagesRequest) (*RetrieveImagesResponse, error)
	DeleteImages(context.Context, *DeleteImagesRequest) (*empty.Empty, error)
	CloneImages(context.Context, *CloneImagesRequest) (*CloneImagesResponse, error)
	ChangeEntityImages(context.Context, *ChangeEntityImagesRequest) (*empty.Empty, error)
}

// UnimplementedImagesServer can be embedded to have forward compatible implementations.
type UnimplementedImagesServer struct {
}

func (*UnimplementedImagesServer) CreateImage(srv Images_CreateImageServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateImage not implemented")
}
func (*UnimplementedImagesServer) RetrieveImages(ctx context.Context, req *RetrieveImagesRequest) (*RetrieveImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveImages not implemented")
}
func (*UnimplementedImagesServer) DeleteImages(ctx context.Context, req *DeleteImagesRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImages not implemented")
}
func (*UnimplementedImagesServer) CloneImages(ctx context.Context, req *CloneImagesRequest) (*CloneImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneImages not implemented")
}
func (*UnimplementedImagesServer) ChangeEntityImages(ctx context.Context, req *ChangeEntityImagesRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeEntityImages not implemented")
}

func RegisterImagesServer(s *grpc.Server, srv ImagesServer) {
	s.RegisterService(&_Images_serviceDesc, srv)
}

func _Images_CreateImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImagesServer).CreateImage(&imagesCreateImageServer{stream})
}

type Images_CreateImageServer interface {
	SendAndClose(*CreateImageResponse) error
	Recv() (*CreateImageRequest, error)
	grpc.ServerStream
}

type imagesCreateImageServer struct {
	grpc.ServerStream
}

func (x *imagesCreateImageServer) SendAndClose(m *CreateImageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imagesCreateImageServer) Recv() (*CreateImageRequest, error) {
	m := new(CreateImageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Images_RetrieveImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).RetrieveImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/images.Images/RetrieveImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).RetrieveImages(ctx, req.(*RetrieveImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_DeleteImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).DeleteImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/images.Images/DeleteImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).DeleteImages(ctx, req.(*DeleteImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_CloneImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).CloneImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/images.Images/CloneImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).CloneImages(ctx, req.(*CloneImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_ChangeEntityImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeEntityImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).ChangeEntityImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/images.Images/ChangeEntityImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).ChangeEntityImages(ctx, req.(*ChangeEntityImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Images_serviceDesc = grpc.ServiceDesc{
	ServiceName: "images.Images",
	HandlerType: (*ImagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveImages",
			Handler:    _Images_RetrieveImages_Handler,
		},
		{
			MethodName: "DeleteImages",
			Handler:    _Images_DeleteImages_Handler,
		},
		{
			MethodName: "CloneImages",
			Handler:    _Images_CloneImages_Handler,
		},
		{
			MethodName: "ChangeEntityImages",
			Handler:    _Images_ChangeEntityImages_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateImage",
			Handler:       _Images_CreateImage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "images.proto",
}
