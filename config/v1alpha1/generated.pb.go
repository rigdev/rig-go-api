// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: config/v1alpha1/generated.proto

package configv1alpha1

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

type OperatorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind                  string    `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	ApiVersion            string    `protobuf:"bytes,2,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	WebhooksEnabled       bool      `protobuf:"varint,3,opt,name=webhooksEnabled,proto3" json:"webhooksEnabled,omitempty"`
	DevModeEnabled        bool      `protobuf:"varint,4,opt,name=devModeEnabled,proto3" json:"devModeEnabled,omitempty"`
	LeaderElectionEnabled bool      `protobuf:"varint,5,opt,name=leaderElectionEnabled,proto3" json:"leaderElectionEnabled,omitempty"`
	Pipeline              *Pipeline `protobuf:"bytes,7,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
}

func (x *OperatorConfig) Reset() {
	*x = OperatorConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1alpha1_generated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorConfig) ProtoMessage() {}

func (x *OperatorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1alpha1_generated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorConfig.ProtoReflect.Descriptor instead.
func (*OperatorConfig) Descriptor() ([]byte, []int) {
	return file_config_v1alpha1_generated_proto_rawDescGZIP(), []int{0}
}

func (x *OperatorConfig) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *OperatorConfig) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *OperatorConfig) GetWebhooksEnabled() bool {
	if x != nil {
		return x.WebhooksEnabled
	}
	return false
}

func (x *OperatorConfig) GetDevModeEnabled() bool {
	if x != nil {
		return x.DevModeEnabled
	}
	return false
}

func (x *OperatorConfig) GetLeaderElectionEnabled() bool {
	if x != nil {
		return x.LeaderElectionEnabled
	}
	return false
}

func (x *OperatorConfig) GetPipeline() *Pipeline {
	if x != nil {
		return x.Pipeline
	}
	return nil
}

type Pipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceAccountStep *CapsuleStep            `protobuf:"bytes,1,opt,name=serviceAccountStep,proto3" json:"serviceAccountStep,omitempty"`
	DeploymentStep     *CapsuleStep            `protobuf:"bytes,2,opt,name=deploymentStep,proto3" json:"deploymentStep,omitempty"`
	RoutesStep         *CapsuleStep            `protobuf:"bytes,3,opt,name=routesStep,proto3" json:"routesStep,omitempty"`
	CronJobsStep       *CapsuleStep            `protobuf:"bytes,4,opt,name=cronJobsStep,proto3" json:"cronJobsStep,omitempty"`
	VpaStep            *CapsuleStep            `protobuf:"bytes,5,opt,name=vpaStep,proto3" json:"vpaStep,omitempty"`
	ServiceMonitorStep *CapsuleStep            `protobuf:"bytes,6,opt,name=serviceMonitorStep,proto3" json:"serviceMonitorStep,omitempty"`
	Steps              []*Step                 `protobuf:"bytes,7,rep,name=steps,proto3" json:"steps,omitempty"`
	CustomPlugins      []*CustomPlugin         `protobuf:"bytes,8,rep,name=customPlugins,proto3" json:"customPlugins,omitempty"`
	CapsuleExtensions  map[string]*CapsuleStep `protobuf:"bytes,9,rep,name=capsuleExtensions,proto3" json:"capsuleExtensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Pipeline) Reset() {
	*x = Pipeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1alpha1_generated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pipeline) ProtoMessage() {}

func (x *Pipeline) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1alpha1_generated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pipeline.ProtoReflect.Descriptor instead.
func (*Pipeline) Descriptor() ([]byte, []int) {
	return file_config_v1alpha1_generated_proto_rawDescGZIP(), []int{1}
}

func (x *Pipeline) GetServiceAccountStep() *CapsuleStep {
	if x != nil {
		return x.ServiceAccountStep
	}
	return nil
}

func (x *Pipeline) GetDeploymentStep() *CapsuleStep {
	if x != nil {
		return x.DeploymentStep
	}
	return nil
}

func (x *Pipeline) GetRoutesStep() *CapsuleStep {
	if x != nil {
		return x.RoutesStep
	}
	return nil
}

func (x *Pipeline) GetCronJobsStep() *CapsuleStep {
	if x != nil {
		return x.CronJobsStep
	}
	return nil
}

func (x *Pipeline) GetVpaStep() *CapsuleStep {
	if x != nil {
		return x.VpaStep
	}
	return nil
}

func (x *Pipeline) GetServiceMonitorStep() *CapsuleStep {
	if x != nil {
		return x.ServiceMonitorStep
	}
	return nil
}

func (x *Pipeline) GetSteps() []*Step {
	if x != nil {
		return x.Steps
	}
	return nil
}

func (x *Pipeline) GetCustomPlugins() []*CustomPlugin {
	if x != nil {
		return x.CustomPlugins
	}
	return nil
}

func (x *Pipeline) GetCapsuleExtensions() map[string]*CapsuleStep {
	if x != nil {
		return x.CapsuleExtensions
	}
	return nil
}

type CapsuleStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin string `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Config string `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *CapsuleStep) Reset() {
	*x = CapsuleStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1alpha1_generated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapsuleStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapsuleStep) ProtoMessage() {}

func (x *CapsuleStep) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1alpha1_generated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapsuleStep.ProtoReflect.Descriptor instead.
func (*CapsuleStep) Descriptor() ([]byte, []int) {
	return file_config_v1alpha1_generated_proto_rawDescGZIP(), []int{2}
}

func (x *CapsuleStep) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *CapsuleStep) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag               string        `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Match             *CapsuleMatch `protobuf:"bytes,2,opt,name=match,proto3" json:"match,omitempty"`
	Plugins           []*Plugin     `protobuf:"bytes,3,rep,name=plugins,proto3" json:"plugins,omitempty"`
	Namespaces        []string      `protobuf:"bytes,4,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	Capsules          []string      `protobuf:"bytes,5,rep,name=capsules,proto3" json:"capsules,omitempty"`
	EnableForPlatform bool          `protobuf:"varint,6,opt,name=enableForPlatform,proto3" json:"enableForPlatform,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1alpha1_generated_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1alpha1_generated_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_config_v1alpha1_generated_proto_rawDescGZIP(), []int{3}
}

func (x *Step) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Step) GetMatch() *CapsuleMatch {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *Step) GetPlugins() []*Plugin {
	if x != nil {
		return x.Plugins
	}
	return nil
}

func (x *Step) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *Step) GetCapsules() []string {
	if x != nil {
		return x.Capsules
	}
	return nil
}

func (x *Step) GetEnableForPlatform() bool {
	if x != nil {
		return x.EnableForPlatform
	}
	return false
}

type CapsuleMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespaces        []string          `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	Names             []string          `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	Annotations       map[string]string `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EnableForPlatform bool              `protobuf:"varint,4,opt,name=enableForPlatform,proto3" json:"enableForPlatform,omitempty"`
}

func (x *CapsuleMatch) Reset() {
	*x = CapsuleMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1alpha1_generated_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapsuleMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapsuleMatch) ProtoMessage() {}

func (x *CapsuleMatch) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1alpha1_generated_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapsuleMatch.ProtoReflect.Descriptor instead.
func (*CapsuleMatch) Descriptor() ([]byte, []int) {
	return file_config_v1alpha1_generated_proto_rawDescGZIP(), []int{4}
}

func (x *CapsuleMatch) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *CapsuleMatch) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *CapsuleMatch) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *CapsuleMatch) GetEnableForPlatform() bool {
	if x != nil {
		return x.EnableForPlatform
	}
	return false
}

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag    string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Plugin string `protobuf:"bytes,3,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Config string `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1alpha1_generated_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1alpha1_generated_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_config_v1alpha1_generated_proto_rawDescGZIP(), []int{5}
}

func (x *Plugin) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Plugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plugin) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *Plugin) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

type CustomPlugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *CustomPlugin) Reset() {
	*x = CustomPlugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1alpha1_generated_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomPlugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomPlugin) ProtoMessage() {}

func (x *CustomPlugin) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1alpha1_generated_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomPlugin.ProtoReflect.Descriptor instead.
func (*CustomPlugin) Descriptor() ([]byte, []int) {
	return file_config_v1alpha1_generated_proto_rawDescGZIP(), []int{6}
}

func (x *CustomPlugin) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

var File_config_v1alpha1_generated_proto protoreflect.FileDescriptor

var file_config_v1alpha1_generated_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x22, 0x83, 0x02, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x69,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61,
	0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x77, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x4d, 0x6f, 0x64, 0x65, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x64, 0x65, 0x76,
	0x4d, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x15, 0x6c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0xda, 0x05, 0x0a, 0x08, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x4c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52,
	0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53,
	0x74, 0x65, 0x70, 0x12, 0x44, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x61,
	0x70, 0x73, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x65, 0x70, 0x12, 0x3c, 0x0a, 0x0a, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x53, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x0a, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x53, 0x74, 0x65, 0x70, 0x12, 0x40, 0x0a, 0x0c, 0x63, 0x72, 0x6f, 0x6e, 0x4a,
	0x6f, 0x62, 0x73, 0x53, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x6f,
	0x6e, 0x4a, 0x6f, 0x62, 0x73, 0x53, 0x74, 0x65, 0x70, 0x12, 0x36, 0x0a, 0x07, 0x76, 0x70, 0x61,
	0x53, 0x74, 0x65, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x61, 0x70,
	0x73, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x07, 0x76, 0x70, 0x61, 0x53, 0x74, 0x65,
	0x70, 0x12, 0x4c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x53, 0x74, 0x65, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x12, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x65, 0x70, 0x12,
	0x2b, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x12, 0x43, 0x0a, 0x0d,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x12, 0x5e, 0x0a, 0x11, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11,
	0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x62, 0x0a, 0x16, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43,
	0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a, 0x0b, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65,
	0x53, 0x74, 0x65, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0xea, 0x01, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12,
	0x33, 0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x07,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x73, 0x75,
	0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x70, 0x73, 0x75,
	0x6c, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x72,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x22, 0x84, 0x02, 0x0a, 0x0c, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x72,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5e, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x74, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x24, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0xbf,
	0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d,
	0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x0f,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2,
	0x02, 0x1b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_v1alpha1_generated_proto_rawDescOnce sync.Once
	file_config_v1alpha1_generated_proto_rawDescData = file_config_v1alpha1_generated_proto_rawDesc
)

func file_config_v1alpha1_generated_proto_rawDescGZIP() []byte {
	file_config_v1alpha1_generated_proto_rawDescOnce.Do(func() {
		file_config_v1alpha1_generated_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_v1alpha1_generated_proto_rawDescData)
	})
	return file_config_v1alpha1_generated_proto_rawDescData
}

var file_config_v1alpha1_generated_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_config_v1alpha1_generated_proto_goTypes = []any{
	(*OperatorConfig)(nil), // 0: config.v1alpha1.OperatorConfig
	(*Pipeline)(nil),       // 1: config.v1alpha1.Pipeline
	(*CapsuleStep)(nil),    // 2: config.v1alpha1.CapsuleStep
	(*Step)(nil),           // 3: config.v1alpha1.Step
	(*CapsuleMatch)(nil),   // 4: config.v1alpha1.CapsuleMatch
	(*Plugin)(nil),         // 5: config.v1alpha1.Plugin
	(*CustomPlugin)(nil),   // 6: config.v1alpha1.CustomPlugin
	nil,                    // 7: config.v1alpha1.Pipeline.CapsuleExtensionsEntry
	nil,                    // 8: config.v1alpha1.CapsuleMatch.AnnotationsEntry
}
var file_config_v1alpha1_generated_proto_depIdxs = []int32{
	1,  // 0: config.v1alpha1.OperatorConfig.pipeline:type_name -> config.v1alpha1.Pipeline
	2,  // 1: config.v1alpha1.Pipeline.serviceAccountStep:type_name -> config.v1alpha1.CapsuleStep
	2,  // 2: config.v1alpha1.Pipeline.deploymentStep:type_name -> config.v1alpha1.CapsuleStep
	2,  // 3: config.v1alpha1.Pipeline.routesStep:type_name -> config.v1alpha1.CapsuleStep
	2,  // 4: config.v1alpha1.Pipeline.cronJobsStep:type_name -> config.v1alpha1.CapsuleStep
	2,  // 5: config.v1alpha1.Pipeline.vpaStep:type_name -> config.v1alpha1.CapsuleStep
	2,  // 6: config.v1alpha1.Pipeline.serviceMonitorStep:type_name -> config.v1alpha1.CapsuleStep
	3,  // 7: config.v1alpha1.Pipeline.steps:type_name -> config.v1alpha1.Step
	6,  // 8: config.v1alpha1.Pipeline.customPlugins:type_name -> config.v1alpha1.CustomPlugin
	7,  // 9: config.v1alpha1.Pipeline.capsuleExtensions:type_name -> config.v1alpha1.Pipeline.CapsuleExtensionsEntry
	4,  // 10: config.v1alpha1.Step.match:type_name -> config.v1alpha1.CapsuleMatch
	5,  // 11: config.v1alpha1.Step.plugins:type_name -> config.v1alpha1.Plugin
	8,  // 12: config.v1alpha1.CapsuleMatch.annotations:type_name -> config.v1alpha1.CapsuleMatch.AnnotationsEntry
	2,  // 13: config.v1alpha1.Pipeline.CapsuleExtensionsEntry.value:type_name -> config.v1alpha1.CapsuleStep
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_config_v1alpha1_generated_proto_init() }
func file_config_v1alpha1_generated_proto_init() {
	if File_config_v1alpha1_generated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_v1alpha1_generated_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*OperatorConfig); i {
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
		file_config_v1alpha1_generated_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Pipeline); i {
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
		file_config_v1alpha1_generated_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CapsuleStep); i {
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
		file_config_v1alpha1_generated_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Step); i {
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
		file_config_v1alpha1_generated_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CapsuleMatch); i {
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
		file_config_v1alpha1_generated_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Plugin); i {
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
		file_config_v1alpha1_generated_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CustomPlugin); i {
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
			RawDescriptor: file_config_v1alpha1_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_v1alpha1_generated_proto_goTypes,
		DependencyIndexes: file_config_v1alpha1_generated_proto_depIdxs,
		MessageInfos:      file_config_v1alpha1_generated_proto_msgTypes,
	}.Build()
	File_config_v1alpha1_generated_proto = out.File
	file_config_v1alpha1_generated_proto_rawDesc = nil
	file_config_v1alpha1_generated_proto_goTypes = nil
	file_config_v1alpha1_generated_proto_depIdxs = nil
}