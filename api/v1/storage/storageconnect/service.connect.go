// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/storage/service.proto

package storageconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	storage "github.com/rigdev/rig-go-api/api/v1/storage"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ServiceName is the fully-qualified name of the Service service.
	ServiceName = "api.v1.storage.Service"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ServiceCreateBucketProcedure is the fully-qualified name of the Service's CreateBucket RPC.
	ServiceCreateBucketProcedure = "/api.v1.storage.Service/CreateBucket"
	// ServiceGetBucketProcedure is the fully-qualified name of the Service's GetBucket RPC.
	ServiceGetBucketProcedure = "/api.v1.storage.Service/GetBucket"
	// ServiceListBucketsProcedure is the fully-qualified name of the Service's ListBuckets RPC.
	ServiceListBucketsProcedure = "/api.v1.storage.Service/ListBuckets"
	// ServiceDeleteBucketProcedure is the fully-qualified name of the Service's DeleteBucket RPC.
	ServiceDeleteBucketProcedure = "/api.v1.storage.Service/DeleteBucket"
	// ServiceUnlinkBucketProcedure is the fully-qualified name of the Service's UnlinkBucket RPC.
	ServiceUnlinkBucketProcedure = "/api.v1.storage.Service/UnlinkBucket"
	// ServiceGetObjectProcedure is the fully-qualified name of the Service's GetObject RPC.
	ServiceGetObjectProcedure = "/api.v1.storage.Service/GetObject"
	// ServiceListObjectsProcedure is the fully-qualified name of the Service's ListObjects RPC.
	ServiceListObjectsProcedure = "/api.v1.storage.Service/ListObjects"
	// ServiceDeleteObjectProcedure is the fully-qualified name of the Service's DeleteObject RPC.
	ServiceDeleteObjectProcedure = "/api.v1.storage.Service/DeleteObject"
	// ServiceCopyObjectProcedure is the fully-qualified name of the Service's CopyObject RPC.
	ServiceCopyObjectProcedure = "/api.v1.storage.Service/CopyObject"
	// ServiceUploadObjectProcedure is the fully-qualified name of the Service's UploadObject RPC.
	ServiceUploadObjectProcedure = "/api.v1.storage.Service/UploadObject"
	// ServiceDownloadObjectProcedure is the fully-qualified name of the Service's DownloadObject RPC.
	ServiceDownloadObjectProcedure = "/api.v1.storage.Service/DownloadObject"
	// ServiceCreateProviderProcedure is the fully-qualified name of the Service's CreateProvider RPC.
	ServiceCreateProviderProcedure = "/api.v1.storage.Service/CreateProvider"
	// ServiceDeleteProviderProcedure is the fully-qualified name of the Service's DeleteProvider RPC.
	ServiceDeleteProviderProcedure = "/api.v1.storage.Service/DeleteProvider"
	// ServiceGetProviderProcedure is the fully-qualified name of the Service's GetProvider RPC.
	ServiceGetProviderProcedure = "/api.v1.storage.Service/GetProvider"
	// ServiceListProvidersProcedure is the fully-qualified name of the Service's ListProviders RPC.
	ServiceListProvidersProcedure = "/api.v1.storage.Service/ListProviders"
	// ServiceLookupProviderProcedure is the fully-qualified name of the Service's LookupProvider RPC.
	ServiceLookupProviderProcedure = "/api.v1.storage.Service/LookupProvider"
)

// ServiceClient is a client for the api.v1.storage.Service service.
type ServiceClient interface {
	CreateBucket(context.Context, *connect_go.Request[storage.CreateBucketRequest]) (*connect_go.Response[storage.CreateBucketResponse], error)
	GetBucket(context.Context, *connect_go.Request[storage.GetBucketRequest]) (*connect_go.Response[storage.GetBucketResponse], error)
	ListBuckets(context.Context, *connect_go.Request[storage.ListBucketsRequest]) (*connect_go.Response[storage.ListBucketsResponse], error)
	DeleteBucket(context.Context, *connect_go.Request[storage.DeleteBucketRequest]) (*connect_go.Response[storage.DeleteBucketResponse], error)
	UnlinkBucket(context.Context, *connect_go.Request[storage.UnlinkBucketRequest]) (*connect_go.Response[storage.UnlinkBucketResponse], error)
	GetObject(context.Context, *connect_go.Request[storage.GetObjectRequest]) (*connect_go.Response[storage.GetObjectResponse], error)
	ListObjects(context.Context, *connect_go.Request[storage.ListObjectsRequest]) (*connect_go.Response[storage.ListObjectsResponse], error)
	DeleteObject(context.Context, *connect_go.Request[storage.DeleteObjectRequest]) (*connect_go.Response[storage.DeleteObjectResponse], error)
	CopyObject(context.Context, *connect_go.Request[storage.CopyObjectRequest]) (*connect_go.Response[storage.CopyObjectResponse], error)
	UploadObject(context.Context) *connect_go.ClientStreamForClient[storage.UploadObjectRequest, storage.UploadObjectResponse]
	DownloadObject(context.Context, *connect_go.Request[storage.DownloadObjectRequest]) (*connect_go.ServerStreamForClient[storage.DownloadObjectResponse], error)
	CreateProvider(context.Context, *connect_go.Request[storage.CreateProviderRequest]) (*connect_go.Response[storage.CreateProviderResponse], error)
	DeleteProvider(context.Context, *connect_go.Request[storage.DeleteProviderRequest]) (*connect_go.Response[storage.DeleteProviderResponse], error)
	GetProvider(context.Context, *connect_go.Request[storage.GetProviderRequest]) (*connect_go.Response[storage.GetProviderResponse], error)
	ListProviders(context.Context, *connect_go.Request[storage.ListProvidersRequest]) (*connect_go.Response[storage.ListProvidersResponse], error)
	LookupProvider(context.Context, *connect_go.Request[storage.LookupProviderRequest]) (*connect_go.Response[storage.LookupProviderResponse], error)
}

// NewServiceClient constructs a client for the api.v1.storage.Service service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &serviceClient{
		createBucket: connect_go.NewClient[storage.CreateBucketRequest, storage.CreateBucketResponse](
			httpClient,
			baseURL+ServiceCreateBucketProcedure,
			opts...,
		),
		getBucket: connect_go.NewClient[storage.GetBucketRequest, storage.GetBucketResponse](
			httpClient,
			baseURL+ServiceGetBucketProcedure,
			opts...,
		),
		listBuckets: connect_go.NewClient[storage.ListBucketsRequest, storage.ListBucketsResponse](
			httpClient,
			baseURL+ServiceListBucketsProcedure,
			opts...,
		),
		deleteBucket: connect_go.NewClient[storage.DeleteBucketRequest, storage.DeleteBucketResponse](
			httpClient,
			baseURL+ServiceDeleteBucketProcedure,
			opts...,
		),
		unlinkBucket: connect_go.NewClient[storage.UnlinkBucketRequest, storage.UnlinkBucketResponse](
			httpClient,
			baseURL+ServiceUnlinkBucketProcedure,
			opts...,
		),
		getObject: connect_go.NewClient[storage.GetObjectRequest, storage.GetObjectResponse](
			httpClient,
			baseURL+ServiceGetObjectProcedure,
			opts...,
		),
		listObjects: connect_go.NewClient[storage.ListObjectsRequest, storage.ListObjectsResponse](
			httpClient,
			baseURL+ServiceListObjectsProcedure,
			opts...,
		),
		deleteObject: connect_go.NewClient[storage.DeleteObjectRequest, storage.DeleteObjectResponse](
			httpClient,
			baseURL+ServiceDeleteObjectProcedure,
			opts...,
		),
		copyObject: connect_go.NewClient[storage.CopyObjectRequest, storage.CopyObjectResponse](
			httpClient,
			baseURL+ServiceCopyObjectProcedure,
			opts...,
		),
		uploadObject: connect_go.NewClient[storage.UploadObjectRequest, storage.UploadObjectResponse](
			httpClient,
			baseURL+ServiceUploadObjectProcedure,
			opts...,
		),
		downloadObject: connect_go.NewClient[storage.DownloadObjectRequest, storage.DownloadObjectResponse](
			httpClient,
			baseURL+ServiceDownloadObjectProcedure,
			opts...,
		),
		createProvider: connect_go.NewClient[storage.CreateProviderRequest, storage.CreateProviderResponse](
			httpClient,
			baseURL+ServiceCreateProviderProcedure,
			opts...,
		),
		deleteProvider: connect_go.NewClient[storage.DeleteProviderRequest, storage.DeleteProviderResponse](
			httpClient,
			baseURL+ServiceDeleteProviderProcedure,
			opts...,
		),
		getProvider: connect_go.NewClient[storage.GetProviderRequest, storage.GetProviderResponse](
			httpClient,
			baseURL+ServiceGetProviderProcedure,
			opts...,
		),
		listProviders: connect_go.NewClient[storage.ListProvidersRequest, storage.ListProvidersResponse](
			httpClient,
			baseURL+ServiceListProvidersProcedure,
			opts...,
		),
		lookupProvider: connect_go.NewClient[storage.LookupProviderRequest, storage.LookupProviderResponse](
			httpClient,
			baseURL+ServiceLookupProviderProcedure,
			opts...,
		),
	}
}

// serviceClient implements ServiceClient.
type serviceClient struct {
	createBucket   *connect_go.Client[storage.CreateBucketRequest, storage.CreateBucketResponse]
	getBucket      *connect_go.Client[storage.GetBucketRequest, storage.GetBucketResponse]
	listBuckets    *connect_go.Client[storage.ListBucketsRequest, storage.ListBucketsResponse]
	deleteBucket   *connect_go.Client[storage.DeleteBucketRequest, storage.DeleteBucketResponse]
	unlinkBucket   *connect_go.Client[storage.UnlinkBucketRequest, storage.UnlinkBucketResponse]
	getObject      *connect_go.Client[storage.GetObjectRequest, storage.GetObjectResponse]
	listObjects    *connect_go.Client[storage.ListObjectsRequest, storage.ListObjectsResponse]
	deleteObject   *connect_go.Client[storage.DeleteObjectRequest, storage.DeleteObjectResponse]
	copyObject     *connect_go.Client[storage.CopyObjectRequest, storage.CopyObjectResponse]
	uploadObject   *connect_go.Client[storage.UploadObjectRequest, storage.UploadObjectResponse]
	downloadObject *connect_go.Client[storage.DownloadObjectRequest, storage.DownloadObjectResponse]
	createProvider *connect_go.Client[storage.CreateProviderRequest, storage.CreateProviderResponse]
	deleteProvider *connect_go.Client[storage.DeleteProviderRequest, storage.DeleteProviderResponse]
	getProvider    *connect_go.Client[storage.GetProviderRequest, storage.GetProviderResponse]
	listProviders  *connect_go.Client[storage.ListProvidersRequest, storage.ListProvidersResponse]
	lookupProvider *connect_go.Client[storage.LookupProviderRequest, storage.LookupProviderResponse]
}

// CreateBucket calls api.v1.storage.Service.CreateBucket.
func (c *serviceClient) CreateBucket(ctx context.Context, req *connect_go.Request[storage.CreateBucketRequest]) (*connect_go.Response[storage.CreateBucketResponse], error) {
	return c.createBucket.CallUnary(ctx, req)
}

// GetBucket calls api.v1.storage.Service.GetBucket.
func (c *serviceClient) GetBucket(ctx context.Context, req *connect_go.Request[storage.GetBucketRequest]) (*connect_go.Response[storage.GetBucketResponse], error) {
	return c.getBucket.CallUnary(ctx, req)
}

// ListBuckets calls api.v1.storage.Service.ListBuckets.
func (c *serviceClient) ListBuckets(ctx context.Context, req *connect_go.Request[storage.ListBucketsRequest]) (*connect_go.Response[storage.ListBucketsResponse], error) {
	return c.listBuckets.CallUnary(ctx, req)
}

// DeleteBucket calls api.v1.storage.Service.DeleteBucket.
func (c *serviceClient) DeleteBucket(ctx context.Context, req *connect_go.Request[storage.DeleteBucketRequest]) (*connect_go.Response[storage.DeleteBucketResponse], error) {
	return c.deleteBucket.CallUnary(ctx, req)
}

// UnlinkBucket calls api.v1.storage.Service.UnlinkBucket.
func (c *serviceClient) UnlinkBucket(ctx context.Context, req *connect_go.Request[storage.UnlinkBucketRequest]) (*connect_go.Response[storage.UnlinkBucketResponse], error) {
	return c.unlinkBucket.CallUnary(ctx, req)
}

// GetObject calls api.v1.storage.Service.GetObject.
func (c *serviceClient) GetObject(ctx context.Context, req *connect_go.Request[storage.GetObjectRequest]) (*connect_go.Response[storage.GetObjectResponse], error) {
	return c.getObject.CallUnary(ctx, req)
}

// ListObjects calls api.v1.storage.Service.ListObjects.
func (c *serviceClient) ListObjects(ctx context.Context, req *connect_go.Request[storage.ListObjectsRequest]) (*connect_go.Response[storage.ListObjectsResponse], error) {
	return c.listObjects.CallUnary(ctx, req)
}

// DeleteObject calls api.v1.storage.Service.DeleteObject.
func (c *serviceClient) DeleteObject(ctx context.Context, req *connect_go.Request[storage.DeleteObjectRequest]) (*connect_go.Response[storage.DeleteObjectResponse], error) {
	return c.deleteObject.CallUnary(ctx, req)
}

// CopyObject calls api.v1.storage.Service.CopyObject.
func (c *serviceClient) CopyObject(ctx context.Context, req *connect_go.Request[storage.CopyObjectRequest]) (*connect_go.Response[storage.CopyObjectResponse], error) {
	return c.copyObject.CallUnary(ctx, req)
}

// UploadObject calls api.v1.storage.Service.UploadObject.
func (c *serviceClient) UploadObject(ctx context.Context) *connect_go.ClientStreamForClient[storage.UploadObjectRequest, storage.UploadObjectResponse] {
	return c.uploadObject.CallClientStream(ctx)
}

// DownloadObject calls api.v1.storage.Service.DownloadObject.
func (c *serviceClient) DownloadObject(ctx context.Context, req *connect_go.Request[storage.DownloadObjectRequest]) (*connect_go.ServerStreamForClient[storage.DownloadObjectResponse], error) {
	return c.downloadObject.CallServerStream(ctx, req)
}

// CreateProvider calls api.v1.storage.Service.CreateProvider.
func (c *serviceClient) CreateProvider(ctx context.Context, req *connect_go.Request[storage.CreateProviderRequest]) (*connect_go.Response[storage.CreateProviderResponse], error) {
	return c.createProvider.CallUnary(ctx, req)
}

// DeleteProvider calls api.v1.storage.Service.DeleteProvider.
func (c *serviceClient) DeleteProvider(ctx context.Context, req *connect_go.Request[storage.DeleteProviderRequest]) (*connect_go.Response[storage.DeleteProviderResponse], error) {
	return c.deleteProvider.CallUnary(ctx, req)
}

// GetProvider calls api.v1.storage.Service.GetProvider.
func (c *serviceClient) GetProvider(ctx context.Context, req *connect_go.Request[storage.GetProviderRequest]) (*connect_go.Response[storage.GetProviderResponse], error) {
	return c.getProvider.CallUnary(ctx, req)
}

// ListProviders calls api.v1.storage.Service.ListProviders.
func (c *serviceClient) ListProviders(ctx context.Context, req *connect_go.Request[storage.ListProvidersRequest]) (*connect_go.Response[storage.ListProvidersResponse], error) {
	return c.listProviders.CallUnary(ctx, req)
}

// LookupProvider calls api.v1.storage.Service.LookupProvider.
func (c *serviceClient) LookupProvider(ctx context.Context, req *connect_go.Request[storage.LookupProviderRequest]) (*connect_go.Response[storage.LookupProviderResponse], error) {
	return c.lookupProvider.CallUnary(ctx, req)
}

// ServiceHandler is an implementation of the api.v1.storage.Service service.
type ServiceHandler interface {
	CreateBucket(context.Context, *connect_go.Request[storage.CreateBucketRequest]) (*connect_go.Response[storage.CreateBucketResponse], error)
	GetBucket(context.Context, *connect_go.Request[storage.GetBucketRequest]) (*connect_go.Response[storage.GetBucketResponse], error)
	ListBuckets(context.Context, *connect_go.Request[storage.ListBucketsRequest]) (*connect_go.Response[storage.ListBucketsResponse], error)
	DeleteBucket(context.Context, *connect_go.Request[storage.DeleteBucketRequest]) (*connect_go.Response[storage.DeleteBucketResponse], error)
	UnlinkBucket(context.Context, *connect_go.Request[storage.UnlinkBucketRequest]) (*connect_go.Response[storage.UnlinkBucketResponse], error)
	GetObject(context.Context, *connect_go.Request[storage.GetObjectRequest]) (*connect_go.Response[storage.GetObjectResponse], error)
	ListObjects(context.Context, *connect_go.Request[storage.ListObjectsRequest]) (*connect_go.Response[storage.ListObjectsResponse], error)
	DeleteObject(context.Context, *connect_go.Request[storage.DeleteObjectRequest]) (*connect_go.Response[storage.DeleteObjectResponse], error)
	CopyObject(context.Context, *connect_go.Request[storage.CopyObjectRequest]) (*connect_go.Response[storage.CopyObjectResponse], error)
	UploadObject(context.Context, *connect_go.ClientStream[storage.UploadObjectRequest]) (*connect_go.Response[storage.UploadObjectResponse], error)
	DownloadObject(context.Context, *connect_go.Request[storage.DownloadObjectRequest], *connect_go.ServerStream[storage.DownloadObjectResponse]) error
	CreateProvider(context.Context, *connect_go.Request[storage.CreateProviderRequest]) (*connect_go.Response[storage.CreateProviderResponse], error)
	DeleteProvider(context.Context, *connect_go.Request[storage.DeleteProviderRequest]) (*connect_go.Response[storage.DeleteProviderResponse], error)
	GetProvider(context.Context, *connect_go.Request[storage.GetProviderRequest]) (*connect_go.Response[storage.GetProviderResponse], error)
	ListProviders(context.Context, *connect_go.Request[storage.ListProvidersRequest]) (*connect_go.Response[storage.ListProvidersResponse], error)
	LookupProvider(context.Context, *connect_go.Request[storage.LookupProviderRequest]) (*connect_go.Response[storage.LookupProviderResponse], error)
}

// NewServiceHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewServiceHandler(svc ServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	serviceCreateBucketHandler := connect_go.NewUnaryHandler(
		ServiceCreateBucketProcedure,
		svc.CreateBucket,
		opts...,
	)
	serviceGetBucketHandler := connect_go.NewUnaryHandler(
		ServiceGetBucketProcedure,
		svc.GetBucket,
		opts...,
	)
	serviceListBucketsHandler := connect_go.NewUnaryHandler(
		ServiceListBucketsProcedure,
		svc.ListBuckets,
		opts...,
	)
	serviceDeleteBucketHandler := connect_go.NewUnaryHandler(
		ServiceDeleteBucketProcedure,
		svc.DeleteBucket,
		opts...,
	)
	serviceUnlinkBucketHandler := connect_go.NewUnaryHandler(
		ServiceUnlinkBucketProcedure,
		svc.UnlinkBucket,
		opts...,
	)
	serviceGetObjectHandler := connect_go.NewUnaryHandler(
		ServiceGetObjectProcedure,
		svc.GetObject,
		opts...,
	)
	serviceListObjectsHandler := connect_go.NewUnaryHandler(
		ServiceListObjectsProcedure,
		svc.ListObjects,
		opts...,
	)
	serviceDeleteObjectHandler := connect_go.NewUnaryHandler(
		ServiceDeleteObjectProcedure,
		svc.DeleteObject,
		opts...,
	)
	serviceCopyObjectHandler := connect_go.NewUnaryHandler(
		ServiceCopyObjectProcedure,
		svc.CopyObject,
		opts...,
	)
	serviceUploadObjectHandler := connect_go.NewClientStreamHandler(
		ServiceUploadObjectProcedure,
		svc.UploadObject,
		opts...,
	)
	serviceDownloadObjectHandler := connect_go.NewServerStreamHandler(
		ServiceDownloadObjectProcedure,
		svc.DownloadObject,
		opts...,
	)
	serviceCreateProviderHandler := connect_go.NewUnaryHandler(
		ServiceCreateProviderProcedure,
		svc.CreateProvider,
		opts...,
	)
	serviceDeleteProviderHandler := connect_go.NewUnaryHandler(
		ServiceDeleteProviderProcedure,
		svc.DeleteProvider,
		opts...,
	)
	serviceGetProviderHandler := connect_go.NewUnaryHandler(
		ServiceGetProviderProcedure,
		svc.GetProvider,
		opts...,
	)
	serviceListProvidersHandler := connect_go.NewUnaryHandler(
		ServiceListProvidersProcedure,
		svc.ListProviders,
		opts...,
	)
	serviceLookupProviderHandler := connect_go.NewUnaryHandler(
		ServiceLookupProviderProcedure,
		svc.LookupProvider,
		opts...,
	)
	return "/api.v1.storage.Service/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ServiceCreateBucketProcedure:
			serviceCreateBucketHandler.ServeHTTP(w, r)
		case ServiceGetBucketProcedure:
			serviceGetBucketHandler.ServeHTTP(w, r)
		case ServiceListBucketsProcedure:
			serviceListBucketsHandler.ServeHTTP(w, r)
		case ServiceDeleteBucketProcedure:
			serviceDeleteBucketHandler.ServeHTTP(w, r)
		case ServiceUnlinkBucketProcedure:
			serviceUnlinkBucketHandler.ServeHTTP(w, r)
		case ServiceGetObjectProcedure:
			serviceGetObjectHandler.ServeHTTP(w, r)
		case ServiceListObjectsProcedure:
			serviceListObjectsHandler.ServeHTTP(w, r)
		case ServiceDeleteObjectProcedure:
			serviceDeleteObjectHandler.ServeHTTP(w, r)
		case ServiceCopyObjectProcedure:
			serviceCopyObjectHandler.ServeHTTP(w, r)
		case ServiceUploadObjectProcedure:
			serviceUploadObjectHandler.ServeHTTP(w, r)
		case ServiceDownloadObjectProcedure:
			serviceDownloadObjectHandler.ServeHTTP(w, r)
		case ServiceCreateProviderProcedure:
			serviceCreateProviderHandler.ServeHTTP(w, r)
		case ServiceDeleteProviderProcedure:
			serviceDeleteProviderHandler.ServeHTTP(w, r)
		case ServiceGetProviderProcedure:
			serviceGetProviderHandler.ServeHTTP(w, r)
		case ServiceListProvidersProcedure:
			serviceListProvidersHandler.ServeHTTP(w, r)
		case ServiceLookupProviderProcedure:
			serviceLookupProviderHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedServiceHandler struct{}

func (UnimplementedServiceHandler) CreateBucket(context.Context, *connect_go.Request[storage.CreateBucketRequest]) (*connect_go.Response[storage.CreateBucketResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.CreateBucket is not implemented"))
}

func (UnimplementedServiceHandler) GetBucket(context.Context, *connect_go.Request[storage.GetBucketRequest]) (*connect_go.Response[storage.GetBucketResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.GetBucket is not implemented"))
}

func (UnimplementedServiceHandler) ListBuckets(context.Context, *connect_go.Request[storage.ListBucketsRequest]) (*connect_go.Response[storage.ListBucketsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.ListBuckets is not implemented"))
}

func (UnimplementedServiceHandler) DeleteBucket(context.Context, *connect_go.Request[storage.DeleteBucketRequest]) (*connect_go.Response[storage.DeleteBucketResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.DeleteBucket is not implemented"))
}

func (UnimplementedServiceHandler) UnlinkBucket(context.Context, *connect_go.Request[storage.UnlinkBucketRequest]) (*connect_go.Response[storage.UnlinkBucketResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.UnlinkBucket is not implemented"))
}

func (UnimplementedServiceHandler) GetObject(context.Context, *connect_go.Request[storage.GetObjectRequest]) (*connect_go.Response[storage.GetObjectResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.GetObject is not implemented"))
}

func (UnimplementedServiceHandler) ListObjects(context.Context, *connect_go.Request[storage.ListObjectsRequest]) (*connect_go.Response[storage.ListObjectsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.ListObjects is not implemented"))
}

func (UnimplementedServiceHandler) DeleteObject(context.Context, *connect_go.Request[storage.DeleteObjectRequest]) (*connect_go.Response[storage.DeleteObjectResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.DeleteObject is not implemented"))
}

func (UnimplementedServiceHandler) CopyObject(context.Context, *connect_go.Request[storage.CopyObjectRequest]) (*connect_go.Response[storage.CopyObjectResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.CopyObject is not implemented"))
}

func (UnimplementedServiceHandler) UploadObject(context.Context, *connect_go.ClientStream[storage.UploadObjectRequest]) (*connect_go.Response[storage.UploadObjectResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.UploadObject is not implemented"))
}

func (UnimplementedServiceHandler) DownloadObject(context.Context, *connect_go.Request[storage.DownloadObjectRequest], *connect_go.ServerStream[storage.DownloadObjectResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.DownloadObject is not implemented"))
}

func (UnimplementedServiceHandler) CreateProvider(context.Context, *connect_go.Request[storage.CreateProviderRequest]) (*connect_go.Response[storage.CreateProviderResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.CreateProvider is not implemented"))
}

func (UnimplementedServiceHandler) DeleteProvider(context.Context, *connect_go.Request[storage.DeleteProviderRequest]) (*connect_go.Response[storage.DeleteProviderResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.DeleteProvider is not implemented"))
}

func (UnimplementedServiceHandler) GetProvider(context.Context, *connect_go.Request[storage.GetProviderRequest]) (*connect_go.Response[storage.GetProviderResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.GetProvider is not implemented"))
}

func (UnimplementedServiceHandler) ListProviders(context.Context, *connect_go.Request[storage.ListProvidersRequest]) (*connect_go.Response[storage.ListProvidersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.ListProviders is not implemented"))
}

func (UnimplementedServiceHandler) LookupProvider(context.Context, *connect_go.Request[storage.LookupProviderRequest]) (*connect_go.Response[storage.LookupProviderResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.storage.Service.LookupProvider is not implemented"))
}
