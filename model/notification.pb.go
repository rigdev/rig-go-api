// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: model/notification.proto

package model

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

type NotificationTopic int32

const (
	NotificationTopic_NOTIFICATION_TOPIC_UNSPECIFIED NotificationTopic = 0
	NotificationTopic_NOTIFICATION_TOPIC_ROLLOUT     NotificationTopic = 1
	NotificationTopic_NOTIFICATION_TOPIC_ISSUE       NotificationTopic = 2
)

// Enum value maps for NotificationTopic.
var (
	NotificationTopic_name = map[int32]string{
		0: "NOTIFICATION_TOPIC_UNSPECIFIED",
		1: "NOTIFICATION_TOPIC_ROLLOUT",
		2: "NOTIFICATION_TOPIC_ISSUE",
	}
	NotificationTopic_value = map[string]int32{
		"NOTIFICATION_TOPIC_UNSPECIFIED": 0,
		"NOTIFICATION_TOPIC_ROLLOUT":     1,
		"NOTIFICATION_TOPIC_ISSUE":       2,
	}
)

func (x NotificationTopic) Enum() *NotificationTopic {
	p := new(NotificationTopic)
	*p = x
	return p
}

func (x NotificationTopic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationTopic) Descriptor() protoreflect.EnumDescriptor {
	return file_model_notification_proto_enumTypes[0].Descriptor()
}

func (NotificationTopic) Type() protoreflect.EnumType {
	return &file_model_notification_proto_enumTypes[0]
}

func (x NotificationTopic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationTopic.Descriptor instead.
func (NotificationTopic) EnumDescriptor() ([]byte, []int) {
	return file_model_notification_proto_rawDescGZIP(), []int{0}
}

type NotificationNotifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target       *NotificationTarget `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Topics       []NotificationTopic `protobuf:"varint,2,rep,packed,name=topics,proto3,enum=model.NotificationTopic" json:"topics,omitempty"`
	Environments *EnvironmentFilter  `protobuf:"bytes,3,opt,name=environments,proto3" json:"environments,omitempty"`
}

func (x *NotificationNotifier) Reset() {
	*x = NotificationNotifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationNotifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationNotifier) ProtoMessage() {}

func (x *NotificationNotifier) ProtoReflect() protoreflect.Message {
	mi := &file_model_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationNotifier.ProtoReflect.Descriptor instead.
func (*NotificationNotifier) Descriptor() ([]byte, []int) {
	return file_model_notification_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationNotifier) GetTarget() *NotificationTarget {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *NotificationNotifier) GetTopics() []NotificationTopic {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *NotificationNotifier) GetEnvironments() *EnvironmentFilter {
	if x != nil {
		return x.Environments
	}
	return nil
}

type NotificationTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Target:
	//
	//	*NotificationTarget_Slack
	//	*NotificationTarget_Email
	Target isNotificationTarget_Target `protobuf_oneof:"target"`
}

func (x *NotificationTarget) Reset() {
	*x = NotificationTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationTarget) ProtoMessage() {}

func (x *NotificationTarget) ProtoReflect() protoreflect.Message {
	mi := &file_model_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationTarget.ProtoReflect.Descriptor instead.
func (*NotificationTarget) Descriptor() ([]byte, []int) {
	return file_model_notification_proto_rawDescGZIP(), []int{1}
}

func (m *NotificationTarget) GetTarget() isNotificationTarget_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *NotificationTarget) GetSlack() *NotificationTarget_SlackTarget {
	if x, ok := x.GetTarget().(*NotificationTarget_Slack); ok {
		return x.Slack
	}
	return nil
}

func (x *NotificationTarget) GetEmail() *NotificationTarget_EmailTarget {
	if x, ok := x.GetTarget().(*NotificationTarget_Email); ok {
		return x.Email
	}
	return nil
}

type isNotificationTarget_Target interface {
	isNotificationTarget_Target()
}

type NotificationTarget_Slack struct {
	Slack *NotificationTarget_SlackTarget `protobuf:"bytes,1,opt,name=slack,proto3,oneof"`
}

type NotificationTarget_Email struct {
	Email *NotificationTarget_EmailTarget `protobuf:"bytes,2,opt,name=email,proto3,oneof"`
}

func (*NotificationTarget_Slack) isNotificationTarget_Target() {}

func (*NotificationTarget_Email) isNotificationTarget_Target() {}

type NotificationTarget_SlackTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace string `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *NotificationTarget_SlackTarget) Reset() {
	*x = NotificationTarget_SlackTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationTarget_SlackTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationTarget_SlackTarget) ProtoMessage() {}

func (x *NotificationTarget_SlackTarget) ProtoReflect() protoreflect.Message {
	mi := &file_model_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationTarget_SlackTarget.ProtoReflect.Descriptor instead.
func (*NotificationTarget_SlackTarget) Descriptor() ([]byte, []int) {
	return file_model_notification_proto_rawDescGZIP(), []int{1, 0}
}

func (x *NotificationTarget_SlackTarget) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

func (x *NotificationTarget_SlackTarget) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

type NotificationTarget_EmailTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FromEmail string   `protobuf:"bytes,2,opt,name=from_email,json=fromEmail,proto3" json:"from_email,omitempty"`
	ToEmails  []string `protobuf:"bytes,3,rep,name=to_emails,json=toEmails,proto3" json:"to_emails,omitempty"`
}

func (x *NotificationTarget_EmailTarget) Reset() {
	*x = NotificationTarget_EmailTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationTarget_EmailTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationTarget_EmailTarget) ProtoMessage() {}

func (x *NotificationTarget_EmailTarget) ProtoReflect() protoreflect.Message {
	mi := &file_model_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationTarget_EmailTarget.ProtoReflect.Descriptor instead.
func (*NotificationTarget_EmailTarget) Descriptor() ([]byte, []int) {
	return file_model_notification_proto_rawDescGZIP(), []int{1, 1}
}

func (x *NotificationTarget_EmailTarget) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NotificationTarget_EmailTarget) GetFromEmail() string {
	if x != nil {
		return x.FromEmail
	}
	return ""
}

func (x *NotificationTarget_EmailTarget) GetToEmails() []string {
	if x != nil {
		return x.ToEmails
	}
	return nil
}

var File_model_notification_proto protoreflect.FileDescriptor

var file_model_notification_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x14, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xc3, 0x02, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x3d, 0x0a,
	0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x3d, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x4a, 0x0a, 0x0b, 0x53,
	0x6c, 0x61, 0x63, 0x6b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x1a, 0x59, 0x0a, 0x0b, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6d,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2a, 0x75, 0x0a, 0x11,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x22, 0x0a, 0x1e, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x5f, 0x52, 0x4f, 0x4c, 0x4c,
	0x4f, 0x55, 0x54, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x5f, 0x49, 0x53, 0x53, 0x55,
	0x45, 0x10, 0x02, 0x42, 0x76, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x42, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa,
	0x02, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2,
	0x02, 0x11, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_model_notification_proto_rawDescOnce sync.Once
	file_model_notification_proto_rawDescData = file_model_notification_proto_rawDesc
)

func file_model_notification_proto_rawDescGZIP() []byte {
	file_model_notification_proto_rawDescOnce.Do(func() {
		file_model_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_notification_proto_rawDescData)
	})
	return file_model_notification_proto_rawDescData
}

var file_model_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_model_notification_proto_goTypes = []interface{}{
	(NotificationTopic)(0),                 // 0: model.NotificationTopic
	(*NotificationNotifier)(nil),           // 1: model.NotificationNotifier
	(*NotificationTarget)(nil),             // 2: model.NotificationTarget
	(*NotificationTarget_SlackTarget)(nil), // 3: model.NotificationTarget.SlackTarget
	(*NotificationTarget_EmailTarget)(nil), // 4: model.NotificationTarget.EmailTarget
	(*EnvironmentFilter)(nil),              // 5: model.EnvironmentFilter
}
var file_model_notification_proto_depIdxs = []int32{
	2, // 0: model.NotificationNotifier.target:type_name -> model.NotificationTarget
	0, // 1: model.NotificationNotifier.topics:type_name -> model.NotificationTopic
	5, // 2: model.NotificationNotifier.environments:type_name -> model.EnvironmentFilter
	3, // 3: model.NotificationTarget.slack:type_name -> model.NotificationTarget.SlackTarget
	4, // 4: model.NotificationTarget.email:type_name -> model.NotificationTarget.EmailTarget
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_model_notification_proto_init() }
func file_model_notification_proto_init() {
	if File_model_notification_proto != nil {
		return
	}
	file_model_environment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationNotifier); i {
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
		file_model_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationTarget); i {
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
		file_model_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationTarget_SlackTarget); i {
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
		file_model_notification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationTarget_EmailTarget); i {
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
	file_model_notification_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*NotificationTarget_Slack)(nil),
		(*NotificationTarget_Email)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_notification_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_notification_proto_goTypes,
		DependencyIndexes: file_model_notification_proto_depIdxs,
		EnumInfos:         file_model_notification_proto_enumTypes,
		MessageInfos:      file_model_notification_proto_msgTypes,
	}.Build()
	File_model_notification_proto = out.File
	file_model_notification_proto_rawDesc = nil
	file_model_notification_proto_goTypes = nil
	file_model_notification_proto_depIdxs = nil
}
