// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: mojo/core/case_style.proto

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

type CaseStyle int32

const (
	CaseStyle_CASE_STYLE_UNSPECIFIED     CaseStyle = 0
	CaseStyle_CASE_STYLE_SNAKE           CaseStyle = 1
	CaseStyle_CASE_STYLE_SCREAMING_SNAKE CaseStyle = 2
	CaseStyle_CASE_STYLE_KEBAB           CaseStyle = 3
	CaseStyle_CASE_STYLE_SCREAMING_KEBAB CaseStyle = 4
	CaseStyle_CASE_STYLE_CAMEL           CaseStyle = 5
	CaseStyle_CASE_STYLE_LOWER_CAMEL     CaseStyle = 6
)

// Enum value maps for CaseStyle.
var (
	CaseStyle_name = map[int32]string{
		0: "CASE_STYLE_UNSPECIFIED",
		1: "CASE_STYLE_SNAKE",
		2: "CASE_STYLE_SCREAMING_SNAKE",
		3: "CASE_STYLE_KEBAB",
		4: "CASE_STYLE_SCREAMING_KEBAB",
		5: "CASE_STYLE_CAMEL",
		6: "CASE_STYLE_LOWER_CAMEL",
	}
	CaseStyle_value = map[string]int32{
		"CASE_STYLE_UNSPECIFIED":     0,
		"CASE_STYLE_SNAKE":           1,
		"CASE_STYLE_SCREAMING_SNAKE": 2,
		"CASE_STYLE_KEBAB":           3,
		"CASE_STYLE_SCREAMING_KEBAB": 4,
		"CASE_STYLE_CAMEL":           5,
		"CASE_STYLE_LOWER_CAMEL":     6,
	}
)

func (x CaseStyle) Enum() *CaseStyle {
	p := new(CaseStyle)
	*p = x
	return p
}

func (x CaseStyle) ToText() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CaseStyle) Descriptor() protoreflect.EnumDescriptor {
	return file_mojo_core_case_style_proto_enumTypes[0].Descriptor()
}

func (CaseStyle) Type() protoreflect.EnumType {
	return &file_mojo_core_case_style_proto_enumTypes[0]
}

func (x CaseStyle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CaseStyle.Descriptor instead.
func (CaseStyle) EnumDescriptor() ([]byte, []int) {
	return file_mojo_core_case_style_proto_rawDescGZIP(), []int{0}
}

var File_mojo_core_case_style_proto protoreflect.FileDescriptor

var file_mojo_core_case_style_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65,
	0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f,
	0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2a, 0xc5, 0x01, 0x0a, 0x09, 0x43, 0x61, 0x73, 0x65,
	0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x54,
	0x59, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x5f,
	0x53, 0x4e, 0x41, 0x4b, 0x45, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x41, 0x53, 0x45, 0x5f,
	0x53, 0x54, 0x59, 0x4c, 0x45, 0x5f, 0x53, 0x43, 0x52, 0x45, 0x41, 0x4d, 0x49, 0x4e, 0x47, 0x5f,
	0x53, 0x4e, 0x41, 0x4b, 0x45, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x53, 0x45, 0x5f,
	0x53, 0x54, 0x59, 0x4c, 0x45, 0x5f, 0x4b, 0x45, 0x42, 0x41, 0x42, 0x10, 0x03, 0x12, 0x1e, 0x0a,
	0x1a, 0x43, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x5f, 0x53, 0x43, 0x52, 0x45,
	0x41, 0x4d, 0x49, 0x4e, 0x47, 0x5f, 0x4b, 0x45, 0x42, 0x41, 0x42, 0x10, 0x04, 0x12, 0x14, 0x0a,
	0x10, 0x43, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x4d, 0x45,
	0x4c, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x59, 0x4c,
	0x45, 0x5f, 0x4c, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x4d, 0x45, 0x4c, 0x10, 0x06, 0x42,
	0x5b, 0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x0e, 0x43, 0x61, 0x73, 0x65, 0x53,
	0x74, 0x79, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f,
	0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_case_style_proto_rawDescOnce sync.Once
	file_mojo_core_case_style_proto_rawDescData = file_mojo_core_case_style_proto_rawDesc
)

func file_mojo_core_case_style_proto_rawDescGZIP() []byte {
	file_mojo_core_case_style_proto_rawDescOnce.Do(func() {
		file_mojo_core_case_style_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_case_style_proto_rawDescData)
	})
	return file_mojo_core_case_style_proto_rawDescData
}

var file_mojo_core_case_style_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mojo_core_case_style_proto_goTypes = []interface{}{
	(CaseStyle)(0), // 0: mojo.core.CaseStyle
}
var file_mojo_core_case_style_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mojo_core_case_style_proto_init() }
func file_mojo_core_case_style_proto_init() {
	if File_mojo_core_case_style_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mojo_core_case_style_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_case_style_proto_goTypes,
		DependencyIndexes: file_mojo_core_case_style_proto_depIdxs,
		EnumInfos:         file_mojo_core_case_style_proto_enumTypes,
	}.Build()
	File_mojo_core_case_style_proto = out.File
	file_mojo_core_case_style_proto_rawDesc = nil
	file_mojo_core_case_style_proto_goTypes = nil
	file_mojo_core_case_style_proto_depIdxs = nil
}
