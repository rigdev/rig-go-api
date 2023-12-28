// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/v1/capsule/build.proto

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

type Build struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId    string                 `protobuf:"bytes,2,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	Digest     string                 `protobuf:"bytes,7,opt,name=digest,proto3" json:"digest,omitempty"`
	Repository string                 `protobuf:"bytes,8,opt,name=repository,proto3" json:"repository,omitempty"`
	Tag        string                 `protobuf:"bytes,9,opt,name=tag,proto3" json:"tag,omitempty"`
	CreatedBy  *model.Author          `protobuf:"bytes,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Origin     *Origin                `protobuf:"bytes,5,opt,name=origin,proto3" json:"origin,omitempty"`
	Labels     map[string]string      `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Build) Reset() {
	*x = Build{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_build_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build) ProtoMessage() {}

func (x *Build) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_build_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Build.ProtoReflect.Descriptor instead.
func (*Build) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_build_proto_rawDescGZIP(), []int{0}
}

func (x *Build) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

func (x *Build) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *Build) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *Build) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Build) GetCreatedBy() *model.Author {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *Build) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Build) GetOrigin() *Origin {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *Build) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type GitReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepositoryUrl string `protobuf:"bytes,1,opt,name=repository_url,json=repositoryUrl,proto3" json:"repository_url,omitempty"`
	CommitSha     string `protobuf:"bytes,2,opt,name=commit_sha,json=commitSha,proto3" json:"commit_sha,omitempty"`
	CommitUrl     string `protobuf:"bytes,3,opt,name=commit_url,json=commitUrl,proto3" json:"commit_url,omitempty"`
}

func (x *GitReference) Reset() {
	*x = GitReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_build_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitReference) ProtoMessage() {}

func (x *GitReference) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_build_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitReference.ProtoReflect.Descriptor instead.
func (*GitReference) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_build_proto_rawDescGZIP(), []int{1}
}

func (x *GitReference) GetRepositoryUrl() string {
	if x != nil {
		return x.RepositoryUrl
	}
	return ""
}

func (x *GitReference) GetCommitSha() string {
	if x != nil {
		return x.CommitSha
	}
	return ""
}

func (x *GitReference) GetCommitUrl() string {
	if x != nil {
		return x.CommitUrl
	}
	return ""
}

type Origin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*Origin_GitReference
	Kind isOrigin_Kind `protobuf_oneof:"kind"`
}

func (x *Origin) Reset() {
	*x = Origin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_build_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Origin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Origin) ProtoMessage() {}

func (x *Origin) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_build_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Origin.ProtoReflect.Descriptor instead.
func (*Origin) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_build_proto_rawDescGZIP(), []int{2}
}

func (m *Origin) GetKind() isOrigin_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Origin) GetGitReference() *GitReference {
	if x, ok := x.GetKind().(*Origin_GitReference); ok {
		return x.GitReference
	}
	return nil
}

type isOrigin_Kind interface {
	isOrigin_Kind()
}

type Origin_GitReference struct {
	GitReference *GitReference `protobuf:"bytes,1,opt,name=git_reference,json=gitReference,proto3,oneof"`
}

func (*Origin_GitReference) isOrigin_Kind() {}

var File_api_v1_capsule_build_proto protoreflect.FileDescriptor

var file_api_v1_capsule_build_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xfb, 0x02, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67,
	0x12, 0x2c, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x39, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x73, 0x0a, 0x0c, 0x47, 0x69, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x5f, 0x73, 0x68, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x53, 0x68, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x55, 0x72, 0x6c, 0x22, 0x55, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x43,
	0x0a, 0x0d, 0x67, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x69, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x67, 0x69, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0xa7, 0x01, 0x0a, 0x12,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75,
	0x6c, 0x65, 0x42, 0x0a, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x67,
	0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0xa2, 0x02, 0x03,
	0x41, 0x56, 0x43, 0xaa, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x61, 0x70,
	0x73, 0x75, 0x6c, 0x65, 0xca, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x61,
	0x70, 0x73, 0x75, 0x6c, 0x65, 0xe2, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x43,
	0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x10, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x61,
	0x70, 0x73, 0x75, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_capsule_build_proto_rawDescOnce sync.Once
	file_api_v1_capsule_build_proto_rawDescData = file_api_v1_capsule_build_proto_rawDesc
)

func file_api_v1_capsule_build_proto_rawDescGZIP() []byte {
	file_api_v1_capsule_build_proto_rawDescOnce.Do(func() {
		file_api_v1_capsule_build_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_capsule_build_proto_rawDescData)
	})
	return file_api_v1_capsule_build_proto_rawDescData
}

var file_api_v1_capsule_build_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_v1_capsule_build_proto_goTypes = []interface{}{
	(*Build)(nil),                 // 0: api.v1.capsule.Build
	(*GitReference)(nil),          // 1: api.v1.capsule.GitReference
	(*Origin)(nil),                // 2: api.v1.capsule.Origin
	nil,                           // 3: api.v1.capsule.Build.LabelsEntry
	(*model.Author)(nil),          // 4: model.Author
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_api_v1_capsule_build_proto_depIdxs = []int32{
	4, // 0: api.v1.capsule.Build.created_by:type_name -> model.Author
	5, // 1: api.v1.capsule.Build.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: api.v1.capsule.Build.origin:type_name -> api.v1.capsule.Origin
	3, // 3: api.v1.capsule.Build.labels:type_name -> api.v1.capsule.Build.LabelsEntry
	1, // 4: api.v1.capsule.Origin.git_reference:type_name -> api.v1.capsule.GitReference
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_v1_capsule_build_proto_init() }
func file_api_v1_capsule_build_proto_init() {
	if File_api_v1_capsule_build_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_capsule_build_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Build); i {
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
		file_api_v1_capsule_build_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitReference); i {
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
		file_api_v1_capsule_build_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Origin); i {
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
	file_api_v1_capsule_build_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Origin_GitReference)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_capsule_build_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_capsule_build_proto_goTypes,
		DependencyIndexes: file_api_v1_capsule_build_proto_depIdxs,
		MessageInfos:      file_api_v1_capsule_build_proto_msgTypes,
	}.Build()
	File_api_v1_capsule_build_proto = out.File
	file_api_v1_capsule_build_proto_rawDesc = nil
	file_api_v1_capsule_build_proto_goTypes = nil
	file_api_v1_capsule_build_proto_depIdxs = nil
}
