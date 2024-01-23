// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/user/service.proto

package userconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	user "github.com/rigdev/rig-go-api/api/v1/user"
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
	ServiceName = "api.v1.user.Service"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ServiceCreateProcedure is the fully-qualified name of the Service's Create RPC.
	ServiceCreateProcedure = "/api.v1.user.Service/Create"
	// ServiceUpdateProcedure is the fully-qualified name of the Service's Update RPC.
	ServiceUpdateProcedure = "/api.v1.user.Service/Update"
	// ServiceListSessionsProcedure is the fully-qualified name of the Service's ListSessions RPC.
	ServiceListSessionsProcedure = "/api.v1.user.Service/ListSessions"
	// ServiceGetProcedure is the fully-qualified name of the Service's Get RPC.
	ServiceGetProcedure = "/api.v1.user.Service/Get"
	// ServiceGetByIdentifierProcedure is the fully-qualified name of the Service's GetByIdentifier RPC.
	ServiceGetByIdentifierProcedure = "/api.v1.user.Service/GetByIdentifier"
	// ServiceListProcedure is the fully-qualified name of the Service's List RPC.
	ServiceListProcedure = "/api.v1.user.Service/List"
	// ServiceDeleteProcedure is the fully-qualified name of the Service's Delete RPC.
	ServiceDeleteProcedure = "/api.v1.user.Service/Delete"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	serviceServiceDescriptor               = user.File_api_v1_user_service_proto.Services().ByName("Service")
	serviceCreateMethodDescriptor          = serviceServiceDescriptor.Methods().ByName("Create")
	serviceUpdateMethodDescriptor          = serviceServiceDescriptor.Methods().ByName("Update")
	serviceListSessionsMethodDescriptor    = serviceServiceDescriptor.Methods().ByName("ListSessions")
	serviceGetMethodDescriptor             = serviceServiceDescriptor.Methods().ByName("Get")
	serviceGetByIdentifierMethodDescriptor = serviceServiceDescriptor.Methods().ByName("GetByIdentifier")
	serviceListMethodDescriptor            = serviceServiceDescriptor.Methods().ByName("List")
	serviceDeleteMethodDescriptor          = serviceServiceDescriptor.Methods().ByName("Delete")
)

// ServiceClient is a client for the api.v1.user.Service service.
type ServiceClient interface {
	// Create a new user.
	Create(context.Context, *connect.Request[user.CreateRequest]) (*connect.Response[user.CreateResponse], error)
	// Update a users profile and info.
	Update(context.Context, *connect.Request[user.UpdateRequest]) (*connect.Response[user.UpdateResponse], error)
	// Get the list of active sessions for the given user.
	ListSessions(context.Context, *connect.Request[user.ListSessionsRequest]) (*connect.Response[user.ListSessionsResponse], error)
	// Get a user by user-id.
	Get(context.Context, *connect.Request[user.GetRequest]) (*connect.Response[user.GetResponse], error)
	// Lookup a user by a unique identifier - email, username, phone number etc.
	GetByIdentifier(context.Context, *connect.Request[user.GetByIdentifierRequest]) (*connect.Response[user.GetByIdentifierResponse], error)
	// List users.
	List(context.Context, *connect.Request[user.ListRequest]) (*connect.Response[user.ListResponse], error)
	// Delete a specific user.
	Delete(context.Context, *connect.Request[user.DeleteRequest]) (*connect.Response[user.DeleteResponse], error)
}

// NewServiceClient constructs a client for the api.v1.user.Service service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &serviceClient{
		create: connect.NewClient[user.CreateRequest, user.CreateResponse](
			httpClient,
			baseURL+ServiceCreateProcedure,
			connect.WithSchema(serviceCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		update: connect.NewClient[user.UpdateRequest, user.UpdateResponse](
			httpClient,
			baseURL+ServiceUpdateProcedure,
			connect.WithSchema(serviceUpdateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listSessions: connect.NewClient[user.ListSessionsRequest, user.ListSessionsResponse](
			httpClient,
			baseURL+ServiceListSessionsProcedure,
			connect.WithSchema(serviceListSessionsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		get: connect.NewClient[user.GetRequest, user.GetResponse](
			httpClient,
			baseURL+ServiceGetProcedure,
			connect.WithSchema(serviceGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getByIdentifier: connect.NewClient[user.GetByIdentifierRequest, user.GetByIdentifierResponse](
			httpClient,
			baseURL+ServiceGetByIdentifierProcedure,
			connect.WithSchema(serviceGetByIdentifierMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		list: connect.NewClient[user.ListRequest, user.ListResponse](
			httpClient,
			baseURL+ServiceListProcedure,
			connect.WithSchema(serviceListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		delete: connect.NewClient[user.DeleteRequest, user.DeleteResponse](
			httpClient,
			baseURL+ServiceDeleteProcedure,
			connect.WithSchema(serviceDeleteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// serviceClient implements ServiceClient.
type serviceClient struct {
	create          *connect.Client[user.CreateRequest, user.CreateResponse]
	update          *connect.Client[user.UpdateRequest, user.UpdateResponse]
	listSessions    *connect.Client[user.ListSessionsRequest, user.ListSessionsResponse]
	get             *connect.Client[user.GetRequest, user.GetResponse]
	getByIdentifier *connect.Client[user.GetByIdentifierRequest, user.GetByIdentifierResponse]
	list            *connect.Client[user.ListRequest, user.ListResponse]
	delete          *connect.Client[user.DeleteRequest, user.DeleteResponse]
}

// Create calls api.v1.user.Service.Create.
func (c *serviceClient) Create(ctx context.Context, req *connect.Request[user.CreateRequest]) (*connect.Response[user.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Update calls api.v1.user.Service.Update.
func (c *serviceClient) Update(ctx context.Context, req *connect.Request[user.UpdateRequest]) (*connect.Response[user.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// ListSessions calls api.v1.user.Service.ListSessions.
func (c *serviceClient) ListSessions(ctx context.Context, req *connect.Request[user.ListSessionsRequest]) (*connect.Response[user.ListSessionsResponse], error) {
	return c.listSessions.CallUnary(ctx, req)
}

// Get calls api.v1.user.Service.Get.
func (c *serviceClient) Get(ctx context.Context, req *connect.Request[user.GetRequest]) (*connect.Response[user.GetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// GetByIdentifier calls api.v1.user.Service.GetByIdentifier.
func (c *serviceClient) GetByIdentifier(ctx context.Context, req *connect.Request[user.GetByIdentifierRequest]) (*connect.Response[user.GetByIdentifierResponse], error) {
	return c.getByIdentifier.CallUnary(ctx, req)
}

// List calls api.v1.user.Service.List.
func (c *serviceClient) List(ctx context.Context, req *connect.Request[user.ListRequest]) (*connect.Response[user.ListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Delete calls api.v1.user.Service.Delete.
func (c *serviceClient) Delete(ctx context.Context, req *connect.Request[user.DeleteRequest]) (*connect.Response[user.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// ServiceHandler is an implementation of the api.v1.user.Service service.
type ServiceHandler interface {
	// Create a new user.
	Create(context.Context, *connect.Request[user.CreateRequest]) (*connect.Response[user.CreateResponse], error)
	// Update a users profile and info.
	Update(context.Context, *connect.Request[user.UpdateRequest]) (*connect.Response[user.UpdateResponse], error)
	// Get the list of active sessions for the given user.
	ListSessions(context.Context, *connect.Request[user.ListSessionsRequest]) (*connect.Response[user.ListSessionsResponse], error)
	// Get a user by user-id.
	Get(context.Context, *connect.Request[user.GetRequest]) (*connect.Response[user.GetResponse], error)
	// Lookup a user by a unique identifier - email, username, phone number etc.
	GetByIdentifier(context.Context, *connect.Request[user.GetByIdentifierRequest]) (*connect.Response[user.GetByIdentifierResponse], error)
	// List users.
	List(context.Context, *connect.Request[user.ListRequest]) (*connect.Response[user.ListResponse], error)
	// Delete a specific user.
	Delete(context.Context, *connect.Request[user.DeleteRequest]) (*connect.Response[user.DeleteResponse], error)
}

// NewServiceHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewServiceHandler(svc ServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	serviceCreateHandler := connect.NewUnaryHandler(
		ServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(serviceCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serviceUpdateHandler := connect.NewUnaryHandler(
		ServiceUpdateProcedure,
		svc.Update,
		connect.WithSchema(serviceUpdateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serviceListSessionsHandler := connect.NewUnaryHandler(
		ServiceListSessionsProcedure,
		svc.ListSessions,
		connect.WithSchema(serviceListSessionsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serviceGetHandler := connect.NewUnaryHandler(
		ServiceGetProcedure,
		svc.Get,
		connect.WithSchema(serviceGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serviceGetByIdentifierHandler := connect.NewUnaryHandler(
		ServiceGetByIdentifierProcedure,
		svc.GetByIdentifier,
		connect.WithSchema(serviceGetByIdentifierMethodDescriptor),
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
	return "/api.v1.user.Service/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ServiceCreateProcedure:
			serviceCreateHandler.ServeHTTP(w, r)
		case ServiceUpdateProcedure:
			serviceUpdateHandler.ServeHTTP(w, r)
		case ServiceListSessionsProcedure:
			serviceListSessionsHandler.ServeHTTP(w, r)
		case ServiceGetProcedure:
			serviceGetHandler.ServeHTTP(w, r)
		case ServiceGetByIdentifierProcedure:
			serviceGetByIdentifierHandler.ServeHTTP(w, r)
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

func (UnimplementedServiceHandler) Create(context.Context, *connect.Request[user.CreateRequest]) (*connect.Response[user.CreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.user.Service.Create is not implemented"))
}

func (UnimplementedServiceHandler) Update(context.Context, *connect.Request[user.UpdateRequest]) (*connect.Response[user.UpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.user.Service.Update is not implemented"))
}

func (UnimplementedServiceHandler) ListSessions(context.Context, *connect.Request[user.ListSessionsRequest]) (*connect.Response[user.ListSessionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.user.Service.ListSessions is not implemented"))
}

func (UnimplementedServiceHandler) Get(context.Context, *connect.Request[user.GetRequest]) (*connect.Response[user.GetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.user.Service.Get is not implemented"))
}

func (UnimplementedServiceHandler) GetByIdentifier(context.Context, *connect.Request[user.GetByIdentifierRequest]) (*connect.Response[user.GetByIdentifierResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.user.Service.GetByIdentifier is not implemented"))
}

func (UnimplementedServiceHandler) List(context.Context, *connect.Request[user.ListRequest]) (*connect.Response[user.ListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.user.Service.List is not implemented"))
}

func (UnimplementedServiceHandler) Delete(context.Context, *connect.Request[user.DeleteRequest]) (*connect.Response[user.DeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.user.Service.Delete is not implemented"))
}
