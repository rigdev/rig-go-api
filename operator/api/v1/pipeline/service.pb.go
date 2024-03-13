// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
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

type ObjectState int32

const (
	ObjectState_OBJECT_STATE_UNSPECIFIED    ObjectState = 0
	ObjectState_OBJECT_STATE_CREATE         ObjectState = 1
	ObjectState_OBJECT_STATE_UPDATE         ObjectState = 2
	ObjectState_OBJECT_STATE_DELETE         ObjectState = 3
	ObjectState_OBJECT_STATE_UNCHANGED      ObjectState = 4
	ObjectState_OBJECT_STATE_ALREADY_EXISTS ObjectState = 5
)

// Enum value maps for ObjectState.
var (
	ObjectState_name = map[int32]string{
		0: "OBJECT_STATE_UNSPECIFIED",
		1: "OBJECT_STATE_CREATE",
		2: "OBJECT_STATE_UPDATE",
		3: "OBJECT_STATE_DELETE",
		4: "OBJECT_STATE_UNCHANGED",
		5: "OBJECT_STATE_ALREADY_EXISTS",
	}
	ObjectState_value = map[string]int32{
		"OBJECT_STATE_UNSPECIFIED":    0,
		"OBJECT_STATE_CREATE":         1,
		"OBJECT_STATE_UPDATE":         2,
		"OBJECT_STATE_DELETE":         3,
		"OBJECT_STATE_UNCHANGED":      4,
		"OBJECT_STATE_ALREADY_EXISTS": 5,
	}
)

func (x ObjectState) Enum() *ObjectState {
	p := new(ObjectState)
	*p = x
	return p
}

func (x ObjectState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectState) Descriptor() protoreflect.EnumDescriptor {
	return file_operator_api_v1_pipeline_service_proto_enumTypes[0].Descriptor()
}

func (ObjectState) Type() protoreflect.EnumType {
	return &file_operator_api_v1_pipeline_service_proto_enumTypes[0]
}

func (x ObjectState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectState.Descriptor instead.
func (ObjectState) EnumDescriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{0}
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
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DryRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DryRunRequest) ProtoMessage() {}

func (x *DryRunRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DryRunRequest.ProtoReflect.Descriptor instead.
func (*DryRunRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{0}
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
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DryRunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DryRunResponse) ProtoMessage() {}

func (x *DryRunResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DryRunResponse.ProtoReflect.Descriptor instead.
func (*DryRunResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{1}
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

	Object *Object     `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	State  ObjectState `protobuf:"varint,2,opt,name=state,proto3,enum=api.v1.pipeline.ObjectState" json:"state,omitempty"`
}

func (x *ObjectChange) Reset() {
	*x = ObjectChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectChange) ProtoMessage() {}

func (x *ObjectChange) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ObjectChange.ProtoReflect.Descriptor instead.
func (*ObjectChange) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{2}
}

func (x *ObjectChange) GetObject() *Object {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *ObjectChange) GetState() ObjectState {
	if x != nil {
		return x.State
	}
	return ObjectState_OBJECT_STATE_UNSPECIFIED
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
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{3}
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

type GVK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group   string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Kind    string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *GVK) Reset() {
	*x = GVK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_pipeline_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GVK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GVK) ProtoMessage() {}

func (x *GVK) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GVK.ProtoReflect.Descriptor instead.
func (*GVK) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_pipeline_service_proto_rawDescGZIP(), []int{4}
}

func (x *GVK) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *GVK) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GVK) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

var File_operator_api_v1_pipeline_service_proto protoreflect.FileDescriptor

var file_operator_api_v1_pipeline_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0xf1, 0x01, 0x0a, 0x0d, 0x44, 0x72,
	0x79, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70,
	0x73, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x73,
	0x75, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x11, 0x61, 0x64, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x94, 0x01,
	0x0a, 0x0e, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x44,
	0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x22, 0x73, 0x0a, 0x0c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x5e, 0x0a, 0x06, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x26, 0x0a, 0x03, 0x67, 0x76, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2e, 0x47, 0x56, 0x4b, 0x52, 0x03, 0x67, 0x76, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x49, 0x0a, 0x03, 0x47, 0x56, 0x4b,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x2a, 0xb3, 0x01, 0x0a, 0x0b, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4f,
	0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x12, 0x1a, 0x0a,
	0x16, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x42, 0x4a,
	0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x05, 0x32, 0x56, 0x0a, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x12,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0xb8, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69,
	0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x50, 0xaa, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31,
	0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0xca, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c,
	0x56, 0x31, 0x5c, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0xe2, 0x02, 0x1b, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_operator_api_v1_pipeline_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_operator_api_v1_pipeline_service_proto_goTypes = []interface{}{
	(ObjectState)(0),       // 0: api.v1.pipeline.ObjectState
	(*DryRunRequest)(nil),  // 1: api.v1.pipeline.DryRunRequest
	(*DryRunResponse)(nil), // 2: api.v1.pipeline.DryRunResponse
	(*ObjectChange)(nil),   // 3: api.v1.pipeline.ObjectChange
	(*Object)(nil),         // 4: api.v1.pipeline.Object
	(*GVK)(nil),            // 5: api.v1.pipeline.GVK
}
var file_operator_api_v1_pipeline_service_proto_depIdxs = []int32{
	4, // 0: api.v1.pipeline.DryRunRequest.additional_objects:type_name -> api.v1.pipeline.Object
	4, // 1: api.v1.pipeline.DryRunResponse.input_objects:type_name -> api.v1.pipeline.Object
	3, // 2: api.v1.pipeline.DryRunResponse.output_objects:type_name -> api.v1.pipeline.ObjectChange
	4, // 3: api.v1.pipeline.ObjectChange.object:type_name -> api.v1.pipeline.Object
	0, // 4: api.v1.pipeline.ObjectChange.state:type_name -> api.v1.pipeline.ObjectState
	5, // 5: api.v1.pipeline.Object.gvk:type_name -> api.v1.pipeline.GVK
	1, // 6: api.v1.pipeline.Service.DryRun:input_type -> api.v1.pipeline.DryRunRequest
	2, // 7: api.v1.pipeline.Service.DryRun:output_type -> api.v1.pipeline.DryRunResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_operator_api_v1_pipeline_service_proto_init() }
func file_operator_api_v1_pipeline_service_proto_init() {
	if File_operator_api_v1_pipeline_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_operator_api_v1_pipeline_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_operator_api_v1_pipeline_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_operator_api_v1_pipeline_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_operator_api_v1_pipeline_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_operator_api_v1_pipeline_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GVK); i {
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
			NumMessages:   5,
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
