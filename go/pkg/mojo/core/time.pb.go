// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: mojo/core/time.proto

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

type Month int32

const (
	Month_MONTH_UNSPECIFIED Month = 0
	Month_MONTH_JANUARY     Month = 1
	Month_MONTH_FEBRUARY    Month = 2
	Month_MONTH_MARCH       Month = 3
	Month_MONTH_APRIL       Month = 4
	Month_MONTH_MAY         Month = 5
	Month_MONTH_JUNE        Month = 6
	Month_MONTH_JULY        Month = 7
	Month_MONTH_AUGUST      Month = 8
	Month_MONTH_SEPTEMBER   Month = 9
	Month_MONTH_OCTOBER     Month = 10
	Month_MONTH_NOVEMBER    Month = 11
	Month_MONTH_DECEMBER    Month = 12
)

// Enum value maps for Month.
var (
	Month_name = map[int32]string{
		0:  "MONTH_UNSPECIFIED",
		1:  "MONTH_JANUARY",
		2:  "MONTH_FEBRUARY",
		3:  "MONTH_MARCH",
		4:  "MONTH_APRIL",
		5:  "MONTH_MAY",
		6:  "MONTH_JUNE",
		7:  "MONTH_JULY",
		8:  "MONTH_AUGUST",
		9:  "MONTH_SEPTEMBER",
		10: "MONTH_OCTOBER",
		11: "MONTH_NOVEMBER",
		12: "MONTH_DECEMBER",
	}
	Month_value = map[string]int32{
		"MONTH_UNSPECIFIED": 0,
		"MONTH_JANUARY":     1,
		"MONTH_FEBRUARY":    2,
		"MONTH_MARCH":       3,
		"MONTH_APRIL":       4,
		"MONTH_MAY":         5,
		"MONTH_JUNE":        6,
		"MONTH_JULY":        7,
		"MONTH_AUGUST":      8,
		"MONTH_SEPTEMBER":   9,
		"MONTH_OCTOBER":     10,
		"MONTH_NOVEMBER":    11,
		"MONTH_DECEMBER":    12,
	}
)

func (x Month) Enum() *Month {
	p := new(Month)
	*p = x
	return p
}

func (x Month) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Month) Descriptor() protoreflect.EnumDescriptor {
	return file_mojo_core_time_proto_enumTypes[0].Descriptor()
}

func (Month) Type() protoreflect.EnumType {
	return &file_mojo_core_time_proto_enumTypes[0]
}

func (x Month) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Month.Descriptor instead.
func (Month) EnumDescriptor() ([]byte, []int) {
	return file_mojo_core_time_proto_rawDescGZIP(), []int{0}
}

type DayOfWeek int32

const (
	DayOfWeek_DAY_OF_WEEK_UNSPECIFIED DayOfWeek = 0
	DayOfWeek_DAY_OF_WEEK_MONDAY      DayOfWeek = 1
	DayOfWeek_DAY_OF_WEEK_TUESDAY     DayOfWeek = 2
	DayOfWeek_DAY_OF_WEEK_WEDNESDAY   DayOfWeek = 3
	DayOfWeek_DAY_OF_WEEK_THURSDAY    DayOfWeek = 4
	DayOfWeek_DAY_OF_WEEK_FRIDAY      DayOfWeek = 5
	DayOfWeek_DAY_OF_WEEK_SATURDAY    DayOfWeek = 6
	DayOfWeek_DAY_OF_WEEK_SUNDAY      DayOfWeek = 7
)

// Enum value maps for DayOfWeek.
var (
	DayOfWeek_name = map[int32]string{
		0: "DAY_OF_WEEK_UNSPECIFIED",
		1: "DAY_OF_WEEK_MONDAY",
		2: "DAY_OF_WEEK_TUESDAY",
		3: "DAY_OF_WEEK_WEDNESDAY",
		4: "DAY_OF_WEEK_THURSDAY",
		5: "DAY_OF_WEEK_FRIDAY",
		6: "DAY_OF_WEEK_SATURDAY",
		7: "DAY_OF_WEEK_SUNDAY",
	}
	DayOfWeek_value = map[string]int32{
		"DAY_OF_WEEK_UNSPECIFIED": 0,
		"DAY_OF_WEEK_MONDAY":      1,
		"DAY_OF_WEEK_TUESDAY":     2,
		"DAY_OF_WEEK_WEDNESDAY":   3,
		"DAY_OF_WEEK_THURSDAY":    4,
		"DAY_OF_WEEK_FRIDAY":      5,
		"DAY_OF_WEEK_SATURDAY":    6,
		"DAY_OF_WEEK_SUNDAY":      7,
	}
)

func (x DayOfWeek) Enum() *DayOfWeek {
	p := new(DayOfWeek)
	*p = x
	return p
}

func (x DayOfWeek) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DayOfWeek) Descriptor() protoreflect.EnumDescriptor {
	return file_mojo_core_time_proto_enumTypes[1].Descriptor()
}

func (DayOfWeek) Type() protoreflect.EnumType {
	return &file_mojo_core_time_proto_enumTypes[1]
}

func (x DayOfWeek) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DayOfWeek.Descriptor instead.
func (DayOfWeek) EnumDescriptor() ([]byte, []int) {
	return file_mojo_core_time_proto_rawDescGZIP(), []int{1}
}

type Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year  int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day   int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *Date) Reset() {
	*x = Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_time_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Date) ProtoMessage() {}

func (x *Date) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_time_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Date.ProtoReflect.Descriptor instead.
func (*Date) Descriptor() ([]byte, []int) {
	return file_mojo_core_time_proto_rawDescGZIP(), []int{0}
}

func (x *Date) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Date) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *Date) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

type DateTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year        int32     `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month       int32     `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day         int32     `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	Hour        int32     `protobuf:"varint,4,opt,name=hour,proto3" json:"hour,omitempty"`
	Minute      int32     `protobuf:"varint,5,opt,name=minute,proto3" json:"minute,omitempty"`
	Seconds     int64     `protobuf:"varint,6,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanoseconds int32     `protobuf:"varint,7,opt,name=nanoseconds,proto3" json:"nanoseconds,omitempty"`
	TimeZone    *TimeZone `protobuf:"bytes,15,opt,name=time_zone,json=timeZone,proto3" json:"timeZone,omitempty"`
}

func (x *DateTime) Reset() {
	*x = DateTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_time_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateTime) ProtoMessage() {}

func (x *DateTime) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_time_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateTime.ProtoReflect.Descriptor instead.
func (*DateTime) Descriptor() ([]byte, []int) {
	return file_mojo_core_time_proto_rawDescGZIP(), []int{1}
}

func (x *DateTime) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *DateTime) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *DateTime) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *DateTime) GetHour() int32 {
	if x != nil {
		return x.Hour
	}
	return 0
}

func (x *DateTime) GetMinute() int32 {
	if x != nil {
		return x.Minute
	}
	return 0
}

func (x *DateTime) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *DateTime) GetNanoseconds() int32 {
	if x != nil {
		return x.Nanoseconds
	}
	return 0
}

func (x *DateTime) GetTimeZone() *TimeZone {
	if x != nil {
		return x.TimeZone
	}
	return nil
}

type TimeZone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TimeZone) Reset() {
	*x = TimeZone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_time_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeZone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeZone) ProtoMessage() {}

func (x *TimeZone) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_time_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeZone.ProtoReflect.Descriptor instead.
func (*TimeZone) Descriptor() ([]byte, []int) {
	return file_mojo_core_time_proto_rawDescGZIP(), []int{2}
}

func (x *TimeZone) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *TimeZone) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Timestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seconds     int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanoseconds int32 `protobuf:"varint,2,opt,name=nanoseconds,proto3" json:"nanoseconds,omitempty"`
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_time_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_time_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_mojo_core_time_proto_rawDescGZIP(), []int{3}
}

func (x *Timestamp) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Timestamp) GetNanoseconds() int32 {
	if x != nil {
		return x.Nanoseconds
	}
	return 0
}

type Duration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seconds     int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanoseconds int32 `protobuf:"varint,2,opt,name=nanoseconds,proto3" json:"nanoseconds,omitempty"`
}

func (x *Duration) Reset() {
	*x = Duration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_time_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Duration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Duration) ProtoMessage() {}

func (x *Duration) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_time_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Duration.ProtoReflect.Descriptor instead.
func (*Duration) Descriptor() ([]byte, []int) {
	return file_mojo_core_time_proto_rawDescGZIP(), []int{4}
}

func (x *Duration) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Duration) GetNanoseconds() int32 {
	if x != nil {
		return x.Nanoseconds
	}
	return 0
}

type TimeOfDay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hours   int32 `protobuf:"varint,1,opt,name=hours,proto3" json:"hours,omitempty"`
	Minutes int32 `protobuf:"varint,2,opt,name=minutes,proto3" json:"minutes,omitempty"`
	Seconds int32 `protobuf:"varint,3,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos   int32 `protobuf:"varint,4,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *TimeOfDay) Reset() {
	*x = TimeOfDay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_time_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeOfDay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeOfDay) ProtoMessage() {}

func (x *TimeOfDay) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_time_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeOfDay.ProtoReflect.Descriptor instead.
func (*TimeOfDay) Descriptor() ([]byte, []int) {
	return file_mojo_core_time_proto_rawDescGZIP(), []int{5}
}

func (x *TimeOfDay) GetHours() int32 {
	if x != nil {
		return x.Hours
	}
	return 0
}

func (x *TimeOfDay) GetMinutes() int32 {
	if x != nil {
		return x.Minutes
	}
	return 0
}

func (x *TimeOfDay) GetSeconds() int32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *TimeOfDay) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

var File_mojo_core_time_proto protoreflect.FileDescriptor

var file_mojo_core_time_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x22, 0x42, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x64, 0x61, 0x79, 0x22, 0xe0, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f,
	0x75, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x61, 0x6e, 0x6f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x30, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a,
	0x6f, 0x6e, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x36, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65,
	0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x47, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x6e, 0x6f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x61,
	0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x46, 0x0a, 0x08, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x22, 0x6b, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x68,
	0x6f, 0x75, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x2a, 0xf8,
	0x01, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x4f, 0x4e, 0x54,
	0x48, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x11, 0x0a, 0x0d, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x5f, 0x4a, 0x41, 0x4e, 0x55, 0x41, 0x52, 0x59,
	0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x5f, 0x46, 0x45, 0x42, 0x52,
	0x55, 0x41, 0x52, 0x59, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x5f,
	0x4d, 0x41, 0x52, 0x43, 0x48, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x4f, 0x4e, 0x54, 0x48,
	0x5f, 0x41, 0x50, 0x52, 0x49, 0x4c, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x4f, 0x4e, 0x54,
	0x48, 0x5f, 0x4d, 0x41, 0x59, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x4f, 0x4e, 0x54, 0x48,
	0x5f, 0x4a, 0x55, 0x4e, 0x45, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x4f, 0x4e, 0x54, 0x48,
	0x5f, 0x4a, 0x55, 0x4c, 0x59, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x4f, 0x4e, 0x54, 0x48,
	0x5f, 0x41, 0x55, 0x47, 0x55, 0x53, 0x54, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x4f, 0x4e,
	0x54, 0x48, 0x5f, 0x53, 0x45, 0x50, 0x54, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x09, 0x12, 0x11,
	0x0a, 0x0d, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x5f, 0x4f, 0x43, 0x54, 0x4f, 0x42, 0x45, 0x52, 0x10,
	0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x5f, 0x4e, 0x4f, 0x56, 0x45, 0x4d,
	0x42, 0x45, 0x52, 0x10, 0x0b, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x5f, 0x44,
	0x45, 0x43, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x0c, 0x2a, 0xd8, 0x01, 0x0a, 0x09, 0x44, 0x61,
	0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x41, 0x59, 0x5f, 0x4f,
	0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57,
	0x45, 0x45, 0x4b, 0x5f, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x54, 0x55, 0x45, 0x53,
	0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f,
	0x57, 0x45, 0x45, 0x4b, 0x5f, 0x57, 0x45, 0x44, 0x4e, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x03,
	0x12, 0x18, 0x0a, 0x14, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f,
	0x54, 0x48, 0x55, 0x52, 0x53, 0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x41,
	0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x46, 0x52, 0x49, 0x44, 0x41, 0x59,
	0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45,
	0x4b, 0x5f, 0x53, 0x41, 0x54, 0x55, 0x52, 0x44, 0x41, 0x59, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12,
	0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x53, 0x55, 0x4e, 0x44,
	0x41, 0x59, 0x10, 0x07, 0x42, 0x56, 0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f,
	0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_time_proto_rawDescOnce sync.Once
	file_mojo_core_time_proto_rawDescData = file_mojo_core_time_proto_rawDesc
)

func file_mojo_core_time_proto_rawDescGZIP() []byte {
	file_mojo_core_time_proto_rawDescOnce.Do(func() {
		file_mojo_core_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_time_proto_rawDescData)
	})
	return file_mojo_core_time_proto_rawDescData
}

var file_mojo_core_time_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mojo_core_time_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mojo_core_time_proto_goTypes = []interface{}{
	(Month)(0),        // 0: mojo.core.Month
	(DayOfWeek)(0),    // 1: mojo.core.DayOfWeek
	(*Date)(nil),      // 2: mojo.core.Date
	(*DateTime)(nil),  // 3: mojo.core.DateTime
	(*TimeZone)(nil),  // 4: mojo.core.TimeZone
	(*Timestamp)(nil), // 5: mojo.core.Timestamp
	(*Duration)(nil),  // 6: mojo.core.Duration
	(*TimeOfDay)(nil), // 7: mojo.core.TimeOfDay
}
var file_mojo_core_time_proto_depIdxs = []int32{
	4, // 0: mojo.core.DateTime.time_zone:type_name -> mojo.core.TimeZone
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mojo_core_time_proto_init() }
func file_mojo_core_time_proto_init() {
	if File_mojo_core_time_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mojo_core_time_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Date); i {
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
		file_mojo_core_time_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateTime); i {
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
		file_mojo_core_time_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeZone); i {
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
		file_mojo_core_time_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timestamp); i {
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
		file_mojo_core_time_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Duration); i {
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
		file_mojo_core_time_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeOfDay); i {
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
			RawDescriptor: file_mojo_core_time_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_time_proto_goTypes,
		DependencyIndexes: file_mojo_core_time_proto_depIdxs,
		EnumInfos:         file_mojo_core_time_proto_enumTypes,
		MessageInfos:      file_mojo_core_time_proto_msgTypes,
	}.Build()
	File_mojo_core_time_proto = out.File
	file_mojo_core_time_proto_rawDesc = nil
	file_mojo_core_time_proto_goTypes = nil
	file_mojo_core_time_proto_depIdxs = nil
}
