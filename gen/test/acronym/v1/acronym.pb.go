// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: test/acronym/v1/acronym.proto

package acronymv1

import (
	_ "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ManageAWSRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Urn           string                 `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ManageAWSRequest) Reset() {
	*x = ManageAWSRequest{}
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ManageAWSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageAWSRequest) ProtoMessage() {}

func (x *ManageAWSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageAWSRequest.ProtoReflect.Descriptor instead.
func (*ManageAWSRequest) Descriptor() ([]byte, []int) {
	return file_test_acronym_v1_acronym_proto_rawDescGZIP(), []int{0}
}

func (x *ManageAWSRequest) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

type ManageAWSResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Urn           string                 `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ManageAWSResponse) Reset() {
	*x = ManageAWSResponse{}
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ManageAWSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageAWSResponse) ProtoMessage() {}

func (x *ManageAWSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageAWSResponse.ProtoReflect.Descriptor instead.
func (*ManageAWSResponse) Descriptor() ([]byte, []int) {
	return file_test_acronym_v1_acronym_proto_rawDescGZIP(), []int{1}
}

func (x *ManageAWSResponse) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

type ManageAWSResourceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Urn           string                 `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	K8SNamespace  string                 `protobuf:"bytes,2,opt,name=k8s_namespace,json=k8sNamespace,proto3" json:"k8s_namespace,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ManageAWSResourceRequest) Reset() {
	*x = ManageAWSResourceRequest{}
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ManageAWSResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageAWSResourceRequest) ProtoMessage() {}

func (x *ManageAWSResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageAWSResourceRequest.ProtoReflect.Descriptor instead.
func (*ManageAWSResourceRequest) Descriptor() ([]byte, []int) {
	return file_test_acronym_v1_acronym_proto_rawDescGZIP(), []int{2}
}

func (x *ManageAWSResourceRequest) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *ManageAWSResourceRequest) GetK8SNamespace() string {
	if x != nil {
		return x.K8SNamespace
	}
	return ""
}

type ManageAWSResourceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Urn           string                 `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ManageAWSResourceResponse) Reset() {
	*x = ManageAWSResourceResponse{}
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ManageAWSResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageAWSResourceResponse) ProtoMessage() {}

func (x *ManageAWSResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageAWSResourceResponse.ProtoReflect.Descriptor instead.
func (*ManageAWSResourceResponse) Descriptor() ([]byte, []int) {
	return file_test_acronym_v1_acronym_proto_rawDescGZIP(), []int{3}
}

func (x *ManageAWSResourceResponse) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

type ManageAWSResourceURNRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Urn           string                 `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ManageAWSResourceURNRequest) Reset() {
	*x = ManageAWSResourceURNRequest{}
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ManageAWSResourceURNRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageAWSResourceURNRequest) ProtoMessage() {}

func (x *ManageAWSResourceURNRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageAWSResourceURNRequest.ProtoReflect.Descriptor instead.
func (*ManageAWSResourceURNRequest) Descriptor() ([]byte, []int) {
	return file_test_acronym_v1_acronym_proto_rawDescGZIP(), []int{4}
}

func (x *ManageAWSResourceURNRequest) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

type ManageAWSResourceURNResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Urn           string                 `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ManageAWSResourceURNResponse) Reset() {
	*x = ManageAWSResourceURNResponse{}
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ManageAWSResourceURNResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageAWSResourceURNResponse) ProtoMessage() {}

func (x *ManageAWSResourceURNResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageAWSResourceURNResponse.ProtoReflect.Descriptor instead.
func (*ManageAWSResourceURNResponse) Descriptor() ([]byte, []int) {
	return file_test_acronym_v1_acronym_proto_rawDescGZIP(), []int{5}
}

func (x *ManageAWSResourceURNResponse) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

type SomethingV1FooBarRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Urn           string                 `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	K8SNamespace  string                 `protobuf:"bytes,2,opt,name=k8s_namespace,json=k8sNamespace,proto3" json:"k8s_namespace,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SomethingV1FooBarRequest) Reset() {
	*x = SomethingV1FooBarRequest{}
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SomethingV1FooBarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomethingV1FooBarRequest) ProtoMessage() {}

func (x *SomethingV1FooBarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomethingV1FooBarRequest.ProtoReflect.Descriptor instead.
func (*SomethingV1FooBarRequest) Descriptor() ([]byte, []int) {
	return file_test_acronym_v1_acronym_proto_rawDescGZIP(), []int{6}
}

func (x *SomethingV1FooBarRequest) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *SomethingV1FooBarRequest) GetK8SNamespace() string {
	if x != nil {
		return x.K8SNamespace
	}
	return ""
}

type SomethingV1FooBarResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Urn           string                 `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SomethingV1FooBarResponse) Reset() {
	*x = SomethingV1FooBarResponse{}
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SomethingV1FooBarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomethingV1FooBarResponse) ProtoMessage() {}

func (x *SomethingV1FooBarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomethingV1FooBarResponse.ProtoReflect.Descriptor instead.
func (*SomethingV1FooBarResponse) Descriptor() ([]byte, []int) {
	return file_test_acronym_v1_acronym_proto_rawDescGZIP(), []int{7}
}

func (x *SomethingV1FooBarResponse) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

type SomethingV2FooBarRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Urn           string                 `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	K8SNamespace  string                 `protobuf:"bytes,2,opt,name=k8s_namespace,json=k8sNamespace,proto3" json:"k8s_namespace,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SomethingV2FooBarRequest) Reset() {
	*x = SomethingV2FooBarRequest{}
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SomethingV2FooBarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomethingV2FooBarRequest) ProtoMessage() {}

func (x *SomethingV2FooBarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomethingV2FooBarRequest.ProtoReflect.Descriptor instead.
func (*SomethingV2FooBarRequest) Descriptor() ([]byte, []int) {
	return file_test_acronym_v1_acronym_proto_rawDescGZIP(), []int{8}
}

func (x *SomethingV2FooBarRequest) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *SomethingV2FooBarRequest) GetK8SNamespace() string {
	if x != nil {
		return x.K8SNamespace
	}
	return ""
}

type SomethingV2FooBarResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Urn           string                 `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SomethingV2FooBarResponse) Reset() {
	*x = SomethingV2FooBarResponse{}
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SomethingV2FooBarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomethingV2FooBarResponse) ProtoMessage() {}

func (x *SomethingV2FooBarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_acronym_v1_acronym_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomethingV2FooBarResponse.ProtoReflect.Descriptor instead.
func (*SomethingV2FooBarResponse) Descriptor() ([]byte, []int) {
	return file_test_acronym_v1_acronym_proto_rawDescGZIP(), []int{9}
}

func (x *SomethingV2FooBarResponse) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

var File_test_acronym_v1_acronym_proto protoreflect.FileDescriptor

var file_test_acronym_v1_acronym_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x2e, 0x76, 0x31,
	0x1a, 0x1a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x10,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6e, 0x22, 0x25, 0x0a, 0x11, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x41, 0x57, 0x53, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x22, 0x51, 0x0a, 0x18, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6b, 0x38, 0x73, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6b, 0x38, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x2d, 0x0a, 0x19,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x22, 0x2f, 0x0a, 0x1b, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x55, 0x52, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x22, 0x30, 0x0a, 0x1c,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x55, 0x52, 0x4e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x22, 0x51,
	0x0a, 0x18, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x56, 0x31, 0x46, 0x6f, 0x6f,
	0x42, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x6b, 0x38, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6b, 0x38, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x22, 0x2d, 0x0a, 0x19, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x56, 0x31,
	0x46, 0x6f, 0x6f, 0x42, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e,
	0x22, 0x86, 0x01, 0x0a, 0x18, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x56, 0x32,
	0x46, 0x6f, 0x6f, 0x42, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x12,
	0x58, 0x0a, 0x0d, 0x6b, 0x38, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0x8a, 0xc4, 0x03, 0x2f, 0x0a, 0x2d, 0x12, 0x0d,
	0x6b, 0x38, 0x73, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x19, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x01, 0x6e, 0x52, 0x0c, 0x6b, 0x38, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x2d, 0x0a, 0x19, 0x53, 0x6f, 0x6d,
	0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x42, 0x61, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x32, 0xcc, 0x06, 0x0a, 0x03, 0x41, 0x57, 0x53,
	0x12, 0x7e, 0x0a, 0x09, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x41, 0x57, 0x53, 0x12, 0x21, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x8a, 0xc4, 0x03, 0x26, 0x2a, 0x24, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x2d, 0x61, 0x77, 0x73, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x75, 0x72, 0x6e, 0x20, 0x7d,
	0x2f, 0x24, 0x7b, 0x21, 0x20, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x76, 0x34, 0x28, 0x29, 0x20, 0x7d,
	0x12, 0xac, 0x01, 0x0a, 0x11, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x29, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x63,
	0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x41,
	0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x8a,
	0xc4, 0x03, 0x34, 0x2a, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x77, 0x73, 0x2d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x75, 0x72, 0x6e,
	0x20, 0x7d, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x76, 0x34, 0x28, 0x29,
	0x20, 0x7d, 0x9a, 0x01, 0x02, 0x08, 0x01, 0x92, 0xc4, 0x03, 0x04, 0x22, 0x02, 0x08, 0x3c, 0x12,
	0x7d, 0x0a, 0x14, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x55, 0x52, 0x4e, 0x12, 0x2c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61,
	0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x52, 0x4e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x63, 0x72,
	0x6f, 0x6e, 0x79, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x41, 0x57,
	0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x52, 0x4e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x08, 0x92, 0xc4, 0x03, 0x04, 0x22, 0x02, 0x08, 0x3c, 0x12, 0xa0,
	0x01, 0x0a, 0x11, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x56, 0x31, 0x46, 0x6f,
	0x6f, 0x42, 0x61, 0x72, 0x12, 0x29, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x63, 0x72, 0x6f,
	0x6e, 0x79, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x56, 0x31, 0x46, 0x6f, 0x6f, 0x42, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x56, 0x31, 0x46, 0x6f, 0x6f,
	0x42, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x8a, 0xc4, 0x03,
	0x30, 0x2a, 0x2e, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2d, 0x76, 0x31, 0x2d,
	0x66, 0x6f, 0x6f, 0x2d, 0x62, 0x61, 0x72, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x75, 0x72, 0x6e, 0x20,
	0x7d, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x76, 0x34, 0x28, 0x29, 0x20,
	0x7d, 0x12, 0xdd, 0x01, 0x0a, 0x11, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x56,
	0x32, 0x46, 0x6f, 0x6f, 0x42, 0x61, 0x72, 0x12, 0x29, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61,
	0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x42, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x56, 0x32,
	0x46, 0x6f, 0x6f, 0x42, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x71,
	0x8a, 0xc4, 0x03, 0x6d, 0x2a, 0x2e, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2d,
	0x76, 0x32, 0x2d, 0x66, 0x6f, 0x6f, 0x2d, 0x62, 0x61, 0x72, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x75,
	0x72, 0x6e, 0x20, 0x7d, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x76, 0x34,
	0x28, 0x29, 0x20, 0x7d, 0x9a, 0x01, 0x3a, 0x12, 0x11, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x2d, 0x66, 0x6f, 0x6f, 0x2d, 0x62, 0x61, 0x72, 0x1a, 0x19, 0x64, 0x6f, 0x20, 0x73,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x66, 0x6f,
	0x6f, 0x20, 0x62, 0x61, 0x72, 0x22, 0x03, 0x73, 0x66, 0x62, 0x22, 0x05, 0x73, 0x66, 0x62, 0x76,
	0x32, 0x1a, 0x14, 0x8a, 0xc4, 0x03, 0x10, 0x0a, 0x0e, 0x61, 0x77, 0x73, 0x2d, 0x74, 0x61, 0x73,
	0x6b, 0x2d, 0x71, 0x75, 0x65, 0x75, 0x65, 0x42, 0xca, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x2e, 0x76, 0x31, 0x42,
	0x0c, 0x41, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x75, 0x64,
	0x64, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2f, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x61,
	0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x41, 0x58, 0xaa, 0x02,
	0x0f, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0f, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x41, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x1b, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x41, 0x63, 0x72, 0x6f, 0x6e, 0x79,
	0x6d, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x11, 0x54, 0x65, 0x73, 0x74, 0x3a, 0x3a, 0x41, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_test_acronym_v1_acronym_proto_rawDescOnce sync.Once
	file_test_acronym_v1_acronym_proto_rawDescData []byte
)

func file_test_acronym_v1_acronym_proto_rawDescGZIP() []byte {
	file_test_acronym_v1_acronym_proto_rawDescOnce.Do(func() {
		file_test_acronym_v1_acronym_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_test_acronym_v1_acronym_proto_rawDesc), len(file_test_acronym_v1_acronym_proto_rawDesc)))
	})
	return file_test_acronym_v1_acronym_proto_rawDescData
}

var file_test_acronym_v1_acronym_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_test_acronym_v1_acronym_proto_goTypes = []any{
	(*ManageAWSRequest)(nil),             // 0: test.acronym.v1.ManageAWSRequest
	(*ManageAWSResponse)(nil),            // 1: test.acronym.v1.ManageAWSResponse
	(*ManageAWSResourceRequest)(nil),     // 2: test.acronym.v1.ManageAWSResourceRequest
	(*ManageAWSResourceResponse)(nil),    // 3: test.acronym.v1.ManageAWSResourceResponse
	(*ManageAWSResourceURNRequest)(nil),  // 4: test.acronym.v1.ManageAWSResourceURNRequest
	(*ManageAWSResourceURNResponse)(nil), // 5: test.acronym.v1.ManageAWSResourceURNResponse
	(*SomethingV1FooBarRequest)(nil),     // 6: test.acronym.v1.SomethingV1FooBarRequest
	(*SomethingV1FooBarResponse)(nil),    // 7: test.acronym.v1.SomethingV1FooBarResponse
	(*SomethingV2FooBarRequest)(nil),     // 8: test.acronym.v1.SomethingV2FooBarRequest
	(*SomethingV2FooBarResponse)(nil),    // 9: test.acronym.v1.SomethingV2FooBarResponse
}
var file_test_acronym_v1_acronym_proto_depIdxs = []int32{
	0, // 0: test.acronym.v1.AWS.ManageAWS:input_type -> test.acronym.v1.ManageAWSRequest
	2, // 1: test.acronym.v1.AWS.ManageAWSResource:input_type -> test.acronym.v1.ManageAWSResourceRequest
	4, // 2: test.acronym.v1.AWS.ManageAWSResourceURN:input_type -> test.acronym.v1.ManageAWSResourceURNRequest
	6, // 3: test.acronym.v1.AWS.SomethingV1FooBar:input_type -> test.acronym.v1.SomethingV1FooBarRequest
	8, // 4: test.acronym.v1.AWS.SomethingV2FooBar:input_type -> test.acronym.v1.SomethingV2FooBarRequest
	1, // 5: test.acronym.v1.AWS.ManageAWS:output_type -> test.acronym.v1.ManageAWSResponse
	3, // 6: test.acronym.v1.AWS.ManageAWSResource:output_type -> test.acronym.v1.ManageAWSResourceResponse
	5, // 7: test.acronym.v1.AWS.ManageAWSResourceURN:output_type -> test.acronym.v1.ManageAWSResourceURNResponse
	7, // 8: test.acronym.v1.AWS.SomethingV1FooBar:output_type -> test.acronym.v1.SomethingV1FooBarResponse
	9, // 9: test.acronym.v1.AWS.SomethingV2FooBar:output_type -> test.acronym.v1.SomethingV2FooBarResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_acronym_v1_acronym_proto_init() }
func file_test_acronym_v1_acronym_proto_init() {
	if File_test_acronym_v1_acronym_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_test_acronym_v1_acronym_proto_rawDesc), len(file_test_acronym_v1_acronym_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_acronym_v1_acronym_proto_goTypes,
		DependencyIndexes: file_test_acronym_v1_acronym_proto_depIdxs,
		MessageInfos:      file_test_acronym_v1_acronym_proto_msgTypes,
	}.Build()
	File_test_acronym_v1_acronym_proto = out.File
	file_test_acronym_v1_acronym_proto_goTypes = nil
	file_test_acronym_v1_acronym_proto_depIdxs = nil
}
