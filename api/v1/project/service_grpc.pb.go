// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package project

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// Create project.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Delete project.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Get project.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Get project list.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Update updates the profile of the project.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Get public key.
	PublicKey(ctx context.Context, in *PublicKeyRequest, opts ...grpc.CallOption) (*PublicKeyResponse, error)
	// Returns all objects of a given kind.
	GetObjectsByKind(ctx context.Context, in *GetObjectsByKindRequest, opts ...grpc.CallOption) (*GetObjectsByKindResponse, error)
	// Returns all metrics of a given custom object.
	GetCustomObjectMetrics(ctx context.Context, in *GetCustomObjectMetricsRequest, opts ...grpc.CallOption) (*GetCustomObjectMetricsResponse, error)
	GetEffectiveGitSettings(ctx context.Context, in *GetEffectiveGitSettingsRequest, opts ...grpc.CallOption) (*GetEffectiveGitSettingsResponse, error)
	GetEffectivePipelineSettings(ctx context.Context, in *GetEffectivePipelineSettingsRequest, opts ...grpc.CallOption) (*GetEffectivePipelineSettingsResponse, error)
	GetEffectiveNotificationSettings(ctx context.Context, in *GetEffectiveNotificationSettingsRequest, opts ...grpc.CallOption) (*GetEffectiveNotificationSettingsResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/api.v1.project.Service/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/api.v1.project.Service/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/api.v1.project.Service/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/api.v1.project.Service/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/api.v1.project.Service/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PublicKey(ctx context.Context, in *PublicKeyRequest, opts ...grpc.CallOption) (*PublicKeyResponse, error) {
	out := new(PublicKeyResponse)
	err := c.cc.Invoke(ctx, "/api.v1.project.Service/PublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetObjectsByKind(ctx context.Context, in *GetObjectsByKindRequest, opts ...grpc.CallOption) (*GetObjectsByKindResponse, error) {
	out := new(GetObjectsByKindResponse)
	err := c.cc.Invoke(ctx, "/api.v1.project.Service/GetObjectsByKind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetCustomObjectMetrics(ctx context.Context, in *GetCustomObjectMetricsRequest, opts ...grpc.CallOption) (*GetCustomObjectMetricsResponse, error) {
	out := new(GetCustomObjectMetricsResponse)
	err := c.cc.Invoke(ctx, "/api.v1.project.Service/GetCustomObjectMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetEffectiveGitSettings(ctx context.Context, in *GetEffectiveGitSettingsRequest, opts ...grpc.CallOption) (*GetEffectiveGitSettingsResponse, error) {
	out := new(GetEffectiveGitSettingsResponse)
	err := c.cc.Invoke(ctx, "/api.v1.project.Service/GetEffectiveGitSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetEffectivePipelineSettings(ctx context.Context, in *GetEffectivePipelineSettingsRequest, opts ...grpc.CallOption) (*GetEffectivePipelineSettingsResponse, error) {
	out := new(GetEffectivePipelineSettingsResponse)
	err := c.cc.Invoke(ctx, "/api.v1.project.Service/GetEffectivePipelineSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetEffectiveNotificationSettings(ctx context.Context, in *GetEffectiveNotificationSettingsRequest, opts ...grpc.CallOption) (*GetEffectiveNotificationSettingsResponse, error) {
	out := new(GetEffectiveNotificationSettingsResponse)
	err := c.cc.Invoke(ctx, "/api.v1.project.Service/GetEffectiveNotificationSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// Create project.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Delete project.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Get project.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Get project list.
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Update updates the profile of the project.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Get public key.
	PublicKey(context.Context, *PublicKeyRequest) (*PublicKeyResponse, error)
	// Returns all objects of a given kind.
	GetObjectsByKind(context.Context, *GetObjectsByKindRequest) (*GetObjectsByKindResponse, error)
	// Returns all metrics of a given custom object.
	GetCustomObjectMetrics(context.Context, *GetCustomObjectMetricsRequest) (*GetCustomObjectMetricsResponse, error)
	GetEffectiveGitSettings(context.Context, *GetEffectiveGitSettingsRequest) (*GetEffectiveGitSettingsResponse, error)
	GetEffectivePipelineSettings(context.Context, *GetEffectivePipelineSettingsRequest) (*GetEffectivePipelineSettingsResponse, error)
	GetEffectiveNotificationSettings(context.Context, *GetEffectiveNotificationSettingsRequest) (*GetEffectiveNotificationSettingsResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedServiceServer) PublicKey(context.Context, *PublicKeyRequest) (*PublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublicKey not implemented")
}
func (UnimplementedServiceServer) GetObjectsByKind(context.Context, *GetObjectsByKindRequest) (*GetObjectsByKindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectsByKind not implemented")
}
func (UnimplementedServiceServer) GetCustomObjectMetrics(context.Context, *GetCustomObjectMetricsRequest) (*GetCustomObjectMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomObjectMetrics not implemented")
}
func (UnimplementedServiceServer) GetEffectiveGitSettings(context.Context, *GetEffectiveGitSettingsRequest) (*GetEffectiveGitSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEffectiveGitSettings not implemented")
}
func (UnimplementedServiceServer) GetEffectivePipelineSettings(context.Context, *GetEffectivePipelineSettingsRequest) (*GetEffectivePipelineSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEffectivePipelineSettings not implemented")
}
func (UnimplementedServiceServer) GetEffectiveNotificationSettings(context.Context, *GetEffectiveNotificationSettingsRequest) (*GetEffectiveNotificationSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEffectiveNotificationSettings not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.project.Service/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.project.Service/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.project.Service/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.project.Service/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.project.Service/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_PublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).PublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.project.Service/PublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).PublicKey(ctx, req.(*PublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetObjectsByKind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectsByKindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetObjectsByKind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.project.Service/GetObjectsByKind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetObjectsByKind(ctx, req.(*GetObjectsByKindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetCustomObjectMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomObjectMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetCustomObjectMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.project.Service/GetCustomObjectMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetCustomObjectMetrics(ctx, req.(*GetCustomObjectMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetEffectiveGitSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEffectiveGitSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetEffectiveGitSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.project.Service/GetEffectiveGitSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetEffectiveGitSettings(ctx, req.(*GetEffectiveGitSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetEffectivePipelineSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEffectivePipelineSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetEffectivePipelineSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.project.Service/GetEffectivePipelineSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetEffectivePipelineSettings(ctx, req.(*GetEffectivePipelineSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetEffectiveNotificationSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEffectiveNotificationSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetEffectiveNotificationSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.project.Service/GetEffectiveNotificationSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetEffectiveNotificationSettings(ctx, req.(*GetEffectiveNotificationSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.project.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Service_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Service_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Service_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Service_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Service_Update_Handler,
		},
		{
			MethodName: "PublicKey",
			Handler:    _Service_PublicKey_Handler,
		},
		{
			MethodName: "GetObjectsByKind",
			Handler:    _Service_GetObjectsByKind_Handler,
		},
		{
			MethodName: "GetCustomObjectMetrics",
			Handler:    _Service_GetCustomObjectMetrics_Handler,
		},
		{
			MethodName: "GetEffectiveGitSettings",
			Handler:    _Service_GetEffectiveGitSettings_Handler,
		},
		{
			MethodName: "GetEffectivePipelineSettings",
			Handler:    _Service_GetEffectivePipelineSettings_Handler,
		},
		{
			MethodName: "GetEffectiveNotificationSettings",
			Handler:    _Service_GetEffectiveNotificationSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/project/service.proto",
}
