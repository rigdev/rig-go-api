// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/v1/capsule/job.proto

package capsule

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// Different states a job execution can be in
type JobState int32

const (
	// Default value.
	JobState_JOB_STATE_UNSPECIFIED JobState = 0
	// The job is running.
	JobState_JOB_STATE_ONGOING JobState = 1
	// The job completed successfully.
	JobState_JOB_STATE_COMPLETED JobState = 2
	// The job failed.
	JobState_JOB_STATE_FAILED JobState = 3
	// The job was terminated.
	JobState_JOB_STATE_TERMINATED JobState = 4
)

// Enum value maps for JobState.
var (
	JobState_name = map[int32]string{
		0: "JOB_STATE_UNSPECIFIED",
		1: "JOB_STATE_ONGOING",
		2: "JOB_STATE_COMPLETED",
		3: "JOB_STATE_FAILED",
		4: "JOB_STATE_TERMINATED",
	}
	JobState_value = map[string]int32{
		"JOB_STATE_UNSPECIFIED": 0,
		"JOB_STATE_ONGOING":     1,
		"JOB_STATE_COMPLETED":   2,
		"JOB_STATE_FAILED":      3,
		"JOB_STATE_TERMINATED":  4,
	}
)

func (x JobState) Enum() *JobState {
	p := new(JobState)
	*p = x
	return p
}

func (x JobState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobState) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_capsule_job_proto_enumTypes[0].Descriptor()
}

func (JobState) Type() protoreflect.EnumType {
	return &file_api_v1_capsule_job_proto_enumTypes[0]
}

func (x JobState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobState.Descriptor instead.
func (JobState) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_capsule_job_proto_rawDescGZIP(), []int{0}
}

// An execution of a cron job.
type JobExecution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the job.
	JobName string `protobuf:"bytes,1,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	// When the job started running.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// When the job finished.
	FinishedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	// The state of the job.
	State JobState `protobuf:"varint,4,opt,name=state,proto3,enum=api.v1.capsule.JobState" json:"state,omitempty"`
	// Number of retries.
	Retries int32 `protobuf:"varint,5,opt,name=retries,proto3" json:"retries,omitempty"`
	// ID of the rollout.
	RolloutId uint64 `protobuf:"varint,6,opt,name=rollout_id,json=rolloutId,proto3" json:"rollout_id,omitempty"`
	// ID of the capsule.
	CapsuleId string `protobuf:"bytes,7,opt,name=capsule_id,json=capsuleId,proto3" json:"capsule_id,omitempty"`
	// ID of the project.
	ProjectId string `protobuf:"bytes,8,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// ID of the execution.
	ExecutionId string `protobuf:"bytes,9,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// ID of the environment.
	EnvironmentId string `protobuf:"bytes,10,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
}

func (x *JobExecution) Reset() {
	*x = JobExecution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobExecution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobExecution) ProtoMessage() {}

func (x *JobExecution) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobExecution.ProtoReflect.Descriptor instead.
func (*JobExecution) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_job_proto_rawDescGZIP(), []int{0}
}

func (x *JobExecution) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *JobExecution) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *JobExecution) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *JobExecution) GetState() JobState {
	if x != nil {
		return x.State
	}
	return JobState_JOB_STATE_UNSPECIFIED
}

func (x *JobExecution) GetRetries() int32 {
	if x != nil {
		return x.Retries
	}
	return 0
}

func (x *JobExecution) GetRolloutId() uint64 {
	if x != nil {
		return x.RolloutId
	}
	return 0
}

func (x *JobExecution) GetCapsuleId() string {
	if x != nil {
		return x.CapsuleId
	}
	return ""
}

func (x *JobExecution) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *JobExecution) GetExecutionId() string {
	if x != nil {
		return x.ExecutionId
	}
	return ""
}

func (x *JobExecution) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

// Specification for a cron job.
type CronJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the job.
	JobName string `protobuf:"bytes,1,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	// Cron schedule.
	Schedule string `protobuf:"bytes,2,opt,name=schedule,proto3" json:"schedule,omitempty"`
	// Maximum number of retries.
	MaxRetries int32 `protobuf:"varint,3,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// Maximum duration of the job.
	Timeout *durationpb.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The type of the job
	//
	// Types that are assignable to JobType:
	//
	//	*CronJob_Url
	//	*CronJob_Command
	JobType isCronJob_JobType `protobuf_oneof:"job_type"`
}

func (x *CronJob) Reset() {
	*x = CronJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CronJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronJob) ProtoMessage() {}

func (x *CronJob) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronJob.ProtoReflect.Descriptor instead.
func (*CronJob) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_job_proto_rawDescGZIP(), []int{1}
}

func (x *CronJob) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *CronJob) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

func (x *CronJob) GetMaxRetries() int32 {
	if x != nil {
		return x.MaxRetries
	}
	return 0
}

func (x *CronJob) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (m *CronJob) GetJobType() isCronJob_JobType {
	if m != nil {
		return m.JobType
	}
	return nil
}

func (x *CronJob) GetUrl() *JobURL {
	if x, ok := x.GetJobType().(*CronJob_Url); ok {
		return x.Url
	}
	return nil
}

func (x *CronJob) GetCommand() *JobCommand {
	if x, ok := x.GetJobType().(*CronJob_Command); ok {
		return x.Command
	}
	return nil
}

type isCronJob_JobType interface {
	isCronJob_JobType()
}

type CronJob_Url struct {
	// URL job.
	Url *JobURL `protobuf:"bytes,5,opt,name=url,proto3,oneof"`
}

type CronJob_Command struct {
	// Command job.
	Command *JobCommand `protobuf:"bytes,6,opt,name=command,proto3,oneof"`
}

func (*CronJob_Url) isCronJob_JobType() {}

func (*CronJob_Command) isCronJob_JobType() {}

// Run a job by making a HTTP request to a URL.
type JobURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Port to make the request to.
	Port uint64 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	// Path to make the request to.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Query parameters to add to the request.
	QueryParameters map[string]string `protobuf:"bytes,3,rep,name=query_parameters,json=queryParameters,proto3" json:"query_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *JobURL) Reset() {
	*x = JobURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobURL) ProtoMessage() {}

func (x *JobURL) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobURL.ProtoReflect.Descriptor instead.
func (*JobURL) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_job_proto_rawDescGZIP(), []int{2}
}

func (x *JobURL) GetPort() uint64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *JobURL) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *JobURL) GetQueryParameters() map[string]string {
	if x != nil {
		return x.QueryParameters
	}
	return nil
}

// Run a job by running a command in an instance of a capsule
type JobCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Command to run.
	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	// Arguments to pass to the command.
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *JobCommand) Reset() {
	*x = JobCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_capsule_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobCommand) ProtoMessage() {}

func (x *JobCommand) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_capsule_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobCommand.ProtoReflect.Descriptor instead.
func (*JobCommand) Descriptor() ([]byte, []int) {
	return file_api_v1_capsule_job_proto_rawDescGZIP(), []int{3}
}

func (x *JobCommand) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *JobCommand) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

var File_api_v1_capsule_job_proto protoreflect.FileDescriptor

var file_api_v1_capsule_job_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65,
	0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x03, 0x0a, 0x0c,
	0x4a, 0x6f, 0x62, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08,
	0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x2e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e,
	0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6c,
	0x6c, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72,
	0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x70, 0x73,
	0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61,
	0x70, 0x73, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x86, 0x02, 0x0a, 0x07, 0x43, 0x72, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x19, 0x0a, 0x08,
	0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x55, 0x52, 0x4c, 0x48, 0x00,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x0a, 0x0a,
	0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xcc, 0x01, 0x0a, 0x06, 0x4a, 0x6f,
	0x62, 0x55, 0x52, 0x4c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x56, 0x0a, 0x10,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x55, 0x52, 0x4c, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x1a, 0x42, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3a, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x72, 0x67, 0x73, 0x2a, 0x85, 0x01, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x0a, 0x15, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11,
	0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4e, 0x47, 0x4f, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10,
	0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x44, 0x10, 0x04, 0x42, 0xa5, 0x01, 0x0a,
	0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x70, 0x73,
	0x75, 0x6c, 0x65, 0x42, 0x08, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x67, 0x64,
	0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0xa2, 0x02, 0x03, 0x41,
	0x56, 0x43, 0xaa, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x73,
	0x75, 0x6c, 0x65, 0xca, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x61, 0x70,
	0x73, 0x75, 0x6c, 0x65, 0xe2, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x61,
	0x70, 0x73, 0x75, 0x6c, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x10, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x61, 0x70,
	0x73, 0x75, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_capsule_job_proto_rawDescOnce sync.Once
	file_api_v1_capsule_job_proto_rawDescData = file_api_v1_capsule_job_proto_rawDesc
)

func file_api_v1_capsule_job_proto_rawDescGZIP() []byte {
	file_api_v1_capsule_job_proto_rawDescOnce.Do(func() {
		file_api_v1_capsule_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_capsule_job_proto_rawDescData)
	})
	return file_api_v1_capsule_job_proto_rawDescData
}

var file_api_v1_capsule_job_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_capsule_job_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_capsule_job_proto_goTypes = []interface{}{
	(JobState)(0),                 // 0: api.v1.capsule.JobState
	(*JobExecution)(nil),          // 1: api.v1.capsule.JobExecution
	(*CronJob)(nil),               // 2: api.v1.capsule.CronJob
	(*JobURL)(nil),                // 3: api.v1.capsule.JobURL
	(*JobCommand)(nil),            // 4: api.v1.capsule.JobCommand
	nil,                           // 5: api.v1.capsule.JobURL.QueryParametersEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 7: google.protobuf.Duration
}
var file_api_v1_capsule_job_proto_depIdxs = []int32{
	6, // 0: api.v1.capsule.JobExecution.created_at:type_name -> google.protobuf.Timestamp
	6, // 1: api.v1.capsule.JobExecution.finished_at:type_name -> google.protobuf.Timestamp
	0, // 2: api.v1.capsule.JobExecution.state:type_name -> api.v1.capsule.JobState
	7, // 3: api.v1.capsule.CronJob.timeout:type_name -> google.protobuf.Duration
	3, // 4: api.v1.capsule.CronJob.url:type_name -> api.v1.capsule.JobURL
	4, // 5: api.v1.capsule.CronJob.command:type_name -> api.v1.capsule.JobCommand
	5, // 6: api.v1.capsule.JobURL.query_parameters:type_name -> api.v1.capsule.JobURL.QueryParametersEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_v1_capsule_job_proto_init() }
func file_api_v1_capsule_job_proto_init() {
	if File_api_v1_capsule_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_capsule_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobExecution); i {
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
		file_api_v1_capsule_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CronJob); i {
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
		file_api_v1_capsule_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobURL); i {
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
		file_api_v1_capsule_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobCommand); i {
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
	file_api_v1_capsule_job_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CronJob_Url)(nil),
		(*CronJob_Command)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_capsule_job_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_capsule_job_proto_goTypes,
		DependencyIndexes: file_api_v1_capsule_job_proto_depIdxs,
		EnumInfos:         file_api_v1_capsule_job_proto_enumTypes,
		MessageInfos:      file_api_v1_capsule_job_proto_msgTypes,
	}.Build()
	File_api_v1_capsule_job_proto = out.File
	file_api_v1_capsule_job_proto_rawDesc = nil
	file_api_v1_capsule_job_proto_goTypes = nil
	file_api_v1_capsule_job_proto_depIdxs = nil
}
