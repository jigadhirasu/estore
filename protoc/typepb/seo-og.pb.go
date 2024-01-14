// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: types/seo-og.proto

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

type SEO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Keywords    string `protobuf:"bytes,2,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Jsonld      string `protobuf:"bytes,4,opt,name=jsonld,proto3" json:"jsonld,omitempty"`
}

func (x *SEO) Reset() {
	*x = SEO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_seo_og_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SEO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SEO) ProtoMessage() {}

func (x *SEO) ProtoReflect() protoreflect.Message {
	mi := &file_types_seo_og_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SEO.ProtoReflect.Descriptor instead.
func (*SEO) Descriptor() ([]byte, []int) {
	return file_types_seo_og_proto_rawDescGZIP(), []int{0}
}

func (x *SEO) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SEO) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *SEO) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SEO) GetJsonld() string {
	if x != nil {
		return x.Jsonld
	}
	return ""
}

type OG struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Url         string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Image       string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *OG) Reset() {
	*x = OG{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_seo_og_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OG) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OG) ProtoMessage() {}

func (x *OG) ProtoReflect() protoreflect.Message {
	mi := &file_types_seo_og_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OG.ProtoReflect.Descriptor instead.
func (*OG) Descriptor() ([]byte, []int) {
	return file_types_seo_og_proto_rawDescGZIP(), []int{1}
}

func (x *OG) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OG) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *OG) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *OG) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *OG) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_types_seo_og_proto protoreflect.FileDescriptor

var file_types_seo_og_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x6f, 0x2d, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x71, 0x0a, 0x03, 0x53,
	0x45, 0x4f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x73, 0x6f, 0x6e, 0x6c, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6a, 0x73, 0x6f, 0x6e, 0x6c, 0x64, 0x22, 0x78,
	0x0a, 0x02, 0x4f, 0x47, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x16, 0x5a, 0x14, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_seo_og_proto_rawDescOnce sync.Once
	file_types_seo_og_proto_rawDescData = file_types_seo_og_proto_rawDesc
)

func file_types_seo_og_proto_rawDescGZIP() []byte {
	file_types_seo_og_proto_rawDescOnce.Do(func() {
		file_types_seo_og_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_seo_og_proto_rawDescData)
	})
	return file_types_seo_og_proto_rawDescData
}

var file_types_seo_og_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_types_seo_og_proto_goTypes = []interface{}{
	(*SEO)(nil), // 0: types.SEO
	(*OG)(nil),  // 1: types.OG
}
var file_types_seo_og_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_seo_og_proto_init() }
func file_types_seo_og_proto_init() {
	if File_types_seo_og_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_seo_og_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SEO); i {
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
		file_types_seo_og_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OG); i {
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
			RawDescriptor: file_types_seo_og_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_seo_og_proto_goTypes,
		DependencyIndexes: file_types_seo_og_proto_depIdxs,
		MessageInfos:      file_types_seo_og_proto_msgTypes,
	}.Build()
	File_types_seo_og_proto = out.File
	file_types_seo_og_proto_rawDesc = nil
	file_types_seo_og_proto_goTypes = nil
	file_types_seo_og_proto_depIdxs = nil
}
