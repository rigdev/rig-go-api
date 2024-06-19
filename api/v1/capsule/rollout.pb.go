// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/v1/capsule/rollout.proto

package capsule

import (
	rollout "github.com/rigdev/rig-go-api/api/v1/capsule/rollout"
	_ "github.com/rigdev/rig-go-api/api/v1/environment"
	_ "github.com/rigdev/rig-go-api/api/v1/project"
	model "github.com/rigdev/rig-go-api/model"
	v1 "github.com/rigdev/rig-go-api/platform/v1"
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

type EventType int32

const (
	EventType_EVENT_TYPE_UNSPECIFIED EventType = 0
	EventType_EVENT_TYPE_ABORT       EventType = 1
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "EVENT_TYPE_UNSPECIFIED",
		1: "EVENT_TYPE_ABORT",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_UNSPECIFIED": 0,
		"EVENT_TYPE_ABORT":       1,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_capsule_rollout_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_api_v1_capsule_rollout_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_capsule_rollout_proto_rawDescGZIP(), []int{0}
}

// The rollout model.
type Rollout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique indentifier for the rollout.
	RolloutId uint64 `protobuf:"varint,1,opt,name=rollout_id,json=rolloutId,proto3" json:"rollout_id,omitempty"`
	// The rollout config.
	Config *RolloutConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// The rollout status.
	Status    *rollout.Status `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Spec      *v1.CapsuleSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	Revisions *Revisions      `protobuf:"bytes,6,opt,name=revisions,proto3" json:"revisions,omitempty"`
}

func (x *Rollout) Reset() {
	*x = Rollout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_rollout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rollout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rollout) ProtoMessage() {}

func (x *Rollout) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_rollout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rollout.ProtoReflect.Descriptor instead.
func (*Rollout) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_rollout_proto_rawDescGZIP(), []int{0}
}

func (x *Rollout) GetRolloutId() uint64 {
	if x != nil {
		return x.RolloutId
	}
	return 0
}

func (x *Rollout) GetConfig() *RolloutConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Rollout) GetStatus() *rollout.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Rollout) GetSpec() *v1.CapsuleSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Rollout) GetRevisions() *Revisions {
	if x != nil {
		return x.Revisions
	}
	return nil
}

type RolloutConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user who initiated the rollout.
	CreatedBy                 *model.Author          `protobuf:"bytes,1,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedAt                 *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Changes                   []*Change              `protobuf:"bytes,3,rep,name=changes,proto3" json:"changes,omitempty"`
	Replicas                  uint32                 `protobuf:"varint,4,opt,name=replicas,proto3" json:"replicas,omitempty"`
	ImageId                   string                 `protobuf:"bytes,5,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	Network                   *Network               `protobuf:"bytes,6,opt,name=network,proto3" json:"network,omitempty"`
	ContainerSettings         *ContainerSettings     `protobuf:"bytes,7,opt,name=container_settings,json=containerSettings,proto3" json:"container_settings,omitempty"`
	AutoAddRigServiceAccounts bool                   `protobuf:"varint,8,opt,name=auto_add_rig_service_accounts,json=autoAddRigServiceAccounts,proto3" json:"auto_add_rig_service_accounts,omitempty"`
	ConfigFiles               []*ConfigFile          `protobuf:"bytes,9,rep,name=config_files,json=configFiles,proto3" json:"config_files,omitempty"`
	HorizontalScale           *HorizontalScale       `protobuf:"bytes,10,opt,name=horizontal_scale,json=horizontalScale,proto3" json:"horizontal_scale,omitempty"`
	CronJobs                  []*CronJob             `protobuf:"bytes,11,rep,name=cron_jobs,json=cronJobs,proto3" json:"cron_jobs,omitempty"`
	EnvironmentId             string                 `protobuf:"bytes,12,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	Message                   string                 `protobuf:"bytes,13,opt,name=message,proto3" json:"message,omitempty"`
	Annotations               map[string]string      `protobuf:"bytes,14,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RolloutConfig) Reset() {
	*x = RolloutConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_rollout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RolloutConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RolloutConfig) ProtoMessage() {}

func (x *RolloutConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_rollout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RolloutConfig.ProtoReflect.Descriptor instead.
func (*RolloutConfig) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_rollout_proto_rawDescGZIP(), []int{1}
}

func (x *RolloutConfig) GetCreatedBy() *model.Author {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *RolloutConfig) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RolloutConfig) GetChanges() []*Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

func (x *RolloutConfig) GetReplicas() uint32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *RolloutConfig) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *RolloutConfig) GetNetwork() *Network {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *RolloutConfig) GetContainerSettings() *ContainerSettings {
	if x != nil {
		return x.ContainerSettings
	}
	return nil
}

func (x *RolloutConfig) GetAutoAddRigServiceAccounts() bool {
	if x != nil {
		return x.AutoAddRigServiceAccounts
	}
	return false
}

func (x *RolloutConfig) GetConfigFiles() []*ConfigFile {
	if x != nil {
		return x.ConfigFiles
	}
	return nil
}

func (x *RolloutConfig) GetHorizontalScale() *HorizontalScale {
	if x != nil {
		return x.HorizontalScale
	}
	return nil
}

func (x *RolloutConfig) GetCronJobs() []*CronJob {
	if x != nil {
		return x.CronJobs
	}
	return nil
}

func (x *RolloutConfig) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *RolloutConfig) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RolloutConfig) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type ConfigFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path      string                 `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Content   []byte                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	UpdatedBy *model.Author          `protobuf:"bytes,3,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsSecret  bool                   `protobuf:"varint,5,opt,name=is_secret,json=isSecret,proto3" json:"is_secret,omitempty"`
}

func (x *ConfigFile) Reset() {
	*x = ConfigFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_rollout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigFile) ProtoMessage() {}

func (x *ConfigFile) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_rollout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigFile.ProtoReflect.Descriptor instead.
func (*ConfigFile) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_rollout_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ConfigFile) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ConfigFile) GetUpdatedBy() *model.Author {
	if x != nil {
		return x.UpdatedBy
	}
	return nil
}

func (x *ConfigFile) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ConfigFile) GetIsSecret() bool {
	if x != nil {
		return x.IsSecret
	}
	return false
}

type Revisions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects     []*model.RevisionMetadata `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	Environments []*model.RevisionMetadata `protobuf:"bytes,2,rep,name=environments,proto3" json:"environments,omitempty"`
	CapsuleSets  []*model.RevisionMetadata `protobuf:"bytes,3,rep,name=capsule_sets,json=capsuleSets,proto3" json:"capsule_sets,omitempty"`
	Capsules     []*model.RevisionMetadata `protobuf:"bytes,4,rep,name=capsules,proto3" json:"capsules,omitempty"`
}

func (x *Revisions) Reset() {
	*x = Revisions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_rollout_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Revisions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Revisions) ProtoMessage() {}

func (x *Revisions) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_rollout_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Revisions.ProtoReflect.Descriptor instead.
func (*Revisions) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_rollout_proto_rawDescGZIP(), []int{3}
}

func (x *Revisions) GetProjects() []*model.RevisionMetadata {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *Revisions) GetEnvironments() []*model.RevisionMetadata {
	if x != nil {
		return x.Environments
	}
	return nil
}

func (x *Revisions) GetCapsuleSets() []*model.RevisionMetadata {
	if x != nil {
		return x.CapsuleSets
	}
	return nil
}

func (x *Revisions) GetCapsules() []*model.RevisionMetadata {
	if x != nil {
		return x.Capsules
	}
	return nil
}

type CapsuleProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec     *v1.Capsule             `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Changes  []*Change               `protobuf:"bytes,2,rep,name=changes,proto3" json:"changes,omitempty"`
	Metadata *model.ProposalMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *CapsuleProposal) Reset() {
	*x = CapsuleProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_rollout_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapsuleProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapsuleProposal) ProtoMessage() {}

func (x *CapsuleProposal) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_rollout_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapsuleProposal.ProtoReflect.Descriptor instead.
func (*CapsuleProposal) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_rollout_proto_rawDescGZIP(), []int{4}
}

func (x *CapsuleProposal) GetSpec() *v1.Capsule {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *CapsuleProposal) GetChanges() []*Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

func (x *CapsuleProposal) GetMetadata() *model.ProposalMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_api_v1_capsule_rollout_proto protoreflect.FileDescriptor

var file_api_v1_capsule_rollout_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65,
	0x2f, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x67, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x6c, 0x6f,
	0x75, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2f,
	0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61,
	0x70, 0x73, 0x75, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x07, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x49, 0x64, 0x12,
	0x35, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65,
	0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x73, 0x75,
	0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x37, 0x0a, 0x09,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65,
	0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xbc, 0x06, 0x0a, 0x0d, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2c, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x30, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75,
	0x6c, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x50, 0x0a, 0x12,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x11, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x40,
	0x0a, 0x1d, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x72, 0x69, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x61, 0x75, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x52, 0x69,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12,
	0x4a, 0x0a, 0x10, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x48, 0x6f, 0x72, 0x69, 0x7a,
	0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x0f, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x63,
	0x72, 0x6f, 0x6e, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e,
	0x43, 0x72, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x08, 0x63, 0x72, 0x6f, 0x6e, 0x4a, 0x6f, 0x62,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xc0, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x2c, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0xee, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0c, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x73, 0x75,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x53,
	0x65, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x70,
	0x73, 0x75, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x3d, 0x0a,
	0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x42, 0x4f, 0x52, 0x54, 0x10, 0x01, 0x42, 0xa9, 0x01, 0x0a,
	0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73,
	0x75, 0x6c, 0x65, 0x42, 0x0c, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65,
	0xa2, 0x02, 0x03, 0x41, 0x56, 0x43, 0xaa, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e,
	0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0xca, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31,
	0x5c, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0xe2, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x5c, 0x56,
	0x31, 0x5c, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a,
	0x3a, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_capsule_rollout_proto_rawDescOnce sync.Once
	file_api_v1_capsule_rollout_proto_rawDescData = file_api_v1_capsule_rollout_proto_rawDesc
)

func file_api_v1_capsule_rollout_proto_rawDescGZIP() []byte {
	file_api_v1_capsule_rollout_proto_rawDescOnce.Do(func() {
		file_api_v1_capsule_rollout_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_capsule_rollout_proto_rawDescData)
	})
	return file_api_v1_capsule_rollout_proto_rawDescData
}

var file_api_v1_capsule_rollout_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_capsule_rollout_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_v1_capsule_rollout_proto_goTypes = []interface{}{
	(EventType)(0),                 // 0: api.v1.capsule.EventType
	(*Rollout)(nil),                // 1: api.v1.capsule.Rollout
	(*RolloutConfig)(nil),          // 2: api.v1.capsule.RolloutConfig
	(*ConfigFile)(nil),             // 3: api.v1.capsule.ConfigFile
	(*Revisions)(nil),              // 4: api.v1.capsule.Revisions
	(*CapsuleProposal)(nil),        // 5: api.v1.capsule.CapsuleProposal
	nil,                            // 6: api.v1.capsule.RolloutConfig.AnnotationsEntry
	(*rollout.Status)(nil),         // 7: api.v1.capsule.rollout.Status
	(*v1.CapsuleSpec)(nil),         // 8: platform.v1.CapsuleSpec
	(*model.Author)(nil),           // 9: model.Author
	(*timestamppb.Timestamp)(nil),  // 10: google.protobuf.Timestamp
	(*Change)(nil),                 // 11: api.v1.capsule.Change
	(*Network)(nil),                // 12: api.v1.capsule.Network
	(*ContainerSettings)(nil),      // 13: api.v1.capsule.ContainerSettings
	(*HorizontalScale)(nil),        // 14: api.v1.capsule.HorizontalScale
	(*CronJob)(nil),                // 15: api.v1.capsule.CronJob
	(*model.RevisionMetadata)(nil), // 16: model.RevisionMetadata
	(*v1.Capsule)(nil),             // 17: platform.v1.Capsule
	(*model.ProposalMetadata)(nil), // 18: model.ProposalMetadata
}
var file_api_v1_capsule_rollout_proto_depIdxs = []int32{
	2,  // 0: api.v1.capsule.Rollout.config:type_name -> api.v1.capsule.RolloutConfig
	7,  // 1: api.v1.capsule.Rollout.status:type_name -> api.v1.capsule.rollout.Status
	8,  // 2: api.v1.capsule.Rollout.spec:type_name -> platform.v1.CapsuleSpec
	4,  // 3: api.v1.capsule.Rollout.revisions:type_name -> api.v1.capsule.Revisions
	9,  // 4: api.v1.capsule.RolloutConfig.created_by:type_name -> model.Author
	10, // 5: api.v1.capsule.RolloutConfig.created_at:type_name -> google.protobuf.Timestamp
	11, // 6: api.v1.capsule.RolloutConfig.changes:type_name -> api.v1.capsule.Change
	12, // 7: api.v1.capsule.RolloutConfig.network:type_name -> api.v1.capsule.Network
	13, // 8: api.v1.capsule.RolloutConfig.container_settings:type_name -> api.v1.capsule.ContainerSettings
	3,  // 9: api.v1.capsule.RolloutConfig.config_files:type_name -> api.v1.capsule.ConfigFile
	14, // 10: api.v1.capsule.RolloutConfig.horizontal_scale:type_name -> api.v1.capsule.HorizontalScale
	15, // 11: api.v1.capsule.RolloutConfig.cron_jobs:type_name -> api.v1.capsule.CronJob
	6,  // 12: api.v1.capsule.RolloutConfig.annotations:type_name -> api.v1.capsule.RolloutConfig.AnnotationsEntry
	9,  // 13: api.v1.capsule.ConfigFile.updated_by:type_name -> model.Author
	10, // 14: api.v1.capsule.ConfigFile.updated_at:type_name -> google.protobuf.Timestamp
	16, // 15: api.v1.capsule.Revisions.projects:type_name -> model.RevisionMetadata
	16, // 16: api.v1.capsule.Revisions.environments:type_name -> model.RevisionMetadata
	16, // 17: api.v1.capsule.Revisions.capsule_sets:type_name -> model.RevisionMetadata
	16, // 18: api.v1.capsule.Revisions.capsules:type_name -> model.RevisionMetadata
	17, // 19: api.v1.capsule.CapsuleProposal.spec:type_name -> platform.v1.Capsule
	11, // 20: api.v1.capsule.CapsuleProposal.changes:type_name -> api.v1.capsule.Change
	18, // 21: api.v1.capsule.CapsuleProposal.metadata:type_name -> model.ProposalMetadata
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_api_v1_capsule_rollout_proto_init() }
func file_api_v1_capsule_rollout_proto_init() {
	if File_api_v1_capsule_rollout_proto != nil {
		return
	}
	file_api_v1_capsule_job_proto_init()
	file_api_v1_capsule_change_proto_init()
	file_api_v1_capsule_revision_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_capsule_rollout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rollout); i {
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
		file_api_v1_capsule_rollout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RolloutConfig); i {
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
		file_api_v1_capsule_rollout_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigFile); i {
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
		file_api_v1_capsule_rollout_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Revisions); i {
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
		file_api_v1_capsule_rollout_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapsuleProposal); i {
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
			RawDescriptor: file_api_v1_capsule_rollout_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_capsule_rollout_proto_goTypes,
		DependencyIndexes: file_api_v1_capsule_rollout_proto_depIdxs,
		EnumInfos:         file_api_v1_capsule_rollout_proto_enumTypes,
		MessageInfos:      file_api_v1_capsule_rollout_proto_msgTypes,
	}.Build()
	File_api_v1_capsule_rollout_proto = out.File
	file_api_v1_capsule_rollout_proto_rawDesc = nil
	file_api_v1_capsule_rollout_proto_goTypes = nil
	file_api_v1_capsule_rollout_proto_depIdxs = nil
}
