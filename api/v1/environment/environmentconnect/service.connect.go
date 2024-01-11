// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/environment/service.proto

package environmentconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	environment "github.com/rigdev/rig-go-api/api/v1/environment"
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
	ServiceName = "api.v1.environment.Service"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ServiceListProcedure is the fully-qualified name of the Service's List RPC.
	ServiceListProcedure = "/api.v1.environment.Service/List"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	serviceServiceDescriptor    = environment.File_api_v1_environment_service_proto.Services().ByName("Service")
	serviceListMethodDescriptor = serviceServiceDescriptor.Methods().ByName("List")
)

// ServiceClient is a client for the api.v1.environment.Service service.
type ServiceClient interface {
	// List available environments.
	List(context.Context, *connect.Request[environment.ListRequest]) (*connect.Response[environment.ListResponse], error)
}

// NewServiceClient constructs a client for the api.v1.environment.Service service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &serviceClient{
		list: connect.NewClient[environment.ListRequest, environment.ListResponse](
			httpClient,
			baseURL+ServiceListProcedure,
			connect.WithSchema(serviceListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// serviceClient implements ServiceClient.
type serviceClient struct {
	list *connect.Client[environment.ListRequest, environment.ListResponse]
}

// List calls api.v1.environment.Service.List.
func (c *serviceClient) List(ctx context.Context, req *connect.Request[environment.ListRequest]) (*connect.Response[environment.ListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// ServiceHandler is an implementation of the api.v1.environment.Service service.
type ServiceHandler interface {
	// List available environments.
	List(context.Context, *connect.Request[environment.ListRequest]) (*connect.Response[environment.ListResponse], error)
}

// NewServiceHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewServiceHandler(svc ServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	serviceListHandler := connect.NewUnaryHandler(
		ServiceListProcedure,
		svc.List,
		connect.WithSchema(serviceListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.environment.Service/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ServiceListProcedure:
			serviceListHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedServiceHandler struct{}

func (UnimplementedServiceHandler) List(context.Context, *connect.Request[environment.ListRequest]) (*connect.Response[environment.ListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.environment.Service.List is not implemented"))
}