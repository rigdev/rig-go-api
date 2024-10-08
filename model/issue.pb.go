// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: model/issue.proto

package model

import (
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

type Level int32

const (
	Level_LEVEL_UNSPECIFIED Level = 0
	Level_LEVEL_INFORMATIVE Level = 1
	Level_LEVEL_MINOR       Level = 2
	Level_LEVEL_MAJOR       Level = 3
	Level_LEVEL_CRITICAL    Level = 4
)

// Enum value maps for Level.
var (
	Level_name = map[int32]string{
		0: "LEVEL_UNSPECIFIED",
		1: "LEVEL_INFORMATIVE",
		2: "LEVEL_MINOR",
		3: "LEVEL_MAJOR",
		4: "LEVEL_CRITICAL",
	}
	Level_value = map[string]int32{
		"LEVEL_UNSPECIFIED": 0,
		"LEVEL_INFORMATIVE": 1,
		"LEVEL_MINOR":       2,
		"LEVEL_MAJOR":       3,
		"LEVEL_CRITICAL":    4,
	}
)

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Level) Descriptor() protoreflect.EnumDescriptor {
	return file_model_issue_proto_enumTypes[0].Descriptor()
}

func (Level) Type() protoreflect.EnumType {
	return &file_model_issue_proto_enumTypes[0]
}

func (x Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Level.Descriptor instead.
func (Level) EnumDescriptor() ([]byte, []int) {
	return file_model_issue_proto_rawDescGZIP(), []int{0}
}

type Issue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IssueId   string                 `protobuf:"bytes,1,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	Type      string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	StaleAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=stale_at,json=staleAt,proto3" json:"stale_at,omitempty"`
	ClosedAt  *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=closed_at,json=closedAt,proto3" json:"closed_at,omitempty"`
	Reference *Reference             `protobuf:"bytes,6,opt,name=reference,proto3" json:"reference,omitempty"`
	Message   string                 `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
	Level     Level                  `protobuf:"varint,8,opt,name=level,proto3,enum=model.Level" json:"level,omitempty"`
	Count     uint32                 `protobuf:"varint,9,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Issue) Reset() {
	*x = Issue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_issue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Issue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Issue) ProtoMessage() {}

func (x *Issue) ProtoReflect() protoreflect.Message {
	mi := &file_model_issue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Issue.ProtoReflect.Descriptor instead.
func (*Issue) Descriptor() ([]byte, []int) {
	return file_model_issue_proto_rawDescGZIP(), []int{0}
}

func (x *Issue) GetIssueId() string {
	if x != nil {
		return x.IssueId
	}
	return ""
}

func (x *Issue) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Issue) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Issue) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Issue) GetStaleAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StaleAt
	}
	return nil
}

func (x *Issue) GetClosedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ClosedAt
	}
	return nil
}

func (x *Issue) GetReference() *Reference {
	if x != nil {
		return x.Reference
	}
	return nil
}

func (x *Issue) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Issue) GetLevel() Level {
	if x != nil {
		return x.Level
	}
	return Level_LEVEL_UNSPECIFIED
}

func (x *Issue) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId     string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CapsuleId     string `protobuf:"bytes,2,opt,name=capsule_id,json=capsuleId,proto3" json:"capsule_id,omitempty"`
	EnvironmentId string `protobuf:"bytes,3,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	RolloutId     uint64 `protobuf:"varint,4,opt,name=rollout_id,json=rolloutId,proto3" json:"rollout_id,omitempty"`
	InstanceId    string `protobuf:"bytes,5,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_issue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_model_issue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_model_issue_proto_rawDescGZIP(), []int{1}
}

func (x *Reference) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Reference) GetCapsuleId() string {
	if x != nil {
		return x.CapsuleId
	}
	return ""
}

func (x *Reference) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *Reference) GetRolloutId() uint64 {
	if x != nil {
		return x.RolloutId
	}
	return 0
}

func (x *Reference) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

var File_model_issue_proto protoreflect.FileDescriptor

var file_model_issue_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x03, 0x0a, 0x05,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x74,
	0x61, 0x6c, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x41,
	0x74, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xb0,
	0x01, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x2a, 0x6b, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x52,
	0x4d, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x4d, 0x49, 0x4e, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x5f, 0x4d, 0x41, 0x4a, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x04, 0x42, 0x6f,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67,
	0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x03,
	0x4d, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x05, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0xe2, 0x02, 0x11, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_issue_proto_rawDescOnce sync.Once
	file_model_issue_proto_rawDescData = file_model_issue_proto_rawDesc
)

func file_model_issue_proto_rawDescGZIP() []byte {
	file_model_issue_proto_rawDescOnce.Do(func() {
		file_model_issue_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_issue_proto_rawDescData)
	})
	return file_model_issue_proto_rawDescData
}

var file_model_issue_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_issue_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_issue_proto_goTypes = []any{
	(Level)(0),                    // 0: model.Level
	(*Issue)(nil),                 // 1: model.Issue
	(*Reference)(nil),             // 2: model.Reference
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_model_issue_proto_depIdxs = []int32{
	3, // 0: model.Issue.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: model.Issue.updated_at:type_name -> google.protobuf.Timestamp
	3, // 2: model.Issue.stale_at:type_name -> google.protobuf.Timestamp
	3, // 3: model.Issue.closed_at:type_name -> google.protobuf.Timestamp
	2, // 4: model.Issue.reference:type_name -> model.Reference
	0, // 5: model.Issue.level:type_name -> model.Level
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_model_issue_proto_init() }
func file_model_issue_proto_init() {
	if File_model_issue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_issue_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Issue); i {
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
		file_model_issue_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Reference); i {
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
			RawDescriptor: file_model_issue_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_issue_proto_goTypes,
		DependencyIndexes: file_model_issue_proto_depIdxs,
		EnumInfos:         file_model_issue_proto_enumTypes,
		MessageInfos:      file_model_issue_proto_msgTypes,
	}.Build()
	File_model_issue_proto = out.File
	file_model_issue_proto_rawDesc = nil
	file_model_issue_proto_goTypes = nil
	file_model_issue_proto_depIdxs = nil
}
