// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: mojo/core/uuid.proto

package core

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

type Uuid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []byte `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Uuid) Reset() {
	*x = Uuid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_uuid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uuid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uuid) ProtoMessage() {}

func (x *Uuid) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_uuid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uuid.ProtoReflect.Descriptor instead.
func (*Uuid) Descriptor() ([]byte, []int) {
	return file_mojo_core_uuid_proto_rawDescGZIP(), []int{0}
}

func (x *Uuid) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

var File_mojo_core_uuid_proto protoreflect.FileDescriptor

var file_mojo_core_uuid_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x75, 0x69, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x22, 0x18, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x42, 0x56, 0x0a, 0x16, 0x6f,
	0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x09, 0x55, 0x75, 0x69, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b, 0x63,
	0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_uuid_proto_rawDescOnce sync.Once
	file_mojo_core_uuid_proto_rawDescData = file_mojo_core_uuid_proto_rawDesc
)

func file_mojo_core_uuid_proto_rawDescGZIP() []byte {
	file_mojo_core_uuid_proto_rawDescOnce.Do(func() {
		file_mojo_core_uuid_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_uuid_proto_rawDescData)
	})
	return file_mojo_core_uuid_proto_rawDescData
}

var file_mojo_core_uuid_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mojo_core_uuid_proto_goTypes = []interface{}{
	(*Uuid)(nil), // 0: mojo.core.Uuid
}
var file_mojo_core_uuid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mojo_core_uuid_proto_init() }
func file_mojo_core_uuid_proto_init() {
	if File_mojo_core_uuid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mojo_core_uuid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uuid); i {
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
			RawDescriptor: file_mojo_core_uuid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_uuid_proto_goTypes,
		DependencyIndexes: file_mojo_core_uuid_proto_depIdxs,
		MessageInfos:      file_mojo_core_uuid_proto_msgTypes,
	}.Build()
	File_mojo_core_uuid_proto = out.File
	file_mojo_core_uuid_proto_rawDesc = nil
	file_mojo_core_uuid_proto_goTypes = nil
	file_mojo_core_uuid_proto_depIdxs = nil
}
