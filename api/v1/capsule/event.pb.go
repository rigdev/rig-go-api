// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/v1/capsule/event.proto

package capsule

import (
	model "github.com/rigdev/rig-go-api/model"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An event is a message from a rollout
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Potential author associated with the event.
	CreatedBy *model.Author `protobuf:"bytes,1,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// When the event was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// The rollout that created the event.
	RolloutId uint64 `protobuf:"varint,3,opt,name=rollout_id,json=rolloutId,proto3" json:"rollout_id,omitempty"`
	// A message associated with the event.
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// The data associated with the event.
	EventData *EventData `protobuf:"bytes,5,opt,name=event_data,json=eventData,proto3" json:"event_data,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetCreatedBy() *model.Author {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *Event) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Event) GetRolloutId() uint64 {
	if x != nil {
		return x.RolloutId
	}
	return 0
}

func (x *Event) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Event) GetEventData() *EventData {
	if x != nil {
		return x.EventData
	}
	return nil
}

// An event that is associated with a rollout.
type RolloutEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RolloutEvent) Reset() {
	*x = RolloutEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RolloutEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RolloutEvent) ProtoMessage() {}

func (x *RolloutEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RolloutEvent.ProtoReflect.Descriptor instead.
func (*RolloutEvent) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_event_proto_rawDescGZIP(), []int{1}
}

// An event that is associated with an abort.
type AbortEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AbortEvent) Reset() {
	*x = AbortEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbortEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbortEvent) ProtoMessage() {}

func (x *AbortEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbortEvent.ProtoReflect.Descriptor instead.
func (*AbortEvent) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_event_proto_rawDescGZIP(), []int{2}
}

// An event that is associated with an error.
type ErrorEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ErrorEvent) Reset() {
	*x = ErrorEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorEvent) ProtoMessage() {}

func (x *ErrorEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorEvent.ProtoReflect.Descriptor instead.
func (*ErrorEvent) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_event_proto_rawDescGZIP(), []int{3}
}

// The data associated with an event.
type EventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*EventData_Rollout
	//	*EventData_Error
	//	*EventData_Abort
	Kind isEventData_Kind `protobuf_oneof:"kind"`
}

func (x *EventData) Reset() {
	*x = EventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventData) ProtoMessage() {}

func (x *EventData) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventData.ProtoReflect.Descriptor instead.
func (*EventData) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_event_proto_rawDescGZIP(), []int{4}
}

func (m *EventData) GetKind() isEventData_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *EventData) GetRollout() *RolloutEvent {
	if x, ok := x.GetKind().(*EventData_Rollout); ok {
		return x.Rollout
	}
	return nil
}

func (x *EventData) GetError() *ErrorEvent {
	if x, ok := x.GetKind().(*EventData_Error); ok {
		return x.Error
	}
	return nil
}

func (x *EventData) GetAbort() *AbortEvent {
	if x, ok := x.GetKind().(*EventData_Abort); ok {
		return x.Abort
	}
	return nil
}

type isEventData_Kind interface {
	isEventData_Kind()
}

type EventData_Rollout struct {
	// If event is a rollout.
	Rollout *RolloutEvent `protobuf:"bytes,1,opt,name=rollout,proto3,oneof"`
}

type EventData_Error struct {
	// if event is an error event.
	Error *ErrorEvent `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type EventData_Abort struct {
	// If event is an abort event.
	Abort *AbortEvent `protobuf:"bytes,3,opt,name=abort,proto3,oneof"`
}

func (*EventData_Rollout) isEventData_Kind() {}

func (*EventData_Error) isEventData_Kind() {}

func (*EventData_Abort) isEventData_Kind() {}

var File_api_v1_capsule_event_proto protoreflect.FileDescriptor

var file_api_v1_capsule_event_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe3, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75,
	0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75,
	0x6c, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x0e, 0x0a, 0x0c, 0x52, 0x6f, 0x6c, 0x6c, 0x6f,
	0x75, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x41, 0x62, 0x6f, 0x72, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0xb5, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x38, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73,
	0x75, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x32, 0x0a, 0x05, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e,
	0x41, 0x62, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x61, 0x62,
	0x6f, 0x72, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0xa7, 0x01, 0x0a, 0x12,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75,
	0x6c, 0x65, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x67,
	0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0xa2, 0x02, 0x03,
	0x41, 0x56, 0x43, 0xaa, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x61, 0x70,
	0x73, 0x75, 0x6c, 0x65, 0xca, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x61,
	0x70, 0x73, 0x75, 0x6c, 0x65, 0xe2, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x43,
	0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x10, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x61,
	0x70, 0x73, 0x75, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_capsule_event_proto_rawDescOnce sync.Once
	file_api_v1_capsule_event_proto_rawDescData = file_api_v1_capsule_event_proto_rawDesc
)

func file_api_v1_capsule_event_proto_rawDescGZIP() []byte {
	file_api_v1_capsule_event_proto_rawDescOnce.Do(func() {
		file_api_v1_capsule_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_capsule_event_proto_rawDescData)
	})
	return file_api_v1_capsule_event_proto_rawDescData
}

var file_api_v1_capsule_event_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_capsule_event_proto_goTypes = []interface{}{
	(*Event)(nil),                 // 0: api.v1.capsule.Event
	(*RolloutEvent)(nil),          // 1: api.v1.capsule.RolloutEvent
	(*AbortEvent)(nil),            // 2: api.v1.capsule.AbortEvent
	(*ErrorEvent)(nil),            // 3: api.v1.capsule.ErrorEvent
	(*EventData)(nil),             // 4: api.v1.capsule.EventData
	(*model.Author)(nil),          // 5: model.Author
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_api_v1_capsule_event_proto_depIdxs = []int32{
	5, // 0: api.v1.capsule.Event.created_by:type_name -> model.Author
	6, // 1: api.v1.capsule.Event.created_at:type_name -> google.protobuf.Timestamp
	4, // 2: api.v1.capsule.Event.event_data:type_name -> api.v1.capsule.EventData
	1, // 3: api.v1.capsule.EventData.rollout:type_name -> api.v1.capsule.RolloutEvent
	3, // 4: api.v1.capsule.EventData.error:type_name -> api.v1.capsule.ErrorEvent
	2, // 5: api.v1.capsule.EventData.abort:type_name -> api.v1.capsule.AbortEvent
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_v1_capsule_event_proto_init() }
func file_api_v1_capsule_event_proto_init() {
	if File_api_v1_capsule_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_capsule_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_api_v1_capsule_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RolloutEvent); i {
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
		file_api_v1_capsule_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbortEvent); i {
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
		file_api_v1_capsule_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorEvent); i {
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
		file_api_v1_capsule_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventData); i {
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
	file_api_v1_capsule_event_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*EventData_Rollout)(nil),
		(*EventData_Error)(nil),
		(*EventData_Abort)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_capsule_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_capsule_event_proto_goTypes,
		DependencyIndexes: file_api_v1_capsule_event_proto_depIdxs,
		MessageInfos:      file_api_v1_capsule_event_proto_msgTypes,
	}.Build()
	File_api_v1_capsule_event_proto = out.File
	file_api_v1_capsule_event_proto_rawDesc = nil
	file_api_v1_capsule_event_proto_goTypes = nil
	file_api_v1_capsule_event_proto_depIdxs = nil
}
