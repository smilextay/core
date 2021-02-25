// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/core/status.proto

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

type Status struct {
	// Types that are valid to be assigned to Status:
	//	*Status_ErrorVal
	//	*Status_NullVal
	Status               isStatus_Status `protobuf_oneof:"status"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef46958bd731057c, []int{0}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

type isStatus_Status interface {
	isStatus_Status()
}

type Status_ErrorVal struct {
	ErrorVal *Error `protobuf:"bytes,1,opt,name=error_val,json=errorVal,proto3,oneof" json:"errorVal,omitempty"`
}
type Status_NullVal struct {
	NullVal *Null `protobuf:"bytes,2,opt,name=null_val,json=nullVal,proto3,oneof" json:"nullVal,omitempty"`
}

func (*Status_ErrorVal) isStatus_Status() {}
func (*Status_NullVal) isStatus_Status()  {}

func (m *Status) GetStatus() isStatus_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Status) GetErrorVal() *Error {
	if x, ok := m.GetStatus().(*Status_ErrorVal); ok {
		return x.ErrorVal
	}
	return nil
}

func (m *Status) GetNullVal() *Null {
	if x, ok := m.GetStatus().(*Status_NullVal); ok {
		return x.NullVal
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Status) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Status_ErrorVal)(nil),
		(*Status_NullVal)(nil),
	}
}

func init() {
	proto.RegisterType((*Status)(nil), "mojo.core.Status")
}

func init() { proto.RegisterFile("mojo/core/status.proto", fileDescriptor_ef46958bd731057c) }

var fileDescriptor_ef46958bd731057c = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0xcd, 0xcf, 0xca,
	0xd7, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x89, 0xeb, 0x81, 0xc4, 0xa5, 0x44, 0x11, 0x4a, 0x52, 0x8b, 0x8a,
	0xf2, 0x8b, 0x20, 0x2a, 0xa4, 0x44, 0x10, 0xc2, 0x79, 0xa5, 0x39, 0x39, 0x10, 0x51, 0xa5, 0x05,
	0x8c, 0x5c, 0x6c, 0xc1, 0x60, 0x83, 0x84, 0xdc, 0xb9, 0x38, 0xc1, 0xea, 0xe3, 0xcb, 0x12, 0x73,
	0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x04, 0xf4, 0xe0, 0xc6, 0xea, 0xb9, 0x82, 0xe4, 0x9c,
	0xc4, 0x5e, 0xdd, 0x93, 0x17, 0x02, 0x2b, 0x0b, 0x4b, 0xcc, 0xd1, 0xc9, 0xcf, 0xcd, 0x2c, 0x49,
	0xcd, 0x2d, 0x28, 0xa9, 0xf4, 0x60, 0x08, 0xe2, 0x80, 0x89, 0x0a, 0x39, 0x73, 0x71, 0x80, 0x6c,
	0x00, 0x9b, 0xc3, 0x04, 0x36, 0x87, 0x1f, 0xc9, 0x1c, 0xbf, 0xd2, 0x9c, 0x1c, 0x27, 0xd1, 0x57,
	0xf7, 0xe4, 0x05, 0x41, 0x8a, 0xd0, 0x4d, 0x61, 0x87, 0x0a, 0x3a, 0x71, 0x70, 0xb1, 0x41, 0x3c,
	0xe8, 0x14, 0xc1, 0x25, 0x96, 0x5f, 0x94, 0x0e, 0x36, 0x21, 0x27, 0x31, 0x2f, 0x1d, 0x61, 0x94,
	0x13, 0x37, 0xc4, 0xe5, 0x01, 0x20, 0x9f, 0x04, 0x30, 0x46, 0xe9, 0xa7, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x83, 0x14, 0xe9, 0x82, 0x94, 0x43, 0x7c, 0x9c, 0x9e, 0xaf,
	0x5f, 0x90, 0x9d, 0xae, 0x0f, 0x0f, 0x02, 0x6b, 0x10, 0x91, 0xc4, 0x06, 0x0e, 0x03, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x80, 0x1d, 0xa9, 0x55, 0x01, 0x00, 0x00,
}
