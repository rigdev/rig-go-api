// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: operator/api/v1/plugin/service.proto

package plugin

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

type InitializeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginConfig   string `protobuf:"bytes,1,opt,name=plugin_config,json=pluginConfig,proto3" json:"plugin_config,omitempty"`
	OperatorConfig []byte `protobuf:"bytes,2,opt,name=operator_config,json=operatorConfig,proto3" json:"operator_config,omitempty"`
	Tag            string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *InitializeRequest) Reset() {
	*x = InitializeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_plugin_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeRequest) ProtoMessage() {}

func (x *InitializeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_plugin_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeRequest.ProtoReflect.Descriptor instead.
func (*InitializeRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_plugin_service_proto_rawDescGZIP(), []int{0}
}

func (x *InitializeRequest) GetPluginConfig() string {
	if x != nil {
		return x.PluginConfig
	}
	return ""
}

func (x *InitializeRequest) GetOperatorConfig() []byte {
	if x != nil {
		return x.OperatorConfig
	}
	return nil
}

func (x *InitializeRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type InitializeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitializeResponse) Reset() {
	*x = InitializeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_plugin_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeResponse) ProtoMessage() {}

func (x *InitializeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_plugin_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeResponse.ProtoReflect.Descriptor instead.
func (*InitializeResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_plugin_service_proto_rawDescGZIP(), []int{1}
}

type RunCapsuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunServer     uint32 `protobuf:"varint,1,opt,name=run_server,json=runServer,proto3" json:"run_server,omitempty"`
	CapsuleObject []byte `protobuf:"bytes,2,opt,name=capsule_object,json=capsuleObject,proto3" json:"capsule_object,omitempty"`
}

func (x *RunCapsuleRequest) Reset() {
	*x = RunCapsuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_plugin_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCapsuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCapsuleRequest) ProtoMessage() {}

func (x *RunCapsuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_plugin_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCapsuleRequest.ProtoReflect.Descriptor instead.
func (*RunCapsuleRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_plugin_service_proto_rawDescGZIP(), []int{2}
}

func (x *RunCapsuleRequest) GetRunServer() uint32 {
	if x != nil {
		return x.RunServer
	}
	return 0
}

func (x *RunCapsuleRequest) GetCapsuleObject() []byte {
	if x != nil {
		return x.CapsuleObject
	}
	return nil
}

type RunCapsuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RunCapsuleResponse) Reset() {
	*x = RunCapsuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_plugin_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCapsuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCapsuleResponse) ProtoMessage() {}

func (x *RunCapsuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_plugin_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCapsuleResponse.ProtoReflect.Descriptor instead.
func (*RunCapsuleResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_plugin_service_proto_rawDescGZIP(), []int{3}
}

type GetObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gvk     *GVK   `protobuf:"bytes,1,opt,name=gvk,proto3" json:"gvk,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Current bool   `protobuf:"varint,3,opt,name=current,proto3" json:"current,omitempty"`
}

func (x *GetObjectRequest) Reset() {
	*x = GetObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_plugin_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectRequest) ProtoMessage() {}

func (x *GetObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_plugin_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectRequest.ProtoReflect.Descriptor instead.
func (*GetObjectRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_plugin_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetObjectRequest) GetGvk() *GVK {
	if x != nil {
		return x.Gvk
	}
	return nil
}

func (x *GetObjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetObjectRequest) GetCurrent() bool {
	if x != nil {
		return x.Current
	}
	return false
}

type GetObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object []byte `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *GetObjectResponse) Reset() {
	*x = GetObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_plugin_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectResponse) ProtoMessage() {}

func (x *GetObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_plugin_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectResponse.ProtoReflect.Descriptor instead.
func (*GetObjectResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_plugin_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetObjectResponse) GetObject() []byte {
	if x != nil {
		return x.Object
	}
	return nil
}

type SetObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gvk    *GVK   `protobuf:"bytes,1,opt,name=gvk,proto3" json:"gvk,omitempty"`
	Object []byte `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *SetObjectRequest) Reset() {
	*x = SetObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_plugin_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetObjectRequest) ProtoMessage() {}

func (x *SetObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_plugin_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetObjectRequest.ProtoReflect.Descriptor instead.
func (*SetObjectRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_plugin_service_proto_rawDescGZIP(), []int{6}
}

func (x *SetObjectRequest) GetGvk() *GVK {
	if x != nil {
		return x.Gvk
	}
	return nil
}

func (x *SetObjectRequest) GetObject() []byte {
	if x != nil {
		return x.Object
	}
	return nil
}

type SetObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetObjectResponse) Reset() {
	*x = SetObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_plugin_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetObjectResponse) ProtoMessage() {}

func (x *SetObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_plugin_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetObjectResponse.ProtoReflect.Descriptor instead.
func (*SetObjectResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_plugin_service_proto_rawDescGZIP(), []int{7}
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
		mi := &file_operator_api_v1_plugin_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GVK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GVK) ProtoMessage() {}

func (x *GVK) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_plugin_service_proto_msgTypes[8]
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
	return file_operator_api_v1_plugin_service_proto_rawDescGZIP(), []int{8}
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

type DeleteObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gvk    *GVK   `protobuf:"bytes,1,opt,name=gvk,proto3" json:"gvk,omitempty"`
	Object []byte `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *DeleteObjectRequest) Reset() {
	*x = DeleteObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_plugin_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteObjectRequest) ProtoMessage() {}

func (x *DeleteObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_plugin_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteObjectRequest.ProtoReflect.Descriptor instead.
func (*DeleteObjectRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_plugin_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteObjectRequest) GetGvk() *GVK {
	if x != nil {
		return x.Gvk
	}
	return nil
}

func (x *DeleteObjectRequest) GetObject() []byte {
	if x != nil {
		return x.Object
	}
	return nil
}

type DeleteObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteObjectResponse) Reset() {
	*x = DeleteObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_plugin_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteObjectResponse) ProtoMessage() {}

func (x *DeleteObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_plugin_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteObjectResponse.ProtoReflect.Descriptor instead.
func (*DeleteObjectResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_plugin_service_proto_rawDescGZIP(), []int{10}
}

type MarkUsedObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gvk     *GVK   `protobuf:"bytes,1,opt,name=gvk,proto3" json:"gvk,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State   string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MarkUsedObjectRequest) Reset() {
	*x = MarkUsedObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_plugin_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkUsedObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkUsedObjectRequest) ProtoMessage() {}

func (x *MarkUsedObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_plugin_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkUsedObjectRequest.ProtoReflect.Descriptor instead.
func (*MarkUsedObjectRequest) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_plugin_service_proto_rawDescGZIP(), []int{11}
}

func (x *MarkUsedObjectRequest) GetGvk() *GVK {
	if x != nil {
		return x.Gvk
	}
	return nil
}

func (x *MarkUsedObjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MarkUsedObjectRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *MarkUsedObjectRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MarkUsedObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkUsedObjectResponse) Reset() {
	*x = MarkUsedObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_api_v1_plugin_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkUsedObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkUsedObjectResponse) ProtoMessage() {}

func (x *MarkUsedObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operator_api_v1_plugin_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkUsedObjectResponse.ProtoReflect.Descriptor instead.
func (*MarkUsedObjectResponse) Descriptor() ([]byte, []int) {
	return file_operator_api_v1_plugin_service_proto_rawDescGZIP(), []int{12}
}

var File_operator_api_v1_plugin_service_proto protoreflect.FileDescriptor

var file_operator_api_v1_plugin_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x22, 0x73, 0x0a, 0x11, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x14, 0x0a, 0x12, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x59, 0x0a, 0x11, 0x52, 0x75, 0x6e, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x75, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x75, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x63, 0x61,
	0x70, 0x73, 0x75, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x52,
	0x75, 0x6e, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x66, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x03, 0x67, 0x76, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x47, 0x56, 0x4b, 0x52, 0x03, 0x67, 0x76, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x50, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x03, 0x67, 0x76,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x56, 0x4b, 0x52, 0x03, 0x67, 0x76, 0x6b,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0x0a,
	0x03, 0x47, 0x56, 0x4b, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x53, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x03, 0x67, 0x76, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x56, 0x4b,
	0x52, 0x03, 0x67, 0x76, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x16, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x15, 0x4d, 0x61, 0x72, 0x6b, 0x55, 0x73,
	0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x03, 0x67, 0x76, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x56, 0x4b,
	0x52, 0x03, 0x67, 0x76, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x4d, 0x61, 0x72,
	0x6b, 0x55, 0x73, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xb9, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0a, 0x52, 0x75,
	0x6e, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x61, 0x70, 0x73,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x61,
	0x70, 0x73, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32,
	0xf0, 0x02, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x53, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5f, 0x0a, 0x0e, 0x4d, 0x61, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x64, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x55, 0x73,
	0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0xac, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d,
	0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0xa2, 0x02, 0x03,
	0x41, 0x56, 0x50, 0xaa, 0x02, 0x0d, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0xca, 0x02, 0x0d, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0xe2, 0x02, 0x19, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0f, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_operator_api_v1_plugin_service_proto_rawDescOnce sync.Once
	file_operator_api_v1_plugin_service_proto_rawDescData = file_operator_api_v1_plugin_service_proto_rawDesc
)

func file_operator_api_v1_plugin_service_proto_rawDescGZIP() []byte {
	file_operator_api_v1_plugin_service_proto_rawDescOnce.Do(func() {
		file_operator_api_v1_plugin_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_operator_api_v1_plugin_service_proto_rawDescData)
	})
	return file_operator_api_v1_plugin_service_proto_rawDescData
}

var file_operator_api_v1_plugin_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_operator_api_v1_plugin_service_proto_goTypes = []interface{}{
	(*InitializeRequest)(nil),      // 0: api.v1.plugin.InitializeRequest
	(*InitializeResponse)(nil),     // 1: api.v1.plugin.InitializeResponse
	(*RunCapsuleRequest)(nil),      // 2: api.v1.plugin.RunCapsuleRequest
	(*RunCapsuleResponse)(nil),     // 3: api.v1.plugin.RunCapsuleResponse
	(*GetObjectRequest)(nil),       // 4: api.v1.plugin.GetObjectRequest
	(*GetObjectResponse)(nil),      // 5: api.v1.plugin.GetObjectResponse
	(*SetObjectRequest)(nil),       // 6: api.v1.plugin.SetObjectRequest
	(*SetObjectResponse)(nil),      // 7: api.v1.plugin.SetObjectResponse
	(*GVK)(nil),                    // 8: api.v1.plugin.GVK
	(*DeleteObjectRequest)(nil),    // 9: api.v1.plugin.DeleteObjectRequest
	(*DeleteObjectResponse)(nil),   // 10: api.v1.plugin.DeleteObjectResponse
	(*MarkUsedObjectRequest)(nil),  // 11: api.v1.plugin.MarkUsedObjectRequest
	(*MarkUsedObjectResponse)(nil), // 12: api.v1.plugin.MarkUsedObjectResponse
}
var file_operator_api_v1_plugin_service_proto_depIdxs = []int32{
	8,  // 0: api.v1.plugin.GetObjectRequest.gvk:type_name -> api.v1.plugin.GVK
	8,  // 1: api.v1.plugin.SetObjectRequest.gvk:type_name -> api.v1.plugin.GVK
	8,  // 2: api.v1.plugin.DeleteObjectRequest.gvk:type_name -> api.v1.plugin.GVK
	8,  // 3: api.v1.plugin.MarkUsedObjectRequest.gvk:type_name -> api.v1.plugin.GVK
	0,  // 4: api.v1.plugin.PluginService.Initialize:input_type -> api.v1.plugin.InitializeRequest
	2,  // 5: api.v1.plugin.PluginService.RunCapsule:input_type -> api.v1.plugin.RunCapsuleRequest
	4,  // 6: api.v1.plugin.RequestService.GetObject:input_type -> api.v1.plugin.GetObjectRequest
	6,  // 7: api.v1.plugin.RequestService.SetObject:input_type -> api.v1.plugin.SetObjectRequest
	9,  // 8: api.v1.plugin.RequestService.DeleteObject:input_type -> api.v1.plugin.DeleteObjectRequest
	11, // 9: api.v1.plugin.RequestService.MarkUsedObject:input_type -> api.v1.plugin.MarkUsedObjectRequest
	1,  // 10: api.v1.plugin.PluginService.Initialize:output_type -> api.v1.plugin.InitializeResponse
	3,  // 11: api.v1.plugin.PluginService.RunCapsule:output_type -> api.v1.plugin.RunCapsuleResponse
	5,  // 12: api.v1.plugin.RequestService.GetObject:output_type -> api.v1.plugin.GetObjectResponse
	7,  // 13: api.v1.plugin.RequestService.SetObject:output_type -> api.v1.plugin.SetObjectResponse
	10, // 14: api.v1.plugin.RequestService.DeleteObject:output_type -> api.v1.plugin.DeleteObjectResponse
	12, // 15: api.v1.plugin.RequestService.MarkUsedObject:output_type -> api.v1.plugin.MarkUsedObjectResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_operator_api_v1_plugin_service_proto_init() }
func file_operator_api_v1_plugin_service_proto_init() {
	if File_operator_api_v1_plugin_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_operator_api_v1_plugin_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeRequest); i {
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
		file_operator_api_v1_plugin_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeResponse); i {
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
		file_operator_api_v1_plugin_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCapsuleRequest); i {
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
		file_operator_api_v1_plugin_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCapsuleResponse); i {
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
		file_operator_api_v1_plugin_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectRequest); i {
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
		file_operator_api_v1_plugin_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectResponse); i {
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
		file_operator_api_v1_plugin_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetObjectRequest); i {
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
		file_operator_api_v1_plugin_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetObjectResponse); i {
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
		file_operator_api_v1_plugin_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_operator_api_v1_plugin_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteObjectRequest); i {
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
		file_operator_api_v1_plugin_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteObjectResponse); i {
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
		file_operator_api_v1_plugin_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkUsedObjectRequest); i {
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
		file_operator_api_v1_plugin_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkUsedObjectResponse); i {
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
			RawDescriptor: file_operator_api_v1_plugin_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_operator_api_v1_plugin_service_proto_goTypes,
		DependencyIndexes: file_operator_api_v1_plugin_service_proto_depIdxs,
		MessageInfos:      file_operator_api_v1_plugin_service_proto_msgTypes,
	}.Build()
	File_operator_api_v1_plugin_service_proto = out.File
	file_operator_api_v1_plugin_service_proto_rawDesc = nil
	file_operator_api_v1_plugin_service_proto_goTypes = nil
	file_operator_api_v1_plugin_service_proto_depIdxs = nil
}
