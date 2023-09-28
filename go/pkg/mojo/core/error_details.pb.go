// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: mojo/core/error_details.proto

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

type RetryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetryDelay *Duration `protobuf:"bytes,1,opt,name=retry_delay,json=retryDelay,proto3" json:"retryDelay,omitempty"`
}

func (x *RetryInfo) Reset() {
	*x = RetryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_error_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryInfo) ProtoMessage() {}

func (x *RetryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_error_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryInfo.ProtoReflect.Descriptor instead.
func (*RetryInfo) Descriptor() ([]byte, []int) {
	return file_mojo_core_error_details_proto_rawDescGZIP(), []int{0}
}

func (x *RetryInfo) GetRetryDelay() *Duration {
	if x != nil {
		return x.RetryDelay
	}
	return nil
}

type DebugInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StackEntries []string `protobuf:"bytes,1,rep,name=stack_entries,json=stackEntries,proto3" json:"stackEntries,omitempty"`
	Detail       string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *DebugInfo) Reset() {
	*x = DebugInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_error_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugInfo) ProtoMessage() {}

func (x *DebugInfo) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_error_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugInfo.ProtoReflect.Descriptor instead.
func (*DebugInfo) Descriptor() ([]byte, []int) {
	return file_mojo_core_error_details_proto_rawDescGZIP(), []int{1}
}

func (x *DebugInfo) GetStackEntries() []string {
	if x != nil {
		return x.StackEntries
	}
	return nil
}

func (x *DebugInfo) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type QuotaFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Violations []*QuotaFailure_Violation `protobuf:"bytes,1,rep,name=violations,proto3" json:"violations,omitempty"`
}

func (x *QuotaFailure) Reset() {
	*x = QuotaFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_error_details_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuotaFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotaFailure) ProtoMessage() {}

func (x *QuotaFailure) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_error_details_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotaFailure.ProtoReflect.Descriptor instead.
func (*QuotaFailure) Descriptor() ([]byte, []int) {
	return file_mojo_core_error_details_proto_rawDescGZIP(), []int{2}
}

func (x *QuotaFailure) GetViolations() []*QuotaFailure_Violation {
	if x != nil {
		return x.Violations
	}
	return nil
}

type ErrorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason   string            `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
	Domain   string            `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Metadata map[string]*Value `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ErrorInfo) Reset() {
	*x = ErrorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_error_details_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorInfo) ProtoMessage() {}

func (x *ErrorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_error_details_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorInfo.ProtoReflect.Descriptor instead.
func (*ErrorInfo) Descriptor() ([]byte, []int) {
	return file_mojo_core_error_details_proto_rawDescGZIP(), []int{3}
}

func (x *ErrorInfo) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *ErrorInfo) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ErrorInfo) GetMetadata() map[string]*Value {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type PreconditionFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Violations []*PreconditionFailure_Violation `protobuf:"bytes,1,rep,name=violations,proto3" json:"violations,omitempty"`
}

func (x *PreconditionFailure) Reset() {
	*x = PreconditionFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_error_details_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreconditionFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreconditionFailure) ProtoMessage() {}

func (x *PreconditionFailure) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_error_details_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreconditionFailure.ProtoReflect.Descriptor instead.
func (*PreconditionFailure) Descriptor() ([]byte, []int) {
	return file_mojo_core_error_details_proto_rawDescGZIP(), []int{4}
}

func (x *PreconditionFailure) GetViolations() []*PreconditionFailure_Violation {
	if x != nil {
		return x.Violations
	}
	return nil
}

type MalformedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldViolations []*MalformedRequest_FieldViolation `protobuf:"bytes,1,rep,name=field_violations,json=fieldViolations,proto3" json:"fieldViolations,omitempty"`
}

func (x *MalformedRequest) Reset() {
	*x = MalformedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_error_details_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MalformedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MalformedRequest) ProtoMessage() {}

func (x *MalformedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_error_details_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MalformedRequest.ProtoReflect.Descriptor instead.
func (*MalformedRequest) Descriptor() ([]byte, []int) {
	return file_mojo_core_error_details_proto_rawDescGZIP(), []int{5}
}

func (x *MalformedRequest) GetFieldViolations() []*MalformedRequest_FieldViolation {
	if x != nil {
		return x.FieldViolations
	}
	return nil
}

type RequestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId   string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"requestId,omitempty"`
	ServingData string `protobuf:"bytes,2,opt,name=serving_data,json=servingData,proto3" json:"servingData,omitempty"`
}

func (x *RequestInfo) Reset() {
	*x = RequestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_error_details_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestInfo) ProtoMessage() {}

func (x *RequestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_error_details_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestInfo.ProtoReflect.Descriptor instead.
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return file_mojo_core_error_details_proto_rawDescGZIP(), []int{6}
}

func (x *RequestInfo) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *RequestInfo) GetServingData() string {
	if x != nil {
		return x.ServingData
	}
	return ""
}

type ResourceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceType string `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resourceType,omitempty"`
	ResourceName string `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resourceName,omitempty"`
	Owner        string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Description  string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ResourceInfo) Reset() {
	*x = ResourceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_error_details_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceInfo) ProtoMessage() {}

func (x *ResourceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_error_details_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceInfo.ProtoReflect.Descriptor instead.
func (*ResourceInfo) Descriptor() ([]byte, []int) {
	return file_mojo_core_error_details_proto_rawDescGZIP(), []int{7}
}

func (x *ResourceInfo) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *ResourceInfo) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ResourceInfo) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ResourceInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Help struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links []*Help_Link `protobuf:"bytes,1,rep,name=links,proto3" json:"links,omitempty"`
}

func (x *Help) Reset() {
	*x = Help{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_error_details_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Help) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Help) ProtoMessage() {}

func (x *Help) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_error_details_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Help.ProtoReflect.Descriptor instead.
func (*Help) Descriptor() ([]byte, []int) {
	return file_mojo_core_error_details_proto_rawDescGZIP(), []int{8}
}

func (x *Help) GetLinks() []*Help_Link {
	if x != nil {
		return x.Links
	}
	return nil
}

type LocalizedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locale  string `protobuf:"bytes,1,opt,name=locale,proto3" json:"locale,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LocalizedMessage) Reset() {
	*x = LocalizedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_error_details_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalizedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalizedMessage) ProtoMessage() {}

func (x *LocalizedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_error_details_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalizedMessage.ProtoReflect.Descriptor instead.
func (*LocalizedMessage) Descriptor() ([]byte, []int) {
	return file_mojo_core_error_details_proto_rawDescGZIP(), []int{9}
}

func (x *LocalizedMessage) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *LocalizedMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type QuotaFailure_Violation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject     string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *QuotaFailure_Violation) Reset() {
	*x = QuotaFailure_Violation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_error_details_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuotaFailure_Violation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotaFailure_Violation) ProtoMessage() {}

func (x *QuotaFailure_Violation) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_error_details_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotaFailure_Violation.ProtoReflect.Descriptor instead.
func (*QuotaFailure_Violation) Descriptor() ([]byte, []int) {
	return file_mojo_core_error_details_proto_rawDescGZIP(), []int{2, 0}
}

func (x *QuotaFailure_Violation) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *QuotaFailure_Violation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type PreconditionFailure_Violation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Subject     string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PreconditionFailure_Violation) Reset() {
	*x = PreconditionFailure_Violation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_error_details_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreconditionFailure_Violation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreconditionFailure_Violation) ProtoMessage() {}

func (x *PreconditionFailure_Violation) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_error_details_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreconditionFailure_Violation.ProtoReflect.Descriptor instead.
func (*PreconditionFailure_Violation) Descriptor() ([]byte, []int) {
	return file_mojo_core_error_details_proto_rawDescGZIP(), []int{4, 0}
}

func (x *PreconditionFailure_Violation) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PreconditionFailure_Violation) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *PreconditionFailure_Violation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type MalformedRequest_FieldViolation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field       string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *MalformedRequest_FieldViolation) Reset() {
	*x = MalformedRequest_FieldViolation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_error_details_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MalformedRequest_FieldViolation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MalformedRequest_FieldViolation) ProtoMessage() {}

func (x *MalformedRequest_FieldViolation) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_error_details_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MalformedRequest_FieldViolation.ProtoReflect.Descriptor instead.
func (*MalformedRequest_FieldViolation) Descriptor() ([]byte, []int) {
	return file_mojo_core_error_details_proto_rawDescGZIP(), []int{5, 0}
}

func (x *MalformedRequest_FieldViolation) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *MalformedRequest_FieldViolation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Help_Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Url         *Url   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Help_Link) Reset() {
	*x = Help_Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_error_details_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Help_Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Help_Link) ProtoMessage() {}

func (x *Help_Link) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_error_details_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Help_Link.ProtoReflect.Descriptor instead.
func (*Help_Link) Descriptor() ([]byte, []int) {
	return file_mojo_core_error_details_proto_rawDescGZIP(), []int{8, 0}
}

func (x *Help_Link) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Help_Link) GetUrl() *Url {
	if x != nil {
		return x.Url
	}
	return nil
}

var File_mojo_core_error_details_proto protoreflect.FileDescriptor

var file_mojo_core_error_details_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x14, 0x6d, 0x6f, 0x6a, 0x6f,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x72, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x09,
	0x52, 0x65, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x0b, 0x72, 0x65, 0x74,
	0x72, 0x79, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x22,
	0x48, 0x0a, 0x09, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x76, 0x69,
	0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x47, 0x0a,
	0x09, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xca, 0x01, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x4d, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xbc, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x76,
	0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e,
	0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x76, 0x69, 0x6f, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x5b, 0x0a, 0x09, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xb3, 0x01, 0x0a, 0x10, 0x4d, 0x61, 0x6c, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x10, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x61,
	0x6c, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x48,
	0x0a, 0x0e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x22, 0x90, 0x01, 0x0a, 0x0c, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7e, 0x0a, 0x04,
	0x48, 0x65, 0x6c, 0x70, 0x12, 0x2a, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x48, 0x65, 0x6c, 0x70, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x1a, 0x4a, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x44, 0x0a, 0x10,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x5e, 0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x11, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f,
	0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b, 0x63, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_error_details_proto_rawDescOnce sync.Once
	file_mojo_core_error_details_proto_rawDescData = file_mojo_core_error_details_proto_rawDesc
)

func file_mojo_core_error_details_proto_rawDescGZIP() []byte {
	file_mojo_core_error_details_proto_rawDescOnce.Do(func() {
		file_mojo_core_error_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_error_details_proto_rawDescData)
	})
	return file_mojo_core_error_details_proto_rawDescData
}

var file_mojo_core_error_details_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_mojo_core_error_details_proto_goTypes = []interface{}{
	(*RetryInfo)(nil),                       // 0: mojo.core.RetryInfo
	(*DebugInfo)(nil),                       // 1: mojo.core.DebugInfo
	(*QuotaFailure)(nil),                    // 2: mojo.core.QuotaFailure
	(*ErrorInfo)(nil),                       // 3: mojo.core.ErrorInfo
	(*PreconditionFailure)(nil),             // 4: mojo.core.PreconditionFailure
	(*MalformedRequest)(nil),                // 5: mojo.core.MalformedRequest
	(*RequestInfo)(nil),                     // 6: mojo.core.RequestInfo
	(*ResourceInfo)(nil),                    // 7: mojo.core.ResourceInfo
	(*Help)(nil),                            // 8: mojo.core.Help
	(*LocalizedMessage)(nil),                // 9: mojo.core.LocalizedMessage
	(*QuotaFailure_Violation)(nil),          // 10: mojo.core.QuotaFailure.Violation
	nil,                                     // 11: mojo.core.ErrorInfo.MetadataEntry
	(*PreconditionFailure_Violation)(nil),   // 12: mojo.core.PreconditionFailure.Violation
	(*MalformedRequest_FieldViolation)(nil), // 13: mojo.core.MalformedRequest.FieldViolation
	(*Help_Link)(nil),                       // 14: mojo.core.Help.Link
	(*Duration)(nil),                        // 15: mojo.core.Duration
	(*Value)(nil),                           // 16: mojo.core.Value
	(*Url)(nil),                             // 17: mojo.core.Url
}
var file_mojo_core_error_details_proto_depIdxs = []int32{
	15, // 0: mojo.core.RetryInfo.retry_delay:type_name -> mojo.core.Duration
	10, // 1: mojo.core.QuotaFailure.violations:type_name -> mojo.core.QuotaFailure.Violation
	11, // 2: mojo.core.ErrorInfo.metadata:type_name -> mojo.core.ErrorInfo.MetadataEntry
	12, // 3: mojo.core.PreconditionFailure.violations:type_name -> mojo.core.PreconditionFailure.Violation
	13, // 4: mojo.core.MalformedRequest.field_violations:type_name -> mojo.core.MalformedRequest.FieldViolation
	14, // 5: mojo.core.Help.links:type_name -> mojo.core.Help.Link
	16, // 6: mojo.core.ErrorInfo.MetadataEntry.value:type_name -> mojo.core.Value
	17, // 7: mojo.core.Help.Link.url:type_name -> mojo.core.Url
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_mojo_core_error_details_proto_init() }
func file_mojo_core_error_details_proto_init() {
	if File_mojo_core_error_details_proto != nil {
		return
	}
	file_mojo_core_time_proto_init()
	file_mojo_core_url_proto_init()
	file_mojo_core_value_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mojo_core_error_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryInfo); i {
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
		file_mojo_core_error_details_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugInfo); i {
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
		file_mojo_core_error_details_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuotaFailure); i {
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
		file_mojo_core_error_details_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorInfo); i {
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
		file_mojo_core_error_details_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreconditionFailure); i {
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
		file_mojo_core_error_details_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MalformedRequest); i {
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
		file_mojo_core_error_details_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestInfo); i {
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
		file_mojo_core_error_details_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceInfo); i {
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
		file_mojo_core_error_details_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Help); i {
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
		file_mojo_core_error_details_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalizedMessage); i {
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
		file_mojo_core_error_details_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuotaFailure_Violation); i {
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
		file_mojo_core_error_details_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreconditionFailure_Violation); i {
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
		file_mojo_core_error_details_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MalformedRequest_FieldViolation); i {
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
		file_mojo_core_error_details_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Help_Link); i {
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
			RawDescriptor: file_mojo_core_error_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_error_details_proto_goTypes,
		DependencyIndexes: file_mojo_core_error_details_proto_depIdxs,
		MessageInfos:      file_mojo_core_error_details_proto_msgTypes,
	}.Build()
	File_mojo_core_error_details_proto = out.File
	file_mojo_core_error_details_proto_rawDesc = nil
	file_mojo_core_error_details_proto_goTypes = nil
	file_mojo_core_error_details_proto_depIdxs = nil
}
