// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: service/messages.proto

package service

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

// The status request message (REST request: GET /api/status)
type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_service_messages_proto_rawDescGZIP(), []int{0}
}

// The response message containing the status (response to StatusRequest)
type StatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StatusReply) Reset() {
	*x = StatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusReply) ProtoMessage() {}

func (x *StatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusReply.ProtoReflect.Descriptor instead.
func (*StatusReply) Descriptor() ([]byte, []int) {
	return file_service_messages_proto_rawDescGZIP(), []int{1}
}

func (x *StatusReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// The status request message (gRPC)
type ProductStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProductStatusRequest) Reset() {
	*x = ProductStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductStatusRequest) ProtoMessage() {}

func (x *ProductStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductStatusRequest.ProtoReflect.Descriptor instead.
func (*ProductStatusRequest) Descriptor() ([]byte, []int) {
	return file_service_messages_proto_rawDescGZIP(), []int{2}
}

// The response to the product status request
type ProductStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ProductStatusReply) Reset() {
	*x = ProductStatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductStatusReply) ProtoMessage() {}

func (x *ProductStatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductStatusReply.ProtoReflect.Descriptor instead.
func (*ProductStatusReply) Descriptor() ([]byte, []int) {
	return file_service_messages_proto_rawDescGZIP(), []int{3}
}

func (x *ProductStatusReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetSingleProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Scope []string `protobuf:"bytes,2,rep,name=scope,proto3" json:"scope,omitempty"`
}

func (x *GetSingleProductRequest) Reset() {
	*x = GetSingleProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSingleProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSingleProductRequest) ProtoMessage() {}

func (x *GetSingleProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSingleProductRequest.ProtoReflect.Descriptor instead.
func (*GetSingleProductRequest) Descriptor() ([]byte, []int) {
	return file_service_messages_proto_rawDescGZIP(), []int{4}
}

func (x *GetSingleProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetSingleProductRequest) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

type GetProductsInScopeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scope []string `protobuf:"bytes,1,rep,name=scope,proto3" json:"scope,omitempty"`
}

func (x *GetProductsInScopeRequest) Reset() {
	*x = GetProductsInScopeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductsInScopeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsInScopeRequest) ProtoMessage() {}

func (x *GetProductsInScopeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsInScopeRequest.ProtoReflect.Descriptor instead.
func (*GetProductsInScopeRequest) Descriptor() ([]byte, []int) {
	return file_service_messages_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductsInScopeRequest) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

type StoredProducts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*StoredProduct `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *StoredProducts) Reset() {
	*x = StoredProducts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredProducts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredProducts) ProtoMessage() {}

func (x *StoredProducts) ProtoReflect() protoreflect.Message {
	mi := &file_service_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredProducts.ProtoReflect.Descriptor instead.
func (*StoredProducts) Descriptor() ([]byte, []int) {
	return file_service_messages_proto_rawDescGZIP(), []int{6}
}

func (x *StoredProducts) GetProducts() []*StoredProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

type PutSingleProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Scope   []string `protobuf:"bytes,2,rep,name=scope,proto3" json:"scope,omitempty"`
	Expires *int64   `protobuf:"varint,3,opt,name=expires,proto3,oneof" json:"expires,omitempty"`
}

func (x *PutSingleProductRequest) Reset() {
	*x = PutSingleProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutSingleProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutSingleProductRequest) ProtoMessage() {}

func (x *PutSingleProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutSingleProductRequest.ProtoReflect.Descriptor instead.
func (*PutSingleProductRequest) Descriptor() ([]byte, []int) {
	return file_service_messages_proto_rawDescGZIP(), []int{7}
}

func (x *PutSingleProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PutSingleProductRequest) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *PutSingleProductRequest) GetExpires() int64 {
	if x != nil && x.Expires != nil {
		return *x.Expires
	}
	return 0
}

type StoredProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Scope   []string `protobuf:"bytes,2,rep,name=scope,proto3" json:"scope,omitempty"`
	Data    string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Expires int64    `protobuf:"varint,4,opt,name=expires,proto3" json:"expires,omitempty"`
}

func (x *StoredProduct) Reset() {
	*x = StoredProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredProduct) ProtoMessage() {}

func (x *StoredProduct) ProtoReflect() protoreflect.Message {
	mi := &file_service_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredProduct.ProtoReflect.Descriptor instead.
func (*StoredProduct) Descriptor() ([]byte, []int) {
	return file_service_messages_proto_rawDescGZIP(), []int{8}
}

func (x *StoredProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoredProduct) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *StoredProduct) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *StoredProduct) GetExpires() int64 {
	if x != nil {
		return x.Expires
	}
	return 0
}

type ClearSingleProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Scope []string `protobuf:"bytes,2,rep,name=scope,proto3" json:"scope,omitempty"`
}

func (x *ClearSingleProductRequest) Reset() {
	*x = ClearSingleProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_messages_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearSingleProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearSingleProductRequest) ProtoMessage() {}

func (x *ClearSingleProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_messages_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearSingleProductRequest.ProtoReflect.Descriptor instead.
func (*ClearSingleProductRequest) Descriptor() ([]byte, []int) {
	return file_service_messages_proto_rawDescGZIP(), []int{9}
}

func (x *ClearSingleProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClearSingleProductRequest) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

type ClearSingleProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeletedName string   `protobuf:"bytes,1,opt,name=deletedName,proto3" json:"deletedName,omitempty"`
	Scope       []string `protobuf:"bytes,2,rep,name=scope,proto3" json:"scope,omitempty"`
}

func (x *ClearSingleProductResponse) Reset() {
	*x = ClearSingleProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_messages_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearSingleProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearSingleProductResponse) ProtoMessage() {}

func (x *ClearSingleProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_messages_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearSingleProductResponse.ProtoReflect.Descriptor instead.
func (*ClearSingleProductResponse) Descriptor() ([]byte, []int) {
	return file_service_messages_proto_rawDescGZIP(), []int{10}
}

func (x *ClearSingleProductResponse) GetDeletedName() string {
	if x != nil {
		return x.DeletedName
	}
	return ""
}

func (x *ClearSingleProductResponse) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

var File_service_messages_proto protoreflect.FileDescriptor

var file_service_messages_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x2c, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x43, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x22, 0x31, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x49, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x44, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x6e, 0x0a,
	0x17, 0x50, 0x75, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x88, 0x01,
	0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x22, 0x67, 0x0a,
	0x0d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x19, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x54, 0x0a,
	0x1a, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x43, 0x53, 0x34, 0x36, 0x37, 0x5f, 0x53, 0x55, 0x32,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_messages_proto_rawDescOnce sync.Once
	file_service_messages_proto_rawDescData = file_service_messages_proto_rawDesc
)

func file_service_messages_proto_rawDescGZIP() []byte {
	file_service_messages_proto_rawDescOnce.Do(func() {
		file_service_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_messages_proto_rawDescData)
	})
	return file_service_messages_proto_rawDescData
}

var file_service_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_service_messages_proto_goTypes = []interface{}{
	(*StatusRequest)(nil),              // 0: service.StatusRequest
	(*StatusReply)(nil),                // 1: service.StatusReply
	(*ProductStatusRequest)(nil),       // 2: service.ProductStatusRequest
	(*ProductStatusReply)(nil),         // 3: service.ProductStatusReply
	(*GetSingleProductRequest)(nil),    // 4: service.GetSingleProductRequest
	(*GetProductsInScopeRequest)(nil),  // 5: service.GetProductsInScopeRequest
	(*StoredProducts)(nil),             // 6: service.StoredProducts
	(*PutSingleProductRequest)(nil),    // 7: service.PutSingleProductRequest
	(*StoredProduct)(nil),              // 8: service.StoredProduct
	(*ClearSingleProductRequest)(nil),  // 9: service.ClearSingleProductRequest
	(*ClearSingleProductResponse)(nil), // 10: service.ClearSingleProductResponse
}
var file_service_messages_proto_depIdxs = []int32{
	8, // 0: service.StoredProducts.products:type_name -> service.StoredProduct
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_messages_proto_init() }
func file_service_messages_proto_init() {
	if File_service_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_service_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusReply); i {
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
		file_service_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductStatusRequest); i {
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
		file_service_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductStatusReply); i {
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
		file_service_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSingleProductRequest); i {
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
		file_service_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductsInScopeRequest); i {
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
		file_service_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredProducts); i {
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
		file_service_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutSingleProductRequest); i {
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
		file_service_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredProduct); i {
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
		file_service_messages_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearSingleProductRequest); i {
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
		file_service_messages_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearSingleProductResponse); i {
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
	file_service_messages_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_messages_proto_goTypes,
		DependencyIndexes: file_service_messages_proto_depIdxs,
		MessageInfos:      file_service_messages_proto_msgTypes,
	}.Build()
	File_service_messages_proto = out.File
	file_service_messages_proto_rawDesc = nil
	file_service_messages_proto_goTypes = nil
	file_service_messages_proto_depIdxs = nil
}
