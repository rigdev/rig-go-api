// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: operator/api/v1/capabilities/service.proto

package capabilities

import (
	v1alpha1 "github.com/rigdev/rig-go-api/config/v1alpha1"
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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_capabilities_service_proto_rawDescGZIP(), []int{0}
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ingress                     bool `protobuf:"varint,1,opt,name=ingress,proto3" json:"ingress,omitempty"`
	HasPrometheusServiceMonitor bool `protobuf:"varint,2,opt,name=has_prometheus_service_monitor,json=hasPrometheusServiceMonitor,proto3" json:"has_prometheus_service_monitor,omitempty"`
	HasCustomMetrics            bool `protobuf:"varint,3,opt,name=has_custom_metrics,json=hasCustomMetrics,proto3" json:"has_custom_metrics,omitempty"`
	HasVerticalPodAutoscaler    bool `protobuf:"varint,4,opt,name=has_vertical_pod_autoscaler,json=hasVerticalPodAutoscaler,proto3" json:"has_vertical_pod_autoscaler,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_capabilities_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetIngress() bool {
	if x != nil {
		return x.Ingress
	}
	return false
}

func (x *GetResponse) GetHasPrometheusServiceMonitor() bool {
	if x != nil {
		return x.HasPrometheusServiceMonitor
	}
	return false
}

func (x *GetResponse) GetHasCustomMetrics() bool {
	if x != nil {
		return x.HasCustomMetrics
	}
	return false
}

func (x *GetResponse) GetHasVerticalPodAutoscaler() bool {
	if x != nil {
		return x.HasVerticalPodAutoscaler
	}
	return false
}

type GetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigRequest.ProtoReflect.Descriptor instead.
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_capabilities_service_proto_rawDescGZIP(), []int{2}
}

type GetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yaml            string                   `protobuf:"bytes,1,opt,name=yaml,proto3" json:"yaml,omitempty"`
	OperatorConfig  *v1alpha1.OperatorConfig `protobuf:"bytes,2,opt,name=operator_config,json=operatorConfig,proto3" json:"operator_config,omitempty"`
	OperatorVersion string                   `protobuf:"bytes,3,opt,name=operator_version,json=operatorVersion,proto3" json:"operator_version,omitempty"`
	K8SVersion      string                   `protobuf:"bytes,4,opt,name=k8s_version,json=k8sVersion,proto3" json:"k8s_version,omitempty"`
}

func (x *GetConfigResponse) Reset() {
	*x = GetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigResponse) ProtoMessage() {}

func (x *GetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigResponse.ProtoReflect.Descriptor instead.
func (*GetConfigResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_capabilities_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetConfigResponse) GetYaml() string {
	if x != nil {
		return x.Yaml
	}
	return ""
}

func (x *GetConfigResponse) GetOperatorConfig() *v1alpha1.OperatorConfig {
	if x != nil {
		return x.OperatorConfig
	}
	return nil
}

func (x *GetConfigResponse) GetOperatorVersion() string {
	if x != nil {
		return x.OperatorVersion
	}
	return ""
}

func (x *GetConfigResponse) GetK8SVersion() string {
	if x != nil {
		return x.K8SVersion
	}
	return ""
}

type GetPluginsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPluginsRequest) Reset() {
	*x = GetPluginsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginsRequest) ProtoMessage() {}

func (x *GetPluginsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginsRequest.ProtoReflect.Descriptor instead.
func (*GetPluginsRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_capabilities_service_proto_rawDescGZIP(), []int{4}
}

type GetPluginsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugins []*GetPluginsResponse_Plugin `protobuf:"bytes,1,rep,name=plugins,proto3" json:"plugins,omitempty"`
}

func (x *GetPluginsResponse) Reset() {
	*x = GetPluginsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginsResponse) ProtoMessage() {}

func (x *GetPluginsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginsResponse.ProtoReflect.Descriptor instead.
func (*GetPluginsResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_capabilities_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetPluginsResponse) GetPlugins() []*GetPluginsResponse_Plugin {
	if x != nil {
		return x.Plugins
	}
	return nil
}

type GetPluginsResponse_Builtin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPluginsResponse_Builtin) Reset() {
	*x = GetPluginsResponse_Builtin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginsResponse_Builtin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginsResponse_Builtin) ProtoMessage() {}

func (x *GetPluginsResponse_Builtin) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginsResponse_Builtin.ProtoReflect.Descriptor instead.
func (*GetPluginsResponse_Builtin) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_capabilities_service_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetPluginsResponse_Builtin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetPluginsResponse_Thirdparty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetPluginsResponse_Thirdparty) Reset() {
	*x = GetPluginsResponse_Thirdparty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginsResponse_Thirdparty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginsResponse_Thirdparty) ProtoMessage() {}

func (x *GetPluginsResponse_Thirdparty) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginsResponse_Thirdparty.ProtoReflect.Descriptor instead.
func (*GetPluginsResponse_Thirdparty) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_capabilities_service_proto_rawDescGZIP(), []int{5, 1}
}

func (x *GetPluginsResponse_Thirdparty) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPluginsResponse_Thirdparty) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type GetPluginsResponse_Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Plugin:
	//
	//	*GetPluginsResponse_Plugin_Builtin
	//	*GetPluginsResponse_Plugin_ThirdParty
	Plugin isGetPluginsResponse_Plugin_Plugin `protobuf_oneof:"plugin"`
}

func (x *GetPluginsResponse_Plugin) Reset() {
	*x = GetPluginsResponse_Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginsResponse_Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginsResponse_Plugin) ProtoMessage() {}

func (x *GetPluginsResponse_Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_capabilities_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginsResponse_Plugin.ProtoReflect.Descriptor instead.
func (*GetPluginsResponse_Plugin) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_capabilities_service_proto_rawDescGZIP(), []int{5, 2}
}

func (m *GetPluginsResponse_Plugin) GetPlugin() isGetPluginsResponse_Plugin_Plugin {
	if m != nil {
		return m.Plugin
	}
	return nil
}

func (x *GetPluginsResponse_Plugin) GetBuiltin() *GetPluginsResponse_Builtin {
	if x, ok := x.GetPlugin().(*GetPluginsResponse_Plugin_Builtin); ok {
		return x.Builtin
	}
	return nil
}

func (x *GetPluginsResponse_Plugin) GetThirdParty() *GetPluginsResponse_Thirdparty {
	if x, ok := x.GetPlugin().(*GetPluginsResponse_Plugin_ThirdParty); ok {
		return x.ThirdParty
	}
	return nil
}

type isGetPluginsResponse_Plugin_Plugin interface {
	isGetPluginsResponse_Plugin_Plugin()
}

type GetPluginsResponse_Plugin_Builtin struct {
	Builtin *GetPluginsResponse_Builtin `protobuf:"bytes,1,opt,name=builtin,proto3,oneof"`
}

type GetPluginsResponse_Plugin_ThirdParty struct {
	ThirdParty *GetPluginsResponse_Thirdparty `protobuf:"bytes,2,opt,name=third_party,json=thirdParty,proto3,oneof"`
}

func (*GetPluginsResponse_Plugin_Builtin) isGetPluginsResponse_Plugin_Plugin() {}

func (*GetPluginsResponse_Plugin_ThirdParty) isGetPluginsResponse_Plugin_Plugin() {}

var File_operator_api_v1_capabilities_service_proto protoreflect.FileDescriptor

var file_operator_api_v1_capabilities_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x1a, 0x1f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xd9, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x43, 0x0a, 0x1e, 0x68, 0x61,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x1b, 0x68, 0x61, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12,
	0x2c, 0x0a, 0x12, 0x68, 0x61, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x68, 0x61, 0x73,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x3d, 0x0a,
	0x1b, 0x68, 0x61, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x6f,
	0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x18, 0x68, 0x61, 0x73, 0x56, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x50,
	0x6f, 0x64, 0x41, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x22, 0x12, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xbd, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x12, 0x48, 0x0a, 0x0f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x6b, 0x38, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x38, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xee, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x07,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x07, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x1a, 0x1d, 0x0a, 0x07, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x36, 0x0a, 0x0a, 0x54, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0xb6, 0x01,
	0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x4b, 0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c,
	0x74, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x07, 0x62, 0x75,
	0x69, 0x6c, 0x74, 0x69, 0x6e, 0x12, 0x55, 0x0a, 0x0b, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x54, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x48, 0x00,
	0x52, 0x0a, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x42, 0x08, 0x0a, 0x06,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x32, 0x94, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xd0, 0x01,
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67,
	0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x43, 0xaa, 0x02, 0x13, 0x41, 0x70, 0x69,
	0x2e, 0x56, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0xca, 0x02, 0x13, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0xe2, 0x02, 0x1f, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c,
	0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x41, 0x70, 0x69, 0x3a, 0x3a,
	0x56, 0x31, 0x3a, 0x3a, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_operator_api_v1_capabilities_service_proto_rawDescOnce sync.Once
	file_operator_api_v1_capabilities_service_proto_rawDescData = file_operator_api_v1_capabilities_service_proto_rawDesc
)

func file_operator_api_v1_capabilities_service_proto_rawDescGZIP() []byte {
	file_operator_api_v1_capabilities_service_proto_rawDescOnce.Do(func() {
		file_operator_api_v1_capabilities_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_operator_api_v1_capabilities_service_proto_rawDescData)
	})
	return file_operator_api_v1_capabilities_service_proto_rawDescData
}

var file_operator_api_v1_capabilities_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_operator_api_v1_capabilities_service_proto_goTypes = []any{
	(*GetRequest)(nil),                    // 0: api.v1.capabilities.GetRequest
	(*GetResponse)(nil),                   // 1: api.v1.capabilities.GetResponse
	(*GetConfigRequest)(nil),              // 2: api.v1.capabilities.GetConfigRequest
	(*GetConfigResponse)(nil),             // 3: api.v1.capabilities.GetConfigResponse
	(*GetPluginsRequest)(nil),             // 4: api.v1.capabilities.GetPluginsRequest
	(*GetPluginsResponse)(nil),            // 5: api.v1.capabilities.GetPluginsResponse
	(*GetPluginsResponse_Builtin)(nil),    // 6: api.v1.capabilities.GetPluginsResponse.Builtin
	(*GetPluginsResponse_Thirdparty)(nil), // 7: api.v1.capabilities.GetPluginsResponse.Thirdparty
	(*GetPluginsResponse_Plugin)(nil),     // 8: api.v1.capabilities.GetPluginsResponse.Plugin
	(*v1alpha1.OperatorConfig)(nil),       // 9: config.v1alpha1.OperatorConfig
}
var file_operator_api_v1_capabilities_service_proto_depIdxs = []int32{
	9, // 0: api.v1.capabilities.GetConfigResponse.operator_config:type_name -> config.v1alpha1.OperatorConfig
	8, // 1: api.v1.capabilities.GetPluginsResponse.plugins:type_name -> api.v1.capabilities.GetPluginsResponse.Plugin
	6, // 2: api.v1.capabilities.GetPluginsResponse.Plugin.builtin:type_name -> api.v1.capabilities.GetPluginsResponse.Builtin
	7, // 3: api.v1.capabilities.GetPluginsResponse.Plugin.third_party:type_name -> api.v1.capabilities.GetPluginsResponse.Thirdparty
	2, // 4: api.v1.capabilities.Service.GetConfig:input_type -> api.v1.capabilities.GetConfigRequest
	0, // 5: api.v1.capabilities.Service.Get:input_type -> api.v1.capabilities.GetRequest
	4, // 6: api.v1.capabilities.Service.GetPlugins:input_type -> api.v1.capabilities.GetPluginsRequest
	3, // 7: api.v1.capabilities.Service.GetConfig:output_type -> api.v1.capabilities.GetConfigResponse
	1, // 8: api.v1.capabilities.Service.Get:output_type -> api.v1.capabilities.GetResponse
	5, // 9: api.v1.capabilities.Service.GetPlugins:output_type -> api.v1.capabilities.GetPluginsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_operator_api_v1_capabilities_service_proto_init() }
func file_operator_api_v1_capabilities_service_proto_init() {
	if File_operator_api_v1_capabilities_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_operator_api_v1_capabilities_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetRequest); i {
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
		file_operator_api_v1_capabilities_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetResponse); i {
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
		file_operator_api_v1_capabilities_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetConfigRequest); i {
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
		file_operator_api_v1_capabilities_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetConfigResponse); i {
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
		file_operator_api_v1_capabilities_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetPluginsRequest); i {
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
		file_operator_api_v1_capabilities_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetPluginsResponse); i {
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
		file_operator_api_v1_capabilities_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetPluginsResponse_Builtin); i {
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
		file_operator_api_v1_capabilities_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetPluginsResponse_Thirdparty); i {
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
		file_operator_api_v1_capabilities_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetPluginsResponse_Plugin); i {
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
	file_operator_api_v1_capabilities_service_proto_msgTypes[8].OneofWrappers = []any{
		(*GetPluginsResponse_Plugin_Builtin)(nil),
		(*GetPluginsResponse_Plugin_ThirdParty)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_operator_api_v1_capabilities_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_operator_api_v1_capabilities_service_proto_goTypes,
		DependencyIndexes: file_operator_api_v1_capabilities_service_proto_depIdxs,
		MessageInfos:      file_operator_api_v1_capabilities_service_proto_msgTypes,
	}.Build()
	File_operator_api_v1_capabilities_service_proto = out.File
	file_operator_api_v1_capabilities_service_proto_rawDesc = nil
	file_operator_api_v1_capabilities_service_proto_goTypes = nil
	file_operator_api_v1_capabilities_service_proto_depIdxs = nil
}
