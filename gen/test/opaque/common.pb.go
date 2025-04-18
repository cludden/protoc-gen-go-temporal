// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: test/opaque/common.proto

package opaque

import (
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

type Status int32

const (
	Status_STATUS_UNSPECIFIED Status = 0
	Status_STATUS_OK          Status = 1
	Status_STATUS_ERROR       Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_OK",
		2: "STATUS_ERROR",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_OK":          1,
		"STATUS_ERROR":       2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_test_opaque_common_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_test_opaque_common_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_test_opaque_common_proto_rawDescGZIP(), []int{0}
}

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Street        string                 `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City          string                 `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State         string                 `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Zip           string                 `protobuf:"bytes,4,opt,name=zip,proto3" json:"zip,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_test_opaque_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_test_opaque_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_test_opaque_common_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

var File_test_opaque_common_proto protoreflect.FileDescriptor

var file_test_opaque_common_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x22, 0x5d, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x7a, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x7a, 0x69, 0x70, 0x2a, 0x41, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x42, 0xa6, 0x01, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x42, 0x0b, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x75, 0x64, 0x64, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0xa2, 0x02, 0x03, 0x54, 0x4f, 0x58, 0xaa, 0x02, 0x0b,
	0x54, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0xca, 0x02, 0x0b, 0x54, 0x65,
	0x73, 0x74, 0x5c, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0xe2, 0x02, 0x17, 0x54, 0x65, 0x73, 0x74,
	0x5c, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x3a, 0x3a, 0x4f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_test_opaque_common_proto_rawDescOnce sync.Once
	file_test_opaque_common_proto_rawDescData []byte
)

func file_test_opaque_common_proto_rawDescGZIP() []byte {
	file_test_opaque_common_proto_rawDescOnce.Do(func() {
		file_test_opaque_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_test_opaque_common_proto_rawDesc), len(file_test_opaque_common_proto_rawDesc)))
	})
	return file_test_opaque_common_proto_rawDescData
}

var file_test_opaque_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_opaque_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_test_opaque_common_proto_goTypes = []any{
	(Status)(0),     // 0: test.opaque.Status
	(*Address)(nil), // 1: test.opaque.Address
}
var file_test_opaque_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_opaque_common_proto_init() }
func file_test_opaque_common_proto_init() {
	if File_test_opaque_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_test_opaque_common_proto_rawDesc), len(file_test_opaque_common_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_opaque_common_proto_goTypes,
		DependencyIndexes: file_test_opaque_common_proto_depIdxs,
		EnumInfos:         file_test_opaque_common_proto_enumTypes,
		MessageInfos:      file_test_opaque_common_proto_msgTypes,
	}.Build()
	File_test_opaque_common_proto = out.File
	file_test_opaque_common_proto_goTypes = nil
	file_test_opaque_common_proto_depIdxs = nil
}
