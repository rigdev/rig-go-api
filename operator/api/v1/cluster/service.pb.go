// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: operator/api/v1/cluster/service.proto

package cluster

import (
	model "github.com/rigdev/rig-go-api/model"
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

type GetNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNodesRequest) Reset() {
	*x = GetNodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_cluster_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodesRequest) ProtoMessage() {}

func (x *GetNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_cluster_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodesRequest.ProtoReflect.Descriptor instead.
func (*GetNodesRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_cluster_service_proto_rawDescGZIP(), []int{0}
}

type GetNodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *GetNodesResponse) Reset() {
	*x = GetNodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_cluster_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodesResponse) ProtoMessage() {}

func (x *GetNodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_cluster_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodesResponse.ProtoReflect.Descriptor instead.
func (*GetNodesResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_cluster_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetNodesResponse) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName     string           `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Allocateable *model.Resources `protobuf:"bytes,2,opt,name=allocateable,proto3" json:"allocateable,omitempty"`
	Usage        *model.Resources `protobuf:"bytes,3,opt,name=usage,proto3" json:"usage,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_cluster_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_cluster_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_cluster_service_proto_rawDescGZIP(), []int{2}
}

func (x *Node) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *Node) GetAllocateable() *model.Resources {
	if x != nil {
		return x.Allocateable
	}
	return nil
}

func (x *Node) GetUsage() *model.Resources {
	if x != nil {
		return x.Usage
	}
	return nil
}

type GetNodePodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
}

func (x *GetNodePodsRequest) Reset() {
	*x = GetNodePodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_cluster_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodePodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodePodsRequest) ProtoMessage() {}

func (x *GetNodePodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_cluster_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodePodsRequest.ProtoReflect.Descriptor instead.
func (*GetNodePodsRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_cluster_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetNodePodsRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

type GetNodePodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pods []*Pod `protobuf:"bytes,1,rep,name=pods,proto3" json:"pods,omitempty"`
}

func (x *GetNodePodsResponse) Reset() {
	*x = GetNodePodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_cluster_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodePodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodePodsResponse) ProtoMessage() {}

func (x *GetNodePodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_cluster_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodePodsResponse.ProtoReflect.Descriptor instead.
func (*GetNodePodsResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_cluster_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetNodePodsResponse) GetPods() []*Pod {
	if x != nil {
		return x.Pods
	}
	return nil
}

type Pod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodName     string           `protobuf:"bytes,1,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	Namespace   string           `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Requested   *model.Resources `protobuf:"bytes,3,opt,name=requested,proto3" json:"requested,omitempty"`
	CapsuleName string           `protobuf:"bytes,4,opt,name=capsule_name,json=capsuleName,proto3" json:"capsule_name,omitempty"`
}

func (x *Pod) Reset() {
	*x = Pod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_cluster_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pod) ProtoMessage() {}

func (x *Pod) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_cluster_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pod.ProtoReflect.Descriptor instead.
func (*Pod) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_cluster_service_proto_rawDescGZIP(), []int{5}
}

func (x *Pod) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *Pod) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Pod) GetRequested() *model.Resources {
	if x != nil {
		return x.Requested
	}
	return nil
}

func (x *Pod) GetCapsuleName() string {
	if x != nil {
		return x.CapsuleName
	}
	return ""
}

var File_operator_api_v1_cluster_service_proto protoreflect.FileDescriptor

var file_operator_api_v1_cluster_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x1a, 0x13, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x47, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x22, 0x81, 0x01, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52,
	0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x26, 0x0a,
	0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x05,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x04, 0x70, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x64, 0x52, 0x04, 0x70, 0x6f, 0x64,
	0x73, 0x22, 0x91, 0x01, 0x0a, 0x03, 0x50, 0x6f, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xd8, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x61, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x28, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x50,
	0x6f, 0x64, 0x73, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0xe1, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x67,
	0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0xa2, 0x02, 0x04, 0x41, 0x56, 0x4f, 0x43, 0xaa, 0x02, 0x17,
	0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0xca, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31,
	0x5c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x5c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56,
	0x31, 0x3a, 0x3a, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x3a, 0x3a, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_operator_api_v1_cluster_service_proto_rawDescOnce sync.Once
	file_operator_api_v1_cluster_service_proto_rawDescData = file_operator_api_v1_cluster_service_proto_rawDesc
)

func file_operator_api_v1_cluster_service_proto_rawDescGZIP() []byte {
	file_operator_api_v1_cluster_service_proto_rawDescOnce.Do(func() {
		file_operator_api_v1_cluster_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_operator_api_v1_cluster_service_proto_rawDescData)
	})
	return file_operator_api_v1_cluster_service_proto_rawDescData
}

var file_operator_api_v1_cluster_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_operator_api_v1_cluster_service_proto_goTypes = []any{
	(*GetNodesRequest)(nil),     // 0: api.v1.operator.cluster.GetNodesRequest
	(*GetNodesResponse)(nil),    // 1: api.v1.operator.cluster.GetNodesResponse
	(*Node)(nil),                // 2: api.v1.operator.cluster.Node
	(*GetNodePodsRequest)(nil),  // 3: api.v1.operator.cluster.GetNodePodsRequest
	(*GetNodePodsResponse)(nil), // 4: api.v1.operator.cluster.GetNodePodsResponse
	(*Pod)(nil),                 // 5: api.v1.operator.cluster.Pod
	(*model.Resources)(nil),     // 6: model.Resources
}
var file_operator_api_v1_cluster_service_proto_depIdxs = []int32{
	2, // 0: api.v1.operator.cluster.GetNodesResponse.nodes:type_name -> api.v1.operator.cluster.Node
	6, // 1: api.v1.operator.cluster.Node.allocateable:type_name -> model.Resources
	6, // 2: api.v1.operator.cluster.Node.usage:type_name -> model.Resources
	5, // 3: api.v1.operator.cluster.GetNodePodsResponse.pods:type_name -> api.v1.operator.cluster.Pod
	6, // 4: api.v1.operator.cluster.Pod.requested:type_name -> model.Resources
	0, // 5: api.v1.operator.cluster.Service.GetNodes:input_type -> api.v1.operator.cluster.GetNodesRequest
	3, // 6: api.v1.operator.cluster.Service.GetNodePods:input_type -> api.v1.operator.cluster.GetNodePodsRequest
	1, // 7: api.v1.operator.cluster.Service.GetNodes:output_type -> api.v1.operator.cluster.GetNodesResponse
	4, // 8: api.v1.operator.cluster.Service.GetNodePods:output_type -> api.v1.operator.cluster.GetNodePodsResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_operator_api_v1_cluster_service_proto_init() }
func file_operator_api_v1_cluster_service_proto_init() {
	if File_operator_api_v1_cluster_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_operator_api_v1_cluster_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetNodesRequest); i {
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
		file_operator_api_v1_cluster_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetNodesResponse); i {
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
		file_operator_api_v1_cluster_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Node); i {
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
		file_operator_api_v1_cluster_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetNodePodsRequest); i {
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
		file_operator_api_v1_cluster_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetNodePodsResponse); i {
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
		file_operator_api_v1_cluster_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Pod); i {
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
			RawDescriptor: file_operator_api_v1_cluster_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_operator_api_v1_cluster_service_proto_goTypes,
		DependencyIndexes: file_operator_api_v1_cluster_service_proto_depIdxs,
		MessageInfos:      file_operator_api_v1_cluster_service_proto_msgTypes,
	}.Build()
	File_operator_api_v1_cluster_service_proto = out.File
	file_operator_api_v1_cluster_service_proto_rawDesc = nil
	file_operator_api_v1_cluster_service_proto_goTypes = nil
	file_operator_api_v1_cluster_service_proto_depIdxs = nil
}
