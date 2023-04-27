// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: mojo/core/url.proto

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

type Url struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scheme    string         `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Authority *Url_Authority `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	Path      string         `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Query     *Url_Query     `protobuf:"bytes,5,opt,name=query,proto3" json:"query,omitempty"`
	Fragment  string         `protobuf:"bytes,7,opt,name=fragment,proto3" json:"fragment,omitempty"`
}

func (x *Url) Reset() {
	*x = Url{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_url_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Url) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Url) ProtoMessage() {}

func (x *Url) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_url_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url.ProtoReflect.Descriptor instead.
func (*Url) Descriptor() ([]byte, []int) {
	return file_mojo_core_url_proto_rawDescGZIP(), []int{0}
}

func (x *Url) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *Url) GetAuthority() *Url_Authority {
	if x != nil {
		return x.Authority
	}
	return nil
}

func (x *Url) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Url) GetQuery() *Url_Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *Url) GetFragment() string {
	if x != nil {
		return x.Fragment
	}
	return ""
}

type Url_Authority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo string `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"userInfo,omitempty"`
	Host     string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port     int64  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Url_Authority) Reset() {
	*x = Url_Authority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_url_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Url_Authority) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Url_Authority) ProtoMessage() {}

func (x *Url_Authority) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_url_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url_Authority.ProtoReflect.Descriptor instead.
func (*Url_Authority) Descriptor() ([]byte, []int) {
	return file_mojo_core_url_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Url_Authority) GetUserInfo() string {
	if x != nil {
		return x.UserInfo
	}
	return ""
}

func (x *Url_Authority) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Url_Authority) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

type Url_Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals map[string]*StringValues `protobuf:"bytes,1,rep,name=vals,proto3" json:"vals,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Url_Query) Reset() {
	*x = Url_Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_url_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Url_Query) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Url_Query) ProtoMessage() {}

func (x *Url_Query) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_url_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url_Query.ProtoReflect.Descriptor instead.
func (*Url_Query) Descriptor() ([]byte, []int) {
	return file_mojo_core_url_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Url_Query) GetVals() map[string]*StringValues {
	if x != nil {
		return x.Vals
	}
	return nil
}

var File_mojo_core_url_proto protoreflect.FileDescriptor

var file_mojo_core_url_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x72, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x1a, 0x15, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x6f, 0x78, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x03, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x6a,
	0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x72, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x2a, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55,
	0x72, 0x6c, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x50, 0x0a, 0x09, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x8d, 0x01,
	0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x55, 0x72, 0x6c, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x56, 0x61, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x1a, 0x50, 0x0a, 0x09, 0x56,
	0x61, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x55, 0x0a,
	0x16, 0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6d, 0x6f,
	0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x08, 0x55, 0x72, 0x6c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b,
	0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_url_proto_rawDescOnce sync.Once
	file_mojo_core_url_proto_rawDescData = file_mojo_core_url_proto_rawDesc
)

func file_mojo_core_url_proto_rawDescGZIP() []byte {
	file_mojo_core_url_proto_rawDescOnce.Do(func() {
		file_mojo_core_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_url_proto_rawDescData)
	})
	return file_mojo_core_url_proto_rawDescData
}

var file_mojo_core_url_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mojo_core_url_proto_goTypes = []interface{}{
	(*Url)(nil),           // 0: mojo.core.Url
	(*Url_Authority)(nil), // 1: mojo.core.Url.Authority
	(*Url_Query)(nil),     // 2: mojo.core.Url.Query
	nil,                   // 3: mojo.core.Url.Query.ValsEntry
	(*StringValues)(nil),  // 4: mojo.core.StringValues
}
var file_mojo_core_url_proto_depIdxs = []int32{
	1, // 0: mojo.core.Url.authority:type_name -> mojo.core.Url.Authority
	2, // 1: mojo.core.Url.query:type_name -> mojo.core.Url.Query
	3, // 2: mojo.core.Url.Query.vals:type_name -> mojo.core.Url.Query.ValsEntry
	4, // 3: mojo.core.Url.Query.ValsEntry.value:type_name -> mojo.core.StringValues
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mojo_core_url_proto_init() }
func file_mojo_core_url_proto_init() {
	if File_mojo_core_url_proto != nil {
		return
	}
	file_mojo_core_boxed_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mojo_core_url_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Url); i {
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
		file_mojo_core_url_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Url_Authority); i {
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
		file_mojo_core_url_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Url_Query); i {
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
			RawDescriptor: file_mojo_core_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_url_proto_goTypes,
		DependencyIndexes: file_mojo_core_url_proto_depIdxs,
		MessageInfos:      file_mojo_core_url_proto_msgTypes,
	}.Build()
	File_mojo_core_url_proto = out.File
	file_mojo_core_url_proto_rawDesc = nil
	file_mojo_core_url_proto_goTypes = nil
	file_mojo_core_url_proto_depIdxs = nil
}
