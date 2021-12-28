// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/license/space.proto

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

// 创建的请求
type SpaceCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // 空间名
}

func (x *SpaceCreateRequest) Reset() {
	*x = SpaceCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_license_space_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceCreateRequest) ProtoMessage() {}

func (x *SpaceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_license_space_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceCreateRequest.ProtoReflect.Descriptor instead.
func (*SpaceCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_license_space_proto_rawDescGZIP(), []int{0}
}

func (x *SpaceCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 搜索的请求
type SpaceSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // 偏移值
	Count  int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`   // 数量
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`      // 空间名
}

func (x *SpaceSearchRequest) Reset() {
	*x = SpaceSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_license_space_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceSearchRequest) ProtoMessage() {}

func (x *SpaceSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_license_space_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceSearchRequest.ProtoReflect.Descriptor instead.
func (*SpaceSearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_license_space_proto_rawDescGZIP(), []int{1}
}

func (x *SpaceSearchRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SpaceSearchRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SpaceSearchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 列举的请求
type SpaceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // 偏移值
	Count  int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`   // 数量
}

func (x *SpaceListRequest) Reset() {
	*x = SpaceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_license_space_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceListRequest) ProtoMessage() {}

func (x *SpaceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_license_space_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceListRequest.ProtoReflect.Descriptor instead.
func (*SpaceListRequest) Descriptor() ([]byte, []int) {
	return file_proto_license_space_proto_rawDescGZIP(), []int{2}
}

func (x *SpaceListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SpaceListRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

// 列举的回复
type SpaceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status        `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Total  int64          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`  // 持久化的总数
	Space  []*SpaceEntity `protobuf:"bytes,3,rep,name=space,proto3" json:"space,omitempty"`   // 空间列表
}

func (x *SpaceListResponse) Reset() {
	*x = SpaceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_license_space_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceListResponse) ProtoMessage() {}

func (x *SpaceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_license_space_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceListResponse.ProtoReflect.Descriptor instead.
func (*SpaceListResponse) Descriptor() ([]byte, []int) {
	return file_proto_license_space_proto_rawDescGZIP(), []int{3}
}

func (x *SpaceListResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *SpaceListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SpaceListResponse) GetSpace() []*SpaceEntity {
	if x != nil {
		return x.Space
	}
	return nil
}

var File_proto_license_space_proto protoreflect.FileDescriptor

var file_proto_license_space_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2f,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x28, 0x0a, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x12, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x40, 0x0a, 0x10, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7e, 0x0a, 0x11, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x32, 0xcd, 0x01, 0x0a, 0x05, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3e,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e,
	0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x3b, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_license_space_proto_rawDescOnce sync.Once
	file_proto_license_space_proto_rawDescData = file_proto_license_space_proto_rawDesc
)

func file_proto_license_space_proto_rawDescGZIP() []byte {
	file_proto_license_space_proto_rawDescOnce.Do(func() {
		file_proto_license_space_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_license_space_proto_rawDescData)
	})
	return file_proto_license_space_proto_rawDescData
}

var file_proto_license_space_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_license_space_proto_goTypes = []interface{}{
	(*SpaceCreateRequest)(nil), // 0: license.SpaceCreateRequest
	(*SpaceSearchRequest)(nil), // 1: license.SpaceSearchRequest
	(*SpaceListRequest)(nil),   // 2: license.SpaceListRequest
	(*SpaceListResponse)(nil),  // 3: license.SpaceListResponse
	(*Status)(nil),             // 4: license.Status
	(*SpaceEntity)(nil),        // 5: license.SpaceEntity
	(*UuidResponse)(nil),       // 6: license.UuidResponse
}
var file_proto_license_space_proto_depIdxs = []int32{
	4, // 0: license.SpaceListResponse.status:type_name -> license.Status
	5, // 1: license.SpaceListResponse.space:type_name -> license.SpaceEntity
	0, // 2: license.Space.Create:input_type -> license.SpaceCreateRequest
	2, // 3: license.Space.List:input_type -> license.SpaceListRequest
	1, // 4: license.Space.Search:input_type -> license.SpaceSearchRequest
	6, // 5: license.Space.Create:output_type -> license.UuidResponse
	3, // 6: license.Space.List:output_type -> license.SpaceListResponse
	3, // 7: license.Space.Search:output_type -> license.SpaceListResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_license_space_proto_init() }
func file_proto_license_space_proto_init() {
	if File_proto_license_space_proto != nil {
		return
	}
	file_proto_license_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_license_space_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceCreateRequest); i {
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
		file_proto_license_space_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceSearchRequest); i {
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
		file_proto_license_space_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceListRequest); i {
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
		file_proto_license_space_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceListResponse); i {
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
			RawDescriptor: file_proto_license_space_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_license_space_proto_goTypes,
		DependencyIndexes: file_proto_license_space_proto_depIdxs,
		MessageInfos:      file_proto_license_space_proto_msgTypes,
	}.Build()
	File_proto_license_space_proto = out.File
	file_proto_license_space_proto_rawDesc = nil
	file_proto_license_space_proto_goTypes = nil
	file_proto_license_space_proto_depIdxs = nil
}
