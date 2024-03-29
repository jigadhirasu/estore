// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: types/phone.proto

package typepb

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

type Phone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Number   string `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Phone) Reset() {
	*x = Phone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_phone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Phone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Phone) ProtoMessage() {}

func (x *Phone) ProtoReflect() protoreflect.Message {
	mi := &file_types_phone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Phone.ProtoReflect.Descriptor instead.
func (*Phone) Descriptor() ([]byte, []int) {
	return file_types_phone_proto_rawDescGZIP(), []int{0}
}

func (x *Phone) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Phone) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

var File_types_phone_proto protoreflect.FileDescriptor

var file_types_phone_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x05, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x16, 0x5a, 0x14, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_phone_proto_rawDescOnce sync.Once
	file_types_phone_proto_rawDescData = file_types_phone_proto_rawDesc
)

func file_types_phone_proto_rawDescGZIP() []byte {
	file_types_phone_proto_rawDescOnce.Do(func() {
		file_types_phone_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_phone_proto_rawDescData)
	})
	return file_types_phone_proto_rawDescData
}

var file_types_phone_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_types_phone_proto_goTypes = []interface{}{
	(*Phone)(nil), // 0: types.Phone
}
var file_types_phone_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_phone_proto_init() }
func file_types_phone_proto_init() {
	if File_types_phone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_phone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Phone); i {
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
			RawDescriptor: file_types_phone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_phone_proto_goTypes,
		DependencyIndexes: file_types_phone_proto_depIdxs,
		MessageInfos:      file_types_phone_proto_msgTypes,
	}.Build()
	File_types_phone_proto = out.File
	file_types_phone_proto_rawDesc = nil
	file_types_phone_proto_goTypes = nil
	file_types_phone_proto_depIdxs = nil
}
