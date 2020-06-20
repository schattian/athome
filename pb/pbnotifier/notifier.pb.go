// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notifier.proto

package pbnotifier

import (
	pbshared "./pbshared"
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type UpdateStatusRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	NotificationId       uint64   `protobuf:"varint,2,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateStatusRequest) Reset()         { *m = UpdateStatusRequest{} }
func (m *UpdateStatusRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateStatusRequest) ProtoMessage()    {}
func (*UpdateStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c0fc606bc4470de, []int{0}
}

func (m *UpdateStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStatusRequest.Unmarshal(m, b)
}
func (m *UpdateStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStatusRequest.Marshal(b, m, deterministic)
}
func (m *UpdateStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStatusRequest.Merge(m, src)
}
func (m *UpdateStatusRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateStatusRequest.Size(m)
}
func (m *UpdateStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStatusRequest proto.InternalMessageInfo

func (m *UpdateStatusRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *UpdateStatusRequest) GetNotificationId() uint64 {
	if m != nil {
		return m.NotificationId
	}
	return 0
}

type CreateRequest struct {
	NotificationToken    string        `protobuf:"bytes,1,opt,name=notification_token,json=notificationToken,proto3" json:"notification_token,omitempty"`
	Notification         *Notification `protobuf:"bytes,2,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c0fc606bc4470de, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetNotificationToken() string {
	if m != nil {
		return m.NotificationToken
	}
	return ""
}

func (m *CreateRequest) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

type CreateResponse struct {
	NotificationId       uint64        `protobuf:"varint,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
	Notification         *Notification `protobuf:"bytes,2,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c0fc606bc4470de, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetNotificationId() uint64 {
	if m != nil {
		return m.NotificationId
	}
	return 0
}

func (m *CreateResponse) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

type RetrieveRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	NotificationId       uint64   `protobuf:"varint,2,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveRequest) Reset()         { *m = RetrieveRequest{} }
func (m *RetrieveRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveRequest) ProtoMessage()    {}
func (*RetrieveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c0fc606bc4470de, []int{3}
}

func (m *RetrieveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveRequest.Unmarshal(m, b)
}
func (m *RetrieveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveRequest.Marshal(b, m, deterministic)
}
func (m *RetrieveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveRequest.Merge(m, src)
}
func (m *RetrieveRequest) XXX_Size() int {
	return xxx_messageInfo_RetrieveRequest.Size(m)
}
func (m *RetrieveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveRequest proto.InternalMessageInfo

func (m *RetrieveRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *RetrieveRequest) GetNotificationId() uint64 {
	if m != nil {
		return m.NotificationId
	}
	return 0
}

type RetrieveStreamRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	TickerMs             uint64   `protobuf:"varint,2,opt,name=ticker_ms,json=tickerMs,proto3" json:"ticker_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveStreamRequest) Reset()         { *m = RetrieveStreamRequest{} }
func (m *RetrieveStreamRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveStreamRequest) ProtoMessage()    {}
func (*RetrieveStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c0fc606bc4470de, []int{4}
}

func (m *RetrieveStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveStreamRequest.Unmarshal(m, b)
}
func (m *RetrieveStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveStreamRequest.Marshal(b, m, deterministic)
}
func (m *RetrieveStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveStreamRequest.Merge(m, src)
}
func (m *RetrieveStreamRequest) XXX_Size() int {
	return xxx_messageInfo_RetrieveStreamRequest.Size(m)
}
func (m *RetrieveStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveStreamRequest proto.InternalMessageInfo

func (m *RetrieveStreamRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *RetrieveStreamRequest) GetTickerMs() uint64 {
	if m != nil {
		return m.TickerMs
	}
	return 0
}

type RetrieveManyRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveManyRequest) Reset()         { *m = RetrieveManyRequest{} }
func (m *RetrieveManyRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveManyRequest) ProtoMessage()    {}
func (*RetrieveManyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c0fc606bc4470de, []int{5}
}

func (m *RetrieveManyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveManyRequest.Unmarshal(m, b)
}
func (m *RetrieveManyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveManyRequest.Marshal(b, m, deterministic)
}
func (m *RetrieveManyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveManyRequest.Merge(m, src)
}
func (m *RetrieveManyRequest) XXX_Size() int {
	return xxx_messageInfo_RetrieveManyRequest.Size(m)
}
func (m *RetrieveManyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveManyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveManyRequest proto.InternalMessageInfo

func (m *RetrieveManyRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type RetrieveManyResponse struct {
	Notifications        map[uint64]*Notification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RetrieveManyResponse) Reset()         { *m = RetrieveManyResponse{} }
func (m *RetrieveManyResponse) String() string { return proto.CompactTextString(m) }
func (*RetrieveManyResponse) ProtoMessage()    {}
func (*RetrieveManyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c0fc606bc4470de, []int{6}
}

func (m *RetrieveManyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveManyResponse.Unmarshal(m, b)
}
func (m *RetrieveManyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveManyResponse.Marshal(b, m, deterministic)
}
func (m *RetrieveManyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveManyResponse.Merge(m, src)
}
func (m *RetrieveManyResponse) XXX_Size() int {
	return xxx_messageInfo_RetrieveManyResponse.Size(m)
}
func (m *RetrieveManyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveManyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveManyResponse proto.InternalMessageInfo

func (m *RetrieveManyResponse) GetNotifications() map[uint64]*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

type Notification struct {
	UserId               uint64           `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Body                 string           `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Entity               *pbshared.Entity `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	Status               *Status          `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Priority             string           `protobuf:"bytes,5,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c0fc606bc4470de, []int{7}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Notification) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Notification) GetEntity() *pbshared.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Notification) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Notification) GetPriority() string {
	if m != nil {
		return m.Priority
	}
	return ""
}

type Status struct {
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ReceivedAt           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
	SeenAt               *timestamp.Timestamp `protobuf:"bytes,3,opt,name=seen_at,json=seenAt,proto3" json:"seen_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c0fc606bc4470de, []int{8}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Status) GetReceivedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ReceivedAt
	}
	return nil
}

func (m *Status) GetSeenAt() *timestamp.Timestamp {
	if m != nil {
		return m.SeenAt
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateStatusRequest)(nil), "notifier.UpdateStatusRequest")
	proto.RegisterType((*CreateRequest)(nil), "notifier.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "notifier.CreateResponse")
	proto.RegisterType((*RetrieveRequest)(nil), "notifier.RetrieveRequest")
	proto.RegisterType((*RetrieveStreamRequest)(nil), "notifier.RetrieveStreamRequest")
	proto.RegisterType((*RetrieveManyRequest)(nil), "notifier.RetrieveManyRequest")
	proto.RegisterType((*RetrieveManyResponse)(nil), "notifier.RetrieveManyResponse")
	proto.RegisterMapType((map[uint64]*Notification)(nil), "notifier.RetrieveManyResponse.NotificationsEntry")
	proto.RegisterType((*Notification)(nil), "notifier.Notification")
	proto.RegisterType((*Status)(nil), "notifier.Status")
}

func init() {
	proto.RegisterFile("notifier.proto", fileDescriptor_1c0fc606bc4470de)
}

var fileDescriptor_1c0fc606bc4470de = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x8e, 0xeb, 0xd4, 0x4d, 0x26, 0x69, 0xda, 0x77, 0xfb, 0xb6, 0xc9, 0xeb, 0xea, 0xa5, 0xc5,
	0x07, 0xc8, 0x01, 0x1c, 0x48, 0x25, 0x54, 0x8a, 0x10, 0x6a, 0x50, 0x91, 0x7a, 0x28, 0x48, 0x9b,
	0xa2, 0x22, 0x24, 0x14, 0x6d, 0xe2, 0x69, 0x6b, 0x35, 0xfe, 0x60, 0x77, 0x13, 0x08, 0xff, 0x8b,
	0x13, 0x67, 0x0e, 0xfc, 0x25, 0x4e, 0xc8, 0x5e, 0xbb, 0xb1, 0x49, 0x4a, 0xf9, 0x10, 0xb7, 0xdd,
	0x99, 0x67, 0x9e, 0xf9, 0xd8, 0x7d, 0x06, 0x6a, 0x7e, 0x20, 0xdd, 0x53, 0x17, 0xb9, 0x1d, 0xf2,
	0x40, 0x06, 0xa4, 0x94, 0xde, 0xcd, 0x35, 0x71, 0xce, 0x38, 0x3a, 0x2d, 0xf4, 0xa5, 0x2b, 0x27,
	0xca, 0x6d, 0xd6, 0xc7, 0x6c, 0xe8, 0x3a, 0x4c, 0x62, 0x2b, 0x3d, 0x24, 0x8e, 0xcd, 0xb3, 0x20,
	0x38, 0x1b, 0x62, 0x2b, 0xbe, 0xf5, 0x47, 0xa7, 0x2d, 0xf4, 0xc2, 0xcb, 0xa8, 0xad, 0xef, 0x9d,
	0xd2, 0xf5, 0x50, 0x48, 0xe6, 0x85, 0x0a, 0x60, 0x31, 0x58, 0x7b, 0x19, 0x46, 0x6c, 0x5d, 0xc9,
	0xe4, 0x48, 0x50, 0x7c, 0x3b, 0x42, 0x21, 0xc9, 0x4d, 0xa8, 0xb2, 0xc1, 0x00, 0x85, 0xe8, 0xc9,
	0xe0, 0x02, 0xfd, 0x86, 0xb6, 0xad, 0x35, 0xcb, 0xb4, 0xa2, 0x6c, 0xc7, 0x91, 0x89, 0xdc, 0x86,
	0x15, 0x55, 0xf1, 0x80, 0x49, 0x37, 0xf0, 0x7b, 0xae, 0xd3, 0x58, 0xd8, 0xd6, 0x9a, 0x45, 0x5a,
	0xcb, 0x9a, 0x0f, 0x1d, 0xeb, 0x03, 0x2c, 0x3f, 0xe5, 0xc8, 0x24, 0xa6, 0xe4, 0x77, 0x81, 0xe4,
	0x22, 0xb3, 0x29, 0xfe, 0xc9, 0x7a, 0x54, 0xa2, 0x3d, 0xa8, 0x66, 0x8d, 0x71, 0x96, 0x4a, 0x7b,
	0xc3, 0xbe, 0x9c, 0xdf, 0xf3, 0x8c, 0x97, 0xe6, 0xb0, 0xd6, 0x08, 0x6a, 0x69, 0x6e, 0x11, 0x06,
	0xbe, 0xc0, 0x79, 0x65, 0x6b, 0xf3, 0xca, 0xfe, 0xa3, 0xb4, 0x6f, 0x60, 0x85, 0xa2, 0xe4, 0x2e,
	0x8e, 0xf1, 0x6f, 0x4c, 0xf4, 0x04, 0xd6, 0x53, 0xfa, 0xae, 0xe4, 0xc8, 0xbc, 0x5f, 0x48, 0xb2,
	0x09, 0x65, 0xe9, 0x0e, 0x2e, 0x90, 0xf7, 0x3c, 0x91, 0xd0, 0x97, 0x94, 0xe1, 0x48, 0x58, 0xbb,
	0xb0, 0x96, 0x12, 0x1f, 0x31, 0x7f, 0xf2, 0xf3, 0xb4, 0xd6, 0x17, 0x0d, 0xfe, 0xcd, 0x87, 0x26,
	0xf3, 0x3e, 0x81, 0xe5, 0x6c, 0xf5, 0xa2, 0xa1, 0x6d, 0xeb, 0xcd, 0x4a, 0xfb, 0xfe, 0x74, 0x8e,
	0xf3, 0xc2, 0x72, 0xc3, 0x15, 0x07, 0xbe, 0xe4, 0x13, 0x9a, 0xe7, 0x31, 0x5f, 0x01, 0x99, 0x05,
	0x91, 0x55, 0xd0, 0x2f, 0x70, 0x92, 0x3c, 0x69, 0x74, 0x24, 0x77, 0x60, 0x71, 0xcc, 0x86, 0x23,
	0xbc, 0xe6, 0x01, 0x15, 0x68, 0x6f, 0x61, 0x57, 0xb3, 0x3e, 0x6b, 0x50, 0xcd, 0xfa, 0x48, 0x1d,
	0x96, 0x46, 0x02, 0xf9, 0xf4, 0xaf, 0x18, 0xd1, 0xf5, 0xd0, 0x21, 0x04, 0x8a, 0xfd, 0xc0, 0x99,
	0xc4, 0xd4, 0x65, 0x1a, 0x9f, 0xc9, 0x2d, 0x30, 0x94, 0x70, 0x1b, 0x7a, 0x9c, 0xb0, 0x66, 0x2b,
	0x39, 0xdb, 0x07, 0xb1, 0x95, 0x26, 0x5e, 0xd2, 0x04, 0x43, 0xc4, 0x9a, 0x6b, 0x14, 0x63, 0xdc,
	0xea, 0xb4, 0xb0, 0x44, 0x8b, 0x89, 0x9f, 0x3c, 0x80, 0x52, 0xc8, 0xdd, 0x80, 0x47, 0x9c, 0x8b,
	0x51, 0xa6, 0x8e, 0xf9, 0xb5, 0x53, 0xe7, 0xeb, 0x54, 0x1f, 0x06, 0xef, 0xa8, 0xee, 0xb9, 0x0e,
	0x2d, 0x9e, 0xbb, 0x67, 0xe7, 0x54, 0xf7, 0xd8, 0x7b, 0x7a, 0x89, 0xb5, 0x3e, 0x6a, 0x60, 0x28,
	0x2a, 0xf2, 0x10, 0x60, 0x10, 0xeb, 0xc0, 0xe9, 0x31, 0x19, 0x37, 0x51, 0x69, 0x9b, 0xb6, 0x5a,
	0x0e, 0x76, 0xba, 0x1c, 0xec, 0xe3, 0x74, 0x39, 0xd0, 0x72, 0x82, 0xde, 0x97, 0xe4, 0x11, 0x54,
	0x38, 0x0e, 0xd0, 0x1d, 0xab, 0xd8, 0x85, 0x6b, 0x63, 0x21, 0x85, 0xef, 0x4b, 0xb2, 0x03, 0x4b,
	0x02, 0xd1, 0x8f, 0x02, 0xf5, 0x6b, 0x03, 0x8d, 0x08, 0xba, 0x2f, 0xdb, 0x9f, 0x74, 0x58, 0xce,
	0x3d, 0x2d, 0x79, 0x0c, 0x86, 0x92, 0x31, 0xa9, 0x4f, 0xa7, 0x94, 0x5b, 0x2a, 0x66, 0x63, 0xd6,
	0xa1, 0xbe, 0x92, 0x55, 0x20, 0x4f, 0xa0, 0x94, 0x7e, 0x32, 0xf2, 0xdf, 0xec, 0xc7, 0x4b, 0x29,
	0xae, 0xf8, 0x1a, 0x56, 0x81, 0x3c, 0x83, 0x4a, 0x17, 0x25, 0x4d, 0xfa, 0x22, 0xff, 0x4f, 0x81,
	0x73, 0x96, 0xa7, 0xb9, 0x31, 0xd3, 0xe3, 0x41, 0xb4, 0x92, 0xad, 0x02, 0xe9, 0xc0, 0x52, 0x17,
	0x65, 0x17, 0xd1, 0xff, 0x7d, 0x8e, 0x23, 0xa8, 0xe5, 0xc5, 0x4f, 0xb6, 0x66, 0x5b, 0xca, 0xad,
	0x85, 0xab, 0x1b, 0xbb, 0xa7, 0x91, 0x17, 0x50, 0xcd, 0x0a, 0x30, 0x5b, 0xd7, 0x9c, 0x55, 0x60,
	0xde, 0xf8, 0xb1, 0x6e, 0xad, 0x42, 0xa7, 0xf6, 0xba, 0x6a, 0xb7, 0xc2, 0x7e, 0x0a, 0xeb, 0x1b,
	0x71, 0x07, 0x3b, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xee, 0x6d, 0xb4, 0x67, 0xf0, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotificationsClient is the client API for Notifications service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationsClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*Notification, error)
	SetReceived(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SetSeen(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RetrieveStream(ctx context.Context, in *RetrieveStreamRequest, opts ...grpc.CallOption) (Notifications_RetrieveStreamClient, error)
	RetrieveMany(ctx context.Context, in *RetrieveManyRequest, opts ...grpc.CallOption) (*RetrieveManyResponse, error)
}

type notificationsClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationsClient(cc grpc.ClientConnInterface) NotificationsClient {
	return &notificationsClient{cc}
}

func (c *notificationsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/notifier.Notifications/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*Notification, error) {
	out := new(Notification)
	err := c.cc.Invoke(ctx, "/notifier.Notifications/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) SetReceived(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/notifier.Notifications/SetReceived", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) SetSeen(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/notifier.Notifications/SetSeen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) RetrieveStream(ctx context.Context, in *RetrieveStreamRequest, opts ...grpc.CallOption) (Notifications_RetrieveStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Notifications_serviceDesc.Streams[0], "/notifier.Notifications/RetrieveStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationsRetrieveStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Notifications_RetrieveStreamClient interface {
	Recv() (*Notification, error)
	grpc.ClientStream
}

type notificationsRetrieveStreamClient struct {
	grpc.ClientStream
}

func (x *notificationsRetrieveStreamClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *notificationsClient) RetrieveMany(ctx context.Context, in *RetrieveManyRequest, opts ...grpc.CallOption) (*RetrieveManyResponse, error) {
	out := new(RetrieveManyResponse)
	err := c.cc.Invoke(ctx, "/notifier.Notifications/RetrieveMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationsServer is the server API for Notifications service.
type NotificationsServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Retrieve(context.Context, *RetrieveRequest) (*Notification, error)
	SetReceived(context.Context, *UpdateStatusRequest) (*empty.Empty, error)
	SetSeen(context.Context, *UpdateStatusRequest) (*empty.Empty, error)
	RetrieveStream(*RetrieveStreamRequest, Notifications_RetrieveStreamServer) error
	RetrieveMany(context.Context, *RetrieveManyRequest) (*RetrieveManyResponse, error)
}

// UnimplementedNotificationsServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationsServer struct {
}

func (*UnimplementedNotificationsServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedNotificationsServer) Retrieve(ctx context.Context, req *RetrieveRequest) (*Notification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (*UnimplementedNotificationsServer) SetReceived(ctx context.Context, req *UpdateStatusRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetReceived not implemented")
}
func (*UnimplementedNotificationsServer) SetSeen(ctx context.Context, req *UpdateStatusRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSeen not implemented")
}
func (*UnimplementedNotificationsServer) RetrieveStream(req *RetrieveStreamRequest, srv Notifications_RetrieveStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method RetrieveStream not implemented")
}
func (*UnimplementedNotificationsServer) RetrieveMany(ctx context.Context, req *RetrieveManyRequest) (*RetrieveManyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveMany not implemented")
}

func RegisterNotificationsServer(s *grpc.Server, srv NotificationsServer) {
	s.RegisterService(&_Notifications_serviceDesc, srv)
}

func _Notifications_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifier.Notifications/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifier.Notifications/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).Retrieve(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_SetReceived_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).SetReceived(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifier.Notifications/SetReceived",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).SetReceived(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_SetSeen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).SetSeen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifier.Notifications/SetSeen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).SetSeen(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_RetrieveStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RetrieveStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationsServer).RetrieveStream(m, &notificationsRetrieveStreamServer{stream})
}

type Notifications_RetrieveStreamServer interface {
	Send(*Notification) error
	grpc.ServerStream
}

type notificationsRetrieveStreamServer struct {
	grpc.ServerStream
}

func (x *notificationsRetrieveStreamServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func _Notifications_RetrieveMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).RetrieveMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifier.Notifications/RetrieveMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).RetrieveMany(ctx, req.(*RetrieveManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notifications_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notifier.Notifications",
	HandlerType: (*NotificationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Notifications_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _Notifications_Retrieve_Handler,
		},
		{
			MethodName: "SetReceived",
			Handler:    _Notifications_SetReceived_Handler,
		},
		{
			MethodName: "SetSeen",
			Handler:    _Notifications_SetSeen_Handler,
		},
		{
			MethodName: "RetrieveMany",
			Handler:    _Notifications_RetrieveMany_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RetrieveStream",
			Handler:       _Notifications_RetrieveStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "notifier.proto",
}
