// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/project/settings/service.proto

package settingsconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	settings "github.com/rigdev/rig-go-api/api/v1/project/settings"
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
	ServiceName = "api.v1.project.settings.Service"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ServiceGetSettingsProcedure is the fully-qualified name of the Service's GetSettings RPC.
	ServiceGetSettingsProcedure = "/api.v1.project.settings.Service/GetSettings"
	// ServiceUpdateSettingsProcedure is the fully-qualified name of the Service's UpdateSettings RPC.
	ServiceUpdateSettingsProcedure = "/api.v1.project.settings.Service/UpdateSettings"
	// ServiceGetLicenseInfoProcedure is the fully-qualified name of the Service's GetLicenseInfo RPC.
	ServiceGetLicenseInfoProcedure = "/api.v1.project.settings.Service/GetLicenseInfo"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	serviceServiceDescriptor              = settings.File_api_v1_project_settings_service_proto.Services().ByName("Service")
	serviceGetSettingsMethodDescriptor    = serviceServiceDescriptor.Methods().ByName("GetSettings")
	serviceUpdateSettingsMethodDescriptor = serviceServiceDescriptor.Methods().ByName("UpdateSettings")
	serviceGetLicenseInfoMethodDescriptor = serviceServiceDescriptor.Methods().ByName("GetLicenseInfo")
)

// ServiceClient is a client for the api.v1.project.settings.Service service.
type ServiceClient interface {
	// Gets the users settings for the current project.
	GetSettings(context.Context, *connect.Request[settings.GetSettingsRequest]) (*connect.Response[settings.GetSettingsResponse], error)
	// Sets the users settings for the current project.
	UpdateSettings(context.Context, *connect.Request[settings.UpdateSettingsRequest]) (*connect.Response[settings.UpdateSettingsResponse], error)
	// Get License Information.
	GetLicenseInfo(context.Context, *connect.Request[settings.GetLicenseInfoRequest]) (*connect.Response[settings.GetLicenseInfoResponse], error)
}

// NewServiceClient constructs a client for the api.v1.project.settings.Service service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &serviceClient{
		getSettings: connect.NewClient[settings.GetSettingsRequest, settings.GetSettingsResponse](
			httpClient,
			baseURL+ServiceGetSettingsProcedure,
			connect.WithSchema(serviceGetSettingsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateSettings: connect.NewClient[settings.UpdateSettingsRequest, settings.UpdateSettingsResponse](
			httpClient,
			baseURL+ServiceUpdateSettingsProcedure,
			connect.WithSchema(serviceUpdateSettingsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getLicenseInfo: connect.NewClient[settings.GetLicenseInfoRequest, settings.GetLicenseInfoResponse](
			httpClient,
			baseURL+ServiceGetLicenseInfoProcedure,
			connect.WithSchema(serviceGetLicenseInfoMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// serviceClient implements ServiceClient.
type serviceClient struct {
	getSettings    *connect.Client[settings.GetSettingsRequest, settings.GetSettingsResponse]
	updateSettings *connect.Client[settings.UpdateSettingsRequest, settings.UpdateSettingsResponse]
	getLicenseInfo *connect.Client[settings.GetLicenseInfoRequest, settings.GetLicenseInfoResponse]
}

// GetSettings calls api.v1.project.settings.Service.GetSettings.
func (c *serviceClient) GetSettings(ctx context.Context, req *connect.Request[settings.GetSettingsRequest]) (*connect.Response[settings.GetSettingsResponse], error) {
	return c.getSettings.CallUnary(ctx, req)
}

// UpdateSettings calls api.v1.project.settings.Service.UpdateSettings.
func (c *serviceClient) UpdateSettings(ctx context.Context, req *connect.Request[settings.UpdateSettingsRequest]) (*connect.Response[settings.UpdateSettingsResponse], error) {
	return c.updateSettings.CallUnary(ctx, req)
}

// GetLicenseInfo calls api.v1.project.settings.Service.GetLicenseInfo.
func (c *serviceClient) GetLicenseInfo(ctx context.Context, req *connect.Request[settings.GetLicenseInfoRequest]) (*connect.Response[settings.GetLicenseInfoResponse], error) {
	return c.getLicenseInfo.CallUnary(ctx, req)
}

// ServiceHandler is an implementation of the api.v1.project.settings.Service service.
type ServiceHandler interface {
	// Gets the users settings for the current project.
	GetSettings(context.Context, *connect.Request[settings.GetSettingsRequest]) (*connect.Response[settings.GetSettingsResponse], error)
	// Sets the users settings for the current project.
	UpdateSettings(context.Context, *connect.Request[settings.UpdateSettingsRequest]) (*connect.Response[settings.UpdateSettingsResponse], error)
	// Get License Information.
	GetLicenseInfo(context.Context, *connect.Request[settings.GetLicenseInfoRequest]) (*connect.Response[settings.GetLicenseInfoResponse], error)
}

// NewServiceHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewServiceHandler(svc ServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	serviceGetSettingsHandler := connect.NewUnaryHandler(
		ServiceGetSettingsProcedure,
		svc.GetSettings,
		connect.WithSchema(serviceGetSettingsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serviceUpdateSettingsHandler := connect.NewUnaryHandler(
		ServiceUpdateSettingsProcedure,
		svc.UpdateSettings,
		connect.WithSchema(serviceUpdateSettingsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serviceGetLicenseInfoHandler := connect.NewUnaryHandler(
		ServiceGetLicenseInfoProcedure,
		svc.GetLicenseInfo,
		connect.WithSchema(serviceGetLicenseInfoMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.project.settings.Service/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ServiceGetSettingsProcedure:
			serviceGetSettingsHandler.ServeHTTP(w, r)
		case ServiceUpdateSettingsProcedure:
			serviceUpdateSettingsHandler.ServeHTTP(w, r)
		case ServiceGetLicenseInfoProcedure:
			serviceGetLicenseInfoHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedServiceHandler struct{}

func (UnimplementedServiceHandler) GetSettings(context.Context, *connect.Request[settings.GetSettingsRequest]) (*connect.Response[settings.GetSettingsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.project.settings.Service.GetSettings is not implemented"))
}

func (UnimplementedServiceHandler) UpdateSettings(context.Context, *connect.Request[settings.UpdateSettingsRequest]) (*connect.Response[settings.UpdateSettingsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.project.settings.Service.UpdateSettings is not implemented"))
}

func (UnimplementedServiceHandler) GetLicenseInfo(context.Context, *connect.Request[settings.GetLicenseInfoRequest]) (*connect.Response[settings.GetLicenseInfoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.project.settings.Service.GetLicenseInfo is not implemented"))
}
