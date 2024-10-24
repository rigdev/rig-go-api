// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/metrics/service.proto

package metricsconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	metrics "github.com/rigdev/rig-go-api/api/v1/metrics"
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
	ServiceName = "api.v1.metrics.Service"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ServiceGetMetricsProcedure is the fully-qualified name of the Service's GetMetrics RPC.
	ServiceGetMetricsProcedure = "/api.v1.metrics.Service/GetMetrics"
	// ServiceGetMetricsManyProcedure is the fully-qualified name of the Service's GetMetricsMany RPC.
	ServiceGetMetricsManyProcedure = "/api.v1.metrics.Service/GetMetricsMany"
	// ServiceGetMetricsExpressionProcedure is the fully-qualified name of the Service's
	// GetMetricsExpression RPC.
	ServiceGetMetricsExpressionProcedure = "/api.v1.metrics.Service/GetMetricsExpression"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	serviceServiceDescriptor                    = metrics.File_api_v1_metrics_service_proto.Services().ByName("Service")
	serviceGetMetricsMethodDescriptor           = serviceServiceDescriptor.Methods().ByName("GetMetrics")
	serviceGetMetricsManyMethodDescriptor       = serviceServiceDescriptor.Methods().ByName("GetMetricsMany")
	serviceGetMetricsExpressionMethodDescriptor = serviceServiceDescriptor.Methods().ByName("GetMetricsExpression")
)

// ServiceClient is a client for the api.v1.metrics.Service service.
type ServiceClient interface {
	// Retrieve metrics. metric_type is mandatory, while the rest of the fields
	// in the tags are optional. If project, env or capsule is not
	// specified, they will be treated as wildcards.
	GetMetrics(context.Context, *connect.Request[metrics.GetMetricsRequest]) (*connect.Response[metrics.GetMetricsResponse], error)
	// Retrive metrics for multiple sets of tags at a time. Metrics within the
	// same set of tags will be in ascending order of timestamp.
	GetMetricsMany(context.Context, *connect.Request[metrics.GetMetricsManyRequest]) (*connect.Response[metrics.GetMetricsManyResponse], error)
	GetMetricsExpression(context.Context, *connect.Request[metrics.GetMetricsExpressionRequest]) (*connect.Response[metrics.GetMetricsExpressionResponse], error)
}

// NewServiceClient constructs a client for the api.v1.metrics.Service service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &serviceClient{
		getMetrics: connect.NewClient[metrics.GetMetricsRequest, metrics.GetMetricsResponse](
			httpClient,
			baseURL+ServiceGetMetricsProcedure,
			connect.WithSchema(serviceGetMetricsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getMetricsMany: connect.NewClient[metrics.GetMetricsManyRequest, metrics.GetMetricsManyResponse](
			httpClient,
			baseURL+ServiceGetMetricsManyProcedure,
			connect.WithSchema(serviceGetMetricsManyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getMetricsExpression: connect.NewClient[metrics.GetMetricsExpressionRequest, metrics.GetMetricsExpressionResponse](
			httpClient,
			baseURL+ServiceGetMetricsExpressionProcedure,
			connect.WithSchema(serviceGetMetricsExpressionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// serviceClient implements ServiceClient.
type serviceClient struct {
	getMetrics           *connect.Client[metrics.GetMetricsRequest, metrics.GetMetricsResponse]
	getMetricsMany       *connect.Client[metrics.GetMetricsManyRequest, metrics.GetMetricsManyResponse]
	getMetricsExpression *connect.Client[metrics.GetMetricsExpressionRequest, metrics.GetMetricsExpressionResponse]
}

// GetMetrics calls api.v1.metrics.Service.GetMetrics.
func (c *serviceClient) GetMetrics(ctx context.Context, req *connect.Request[metrics.GetMetricsRequest]) (*connect.Response[metrics.GetMetricsResponse], error) {
	return c.getMetrics.CallUnary(ctx, req)
}

// GetMetricsMany calls api.v1.metrics.Service.GetMetricsMany.
func (c *serviceClient) GetMetricsMany(ctx context.Context, req *connect.Request[metrics.GetMetricsManyRequest]) (*connect.Response[metrics.GetMetricsManyResponse], error) {
	return c.getMetricsMany.CallUnary(ctx, req)
}

// GetMetricsExpression calls api.v1.metrics.Service.GetMetricsExpression.
func (c *serviceClient) GetMetricsExpression(ctx context.Context, req *connect.Request[metrics.GetMetricsExpressionRequest]) (*connect.Response[metrics.GetMetricsExpressionResponse], error) {
	return c.getMetricsExpression.CallUnary(ctx, req)
}

// ServiceHandler is an implementation of the api.v1.metrics.Service service.
type ServiceHandler interface {
	// Retrieve metrics. metric_type is mandatory, while the rest of the fields
	// in the tags are optional. If project, env or capsule is not
	// specified, they will be treated as wildcards.
	GetMetrics(context.Context, *connect.Request[metrics.GetMetricsRequest]) (*connect.Response[metrics.GetMetricsResponse], error)
	// Retrive metrics for multiple sets of tags at a time. Metrics within the
	// same set of tags will be in ascending order of timestamp.
	GetMetricsMany(context.Context, *connect.Request[metrics.GetMetricsManyRequest]) (*connect.Response[metrics.GetMetricsManyResponse], error)
	GetMetricsExpression(context.Context, *connect.Request[metrics.GetMetricsExpressionRequest]) (*connect.Response[metrics.GetMetricsExpressionResponse], error)
}

// NewServiceHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewServiceHandler(svc ServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	serviceGetMetricsHandler := connect.NewUnaryHandler(
		ServiceGetMetricsProcedure,
		svc.GetMetrics,
		connect.WithSchema(serviceGetMetricsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serviceGetMetricsManyHandler := connect.NewUnaryHandler(
		ServiceGetMetricsManyProcedure,
		svc.GetMetricsMany,
		connect.WithSchema(serviceGetMetricsManyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serviceGetMetricsExpressionHandler := connect.NewUnaryHandler(
		ServiceGetMetricsExpressionProcedure,
		svc.GetMetricsExpression,
		connect.WithSchema(serviceGetMetricsExpressionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.metrics.Service/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ServiceGetMetricsProcedure:
			serviceGetMetricsHandler.ServeHTTP(w, r)
		case ServiceGetMetricsManyProcedure:
			serviceGetMetricsManyHandler.ServeHTTP(w, r)
		case ServiceGetMetricsExpressionProcedure:
			serviceGetMetricsExpressionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedServiceHandler struct{}

func (UnimplementedServiceHandler) GetMetrics(context.Context, *connect.Request[metrics.GetMetricsRequest]) (*connect.Response[metrics.GetMetricsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.metrics.Service.GetMetrics is not implemented"))
}

func (UnimplementedServiceHandler) GetMetricsMany(context.Context, *connect.Request[metrics.GetMetricsManyRequest]) (*connect.Response[metrics.GetMetricsManyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.metrics.Service.GetMetricsMany is not implemented"))
}

func (UnimplementedServiceHandler) GetMetricsExpression(context.Context, *connect.Request[metrics.GetMetricsExpressionRequest]) (*connect.Response[metrics.GetMetricsExpressionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.metrics.Service.GetMetricsExpression is not implemented"))
}
