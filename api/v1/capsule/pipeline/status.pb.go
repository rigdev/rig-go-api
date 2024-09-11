// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/v1/capsule/pipeline/status.proto

package pipeline

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

type State int32

const (
	// The state is unspecified.
	State_STATE_UNSPECIFIED State = 0
	// The pipeline has started.
	State_STATE_RUNNING State = 1
	// The pipeline is aborted.
	State_STATE_ABORTED State = 2
	// The pipeline is completed.
	State_STATE_COMPLETED State = 3
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_RUNNING",
		2: "STATE_ABORTED",
		3: "STATE_COMPLETED",
	}
	State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STATE_RUNNING":     1,
		"STATE_ABORTED":     2,
		"STATE_COMPLETED":   3,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_capsule_pipeline_status_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_api_v1_capsule_pipeline_status_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_capsule_pipeline_status_proto_rawDescGZIP(), []int{0}
}

type PhaseState int32

const (
	// The state is unspecified.
	PhaseState_PHASE_STATE_UNSPECIFIED PhaseState = 0
	// The phase is not ready for promotion
	PhaseState_PHASE_STATE_NOT_READY PhaseState = 1
	// The phase is ready for promotion
	PhaseState_PHASE_STATE_READY PhaseState = 2
	// The phase is promoted
	PhaseState_PHASE_STATE_PROMOTED PhaseState = 3
)

// Enum value maps for PhaseState.
var (
	PhaseState_name = map[int32]string{
		0: "PHASE_STATE_UNSPECIFIED",
		1: "PHASE_STATE_NOT_READY",
		2: "PHASE_STATE_READY",
		3: "PHASE_STATE_PROMOTED",
	}
	PhaseState_value = map[string]int32{
		"PHASE_STATE_UNSPECIFIED": 0,
		"PHASE_STATE_NOT_READY":   1,
		"PHASE_STATE_READY":       2,
		"PHASE_STATE_PROMOTED":    3,
	}
)

func (x PhaseState) Enum() *PhaseState {
	p := new(PhaseState)
	*p = x
	return p
}

func (x PhaseState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhaseState) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_capsule_pipeline_status_proto_enumTypes[1].Descriptor()
}

func (PhaseState) Type() protoreflect.EnumType {
	return &file_api_v1_capsule_pipeline_status_proto_enumTypes[1]
}

func (x PhaseState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhaseState.Descriptor instead.
func (PhaseState) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_capsule_pipeline_status_proto_rawDescGZIP(), []int{1}
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the pipeline.
	PipelineName string `protobuf:"bytes,1,opt,name=pipeline_name,json=pipelineName,proto3" json:"pipeline_name,omitempty"`
	// The capsule that is executing the pipeline.
	CapsuleId string `protobuf:"bytes,2,opt,name=capsule_id,json=capsuleId,proto3" json:"capsule_id,omitempty"`
	// The ID of the pipeline execution
	ExecutionId uint64 `protobuf:"varint,3,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// The overall state of the pipeline execution.
	State State `protobuf:"varint,4,opt,name=state,proto3,enum=api.v1.capsule.pipeline.State" json:"state,omitempty"`
	// The statuses of the phases in the pipeline.
	PhaseStatuses []*PhaseStatus `protobuf:"bytes,5,rep,name=phase_statuses,json=phaseStatuses,proto3" json:"phase_statuses,omitempty"`
	// When the pipeline was started.
	StartedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// When the pipeline was last updated.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_pipeline_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_pipeline_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_pipeline_status_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetPipelineName() string {
	if x != nil {
		return x.PipelineName
	}
	return ""
}

func (x *Status) GetCapsuleId() string {
	if x != nil {
		return x.CapsuleId
	}
	return ""
}

func (x *Status) GetExecutionId() uint64 {
	if x != nil {
		return x.ExecutionId
	}
	return 0
}

func (x *Status) GetState() State {
	if x != nil {
		return x.State
	}
	return State_STATE_UNSPECIFIED
}

func (x *Status) GetPhaseStatuses() []*PhaseStatus {
	if x != nil {
		return x.PhaseStatuses
	}
	return nil
}

func (x *Status) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Status) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type PhaseStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvironmentId string          `protobuf:"bytes,1,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	State         PhaseState      `protobuf:"varint,2,opt,name=state,proto3,enum=api.v1.capsule.pipeline.PhaseState" json:"state,omitempty"`
	RolloutId     uint64          `protobuf:"varint,3,opt,name=rollout_id,json=rolloutId,proto3" json:"rollout_id,omitempty"`
	Messages      []*PhaseMessage `protobuf:"bytes,4,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *PhaseStatus) Reset() {
	*x = PhaseStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_pipeline_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhaseStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhaseStatus) ProtoMessage() {}

func (x *PhaseStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_pipeline_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhaseStatus.ProtoReflect.Descriptor instead.
func (*PhaseStatus) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_pipeline_status_proto_rawDescGZIP(), []int{1}
}

func (x *PhaseStatus) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *PhaseStatus) GetState() PhaseState {
	if x != nil {
		return x.State
	}
	return PhaseState_PHASE_STATE_UNSPECIFIED
}

func (x *PhaseStatus) GetRolloutId() uint64 {
	if x != nil {
		return x.RolloutId
	}
	return 0
}

func (x *PhaseStatus) GetMessages() []*PhaseMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

type PhaseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PhaseMessage) Reset() {
	*x = PhaseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_pipeline_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhaseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhaseMessage) ProtoMessage() {}

func (x *PhaseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_pipeline_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhaseMessage.ProtoReflect.Descriptor instead.
func (*PhaseMessage) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_pipeline_status_proto_rawDescGZIP(), []int{2}
}

func (x *PhaseMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PhaseMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_api_v1_capsule_pipeline_status_proto protoreflect.FileDescriptor

var file_api_v1_capsule_pipeline_status_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65,
	0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe8, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x34, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75,
	0x6c, 0x65, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x70, 0x68, 0x61, 0x73,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c,
	0x65, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x68, 0x61, 0x73, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x70, 0x68, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd1, 0x01, 0x0a, 0x0b,
	0x50, 0x68, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75,
	0x6c, 0x65, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x68, 0x61, 0x73,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x68, 0x61, 0x73, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22,
	0x62, 0x0a, 0x0c, 0x50, 0x68, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2a, 0x59, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x55, 0x4e,
	0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x41, 0x42, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x75,
	0x0a, 0x0a, 0x50, 0x68, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x17,
	0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x48, 0x41,
	0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x41,
	0x44, 0x59, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x50,
	0x48, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x4f,
	0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0xe0, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c,
	0x65, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0xa2, 0x02, 0x04, 0x41, 0x56, 0x43,
	0x50, 0xaa, 0x02, 0x17, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x73, 0x75,
	0x6c, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0xca, 0x02, 0x17, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5c, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x43,
	0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5c, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x41, 0x70,
	0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x3a, 0x3a,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_capsule_pipeline_status_proto_rawDescOnce sync.Once
	file_api_v1_capsule_pipeline_status_proto_rawDescData = file_api_v1_capsule_pipeline_status_proto_rawDesc
)

func file_api_v1_capsule_pipeline_status_proto_rawDescGZIP() []byte {
	file_api_v1_capsule_pipeline_status_proto_rawDescOnce.Do(func() {
		file_api_v1_capsule_pipeline_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_capsule_pipeline_status_proto_rawDescData)
	})
	return file_api_v1_capsule_pipeline_status_proto_rawDescData
}

var file_api_v1_capsule_pipeline_status_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_v1_capsule_pipeline_status_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_capsule_pipeline_status_proto_goTypes = []interface{}{
	(State)(0),                    // 0: api.v1.capsule.pipeline.State
	(PhaseState)(0),               // 1: api.v1.capsule.pipeline.PhaseState
	(*Status)(nil),                // 2: api.v1.capsule.pipeline.Status
	(*PhaseStatus)(nil),           // 3: api.v1.capsule.pipeline.PhaseStatus
	(*PhaseMessage)(nil),          // 4: api.v1.capsule.pipeline.PhaseMessage
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_api_v1_capsule_pipeline_status_proto_depIdxs = []int32{
	0, // 0: api.v1.capsule.pipeline.Status.state:type_name -> api.v1.capsule.pipeline.State
	3, // 1: api.v1.capsule.pipeline.Status.phase_statuses:type_name -> api.v1.capsule.pipeline.PhaseStatus
	5, // 2: api.v1.capsule.pipeline.Status.started_at:type_name -> google.protobuf.Timestamp
	5, // 3: api.v1.capsule.pipeline.Status.updated_at:type_name -> google.protobuf.Timestamp
	1, // 4: api.v1.capsule.pipeline.PhaseStatus.state:type_name -> api.v1.capsule.pipeline.PhaseState
	4, // 5: api.v1.capsule.pipeline.PhaseStatus.messages:type_name -> api.v1.capsule.pipeline.PhaseMessage
	5, // 6: api.v1.capsule.pipeline.PhaseMessage.timestamp:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_v1_capsule_pipeline_status_proto_init() }
func file_api_v1_capsule_pipeline_status_proto_init() {
	if File_api_v1_capsule_pipeline_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_capsule_pipeline_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_api_v1_capsule_pipeline_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhaseStatus); i {
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
		file_api_v1_capsule_pipeline_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhaseMessage); i {
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
			RawDescriptor: file_api_v1_capsule_pipeline_status_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_capsule_pipeline_status_proto_goTypes,
		DependencyIndexes: file_api_v1_capsule_pipeline_status_proto_depIdxs,
		EnumInfos:         file_api_v1_capsule_pipeline_status_proto_enumTypes,
		MessageInfos:      file_api_v1_capsule_pipeline_status_proto_msgTypes,
	}.Build()
	File_api_v1_capsule_pipeline_status_proto = out.File
	file_api_v1_capsule_pipeline_status_proto_rawDesc = nil
	file_api_v1_capsule_pipeline_status_proto_goTypes = nil
	file_api_v1_capsule_pipeline_status_proto_depIdxs = nil
}
