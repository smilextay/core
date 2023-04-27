// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: mojo/core/boxed.proto

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

type BoolValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val bool `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *BoolValue) Reset() {
	*x = BoolValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolValue) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolValue) ProtoMessage() {}

func (x *BoolValue) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolValue.ProtoReflect.Descriptor instead.
func (*BoolValue) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{0}
}

func (x *BoolValue) GetVal() bool {
	if x != nil {
		return x.Val
	}
	return false
}

type Int32Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val int32 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Int32Value) Reset() {
	*x = Int32Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int32Value) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32Value) ProtoMessage() {}

func (x *Int32Value) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32Value.ProtoReflect.Descriptor instead.
func (*Int32Value) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{1}
}

func (x *Int32Value) GetVal() int32 {
	if x != nil {
		return x.Val
	}
	return 0
}

type UInt32Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val uint32 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *UInt32Value) Reset() {
	*x = UInt32Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UInt32Value) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UInt32Value) ProtoMessage() {}

func (x *UInt32Value) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UInt32Value.ProtoReflect.Descriptor instead.
func (*UInt32Value) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{2}
}

func (x *UInt32Value) GetVal() uint32 {
	if x != nil {
		return x.Val
	}
	return 0
}

type Int64Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val int64 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Int64Value) Reset() {
	*x = Int64Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64Value) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64Value) ProtoMessage() {}

func (x *Int64Value) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64Value.ProtoReflect.Descriptor instead.
func (*Int64Value) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{3}
}

func (x *Int64Value) GetVal() int64 {
	if x != nil {
		return x.Val
	}
	return 0
}

type UInt64Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val uint64 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *UInt64Value) Reset() {
	*x = UInt64Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UInt64Value) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UInt64Value) ProtoMessage() {}

func (x *UInt64Value) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UInt64Value.ProtoReflect.Descriptor instead.
func (*UInt64Value) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{4}
}

func (x *UInt64Value) GetVal() uint64 {
	if x != nil {
		return x.Val
	}
	return 0
}

type Float32Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val float32 `protobuf:"fixed32,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Float32Value) Reset() {
	*x = Float32Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Float32Value) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Float32Value) ProtoMessage() {}

func (x *Float32Value) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Float32Value.ProtoReflect.Descriptor instead.
func (*Float32Value) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{5}
}

func (x *Float32Value) GetVal() float32 {
	if x != nil {
		return x.Val
	}
	return 0
}

type Float64Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val float64 `protobuf:"fixed64,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Float64Value) Reset() {
	*x = Float64Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Float64Value) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Float64Value) ProtoMessage() {}

func (x *Float64Value) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Float64Value.ProtoReflect.Descriptor instead.
func (*Float64Value) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{6}
}

func (x *Float64Value) GetVal() float64 {
	if x != nil {
		return x.Val
	}
	return 0
}

type StringValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val string `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *StringValue) Reset() {
	*x = StringValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringValue) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringValue) ProtoMessage() {}

func (x *StringValue) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringValue.ProtoReflect.Descriptor instead.
func (*StringValue) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{7}
}

func (x *StringValue) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type BytesValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []byte `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *BytesValue) Reset() {
	*x = BytesValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesValue) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesValue) ProtoMessage() {}

func (x *BytesValue) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesValue.ProtoReflect.Descriptor instead.
func (*BytesValue) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{8}
}

func (x *BytesValue) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

type BoolValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals []bool `protobuf:"varint,1,rep,packed,name=vals,proto3" json:"vals,omitempty"`
}

func (x *BoolValues) Reset() {
	*x = BoolValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolValues) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolValues) ProtoMessage() {}

func (x *BoolValues) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolValues.ProtoReflect.Descriptor instead.
func (*BoolValues) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{9}
}

func (x *BoolValues) GetVals() []bool {
	if x != nil {
		return x.Vals
	}
	return nil
}

type Int32Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals []int32 `protobuf:"varint,1,rep,packed,name=vals,proto3" json:"vals,omitempty"`
}

func (x *Int32Values) Reset() {
	*x = Int32Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int32Values) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32Values) ProtoMessage() {}

func (x *Int32Values) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32Values.ProtoReflect.Descriptor instead.
func (*Int32Values) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{10}
}

func (x *Int32Values) GetVals() []int32 {
	if x != nil {
		return x.Vals
	}
	return nil
}

type UInt32Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals []uint32 `protobuf:"varint,1,rep,packed,name=vals,proto3" json:"vals,omitempty"`
}

func (x *UInt32Values) Reset() {
	*x = UInt32Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UInt32Values) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UInt32Values) ProtoMessage() {}

func (x *UInt32Values) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UInt32Values.ProtoReflect.Descriptor instead.
func (*UInt32Values) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{11}
}

func (x *UInt32Values) GetVals() []uint32 {
	if x != nil {
		return x.Vals
	}
	return nil
}

type Int64Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals []int64 `protobuf:"varint,1,rep,packed,name=vals,proto3" json:"vals,omitempty"`
}

func (x *Int64Values) Reset() {
	*x = Int64Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64Values) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64Values) ProtoMessage() {}

func (x *Int64Values) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64Values.ProtoReflect.Descriptor instead.
func (*Int64Values) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{12}
}

func (x *Int64Values) GetVals() []int64 {
	if x != nil {
		return x.Vals
	}
	return nil
}

type UInt64Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals []uint64 `protobuf:"varint,1,rep,packed,name=vals,proto3" json:"vals,omitempty"`
}

func (x *UInt64Values) Reset() {
	*x = UInt64Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UInt64Values) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UInt64Values) ProtoMessage() {}

func (x *UInt64Values) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UInt64Values.ProtoReflect.Descriptor instead.
func (*UInt64Values) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{13}
}

func (x *UInt64Values) GetVals() []uint64 {
	if x != nil {
		return x.Vals
	}
	return nil
}

type Float32Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals []float32 `protobuf:"fixed32,1,rep,packed,name=vals,proto3" json:"vals,omitempty"`
}

func (x *Float32Values) Reset() {
	*x = Float32Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Float32Values) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Float32Values) ProtoMessage() {}

func (x *Float32Values) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Float32Values.ProtoReflect.Descriptor instead.
func (*Float32Values) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{14}
}

func (x *Float32Values) GetVals() []float32 {
	if x != nil {
		return x.Vals
	}
	return nil
}

type Float64Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals []float64 `protobuf:"fixed64,1,rep,packed,name=vals,proto3" json:"vals,omitempty"`
}

func (x *Float64Values) Reset() {
	*x = Float64Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Float64Values) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Float64Values) ProtoMessage() {}

func (x *Float64Values) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Float64Values.ProtoReflect.Descriptor instead.
func (*Float64Values) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{15}
}

func (x *Float64Values) GetVals() []float64 {
	if x != nil {
		return x.Vals
	}
	return nil
}

type StringValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals []string `protobuf:"bytes,1,rep,name=vals,proto3" json:"vals,omitempty"`
}

func (x *StringValues) Reset() {
	*x = StringValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringValues) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringValues) ProtoMessage() {}

func (x *StringValues) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringValues.ProtoReflect.Descriptor instead.
func (*StringValues) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{16}
}

func (x *StringValues) GetVals() []string {
	if x != nil {
		return x.Vals
	}
	return nil
}

type StringMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals map[string]string `protobuf:"bytes,1,rep,name=vals,proto3" json:"vals,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *StringMap) Reset() {
	*x = StringMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMap) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMap) ProtoMessage() {}

func (x *StringMap) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMap.ProtoReflect.Descriptor instead.
func (*StringMap) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{17}
}

func (x *StringMap) GetVals() map[string]string {
	if x != nil {
		return x.Vals
	}
	return nil
}

type StringMultiMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals map[string]*StringValues `protobuf:"bytes,1,rep,name=vals,proto3" json:"vals,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *StringMultiMap) Reset() {
	*x = StringMultiMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_boxed_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMultiMap) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMultiMap) ProtoMessage() {}

func (x *StringMultiMap) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_boxed_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMultiMap.ProtoReflect.Descriptor instead.
func (*StringMultiMap) Descriptor() ([]byte, []int) {
	return file_mojo_core_boxed_proto_rawDescGZIP(), []int{18}
}

func (x *StringMultiMap) GetVals() map[string]*StringValues {
	if x != nil {
		return x.Vals
	}
	return nil
}

var File_mojo_core_boxed_proto protoreflect.FileDescriptor

var file_mojo_core_boxed_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x6f, 0x78, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0x1d, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x22, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x22, 0x1f, 0x0a, 0x0b, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x76,
	0x61, 0x6c, 0x22, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x76,
	0x61, 0x6c, 0x22, 0x1f, 0x0a, 0x0b, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x22, 0x20, 0x0a, 0x0c, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x20, 0x0a, 0x0c, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x1f, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x1e, 0x0a, 0x0a, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x20, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x08, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x22, 0x21, 0x0a, 0x0b, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x61, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x22, 0x22, 0x0a,
	0x0c, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x04, 0x76, 0x61, 0x6c,
	0x73, 0x22, 0x21, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04,
	0x76, 0x61, 0x6c, 0x73, 0x22, 0x22, 0x0a, 0x0c, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x04, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x22, 0x23, 0x0a, 0x0d, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x61, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x22, 0x23, 0x0a,
	0x0d, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x04, 0x76, 0x61,
	0x6c, 0x73, 0x22, 0x22, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x22, 0x78, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x4d, 0x61, 0x70, 0x12, 0x32, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x9b, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x4d, 0x61, 0x70, 0x12, 0x37, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x61, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x1a, 0x50, 0x0a, 0x09,
	0x56, 0x61, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x6a,
	0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x57,
	0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6d,
	0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x0a, 0x42, 0x6f, 0x78, 0x65, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_boxed_proto_rawDescOnce sync.Once
	file_mojo_core_boxed_proto_rawDescData = file_mojo_core_boxed_proto_rawDesc
)

func file_mojo_core_boxed_proto_rawDescGZIP() []byte {
	file_mojo_core_boxed_proto_rawDescOnce.Do(func() {
		file_mojo_core_boxed_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_boxed_proto_rawDescData)
	})
	return file_mojo_core_boxed_proto_rawDescData
}

var file_mojo_core_boxed_proto_msgTypes = make([]protoimpl.MessageInfo, 21)
var file_mojo_core_boxed_proto_goTypes = []interface{}{
	(*BoolValue)(nil),      // 0: mojo.core.BoolValue
	(*Int32Value)(nil),     // 1: mojo.core.Int32Value
	(*UInt32Value)(nil),    // 2: mojo.core.UInt32Value
	(*Int64Value)(nil),     // 3: mojo.core.Int64Value
	(*UInt64Value)(nil),    // 4: mojo.core.UInt64Value
	(*Float32Value)(nil),   // 5: mojo.core.Float32Value
	(*Float64Value)(nil),   // 6: mojo.core.Float64Value
	(*StringValue)(nil),    // 7: mojo.core.StringValue
	(*BytesValue)(nil),     // 8: mojo.core.BytesValue
	(*BoolValues)(nil),     // 9: mojo.core.BoolValues
	(*Int32Values)(nil),    // 10: mojo.core.Int32Values
	(*UInt32Values)(nil),   // 11: mojo.core.UInt32Values
	(*Int64Values)(nil),    // 12: mojo.core.Int64Values
	(*UInt64Values)(nil),   // 13: mojo.core.UInt64Values
	(*Float32Values)(nil),  // 14: mojo.core.Float32Values
	(*Float64Values)(nil),  // 15: mojo.core.Float64Values
	(*StringValues)(nil),   // 16: mojo.core.StringValues
	(*StringMap)(nil),      // 17: mojo.core.StringMap
	(*StringMultiMap)(nil), // 18: mojo.core.StringMultiMap
	nil,                    // 19: mojo.core.StringMap.ValsEntry
	nil,                    // 20: mojo.core.StringMultiMap.ValsEntry
}
var file_mojo_core_boxed_proto_depIdxs = []int32{
	19, // 0: mojo.core.StringMap.vals:type_name -> mojo.core.StringMap.ValsEntry
	20, // 1: mojo.core.StringMultiMap.vals:type_name -> mojo.core.StringMultiMap.ValsEntry
	16, // 2: mojo.core.StringMultiMap.ValsEntry.value:type_name -> mojo.core.StringValues
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_mojo_core_boxed_proto_init() }
func file_mojo_core_boxed_proto_init() {
	if File_mojo_core_boxed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mojo_core_boxed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolValue); i {
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
		file_mojo_core_boxed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int32Value); i {
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
		file_mojo_core_boxed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UInt32Value); i {
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
		file_mojo_core_boxed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64Value); i {
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
		file_mojo_core_boxed_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UInt64Value); i {
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
		file_mojo_core_boxed_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Float32Value); i {
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
		file_mojo_core_boxed_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Float64Value); i {
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
		file_mojo_core_boxed_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringValue); i {
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
		file_mojo_core_boxed_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesValue); i {
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
		file_mojo_core_boxed_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolValues); i {
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
		file_mojo_core_boxed_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int32Values); i {
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
		file_mojo_core_boxed_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UInt32Values); i {
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
		file_mojo_core_boxed_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64Values); i {
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
		file_mojo_core_boxed_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UInt64Values); i {
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
		file_mojo_core_boxed_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Float32Values); i {
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
		file_mojo_core_boxed_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Float64Values); i {
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
		file_mojo_core_boxed_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringValues); i {
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
		file_mojo_core_boxed_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMap); i {
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
		file_mojo_core_boxed_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMultiMap); i {
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
			RawDescriptor: file_mojo_core_boxed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   21,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_boxed_proto_goTypes,
		DependencyIndexes: file_mojo_core_boxed_proto_depIdxs,
		MessageInfos:      file_mojo_core_boxed_proto_msgTypes,
	}.Build()
	File_mojo_core_boxed_proto = out.File
	file_mojo_core_boxed_proto_rawDesc = nil
	file_mojo_core_boxed_proto_goTypes = nil
	file_mojo_core_boxed_proto_depIdxs = nil
}
