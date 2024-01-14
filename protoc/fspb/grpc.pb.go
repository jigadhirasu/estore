// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: fs/grpc.proto

package fspb

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

type PingPong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingPong) Reset() {
	*x = PingPong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fs_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingPong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingPong) ProtoMessage() {}

func (x *PingPong) ProtoReflect() protoreflect.Message {
	mi := &file_fs_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingPong.ProtoReflect.Descriptor instead.
func (*PingPong) Descriptor() ([]byte, []int) {
	return file_fs_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *PingPong) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_fs_grpc_proto protoreflect.FileDescriptor

var file_fs_grpc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x66, 0x73, 0x1a, 0x11, 0x66, 0x73, 0x2f, 0x64, 0x69, 0x72, 0x2d, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f,
	0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc0, 0x03, 0x0a,
	0x0a, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x0c, 0x2e, 0x66, 0x73, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e,
	0x67, 0x1a, 0x0c, 0x2e, 0x66, 0x73, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12,
	0x54, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x20, 0x2e, 0x66, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x69, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x66, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x73, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x66, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x69, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x66, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x69, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x69, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16,
	0x2e, 0x66, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x69, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x66, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x44, 0x69, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x14, 0x5a, 0x12, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2f, 0x66, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fs_grpc_proto_rawDescOnce sync.Once
	file_fs_grpc_proto_rawDescData = file_fs_grpc_proto_rawDesc
)

func file_fs_grpc_proto_rawDescGZIP() []byte {
	file_fs_grpc_proto_rawDescOnce.Do(func() {
		file_fs_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_fs_grpc_proto_rawDescData)
	})
	return file_fs_grpc_proto_rawDescData
}

var file_fs_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fs_grpc_proto_goTypes = []interface{}{
	(*PingPong)(nil),                      // 0: fs.PingPong
	(*CreateOrUpdateDirFileRequest)(nil),  // 1: fs.CreateOrUpdateDirFileRequest
	(*DeleteDirFileRequest)(nil),          // 2: fs.DeleteDirFileRequest
	(*FindDirFileRequest)(nil),            // 3: fs.FindDirFileRequest
	(*CreateOrUpdateDirFileResponse)(nil), // 4: fs.CreateOrUpdateDirFileResponse
	(*DeleteDirFileResponse)(nil),         // 5: fs.DeleteDirFileResponse
	(*FindDirFileResponse)(nil),           // 6: fs.FindDirFileResponse
}
var file_fs_grpc_proto_depIdxs = []int32{
	0, // 0: fs.FileSystem.Ping:input_type -> fs.PingPong
	1, // 1: fs.FileSystem.CreateDirFile:input_type -> fs.CreateOrUpdateDirFileRequest
	1, // 2: fs.FileSystem.UpdateDirFile:input_type -> fs.CreateOrUpdateDirFileRequest
	1, // 3: fs.FileSystem.CreateOrUpdateDirFile:input_type -> fs.CreateOrUpdateDirFileRequest
	2, // 4: fs.FileSystem.DeleteDirFile:input_type -> fs.DeleteDirFileRequest
	3, // 5: fs.FileSystem.FindDirFile:input_type -> fs.FindDirFileRequest
	0, // 6: fs.FileSystem.Ping:output_type -> fs.PingPong
	4, // 7: fs.FileSystem.CreateDirFile:output_type -> fs.CreateOrUpdateDirFileResponse
	4, // 8: fs.FileSystem.UpdateDirFile:output_type -> fs.CreateOrUpdateDirFileResponse
	4, // 9: fs.FileSystem.CreateOrUpdateDirFile:output_type -> fs.CreateOrUpdateDirFileResponse
	5, // 10: fs.FileSystem.DeleteDirFile:output_type -> fs.DeleteDirFileResponse
	6, // 11: fs.FileSystem.FindDirFile:output_type -> fs.FindDirFileResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fs_grpc_proto_init() }
func file_fs_grpc_proto_init() {
	if File_fs_grpc_proto != nil {
		return
	}
	file_fs_dir_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fs_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingPong); i {
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
			RawDescriptor: file_fs_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fs_grpc_proto_goTypes,
		DependencyIndexes: file_fs_grpc_proto_depIdxs,
		MessageInfos:      file_fs_grpc_proto_msgTypes,
	}.Build()
	File_fs_grpc_proto = out.File
	file_fs_grpc_proto_rawDesc = nil
	file_fs_grpc_proto_goTypes = nil
	file_fs_grpc_proto_depIdxs = nil
}
