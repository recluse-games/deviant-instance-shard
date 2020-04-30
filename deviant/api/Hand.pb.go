// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: Hand.proto

package Deviant

import (
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

type Hand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cards []*Card `protobuf:"bytes,2,rep,name=cards,proto3" json:"cards,omitempty"`
}

func (x *Hand) Reset() {
	*x = Hand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Hand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hand) ProtoMessage() {}

func (x *Hand) ProtoReflect() protoreflect.Message {
	mi := &file_Hand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hand.ProtoReflect.Descriptor instead.
func (*Hand) Descriptor() ([]byte, []int) {
	return file_Hand_proto_rawDescGZIP(), []int{0}
}

func (x *Hand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Hand) GetCards() []*Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

var File_Hand_proto protoreflect.FileDescriptor

var file_Hand_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x48, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x44, 0x65,
	0x76, 0x69, 0x61, 0x6e, 0x74, 0x1a, 0x0a, 0x43, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3b, 0x0a, 0x04, 0x48, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x05, 0x63, 0x61, 0x72,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61,
	0x6e, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Hand_proto_rawDescOnce sync.Once
	file_Hand_proto_rawDescData = file_Hand_proto_rawDesc
)

func file_Hand_proto_rawDescGZIP() []byte {
	file_Hand_proto_rawDescOnce.Do(func() {
		file_Hand_proto_rawDescData = protoimpl.X.CompressGZIP(file_Hand_proto_rawDescData)
	})
	return file_Hand_proto_rawDescData
}

var file_Hand_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Hand_proto_goTypes = []interface{}{
	(*Hand)(nil), // 0: Deviant.Hand
	(*Card)(nil), // 1: Deviant.Card
}
var file_Hand_proto_depIdxs = []int32{
	1, // 0: Deviant.Hand.cards:type_name -> Deviant.Card
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Hand_proto_init() }
func file_Hand_proto_init() {
	if File_Hand_proto != nil {
		return
	}
	file_Card_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Hand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hand); i {
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
			RawDescriptor: file_Hand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Hand_proto_goTypes,
		DependencyIndexes: file_Hand_proto_depIdxs,
		MessageInfos:      file_Hand_proto_msgTypes,
	}.Build()
	File_Hand_proto = out.File
	file_Hand_proto_rawDesc = nil
	file_Hand_proto_goTypes = nil
	file_Hand_proto_depIdxs = nil
}
