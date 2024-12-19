// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: example/xns/v1/xns.proto

package xnsv1

import (
	_ "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Foo_Status int32

const (
	Foo_FOO_STATUS_UNSPECIFIED Foo_Status = 0
	Foo_FOO_STATUS_READY       Foo_Status = 1
	Foo_FOO_STATUS_CREATING    Foo_Status = 2
)

// Enum value maps for Foo_Status.
var (
	Foo_Status_name = map[int32]string{
		0: "FOO_STATUS_UNSPECIFIED",
		1: "FOO_STATUS_READY",
		2: "FOO_STATUS_CREATING",
	}
	Foo_Status_value = map[string]int32{
		"FOO_STATUS_UNSPECIFIED": 0,
		"FOO_STATUS_READY":       1,
		"FOO_STATUS_CREATING":    2,
	}
)

func (x Foo_Status) Enum() *Foo_Status {
	p := new(Foo_Status)
	*p = x
	return p
}

func (x Foo_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Foo_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_example_xns_v1_xns_proto_enumTypes[0].Descriptor()
}

func (Foo_Status) Type() protoreflect.EnumType {
	return &file_example_xns_v1_xns_proto_enumTypes[0]
}

func (x Foo_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Foo_Status.Descriptor instead.
func (Foo_Status) EnumDescriptor() ([]byte, []int) {
	return file_example_xns_v1_xns_proto_rawDescGZIP(), []int{2, 0}
}

// CreateFooRequest describes the input to a CreateFoo workflow
type CreateFooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique foo name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateFooRequest) Reset() {
	*x = CreateFooRequest{}
	mi := &file_example_xns_v1_xns_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFooRequest) ProtoMessage() {}

func (x *CreateFooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_xns_v1_xns_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFooRequest.ProtoReflect.Descriptor instead.
func (*CreateFooRequest) Descriptor() ([]byte, []int) {
	return file_example_xns_v1_xns_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFooRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// SampleWorkflowWithMutexResponse describes the output from a CreateFoo workflow
type CreateFooResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo *Foo `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
}

func (x *CreateFooResponse) Reset() {
	*x = CreateFooResponse{}
	mi := &file_example_xns_v1_xns_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFooResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFooResponse) ProtoMessage() {}

func (x *CreateFooResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_xns_v1_xns_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFooResponse.ProtoReflect.Descriptor instead.
func (*CreateFooResponse) Descriptor() ([]byte, []int) {
	return file_example_xns_v1_xns_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFooResponse) GetFoo() *Foo {
	if x != nil {
		return x.Foo
	}
	return nil
}

// Foo describes an illustrative foo resource
type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status Foo_Status `protobuf:"varint,2,opt,name=status,proto3,enum=example.xns.v1.Foo_Status" json:"status,omitempty"`
}

func (x *Foo) Reset() {
	*x = Foo{}
	mi := &file_example_xns_v1_xns_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_example_xns_v1_xns_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_example_xns_v1_xns_proto_rawDescGZIP(), []int{2}
}

func (x *Foo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Foo) GetStatus() Foo_Status {
	if x != nil {
		return x.Status
	}
	return Foo_FOO_STATUS_UNSPECIFIED
}

// GetFooProgressResponse describes the output from a GetFooProgress query
type GetFooProgressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Progress float32    `protobuf:"fixed32,1,opt,name=progress,proto3" json:"progress,omitempty"`
	Status   Foo_Status `protobuf:"varint,2,opt,name=status,proto3,enum=example.xns.v1.Foo_Status" json:"status,omitempty"`
}

func (x *GetFooProgressResponse) Reset() {
	*x = GetFooProgressResponse{}
	mi := &file_example_xns_v1_xns_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFooProgressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFooProgressResponse) ProtoMessage() {}

func (x *GetFooProgressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_xns_v1_xns_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFooProgressResponse.ProtoReflect.Descriptor instead.
func (*GetFooProgressResponse) Descriptor() ([]byte, []int) {
	return file_example_xns_v1_xns_proto_rawDescGZIP(), []int{3}
}

func (x *GetFooProgressResponse) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *GetFooProgressResponse) GetStatus() Foo_Status {
	if x != nil {
		return x.Status
	}
	return Foo_FOO_STATUS_UNSPECIFIED
}

// NotifyRequest describes the input to a Notify activity
type NotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NotifyRequest) Reset() {
	*x = NotifyRequest{}
	mi := &file_example_xns_v1_xns_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyRequest) ProtoMessage() {}

func (x *NotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_xns_v1_xns_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyRequest.ProtoReflect.Descriptor instead.
func (*NotifyRequest) Descriptor() ([]byte, []int) {
	return file_example_xns_v1_xns_proto_rawDescGZIP(), []int{4}
}

func (x *NotifyRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// ProvisionFooRequest describes the input to a ProvisionFoo workflow
type ProvisionFooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique foo name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProvisionFooRequest) Reset() {
	*x = ProvisionFooRequest{}
	mi := &file_example_xns_v1_xns_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProvisionFooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisionFooRequest) ProtoMessage() {}

func (x *ProvisionFooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_xns_v1_xns_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisionFooRequest.ProtoReflect.Descriptor instead.
func (*ProvisionFooRequest) Descriptor() ([]byte, []int) {
	return file_example_xns_v1_xns_proto_rawDescGZIP(), []int{5}
}

func (x *ProvisionFooRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// SampleWorkflowWithMutexResponse describes the output from a ProvisionFoo workflow
type ProvisionFooResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo *Foo `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
}

func (x *ProvisionFooResponse) Reset() {
	*x = ProvisionFooResponse{}
	mi := &file_example_xns_v1_xns_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProvisionFooResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisionFooResponse) ProtoMessage() {}

func (x *ProvisionFooResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_xns_v1_xns_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisionFooResponse.ProtoReflect.Descriptor instead.
func (*ProvisionFooResponse) Descriptor() ([]byte, []int) {
	return file_example_xns_v1_xns_proto_rawDescGZIP(), []int{6}
}

func (x *ProvisionFooResponse) GetFoo() *Foo {
	if x != nil {
		return x.Foo
	}
	return nil
}

// SetFooProgressRequest describes the input to a SetFooProgress signal
type SetFooProgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// value of current workflow progress
	Progress float32 `protobuf:"fixed32,1,opt,name=progress,proto3" json:"progress,omitempty"`
}

func (x *SetFooProgressRequest) Reset() {
	*x = SetFooProgressRequest{}
	mi := &file_example_xns_v1_xns_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetFooProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFooProgressRequest) ProtoMessage() {}

func (x *SetFooProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_xns_v1_xns_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFooProgressRequest.ProtoReflect.Descriptor instead.
func (*SetFooProgressRequest) Descriptor() ([]byte, []int) {
	return file_example_xns_v1_xns_proto_rawDescGZIP(), []int{7}
}

func (x *SetFooProgressRequest) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

var File_example_xns_v1_xns_proto protoreflect.FileDescriptor

var file_example_xns_v1_xns_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x78, 0x6e, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x78, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x78, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x78, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x6f, 0x6f, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x78, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x53, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1a, 0x0a, 0x16, 0x46, 0x4f, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x46, 0x4f, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x4f, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x22, 0x68, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x78, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x29, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x29, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x6f,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x14,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x78, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x22, 0x33, 0x0a, 0x15, 0x53,
	0x65, 0x74, 0x46, 0x6f, 0x6f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x32, 0x96, 0x01, 0x0a, 0x03, 0x58, 0x6e, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6f, 0x12, 0x23, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x78, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x78, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x8a, 0xc4, 0x03, 0x21, 0x2a, 0x1f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x66, 0x6f, 0x6f, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x6e, 0x61,
	0x6d, 0x65, 0x2e, 0x73, 0x6c, 0x75, 0x67, 0x28, 0x29, 0x20, 0x7d, 0x1a, 0x0c, 0x8a, 0xc4, 0x03,
	0x08, 0x0a, 0x06, 0x78, 0x6e, 0x73, 0x2d, 0x76, 0x31, 0x32, 0xa0, 0x05, 0x0a, 0x07, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0xc6, 0x01, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x6f, 0x6f, 0x12, 0x20, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x78, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x78, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x74, 0x8a, 0xc4, 0x03, 0x70, 0x0a, 0x10,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x10, 0x01, 0x1a, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f,
	0x6f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x03, 0x08, 0x90, 0x1c, 0x2a, 0x1c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2d, 0x66, 0x6f, 0x6f, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x6e,
	0x61, 0x6d, 0x65, 0x2e, 0x73, 0x6c, 0x75, 0x67, 0x28, 0x29, 0x20, 0x7d, 0x30, 0x01, 0x82, 0x01,
	0x0d, 0x22, 0x03, 0x08, 0xae, 0x1c, 0x2a, 0x02, 0x08, 0x14, 0x42, 0x02, 0x08, 0x0a, 0x12, 0x64,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x26, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x78, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x12, 0x9a, 0xc4, 0x03, 0x0e, 0x12, 0x0c, 0x22, 0x02, 0x08, 0x3c, 0x2a, 0x02, 0x08, 0x14,
	0x42, 0x02, 0x08, 0x0a, 0x12, 0x4d, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1d,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x78, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0c, 0x92, 0xc4, 0x03, 0x08, 0x22, 0x02, 0x08, 0x1e, 0x32,
	0x02, 0x20, 0x03, 0x12, 0x63, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x78, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x12, 0xa2, 0xc4, 0x03, 0x0e, 0x12, 0x0c, 0x22, 0x02, 0x08, 0x3c,
	0x2a, 0x02, 0x08, 0x14, 0x42, 0x02, 0x08, 0x0a, 0x12, 0x9f, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x78, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x78, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0xaa,
	0xc4, 0x03, 0x37, 0x0a, 0x27, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x28, 0x29, 0x20, 0x7d, 0x2a, 0x0c, 0x22, 0x02,
	0x08, 0x3c, 0x2a, 0x02, 0x08, 0x14, 0x42, 0x02, 0x08, 0x0a, 0x1a, 0x10, 0x8a, 0xc4, 0x03, 0x0c,
	0x0a, 0x0a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x76, 0x31, 0x42, 0xbc, 0x01, 0x0a,
	0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x78, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x08, 0x58, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x75, 0x64,
	0x64, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x78, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x78, 0x6e,
	0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x0e, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x58, 0x6e, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x58, 0x6e, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x58, 0x6e, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x3a, 0x3a, 0x58, 0x6e, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_example_xns_v1_xns_proto_rawDescOnce sync.Once
	file_example_xns_v1_xns_proto_rawDescData = file_example_xns_v1_xns_proto_rawDesc
)

func file_example_xns_v1_xns_proto_rawDescGZIP() []byte {
	file_example_xns_v1_xns_proto_rawDescOnce.Do(func() {
		file_example_xns_v1_xns_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_xns_v1_xns_proto_rawDescData)
	})
	return file_example_xns_v1_xns_proto_rawDescData
}

var file_example_xns_v1_xns_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_example_xns_v1_xns_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_example_xns_v1_xns_proto_goTypes = []any{
	(Foo_Status)(0),                // 0: example.xns.v1.Foo.Status
	(*CreateFooRequest)(nil),       // 1: example.xns.v1.CreateFooRequest
	(*CreateFooResponse)(nil),      // 2: example.xns.v1.CreateFooResponse
	(*Foo)(nil),                    // 3: example.xns.v1.Foo
	(*GetFooProgressResponse)(nil), // 4: example.xns.v1.GetFooProgressResponse
	(*NotifyRequest)(nil),          // 5: example.xns.v1.NotifyRequest
	(*ProvisionFooRequest)(nil),    // 6: example.xns.v1.ProvisionFooRequest
	(*ProvisionFooResponse)(nil),   // 7: example.xns.v1.ProvisionFooResponse
	(*SetFooProgressRequest)(nil),  // 8: example.xns.v1.SetFooProgressRequest
	(*emptypb.Empty)(nil),          // 9: google.protobuf.Empty
}
var file_example_xns_v1_xns_proto_depIdxs = []int32{
	3,  // 0: example.xns.v1.CreateFooResponse.foo:type_name -> example.xns.v1.Foo
	0,  // 1: example.xns.v1.Foo.status:type_name -> example.xns.v1.Foo.Status
	0,  // 2: example.xns.v1.GetFooProgressResponse.status:type_name -> example.xns.v1.Foo.Status
	3,  // 3: example.xns.v1.ProvisionFooResponse.foo:type_name -> example.xns.v1.Foo
	6,  // 4: example.xns.v1.Xns.ProvisionFoo:input_type -> example.xns.v1.ProvisionFooRequest
	1,  // 5: example.xns.v1.Example.CreateFoo:input_type -> example.xns.v1.CreateFooRequest
	9,  // 6: example.xns.v1.Example.GetFooProgress:input_type -> google.protobuf.Empty
	5,  // 7: example.xns.v1.Example.Notify:input_type -> example.xns.v1.NotifyRequest
	8,  // 8: example.xns.v1.Example.SetFooProgress:input_type -> example.xns.v1.SetFooProgressRequest
	8,  // 9: example.xns.v1.Example.UpdateFooProgress:input_type -> example.xns.v1.SetFooProgressRequest
	7,  // 10: example.xns.v1.Xns.ProvisionFoo:output_type -> example.xns.v1.ProvisionFooResponse
	2,  // 11: example.xns.v1.Example.CreateFoo:output_type -> example.xns.v1.CreateFooResponse
	4,  // 12: example.xns.v1.Example.GetFooProgress:output_type -> example.xns.v1.GetFooProgressResponse
	9,  // 13: example.xns.v1.Example.Notify:output_type -> google.protobuf.Empty
	9,  // 14: example.xns.v1.Example.SetFooProgress:output_type -> google.protobuf.Empty
	4,  // 15: example.xns.v1.Example.UpdateFooProgress:output_type -> example.xns.v1.GetFooProgressResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_example_xns_v1_xns_proto_init() }
func file_example_xns_v1_xns_proto_init() {
	if File_example_xns_v1_xns_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_xns_v1_xns_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_example_xns_v1_xns_proto_goTypes,
		DependencyIndexes: file_example_xns_v1_xns_proto_depIdxs,
		EnumInfos:         file_example_xns_v1_xns_proto_enumTypes,
		MessageInfos:      file_example_xns_v1_xns_proto_msgTypes,
	}.Build()
	File_example_xns_v1_xns_proto = out.File
	file_example_xns_v1_xns_proto_rawDesc = nil
	file_example_xns_v1_xns_proto_goTypes = nil
	file_example_xns_v1_xns_proto_depIdxs = nil
}
