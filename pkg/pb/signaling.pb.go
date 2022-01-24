// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: signaling.proto

package pb

import (
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

type SubscribeOffersParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SharedKey []byte `protobuf:"bytes,1,opt,name=shared_key,json=sharedKey,proto3" json:"shared_key,omitempty"`
}

func (x *SubscribeOffersParams) Reset() {
	*x = SubscribeOffersParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signaling_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeOffersParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeOffersParams) ProtoMessage() {}

func (x *SubscribeOffersParams) ProtoReflect() protoreflect.Message {
	mi := &file_signaling_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeOffersParams.ProtoReflect.Descriptor instead.
func (*SubscribeOffersParams) Descriptor() ([]byte, []int) {
	return file_signaling_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeOffersParams) GetSharedKey() []byte {
	if x != nil {
		return x.SharedKey
	}
	return nil
}

type PublishOffersParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SharedKey []byte `protobuf:"bytes,1,opt,name=shared_key,json=sharedKey,proto3" json:"shared_key,omitempty"`
	Offer     *Offer `protobuf:"bytes,2,opt,name=offer,proto3" json:"offer,omitempty"`
}

func (x *PublishOffersParams) Reset() {
	*x = PublishOffersParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signaling_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishOffersParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishOffersParams) ProtoMessage() {}

func (x *PublishOffersParams) ProtoReflect() protoreflect.Message {
	mi := &file_signaling_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishOffersParams.ProtoReflect.Descriptor instead.
func (*PublishOffersParams) Descriptor() ([]byte, []int) {
	return file_signaling_proto_rawDescGZIP(), []int{1}
}

func (x *PublishOffersParams) GetSharedKey() []byte {
	if x != nil {
		return x.SharedKey
	}
	return nil
}

func (x *PublishOffersParams) GetOffer() *Offer {
	if x != nil {
		return x.Offer
	}
	return nil
}

var File_signaling_proto protoreflect.FileDescriptor

var file_signaling_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x77, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x36, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x22, 0x57, 0x0a, 0x13, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79,
	0x12, 0x21, 0x0a, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x32, 0x81, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e,
	0x67, 0x12, 0x3f, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x33, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x12, 0x19, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x06, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x72, 0x69, 0x61, 0x73, 0x63,
	0x2e, 0x65, 0x75, 0x2f, 0x77, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_signaling_proto_rawDescOnce sync.Once
	file_signaling_proto_rawDescData = file_signaling_proto_rawDesc
)

func file_signaling_proto_rawDescGZIP() []byte {
	file_signaling_proto_rawDescOnce.Do(func() {
		file_signaling_proto_rawDescData = protoimpl.X.CompressGZIP(file_signaling_proto_rawDescData)
	})
	return file_signaling_proto_rawDescData
}

var file_signaling_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_signaling_proto_goTypes = []interface{}{
	(*SubscribeOffersParams)(nil), // 0: wice.SubscribeOffersParams
	(*PublishOffersParams)(nil),   // 1: wice.PublishOffersParams
	(*Offer)(nil),                 // 2: wice.Offer
	(*Error)(nil),                 // 3: Error
}
var file_signaling_proto_depIdxs = []int32{
	2, // 0: wice.PublishOffersParams.offer:type_name -> wice.Offer
	0, // 1: wice.Signaling.SubscribeOffers:input_type -> wice.SubscribeOffersParams
	1, // 2: wice.Signaling.PublishOffer:input_type -> wice.PublishOffersParams
	2, // 3: wice.Signaling.SubscribeOffers:output_type -> wice.Offer
	3, // 4: wice.Signaling.PublishOffer:output_type -> Error
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_signaling_proto_init() }
func file_signaling_proto_init() {
	if File_signaling_proto != nil {
		return
	}
	file_common_proto_init()
	file_offer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_signaling_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeOffersParams); i {
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
		file_signaling_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishOffersParams); i {
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
			RawDescriptor: file_signaling_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_signaling_proto_goTypes,
		DependencyIndexes: file_signaling_proto_depIdxs,
		MessageInfos:      file_signaling_proto_msgTypes,
	}.Build()
	File_signaling_proto = out.File
	file_signaling_proto_rawDesc = nil
	file_signaling_proto_goTypes = nil
	file_signaling_proto_depIdxs = nil
}