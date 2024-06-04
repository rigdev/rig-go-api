// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/v1/capsule/capsule.proto

package capsule

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

// Environment wide capsule abstraction.
type Capsule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the capsule.
	CapsuleId string `protobuf:"bytes,1,opt,name=capsule_id,json=capsuleId,proto3" json:"capsule_id,omitempty"`
	// Last time the capsule was updated.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Author of the last update.
	UpdatedBy *model.Author   `protobuf:"bytes,7,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	GitStore  *model.GitStore `protobuf:"bytes,8,opt,name=git_store,json=gitStore,proto3" json:"git_store,omitempty"`
}

func (x *Capsule) Reset() {
	*x = Capsule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_capsule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capsule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capsule) ProtoMessage() {}

func (x *Capsule) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_capsule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capsule.ProtoReflect.Descriptor instead.
func (*Capsule) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_capsule_proto_rawDescGZIP(), []int{0}
}

func (x *Capsule) GetCapsuleId() string {
	if x != nil {
		return x.CapsuleId
	}
	return ""
}

func (x *Capsule) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Capsule) GetUpdatedBy() *model.Author {
	if x != nil {
		return x.UpdatedBy
	}
	return nil
}

func (x *Capsule) GetGitStore() *model.GitStore {
	if x != nil {
		return x.GitStore
	}
	return nil
}

// Legacy update message
type Update struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Update) Reset() {
	*x = Update{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_capsule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Update) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Update) ProtoMessage() {}

func (x *Update) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_capsule_proto_msgTypes[1]
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
	return file_api_v1_capsule_capsule_proto_rawDescGZIP(), []int{1}
}

var File_api_v1_capsule_capsule_proto protoreflect.FileDescriptor

var file_api_v1_capsule_capsule_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65,
	0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x1a, 0x12,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x67, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x07, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x2c, 0x0a, 0x09, 0x67, 0x69, 0x74, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x47, 0x69, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x08, 0x67, 0x69,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x08, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0xa9, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x42, 0x0c, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67,
	0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70,
	0x73, 0x75, 0x6c, 0x65, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x43, 0xaa, 0x02, 0x0e, 0x41, 0x70, 0x69,
	0x2e, 0x56, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0xca, 0x02, 0x0e, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0xe2, 0x02, 0x1a, 0x41,
	0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_capsule_capsule_proto_rawDescOnce sync.Once
	file_api_v1_capsule_capsule_proto_rawDescData = file_api_v1_capsule_capsule_proto_rawDesc
)

func file_api_v1_capsule_capsule_proto_rawDescGZIP() []byte {
	file_api_v1_capsule_capsule_proto_rawDescOnce.Do(func() {
		file_api_v1_capsule_capsule_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_capsule_capsule_proto_rawDescData)
	})
	return file_api_v1_capsule_capsule_proto_rawDescData
}

var file_api_v1_capsule_capsule_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_capsule_capsule_proto_goTypes = []interface{}{
	(*Capsule)(nil),               // 0: api.v1.capsule.Capsule
	(*Update)(nil),                // 1: api.v1.capsule.Update
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*model.Author)(nil),          // 3: model.Author
	(*model.GitStore)(nil),        // 4: model.GitStore
}
var file_api_v1_capsule_capsule_proto_depIdxs = []int32{
	2, // 0: api.v1.capsule.Capsule.updated_at:type_name -> google.protobuf.Timestamp
	3, // 1: api.v1.capsule.Capsule.updated_by:type_name -> model.Author
	4, // 2: api.v1.capsule.Capsule.git_store:type_name -> model.GitStore
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_v1_capsule_capsule_proto_init() }
func file_api_v1_capsule_capsule_proto_init() {
	if File_api_v1_capsule_capsule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_capsule_capsule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capsule); i {
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
		file_api_v1_capsule_capsule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_capsule_capsule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_capsule_capsule_proto_goTypes,
		DependencyIndexes: file_api_v1_capsule_capsule_proto_depIdxs,
		MessageInfos:      file_api_v1_capsule_capsule_proto_msgTypes,
	}.Build()
	File_api_v1_capsule_capsule_proto = out.File
	file_api_v1_capsule_capsule_proto_rawDesc = nil
	file_api_v1_capsule_capsule_proto_goTypes = nil
	file_api_v1_capsule_capsule_proto_depIdxs = nil
}
