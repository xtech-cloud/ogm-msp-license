// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/license/key.proto

package license

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

// 生成的请求
type KeyGenerateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Space    string `protobuf:"bytes,1,opt,name=space,proto3" json:"space,omitempty"`        // 空间名
	Count    int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`       // 生成数量
	Capacity int32  `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"` // 可激活的设备数量
	Expiry   int32  `protobuf:"varint,4,opt,name=expiry,proto3" json:"expiry,omitempty"`     // 有效期，以天为单位，默认为0（永久）
	Storage  string `protobuf:"bytes,5,opt,name=storage,proto3" json:"storage,omitempty"`    // 自定义数据
	Profile  string `protobuf:"bytes,6,opt,name=profile,proto3" json:"profile,omitempty"`    // 激活码简介
}

func (x *KeyGenerateRequest) Reset() {
	*x = KeyGenerateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_license_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyGenerateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyGenerateRequest) ProtoMessage() {}

func (x *KeyGenerateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_license_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyGenerateRequest.ProtoReflect.Descriptor instead.
func (*KeyGenerateRequest) Descriptor() ([]byte, []int) {
	return file_proto_license_key_proto_rawDescGZIP(), []int{0}
}

func (x *KeyGenerateRequest) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

func (x *KeyGenerateRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *KeyGenerateRequest) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *KeyGenerateRequest) GetExpiry() int32 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

func (x *KeyGenerateRequest) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

func (x *KeyGenerateRequest) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

// 生成的回复
type KeyGenerateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Number []string `protobuf:"bytes,2,rep,name=number,proto3" json:"number,omitempty"` // 序列号
}

func (x *KeyGenerateResponse) Reset() {
	*x = KeyGenerateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_license_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyGenerateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyGenerateResponse) ProtoMessage() {}

func (x *KeyGenerateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_license_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyGenerateResponse.ProtoReflect.Descriptor instead.
func (*KeyGenerateResponse) Descriptor() ([]byte, []int) {
	return file_proto_license_key_proto_rawDescGZIP(), []int{1}
}

func (x *KeyGenerateResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *KeyGenerateResponse) GetNumber() []string {
	if x != nil {
		return x.Number
	}
	return nil
}

// 查询的请求
type KeyQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"` // 序列号
}

func (x *KeyQueryRequest) Reset() {
	*x = KeyQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_license_key_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyQueryRequest) ProtoMessage() {}

func (x *KeyQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_license_key_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyQueryRequest.ProtoReflect.Descriptor instead.
func (*KeyQueryRequest) Descriptor() ([]byte, []int) {
	return file_proto_license_key_proto_rawDescGZIP(), []int{2}
}

func (x *KeyQueryRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

// 查询的回复
type KeyQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Key    *KeyEntity `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *KeyQueryResponse) Reset() {
	*x = KeyQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_license_key_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyQueryResponse) ProtoMessage() {}

func (x *KeyQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_license_key_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyQueryResponse.ProtoReflect.Descriptor instead.
func (*KeyQueryResponse) Descriptor() ([]byte, []int) {
	return file_proto_license_key_proto_rawDescGZIP(), []int{3}
}

func (x *KeyQueryResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *KeyQueryResponse) GetKey() *KeyEntity {
	if x != nil {
		return x.Key
	}
	return nil
}

// 激活的请求
type KeyActivateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number   string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`     // 序列号
	Consumer string `protobuf:"bytes,2,opt,name=consumer,proto3" json:"consumer,omitempty"` // 消费者
	Space    string `protobuf:"bytes,3,opt,name=space,proto3" json:"space,omitempty"`       //空间
}

func (x *KeyActivateRequest) Reset() {
	*x = KeyActivateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_license_key_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyActivateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyActivateRequest) ProtoMessage() {}

func (x *KeyActivateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_license_key_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyActivateRequest.ProtoReflect.Descriptor instead.
func (*KeyActivateRequest) Descriptor() ([]byte, []int) {
	return file_proto_license_key_proto_rawDescGZIP(), []int{4}
}

func (x *KeyActivateRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *KeyActivateRequest) GetConsumer() string {
	if x != nil {
		return x.Consumer
	}
	return ""
}

func (x *KeyActivateRequest) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

// 激活的回复
type KeyActivateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`         // 状态
	CerUID     string  `protobuf:"bytes,2,opt,name=cerUID,proto3" json:"cerUID,omitempty"`         // 授权文件ID
	CerContent string  `protobuf:"bytes,3,opt,name=cerContent,proto3" json:"cerContent,omitempty"` // 授权文件内容
}

func (x *KeyActivateResponse) Reset() {
	*x = KeyActivateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_license_key_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyActivateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyActivateResponse) ProtoMessage() {}

func (x *KeyActivateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_license_key_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyActivateResponse.ProtoReflect.Descriptor instead.
func (*KeyActivateResponse) Descriptor() ([]byte, []int) {
	return file_proto_license_key_proto_rawDescGZIP(), []int{5}
}

func (x *KeyActivateResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *KeyActivateResponse) GetCerUID() string {
	if x != nil {
		return x.CerUID
	}
	return ""
}

func (x *KeyActivateResponse) GetCerContent() string {
	if x != nil {
		return x.CerContent
	}
	return ""
}

// 挂起的请求
type KeySuspendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Space  string `protobuf:"bytes,1,opt,name=space,proto3" json:"space,omitempty"`   // 空间名
	Number string `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"` // 序列号
	Ban    int32  `protobuf:"varint,3,opt,name=ban,proto3" json:"ban,omitempty"`      // 禁用码, 0表示重新启用
	Reason string `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"` // 原因
}

func (x *KeySuspendRequest) Reset() {
	*x = KeySuspendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_license_key_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeySuspendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeySuspendRequest) ProtoMessage() {}

func (x *KeySuspendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_license_key_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeySuspendRequest.ProtoReflect.Descriptor instead.
func (*KeySuspendRequest) Descriptor() ([]byte, []int) {
	return file_proto_license_key_proto_rawDescGZIP(), []int{6}
}

func (x *KeySuspendRequest) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

func (x *KeySuspendRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *KeySuspendRequest) GetBan() int32 {
	if x != nil {
		return x.Ban
	}
	return 0
}

func (x *KeySuspendRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

// 列举的请求
type KeyListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // 偏移值
	Count  int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`   // 数量
	Space  string `protobuf:"bytes,3,opt,name=space,proto3" json:"space,omitempty"`    // 空间名
}

func (x *KeyListRequest) Reset() {
	*x = KeyListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_license_key_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyListRequest) ProtoMessage() {}

func (x *KeyListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_license_key_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyListRequest.ProtoReflect.Descriptor instead.
func (*KeyListRequest) Descriptor() ([]byte, []int) {
	return file_proto_license_key_proto_rawDescGZIP(), []int{7}
}

func (x *KeyListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *KeyListRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *KeyListRequest) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

// 列举的回复
type KeyListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status      `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Total  int64        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`  //持久化的总数
	Key    []*KeyEntity `protobuf:"bytes,3,rep,name=key,proto3" json:"key,omitempty"`       // 激活码列表
}

func (x *KeyListResponse) Reset() {
	*x = KeyListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_license_key_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyListResponse) ProtoMessage() {}

func (x *KeyListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_license_key_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyListResponse.ProtoReflect.Descriptor instead.
func (*KeyListResponse) Descriptor() ([]byte, []int) {
	return file_proto_license_key_proto_rawDescGZIP(), []int{8}
}

func (x *KeyListResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *KeyListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *KeyListResponse) GetKey() []*KeyEntity {
	if x != nil {
		return x.Key
	}
	return nil
}

var File_proto_license_key_proto protoreflect.FileDescriptor

var file_proto_license_key_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2f,
	0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8,
	0x01, 0x0a, 0x12, 0x4b, 0x65, 0x79, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x56, 0x0a, 0x13, 0x4b, 0x65, 0x79,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x29, 0x0a, 0x0f, 0x4b, 0x65, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x61, 0x0a, 0x10,
	0x4b, 0x65, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x2e, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x5e, 0x0a, 0x12, 0x4b, 0x65, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22,
	0x76, 0x0a, 0x13, 0x4b, 0x65, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x65, 0x72, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x65, 0x72, 0x55, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x6b, 0x0a, 0x11, 0x4b, 0x65, 0x79, 0x53, 0x75,
	0x73, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x61, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x0e, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x76, 0x0a, 0x0f, 0x4b, 0x65,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x32, 0xd5, 0x02, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x47, 0x0a, 0x08, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x2e, 0x4b, 0x65, 0x79, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65,
	0x79, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x18, 0x2e, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x2e, 0x4b, 0x65, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12,
	0x1b, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x07,
	0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x1a, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x2e, 0x4b, 0x65, 0x79, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x6c,
	0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e,
	0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x3b, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_license_key_proto_rawDescOnce sync.Once
	file_proto_license_key_proto_rawDescData = file_proto_license_key_proto_rawDesc
)

func file_proto_license_key_proto_rawDescGZIP() []byte {
	file_proto_license_key_proto_rawDescOnce.Do(func() {
		file_proto_license_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_license_key_proto_rawDescData)
	})
	return file_proto_license_key_proto_rawDescData
}

var file_proto_license_key_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_license_key_proto_goTypes = []interface{}{
	(*KeyGenerateRequest)(nil),  // 0: license.KeyGenerateRequest
	(*KeyGenerateResponse)(nil), // 1: license.KeyGenerateResponse
	(*KeyQueryRequest)(nil),     // 2: license.KeyQueryRequest
	(*KeyQueryResponse)(nil),    // 3: license.KeyQueryResponse
	(*KeyActivateRequest)(nil),  // 4: license.KeyActivateRequest
	(*KeyActivateResponse)(nil), // 5: license.KeyActivateResponse
	(*KeySuspendRequest)(nil),   // 6: license.KeySuspendRequest
	(*KeyListRequest)(nil),      // 7: license.KeyListRequest
	(*KeyListResponse)(nil),     // 8: license.KeyListResponse
	(*Status)(nil),              // 9: license.Status
	(*KeyEntity)(nil),           // 10: license.KeyEntity
	(*BlankResponse)(nil),       // 11: license.BlankResponse
}
var file_proto_license_key_proto_depIdxs = []int32{
	9,  // 0: license.KeyGenerateResponse.status:type_name -> license.Status
	9,  // 1: license.KeyQueryResponse.status:type_name -> license.Status
	10, // 2: license.KeyQueryResponse.key:type_name -> license.KeyEntity
	9,  // 3: license.KeyActivateResponse.status:type_name -> license.Status
	9,  // 4: license.KeyListResponse.status:type_name -> license.Status
	10, // 5: license.KeyListResponse.key:type_name -> license.KeyEntity
	0,  // 6: license.Key.Generate:input_type -> license.KeyGenerateRequest
	2,  // 7: license.Key.Query:input_type -> license.KeyQueryRequest
	4,  // 8: license.Key.Activate:input_type -> license.KeyActivateRequest
	6,  // 9: license.Key.Suspend:input_type -> license.KeySuspendRequest
	7,  // 10: license.Key.List:input_type -> license.KeyListRequest
	1,  // 11: license.Key.Generate:output_type -> license.KeyGenerateResponse
	3,  // 12: license.Key.Query:output_type -> license.KeyQueryResponse
	5,  // 13: license.Key.Activate:output_type -> license.KeyActivateResponse
	11, // 14: license.Key.Suspend:output_type -> license.BlankResponse
	8,  // 15: license.Key.List:output_type -> license.KeyListResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_license_key_proto_init() }
func file_proto_license_key_proto_init() {
	if File_proto_license_key_proto != nil {
		return
	}
	file_proto_license_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_license_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyGenerateRequest); i {
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
		file_proto_license_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyGenerateResponse); i {
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
		file_proto_license_key_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyQueryRequest); i {
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
		file_proto_license_key_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyQueryResponse); i {
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
		file_proto_license_key_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyActivateRequest); i {
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
		file_proto_license_key_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyActivateResponse); i {
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
		file_proto_license_key_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeySuspendRequest); i {
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
		file_proto_license_key_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyListRequest); i {
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
		file_proto_license_key_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyListResponse); i {
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
			RawDescriptor: file_proto_license_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_license_key_proto_goTypes,
		DependencyIndexes: file_proto_license_key_proto_depIdxs,
		MessageInfos:      file_proto_license_key_proto_msgTypes,
	}.Build()
	File_proto_license_key_proto = out.File
	file_proto_license_key_proto_rawDesc = nil
	file_proto_license_key_proto_goTypes = nil
	file_proto_license_key_proto_depIdxs = nil
}
