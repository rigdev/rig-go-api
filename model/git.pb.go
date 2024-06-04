// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: model/git.proto

package model

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

type GitStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disabled       bool               `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Repository     string             `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	Branch         string             `protobuf:"bytes,3,opt,name=branch,proto3" json:"branch,omitempty"`
	CapsulePath    string             `protobuf:"bytes,4,opt,name=capsule_path,json=capsulePath,proto3" json:"capsule_path,omitempty"`
	CommitTemplate string             `protobuf:"bytes,5,opt,name=commit_template,json=commitTemplate,proto3" json:"commit_template,omitempty"`
	Environments   *EnvironmentFilter `protobuf:"bytes,6,opt,name=environments,proto3" json:"environments,omitempty"`
}

func (x *GitStore) Reset() {
	*x = GitStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_git_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitStore) ProtoMessage() {}

func (x *GitStore) ProtoReflect() protoreflect.Message {
	mi := &file_model_git_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitStore.ProtoReflect.Descriptor instead.
func (*GitStore) Descriptor() ([]byte, []int) {
	return file_model_git_proto_rawDescGZIP(), []int{0}
}

func (x *GitStore) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *GitStore) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *GitStore) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *GitStore) GetCapsulePath() string {
	if x != nil {
		return x.CapsulePath
	}
	return ""
}

func (x *GitStore) GetCommitTemplate() string {
	if x != nil {
		return x.CommitTemplate
	}
	return ""
}

func (x *GitStore) GetEnvironments() *EnvironmentFilter {
	if x != nil {
		return x.Environments
	}
	return nil
}

var File_model_git_proto protoreflect.FileDescriptor

var file_model_git_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x67, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x08, 0x47, 0x69, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x3c,
	0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0c,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x6d, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x08, 0x47, 0x69, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa,
	0x02, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2,
	0x02, 0x11, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_model_git_proto_rawDescOnce sync.Once
	file_model_git_proto_rawDescData = file_model_git_proto_rawDesc
)

func file_model_git_proto_rawDescGZIP() []byte {
	file_model_git_proto_rawDescOnce.Do(func() {
		file_model_git_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_git_proto_rawDescData)
	})
	return file_model_git_proto_rawDescData
}

var file_model_git_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_git_proto_goTypes = []interface{}{
	(*GitStore)(nil),          // 0: model.GitStore
	(*EnvironmentFilter)(nil), // 1: model.EnvironmentFilter
}
var file_model_git_proto_depIdxs = []int32{
	1, // 0: model.GitStore.environments:type_name -> model.EnvironmentFilter
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_model_git_proto_init() }
func file_model_git_proto_init() {
	if File_model_git_proto != nil {
		return
	}
	file_model_environment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_git_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitStore); i {
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
			RawDescriptor: file_model_git_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_git_proto_goTypes,
		DependencyIndexes: file_model_git_proto_depIdxs,
		MessageInfos:      file_model_git_proto_msgTypes,
	}.Build()
	File_model_git_proto = out.File
	file_model_git_proto_rawDesc = nil
	file_model_git_proto_goTypes = nil
	file_model_git_proto_depIdxs = nil
}
