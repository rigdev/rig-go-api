// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/v1/capsule/status.proto

package capsule

import (
	rollout "github.com/rigdev/rig-go-api/api/v1/capsule/rollout"
	pipeline "github.com/rigdev/rig-go-api/operator/api/v1/pipeline"
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

type Transition int32

const (
	Transition_TRANSITION_UNSPECIFIED   Transition = 0
	Transition_TRANSITION_BEING_CREATED Transition = 1
	Transition_TRANSITION_UP_TO_DATE    Transition = 2
	Transition_TRANSITION_BEING_DELETED Transition = 3
)

// Enum value maps for Transition.
var (
	Transition_name = map[int32]string{
		0: "TRANSITION_UNSPECIFIED",
		1: "TRANSITION_BEING_CREATED",
		2: "TRANSITION_UP_TO_DATE",
		3: "TRANSITION_BEING_DELETED",
	}
	Transition_value = map[string]int32{
		"TRANSITION_UNSPECIFIED":   0,
		"TRANSITION_BEING_CREATED": 1,
		"TRANSITION_UP_TO_DATE":    2,
		"TRANSITION_BEING_DELETED": 3,
	}
)

func (x Transition) Enum() *Transition {
	p := new(Transition)
	*p = x
	return p
}

func (x Transition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Transition) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_capsule_status_proto_enumTypes[0].Descriptor()
}

func (Transition) Type() protoreflect.EnumType {
	return &file_api_v1_capsule_status_proto_enumTypes[0]
}

func (x Transition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Transition.Descriptor instead.
func (Transition) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_capsule_status_proto_rawDescGZIP(), []int{0}
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace       string              `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Capsule         *CapsuleStatus      `protobuf:"bytes,2,opt,name=capsule,proto3" json:"capsule,omitempty"`
	Rollout         *rollout.Status     `protobuf:"bytes,3,opt,name=rollout,proto3" json:"rollout,omitempty"`
	ContainerConfig *ContainerConfig    `protobuf:"bytes,4,opt,name=container_config,json=containerConfig,proto3" json:"container_config,omitempty"`
	Instances       *InstancesStatus    `protobuf:"bytes,5,opt,name=instances,proto3" json:"instances,omitempty"`
	Interfaces      []*InterfaceStatus  `protobuf:"bytes,6,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	ConfigFiles     []*ConfigFileStatus `protobuf:"bytes,7,rep,name=config_files,json=configFiles,proto3" json:"config_files,omitempty"`
	CronJobs        []*CronJobStatus    `protobuf:"bytes,8,rep,name=cron_jobs,json=cronJobs,proto3" json:"cron_jobs,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_status_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Status) GetCapsule() *CapsuleStatus {
	if x != nil {
		return x.Capsule
	}
	return nil
}

func (x *Status) GetRollout() *rollout.Status {
	if x != nil {
		return x.Rollout
	}
	return nil
}

func (x *Status) GetContainerConfig() *ContainerConfig {
	if x != nil {
		return x.ContainerConfig
	}
	return nil
}

func (x *Status) GetInstances() *InstancesStatus {
	if x != nil {
		return x.Instances
	}
	return nil
}

func (x *Status) GetInterfaces() []*InterfaceStatus {
	if x != nil {
		return x.Interfaces
	}
	return nil
}

func (x *Status) GetConfigFiles() []*ConfigFileStatus {
	if x != nil {
		return x.ConfigFiles
	}
	return nil
}

func (x *Status) GetCronJobs() []*CronJobStatus {
	if x != nil {
		return x.CronJobs
	}
	return nil
}

type CapsuleStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statuses []*pipeline.ObjectStatus `protobuf:"bytes,1,rep,name=statuses,proto3" json:"statuses,omitempty"`
}

func (x *CapsuleStatus) Reset() {
	*x = CapsuleStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapsuleStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapsuleStatus) ProtoMessage() {}

func (x *CapsuleStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapsuleStatus.ProtoReflect.Descriptor instead.
func (*CapsuleStatus) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_status_proto_rawDescGZIP(), []int{1}
}

func (x *CapsuleStatus) GetStatuses() []*pipeline.ObjectStatus {
	if x != nil {
		return x.Statuses
	}
	return nil
}

type InstancesStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of updated instances.
	NumUpdated uint32 `protobuf:"varint,1,opt,name=num_updated,json=numUpdated,proto3" json:"num_updated,omitempty"`
	// The number of ready instances.
	NumReady uint32 `protobuf:"varint,2,opt,name=num_ready,json=numReady,proto3" json:"num_ready,omitempty"`
	// The number of stuck instances.
	NumStuck uint32 `protobuf:"varint,3,opt,name=num_stuck,json=numStuck,proto3" json:"num_stuck,omitempty"`
	// The number of instances with the wrong version.
	NumWrongVersion uint32 `protobuf:"varint,4,opt,name=num_wrong_version,json=numWrongVersion,proto3" json:"num_wrong_version,omitempty"`
}

func (x *InstancesStatus) Reset() {
	*x = InstancesStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstancesStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstancesStatus) ProtoMessage() {}

func (x *InstancesStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstancesStatus.ProtoReflect.Descriptor instead.
func (*InstancesStatus) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_status_proto_rawDescGZIP(), []int{2}
}

func (x *InstancesStatus) GetNumUpdated() uint32 {
	if x != nil {
		return x.NumUpdated
	}
	return 0
}

func (x *InstancesStatus) GetNumReady() uint32 {
	if x != nil {
		return x.NumReady
	}
	return 0
}

func (x *InstancesStatus) GetNumStuck() uint32 {
	if x != nil {
		return x.NumStuck
	}
	return 0
}

func (x *InstancesStatus) GetNumWrongVersion() uint32 {
	if x != nil {
		return x.NumWrongVersion
	}
	return 0
}

type ContainerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image                string            `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Command              string            `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Args                 []string          `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	EnvironmentVariables map[string]string `protobuf:"bytes,4,rep,name=environment_variables,json=environmentVariables,proto3" json:"environment_variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Scale                *HorizontalScale  `protobuf:"bytes,5,opt,name=scale,proto3" json:"scale,omitempty"`
}

func (x *ContainerConfig) Reset() {
	*x = ContainerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_status_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerConfig) ProtoMessage() {}

func (x *ContainerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_status_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerConfig.ProtoReflect.Descriptor instead.
func (*ContainerConfig) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_status_proto_rawDescGZIP(), []int{3}
}

func (x *ContainerConfig) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ContainerConfig) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *ContainerConfig) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *ContainerConfig) GetEnvironmentVariables() map[string]string {
	if x != nil {
		return x.EnvironmentVariables
	}
	return nil
}

func (x *ContainerConfig) GetScale() *HorizontalScale {
	if x != nil {
		return x.Scale
	}
	return nil
}

type InterfaceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port       uint32                   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Routes     []*InterfaceStatus_Route `protobuf:"bytes,3,rep,name=routes,proto3" json:"routes,omitempty"`
	Status     []*pipeline.ObjectStatus `protobuf:"bytes,4,rep,name=status,proto3" json:"status,omitempty"`
	Transition Transition               `protobuf:"varint,5,opt,name=transition,proto3,enum=api.v1.capsule.Transition" json:"transition,omitempty"`
}

func (x *InterfaceStatus) Reset() {
	*x = InterfaceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_status_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceStatus) ProtoMessage() {}

func (x *InterfaceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_status_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfaceStatus.ProtoReflect.Descriptor instead.
func (*InterfaceStatus) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_status_proto_rawDescGZIP(), []int{4}
}

func (x *InterfaceStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InterfaceStatus) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *InterfaceStatus) GetRoutes() []*InterfaceStatus_Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *InterfaceStatus) GetStatus() []*pipeline.ObjectStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *InterfaceStatus) GetTransition() Transition {
	if x != nil {
		return x.Transition
	}
	return Transition_TRANSITION_UNSPECIFIED
}

type ConfigFileStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path       string                   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	IsSecret   bool                     `protobuf:"varint,2,opt,name=isSecret,proto3" json:"isSecret,omitempty"`
	Status     []*pipeline.ObjectStatus `protobuf:"bytes,3,rep,name=status,proto3" json:"status,omitempty"`
	Transition Transition               `protobuf:"varint,4,opt,name=transition,proto3,enum=api.v1.capsule.Transition" json:"transition,omitempty"`
}

func (x *ConfigFileStatus) Reset() {
	*x = ConfigFileStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_status_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigFileStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigFileStatus) ProtoMessage() {}

func (x *ConfigFileStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_status_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigFileStatus.ProtoReflect.Descriptor instead.
func (*ConfigFileStatus) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_status_proto_rawDescGZIP(), []int{5}
}

func (x *ConfigFileStatus) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ConfigFileStatus) GetIsSecret() bool {
	if x != nil {
		return x.IsSecret
	}
	return false
}

func (x *ConfigFileStatus) GetStatus() []*pipeline.ObjectStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ConfigFileStatus) GetTransition() Transition {
	if x != nil {
		return x.Transition
	}
	return Transition_TRANSITION_UNSPECIFIED
}

type CronJobStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobName       string        `protobuf:"bytes,1,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	Schedule      string        `protobuf:"bytes,2,opt,name=schedule,proto3" json:"schedule,omitempty"`
	LastExecution *JobExecution `protobuf:"bytes,3,opt,name=last_execution,json=lastExecution,proto3" json:"last_execution,omitempty"`
	Transition    Transition    `protobuf:"varint,4,opt,name=transition,proto3,enum=api.v1.capsule.Transition" json:"transition,omitempty"`
}

func (x *CronJobStatus) Reset() {
	*x = CronJobStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_status_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CronJobStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronJobStatus) ProtoMessage() {}

func (x *CronJobStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_status_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronJobStatus.ProtoReflect.Descriptor instead.
func (*CronJobStatus) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_status_proto_rawDescGZIP(), []int{6}
}

func (x *CronJobStatus) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *CronJobStatus) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

func (x *CronJobStatus) GetLastExecution() *JobExecution {
	if x != nil {
		return x.LastExecution
	}
	return nil
}

func (x *CronJobStatus) GetTransition() Transition {
	if x != nil {
		return x.Transition
	}
	return Transition_TRANSITION_UNSPECIFIED
}

type InterfaceStatus_Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route      *HostRoute               `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	Status     []*pipeline.ObjectStatus `protobuf:"bytes,2,rep,name=status,proto3" json:"status,omitempty"`
	Transition Transition               `protobuf:"varint,3,opt,name=transition,proto3,enum=api.v1.capsule.Transition" json:"transition,omitempty"`
}

func (x *InterfaceStatus_Route) Reset() {
	*x = InterfaceStatus_Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_status_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceStatus_Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceStatus_Route) ProtoMessage() {}

func (x *InterfaceStatus_Route) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_status_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfaceStatus_Route.ProtoReflect.Descriptor instead.
func (*InterfaceStatus_Route) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_status_proto_rawDescGZIP(), []int{4, 0}
}

func (x *InterfaceStatus_Route) GetRoute() *HostRoute {
	if x != nil {
		return x.Route
	}
	return nil
}

func (x *InterfaceStatus_Route) GetStatus() []*pipeline.ObjectStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *InterfaceStatus_Route) GetTransition() Transition {
	if x != nil {
		return x.Transition
	}
	return Transition_TRANSITION_UNSPECIFIED
}

var File_api_v1_capsule_status_proto protoreflect.FileDescriptor

var file_api_v1_capsule_status_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x1a, 0x23, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2f, 0x72, 0x6f,
	0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75,
	0x6c, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2f, 0x6a, 0x6f,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x03, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x37, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75,
	0x6c, 0x65, 0x2e, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x07, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x72, 0x6f, 0x6c,
	0x6c, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x6c, 0x6c,
	0x6f, 0x75, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x6c,
	0x6f, 0x75, 0x74, 0x12, 0x4a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x3d, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73,
	0x75, 0x6c, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x3f,
	0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73,
	0x75, 0x6c, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12,
	0x43, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x6a, 0x6f, 0x62,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x6f, 0x6e, 0x4a, 0x6f, 0x62,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x08, 0x63, 0x72, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x73,
	0x22, 0x4a, 0x0a, 0x0d, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x39, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0x98, 0x01, 0x0a,
	0x0f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x74, 0x75, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x53, 0x74, 0x75, 0x63, 0x6b, 0x12, 0x2a, 0x0a, 0x11, 0x6e,
	0x75, 0x6d, 0x5f, 0x77, 0x72, 0x6f, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x57, 0x72, 0x6f, 0x6e, 0x67,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xc5, 0x02, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12,
	0x6e, 0x0a, 0x15, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12,
	0x35, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e,
	0x48, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x52,
	0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x1a, 0x47, 0x0a, 0x19, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x99, 0x03, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xab, 0x01,
	0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x3a, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70,
	0x73, 0x75, 0x6c, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb5, 0x01, 0x0a, 0x10,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xc7, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x0e,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61,
	0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3a, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x7f, 0x0a,
	0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x45, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x5f, 0x54, 0x4f, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02,
	0x12, 0x1c, 0x0a, 0x18, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42,
	0x45, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0xa8,
	0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61,
	0x70, 0x73, 0x75, 0x6c, 0x65, 0x42, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c,
	0x65, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x43, 0xaa, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31,
	0x2e, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0xca, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x5c, 0x56,
	0x31, 0x5c, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0xe2, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x5c,
	0x56, 0x31, 0x5c, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31,
	0x3a, 0x3a, 0x43, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v1_capsule_status_proto_rawDescOnce sync.Once
	file_api_v1_capsule_status_proto_rawDescData = file_api_v1_capsule_status_proto_rawDesc
)

func file_api_v1_capsule_status_proto_rawDescGZIP() []byte {
	file_api_v1_capsule_status_proto_rawDescOnce.Do(func() {
		file_api_v1_capsule_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_capsule_status_proto_rawDescData)
	})
	return file_api_v1_capsule_status_proto_rawDescData
}

var file_api_v1_capsule_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_capsule_status_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_v1_capsule_status_proto_goTypes = []interface{}{
	(Transition)(0),               // 0: api.v1.capsule.Transition
	(*Status)(nil),                // 1: api.v1.capsule.Status
	(*CapsuleStatus)(nil),         // 2: api.v1.capsule.CapsuleStatus
	(*InstancesStatus)(nil),       // 3: api.v1.capsule.InstancesStatus
	(*ContainerConfig)(nil),       // 4: api.v1.capsule.ContainerConfig
	(*InterfaceStatus)(nil),       // 5: api.v1.capsule.InterfaceStatus
	(*ConfigFileStatus)(nil),      // 6: api.v1.capsule.ConfigFileStatus
	(*CronJobStatus)(nil),         // 7: api.v1.capsule.CronJobStatus
	nil,                           // 8: api.v1.capsule.ContainerConfig.EnvironmentVariablesEntry
	(*InterfaceStatus_Route)(nil), // 9: api.v1.capsule.InterfaceStatus.Route
	(*rollout.Status)(nil),        // 10: api.v1.capsule.rollout.Status
	(*pipeline.ObjectStatus)(nil), // 11: api.v1.pipeline.ObjectStatus
	(*HorizontalScale)(nil),       // 12: api.v1.capsule.HorizontalScale
	(*JobExecution)(nil),          // 13: api.v1.capsule.JobExecution
	(*HostRoute)(nil),             // 14: api.v1.capsule.HostRoute
}
var file_api_v1_capsule_status_proto_depIdxs = []int32{
	2,  // 0: api.v1.capsule.Status.capsule:type_name -> api.v1.capsule.CapsuleStatus
	10, // 1: api.v1.capsule.Status.rollout:type_name -> api.v1.capsule.rollout.Status
	4,  // 2: api.v1.capsule.Status.container_config:type_name -> api.v1.capsule.ContainerConfig
	3,  // 3: api.v1.capsule.Status.instances:type_name -> api.v1.capsule.InstancesStatus
	5,  // 4: api.v1.capsule.Status.interfaces:type_name -> api.v1.capsule.InterfaceStatus
	6,  // 5: api.v1.capsule.Status.config_files:type_name -> api.v1.capsule.ConfigFileStatus
	7,  // 6: api.v1.capsule.Status.cron_jobs:type_name -> api.v1.capsule.CronJobStatus
	11, // 7: api.v1.capsule.CapsuleStatus.statuses:type_name -> api.v1.pipeline.ObjectStatus
	8,  // 8: api.v1.capsule.ContainerConfig.environment_variables:type_name -> api.v1.capsule.ContainerConfig.EnvironmentVariablesEntry
	12, // 9: api.v1.capsule.ContainerConfig.scale:type_name -> api.v1.capsule.HorizontalScale
	9,  // 10: api.v1.capsule.InterfaceStatus.routes:type_name -> api.v1.capsule.InterfaceStatus.Route
	11, // 11: api.v1.capsule.InterfaceStatus.status:type_name -> api.v1.pipeline.ObjectStatus
	0,  // 12: api.v1.capsule.InterfaceStatus.transition:type_name -> api.v1.capsule.Transition
	11, // 13: api.v1.capsule.ConfigFileStatus.status:type_name -> api.v1.pipeline.ObjectStatus
	0,  // 14: api.v1.capsule.ConfigFileStatus.transition:type_name -> api.v1.capsule.Transition
	13, // 15: api.v1.capsule.CronJobStatus.last_execution:type_name -> api.v1.capsule.JobExecution
	0,  // 16: api.v1.capsule.CronJobStatus.transition:type_name -> api.v1.capsule.Transition
	14, // 17: api.v1.capsule.InterfaceStatus.Route.route:type_name -> api.v1.capsule.HostRoute
	11, // 18: api.v1.capsule.InterfaceStatus.Route.status:type_name -> api.v1.pipeline.ObjectStatus
	0,  // 19: api.v1.capsule.InterfaceStatus.Route.transition:type_name -> api.v1.capsule.Transition
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_api_v1_capsule_status_proto_init() }
func file_api_v1_capsule_status_proto_init() {
	if File_api_v1_capsule_status_proto != nil {
		return
	}
	file_api_v1_capsule_change_proto_init()
	file_api_v1_capsule_job_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_capsule_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_api_v1_capsule_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapsuleStatus); i {
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
		file_api_v1_capsule_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstancesStatus); i {
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
		file_api_v1_capsule_status_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerConfig); i {
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
		file_api_v1_capsule_status_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfaceStatus); i {
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
		file_api_v1_capsule_status_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigFileStatus); i {
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
		file_api_v1_capsule_status_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CronJobStatus); i {
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
		file_api_v1_capsule_status_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfaceStatus_Route); i {
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
			RawDescriptor: file_api_v1_capsule_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_capsule_status_proto_goTypes,
		DependencyIndexes: file_api_v1_capsule_status_proto_depIdxs,
		EnumInfos:         file_api_v1_capsule_status_proto_enumTypes,
		MessageInfos:      file_api_v1_capsule_status_proto_msgTypes,
	}.Build()
	File_api_v1_capsule_status_proto = out.File
	file_api_v1_capsule_status_proto_rawDesc = nil
	file_api_v1_capsule_status_proto_goTypes = nil
	file_api_v1_capsule_status_proto_depIdxs = nil
}