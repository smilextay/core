// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/core/version.proto

package core

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Version struct {
	Major                uint64   `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor                uint64   `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch                uint64   `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	PreReleases          []string `protobuf:"bytes,5,rep,name=pre_releases,json=preReleases,proto3" json:"preReleases,omitempty"`
	Builds               []string `protobuf:"bytes,6,rep,name=builds,proto3" json:"builds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cad9cf0c4e8238b, []int{0}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetMajor() uint64 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *Version) GetMinor() uint64 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *Version) GetPatch() uint64 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func (m *Version) GetPreReleases() []string {
	if m != nil {
		return m.PreReleases
	}
	return nil
}

func (m *Version) GetBuilds() []string {
	if m != nil {
		return m.Builds
	}
	return nil
}

func init() {
	proto.RegisterType((*Version)(nil), "mojo.core.Version")
}

func init() { proto.RegisterFile("mojo/core/version.proto", fileDescriptor_7cad9cf0c4e8238b) }

var fileDescriptor_7cad9cf0c4e8238b = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x46, 0x15, 0x4a, 0x83, 0x6a, 0x2a, 0x81, 0xc2, 0x9f, 0x99, 0x52, 0x31, 0x15, 0xa9, 0xc4,
	0x03, 0x23, 0x4c, 0x79, 0x82, 0xaa, 0x03, 0x12, 0x2c, 0x28, 0x09, 0x56, 0xea, 0x12, 0xe7, 0x5a,
	0x37, 0x2e, 0x12, 0x2f, 0x9b, 0x07, 0xc8, 0x03, 0x30, 0xa3, 0x6b, 0x47, 0xe0, 0x2e, 0x96, 0x75,
	0xee, 0xf9, 0x96, 0xc3, 0x6e, 0x34, 0xec, 0x40, 0x54, 0x80, 0x52, 0x7c, 0x49, 0xec, 0x14, 0xb4,
	0x99, 0x41, 0xb0, 0x90, 0xcc, 0xe8, 0x90, 0xd1, 0xe1, 0xee, 0x27, 0x62, 0x27, 0x2f, 0xfe, 0x98,
	0xdc, 0xb3, 0xa9, 0x2e, 0x76, 0x80, 0x3c, 0x5a, 0x44, 0xcb, 0xe3, 0xfc, 0x62, 0xe8, 0xd3, 0x33,
	0x07, 0x56, 0xa0, 0x95, 0x95, 0xda, 0xd8, 0xef, 0x8d, 0x37, 0x9c, 0xaa, 0x5a, 0x40, 0x7e, 0x14,
	0xa8, 0x04, 0x0e, 0x54, 0x02, 0xa4, 0x9a, 0xc2, 0x56, 0x5b, 0x3e, 0xf9, 0x57, 0x1d, 0x08, 0x55,
	0x07, 0x92, 0x67, 0x36, 0x37, 0x28, 0xdf, 0x51, 0x36, 0xb2, 0xe8, 0x64, 0xc7, 0xa7, 0x8b, 0xc9,
	0x72, 0x96, 0xdf, 0x0e, 0x7d, 0x7a, 0x65, 0x50, 0x6e, 0x46, 0x1c, 0xec, 0x4e, 0x03, 0x9c, 0xac,
	0x58, 0x5c, 0xee, 0x55, 0xf3, 0xd1, 0xf1, 0xd8, 0xed, 0x2e, 0x87, 0x3e, 0x3d, 0xf7, 0x24, 0x98,
	0x8c, 0x4e, 0xfe, 0xca, 0xae, 0x01, 0xeb, 0x8c, 0x4a, 0x34, 0x45, 0xeb, 0x3f, 0x2e, 0x49, 0x3e,
	0x1f, 0x7b, 0xac, 0xa9, 0xd5, 0x3a, 0x7a, 0x13, 0xb5, 0xb2, 0xdb, 0x7d, 0x99, 0x55, 0xa0, 0x05,
	0x59, 0x0f, 0xe4, 0xfb, 0xac, 0x35, 0x08, 0xf3, 0x59, 0x8b, 0xbf, 0xce, 0x4f, 0xf4, 0x94, 0xb1,
	0xab, 0xfc, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xff, 0x4e, 0x8a, 0x10, 0x80, 0x01, 0x00, 0x00,
}
