// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: operator/api/v1/capabilities/service.proto

package capabilitiesconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	capabilities "github.com/rigdev/rig-go-api/operator/api/v1/capabilities"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ServiceName is the fully-qualified name of the Service service.
	ServiceName = "api.v1.capabilities.Service"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ServiceGetProcedure is the fully-qualified name of the Service's Get RPC.
	ServiceGetProcedure = "/api.v1.capabilities.Service/Get"
	// ServiceGetConfigProcedure is the fully-qualified name of the Service's GetConfig RPC.
	ServiceGetConfigProcedure = "/api.v1.capabilities.Service/GetConfig"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	serviceServiceDescriptor         = capabilities.File_operator_api_v1_capabilities_service_proto.Services().ByName("Service")
	serviceGetMethodDescriptor       = serviceServiceDescriptor.Methods().ByName("Get")
	serviceGetConfigMethodDescriptor = serviceServiceDescriptor.Methods().ByName("GetConfig")
)

// ServiceClient is a client for the api.v1.capabilities.Service service.
type ServiceClient interface {
	Get(context.Context, *connect.Request[capabilities.GetRequest]) (*connect.Response[capabilities.GetResponse], error)
	GetConfig(context.Context, *connect.Request[capabilities.GetConfigRequest]) (*connect.Response[capabilities.GetConfigResponse], error)
}

// NewServiceClient constructs a client for the api.v1.capabilities.Service service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &serviceClient{
		get: connect.NewClient[capabilities.GetRequest, capabilities.GetResponse](
			httpClient,
			baseURL+ServiceGetProcedure,
			connect.WithSchema(serviceGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getConfig: connect.NewClient[capabilities.GetConfigRequest, capabilities.GetConfigResponse](
			httpClient,
			baseURL+ServiceGetConfigProcedure,
			connect.WithSchema(serviceGetConfigMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// serviceClient implements ServiceClient.
type serviceClient struct {
	get       *connect.Client[capabilities.GetRequest, capabilities.GetResponse]
	getConfig *connect.Client[capabilities.GetConfigRequest, capabilities.GetConfigResponse]
}

// Get calls api.v1.capabilities.Service.Get.
func (c *serviceClient) Get(ctx context.Context, req *connect.Request[capabilities.GetRequest]) (*connect.Response[capabilities.GetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// GetConfig calls api.v1.capabilities.Service.GetConfig.
func (c *serviceClient) GetConfig(ctx context.Context, req *connect.Request[capabilities.GetConfigRequest]) (*connect.Response[capabilities.GetConfigResponse], error) {
	return c.getConfig.CallUnary(ctx, req)
}

// ServiceHandler is an implementation of the api.v1.capabilities.Service service.
type ServiceHandler interface {
	Get(context.Context, *connect.Request[capabilities.GetRequest]) (*connect.Response[capabilities.GetResponse], error)
	GetConfig(context.Context, *connect.Request[capabilities.GetConfigRequest]) (*connect.Response[capabilities.GetConfigResponse], error)
}

// NewServiceHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewServiceHandler(svc ServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	serviceGetHandler := connect.NewUnaryHandler(
		ServiceGetProcedure,
		svc.Get,
		connect.WithSchema(serviceGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serviceGetConfigHandler := connect.NewUnaryHandler(
		ServiceGetConfigProcedure,
		svc.GetConfig,
		connect.WithSchema(serviceGetConfigMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.capabilities.Service/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ServiceGetProcedure:
			serviceGetHandler.ServeHTTP(w, r)
		case ServiceGetConfigProcedure:
			serviceGetConfigHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedServiceHandler struct{}

func (UnimplementedServiceHandler) Get(context.Context, *connect.Request[capabilities.GetRequest]) (*connect.Response[capabilities.GetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.capabilities.Service.Get is not implemented"))
}

func (UnimplementedServiceHandler) GetConfig(context.Context, *connect.Request[capabilities.GetConfigRequest]) (*connect.Response[capabilities.GetConfigResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.capabilities.Service.GetConfig is not implemented"))
}
