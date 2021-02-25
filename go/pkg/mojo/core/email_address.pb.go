// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/core/email_address.proto

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

type EmailAddress struct {
	LocalPart            string   `protobuf:"bytes,1,opt,name=local_part,json=localPart,proto3" json:"localPart,omitempty"`
	Domain               *Domain  `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailAddress) Reset()         { *m = EmailAddress{} }
func (m *EmailAddress) String() string { return proto.CompactTextString(m) }
func (*EmailAddress) ProtoMessage()    {}
func (*EmailAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e9334947d87342d, []int{0}
}
func (m *EmailAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailAddress.Unmarshal(m, b)
}
func (m *EmailAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailAddress.Marshal(b, m, deterministic)
}
func (m *EmailAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailAddress.Merge(m, src)
}
func (m *EmailAddress) XXX_Size() int {
	return xxx_messageInfo_EmailAddress.Size(m)
}
func (m *EmailAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailAddress.DiscardUnknown(m)
}

var xxx_messageInfo_EmailAddress proto.InternalMessageInfo

func (m *EmailAddress) GetLocalPart() string {
	if m != nil {
		return m.LocalPart
	}
	return ""
}

func (m *EmailAddress) GetDomain() *Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func init() {
	proto.RegisterType((*EmailAddress)(nil), "mojo.core.EmailAddress")
}

func init() { proto.RegisterFile("mojo/core/email_address.proto", fileDescriptor_5e9334947d87342d) }

var fileDescriptor_5e9334947d87342d = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0xcd, 0xcf, 0xca,
	0xd7, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0xcd, 0x4d, 0xcc, 0xcc, 0x89, 0x4f, 0x4c, 0x49, 0x29,
	0x4a, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x49, 0xeb, 0x81, 0xa4,
	0xa5, 0xc4, 0x10, 0x2a, 0x53, 0xf2, 0x73, 0x13, 0x33, 0xf3, 0x20, 0x4a, 0x94, 0xda, 0x19, 0xb9,
	0x78, 0x5c, 0x41, 0x5a, 0x1d, 0x21, 0x3a, 0x85, 0xcc, 0xb8, 0xb8, 0x72, 0xf2, 0x93, 0x13, 0x73,
	0xe2, 0x0b, 0x12, 0x8b, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0xc4, 0x5f, 0xdd, 0x93,
	0x17, 0x06, 0x8b, 0x06, 0x24, 0x16, 0x95, 0xe8, 0xe4, 0xe7, 0x66, 0x96, 0xa4, 0xe6, 0x16, 0x94,
	0x54, 0x06, 0x71, 0xc2, 0x05, 0x85, 0xec, 0xb9, 0xd8, 0x20, 0x06, 0x4b, 0x30, 0x29, 0x30, 0x6a,
	0x70, 0x1b, 0x09, 0xea, 0xc1, 0x2d, 0xd7, 0x73, 0x01, 0x4b, 0x38, 0x89, 0xbc, 0xba, 0x27, 0x2f,
	0x00, 0x51, 0x84, 0x64, 0x06, 0x54, 0x9b, 0x53, 0x1c, 0x97, 0x58, 0x7e, 0x51, 0x3a, 0x58, 0x57,
	0x4e, 0x62, 0x5e, 0x3a, 0x42, 0xbb, 0x93, 0x20, 0xb2, 0x03, 0x03, 0x40, 0xce, 0x0e, 0x60, 0x8c,
	0xd2, 0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0x29, 0xd5, 0x05,
	0x69, 0x82, 0x78, 0x30, 0x3d, 0x5f, 0xbf, 0x20, 0x3b, 0x5d, 0x1f, 0xee, 0x63, 0x6b, 0x10, 0x91,
	0xc4, 0x06, 0xf6, 0xb0, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x56, 0x49, 0x50, 0x0f, 0x34, 0x01,
	0x00, 0x00,
}
