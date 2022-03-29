// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/core/target_attribute.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type DeclType int32

const (
	DeclType_DECL_TYPE_UNSPECIFIED DeclType = 0
	DeclType_DECL_TYPE_TYPE        DeclType = 1
	DeclType_DECL_TYPE_VALUE       DeclType = 2
	DeclType_DECL_TYPE_FUNCTION    DeclType = 3
	DeclType_DECL_TYPE_CONSTRUCTOR DeclType = 4
	DeclType_DECL_TYPE_ATTRIBUTE   DeclType = 5
	DeclType_DECL_TYPE_PACKAGE     DeclType = 6
	DeclType_DECL_TYPE_GENERIC     DeclType = 7
)

var DeclType_name = map[int32]string{
	0: "DECL_TYPE_UNSPECIFIED",
	1: "DECL_TYPE_TYPE",
	2: "DECL_TYPE_VALUE",
	3: "DECL_TYPE_FUNCTION",
	4: "DECL_TYPE_CONSTRUCTOR",
	5: "DECL_TYPE_ATTRIBUTE",
	6: "DECL_TYPE_PACKAGE",
	7: "DECL_TYPE_GENERIC",
}

var DeclType_value = map[string]int32{
	"DECL_TYPE_UNSPECIFIED": 0,
	"DECL_TYPE_TYPE":        1,
	"DECL_TYPE_VALUE":       2,
	"DECL_TYPE_FUNCTION":    3,
	"DECL_TYPE_CONSTRUCTOR": 4,
	"DECL_TYPE_ATTRIBUTE":   5,
	"DECL_TYPE_PACKAGE":     6,
	"DECL_TYPE_GENERIC":     7,
}

func (x DeclType) String() string {
	return proto.EnumName(DeclType_name, int32(x))
}

func (DeclType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d4b300abcdb30f4, []int{0}
}

func init() {
	proto.RegisterEnum("mojo.core.DeclType", DeclType_name, DeclType_value)
}

func init() { proto.RegisterFile("mojo/core/target_attribute.proto", fileDescriptor_1d4b300abcdb30f4) }

var fileDescriptor_1d4b300abcdb30f4 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x1c, 0xc4, 0x1b, 0x3e, 0x0a, 0x78, 0x00, 0xe3, 0xd2, 0x22, 0x16, 0xc4, 0x8c, 0x44, 0x3c, 0x30,
	0x32, 0x25, 0x8e, 0x5b, 0x45, 0x54, 0x8e, 0x95, 0xda, 0x48, 0xb0, 0x44, 0x49, 0x64, 0x99, 0x42,
	0x8b, 0x23, 0xcb, 0x1d, 0x78, 0x13, 0x1e, 0x89, 0xf7, 0xe0, 0x45, 0x90, 0x53, 0x81, 0xd5, 0xe5,
	0x64, 0xdd, 0x9d, 0xff, 0xd2, 0xfd, 0xc0, 0xcd, 0xda, 0xbc, 0x19, 0xdc, 0x1a, 0xab, 0xb0, 0xab,
	0xad, 0x56, 0xae, 0xaa, 0x9d, 0xb3, 0xcb, 0x66, 0xe3, 0x54, 0xdc, 0x59, 0xe3, 0x0c, 0x3a, 0xf1,
	0x8d, 0xd8, 0x37, 0x6e, 0xbf, 0x23, 0x70, 0x9c, 0xa9, 0x76, 0x25, 0x3e, 0x3b, 0x85, 0xae, 0xc0,
	0x38, 0xa3, 0x64, 0x5e, 0x89, 0x67, 0x4e, 0x2b, 0xc9, 0x16, 0x9c, 0x92, 0x7c, 0x9a, 0xd3, 0x0c,
	0x0e, 0x10, 0x02, 0xa7, 0x21, 0xf2, 0x02, 0x23, 0x34, 0x02, 0x67, 0xc1, 0x7b, 0x4a, 0xe6, 0x92,
	0xc2, 0x3d, 0x34, 0x01, 0x28, 0x98, 0x53, 0xc9, 0x88, 0xc8, 0x0b, 0x06, 0xf7, 0x77, 0x6f, 0x93,
	0x82, 0x2d, 0x44, 0x29, 0x89, 0x28, 0x4a, 0x78, 0x80, 0x2e, 0xc1, 0x28, 0x44, 0x89, 0x10, 0x65,
	0x9e, 0x4a, 0x41, 0xe1, 0x21, 0x1a, 0x83, 0xf3, 0x10, 0xf0, 0x84, 0x3c, 0x26, 0x33, 0x0a, 0x87,
	0xbb, 0xf6, 0x8c, 0x32, 0x5a, 0xe6, 0x04, 0x1e, 0xa5, 0xea, 0xeb, 0xe7, 0x7a, 0x00, 0x26, 0xc6,
	0xea, 0xd8, 0xef, 0x5b, 0xd5, 0x1f, 0xdb, 0x47, 0x3f, 0x34, 0xbd, 0x10, 0x3d, 0x8b, 0xe4, 0x0f,
	0x05, 0xf7, 0x24, 0x78, 0xf4, 0x82, 0xf5, 0xd2, 0xbd, 0x6e, 0x9a, 0xb8, 0x35, 0x6b, 0xec, 0xdb,
	0x77, 0xfe, 0xdf, 0x96, 0x9e, 0x36, 0xb8, 0x7b, 0xd7, 0xf8, 0x1f, 0xe7, 0x83, 0x97, 0x66, 0xd8,
	0x33, 0xbc, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x84, 0x3e, 0xdd, 0x67, 0x01, 0x00, 0x00,
}
