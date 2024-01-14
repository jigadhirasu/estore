// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: stock/supplier.proto

package stockpb

import (
	typepb "estore/protoc/typepb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Supplier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                // 供應商編號
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                            // 供應商名稱
	Phone     *typepb.Phone          `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`                          // 供應商聯繫電話
	Email     *typepb.Email          `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`                          // 供應商聯繫電子郵件
	Address   *typepb.Address        `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`                      // 供應商聯繫地址
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // 建立時間
}

func (x *Supplier) Reset() {
	*x = Supplier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_supplier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Supplier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Supplier) ProtoMessage() {}

func (x *Supplier) ProtoReflect() protoreflect.Message {
	mi := &file_stock_supplier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Supplier.ProtoReflect.Descriptor instead.
func (*Supplier) Descriptor() ([]byte, []int) {
	return file_stock_supplier_proto_rawDescGZIP(), []int{0}
}

func (x *Supplier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Supplier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Supplier) GetPhone() *typepb.Phone {
	if x != nil {
		return x.Phone
	}
	return nil
}

func (x *Supplier) GetEmail() *typepb.Email {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *Supplier) GetAddress() *typepb.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Supplier) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Supplier `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateSupplierRequest) Reset() {
	*x = CreateSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_supplier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSupplierRequest) ProtoMessage() {}

func (x *CreateSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_supplier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSupplierRequest.ProtoReflect.Descriptor instead.
func (*CreateSupplierRequest) Descriptor() ([]byte, []int) {
	return file_stock_supplier_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSupplierRequest) GetData() []*Supplier {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *typepb.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CreateSupplierResponse) Reset() {
	*x = CreateSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_supplier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSupplierResponse) ProtoMessage() {}

func (x *CreateSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_supplier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSupplierResponse.ProtoReflect.Descriptor instead.
func (*CreateSupplierResponse) Descriptor() ([]byte, []int) {
	return file_stock_supplier_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSupplierResponse) GetResponse() *typepb.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type UpdateSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Supplier `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateSupplierRequest) Reset() {
	*x = UpdateSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_supplier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSupplierRequest) ProtoMessage() {}

func (x *UpdateSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_supplier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSupplierRequest.ProtoReflect.Descriptor instead.
func (*UpdateSupplierRequest) Descriptor() ([]byte, []int) {
	return file_stock_supplier_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSupplierRequest) GetData() []*Supplier {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *typepb.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *UpdateSupplierResponse) Reset() {
	*x = UpdateSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_supplier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSupplierResponse) ProtoMessage() {}

func (x *UpdateSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_supplier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSupplierResponse.ProtoReflect.Descriptor instead.
func (*UpdateSupplierResponse) Descriptor() ([]byte, []int) {
	return file_stock_supplier_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSupplierResponse) GetResponse() *typepb.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type DeleteSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSupplierRequest) Reset() {
	*x = DeleteSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_supplier_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSupplierRequest) ProtoMessage() {}

func (x *DeleteSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_supplier_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSupplierRequest.ProtoReflect.Descriptor instead.
func (*DeleteSupplierRequest) Descriptor() ([]byte, []int) {
	return file_stock_supplier_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSupplierRequest) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

type DeleteSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *typepb.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DeleteSupplierResponse) Reset() {
	*x = DeleteSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_supplier_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSupplierResponse) ProtoMessage() {}

func (x *DeleteSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_supplier_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSupplierResponse.ProtoReflect.Descriptor instead.
func (*DeleteSupplierResponse) Descriptor() ([]byte, []int) {
	return file_stock_supplier_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSupplierResponse) GetResponse() *typepb.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type FindSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request   *typepb.Request    `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Id        []string           `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
	LikeName  string             `protobuf:"bytes,3,opt,name=like_name,json=likeName,proto3" json:"like_name,omitempty"`
	CreatedAt *typepb.Datepicker `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *FindSupplierRequest) Reset() {
	*x = FindSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_supplier_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindSupplierRequest) ProtoMessage() {}

func (x *FindSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_supplier_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindSupplierRequest.ProtoReflect.Descriptor instead.
func (*FindSupplierRequest) Descriptor() ([]byte, []int) {
	return file_stock_supplier_proto_rawDescGZIP(), []int{7}
}

func (x *FindSupplierRequest) GetRequest() *typepb.Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *FindSupplierRequest) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *FindSupplierRequest) GetLikeName() string {
	if x != nil {
		return x.LikeName
	}
	return ""
}

func (x *FindSupplierRequest) GetCreatedAt() *typepb.Datepicker {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type FindSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *typepb.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Data     []*Supplier      `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *FindSupplierResponse) Reset() {
	*x = FindSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_supplier_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindSupplierResponse) ProtoMessage() {}

func (x *FindSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_supplier_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindSupplierResponse.ProtoReflect.Descriptor instead.
func (*FindSupplierResponse) Descriptor() ([]byte, []int) {
	return file_stock_supplier_proto_rawDescGZIP(), []int{8}
}

func (x *FindSupplierResponse) GetResponse() *typepb.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *FindSupplierResponse) GetData() []*Supplier {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_stock_supplier_proto protoreflect.FileDescriptor

var file_stock_supplier_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x1a, 0x11, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb,
	0x01, 0x0a, 0x08, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3c, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x45, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3c, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x45, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x45, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x69, 0x6b, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x69, 0x6b,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x70, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x68, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x17, 0x5a, 0x15, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_stock_supplier_proto_rawDescOnce sync.Once
	file_stock_supplier_proto_rawDescData = file_stock_supplier_proto_rawDesc
)

func file_stock_supplier_proto_rawDescGZIP() []byte {
	file_stock_supplier_proto_rawDescOnce.Do(func() {
		file_stock_supplier_proto_rawDescData = protoimpl.X.CompressGZIP(file_stock_supplier_proto_rawDescData)
	})
	return file_stock_supplier_proto_rawDescData
}

var file_stock_supplier_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_stock_supplier_proto_goTypes = []interface{}{
	(*Supplier)(nil),               // 0: stock.Supplier
	(*CreateSupplierRequest)(nil),  // 1: stock.CreateSupplierRequest
	(*CreateSupplierResponse)(nil), // 2: stock.CreateSupplierResponse
	(*UpdateSupplierRequest)(nil),  // 3: stock.UpdateSupplierRequest
	(*UpdateSupplierResponse)(nil), // 4: stock.UpdateSupplierResponse
	(*DeleteSupplierRequest)(nil),  // 5: stock.DeleteSupplierRequest
	(*DeleteSupplierResponse)(nil), // 6: stock.DeleteSupplierResponse
	(*FindSupplierRequest)(nil),    // 7: stock.FindSupplierRequest
	(*FindSupplierResponse)(nil),   // 8: stock.FindSupplierResponse
	(*typepb.Phone)(nil),           // 9: types.Phone
	(*typepb.Email)(nil),           // 10: types.Email
	(*typepb.Address)(nil),         // 11: types.Address
	(*timestamppb.Timestamp)(nil),  // 12: google.protobuf.Timestamp
	(*typepb.Response)(nil),        // 13: types.Response
	(*typepb.Request)(nil),         // 14: types.Request
	(*typepb.Datepicker)(nil),      // 15: types.Datepicker
}
var file_stock_supplier_proto_depIdxs = []int32{
	9,  // 0: stock.Supplier.phone:type_name -> types.Phone
	10, // 1: stock.Supplier.email:type_name -> types.Email
	11, // 2: stock.Supplier.address:type_name -> types.Address
	12, // 3: stock.Supplier.created_at:type_name -> google.protobuf.Timestamp
	0,  // 4: stock.CreateSupplierRequest.data:type_name -> stock.Supplier
	13, // 5: stock.CreateSupplierResponse.response:type_name -> types.Response
	0,  // 6: stock.UpdateSupplierRequest.data:type_name -> stock.Supplier
	13, // 7: stock.UpdateSupplierResponse.response:type_name -> types.Response
	13, // 8: stock.DeleteSupplierResponse.response:type_name -> types.Response
	14, // 9: stock.FindSupplierRequest.request:type_name -> types.Request
	15, // 10: stock.FindSupplierRequest.created_at:type_name -> types.Datepicker
	13, // 11: stock.FindSupplierResponse.response:type_name -> types.Response
	0,  // 12: stock.FindSupplierResponse.data:type_name -> stock.Supplier
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_stock_supplier_proto_init() }
func file_stock_supplier_proto_init() {
	if File_stock_supplier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stock_supplier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Supplier); i {
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
		file_stock_supplier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSupplierRequest); i {
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
		file_stock_supplier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSupplierResponse); i {
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
		file_stock_supplier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSupplierRequest); i {
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
		file_stock_supplier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSupplierResponse); i {
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
		file_stock_supplier_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSupplierRequest); i {
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
		file_stock_supplier_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSupplierResponse); i {
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
		file_stock_supplier_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindSupplierRequest); i {
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
		file_stock_supplier_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindSupplierResponse); i {
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
			RawDescriptor: file_stock_supplier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stock_supplier_proto_goTypes,
		DependencyIndexes: file_stock_supplier_proto_depIdxs,
		MessageInfos:      file_stock_supplier_proto_msgTypes,
	}.Build()
	File_stock_supplier_proto = out.File
	file_stock_supplier_proto_rawDesc = nil
	file_stock_supplier_proto_goTypes = nil
	file_stock_supplier_proto_depIdxs = nil
}