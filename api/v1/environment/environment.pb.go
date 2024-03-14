// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/v1/environment/environment.proto

package environment

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

type Update struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Field:
	//
	//	*Update_Default
	Field isUpdate_Field `protobuf_oneof:"field"`
}

func (x *Update) Reset() {
	*x = Update{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_environment_environment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Update) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Update) ProtoMessage() {}

func (x *Update) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_environment_environment_proto_msgTypes[0]
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
	return file_api_v1_environment_environment_proto_rawDescGZIP(), []int{0}
}

func (m *Update) GetField() isUpdate_Field {
	if m != nil {
		return m.Field
	}
	return nil
}

func (x *Update) GetDefault() bool {
	if x, ok := x.GetField().(*Update_Default); ok {
		return x.Default
	}
	return false
}

type isUpdate_Field interface {
	isUpdate_Field()
}

type Update_Default struct {
	Default bool `protobuf:"varint,1,opt,name=default,proto3,oneof"`
}

func (*Update_Default) isUpdate_Field() {}

// Environment model.
type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the environment.
	EnvironmentId string `protobuf:"bytes,1,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	// If true, this is the default environment.
	Default bool `protobuf:"varint,2,opt,name=default,proto3" json:"default,omitempty"`
	// The version of the Rig-operator CRD for this environment.
	OperatorVersion string `protobuf:"bytes,3,opt,name=operator_version,json=operatorVersion,proto3" json:"operator_version,omitempty"`
	// ID of the backing cluster.
	ClusterId string `protobuf:"bytes,4,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_environment_environment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_environment_environment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_api_v1_environment_environment_proto_rawDescGZIP(), []int{1}
}

func (x *Environment) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *Environment) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

func (x *Environment) GetOperatorVersion() string {
	if x != nil {
		return x.OperatorVersion
	}
	return ""
}

func (x *Environment) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

var File_api_v1_environment_environment_proto protoreflect.FileDescriptor

var file_api_v1_environment_environment_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x42, 0x07, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x98, 0x01, 0x0a, 0x0b, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x42, 0xc5, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x42,
	0x10, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x45, 0xaa, 0x02, 0x12, 0x41, 0x70, 0x69,
	0x2e, 0x56, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0xca,
	0x02, 0x12, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0xe2, 0x02, 0x1e, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a,
	0x3a, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_environment_environment_proto_rawDescOnce sync.Once
	file_api_v1_environment_environment_proto_rawDescData = file_api_v1_environment_environment_proto_rawDesc
)

func file_api_v1_environment_environment_proto_rawDescGZIP() []byte {
	file_api_v1_environment_environment_proto_rawDescOnce.Do(func() {
		file_api_v1_environment_environment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_environment_environment_proto_rawDescData)
	})
	return file_api_v1_environment_environment_proto_rawDescData
}

var file_api_v1_environment_environment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_environment_environment_proto_goTypes = []interface{}{
	(*Update)(nil),      // 0: api.v1.environment.Update
	(*Environment)(nil), // 1: api.v1.environment.Environment
}
var file_api_v1_environment_environment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_environment_environment_proto_init() }
func file_api_v1_environment_environment_proto_init() {
	if File_api_v1_environment_environment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_environment_environment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_v1_environment_environment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Environment); i {
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
	file_api_v1_environment_environment_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Update_Default)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_environment_environment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_environment_environment_proto_goTypes,
		DependencyIndexes: file_api_v1_environment_environment_proto_depIdxs,
		MessageInfos:      file_api_v1_environment_environment_proto_msgTypes,
	}.Build()
	File_api_v1_environment_environment_proto = out.File
	file_api_v1_environment_environment_proto_rawDesc = nil
	file_api_v1_environment_environment_proto_goTypes = nil
	file_api_v1_environment_environment_proto_depIdxs = nil
}
