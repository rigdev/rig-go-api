// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/issue/service.proto

package issueconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	issue "github.com/rigdev/rig-go-api/api/v1/issue"
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
	ServiceName = "api.v1.issue.Service"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ServiceGetIssuesProcedure is the fully-qualified name of the Service's GetIssues RPC.
	ServiceGetIssuesProcedure = "/api.v1.issue.Service/GetIssues"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	serviceServiceDescriptor         = issue.File_api_v1_issue_service_proto.Services().ByName("Service")
	serviceGetIssuesMethodDescriptor = serviceServiceDescriptor.Methods().ByName("GetIssues")
)

// ServiceClient is a client for the api.v1.issue.Service service.
type ServiceClient interface {
	// Get issues
	GetIssues(context.Context, *connect.Request[issue.GetIssuesRequest]) (*connect.Response[issue.GetIssuesResponse], error)
}

// NewServiceClient constructs a client for the api.v1.issue.Service service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &serviceClient{
		getIssues: connect.NewClient[issue.GetIssuesRequest, issue.GetIssuesResponse](
			httpClient,
			baseURL+ServiceGetIssuesProcedure,
			connect.WithSchema(serviceGetIssuesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// serviceClient implements ServiceClient.
type serviceClient struct {
	getIssues *connect.Client[issue.GetIssuesRequest, issue.GetIssuesResponse]
}

// GetIssues calls api.v1.issue.Service.GetIssues.
func (c *serviceClient) GetIssues(ctx context.Context, req *connect.Request[issue.GetIssuesRequest]) (*connect.Response[issue.GetIssuesResponse], error) {
	return c.getIssues.CallUnary(ctx, req)
}

// ServiceHandler is an implementation of the api.v1.issue.Service service.
type ServiceHandler interface {
	// Get issues
	GetIssues(context.Context, *connect.Request[issue.GetIssuesRequest]) (*connect.Response[issue.GetIssuesResponse], error)
}

// NewServiceHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewServiceHandler(svc ServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	serviceGetIssuesHandler := connect.NewUnaryHandler(
		ServiceGetIssuesProcedure,
		svc.GetIssues,
		connect.WithSchema(serviceGetIssuesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.issue.Service/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ServiceGetIssuesProcedure:
			serviceGetIssuesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedServiceHandler struct{}

func (UnimplementedServiceHandler) GetIssues(context.Context, *connect.Request[issue.GetIssuesRequest]) (*connect.Response[issue.GetIssuesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.issue.Service.GetIssues is not implemented"))
}
