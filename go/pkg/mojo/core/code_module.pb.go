// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: mojo/core/code_module.proto

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

type CodeModule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAddress     int64  `protobuf:"varint,1,opt,name=base_address,json=baseAddress,proto3" json:"baseAddress,omitempty"`
	Size            int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	CodeFile        string `protobuf:"bytes,3,opt,name=code_file,json=codeFile,proto3" json:"codeFile,omitempty"`
	CodeIdentifier  string `protobuf:"bytes,4,opt,name=code_identifier,json=codeIdentifier,proto3" json:"codeIdentifier,omitempty"`
	DebugFile       string `protobuf:"bytes,5,opt,name=debug_file,json=debugFile,proto3" json:"debugFile,omitempty"`
	DebugIdentifier string `protobuf:"bytes,6,opt,name=debug_identifier,json=debugIdentifier,proto3" json:"debugIdentifier,omitempty"`
	Version         string `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *CodeModule) Reset() {
	*x = CodeModule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_code_module_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeModule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeModule) ProtoMessage() {}

func (x *CodeModule) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_code_module_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeModule.ProtoReflect.Descriptor instead.
func (*CodeModule) Descriptor() ([]byte, []int) {
	return file_mojo_core_code_module_proto_rawDescGZIP(), []int{0}
}

func (x *CodeModule) GetBaseAddress() int64 {
	if x != nil {
		return x.BaseAddress
	}
	return 0
}

func (x *CodeModule) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CodeModule) GetCodeFile() string {
	if x != nil {
		return x.CodeFile
	}
	return ""
}

func (x *CodeModule) GetCodeIdentifier() string {
	if x != nil {
		return x.CodeIdentifier
	}
	return ""
}

func (x *CodeModule) GetDebugFile() string {
	if x != nil {
		return x.DebugFile
	}
	return ""
}

func (x *CodeModule) GetDebugIdentifier() string {
	if x != nil {
		return x.DebugIdentifier
	}
	return ""
}

func (x *CodeModule) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_mojo_core_code_module_proto protoreflect.FileDescriptor

var file_mojo_core_code_module_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d,
	0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x22, 0xed, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x64,
	0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62,
	0x61, 0x73, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x62, 0x75, 0x67, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64,
	0x65, 0x62, 0x75, 0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x5c, 0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e,
	0x6d, 0x6f, 0x6a, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x42, 0x0f, 0x43, 0x6f, 0x64, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_code_module_proto_rawDescOnce sync.Once
	file_mojo_core_code_module_proto_rawDescData = file_mojo_core_code_module_proto_rawDesc
)

func file_mojo_core_code_module_proto_rawDescGZIP() []byte {
	file_mojo_core_code_module_proto_rawDescOnce.Do(func() {
		file_mojo_core_code_module_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_code_module_proto_rawDescData)
	})
	return file_mojo_core_code_module_proto_rawDescData
}

var file_mojo_core_code_module_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mojo_core_code_module_proto_goTypes = []interface{}{
	(*CodeModule)(nil), // 0: mojo.core.CodeModule
}
var file_mojo_core_code_module_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mojo_core_code_module_proto_init() }
func file_mojo_core_code_module_proto_init() {
	if File_mojo_core_code_module_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mojo_core_code_module_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeModule); i {
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
			RawDescriptor: file_mojo_core_code_module_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_code_module_proto_goTypes,
		DependencyIndexes: file_mojo_core_code_module_proto_depIdxs,
		MessageInfos:      file_mojo_core_code_module_proto_msgTypes,
	}.Build()
	File_mojo_core_code_module_proto = out.File
	file_mojo_core_code_module_proto_rawDesc = nil
	file_mojo_core_code_module_proto_goTypes = nil
	file_mojo_core_code_module_proto_depIdxs = nil
}
