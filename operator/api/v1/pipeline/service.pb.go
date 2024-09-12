// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: operator/api/v1/pipeline/service.proto

package pipeline

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

type ObjectOutcome int32

const (
	ObjectOutcome_OBJECT_OUTCOME_UNSPECIFIED    ObjectOutcome = 0
	ObjectOutcome_OBJECT_OUTCOME_CREATE         ObjectOutcome = 1
	ObjectOutcome_OBJECT_OUTCOME_UPDATE         ObjectOutcome = 2
	ObjectOutcome_OBJECT_OUTCOME_DELETE         ObjectOutcome = 3
	ObjectOutcome_OBJECT_OUTCOME_UNCHANGED      ObjectOutcome = 4
	ObjectOutcome_OBJECT_OUTCOME_ALREADY_EXISTS ObjectOutcome = 5
)

// Enum value maps for ObjectOutcome.
var (
	ObjectOutcome_name = map[int32]string{
		0: "OBJECT_OUTCOME_UNSPECIFIED",
		1: "OBJECT_OUTCOME_CREATE",
		2: "OBJECT_OUTCOME_UPDATE",
		3: "OBJECT_OUTCOME_DELETE",
		4: "OBJECT_OUTCOME_UNCHANGED",
		5: "OBJECT_OUTCOME_ALREADY_EXISTS",
	}
	ObjectOutcome_value = map[string]int32{
		"OBJECT_OUTCOME_UNSPECIFIED":    0,
		"OBJECT_OUTCOME_CREATE":         1,
		"OBJECT_OUTCOME_UPDATE":         2,
		"OBJECT_OUTCOME_DELETE":         3,
		"OBJECT_OUTCOME_UNCHANGED":      4,
		"OBJECT_OUTCOME_ALREADY_EXISTS": 5,
	}
)

func (x ObjectOutcome) Enum() *ObjectOutcome {
	p := new(ObjectOutcome)
	*p = x
	return p
}

func (x ObjectOutcome) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectOutcome) Descriptor() protoreflect.EnumDescriptor {
	return file_operator_api_v1_pipeline_service_proto_enumTypes[0].Descriptor()
}

func (ObjectOutcome) Type() protoreflect.EnumType {
	return &file_operator_api_v1_pipeline_service_proto_enumTypes[0]
}

func (x ObjectOutcome) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectOutcome.Descriptor instead.
func (ObjectOutcome) EnumDescriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{0}
}

type WatchObjectStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *WatchObjectStatusRequest) Reset() {
	*x = WatchObjectStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchObjectStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchObjectStatusRequest) ProtoMessage() {}

func (x *WatchObjectStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchObjectStatusRequest.ProtoReflect.Descriptor instead.
func (*WatchObjectStatusRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{0}
}

func (x *WatchObjectStatusRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type WatchObjectStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Change *ObjectStatusChange `protobuf:"bytes,1,opt,name=change,proto3" json:"change,omitempty"`
}

func (x *WatchObjectStatusResponse) Reset() {
	*x = WatchObjectStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchObjectStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchObjectStatusResponse) ProtoMessage() {}

func (x *WatchObjectStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchObjectStatusResponse.ProtoReflect.Descriptor instead.
func (*WatchObjectStatusResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{1}
}

func (x *WatchObjectStatusResponse) GetChange() *ObjectStatusChange {
	if x != nil {
		return x.Change
	}
	return nil
}

type DryRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Capsule   string `protobuf:"bytes,2,opt,name=capsule,proto3" json:"capsule,omitempty"`
	// A YAML encoded OperatorConfig, to be used when processing the pipeline.
	OperatorConfig string `protobuf:"bytes,3,opt,name=operator_config,json=operatorConfig,proto3" json:"operator_config,omitempty"`
	// An optional YAML encoded capsule spec, to be used instead of the current
	// one.
	CapsuleSpec string `protobuf:"bytes,4,opt,name=capsule_spec,json=capsuleSpec,proto3" json:"capsule_spec,omitempty"`
	// If force is enabled, existing resources will be handled as if they were
	// supposed to be replaced.
	Force bool `protobuf:"varint,5,opt,name=force,proto3" json:"force,omitempty"`
	// Additional objects to be considered materialized when performing the dryrun
	AdditionalObjects []*Object `protobuf:"bytes,6,rep,name=additional_objects,json=additionalObjects,proto3" json:"additional_objects,omitempty"`
}

func (x *DryRunRequest) Reset() {
	*x = DryRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DryRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DryRunRequest) ProtoMessage() {}

func (x *DryRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DryRunRequest.ProtoReflect.Descriptor instead.
func (*DryRunRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{2}
}

func (x *DryRunRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DryRunRequest) GetCapsule() string {
	if x != nil {
		return x.Capsule
	}
	return ""
}

func (x *DryRunRequest) GetOperatorConfig() string {
	if x != nil {
		return x.OperatorConfig
	}
	return ""
}

func (x *DryRunRequest) GetCapsuleSpec() string {
	if x != nil {
		return x.CapsuleSpec
	}
	return ""
}

func (x *DryRunRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

func (x *DryRunRequest) GetAdditionalObjects() []*Object {
	if x != nil {
		return x.AdditionalObjects
	}
	return nil
}

type DryRunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputObjects  []*Object       `protobuf:"bytes,1,rep,name=input_objects,json=inputObjects,proto3" json:"input_objects,omitempty"`
	OutputObjects []*ObjectChange `protobuf:"bytes,2,rep,name=output_objects,json=outputObjects,proto3" json:"output_objects,omitempty"`
}

func (x *DryRunResponse) Reset() {
	*x = DryRunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DryRunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DryRunResponse) ProtoMessage() {}

func (x *DryRunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DryRunResponse.ProtoReflect.Descriptor instead.
func (*DryRunResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{3}
}

func (x *DryRunResponse) GetInputObjects() []*Object {
	if x != nil {
		return x.InputObjects
	}
	return nil
}

func (x *DryRunResponse) GetOutputObjects() []*ObjectChange {
	if x != nil {
		return x.OutputObjects
	}
	return nil
}

type ObjectChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object  *Object       `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Outcome ObjectOutcome `protobuf:"varint,2,opt,name=outcome,proto3,enum=api.v1.pipeline.ObjectOutcome" json:"outcome,omitempty"`
}

func (x *ObjectChange) Reset() {
	*x = ObjectChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectChange) ProtoMessage() {}

func (x *ObjectChange) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectChange.ProtoReflect.Descriptor instead.
func (*ObjectChange) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{4}
}

func (x *ObjectChange) GetObject() *Object {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *ObjectChange) GetOutcome() ObjectOutcome {
	if x != nil {
		return x.Outcome
	}
	return ObjectOutcome_OBJECT_OUTCOME_UNSPECIFIED
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gvk     *GVK   `protobuf:"bytes,1,opt,name=gvk,proto3" json:"gvk,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{5}
}

func (x *Object) GetGvk() *GVK {
	if x != nil {
		return x.Gvk
	}
	return nil
}

func (x *Object) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Object) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type DryRunPluginConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace      string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Capsule        string `protobuf:"bytes,2,opt,name=capsule,proto3" json:"capsule,omitempty"`
	OperatorConfig string `protobuf:"bytes,3,opt,name=operator_config,json=operatorConfig,proto3" json:"operator_config,omitempty"`
	CapsuleSpec    string `protobuf:"bytes,4,opt,name=capsule_spec,json=capsuleSpec,proto3" json:"capsule_spec,omitempty"`
}

func (x *DryRunPluginConfigRequest) Reset() {
	*x = DryRunPluginConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DryRunPluginConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DryRunPluginConfigRequest) ProtoMessage() {}

func (x *DryRunPluginConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DryRunPluginConfigRequest.ProtoReflect.Descriptor instead.
func (*DryRunPluginConfigRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{6}
}

func (x *DryRunPluginConfigRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DryRunPluginConfigRequest) GetCapsule() string {
	if x != nil {
		return x.Capsule
	}
	return ""
}

func (x *DryRunPluginConfigRequest) GetOperatorConfig() string {
	if x != nil {
		return x.OperatorConfig
	}
	return ""
}

func (x *DryRunPluginConfigRequest) GetCapsuleSpec() string {
	if x != nil {
		return x.CapsuleSpec
	}
	return ""
}

type DryRunPluginConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Steps []*StepConfig `protobuf:"bytes,1,rep,name=steps,proto3" json:"steps,omitempty"`
}

func (x *DryRunPluginConfigResponse) Reset() {
	*x = DryRunPluginConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DryRunPluginConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DryRunPluginConfigResponse) ProtoMessage() {}

func (x *DryRunPluginConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DryRunPluginConfigResponse.ProtoReflect.Descriptor instead.
func (*DryRunPluginConfigResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{7}
}

func (x *DryRunPluginConfigResponse) GetSteps() []*StepConfig {
	if x != nil {
		return x.Steps
	}
	return nil
}

type StepConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Plugins []*PluginConfig `protobuf:"bytes,2,rep,name=plugins,proto3" json:"plugins,omitempty"`
}

func (x *StepConfig) Reset() {
	*x = StepConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepConfig) ProtoMessage() {}

func (x *StepConfig) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepConfig.ProtoReflect.Descriptor instead.
func (*StepConfig) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{8}
}

func (x *StepConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StepConfig) GetPlugins() []*PluginConfig {
	if x != nil {
		return x.Plugins
	}
	return nil
}

type PluginConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Config string `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Err    string `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *PluginConfig) Reset() {
	*x = PluginConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginConfig) ProtoMessage() {}

func (x *PluginConfig) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginConfig.ProtoReflect.Descriptor instead.
func (*PluginConfig) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{9}
}

func (x *PluginConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginConfig) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *PluginConfig) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_operator_api_v1_pipeline_service_proto protoreflect.FileDescriptor

var file_operator_api_v1_pipeline_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x18, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x22, 0x58, 0x0a, 0x19, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0xf1, 0x01, 0x0a, 0x0d,
	0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61,
	0x70, 0x73, 0x75, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x11, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22,
	0x94, 0x01, 0x0a, 0x0e, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x12, 0x44, 0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x79, 0x0a, 0x0c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d,
	0x65, 0x22, 0x5e, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x26, 0x0a, 0x03, 0x67,
	0x76, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x56, 0x4b, 0x52, 0x03,
	0x67, 0x76, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x9f, 0x01, 0x0a, 0x19, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x22, 0x4f, 0x0a, 0x1a, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x73,
	0x74, 0x65, 0x70, 0x73, 0x22, 0x59, 0x0a, 0x0a, 0x53, 0x74, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x22,
	0x4c, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x2a, 0xc1, 0x01,
	0x0a, 0x0d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x1a, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4f, 0x55, 0x54, 0x43, 0x4f, 0x4d,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x19, 0x0a, 0x15, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4f, 0x55, 0x54, 0x43, 0x4f, 0x4d,
	0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x42,
	0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4f, 0x55, 0x54, 0x43, 0x4f, 0x4d, 0x45, 0x5f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f,
	0x4f, 0x55, 0x54, 0x43, 0x4f, 0x4d, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03,
	0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4f, 0x55, 0x54, 0x43, 0x4f,
	0x4d, 0x45, 0x5f, 0x55, 0x4e, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x04, 0x12, 0x21,
	0x0a, 0x1d, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4f, 0x55, 0x54, 0x43, 0x4f, 0x4d, 0x45,
	0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10,
	0x05, 0x32, 0xb7, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a,
	0x11, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4b, 0x0a,
	0x06, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x12, 0x44, 0x72,
	0x79, 0x52, 0x75, 0x6e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2e, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x44,
	0x72, 0x79, 0x52, 0x75, 0x6e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb8, 0x01, 0x0a, 0x13,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x50,
	0xaa, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0xca, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0xe2, 0x02, 0x1b, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x11, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_operator_api_v1_pipeline_service_proto_rawDescOnce sync.Once
	file_operator_api_v1_pipeline_service_proto_rawDescData = file_operator_api_v1_pipeline_service_proto_rawDesc
)

func file_operator_api_v1_pipeline_service_proto_rawDescGZIP() []byte {
	file_operator_api_v1_pipeline_service_proto_rawDescOnce.Do(func() {
		file_operator_api_v1_pipeline_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_operator_api_v1_pipeline_service_proto_rawDescData)
	})
	return file_operator_api_v1_pipeline_service_proto_rawDescData
}

var file_operator_api_v1_pipeline_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_operator_api_v1_pipeline_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_operator_api_v1_pipeline_service_proto_goTypes = []interface{}{
	(ObjectOutcome)(0),                 // 0: api.v1.pipeline.ObjectOutcome
	(*WatchObjectStatusRequest)(nil),   // 1: api.v1.pipeline.WatchObjectStatusRequest
	(*WatchObjectStatusResponse)(nil),  // 2: api.v1.pipeline.WatchObjectStatusResponse
	(*DryRunRequest)(nil),              // 3: api.v1.pipeline.DryRunRequest
	(*DryRunResponse)(nil),             // 4: api.v1.pipeline.DryRunResponse
	(*ObjectChange)(nil),               // 5: api.v1.pipeline.ObjectChange
	(*Object)(nil),                     // 6: api.v1.pipeline.Object
	(*DryRunPluginConfigRequest)(nil),  // 7: api.v1.pipeline.DryRunPluginConfigRequest
	(*DryRunPluginConfigResponse)(nil), // 8: api.v1.pipeline.DryRunPluginConfigResponse
	(*StepConfig)(nil),                 // 9: api.v1.pipeline.StepConfig
	(*PluginConfig)(nil),               // 10: api.v1.pipeline.PluginConfig
	(*ObjectStatusChange)(nil),         // 11: api.v1.pipeline.ObjectStatusChange
	(*GVK)(nil),                        // 12: api.v1.pipeline.GVK
}
var file_operator_api_v1_pipeline_service_proto_depIdxs = []int32{
	11, // 0: api.v1.pipeline.WatchObjectStatusResponse.change:type_name -> api.v1.pipeline.ObjectStatusChange
	6,  // 1: api.v1.pipeline.DryRunRequest.additional_objects:type_name -> api.v1.pipeline.Object
	6,  // 2: api.v1.pipeline.DryRunResponse.input_objects:type_name -> api.v1.pipeline.Object
	5,  // 3: api.v1.pipeline.DryRunResponse.output_objects:type_name -> api.v1.pipeline.ObjectChange
	6,  // 4: api.v1.pipeline.ObjectChange.object:type_name -> api.v1.pipeline.Object
	0,  // 5: api.v1.pipeline.ObjectChange.outcome:type_name -> api.v1.pipeline.ObjectOutcome
	12, // 6: api.v1.pipeline.Object.gvk:type_name -> api.v1.pipeline.GVK
	9,  // 7: api.v1.pipeline.DryRunPluginConfigResponse.steps:type_name -> api.v1.pipeline.StepConfig
	10, // 8: api.v1.pipeline.StepConfig.plugins:type_name -> api.v1.pipeline.PluginConfig
	1,  // 9: api.v1.pipeline.Service.WatchObjectStatus:input_type -> api.v1.pipeline.WatchObjectStatusRequest
	3,  // 10: api.v1.pipeline.Service.DryRun:input_type -> api.v1.pipeline.DryRunRequest
	7,  // 11: api.v1.pipeline.Service.DryRunPluginConfig:input_type -> api.v1.pipeline.DryRunPluginConfigRequest
	2,  // 12: api.v1.pipeline.Service.WatchObjectStatus:output_type -> api.v1.pipeline.WatchObjectStatusResponse
	4,  // 13: api.v1.pipeline.Service.DryRun:output_type -> api.v1.pipeline.DryRunResponse
	8,  // 14: api.v1.pipeline.Service.DryRunPluginConfig:output_type -> api.v1.pipeline.DryRunPluginConfigResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_operator_api_v1_pipeline_service_proto_init() }
func file_operator_api_v1_pipeline_service_proto_init() {
	if File_operator_api_v1_pipeline_service_proto != nil {
		return
	}
	file_operator_api_v1_pipeline_object_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_operator_api_v1_pipeline_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchObjectStatusRequest); i {
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
		file_operator_api_v1_pipeline_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchObjectStatusResponse); i {
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
		file_operator_api_v1_pipeline_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DryRunRequest); i {
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
		file_operator_api_v1_pipeline_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DryRunResponse); i {
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
		file_operator_api_v1_pipeline_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectChange); i {
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
		file_operator_api_v1_pipeline_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_operator_api_v1_pipeline_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DryRunPluginConfigRequest); i {
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
		file_operator_api_v1_pipeline_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DryRunPluginConfigResponse); i {
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
		file_operator_api_v1_pipeline_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepConfig); i {
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
		file_operator_api_v1_pipeline_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginConfig); i {
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
			RawDescriptor: file_operator_api_v1_pipeline_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_operator_api_v1_pipeline_service_proto_goTypes,
		DependencyIndexes: file_operator_api_v1_pipeline_service_proto_depIdxs,
		EnumInfos:         file_operator_api_v1_pipeline_service_proto_enumTypes,
		MessageInfos:      file_operator_api_v1_pipeline_service_proto_msgTypes,
	}.Build()
	File_operator_api_v1_pipeline_service_proto = out.File
	file_operator_api_v1_pipeline_service_proto_rawDesc = nil
	file_operator_api_v1_pipeline_service_proto_goTypes = nil
	file_operator_api_v1_pipeline_service_proto_depIdxs = nil
}
