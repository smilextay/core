// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: mojo/core/file.proto

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

type File_Mode int32

const (
	File_MODE_UNSPECIFIED File_Mode = 0
	File_MODE_DIR         File_Mode = 10
)

// Enum value maps for File_Mode.
var (
	File_Mode_name = map[int32]string{
		0:  "MODE_UNSPECIFIED",
		10: "MODE_DIR",
	}
	File_Mode_value = map[string]int32{
		"MODE_UNSPECIFIED": 0,
		"MODE_DIR":         10,
	}
)

func (x File_Mode) Enum() *File_Mode {
	p := new(File_Mode)
	*p = x
	return p
}

func (x File_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (File_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_mojo_core_file_proto_enumTypes[0].Descriptor()
}

func (File_Mode) Type() protoreflect.EnumType {
	return &file_mojo_core_file_proto_enumTypes[0]
}

func (x File_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use File_Mode.Descriptor instead.
func (File_Mode) EnumDescriptor() ([]byte, []int) {
	return file_mojo_core_file_proto_rawDescGZIP(), []int{0, 0}
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsDir bool       `protobuf:"varint,2,opt,name=is_dir,json=isDir,proto3" json:"isDir,omitempty"`
	Mode  File_Mode  `protobuf:"varint,5,opt,name=mode,proto3,enum=mojo.core.File_Mode" json:"mode,omitempty"`
	Info  *File_Info `protobuf:"bytes,6,opt,name=info,proto3" json:"info,omitempty"`
	Files []*File    `protobuf:"bytes,15,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_mojo_core_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetIsDir() bool {
	if x != nil {
		return x.IsDir
	}
	return false
}

func (x *File) GetMode() File_Mode {
	if x != nil {
		return x.Mode
	}
	return File_MODE_UNSPECIFIED
}

func (x *File) GetInfo() *File_Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *File) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

type File_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Suffix     string     `protobuf:"bytes,2,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Size       int64      `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	ChangeTime *Timestamp `protobuf:"bytes,5,opt,name=change_time,json=changeTime,proto3" json:"changeTime,omitempty"`
	ModifyTime *Timestamp `protobuf:"bytes,6,opt,name=modify_time,json=modifyTime,proto3" json:"modifyTime,omitempty"`
}

func (x *File_Info) Reset() {
	*x = File_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File_Info) ProtoMessage() {}

func (x *File_Info) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File_Info.ProtoReflect.Descriptor instead.
func (*File_Info) Descriptor() ([]byte, []int) {
	return file_mojo_core_file_proto_rawDescGZIP(), []int{0, 0}
}

func (x *File_Info) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File_Info) GetSuffix() string {
	if x != nil {
		return x.Suffix
	}
	return ""
}

func (x *File_Info) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File_Info) GetChangeTime() *Timestamp {
	if x != nil {
		return x.ChangeTime
	}
	return nil
}

func (x *File_Info) GetModifyTime() *Timestamp {
	if x != nil {
		return x.ModifyTime
	}
	return nil
}

var File_mojo_core_file_proto protoreflect.FileDescriptor

var file_mojo_core_file_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x1a, 0x14, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x03, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x44, 0x69, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12,
	0x25, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0xb4, 0x01, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x35, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f,
	0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2a, 0x0a,
	0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d,
	0x4f, 0x44, 0x45, 0x5f, 0x44, 0x49, 0x52, 0x10, 0x0a, 0x42, 0x56, 0x0a, 0x16, 0x6f, 0x72, 0x67,
	0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x42, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6a,
	0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b, 0x63, 0x6f, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_file_proto_rawDescOnce sync.Once
	file_mojo_core_file_proto_rawDescData = file_mojo_core_file_proto_rawDesc
)

func file_mojo_core_file_proto_rawDescGZIP() []byte {
	file_mojo_core_file_proto_rawDescOnce.Do(func() {
		file_mojo_core_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_file_proto_rawDescData)
	})
	return file_mojo_core_file_proto_rawDescData
}

var file_mojo_core_file_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mojo_core_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mojo_core_file_proto_goTypes = []interface{}{
	(File_Mode)(0),    // 0: mojo.core.File.Mode
	(*File)(nil),      // 1: mojo.core.File
	(*File_Info)(nil), // 2: mojo.core.File.Info
	(*Timestamp)(nil), // 3: mojo.core.Timestamp
}
var file_mojo_core_file_proto_depIdxs = []int32{
	0, // 0: mojo.core.File.mode:type_name -> mojo.core.File.Mode
	2, // 1: mojo.core.File.info:type_name -> mojo.core.File.Info
	1, // 2: mojo.core.File.files:type_name -> mojo.core.File
	3, // 3: mojo.core.File.Info.change_time:type_name -> mojo.core.Timestamp
	3, // 4: mojo.core.File.Info.modify_time:type_name -> mojo.core.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_mojo_core_file_proto_init() }
func file_mojo_core_file_proto_init() {
	if File_mojo_core_file_proto != nil {
		return
	}
	file_mojo_core_time_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mojo_core_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_mojo_core_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File_Info); i {
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
			RawDescriptor: file_mojo_core_file_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_file_proto_goTypes,
		DependencyIndexes: file_mojo_core_file_proto_depIdxs,
		EnumInfos:         file_mojo_core_file_proto_enumTypes,
		MessageInfos:      file_mojo_core_file_proto_msgTypes,
	}.Build()
	File_mojo_core_file_proto = out.File
	file_mojo_core_file_proto_rawDesc = nil
	file_mojo_core_file_proto_goTypes = nil
	file_mojo_core_file_proto_depIdxs = nil
}
