// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/v1/project/settings/service.proto

package settings

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

// Request to get the license information of the Rig installation.
type GetLicenseInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLicenseInfoRequest) Reset() {
	*x = GetLicenseInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_project_settings_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLicenseInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLicenseInfoRequest) ProtoMessage() {}

func (x *GetLicenseInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_project_settings_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLicenseInfoRequest.ProtoReflect.Descriptor instead.
func (*GetLicenseInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_project_settings_service_proto_rawDescGZIP(), []int{0}
}

// Response for getting the license information of the Rig installation.
type GetLicenseInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLicenseInfoResponse) Reset() {
	*x = GetLicenseInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_project_settings_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLicenseInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLicenseInfoResponse) ProtoMessage() {}

func (x *GetLicenseInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_project_settings_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLicenseInfoResponse.ProtoReflect.Descriptor instead.
func (*GetLicenseInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_project_settings_service_proto_rawDescGZIP(), []int{1}
}

// Empty get settings request
type GetSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSettingsRequest) Reset() {
	*x = GetSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_project_settings_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingsRequest) ProtoMessage() {}

func (x *GetSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_project_settings_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetSettingsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_project_settings_service_proto_rawDescGZIP(), []int{2}
}

// Response for getting settings for the project.
type GetSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The settings.
	Settings *Settings `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *GetSettingsResponse) Reset() {
	*x = GetSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_project_settings_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingsResponse) ProtoMessage() {}

func (x *GetSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_project_settings_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingsResponse.ProtoReflect.Descriptor instead.
func (*GetSettingsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_project_settings_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetSettingsResponse) GetSettings() *Settings {
	if x != nil {
		return x.Settings
	}
	return nil
}

// Request for  updating settings for a project.
type UpdateSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The updates to apply.
	Updates []*Update `protobuf:"bytes,1,rep,name=updates,proto3" json:"updates,omitempty"`
}

func (x *UpdateSettingsRequest) Reset() {
	*x = UpdateSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_project_settings_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSettingsRequest) ProtoMessage() {}

func (x *UpdateSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_project_settings_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSettingsRequest.ProtoReflect.Descriptor instead.
func (*UpdateSettingsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_project_settings_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSettingsRequest) GetUpdates() []*Update {
	if x != nil {
		return x.Updates
	}
	return nil
}

// Empty response for updating a project's settings.
type UpdateSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSettingsResponse) Reset() {
	*x = UpdateSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_project_settings_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSettingsResponse) ProtoMessage() {}

func (x *UpdateSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_project_settings_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSettingsResponse.ProtoReflect.Descriptor instead.
func (*UpdateSettingsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_project_settings_service_proto_rawDescGZIP(), []int{5}
}

var File_api_v1_project_settings_service_proto protoreflect.FileDescriptor

var file_api_v1_project_settings_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x1a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x54, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x52, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x39, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xdf, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x73, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x73, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xe1, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67,
	0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0xa2, 0x02, 0x04,
	0x41, 0x56, 0x50, 0x53, 0xaa, 0x02, 0x17, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0xca, 0x02,
	0x17, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69, 0x5c, 0x56,
	0x31, 0x5c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x1a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x3a, 0x3a, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_v1_project_settings_service_proto_rawDescOnce sync.Once
	file_api_v1_project_settings_service_proto_rawDescData = file_api_v1_project_settings_service_proto_rawDesc
)

func file_api_v1_project_settings_service_proto_rawDescGZIP() []byte {
	file_api_v1_project_settings_service_proto_rawDescOnce.Do(func() {
		file_api_v1_project_settings_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_project_settings_service_proto_rawDescData)
	})
	return file_api_v1_project_settings_service_proto_rawDescData
}

var file_api_v1_project_settings_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_v1_project_settings_service_proto_goTypes = []interface{}{
	(*GetLicenseInfoRequest)(nil),  // 0: api.v1.project.settings.GetLicenseInfoRequest
	(*GetLicenseInfoResponse)(nil), // 1: api.v1.project.settings.GetLicenseInfoResponse
	(*GetSettingsRequest)(nil),     // 2: api.v1.project.settings.GetSettingsRequest
	(*GetSettingsResponse)(nil),    // 3: api.v1.project.settings.GetSettingsResponse
	(*UpdateSettingsRequest)(nil),  // 4: api.v1.project.settings.UpdateSettingsRequest
	(*UpdateSettingsResponse)(nil), // 5: api.v1.project.settings.UpdateSettingsResponse
	(*Settings)(nil),               // 6: api.v1.project.settings.Settings
	(*Update)(nil),                 // 7: api.v1.project.settings.Update
}
var file_api_v1_project_settings_service_proto_depIdxs = []int32{
	6, // 0: api.v1.project.settings.GetSettingsResponse.settings:type_name -> api.v1.project.settings.Settings
	7, // 1: api.v1.project.settings.UpdateSettingsRequest.updates:type_name -> api.v1.project.settings.Update
	2, // 2: api.v1.project.settings.Service.GetSettings:input_type -> api.v1.project.settings.GetSettingsRequest
	4, // 3: api.v1.project.settings.Service.UpdateSettings:input_type -> api.v1.project.settings.UpdateSettingsRequest
	0, // 4: api.v1.project.settings.Service.GetLicenseInfo:input_type -> api.v1.project.settings.GetLicenseInfoRequest
	3, // 5: api.v1.project.settings.Service.GetSettings:output_type -> api.v1.project.settings.GetSettingsResponse
	5, // 6: api.v1.project.settings.Service.UpdateSettings:output_type -> api.v1.project.settings.UpdateSettingsResponse
	1, // 7: api.v1.project.settings.Service.GetLicenseInfo:output_type -> api.v1.project.settings.GetLicenseInfoResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_project_settings_service_proto_init() }
func file_api_v1_project_settings_service_proto_init() {
	if File_api_v1_project_settings_service_proto != nil {
		return
	}
	file_api_v1_project_settings_settings_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_project_settings_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLicenseInfoRequest); i {
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
		file_api_v1_project_settings_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLicenseInfoResponse); i {
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
		file_api_v1_project_settings_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingsRequest); i {
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
		file_api_v1_project_settings_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingsResponse); i {
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
		file_api_v1_project_settings_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSettingsRequest); i {
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
		file_api_v1_project_settings_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSettingsResponse); i {
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
			RawDescriptor: file_api_v1_project_settings_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_project_settings_service_proto_goTypes,
		DependencyIndexes: file_api_v1_project_settings_service_proto_depIdxs,
		MessageInfos:      file_api_v1_project_settings_service_proto_msgTypes,
	}.Build()
	File_api_v1_project_settings_service_proto = out.File
	file_api_v1_project_settings_service_proto_rawDesc = nil
	file_api_v1_project_settings_service_proto_goTypes = nil
	file_api_v1_project_settings_service_proto_depIdxs = nil
}
