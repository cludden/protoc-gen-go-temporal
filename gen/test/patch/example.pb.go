// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: test/patch/example.proto

package patch

import (
	_ "github.com/alta/protopatch/patch/gopb"
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

type FooInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FooID         string                 `protobuf:"bytes,1,opt,name=foo_id,json=fooID,proto3" json:"fooID"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FooInput) Reset() {
	*x = FooInput{}
	mi := &file_test_patch_example_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FooInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooInput) ProtoMessage() {}

func (x *FooInput) ProtoReflect() protoreflect.Message {
	mi := &file_test_patch_example_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooInput.ProtoReflect.Descriptor instead.
func (*FooInput) Descriptor() ([]byte, []int) {
	return file_test_patch_example_proto_rawDescGZIP(), []int{0}
}

func (x *FooInput) GetFooID() string {
	if x != nil {
		return x.FooID
	}
	return ""
}

type FooOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FooOutput) Reset() {
	*x = FooOutput{}
	mi := &file_test_patch_example_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FooOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooOutput) ProtoMessage() {}

func (x *FooOutput) ProtoReflect() protoreflect.Message {
	mi := &file_test_patch_example_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooOutput.ProtoReflect.Descriptor instead.
func (*FooOutput) Descriptor() ([]byte, []int) {
	return file_test_patch_example_proto_rawDescGZIP(), []int{1}
}

var File_test_patch_example_proto protoreflect.FileDescriptor

var file_test_patch_example_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x0e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x36, 0x0a, 0x08, 0x46, 0x6f, 0x6f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x2a,
	0x0a, 0x06, 0x66, 0x6f, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13,
	0xca, 0xb5, 0x03, 0x0f, 0xa2, 0x01, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x66, 0x6f, 0x6f,
	0x49, 0x44, 0x22, 0x52, 0x05, 0x66, 0x6f, 0x6f, 0x49, 0x44, 0x22, 0x0b, 0x0a, 0x09, 0x46, 0x6f,
	0x6f, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0x5f, 0x0a, 0x0a, 0x46, 0x6f, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x14, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x46, 0x6f, 0x6f, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x46, 0x6f, 0x6f, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x0c, 0x8a, 0xc4, 0x03, 0x00, 0x92,
	0xc4, 0x03, 0x04, 0x22, 0x02, 0x08, 0x02, 0x1a, 0x0f, 0x8a, 0xc4, 0x03, 0x0b, 0x0a, 0x09, 0x66,
	0x6f, 0x6f, 0x2d, 0x71, 0x75, 0x65, 0x75, 0x65, 0x42, 0xa7, 0x01, 0xca, 0xb5, 0x03, 0x02, 0x08,
	0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x42, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x75, 0x64, 0x64, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0xa2, 0x02, 0x03, 0x54, 0x50,
	0x58, 0xaa, 0x02, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0xca, 0x02,
	0x0a, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x50, 0x61, 0x74, 0x63, 0x68, 0xe2, 0x02, 0x16, 0x54, 0x65,
	0x73, 0x74, 0x5c, 0x50, 0x61, 0x74, 0x63, 0x68, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x3a, 0x3a, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_test_patch_example_proto_rawDescOnce sync.Once
	file_test_patch_example_proto_rawDescData []byte
)

func file_test_patch_example_proto_rawDescGZIP() []byte {
	file_test_patch_example_proto_rawDescOnce.Do(func() {
		file_test_patch_example_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_test_patch_example_proto_rawDesc), len(file_test_patch_example_proto_rawDesc)))
	})
	return file_test_patch_example_proto_rawDescData
}

var file_test_patch_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_patch_example_proto_goTypes = []any{
	(*FooInput)(nil),  // 0: test.patch.FooInput
	(*FooOutput)(nil), // 1: test.patch.FooOutput
}
var file_test_patch_example_proto_depIdxs = []int32{
	0, // 0: test.patch.FooService.Foo:input_type -> test.patch.FooInput
	1, // 1: test.patch.FooService.Foo:output_type -> test.patch.FooOutput
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_patch_example_proto_init() }
func file_test_patch_example_proto_init() {
	if File_test_patch_example_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_test_patch_example_proto_rawDesc), len(file_test_patch_example_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_patch_example_proto_goTypes,
		DependencyIndexes: file_test_patch_example_proto_depIdxs,
		MessageInfos:      file_test_patch_example_proto_msgTypes,
	}.Build()
	File_test_patch_example_proto = out.File
	file_test_patch_example_proto_goTypes = nil
	file_test_patch_example_proto_depIdxs = nil
}
