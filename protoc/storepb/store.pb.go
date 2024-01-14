// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: store/store.proto

package storepb

import (
	typepb "estore/protoc/typepb"
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

type Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                  // 商店編號
	Name     string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                              // 商店名稱
	Type     Type            `protobuf:"varint,3,opt,name=type,proto3,enum=store.Type" json:"type,omitempty"`             // 商店類型
	State    State           `protobuf:"varint,4,opt,name=state,proto3,enum=store.State" json:"state,omitempty"`          // 商店狀態
	ClubId   string          `protobuf:"bytes,5,opt,name=club_id,json=clubId,proto3" json:"club_id,omitempty"`            // 群組編號
	Currency typepb.Currency `protobuf:"varint,6,opt,name=currency,proto3,enum=types.Currency" json:"currency,omitempty"` // 商店貨幣
	Info     *StoreInfo      `protobuf:"bytes,7,opt,name=info,proto3" json:"info,omitempty"`                              // 商店資訊
}

func (x *Store) Reset() {
	*x = Store{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store) ProtoMessage() {}

func (x *Store) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store.ProtoReflect.Descriptor instead.
func (*Store) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{0}
}

func (x *Store) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Store) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Store) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_NoType
}

func (x *Store) GetState() State {
	if x != nil {
		return x.State
	}
	return State_None
}

func (x *Store) GetClubId() string {
	if x != nil {
		return x.ClubId
	}
	return ""
}

func (x *Store) GetCurrency() typepb.Currency {
	if x != nil {
		return x.Currency
	}
	return typepb.Currency(0)
}

func (x *Store) GetInfo() *StoreInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type StoreInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Copyright     string          `protobuf:"bytes,1,opt,name=copyright,proto3" json:"copyright,omitempty"`                              // 版權聲明
	Phone         *typepb.Phone   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`                                      // 聯絡電話
	Email         *typepb.Email   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`                                      // 聯絡信箱
	Address       *typepb.Address `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`                                  // 聯絡地址
	BusinessHours string          `protobuf:"bytes,5,opt,name=business_hours,json=businessHours,proto3" json:"business_hours,omitempty"` // 營業時間
	MaxCount      int32           `protobuf:"varint,6,opt,name=max_count,json=maxCount,proto3" json:"max_count,omitempty"`               // 單品最大數量
	AllowCancel   bool            `protobuf:"varint,7,opt,name=allow_cancel,json=allowCancel,proto3" json:"allow_cancel,omitempty"`      // 是否允許會員取消訂單
	AllowReturn   bool            `protobuf:"varint,8,opt,name=allow_return,json=allowReturn,proto3" json:"allow_return,omitempty"`      // 是否允許會員申請退貨
	DisplayStock  bool            `protobuf:"varint,9,opt,name=display_stock,json=displayStock,proto3" json:"display_stock,omitempty"`   // 是否顯示庫存
}

func (x *StoreInfo) Reset() {
	*x = StoreInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreInfo) ProtoMessage() {}

func (x *StoreInfo) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreInfo.ProtoReflect.Descriptor instead.
func (*StoreInfo) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{1}
}

func (x *StoreInfo) GetCopyright() string {
	if x != nil {
		return x.Copyright
	}
	return ""
}

func (x *StoreInfo) GetPhone() *typepb.Phone {
	if x != nil {
		return x.Phone
	}
	return nil
}

func (x *StoreInfo) GetEmail() *typepb.Email {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *StoreInfo) GetAddress() *typepb.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *StoreInfo) GetBusinessHours() string {
	if x != nil {
		return x.BusinessHours
	}
	return ""
}

func (x *StoreInfo) GetMaxCount() int32 {
	if x != nil {
		return x.MaxCount
	}
	return 0
}

func (x *StoreInfo) GetAllowCancel() bool {
	if x != nil {
		return x.AllowCancel
	}
	return false
}

func (x *StoreInfo) GetAllowReturn() bool {
	if x != nil {
		return x.AllowReturn
	}
	return false
}

func (x *StoreInfo) GetDisplayStock() bool {
	if x != nil {
		return x.DisplayStock
	}
	return false
}

var File_store_store_proto protoreflect.FileDescriptor

var file_store_store_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x11, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdc, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x22, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x75, 0x62, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0xca, 0x02, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x22, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48,
	0x6f, 0x75, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x42, 0x17, 0x5a, 0x15,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_store_proto_rawDescOnce sync.Once
	file_store_store_proto_rawDescData = file_store_store_proto_rawDesc
)

func file_store_store_proto_rawDescGZIP() []byte {
	file_store_store_proto_rawDescOnce.Do(func() {
		file_store_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_store_proto_rawDescData)
	})
	return file_store_store_proto_rawDescData
}

var file_store_store_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_store_store_proto_goTypes = []interface{}{
	(*Store)(nil),          // 0: store.Store
	(*StoreInfo)(nil),      // 1: store.StoreInfo
	(Type)(0),              // 2: store.Type
	(State)(0),             // 3: store.State
	(typepb.Currency)(0),   // 4: types.Currency
	(*typepb.Phone)(nil),   // 5: types.Phone
	(*typepb.Email)(nil),   // 6: types.Email
	(*typepb.Address)(nil), // 7: types.Address
}
var file_store_store_proto_depIdxs = []int32{
	2, // 0: store.Store.type:type_name -> store.Type
	3, // 1: store.Store.state:type_name -> store.State
	4, // 2: store.Store.currency:type_name -> types.Currency
	1, // 3: store.Store.info:type_name -> store.StoreInfo
	5, // 4: store.StoreInfo.phone:type_name -> types.Phone
	6, // 5: store.StoreInfo.email:type_name -> types.Email
	7, // 6: store.StoreInfo.address:type_name -> types.Address
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_store_store_proto_init() }
func file_store_store_proto_init() {
	if File_store_store_proto != nil {
		return
	}
	file_store_state_proto_init()
	file_store_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_store_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store); i {
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
		file_store_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreInfo); i {
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
			RawDescriptor: file_store_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_store_proto_goTypes,
		DependencyIndexes: file_store_store_proto_depIdxs,
		MessageInfos:      file_store_store_proto_msgTypes,
	}.Build()
	File_store_store_proto = out.File
	file_store_store_proto_rawDesc = nil
	file_store_store_proto_goTypes = nil
	file_store_store_proto_depIdxs = nil
}