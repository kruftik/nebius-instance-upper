// Code generated by sdkgen. DO NOT EDIT.

// nolint
package endpoint

import (
	"context"

	"google.golang.org/grpc"

	endpoint "github.com/yandex-cloud/go-genproto/yandex/cloud/endpoint"
)

//revive:disable

// ApiEndpointServiceClient is a endpoint.ApiEndpointServiceClient with
// lazy GRPC connection initialization.
type ApiEndpointServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Get implements endpoint.ApiEndpointServiceClient
func (c *ApiEndpointServiceClient) Get(ctx context.Context, in *endpoint.GetApiEndpointRequest, opts ...grpc.CallOption) (*endpoint.ApiEndpoint, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return endpoint.NewApiEndpointServiceClient(conn).Get(ctx, in, opts...)
}

// List implements endpoint.ApiEndpointServiceClient
func (c *ApiEndpointServiceClient) List(ctx context.Context, in *endpoint.ListApiEndpointsRequest, opts ...grpc.CallOption) (*endpoint.ListApiEndpointsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return endpoint.NewApiEndpointServiceClient(conn).List(ctx, in, opts...)
}

type ApiEndpointIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ApiEndpointServiceClient
	request *endpoint.ListApiEndpointsRequest

	items []*endpoint.ApiEndpoint
}

func (c *ApiEndpointServiceClient) ApiEndpointIterator(ctx context.Context, req *endpoint.ListApiEndpointsRequest, opts ...grpc.CallOption) *ApiEndpointIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ApiEndpointIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ApiEndpointIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Endpoints
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ApiEndpointIterator) Take(size int64) ([]*endpoint.ApiEndpoint, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*endpoint.ApiEndpoint

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ApiEndpointIterator) TakeAll() ([]*endpoint.ApiEndpoint, error) {
	return it.Take(0)
}

func (it *ApiEndpointIterator) Value() *endpoint.ApiEndpoint {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ApiEndpointIterator) Error() error {
	return it.err
}
