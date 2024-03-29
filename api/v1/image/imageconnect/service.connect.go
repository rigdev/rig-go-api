// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/image/service.proto

package imageconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	image "github.com/rigdev/rig-go-api/api/v1/image"
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
	ServiceName = "api.v1.image.Service"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ServiceGetImageInfoProcedure is the fully-qualified name of the Service's GetImageInfo RPC.
	ServiceGetImageInfoProcedure = "/api.v1.image.Service/GetImageInfo"
	// ServiceGetRepositoryInfoProcedure is the fully-qualified name of the Service's GetRepositoryInfo
	// RPC.
	ServiceGetRepositoryInfoProcedure = "/api.v1.image.Service/GetRepositoryInfo"
	// ServiceGetProcedure is the fully-qualified name of the Service's Get RPC.
	ServiceGetProcedure = "/api.v1.image.Service/Get"
	// ServiceAddProcedure is the fully-qualified name of the Service's Add RPC.
	ServiceAddProcedure = "/api.v1.image.Service/Add"
	// ServiceListProcedure is the fully-qualified name of the Service's List RPC.
	ServiceListProcedure = "/api.v1.image.Service/List"
	// ServiceDeleteProcedure is the fully-qualified name of the Service's Delete RPC.
	ServiceDeleteProcedure = "/api.v1.image.Service/Delete"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	serviceServiceDescriptor                 = image.File_api_v1_image_service_proto.Services().ByName("Service")
	serviceGetImageInfoMethodDescriptor      = serviceServiceDescriptor.Methods().ByName("GetImageInfo")
	serviceGetRepositoryInfoMethodDescriptor = serviceServiceDescriptor.Methods().ByName("GetRepositoryInfo")
	serviceGetMethodDescriptor               = serviceServiceDescriptor.Methods().ByName("Get")
	serviceAddMethodDescriptor               = serviceServiceDescriptor.Methods().ByName("Add")
	serviceListMethodDescriptor              = serviceServiceDescriptor.Methods().ByName("List")
	serviceDeleteMethodDescriptor            = serviceServiceDescriptor.Methods().ByName("Delete")
)

// ServiceClient is a client for the api.v1.image.Service service.
type ServiceClient interface {
	// Get Information about an image in a image.
	GetImageInfo(context.Context, *connect.Request[image.GetImageInfoRequest]) (*connect.Response[image.GetImageInfoResponse], error)
	// Get Information about a docker registry repository.
	GetRepositoryInfo(context.Context, *connect.Request[image.GetRepositoryInfoRequest]) (*connect.Response[image.GetRepositoryInfoResponse], error)
	// Get a image.
	Get(context.Context, *connect.Request[image.GetRequest]) (*connect.Response[image.GetResponse], error)
	// Add a new image.
	// Images are immutable and cannot change. Add a new image to make
	// changes from an existing one.
	Add(context.Context, *connect.Request[image.AddRequest]) (*connect.Response[image.AddResponse], error)
	// List images for a capsule.
	List(context.Context, *connect.Request[image.ListRequest]) (*connect.Response[image.ListResponse], error)
	// Delete a image.
	Delete(context.Context, *connect.Request[image.DeleteRequest]) (*connect.Response[image.DeleteResponse], error)
}

// NewServiceClient constructs a client for the api.v1.image.Service service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &serviceClient{
		getImageInfo: connect.NewClient[image.GetImageInfoRequest, image.GetImageInfoResponse](
			httpClient,
			baseURL+ServiceGetImageInfoProcedure,
			connect.WithSchema(serviceGetImageInfoMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getRepositoryInfo: connect.NewClient[image.GetRepositoryInfoRequest, image.GetRepositoryInfoResponse](
			httpClient,
			baseURL+ServiceGetRepositoryInfoProcedure,
			connect.WithSchema(serviceGetRepositoryInfoMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		get: connect.NewClient[image.GetRequest, image.GetResponse](
			httpClient,
			baseURL+ServiceGetProcedure,
			connect.WithSchema(serviceGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		add: connect.NewClient[image.AddRequest, image.AddResponse](
			httpClient,
			baseURL+ServiceAddProcedure,
			connect.WithSchema(serviceAddMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		list: connect.NewClient[image.ListRequest, image.ListResponse](
			httpClient,
			baseURL+ServiceListProcedure,
			connect.WithSchema(serviceListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		delete: connect.NewClient[image.DeleteRequest, image.DeleteResponse](
			httpClient,
			baseURL+ServiceDeleteProcedure,
			connect.WithSchema(serviceDeleteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// serviceClient implements ServiceClient.
type serviceClient struct {
	getImageInfo      *connect.Client[image.GetImageInfoRequest, image.GetImageInfoResponse]
	getRepositoryInfo *connect.Client[image.GetRepositoryInfoRequest, image.GetRepositoryInfoResponse]
	get               *connect.Client[image.GetRequest, image.GetResponse]
	add               *connect.Client[image.AddRequest, image.AddResponse]
	list              *connect.Client[image.ListRequest, image.ListResponse]
	delete            *connect.Client[image.DeleteRequest, image.DeleteResponse]
}

// GetImageInfo calls api.v1.image.Service.GetImageInfo.
func (c *serviceClient) GetImageInfo(ctx context.Context, req *connect.Request[image.GetImageInfoRequest]) (*connect.Response[image.GetImageInfoResponse], error) {
	return c.getImageInfo.CallUnary(ctx, req)
}

// GetRepositoryInfo calls api.v1.image.Service.GetRepositoryInfo.
func (c *serviceClient) GetRepositoryInfo(ctx context.Context, req *connect.Request[image.GetRepositoryInfoRequest]) (*connect.Response[image.GetRepositoryInfoResponse], error) {
	return c.getRepositoryInfo.CallUnary(ctx, req)
}

// Get calls api.v1.image.Service.Get.
func (c *serviceClient) Get(ctx context.Context, req *connect.Request[image.GetRequest]) (*connect.Response[image.GetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// Add calls api.v1.image.Service.Add.
func (c *serviceClient) Add(ctx context.Context, req *connect.Request[image.AddRequest]) (*connect.Response[image.AddResponse], error) {
	return c.add.CallUnary(ctx, req)
}

// List calls api.v1.image.Service.List.
func (c *serviceClient) List(ctx context.Context, req *connect.Request[image.ListRequest]) (*connect.Response[image.ListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Delete calls api.v1.image.Service.Delete.
func (c *serviceClient) Delete(ctx context.Context, req *connect.Request[image.DeleteRequest]) (*connect.Response[image.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// ServiceHandler is an implementation of the api.v1.image.Service service.
type ServiceHandler interface {
	// Get Information about an image in a image.
	GetImageInfo(context.Context, *connect.Request[image.GetImageInfoRequest]) (*connect.Response[image.GetImageInfoResponse], error)
	// Get Information about a docker registry repository.
	GetRepositoryInfo(context.Context, *connect.Request[image.GetRepositoryInfoRequest]) (*connect.Response[image.GetRepositoryInfoResponse], error)
	// Get a image.
	Get(context.Context, *connect.Request[image.GetRequest]) (*connect.Response[image.GetResponse], error)
	// Add a new image.
	// Images are immutable and cannot change. Add a new image to make
	// changes from an existing one.
	Add(context.Context, *connect.Request[image.AddRequest]) (*connect.Response[image.AddResponse], error)
	// List images for a capsule.
	List(context.Context, *connect.Request[image.ListRequest]) (*connect.Response[image.ListResponse], error)
	// Delete a image.
	Delete(context.Context, *connect.Request[image.DeleteRequest]) (*connect.Response[image.DeleteResponse], error)
}

// NewServiceHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewServiceHandler(svc ServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	serviceGetImageInfoHandler := connect.NewUnaryHandler(
		ServiceGetImageInfoProcedure,
		svc.GetImageInfo,
		connect.WithSchema(serviceGetImageInfoMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serviceGetRepositoryInfoHandler := connect.NewUnaryHandler(
		ServiceGetRepositoryInfoProcedure,
		svc.GetRepositoryInfo,
		connect.WithSchema(serviceGetRepositoryInfoMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serviceGetHandler := connect.NewUnaryHandler(
		ServiceGetProcedure,
		svc.Get,
		connect.WithSchema(serviceGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serviceAddHandler := connect.NewUnaryHandler(
		ServiceAddProcedure,
		svc.Add,
		connect.WithSchema(serviceAddMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serviceListHandler := connect.NewUnaryHandler(
		ServiceListProcedure,
		svc.List,
		connect.WithSchema(serviceListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serviceDeleteHandler := connect.NewUnaryHandler(
		ServiceDeleteProcedure,
		svc.Delete,
		connect.WithSchema(serviceDeleteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.image.Service/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ServiceGetImageInfoProcedure:
			serviceGetImageInfoHandler.ServeHTTP(w, r)
		case ServiceGetRepositoryInfoProcedure:
			serviceGetRepositoryInfoHandler.ServeHTTP(w, r)
		case ServiceGetProcedure:
			serviceGetHandler.ServeHTTP(w, r)
		case ServiceAddProcedure:
			serviceAddHandler.ServeHTTP(w, r)
		case ServiceListProcedure:
			serviceListHandler.ServeHTTP(w, r)
		case ServiceDeleteProcedure:
			serviceDeleteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedServiceHandler struct{}

func (UnimplementedServiceHandler) GetImageInfo(context.Context, *connect.Request[image.GetImageInfoRequest]) (*connect.Response[image.GetImageInfoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.image.Service.GetImageInfo is not implemented"))
}

func (UnimplementedServiceHandler) GetRepositoryInfo(context.Context, *connect.Request[image.GetRepositoryInfoRequest]) (*connect.Response[image.GetRepositoryInfoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.image.Service.GetRepositoryInfo is not implemented"))
}

func (UnimplementedServiceHandler) Get(context.Context, *connect.Request[image.GetRequest]) (*connect.Response[image.GetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.image.Service.Get is not implemented"))
}

func (UnimplementedServiceHandler) Add(context.Context, *connect.Request[image.AddRequest]) (*connect.Response[image.AddResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.image.Service.Add is not implemented"))
}

func (UnimplementedServiceHandler) List(context.Context, *connect.Request[image.ListRequest]) (*connect.Response[image.ListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.image.Service.List is not implemented"))
}

func (UnimplementedServiceHandler) Delete(context.Context, *connect.Request[image.DeleteRequest]) (*connect.Response[image.DeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.image.Service.Delete is not implemented"))
}
