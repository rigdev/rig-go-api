// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/v1/settings/settings.proto

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
	return file_api_v1_settings_settings_proto_enumTypes[0].Descriptor()
}

func (NotificationTopic) Type() protoreflect.EnumType {
	return &file_api_v1_settings_settings_proto_enumTypes[0]
}

func (x NotificationTopic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationTopic.Descriptor instead.
func (NotificationTopic) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_settings_settings_proto_rawDescGZIP(), []int{0}
}

// The plan for a rig installation
type Plan int32

const (
	// Unspecified / unactivated plan.
	Plan_PLAN_UNSPECIFIED Plan = 0
	// Free tier.
	Plan_PLAN_FREE Plan = 1
	// Team / Pro tier.
	Plan_PLAN_TEAM Plan = 2
	// Enterprise tier.
	Plan_PLAN_ENTERPRISE Plan = 3
)

// Enum value maps for Plan.
var (
	Plan_name = map[int32]string{
		0: "PLAN_UNSPECIFIED",
		1: "PLAN_FREE",
		2: "PLAN_TEAM",
		3: "PLAN_ENTERPRISE",
	}
	Plan_value = map[string]int32{
		"PLAN_UNSPECIFIED": 0,
		"PLAN_FREE":        1,
		"PLAN_TEAM":        2,
		"PLAN_ENTERPRISE":  3,
	}
)

func (x Plan) Enum() *Plan {
	p := new(Plan)
	*p = x
	return p
}

func (x Plan) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Plan) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_settings_settings_proto_enumTypes[1].Descriptor()
}

func (Plan) Type() protoreflect.EnumType {
	return &file_api_v1_settings_settings_proto_enumTypes[1]
}

func (x Plan) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Plan.Descriptor instead.
func (Plan) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_settings_settings_proto_rawDescGZIP(), []int{1}
}

// Platform wide settings.
type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationNotifiers []*NotificationNotifier `protobuf:"bytes,1,rep,name=notification_notifiers,json=notificationNotifiers,proto3" json:"notification_notifiers,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_settings_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_settings_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_api_v1_settings_settings_proto_rawDescGZIP(), []int{0}
}

func (x *Settings) GetNotificationNotifiers() []*NotificationNotifier {
	if x != nil {
		return x.NotificationNotifiers
	}
	return nil
}

// Update message for platform settings.
type Update struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Field:
	//
	//	*Update_SetNotificationNotifiers_
	Field isUpdate_Field `protobuf_oneof:"field"`
}

func (x *Update) Reset() {
	*x = Update{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_settings_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Update) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Update) ProtoMessage() {}

func (x *Update) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_settings_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Update.ProtoReflect.Descriptor instead.
func (*Update) Descriptor() ([]byte, []int) {
	return file_api_v1_settings_settings_proto_rawDescGZIP(), []int{1}
}

func (m *Update) GetField() isUpdate_Field {
	if m != nil {
		return m.Field
	}
	return nil
}

func (x *Update) GetSetNotificationNotifiers() *Update_SetNotificationNotifiers {
	if x, ok := x.GetField().(*Update_SetNotificationNotifiers_); ok {
		return x.SetNotificationNotifiers
	}
	return nil
}

type isUpdate_Field interface {
	isUpdate_Field()
}

type Update_SetNotificationNotifiers_ struct {
	SetNotificationNotifiers *Update_SetNotificationNotifiers `protobuf:"bytes,1,opt,name=set_notification_notifiers,json=setNotificationNotifiers,proto3,oneof"`
}

func (*Update_SetNotificationNotifiers_) isUpdate_Field() {}

type NotificationNotifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target       *NotificationTarget       `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Topics       []NotificationTopic       `protobuf:"varint,2,rep,packed,name=topics,proto3,enum=api.v1.settings.NotificationTopic" json:"topics,omitempty"`
	Environments *NotificationEnvironments `protobuf:"bytes,3,opt,name=environments,proto3" json:"environments,omitempty"`
}

func (x *NotificationNotifier) Reset() {
	*x = NotificationNotifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_settings_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationNotifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationNotifier) ProtoMessage() {}

func (x *NotificationNotifier) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_settings_settings_proto_msgTypes[2]
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
	return file_api_v1_settings_settings_proto_rawDescGZIP(), []int{2}
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

func (x *NotificationNotifier) GetEnvironments() *NotificationEnvironments {
	if x != nil {
		return x.Environments
	}
	return nil
}

type NotificationEnvironments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Filter:
	//
	//	*NotificationEnvironments_All_
	//	*NotificationEnvironments_Selected_
	Filter isNotificationEnvironments_Filter `protobuf_oneof:"filter"`
}

func (x *NotificationEnvironments) Reset() {
	*x = NotificationEnvironments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_settings_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationEnvironments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationEnvironments) ProtoMessage() {}

func (x *NotificationEnvironments) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_settings_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationEnvironments.ProtoReflect.Descriptor instead.
func (*NotificationEnvironments) Descriptor() ([]byte, []int) {
	return file_api_v1_settings_settings_proto_rawDescGZIP(), []int{3}
}

func (m *NotificationEnvironments) GetFilter() isNotificationEnvironments_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (x *NotificationEnvironments) GetAll() *NotificationEnvironments_All {
	if x, ok := x.GetFilter().(*NotificationEnvironments_All_); ok {
		return x.All
	}
	return nil
}

func (x *NotificationEnvironments) GetSelected() *NotificationEnvironments_Selected {
	if x, ok := x.GetFilter().(*NotificationEnvironments_Selected_); ok {
		return x.Selected
	}
	return nil
}

type isNotificationEnvironments_Filter interface {
	isNotificationEnvironments_Filter()
}

type NotificationEnvironments_All_ struct {
	All *NotificationEnvironments_All `protobuf:"bytes,1,opt,name=all,proto3,oneof"`
}

type NotificationEnvironments_Selected_ struct {
	Selected *NotificationEnvironments_Selected `protobuf:"bytes,2,opt,name=selected,proto3,oneof"`
}

func (*NotificationEnvironments_All_) isNotificationEnvironments_Filter() {}

func (*NotificationEnvironments_Selected_) isNotificationEnvironments_Filter() {}

type NotificationTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Target:
	//
	//	*NotificationTarget_Slack
	Target isNotificationTarget_Target `protobuf_oneof:"target"`
}

func (x *NotificationTarget) Reset() {
	*x = NotificationTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_settings_settings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationTarget) ProtoMessage() {}

func (x *NotificationTarget) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_settings_settings_proto_msgTypes[4]
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
	return file_api_v1_settings_settings_proto_rawDescGZIP(), []int{4}
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

type isNotificationTarget_Target interface {
	isNotificationTarget_Target()
}

type NotificationTarget_Slack struct {
	Slack *NotificationTarget_SlackTarget `protobuf:"bytes,1,opt,name=slack,proto3,oneof"`
}

func (*NotificationTarget_Slack) isNotificationTarget_Target() {}

type Update_SetNotificationNotifiers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notifiers []*NotificationNotifier `protobuf:"bytes,1,rep,name=notifiers,proto3" json:"notifiers,omitempty"`
}

func (x *Update_SetNotificationNotifiers) Reset() {
	*x = Update_SetNotificationNotifiers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_settings_settings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Update_SetNotificationNotifiers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Update_SetNotificationNotifiers) ProtoMessage() {}

func (x *Update_SetNotificationNotifiers) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_settings_settings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Update_SetNotificationNotifiers.ProtoReflect.Descriptor instead.
func (*Update_SetNotificationNotifiers) Descriptor() ([]byte, []int) {
	return file_api_v1_settings_settings_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Update_SetNotificationNotifiers) GetNotifiers() []*NotificationNotifier {
	if x != nil {
		return x.Notifiers
	}
	return nil
}

type NotificationEnvironments_All struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncludeEphemeral bool `protobuf:"varint,1,opt,name=include_ephemeral,json=includeEphemeral,proto3" json:"include_ephemeral,omitempty"`
}

func (x *NotificationEnvironments_All) Reset() {
	*x = NotificationEnvironments_All{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_settings_settings_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationEnvironments_All) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationEnvironments_All) ProtoMessage() {}

func (x *NotificationEnvironments_All) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_settings_settings_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationEnvironments_All.ProtoReflect.Descriptor instead.
func (*NotificationEnvironments_All) Descriptor() ([]byte, []int) {
	return file_api_v1_settings_settings_proto_rawDescGZIP(), []int{3, 0}
}

func (x *NotificationEnvironments_All) GetIncludeEphemeral() bool {
	if x != nil {
		return x.IncludeEphemeral
	}
	return false
}

type NotificationEnvironments_Selected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvironmentIds []string `protobuf:"bytes,1,rep,name=environment_ids,json=environmentIds,proto3" json:"environment_ids,omitempty"`
}

func (x *NotificationEnvironments_Selected) Reset() {
	*x = NotificationEnvironments_Selected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_settings_settings_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationEnvironments_Selected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationEnvironments_Selected) ProtoMessage() {}

func (x *NotificationEnvironments_Selected) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_settings_settings_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationEnvironments_Selected.ProtoReflect.Descriptor instead.
func (*NotificationEnvironments_Selected) Descriptor() ([]byte, []int) {
	return file_api_v1_settings_settings_proto_rawDescGZIP(), []int{3, 1}
}

func (x *NotificationEnvironments_Selected) GetEnvironmentIds() []string {
	if x != nil {
		return x.EnvironmentIds
	}
	return nil
}

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
		mi := &file_api_v1_settings_settings_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationTarget_SlackTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationTarget_SlackTarget) ProtoMessage() {}

func (x *NotificationTarget_SlackTarget) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_settings_settings_proto_msgTypes[8]
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
	return file_api_v1_settings_settings_proto_rawDescGZIP(), []int{4, 0}
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

var File_api_v1_settings_settings_proto protoreflect.FileDescriptor

var file_api_v1_settings_settings_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x22, 0x68, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x5c, 0x0a,
	0x16, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x15, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x22, 0xe4, 0x01, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x70, 0x0a, 0x1a, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x18,
	0x73, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x5f, 0x0a, 0x18, 0x53, 0x65, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x73, 0x12, 0x43, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x09,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x22, 0xde, 0x01, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x06, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x12, 0x4d, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0xa2, 0x02, 0x0a, 0x18, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x41, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x6c, 0x6c, 0x48, 0x00, 0x52, 0x03,
	0x61, 0x6c, 0x6c, 0x12, 0x50, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x1a, 0x32, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x12, 0x2b, 0x0a, 0x11,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x45, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x1a, 0x33, 0x0a, 0x08, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x42, 0x08,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xb3, 0x01, 0x0a, 0x12, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x47, 0x0a, 0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48,
	0x00, 0x52, 0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x1a, 0x4a, 0x0a, 0x0b, 0x53, 0x6c, 0x61, 0x63,
	0x6b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2a, 0x75,
	0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x22, 0x0a, 0x1e, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x4e, 0x4f, 0x54, 0x49, 0x46,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x5f, 0x52, 0x4f,
	0x4c, 0x4c, 0x4f, 0x55, 0x54, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x4e, 0x4f, 0x54, 0x49, 0x46,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x5f, 0x49, 0x53,
	0x53, 0x55, 0x45, 0x10, 0x02, 0x2a, 0x4f, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x14, 0x0a,
	0x10, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x46, 0x52, 0x45, 0x45,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x10,
	0x02, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x50,
	0x52, 0x49, 0x53, 0x45, 0x10, 0x03, 0x42, 0xb0, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x0d,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x67, 0x64,
	0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0xa2, 0x02, 0x03,
	0x41, 0x56, 0x53, 0xaa, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0xca, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0xe2, 0x02, 0x1b, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31,
	0x5c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a,
	0x3a, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v1_settings_settings_proto_rawDescOnce sync.Once
	file_api_v1_settings_settings_proto_rawDescData = file_api_v1_settings_settings_proto_rawDesc
)

func file_api_v1_settings_settings_proto_rawDescGZIP() []byte {
	file_api_v1_settings_settings_proto_rawDescOnce.Do(func() {
		file_api_v1_settings_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_settings_settings_proto_rawDescData)
	})
	return file_api_v1_settings_settings_proto_rawDescData
}

var file_api_v1_settings_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_v1_settings_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_v1_settings_settings_proto_goTypes = []interface{}{
	(NotificationTopic)(0),                    // 0: api.v1.settings.NotificationTopic
	(Plan)(0),                                 // 1: api.v1.settings.Plan
	(*Settings)(nil),                          // 2: api.v1.settings.Settings
	(*Update)(nil),                            // 3: api.v1.settings.Update
	(*NotificationNotifier)(nil),              // 4: api.v1.settings.NotificationNotifier
	(*NotificationEnvironments)(nil),          // 5: api.v1.settings.NotificationEnvironments
	(*NotificationTarget)(nil),                // 6: api.v1.settings.NotificationTarget
	(*Update_SetNotificationNotifiers)(nil),   // 7: api.v1.settings.Update.SetNotificationNotifiers
	(*NotificationEnvironments_All)(nil),      // 8: api.v1.settings.NotificationEnvironments.All
	(*NotificationEnvironments_Selected)(nil), // 9: api.v1.settings.NotificationEnvironments.Selected
	(*NotificationTarget_SlackTarget)(nil),    // 10: api.v1.settings.NotificationTarget.SlackTarget
}
var file_api_v1_settings_settings_proto_depIdxs = []int32{
	4,  // 0: api.v1.settings.Settings.notification_notifiers:type_name -> api.v1.settings.NotificationNotifier
	7,  // 1: api.v1.settings.Update.set_notification_notifiers:type_name -> api.v1.settings.Update.SetNotificationNotifiers
	6,  // 2: api.v1.settings.NotificationNotifier.target:type_name -> api.v1.settings.NotificationTarget
	0,  // 3: api.v1.settings.NotificationNotifier.topics:type_name -> api.v1.settings.NotificationTopic
	5,  // 4: api.v1.settings.NotificationNotifier.environments:type_name -> api.v1.settings.NotificationEnvironments
	8,  // 5: api.v1.settings.NotificationEnvironments.all:type_name -> api.v1.settings.NotificationEnvironments.All
	9,  // 6: api.v1.settings.NotificationEnvironments.selected:type_name -> api.v1.settings.NotificationEnvironments.Selected
	10, // 7: api.v1.settings.NotificationTarget.slack:type_name -> api.v1.settings.NotificationTarget.SlackTarget
	4,  // 8: api.v1.settings.Update.SetNotificationNotifiers.notifiers:type_name -> api.v1.settings.NotificationNotifier
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_v1_settings_settings_proto_init() }
func file_api_v1_settings_settings_proto_init() {
	if File_api_v1_settings_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_settings_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
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
		file_api_v1_settings_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Update); i {
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
		file_api_v1_settings_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_v1_settings_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationEnvironments); i {
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
		file_api_v1_settings_settings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_v1_settings_settings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Update_SetNotificationNotifiers); i {
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
		file_api_v1_settings_settings_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationEnvironments_All); i {
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
		file_api_v1_settings_settings_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationEnvironments_Selected); i {
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
		file_api_v1_settings_settings_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_api_v1_settings_settings_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Update_SetNotificationNotifiers_)(nil),
	}
	file_api_v1_settings_settings_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*NotificationEnvironments_All_)(nil),
		(*NotificationEnvironments_Selected_)(nil),
	}
	file_api_v1_settings_settings_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*NotificationTarget_Slack)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_settings_settings_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_settings_settings_proto_goTypes,
		DependencyIndexes: file_api_v1_settings_settings_proto_depIdxs,
		EnumInfos:         file_api_v1_settings_settings_proto_enumTypes,
		MessageInfos:      file_api_v1_settings_settings_proto_msgTypes,
	}.Build()
	File_api_v1_settings_settings_proto = out.File
	file_api_v1_settings_settings_proto_rawDesc = nil
	file_api_v1_settings_settings_proto_goTypes = nil
	file_api_v1_settings_settings_proto_depIdxs = nil
}
