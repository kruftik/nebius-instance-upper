// Code generated by sdkgen. DO NOT EDIT.

// nolint
package apploadbalancer

import (
	"context"

	"google.golang.org/grpc"

	apploadbalancer "github.com/yandex-cloud/go-genproto/yandex/cloud/apploadbalancer/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// BackendGroupServiceClient is a apploadbalancer.BackendGroupServiceClient with
// lazy GRPC connection initialization.
type BackendGroupServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// AddBackend implements apploadbalancer.BackendGroupServiceClient
func (c *BackendGroupServiceClient) AddBackend(ctx context.Context, in *apploadbalancer.AddBackendRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewBackendGroupServiceClient(conn).AddBackend(ctx, in, opts...)
}

// Create implements apploadbalancer.BackendGroupServiceClient
func (c *BackendGroupServiceClient) Create(ctx context.Context, in *apploadbalancer.CreateBackendGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewBackendGroupServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements apploadbalancer.BackendGroupServiceClient
func (c *BackendGroupServiceClient) Delete(ctx context.Context, in *apploadbalancer.DeleteBackendGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewBackendGroupServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements apploadbalancer.BackendGroupServiceClient
func (c *BackendGroupServiceClient) Get(ctx context.Context, in *apploadbalancer.GetBackendGroupRequest, opts ...grpc.CallOption) (*apploadbalancer.BackendGroup, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewBackendGroupServiceClient(conn).Get(ctx, in, opts...)
}

// List implements apploadbalancer.BackendGroupServiceClient
func (c *BackendGroupServiceClient) List(ctx context.Context, in *apploadbalancer.ListBackendGroupsRequest, opts ...grpc.CallOption) (*apploadbalancer.ListBackendGroupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewBackendGroupServiceClient(conn).List(ctx, in, opts...)
}

type BackendGroupIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *BackendGroupServiceClient
	request *apploadbalancer.ListBackendGroupsRequest

	items []*apploadbalancer.BackendGroup
}

func (c *BackendGroupServiceClient) BackendGroupIterator(ctx context.Context, req *apploadbalancer.ListBackendGroupsRequest, opts ...grpc.CallOption) *BackendGroupIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &BackendGroupIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *BackendGroupIterator) Next() bool {
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

	it.items = response.BackendGroups
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *BackendGroupIterator) Take(size int64) ([]*apploadbalancer.BackendGroup, error) {
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

	var result []*apploadbalancer.BackendGroup

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *BackendGroupIterator) TakeAll() ([]*apploadbalancer.BackendGroup, error) {
	return it.Take(0)
}

func (it *BackendGroupIterator) Value() *apploadbalancer.BackendGroup {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *BackendGroupIterator) Error() error {
	return it.err
}

// ListOperations implements apploadbalancer.BackendGroupServiceClient
func (c *BackendGroupServiceClient) ListOperations(ctx context.Context, in *apploadbalancer.ListBackendGroupOperationsRequest, opts ...grpc.CallOption) (*apploadbalancer.ListBackendGroupOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewBackendGroupServiceClient(conn).ListOperations(ctx, in, opts...)
}

type BackendGroupOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *BackendGroupServiceClient
	request *apploadbalancer.ListBackendGroupOperationsRequest

	items []*operation.Operation
}

func (c *BackendGroupServiceClient) BackendGroupOperationsIterator(ctx context.Context, req *apploadbalancer.ListBackendGroupOperationsRequest, opts ...grpc.CallOption) *BackendGroupOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &BackendGroupOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *BackendGroupOperationsIterator) Next() bool {
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

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *BackendGroupOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

	var result []*operation.Operation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *BackendGroupOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *BackendGroupOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *BackendGroupOperationsIterator) Error() error {
	return it.err
}

// RemoveBackend implements apploadbalancer.BackendGroupServiceClient
func (c *BackendGroupServiceClient) RemoveBackend(ctx context.Context, in *apploadbalancer.RemoveBackendRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewBackendGroupServiceClient(conn).RemoveBackend(ctx, in, opts...)
}

// Update implements apploadbalancer.BackendGroupServiceClient
func (c *BackendGroupServiceClient) Update(ctx context.Context, in *apploadbalancer.UpdateBackendGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewBackendGroupServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateBackend implements apploadbalancer.BackendGroupServiceClient
func (c *BackendGroupServiceClient) UpdateBackend(ctx context.Context, in *apploadbalancer.UpdateBackendRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewBackendGroupServiceClient(conn).UpdateBackend(ctx, in, opts...)
}
