// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: mojo/core/ordering.proto

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

type Ordering_Sort int32

const (
	Ordering_SORT_UNSPECIFIED Ordering_Sort = 0
	Ordering_SORT_ASC         Ordering_Sort = 1
	Ordering_SORT_DESC        Ordering_Sort = 2
)

// Enum value maps for Ordering_Sort.
var (
	Ordering_Sort_name = map[int32]string{
		0: "SORT_UNSPECIFIED",
		1: "SORT_ASC",
		2: "SORT_DESC",
	}
	Ordering_Sort_value = map[string]int32{
		"SORT_UNSPECIFIED": 0,
		"SORT_ASC":         1,
		"SORT_DESC":        2,
	}
)

func (x Ordering_Sort) Enum() *Ordering_Sort {
	p := new(Ordering_Sort)
	*p = x
	return p
}

func (x Ordering_Sort) Text() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ordering_Sort) Descriptor() protoreflect.EnumDescriptor {
	return file_mojo_core_ordering_proto_enumTypes[0].Descriptor()
}

func (Ordering_Sort) Type() protoreflect.EnumType {
	return &file_mojo_core_ordering_proto_enumTypes[0]
}

func (x Ordering_Sort) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ordering_Sort.Descriptor instead.
func (Ordering_Sort) EnumDescriptor() ([]byte, []int) {
	return file_mojo_core_ordering_proto_rawDescGZIP(), []int{0, 0}
}

type Ordering struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals []*Ordering_Value `protobuf:"bytes,1,rep,name=vals,proto3" json:"vals,omitempty"`
}

func (x *Ordering) Reset() {
	*x = Ordering{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_ordering_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ordering) Text() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ordering) ProtoMessage() {}

func (x *Ordering) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_ordering_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ordering.ProtoReflect.Descriptor instead.
func (*Ordering) Descriptor() ([]byte, []int) {
	return file_mojo_core_ordering_proto_rawDescGZIP(), []int{0}
}

func (x *Ordering) GetVals() []*Ordering_Value {
	if x != nil {
		return x.Vals
	}
	return nil
}

type Ordering_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string        `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Sort  Ordering_Sort `protobuf:"varint,2,opt,name=sort,proto3,enum=mojo.core.Ordering_Sort" json:"sort,omitempty"`
}

func (x *Ordering_Value) Reset() {
	*x = Ordering_Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_ordering_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ordering_Value) Text() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ordering_Value) ProtoMessage() {}

func (x *Ordering_Value) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_ordering_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ordering_Value.ProtoReflect.Descriptor instead.
func (*Ordering_Value) Descriptor() ([]byte, []int) {
	return file_mojo_core_ordering_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Ordering_Value) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Ordering_Value) GetSort() Ordering_Sort {
	if x != nil {
		return x.Sort
	}
	return Ordering_SORT_UNSPECIFIED
}

var File_mojo_core_ordering_proto protoreflect.FileDescriptor

var file_mojo_core_ordering_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x6a, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x22, 0xc1, 0x01, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x2d, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x76, 0x61, 0x6c,
	0x73, 0x1a, 0x4b, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x2c, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x39,
	0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x4f, 0x52, 0x54, 0x5f, 0x41, 0x53, 0x43, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x4f,
	0x52, 0x54, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x02, 0x42, 0x5a, 0x0a, 0x16, 0x6f, 0x72, 0x67,
	0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x42, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_ordering_proto_rawDescOnce sync.Once
	file_mojo_core_ordering_proto_rawDescData = file_mojo_core_ordering_proto_rawDesc
)

func file_mojo_core_ordering_proto_rawDescGZIP() []byte {
	file_mojo_core_ordering_proto_rawDescOnce.Do(func() {
		file_mojo_core_ordering_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_ordering_proto_rawDescData)
	})
	return file_mojo_core_ordering_proto_rawDescData
}

var file_mojo_core_ordering_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mojo_core_ordering_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mojo_core_ordering_proto_goTypes = []interface{}{
	(Ordering_Sort)(0),     // 0: mojo.core.Ordering.Sort
	(*Ordering)(nil),       // 1: mojo.core.Ordering
	(*Ordering_Value)(nil), // 2: mojo.core.Ordering.Value
}
var file_mojo_core_ordering_proto_depIdxs = []int32{
	2, // 0: mojo.core.Ordering.vals:type_name -> mojo.core.Ordering.Value
	0, // 1: mojo.core.Ordering.Value.sort:type_name -> mojo.core.Ordering.Sort
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mojo_core_ordering_proto_init() }
func file_mojo_core_ordering_proto_init() {
	if File_mojo_core_ordering_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mojo_core_ordering_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ordering); i {
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
		file_mojo_core_ordering_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ordering_Value); i {
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
			RawDescriptor: file_mojo_core_ordering_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_ordering_proto_goTypes,
		DependencyIndexes: file_mojo_core_ordering_proto_depIdxs,
		EnumInfos:         file_mojo_core_ordering_proto_enumTypes,
		MessageInfos:      file_mojo_core_ordering_proto_msgTypes,
	}.Build()
	File_mojo_core_ordering_proto = out.File
	file_mojo_core_ordering_proto_rawDesc = nil
	file_mojo_core_ordering_proto_goTypes = nil
	file_mojo_core_ordering_proto_depIdxs = nil
}
