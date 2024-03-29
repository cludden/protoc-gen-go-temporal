// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: example/helloworld/v1/helloworld.proto

package helloworldv1

import (
	_ "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
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

type HelloWorldInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloWorldInput) Reset() {
	*x = HelloWorldInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_helloworld_v1_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldInput) ProtoMessage() {}

func (x *HelloWorldInput) ProtoReflect() protoreflect.Message {
	mi := &file_example_helloworld_v1_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldInput.ProtoReflect.Descriptor instead.
func (*HelloWorldInput) Descriptor() ([]byte, []int) {
	return file_example_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *HelloWorldInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloWorldOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *HelloWorldOutput) Reset() {
	*x = HelloWorldOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_helloworld_v1_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldOutput) ProtoMessage() {}

func (x *HelloWorldOutput) ProtoReflect() protoreflect.Message {
	mi := &file_example_helloworld_v1_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldOutput.ProtoReflect.Descriptor instead.
func (*HelloWorldOutput) Descriptor() ([]byte, []int) {
	return file_example_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *HelloWorldOutput) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_example_helloworld_v1_helloworld_proto protoreflect.FileDescriptor

var file_example_helloworld_v1_helloworld_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a,
	0x1a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0f, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x10, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xa7,
	0x01, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x0a, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x26, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x27, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x29, 0x8a, 0xc4, 0x03, 0x1d,
	0x2a, 0x1b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x24, 0x7b,
	0x21, 0x20, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x76, 0x34, 0x28, 0x29, 0x20, 0x7d, 0x92, 0xc4, 0x03,
	0x04, 0x22, 0x02, 0x08, 0x0a, 0x1a, 0x11, 0x8a, 0xc4, 0x03, 0x0d, 0x0a, 0x0b, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2d, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x42, 0xf4, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x75, 0x64, 0x64, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x48,
	0x58, 0xaa, 0x02, 0x15, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x5c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x21, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a,
	0x3a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_helloworld_v1_helloworld_proto_rawDescOnce sync.Once
	file_example_helloworld_v1_helloworld_proto_rawDescData = file_example_helloworld_v1_helloworld_proto_rawDesc
)

func file_example_helloworld_v1_helloworld_proto_rawDescGZIP() []byte {
	file_example_helloworld_v1_helloworld_proto_rawDescOnce.Do(func() {
		file_example_helloworld_v1_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_helloworld_v1_helloworld_proto_rawDescData)
	})
	return file_example_helloworld_v1_helloworld_proto_rawDescData
}

var file_example_helloworld_v1_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_helloworld_v1_helloworld_proto_goTypes = []interface{}{
	(*HelloWorldInput)(nil),  // 0: example.helloworld.v1.HelloWorldInput
	(*HelloWorldOutput)(nil), // 1: example.helloworld.v1.HelloWorldOutput
}
var file_example_helloworld_v1_helloworld_proto_depIdxs = []int32{
	0, // 0: example.helloworld.v1.Example.HelloWorld:input_type -> example.helloworld.v1.HelloWorldInput
	1, // 1: example.helloworld.v1.Example.HelloWorld:output_type -> example.helloworld.v1.HelloWorldOutput
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_helloworld_v1_helloworld_proto_init() }
func file_example_helloworld_v1_helloworld_proto_init() {
	if File_example_helloworld_v1_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_helloworld_v1_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldInput); i {
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
		file_example_helloworld_v1_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldOutput); i {
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
			RawDescriptor: file_example_helloworld_v1_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_helloworld_v1_helloworld_proto_goTypes,
		DependencyIndexes: file_example_helloworld_v1_helloworld_proto_depIdxs,
		MessageInfos:      file_example_helloworld_v1_helloworld_proto_msgTypes,
	}.Build()
	File_example_helloworld_v1_helloworld_proto = out.File
	file_example_helloworld_v1_helloworld_proto_rawDesc = nil
	file_example_helloworld_v1_helloworld_proto_goTypes = nil
	file_example_helloworld_v1_helloworld_proto_depIdxs = nil
}
