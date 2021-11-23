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
	Level                int32    `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
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

func (m *Version) GetLevel() int32 {
	if m != nil {
		return m.Level
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
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcd, 0x4e, 0x83, 0x40,
	0x14, 0x85, 0x43, 0x5b, 0x30, 0x1d, 0x9b, 0x68, 0xf0, 0x6f, 0x5c, 0x41, 0x5c, 0x61, 0x52, 0x99,
	0x85, 0x4b, 0x5d, 0xf1, 0x04, 0x4d, 0x17, 0x26, 0xba, 0x31, 0x80, 0x37, 0x94, 0x0a, 0xdc, 0xc9,
	0x30, 0x6d, 0xe2, 0xe3, 0xf8, 0x62, 0x3c, 0x00, 0x4f, 0x61, 0xee, 0x0c, 0xd1, 0x71, 0x33, 0x99,
	0x7c, 0xf7, 0x3b, 0x9b, 0x73, 0xd8, 0x4d, 0x8b, 0x7b, 0x14, 0x25, 0x2a, 0x10, 0x47, 0x50, 0x7d,
	0x8d, 0x5d, 0x2a, 0x15, 0x6a, 0x0c, 0x97, 0x74, 0x48, 0xe9, 0x70, 0xf7, 0x3d, 0x63, 0x27, 0x2f,
	0xf6, 0x18, 0xde, 0x33, 0xbf, 0xcd, 0xf7, 0xa8, 0xb8, 0x17, 0x7b, 0xc9, 0x22, 0xbb, 0x18, 0x87,
	0xe8, 0xcc, 0x80, 0x35, 0xb6, 0xb5, 0x86, 0x56, 0xea, 0xaf, 0xad, 0x35, 0x8c, 0x5a, 0x77, 0xa8,
	0xf8, 0xcc, 0x51, 0x09, 0xfc, 0x53, 0x09, 0x90, 0x2a, 0x73, 0x5d, 0xee, 0xf8, 0xfc, 0x4f, 0x35,
	0xc0, 0x55, 0x0d, 0x20, 0xb5, 0x81, 0x23, 0x34, 0x7c, 0x11, 0x7b, 0x89, 0x6f, 0x55, 0x03, 0x5c,
	0xd5, 0x80, 0xf0, 0x99, 0xad, 0xa4, 0x82, 0x77, 0x05, 0x0d, 0xe4, 0x3d, 0xf4, 0xdc, 0x8f, 0xe7,
	0xc9, 0x32, 0xbb, 0x1d, 0x87, 0xe8, 0x4a, 0x2a, 0xd8, 0x4e, 0xd8, 0xc9, 0x9d, 0x3a, 0x38, 0x5c,
	0xb3, 0xa0, 0x38, 0xd4, 0xcd, 0x47, 0xcf, 0x03, 0x93, 0xbb, 0x1c, 0x87, 0xe8, 0xdc, 0x12, 0x27,
	0x32, 0x39, 0xd9, 0x2b, 0xbb, 0x46, 0x55, 0xa5, 0x54, 0x5a, 0x93, 0x77, 0xf6, 0x63, 0xda, 0xcb,
	0x56, 0x53, 0x75, 0x1b, 0xaa, 0x75, 0xe3, 0xbd, 0x89, 0xaa, 0xd6, 0xbb, 0x43, 0x91, 0x96, 0xd8,
	0x0a, 0xb2, 0x1e, 0xc8, 0xb7, 0x0b, 0x54, 0x28, 0xe4, 0x67, 0x25, 0x7e, 0x27, 0x79, 0xa2, 0xa7,
	0x08, 0xcc, 0x20, 0x8f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x03, 0x4c, 0x82, 0xab, 0x01,
	0x00, 0x00,
}
