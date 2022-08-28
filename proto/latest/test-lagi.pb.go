// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: latest/test-lagi.proto

package proto

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TestRequestLagi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UserID int64  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *TestRequestLagi) Reset() {
	*x = TestRequestLagi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_latest_test_lagi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRequestLagi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequestLagi) ProtoMessage() {}

func (x *TestRequestLagi) ProtoReflect() protoreflect.Message {
	mi := &file_latest_test_lagi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequestLagi.ProtoReflect.Descriptor instead.
func (*TestRequestLagi) Descriptor() ([]byte, []int) {
	return file_latest_test_lagi_proto_rawDescGZIP(), []int{0}
}

func (x *TestRequestLagi) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestRequestLagi) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type TestResponseLagi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TestResponseLagi) Reset() {
	*x = TestResponseLagi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_latest_test_lagi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResponseLagi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResponseLagi) ProtoMessage() {}

func (x *TestResponseLagi) ProtoReflect() protoreflect.Message {
	mi := &file_latest_test_lagi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResponseLagi.ProtoReflect.Descriptor instead.
func (*TestResponseLagi) Descriptor() ([]byte, []int) {
	return file_latest_test_lagi_proto_rawDescGZIP(), []int{1}
}

func (x *TestResponseLagi) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_latest_test_lagi_proto protoreflect.FileDescriptor

var file_latest_test_lagi_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x6c, 0x61,
	0x67, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a,
	0x0f, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x67, 0x69,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2a, 0x0a, 0x10,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x61, 0x67, 0x69,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x9b, 0x01, 0x0a, 0x0f, 0x54, 0x65, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x61, 0x67, 0x69, 0x12, 0x87, 0x01, 0x0a,
	0x14, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x57, 0x69, 0x74, 0x68,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x67, 0x69, 0x1a, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4c, 0x61, 0x67, 0x69, 0x22, 0x3e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x6c, 0x61, 0x67, 0x69,
	0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x7d, 0x92, 0x41, 0x1a, 0x0a, 0x04, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x08, 0x54, 0x65, 0x73, 0x74, 0x20, 0x41, 0x50, 0x49, 0x1a, 0x08, 0x54, 0x65,
	0x73, 0x74, 0x20, 0x41, 0x50, 0x49, 0x42, 0x90, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x79, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2d,
	0x68, 0x74, 0x74, 0x70, 0x2d, 0x6f, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x92, 0x41, 0x55, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x72, 0x49,
	0x0a, 0x1a, 0x47, 0x4f, 0x20, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2b, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x77, 0x69, 0x64, 0x79, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x68, 0x74, 0x74, 0x70, 0x2d,
	0x6f, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_latest_test_lagi_proto_rawDescOnce sync.Once
	file_latest_test_lagi_proto_rawDescData = file_latest_test_lagi_proto_rawDesc
)

func file_latest_test_lagi_proto_rawDescGZIP() []byte {
	file_latest_test_lagi_proto_rawDescOnce.Do(func() {
		file_latest_test_lagi_proto_rawDescData = protoimpl.X.CompressGZIP(file_latest_test_lagi_proto_rawDescData)
	})
	return file_latest_test_lagi_proto_rawDescData
}

var file_latest_test_lagi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_latest_test_lagi_proto_goTypes = []interface{}{
	(*TestRequestLagi)(nil),  // 0: proto.TestRequestLagi
	(*TestResponseLagi)(nil), // 1: proto.TestResponseLagi
}
var file_latest_test_lagi_proto_depIdxs = []int32{
	0, // 0: proto.TestServiceLagi.TestServiceWithParam:input_type -> proto.TestRequestLagi
	1, // 1: proto.TestServiceLagi.TestServiceWithParam:output_type -> proto.TestResponseLagi
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_latest_test_lagi_proto_init() }
func file_latest_test_lagi_proto_init() {
	if File_latest_test_lagi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_latest_test_lagi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRequestLagi); i {
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
		file_latest_test_lagi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResponseLagi); i {
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
			RawDescriptor: file_latest_test_lagi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_latest_test_lagi_proto_goTypes,
		DependencyIndexes: file_latest_test_lagi_proto_depIdxs,
		MessageInfos:      file_latest_test_lagi_proto_msgTypes,
	}.Build()
	File_latest_test_lagi_proto = out.File
	file_latest_test_lagi_proto_rawDesc = nil
	file_latest_test_lagi_proto_goTypes = nil
	file_latest_test_lagi_proto_depIdxs = nil
}