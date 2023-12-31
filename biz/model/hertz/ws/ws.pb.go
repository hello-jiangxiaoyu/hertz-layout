// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.9
// source: ws.proto

package ws

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "hertz/demo/biz/model/api"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoReq) Reset() {
	*x = NoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ws_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoReq) ProtoMessage() {}

func (x *NoReq) ProtoReflect() protoreflect.Message {
	mi := &file_ws_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoReq.ProtoReflect.Descriptor instead.
func (*NoReq) Descriptor() ([]byte, []int) {
	return file_ws_proto_rawDescGZIP(), []int{0}
}

var File_ws_proto protoreflect.FileDescriptor

var file_ws_proto_rawDesc = []byte{
	0x0a, 0x08, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x77, 0x73, 0x1a, 0x09,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x4e, 0x6f, 0x52,
	0x65, 0x71, 0x32, 0x39, 0x0a, 0x03, 0x41, 0x70, 0x69, 0x12, 0x32, 0x0a, 0x08, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x09, 0x2e, 0x77, 0x73, 0x2e, 0x4e, 0x6f, 0x52, 0x65, 0x71,
	0x1a, 0x09, 0x2e, 0x77, 0x73, 0x2e, 0x4e, 0x6f, 0x52, 0x65, 0x71, 0x22, 0x10, 0xca, 0xc1, 0x18,
	0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x2f, 0x77, 0x73, 0x42, 0x1f, 0x5a,
	0x1d, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x62, 0x69, 0x7a, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x2f, 0x77, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ws_proto_rawDescOnce sync.Once
	file_ws_proto_rawDescData = file_ws_proto_rawDesc
)

func file_ws_proto_rawDescGZIP() []byte {
	file_ws_proto_rawDescOnce.Do(func() {
		file_ws_proto_rawDescData = protoimpl.X.CompressGZIP(file_ws_proto_rawDescData)
	})
	return file_ws_proto_rawDescData
}

var file_ws_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ws_proto_goTypes = []interface{}{
	(*NoReq)(nil), // 0: ws.NoReq
}
var file_ws_proto_depIdxs = []int32{
	0, // 0: ws.Api.ChatRoom:input_type -> ws.NoReq
	0, // 1: ws.Api.ChatRoom:output_type -> ws.NoReq
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ws_proto_init() }
func file_ws_proto_init() {
	if File_ws_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ws_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoReq); i {
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
			RawDescriptor: file_ws_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ws_proto_goTypes,
		DependencyIndexes: file_ws_proto_depIdxs,
		MessageInfos:      file_ws_proto_msgTypes,
	}.Build()
	File_ws_proto = out.File
	file_ws_proto_rawDesc = nil
	file_ws_proto_goTypes = nil
	file_ws_proto_depIdxs = nil
}
