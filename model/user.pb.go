// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: model/user.proto

package model

import (
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

type UserIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Identifier:
	//
	//	*UserIdentifier_Username
	//	*UserIdentifier_Email
	//	*UserIdentifier_PhoneNumber
	Identifier isUserIdentifier_Identifier `protobuf_oneof:"identifier"`
}

func (x *UserIdentifier) Reset() {
	*x = UserIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdentifier) ProtoMessage() {}

func (x *UserIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdentifier.ProtoReflect.Descriptor instead.
func (*UserIdentifier) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{0}
}

func (m *UserIdentifier) GetIdentifier() isUserIdentifier_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (x *UserIdentifier) GetUsername() string {
	if x, ok := x.GetIdentifier().(*UserIdentifier_Username); ok {
		return x.Username
	}
	return ""
}

func (x *UserIdentifier) GetEmail() string {
	if x, ok := x.GetIdentifier().(*UserIdentifier_Email); ok {
		return x.Email
	}
	return ""
}

func (x *UserIdentifier) GetPhoneNumber() string {
	if x, ok := x.GetIdentifier().(*UserIdentifier_PhoneNumber); ok {
		return x.PhoneNumber
	}
	return ""
}

type isUserIdentifier_Identifier interface {
	isUserIdentifier_Identifier()
}

type UserIdentifier_Username struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3,oneof"`
}

type UserIdentifier_Email struct {
	Email string `protobuf:"bytes,2,opt,name=email,proto3,oneof"`
}

type UserIdentifier_PhoneNumber struct {
	PhoneNumber string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3,oneof"`
}

func (*UserIdentifier_Username) isUserIdentifier_Identifier() {}

func (*UserIdentifier_Email) isUserIdentifier_Identifier() {}

func (*UserIdentifier_PhoneNumber) isUserIdentifier_Identifier() {}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username    string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	PhoneNumber string                 `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	GroupIds    []string               `protobuf:"bytes,6,rep,name=group_ids,json=groupIds,proto3" json:"group_ids,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfo) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserInfo) GetGroupIds() []string {
	if x != nil {
		return x.GroupIds
	}
	return nil
}

type UserEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PrintableName string                 `protobuf:"bytes,2,opt,name=printable_name,json=printableName,proto3" json:"printable_name,omitempty"`
	RegisterInfo  *RegisterInfo          `protobuf:"bytes,3,opt,name=register_info,json=registerInfo,proto3" json:"register_info,omitempty"`
	Verified      bool                   `protobuf:"varint,4,opt,name=verified,proto3" json:"verified,omitempty"`
	GroupIds      []string               `protobuf:"bytes,5,rep,name=group_ids,json=groupIds,proto3" json:"group_ids,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *UserEntry) Reset() {
	*x = UserEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEntry) ProtoMessage() {}

func (x *UserEntry) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEntry.ProtoReflect.Descriptor instead.
func (*UserEntry) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserEntry) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserEntry) GetPrintableName() string {
	if x != nil {
		return x.PrintableName
	}
	return ""
}

func (x *UserEntry) GetRegisterInfo() *RegisterInfo {
	if x != nil {
		return x.RegisterInfo
	}
	return nil
}

func (x *UserEntry) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *UserEntry) GetGroupIds() []string {
	if x != nil {
		return x.GroupIds
	}
	return nil
}

func (x *UserEntry) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type MemberEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     *UserEntry             `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	JoinedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=joined_at,json=joinedAt,proto3" json:"joined_at,omitempty"`
}

func (x *MemberEntry) Reset() {
	*x = MemberEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberEntry) ProtoMessage() {}

func (x *MemberEntry) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberEntry.ProtoReflect.Descriptor instead.
func (*MemberEntry) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{3}
}

func (x *MemberEntry) GetUser() *UserEntry {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *MemberEntry) GetJoinedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.JoinedAt
	}
	return nil
}

type RegisterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreaterId string          `protobuf:"bytes,1,opt,name=creater_id,json=createrId,proto3" json:"creater_id,omitempty"`
	Method    *RegisterMethod `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *RegisterInfo) Reset() {
	*x = RegisterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterInfo) ProtoMessage() {}

func (x *RegisterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterInfo.ProtoReflect.Descriptor instead.
func (*RegisterInfo) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterInfo) GetCreaterId() string {
	if x != nil {
		return x.CreaterId
	}
	return ""
}

func (x *RegisterInfo) GetMethod() *RegisterMethod {
	if x != nil {
		return x.Method
	}
	return nil
}

type RegisterMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Method:
	//
	//	*RegisterMethod_System_
	//	*RegisterMethod_Signup_
	//	*RegisterMethod_OauthProvider
	//	*RegisterMethod_Migration_
	Method isRegisterMethod_Method `protobuf_oneof:"method"`
}

func (x *RegisterMethod) Reset() {
	*x = RegisterMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterMethod) ProtoMessage() {}

func (x *RegisterMethod) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterMethod.ProtoReflect.Descriptor instead.
func (*RegisterMethod) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{5}
}

func (m *RegisterMethod) GetMethod() isRegisterMethod_Method {
	if m != nil {
		return m.Method
	}
	return nil
}

func (x *RegisterMethod) GetSystem() *RegisterMethod_System {
	if x, ok := x.GetMethod().(*RegisterMethod_System_); ok {
		return x.System
	}
	return nil
}

func (x *RegisterMethod) GetSignup() *RegisterMethod_Signup {
	if x, ok := x.GetMethod().(*RegisterMethod_Signup_); ok {
		return x.Signup
	}
	return nil
}

func (x *RegisterMethod) GetOauthProvider() OauthProvider {
	if x, ok := x.GetMethod().(*RegisterMethod_OauthProvider); ok {
		return x.OauthProvider
	}
	return OauthProvider_OAUTH_PROVIDER_UNSPECIFIED
}

func (x *RegisterMethod) GetMigration() *RegisterMethod_Migration {
	if x, ok := x.GetMethod().(*RegisterMethod_Migration_); ok {
		return x.Migration
	}
	return nil
}

type isRegisterMethod_Method interface {
	isRegisterMethod_Method()
}

type RegisterMethod_System_ struct {
	System *RegisterMethod_System `protobuf:"bytes,1,opt,name=system,proto3,oneof"`
}

type RegisterMethod_Signup_ struct {
	Signup *RegisterMethod_Signup `protobuf:"bytes,2,opt,name=signup,proto3,oneof"`
}

type RegisterMethod_OauthProvider struct {
	OauthProvider OauthProvider `protobuf:"varint,3,opt,name=oauth_provider,json=oauthProvider,proto3,enum=model.OauthProvider,oneof"`
}

type RegisterMethod_Migration_ struct {
	Migration *RegisterMethod_Migration `protobuf:"bytes,4,opt,name=migration,proto3,oneof"`
}

func (*RegisterMethod_System_) isRegisterMethod_Method() {}

func (*RegisterMethod_Signup_) isRegisterMethod_Method() {}

func (*RegisterMethod_OauthProvider) isRegisterMethod_Method() {}

func (*RegisterMethod_Migration_) isRegisterMethod_Method() {}

type RegisterMethod_System struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterMethod_System) Reset() {
	*x = RegisterMethod_System{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterMethod_System) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterMethod_System) ProtoMessage() {}

func (x *RegisterMethod_System) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterMethod_System.ProtoReflect.Descriptor instead.
func (*RegisterMethod_System) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{5, 0}
}

type RegisterMethod_Signup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginType LoginType `protobuf:"varint,1,opt,name=login_type,json=loginType,proto3,enum=model.LoginType" json:"login_type,omitempty"`
}

func (x *RegisterMethod_Signup) Reset() {
	*x = RegisterMethod_Signup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterMethod_Signup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterMethod_Signup) ProtoMessage() {}

func (x *RegisterMethod_Signup) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterMethod_Signup.ProtoReflect.Descriptor instead.
func (*RegisterMethod_Signup) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{5, 1}
}

func (x *RegisterMethod_Signup) GetLoginType() LoginType {
	if x != nil {
		return x.LoginType
	}
	return LoginType_LOGIN_TYPE_UNSPECIFIED
}

type RegisterMethod_Migration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *RegisterMethod_Migration) Reset() {
	*x = RegisterMethod_Migration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterMethod_Migration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterMethod_Migration) ProtoMessage() {}

func (x *RegisterMethod_Migration) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterMethod_Migration.ProtoReflect.Descriptor instead.
func (*RegisterMethod_Migration) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{5, 2}
}

func (x *RegisterMethod_Migration) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

var File_model_user_proto protoreflect.FileDescriptor

var file_model_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x0e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xb7, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x73, 0x22, 0xf9, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x38, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x6c, 0x0a,
	0x0b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5c, 0x0a, 0x0c, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0xf0, 0x02, 0x0a, 0x0e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x36, 0x0a, 0x06,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x48, 0x00, 0x52, 0x06, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x48, 0x00, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x3d, 0x0a, 0x0e,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x61, 0x75,
	0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0d, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x09, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x09, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x08, 0x0a, 0x06,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x39, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70,
	0x12, 0x2f, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x1a, 0x1f, 0x0a, 0x09, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x42, 0x08, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x6e, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58,
	0xaa, 0x02, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0xe2, 0x02, 0x11, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_user_proto_rawDescOnce sync.Once
	file_model_user_proto_rawDescData = file_model_user_proto_rawDesc
)

func file_model_user_proto_rawDescGZIP() []byte {
	file_model_user_proto_rawDescOnce.Do(func() {
		file_model_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_user_proto_rawDescData)
	})
	return file_model_user_proto_rawDescData
}

var file_model_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_model_user_proto_goTypes = []interface{}{
	(*UserIdentifier)(nil),           // 0: model.UserIdentifier
	(*UserInfo)(nil),                 // 1: model.UserInfo
	(*UserEntry)(nil),                // 2: model.UserEntry
	(*MemberEntry)(nil),              // 3: model.MemberEntry
	(*RegisterInfo)(nil),             // 4: model.RegisterInfo
	(*RegisterMethod)(nil),           // 5: model.RegisterMethod
	(*RegisterMethod_System)(nil),    // 6: model.RegisterMethod.System
	(*RegisterMethod_Signup)(nil),    // 7: model.RegisterMethod.Signup
	(*RegisterMethod_Migration)(nil), // 8: model.RegisterMethod.Migration
	(*timestamppb.Timestamp)(nil),    // 9: google.protobuf.Timestamp
	(OauthProvider)(0),               // 10: model.OauthProvider
	(LoginType)(0),                   // 11: model.LoginType
}
var file_model_user_proto_depIdxs = []int32{
	9,  // 0: model.UserInfo.created_at:type_name -> google.protobuf.Timestamp
	4,  // 1: model.UserEntry.register_info:type_name -> model.RegisterInfo
	9,  // 2: model.UserEntry.created_at:type_name -> google.protobuf.Timestamp
	2,  // 3: model.MemberEntry.user:type_name -> model.UserEntry
	9,  // 4: model.MemberEntry.joined_at:type_name -> google.protobuf.Timestamp
	5,  // 5: model.RegisterInfo.method:type_name -> model.RegisterMethod
	6,  // 6: model.RegisterMethod.system:type_name -> model.RegisterMethod.System
	7,  // 7: model.RegisterMethod.signup:type_name -> model.RegisterMethod.Signup
	10, // 8: model.RegisterMethod.oauth_provider:type_name -> model.OauthProvider
	8,  // 9: model.RegisterMethod.migration:type_name -> model.RegisterMethod.Migration
	11, // 10: model.RegisterMethod.Signup.login_type:type_name -> model.LoginType
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_model_user_proto_init() }
func file_model_user_proto_init() {
	if File_model_user_proto != nil {
		return
	}
	file_model_auth_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdentifier); i {
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
		file_model_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_model_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEntry); i {
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
		file_model_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberEntry); i {
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
		file_model_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterInfo); i {
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
		file_model_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterMethod); i {
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
		file_model_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterMethod_System); i {
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
		file_model_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterMethod_Signup); i {
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
		file_model_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterMethod_Migration); i {
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
	file_model_user_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UserIdentifier_Username)(nil),
		(*UserIdentifier_Email)(nil),
		(*UserIdentifier_PhoneNumber)(nil),
	}
	file_model_user_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*RegisterMethod_System_)(nil),
		(*RegisterMethod_Signup_)(nil),
		(*RegisterMethod_OauthProvider)(nil),
		(*RegisterMethod_Migration_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_user_proto_goTypes,
		DependencyIndexes: file_model_user_proto_depIdxs,
		MessageInfos:      file_model_user_proto_msgTypes,
	}.Build()
	File_model_user_proto = out.File
	file_model_user_proto_rawDesc = nil
	file_model_user_proto_goTypes = nil
	file_model_user_proto_depIdxs = nil
}
