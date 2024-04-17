// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: model/auth.proto

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

// The LoginType enum defines the type of login.
type LoginType int32

const (
	// Default value.
	LoginType_LOGIN_TYPE_UNSPECIFIED LoginType = 0
	// Email and password login.
	LoginType_LOGIN_TYPE_EMAIL_PASSWORD LoginType = 1
	// deprecated: text is not supported - Phone number and password login.
	LoginType_LOGIN_TYPE_PHONE_PASSWORD LoginType = 2
	// Username and password login.
	LoginType_LOGIN_TYPE_USERNAME_PASSWORD LoginType = 3
	// SSO Login
	LoginType_LOGIN_TYPE_SSO LoginType = 4
)

// Enum value maps for LoginType.
var (
	LoginType_name = map[int32]string{
		0: "LOGIN_TYPE_UNSPECIFIED",
		1: "LOGIN_TYPE_EMAIL_PASSWORD",
		2: "LOGIN_TYPE_PHONE_PASSWORD",
		3: "LOGIN_TYPE_USERNAME_PASSWORD",
		4: "LOGIN_TYPE_SSO",
	}
	LoginType_value = map[string]int32{
		"LOGIN_TYPE_UNSPECIFIED":       0,
		"LOGIN_TYPE_EMAIL_PASSWORD":    1,
		"LOGIN_TYPE_PHONE_PASSWORD":    2,
		"LOGIN_TYPE_USERNAME_PASSWORD": 3,
		"LOGIN_TYPE_SSO":               4,
	}
)

func (x LoginType) Enum() *LoginType {
	p := new(LoginType)
	*p = x
	return p
}

func (x LoginType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginType) Descriptor() protoreflect.EnumDescriptor {
	return file_model_auth_proto_enumTypes[0].Descriptor()
}

func (LoginType) Type() protoreflect.EnumType {
	return &file_model_auth_proto_enumTypes[0]
}

func (x LoginType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginType.Descriptor instead.
func (LoginType) EnumDescriptor() ([]byte, []int) {
	return file_model_auth_proto_rawDescGZIP(), []int{0}
}

var File_model_auth_proto protoreflect.FileDescriptor

var file_model_auth_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2a, 0x9b, 0x01, 0x0a, 0x09, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x4f, 0x47, 0x49, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44,
	0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10,
	0x02, 0x12, 0x20, 0x0a, 0x1c, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x53, 0x45, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52,
	0x44, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x53, 0x4f, 0x10, 0x04, 0x42, 0x6e, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x42, 0x09, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69,
	0x67, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0xca, 0x02, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x11, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_auth_proto_rawDescOnce sync.Once
	file_model_auth_proto_rawDescData = file_model_auth_proto_rawDesc
)

func file_model_auth_proto_rawDescGZIP() []byte {
	file_model_auth_proto_rawDescOnce.Do(func() {
		file_model_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_auth_proto_rawDescData)
	})
	return file_model_auth_proto_rawDescData
}

var file_model_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_auth_proto_goTypes = []interface{}{
	(LoginType)(0), // 0: model.LoginType
}
var file_model_auth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_model_auth_proto_init() }
func file_model_auth_proto_init() {
	if File_model_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_auth_proto_goTypes,
		DependencyIndexes: file_model_auth_proto_depIdxs,
		EnumInfos:         file_model_auth_proto_enumTypes,
	}.Build()
	File_model_auth_proto = out.File
	file_model_auth_proto_rawDesc = nil
	file_model_auth_proto_goTypes = nil
	file_model_auth_proto_depIdxs = nil
}
