// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: common.proto

package myapi

import (
	_ "github.com/golang/protobuf/ptypes/struct"
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

//Request mapped to URL path by option (google.api.http)
type BulkFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BulkFields) Reset() {
	*x = BulkFields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkFields) ProtoMessage() {}

func (x *BulkFields) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkFields.ProtoReflect.Descriptor instead.
func (*BulkFields) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type GetFieldParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//field ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFieldParams) Reset() {
	*x = GetFieldParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFieldParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFieldParams) ProtoMessage() {}

func (x *GetFieldParams) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFieldParams.ProtoReflect.Descriptor instead.
func (*GetFieldParams) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *GetFieldParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AddFieldParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemName string `protobuf:"bytes,2,opt,name=systemName,json=fieldSystemName,proto3" json:"systemName,omitempty"`
}

func (x *AddFieldParams) Reset() {
	*x = AddFieldParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFieldParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFieldParams) ProtoMessage() {}

func (x *AddFieldParams) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFieldParams.ProtoReflect.Descriptor instead.
func (*AddFieldParams) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *AddFieldParams) GetSystemName() string {
	if x != nil {
		return x.SystemName
	}
	return ""
}

//Response
type FieldObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SystemName string `protobuf:"bytes,2,opt,name=systemName,proto3" json:"systemName,omitempty"`
}

func (x *FieldObject) Reset() {
	*x = FieldObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldObject) ProtoMessage() {}

func (x *FieldObject) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldObject.ProtoReflect.Descriptor instead.
func (*FieldObject) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *FieldObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FieldObject) GetSystemName() string {
	if x != nil {
		return x.SystemName
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x42, 0x75,
	0x6c, 0x6b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x0a,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x3d, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x32, 0xe5, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x46, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0f,
	0x2e, 0x41, 0x64, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x0c, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x64, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x44, 0x0a, 0x0d, 0x42, 0x75,
	0x6c, 0x6b, 0x47, 0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x0b, 0x2e, 0x42, 0x75,
	0x6c, 0x6b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x0c, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x30, 0x01,
	0x12, 0x46, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0f, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0c, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x1b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x24, 0x5a, 0x15, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x6d, 0x79, 0x61, 0x70,
	0x69, 0x92, 0x41, 0x0a, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_proto_goTypes = []interface{}{
	(*BulkFields)(nil),     // 0: BulkFields
	(*GetFieldParams)(nil), // 1: GetFieldParams
	(*AddFieldParams)(nil), // 2: AddFieldParams
	(*FieldObject)(nil),    // 3: FieldObject
}
var file_common_proto_depIdxs = []int32{
	2, // 0: CommonService.AddField:input_type -> AddFieldParams
	0, // 1: CommonService.BulkGetFields:input_type -> BulkFields
	1, // 2: CommonService.GetField:input_type -> GetFieldParams
	3, // 3: CommonService.AddField:output_type -> FieldObject
	3, // 4: CommonService.BulkGetFields:output_type -> FieldObject
	3, // 5: CommonService.GetField:output_type -> FieldObject
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkFields); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFieldParams); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFieldParams); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldObject); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
