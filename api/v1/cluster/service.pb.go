// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/v1/cluster/service.proto

package cluster

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

// Cluster type - Docker or kubernetes.
type ClusterType int32

const (
	ClusterType_CLUSTER_TYPE_UNSPECIFIED ClusterType = 0
	ClusterType_CLUSTER_TYPE_DOCKER      ClusterType = 1
	ClusterType_CLUSTER_TYPE_KUBERNETES  ClusterType = 2
)

// Enum value maps for ClusterType.
var (
	ClusterType_name = map[int32]string{
		0: "CLUSTER_TYPE_UNSPECIFIED",
		1: "CLUSTER_TYPE_DOCKER",
		2: "CLUSTER_TYPE_KUBERNETES",
	}
	ClusterType_value = map[string]int32{
		"CLUSTER_TYPE_UNSPECIFIED": 0,
		"CLUSTER_TYPE_DOCKER":      1,
		"CLUSTER_TYPE_KUBERNETES":  2,
	}
)

func (x ClusterType) Enum() *ClusterType {
	p := new(ClusterType)
	*p = x
	return p
}

func (x ClusterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClusterType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_cluster_service_proto_enumTypes[0].Descriptor()
}

func (ClusterType) Type() protoreflect.EnumType {
	return &file_api_v1_cluster_service_proto_enumTypes[0]
}

func (x ClusterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClusterType.Descriptor instead.
func (ClusterType) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{0}
}

// request for getting cluster config for an environment.
type GetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvironmentId string `protobuf:"bytes,1,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"` // The environment to get cluster config for.
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cluster_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cluster_service_proto_msgTypes[0]
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
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetConfigRequest) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

// response for getting cluster config for an environment.
type GetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterType ClusterType `protobuf:"varint,1,opt,name=cluster_type,json=clusterType,proto3,enum=api.v1.cluster.ClusterType" json:"cluster_type,omitempty"` // Type of the cluster.
	// dev registry of the cluster. This is either a Docker daemon or a registry.
	//
	// Types that are assignable to DevRegistry:
	//
	//	*GetConfigResponse_Docker
	//	*GetConfigResponse_Registry
	DevRegistry isGetConfigResponse_DevRegistry `protobuf_oneof:"dev_registry"`
	Ingress     bool                            `protobuf:"varint,4,opt,name=ingress,proto3" json:"ingress,omitempty"` // if true, the cluster has an ingress controller.
}

func (x *GetConfigResponse) Reset() {
	*x = GetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cluster_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigResponse) ProtoMessage() {}

func (x *GetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cluster_service_proto_msgTypes[1]
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
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetConfigResponse) GetClusterType() ClusterType {
	if x != nil {
		return x.ClusterType
	}
	return ClusterType_CLUSTER_TYPE_UNSPECIFIED
}

func (m *GetConfigResponse) GetDevRegistry() isGetConfigResponse_DevRegistry {
	if m != nil {
		return m.DevRegistry
	}
	return nil
}

func (x *GetConfigResponse) GetDocker() *DockerDaemon {
	if x, ok := x.GetDevRegistry().(*GetConfigResponse_Docker); ok {
		return x.Docker
	}
	return nil
}

func (x *GetConfigResponse) GetRegistry() *Registry {
	if x, ok := x.GetDevRegistry().(*GetConfigResponse_Registry); ok {
		return x.Registry
	}
	return nil
}

func (x *GetConfigResponse) GetIngress() bool {
	if x != nil {
		return x.Ingress
	}
	return false
}

type isGetConfigResponse_DevRegistry interface {
	isGetConfigResponse_DevRegistry()
}

type GetConfigResponse_Docker struct {
	Docker *DockerDaemon `protobuf:"bytes,2,opt,name=docker,proto3,oneof"` // Docker.
}

type GetConfigResponse_Registry struct {
	Registry *Registry `protobuf:"bytes,3,opt,name=registry,proto3,oneof"` // Registry.
}

func (*GetConfigResponse_Docker) isGetConfigResponse_DevRegistry() {}

func (*GetConfigResponse_Registry) isGetConfigResponse_DevRegistry() {}

// Empty Request for getting the configs of all clusters.
type GetConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConfigsRequest) Reset() {
	*x = GetConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cluster_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigsRequest) ProtoMessage() {}

func (x *GetConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cluster_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigsRequest.ProtoReflect.Descriptor instead.
func (*GetConfigsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{2}
}

// Empty Response for getting the configs of all clusters.
type GetConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters []*GetConfigResponse `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *GetConfigsResponse) Reset() {
	*x = GetConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cluster_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigsResponse) ProtoMessage() {}

func (x *GetConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cluster_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigsResponse.ProtoReflect.Descriptor instead.
func (*GetConfigsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetConfigsResponse) GetClusters() []*GetConfigResponse {
	if x != nil {
		return x.Clusters
	}
	return nil
}

// Docker daemon dev registry
type DockerDaemon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DockerDaemon) Reset() {
	*x = DockerDaemon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cluster_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerDaemon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerDaemon) ProtoMessage() {}

func (x *DockerDaemon) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cluster_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerDaemon.ProtoReflect.Descriptor instead.
func (*DockerDaemon) Descriptor() ([]byte, []int) {
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{4}
}

// Registry dev registry
type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *Registry) Reset() {
	*x = Registry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cluster_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cluster_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{5}
}

func (x *Registry) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

var File_api_v1_cluster_service_proto protoreflect.FileDescriptor

var file_api_v1_cluster_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x39,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xed, 0x01, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x36, 0x0a, 0x06, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x06, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x64, 0x65, 0x76,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x53,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x44, 0x61, 0x65,
	0x6d, 0x6f, 0x6e, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x2a, 0x61, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x17, 0x0a, 0x13, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x4f, 0x43, 0x4b, 0x45, 0x52, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4c, 0x55,
	0x53, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e,
	0x45, 0x54, 0x45, 0x53, 0x10, 0x02, 0x32, 0xb4, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x52, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xa9, 0x01,
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x43, 0xaa, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0xca, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x5c, 0x56,
	0x31, 0x5c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0xe2, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x5c,
	0x56, 0x31, 0x5c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31,
	0x3a, 0x3a, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v1_cluster_service_proto_rawDescOnce sync.Once
	file_api_v1_cluster_service_proto_rawDescData = file_api_v1_cluster_service_proto_rawDesc
)

func file_api_v1_cluster_service_proto_rawDescGZIP() []byte {
	file_api_v1_cluster_service_proto_rawDescOnce.Do(func() {
		file_api_v1_cluster_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_cluster_service_proto_rawDescData)
	})
	return file_api_v1_cluster_service_proto_rawDescData
}

var file_api_v1_cluster_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_cluster_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_v1_cluster_service_proto_goTypes = []interface{}{
	(ClusterType)(0),           // 0: api.v1.cluster.ClusterType
	(*GetConfigRequest)(nil),   // 1: api.v1.cluster.GetConfigRequest
	(*GetConfigResponse)(nil),  // 2: api.v1.cluster.GetConfigResponse
	(*GetConfigsRequest)(nil),  // 3: api.v1.cluster.GetConfigsRequest
	(*GetConfigsResponse)(nil), // 4: api.v1.cluster.GetConfigsResponse
	(*DockerDaemon)(nil),       // 5: api.v1.cluster.DockerDaemon
	(*Registry)(nil),           // 6: api.v1.cluster.Registry
}
var file_api_v1_cluster_service_proto_depIdxs = []int32{
	0, // 0: api.v1.cluster.GetConfigResponse.cluster_type:type_name -> api.v1.cluster.ClusterType
	5, // 1: api.v1.cluster.GetConfigResponse.docker:type_name -> api.v1.cluster.DockerDaemon
	6, // 2: api.v1.cluster.GetConfigResponse.registry:type_name -> api.v1.cluster.Registry
	2, // 3: api.v1.cluster.GetConfigsResponse.clusters:type_name -> api.v1.cluster.GetConfigResponse
	1, // 4: api.v1.cluster.Service.GetConfig:input_type -> api.v1.cluster.GetConfigRequest
	3, // 5: api.v1.cluster.Service.GetConfigs:input_type -> api.v1.cluster.GetConfigsRequest
	2, // 6: api.v1.cluster.Service.GetConfig:output_type -> api.v1.cluster.GetConfigResponse
	4, // 7: api.v1.cluster.Service.GetConfigs:output_type -> api.v1.cluster.GetConfigsResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1_cluster_service_proto_init() }
func file_api_v1_cluster_service_proto_init() {
	if File_api_v1_cluster_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_cluster_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_v1_cluster_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_v1_cluster_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigsRequest); i {
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
		file_api_v1_cluster_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigsResponse); i {
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
		file_api_v1_cluster_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerDaemon); i {
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
		file_api_v1_cluster_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry); i {
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
	file_api_v1_cluster_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GetConfigResponse_Docker)(nil),
		(*GetConfigResponse_Registry)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_cluster_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_cluster_service_proto_goTypes,
		DependencyIndexes: file_api_v1_cluster_service_proto_depIdxs,
		EnumInfos:         file_api_v1_cluster_service_proto_enumTypes,
		MessageInfos:      file_api_v1_cluster_service_proto_msgTypes,
	}.Build()
	File_api_v1_cluster_service_proto = out.File
	file_api_v1_cluster_service_proto_rawDesc = nil
	file_api_v1_cluster_service_proto_goTypes = nil
	file_api_v1_cluster_service_proto_depIdxs = nil
}
