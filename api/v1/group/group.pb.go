// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v1/group/group.proto

package group

import (
	model "github.com/rigdev/rig-go-api/model"
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

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId    string                 `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Name       string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NumMembers uint64                 `protobuf:"varint,5,opt,name=num_members,json=numMembers,proto3" json:"num_members,omitempty"`
	Metadata   map[string][]byte      `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_api_v1_group_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetNumMembers() uint64 {
	if x != nil {
		return x.NumMembers
	}
	return 0
}

func (x *Group) GetMetadata() map[string][]byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Group) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Group) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Update struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Field:
	//
	//	*Update_Name
	//	*Update_SetMetadata
	//	*Update_DeleteMetadataKey
	Field isUpdate_Field `protobuf_oneof:"field"`
}

func (x *Update) Reset() {
	*x = Update{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Update) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Update) ProtoMessage() {}

func (x *Update) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Update.ProtoReflect.Descriptor instead.
func (*Update) Descriptor() ([]byte, []int) {
	return file_api_v1_group_group_proto_rawDescGZIP(), []int{1}
}

func (m *Update) GetField() isUpdate_Field {
	if m != nil {
		return m.Field
	}
	return nil
}

func (x *Update) GetName() string {
	if x, ok := x.GetField().(*Update_Name); ok {
		return x.Name
	}
	return ""
}

func (x *Update) GetSetMetadata() *model.Metadata {
	if x, ok := x.GetField().(*Update_SetMetadata); ok {
		return x.SetMetadata
	}
	return nil
}

func (x *Update) GetDeleteMetadataKey() string {
	if x, ok := x.GetField().(*Update_DeleteMetadataKey); ok {
		return x.DeleteMetadataKey
	}
	return ""
}

type isUpdate_Field interface {
	isUpdate_Field()
}

type Update_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

type Update_SetMetadata struct {
	SetMetadata *model.Metadata `protobuf:"bytes,2,opt,name=set_metadata,json=setMetadata,proto3,oneof"`
}

type Update_DeleteMetadataKey struct {
	DeleteMetadataKey string `protobuf:"bytes,3,opt,name=delete_metadata_key,json=deleteMetadataKey,proto3,oneof"`
}

func (*Update_Name) isUpdate_Field() {}

func (*Update_SetMetadata) isUpdate_Field() {}

func (*Update_DeleteMetadataKey) isUpdate_Field() {}

var File_api_v1_group_group_proto protoreflect.FileDescriptor

var file_api_v1_group_group_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc9, 0x02, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6e,
	0x75, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x3b,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8f, 0x01, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x0c,
	0x73, 0x65, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x30, 0x0a, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x4b, 0x65, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x9b, 0x01,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x42, 0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x67,
	0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0xa2, 0x02, 0x03, 0x41, 0x56,
	0x47, 0xaa, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0xca, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0xe2,
	0x02, 0x18, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_v1_group_group_proto_rawDescOnce sync.Once
	file_api_v1_group_group_proto_rawDescData = file_api_v1_group_group_proto_rawDesc
)

func file_api_v1_group_group_proto_rawDescGZIP() []byte {
	file_api_v1_group_group_proto_rawDescOnce.Do(func() {
		file_api_v1_group_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_group_group_proto_rawDescData)
	})
	return file_api_v1_group_group_proto_rawDescData
}

var file_api_v1_group_group_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_group_group_proto_goTypes = []interface{}{
	(*Group)(nil),                 // 0: api.v1.group.Group
	(*Update)(nil),                // 1: api.v1.group.Update
	nil,                           // 2: api.v1.group.Group.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*model.Metadata)(nil),        // 4: model.Metadata
}
var file_api_v1_group_group_proto_depIdxs = []int32{
	2, // 0: api.v1.group.Group.metadata:type_name -> api.v1.group.Group.MetadataEntry
	3, // 1: api.v1.group.Group.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: api.v1.group.Group.updated_at:type_name -> google.protobuf.Timestamp
	4, // 3: api.v1.group.Update.set_metadata:type_name -> model.Metadata
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1_group_group_proto_init() }
func file_api_v1_group_group_proto_init() {
	if File_api_v1_group_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_group_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_api_v1_group_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Update); i {
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
	file_api_v1_group_group_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Update_Name)(nil),
		(*Update_SetMetadata)(nil),
		(*Update_DeleteMetadataKey)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_group_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_group_group_proto_goTypes,
		DependencyIndexes: file_api_v1_group_group_proto_depIdxs,
		MessageInfos:      file_api_v1_group_group_proto_msgTypes,
	}.Build()
	File_api_v1_group_group_proto = out.File
	file_api_v1_group_group_proto_rawDesc = nil
	file_api_v1_group_group_proto_goTypes = nil
	file_api_v1_group_group_proto_depIdxs = nil
}
