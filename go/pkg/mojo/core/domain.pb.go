// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/core/domain.proto

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

type Domain struct {
	Labels               []string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Domain) Reset()         { *m = Domain{} }
func (m *Domain) String() string { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()    {}
func (*Domain) Descriptor() ([]byte, []int) {
	return fileDescriptor_a039d3f244d54c32, []int{0}
}
func (m *Domain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Domain.Unmarshal(m, b)
}
func (m *Domain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Domain.Marshal(b, m, deterministic)
}
func (m *Domain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Domain.Merge(m, src)
}
func (m *Domain) XXX_Size() int {
	return xxx_messageInfo_Domain.Size(m)
}
func (m *Domain) XXX_DiscardUnknown() {
	xxx_messageInfo_Domain.DiscardUnknown(m)
}

var xxx_messageInfo_Domain proto.InternalMessageInfo

func (m *Domain) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Domain)(nil), "mojo.core.Domain")
}

func init() { proto.RegisterFile("mojo/core/domain.proto", fileDescriptor_a039d3f244d54c32) }

var fileDescriptor_a039d3f244d54c32 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0xcd, 0xcf, 0xca,
	0xd7, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0xd3, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x89, 0xeb, 0x81, 0xc4, 0x95, 0xcc, 0xb8, 0xd8, 0x5c, 0xc0, 0x52,
	0x42, 0x3a, 0x5c, 0x6c, 0x39, 0x89, 0x49, 0xa9, 0x39, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0x9c,
	0x4e, 0x22, 0xaf, 0xee, 0xc9, 0x0b, 0x40, 0x44, 0x74, 0xf2, 0x73, 0x33, 0x4b, 0x52, 0x73, 0x0b,
	0x4a, 0x2a, 0x83, 0xa0, 0x6a, 0x9c, 0x22, 0xb8, 0xc4, 0xf2, 0x8b, 0xd2, 0xf5, 0x40, 0x06, 0xe5,
	0x24, 0xe6, 0x41, 0x18, 0x60, 0x13, 0x9d, 0xb8, 0x21, 0xe6, 0x05, 0x80, 0x6c, 0x0a, 0x60, 0x8c,
	0xd2, 0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0x29, 0xd2, 0x05,
	0x29, 0x87, 0xb8, 0x29, 0x3d, 0x5f, 0xbf, 0x20, 0x3b, 0x5d, 0x1f, 0xee, 0x48, 0x6b, 0x10, 0x91,
	0xc4, 0x06, 0x76, 0xa3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xde, 0xa9, 0x38, 0xee, 0xbd, 0x00,
	0x00, 0x00,
}
