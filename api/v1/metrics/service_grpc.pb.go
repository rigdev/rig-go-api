// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package metrics

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
	// Retrieve metrics. metric_type is mandatory, while the rest of the fields
	// in the tags are optional. If project, env or capsule is not
	// specified, they will be treated as wildcards.
	GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
	// Retrive metrics for multiple sets of tags at a time. Metrics within the
	// same set of tags will be in ascending order of timestamp.
	GetMetricsMany(ctx context.Context, in *GetMetricsManyRequest, opts ...grpc.CallOption) (*GetMetricsManyResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, "/api.v1.metrics.Service/GetMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetMetricsMany(ctx context.Context, in *GetMetricsManyRequest, opts ...grpc.CallOption) (*GetMetricsManyResponse, error) {
	out := new(GetMetricsManyResponse)
	err := c.cc.Invoke(ctx, "/api.v1.metrics.Service/GetMetricsMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// Retrieve metrics. metric_type is mandatory, while the rest of the fields
	// in the tags are optional. If project, env or capsule is not
	// specified, they will be treated as wildcards.
	GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
	// Retrive metrics for multiple sets of tags at a time. Metrics within the
	// same set of tags will be in ascending order of timestamp.
	GetMetricsMany(context.Context, *GetMetricsManyRequest) (*GetMetricsManyResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}
func (UnimplementedServiceServer) GetMetricsMany(context.Context, *GetMetricsManyRequest) (*GetMetricsManyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricsMany not implemented")
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

func _Service_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.metrics.Service/GetMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetMetrics(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetMetricsMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetMetricsMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.metrics.Service/GetMetricsMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetMetricsMany(ctx, req.(*GetMetricsManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.metrics.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetrics",
			Handler:    _Service_GetMetrics_Handler,
		},
		{
			MethodName: "GetMetricsMany",
			Handler:    _Service_GetMetricsMany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/metrics/service.proto",
}
