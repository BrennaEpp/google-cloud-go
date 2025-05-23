// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package dataflow

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	dataflowpb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newMetricsV1Beta3ClientHook clientHook

// MetricsV1Beta3CallOptions contains the retry settings for each method of MetricsV1Beta3Client.
type MetricsV1Beta3CallOptions struct {
	GetJobMetrics            []gax.CallOption
	GetJobExecutionDetails   []gax.CallOption
	GetStageExecutionDetails []gax.CallOption
}

func defaultMetricsV1Beta3GRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataflow.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("dataflow.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("dataflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://dataflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultMetricsV1Beta3CallOptions() *MetricsV1Beta3CallOptions {
	return &MetricsV1Beta3CallOptions{
		GetJobMetrics: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetJobExecutionDetails: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetStageExecutionDetails: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultMetricsV1Beta3RESTCallOptions() *MetricsV1Beta3CallOptions {
	return &MetricsV1Beta3CallOptions{
		GetJobMetrics: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetJobExecutionDetails: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetStageExecutionDetails: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalMetricsV1Beta3Client is an interface that defines the methods available from Dataflow API.
type internalMetricsV1Beta3Client interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetJobMetrics(context.Context, *dataflowpb.GetJobMetricsRequest, ...gax.CallOption) (*dataflowpb.JobMetrics, error)
	GetJobExecutionDetails(context.Context, *dataflowpb.GetJobExecutionDetailsRequest, ...gax.CallOption) *StageSummaryIterator
	GetStageExecutionDetails(context.Context, *dataflowpb.GetStageExecutionDetailsRequest, ...gax.CallOption) *WorkerDetailsIterator
}

// MetricsV1Beta3Client is a client for interacting with Dataflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Dataflow Metrics API lets you monitor the progress of Dataflow
// jobs.
type MetricsV1Beta3Client struct {
	// The internal transport-dependent client.
	internalClient internalMetricsV1Beta3Client

	// The call options for this service.
	CallOptions *MetricsV1Beta3CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MetricsV1Beta3Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MetricsV1Beta3Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *MetricsV1Beta3Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetJobMetrics request the job status.
//
// To request the status of a job, we recommend using
// projects.locations.jobs.getMetrics with a [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). Using
// projects.jobs.getMetrics is not recommended, as you can only request the
// status of jobs that are running in us-central1.
func (c *MetricsV1Beta3Client) GetJobMetrics(ctx context.Context, req *dataflowpb.GetJobMetricsRequest, opts ...gax.CallOption) (*dataflowpb.JobMetrics, error) {
	return c.internalClient.GetJobMetrics(ctx, req, opts...)
}

// GetJobExecutionDetails request detailed information about the execution status of the job.
//
// EXPERIMENTAL.  This API is subject to change or removal without notice.
func (c *MetricsV1Beta3Client) GetJobExecutionDetails(ctx context.Context, req *dataflowpb.GetJobExecutionDetailsRequest, opts ...gax.CallOption) *StageSummaryIterator {
	return c.internalClient.GetJobExecutionDetails(ctx, req, opts...)
}

// GetStageExecutionDetails request detailed information about the execution status of a stage of the
// job.
//
// EXPERIMENTAL.  This API is subject to change or removal without notice.
func (c *MetricsV1Beta3Client) GetStageExecutionDetails(ctx context.Context, req *dataflowpb.GetStageExecutionDetailsRequest, opts ...gax.CallOption) *WorkerDetailsIterator {
	return c.internalClient.GetStageExecutionDetails(ctx, req, opts...)
}

// metricsV1Beta3GRPCClient is a client for interacting with Dataflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type metricsV1Beta3GRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing MetricsV1Beta3Client
	CallOptions **MetricsV1Beta3CallOptions

	// The gRPC API client.
	metricsV1Beta3Client dataflowpb.MetricsV1Beta3Client

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewMetricsV1Beta3Client creates a new metrics v1 beta3 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The Dataflow Metrics API lets you monitor the progress of Dataflow
// jobs.
func NewMetricsV1Beta3Client(ctx context.Context, opts ...option.ClientOption) (*MetricsV1Beta3Client, error) {
	clientOpts := defaultMetricsV1Beta3GRPCClientOptions()
	if newMetricsV1Beta3ClientHook != nil {
		hookOpts, err := newMetricsV1Beta3ClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := MetricsV1Beta3Client{CallOptions: defaultMetricsV1Beta3CallOptions()}

	c := &metricsV1Beta3GRPCClient{
		connPool:             connPool,
		metricsV1Beta3Client: dataflowpb.NewMetricsV1Beta3Client(connPool),
		CallOptions:          &client.CallOptions,
		logger:               internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *metricsV1Beta3GRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *metricsV1Beta3GRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version, "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *metricsV1Beta3GRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type metricsV1Beta3RESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing MetricsV1Beta3Client
	CallOptions **MetricsV1Beta3CallOptions

	logger *slog.Logger
}

// NewMetricsV1Beta3RESTClient creates a new metrics v1 beta3 rest client.
//
// The Dataflow Metrics API lets you monitor the progress of Dataflow
// jobs.
func NewMetricsV1Beta3RESTClient(ctx context.Context, opts ...option.ClientOption) (*MetricsV1Beta3Client, error) {
	clientOpts := append(defaultMetricsV1Beta3RESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultMetricsV1Beta3RESTCallOptions()
	c := &metricsV1Beta3RESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &MetricsV1Beta3Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultMetricsV1Beta3RESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://dataflow.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://dataflow.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://dataflow.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://dataflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *metricsV1Beta3RESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN", "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *metricsV1Beta3RESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *metricsV1Beta3RESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *metricsV1Beta3GRPCClient) GetJobMetrics(ctx context.Context, req *dataflowpb.GetJobMetricsRequest, opts ...gax.CallOption) (*dataflowpb.JobMetrics, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetJobMetrics[0:len((*c.CallOptions).GetJobMetrics):len((*c.CallOptions).GetJobMetrics)], opts...)
	var resp *dataflowpb.JobMetrics
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.metricsV1Beta3Client.GetJobMetrics, req, settings.GRPC, c.logger, "GetJobMetrics")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metricsV1Beta3GRPCClient) GetJobExecutionDetails(ctx context.Context, req *dataflowpb.GetJobExecutionDetailsRequest, opts ...gax.CallOption) *StageSummaryIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetJobExecutionDetails[0:len((*c.CallOptions).GetJobExecutionDetails):len((*c.CallOptions).GetJobExecutionDetails)], opts...)
	it := &StageSummaryIterator{}
	req = proto.Clone(req).(*dataflowpb.GetJobExecutionDetailsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataflowpb.StageSummary, string, error) {
		resp := &dataflowpb.JobExecutionDetails{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = executeRPC(ctx, c.metricsV1Beta3Client.GetJobExecutionDetails, req, settings.GRPC, c.logger, "GetJobExecutionDetails")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetStages(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *metricsV1Beta3GRPCClient) GetStageExecutionDetails(ctx context.Context, req *dataflowpb.GetStageExecutionDetailsRequest, opts ...gax.CallOption) *WorkerDetailsIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId()), "stage_id", url.QueryEscape(req.GetStageId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetStageExecutionDetails[0:len((*c.CallOptions).GetStageExecutionDetails):len((*c.CallOptions).GetStageExecutionDetails)], opts...)
	it := &WorkerDetailsIterator{}
	req = proto.Clone(req).(*dataflowpb.GetStageExecutionDetailsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataflowpb.WorkerDetails, string, error) {
		resp := &dataflowpb.StageExecutionDetails{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = executeRPC(ctx, c.metricsV1Beta3Client.GetStageExecutionDetails, req, settings.GRPC, c.logger, "GetStageExecutionDetails")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetWorkers(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// GetJobMetrics request the job status.
//
// To request the status of a job, we recommend using
// projects.locations.jobs.getMetrics with a [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). Using
// projects.jobs.getMetrics is not recommended, as you can only request the
// status of jobs that are running in us-central1.
func (c *metricsV1Beta3RESTClient) GetJobMetrics(ctx context.Context, req *dataflowpb.GetJobMetricsRequest, opts ...gax.CallOption) (*dataflowpb.JobMetrics, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/jobs/%v/metrics", req.GetProjectId(), req.GetLocation(), req.GetJobId())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetStartTime() != nil {
		field, err := protojson.Marshal(req.GetStartTime())
		if err != nil {
			return nil, err
		}
		params.Add("startTime", string(field[1:len(field)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetJobMetrics[0:len((*c.CallOptions).GetJobMetrics):len((*c.CallOptions).GetJobMetrics)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataflowpb.JobMetrics{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetJobMetrics")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// GetJobExecutionDetails request detailed information about the execution status of the job.
//
// EXPERIMENTAL.  This API is subject to change or removal without notice.
func (c *metricsV1Beta3RESTClient) GetJobExecutionDetails(ctx context.Context, req *dataflowpb.GetJobExecutionDetailsRequest, opts ...gax.CallOption) *StageSummaryIterator {
	it := &StageSummaryIterator{}
	req = proto.Clone(req).(*dataflowpb.GetJobExecutionDetailsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataflowpb.StageSummary, string, error) {
		resp := &dataflowpb.JobExecutionDetails{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/jobs/%v/executionDetails", req.GetProjectId(), req.GetLocation(), req.GetJobId())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetJobExecutionDetails")
			if err != nil {
				return err
			}
			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetStages(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// GetStageExecutionDetails request detailed information about the execution status of a stage of the
// job.
//
// EXPERIMENTAL.  This API is subject to change or removal without notice.
func (c *metricsV1Beta3RESTClient) GetStageExecutionDetails(ctx context.Context, req *dataflowpb.GetStageExecutionDetailsRequest, opts ...gax.CallOption) *WorkerDetailsIterator {
	it := &WorkerDetailsIterator{}
	req = proto.Clone(req).(*dataflowpb.GetStageExecutionDetailsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataflowpb.WorkerDetails, string, error) {
		resp := &dataflowpb.StageExecutionDetails{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/jobs/%v/stages/%v/executionDetails", req.GetProjectId(), req.GetLocation(), req.GetJobId(), req.GetStageId())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetEndTime() != nil {
			field, err := protojson.Marshal(req.GetEndTime())
			if err != nil {
				return nil, "", err
			}
			params.Add("endTime", string(field[1:len(field)-1]))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req.GetStartTime() != nil {
			field, err := protojson.Marshal(req.GetStartTime())
			if err != nil {
				return nil, "", err
			}
			params.Add("startTime", string(field[1:len(field)-1]))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetStageExecutionDetails")
			if err != nil {
				return err
			}
			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetWorkers(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}
