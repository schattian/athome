// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: checkout/reservations.proto

package reservations

import (
	checkout "github.com/athomecomar/athome/pb/checkout/checkout"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     uint64              `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Timestamp  *checkout.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Items      map[uint64]uint64   `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // { product_id: quantity }
	Days       uint64              `protobuf:"varint,4,opt,name=days,proto3" json:"days,omitempty"`
	Amount     float64             `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"` // calculated at req time
	MerchantId uint64              `protobuf:"varint,6,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_reservations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_reservations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_checkout_reservations_proto_rawDescGZIP(), []int{0}
}

func (x *Reservation) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Reservation) GetTimestamp() *checkout.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Reservation) GetItems() map[uint64]uint64 {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Reservation) GetDays() uint64 {
	if x != nil {
		return x.Days
	}
	return 0
}

func (x *Reservation) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Reservation) GetMerchantId() uint64 {
	if x != nil {
		return x.MerchantId
	}
	return 0
}

type RetrieveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationId uint64       `protobuf:"varint,1,opt,name=reservation_id,json=reservationId,proto3" json:"reservation_id,omitempty"`
	Reservation   *Reservation `protobuf:"bytes,2,opt,name=reservation,proto3" json:"reservation,omitempty"`
}

func (x *RetrieveResponse) Reset() {
	*x = RetrieveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_reservations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveResponse) ProtoMessage() {}

func (x *RetrieveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_reservations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveResponse.ProtoReflect.Descriptor instead.
func (*RetrieveResponse) Descriptor() ([]byte, []int) {
	return file_checkout_reservations_proto_rawDescGZIP(), []int{1}
}

func (x *RetrieveResponse) GetReservationId() uint64 {
	if x != nil {
		return x.ReservationId
	}
	return 0
}

func (x *RetrieveResponse) GetReservation() *Reservation {
	if x != nil {
		return x.Reservation
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ProductId   uint64 `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Days        uint64 `protobuf:"varint,3,opt,name=days,proto3" json:"days,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_reservations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_reservations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_checkout_reservations_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *CreateRequest) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CreateRequest) GetDays() uint64 {
	if x != nil {
		return x.Days
	}
	return 0
}

var File_checkout_reservations_proto protoreflect.FileDescriptor

var file_checkout_reservations_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x02,
	0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3a, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x1a, 0x38, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x76, 0x0a, 0x10,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x65, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x74, 0x68, 0x6f, 0x6d, 0x65,
	0x63, 0x6f, 0x6d, 0x61, 0x72, 0x2f, 0x61, 0x74, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_checkout_reservations_proto_rawDescOnce sync.Once
	file_checkout_reservations_proto_rawDescData = file_checkout_reservations_proto_rawDesc
)

func file_checkout_reservations_proto_rawDescGZIP() []byte {
	file_checkout_reservations_proto_rawDescOnce.Do(func() {
		file_checkout_reservations_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkout_reservations_proto_rawDescData)
	})
	return file_checkout_reservations_proto_rawDescData
}

var file_checkout_reservations_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_checkout_reservations_proto_goTypes = []interface{}{
	(*Reservation)(nil),        // 0: reservations.Reservation
	(*RetrieveResponse)(nil),   // 1: reservations.RetrieveResponse
	(*CreateRequest)(nil),      // 2: reservations.CreateRequest
	nil,                        // 3: reservations.Reservation.ItemsEntry
	(*checkout.Timestamp)(nil), // 4: checkout.Timestamp
}
var file_checkout_reservations_proto_depIdxs = []int32{
	4, // 0: reservations.Reservation.timestamp:type_name -> checkout.Timestamp
	3, // 1: reservations.Reservation.items:type_name -> reservations.Reservation.ItemsEntry
	0, // 2: reservations.RetrieveResponse.reservation:type_name -> reservations.Reservation
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_checkout_reservations_proto_init() }
func file_checkout_reservations_proto_init() {
	if File_checkout_reservations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_checkout_reservations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reservation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_checkout_reservations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_checkout_reservations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_checkout_reservations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_checkout_reservations_proto_goTypes,
		DependencyIndexes: file_checkout_reservations_proto_depIdxs,
		MessageInfos:      file_checkout_reservations_proto_msgTypes,
	}.Build()
	File_checkout_reservations_proto = out.File
	file_checkout_reservations_proto_rawDesc = nil
	file_checkout_reservations_proto_goTypes = nil
	file_checkout_reservations_proto_depIdxs = nil
}
