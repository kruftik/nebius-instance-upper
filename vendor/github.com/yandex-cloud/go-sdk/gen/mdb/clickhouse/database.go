// Code generated by sdkgen. DO NOT EDIT.

// nolint
package clickhouse

import (
	"context"

	"google.golang.org/grpc"

	clickhouse "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/clickhouse/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// DatabaseServiceClient is a clickhouse.DatabaseServiceClient with
// lazy GRPC connection initialization.
type DatabaseServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements clickhouse.DatabaseServiceClient
func (c *DatabaseServiceClient) Create(ctx context.Context, in *clickhouse.CreateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewDatabaseServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements clickhouse.DatabaseServiceClient
func (c *DatabaseServiceClient) Delete(ctx context.Context, in *clickhouse.DeleteDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewDatabaseServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements clickhouse.DatabaseServiceClient
func (c *DatabaseServiceClient) Get(ctx context.Context, in *clickhouse.GetDatabaseRequest, opts ...grpc.CallOption) (*clickhouse.Database, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewDatabaseServiceClient(conn).Get(ctx, in, opts...)
}

// List implements clickhouse.DatabaseServiceClient
func (c *DatabaseServiceClient) List(ctx context.Context, in *clickhouse.ListDatabasesRequest, opts ...grpc.CallOption) (*clickhouse.ListDatabasesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewDatabaseServiceClient(conn).List(ctx, in, opts...)
}

type DatabaseIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *DatabaseServiceClient
	request *clickhouse.ListDatabasesRequest

	items []*clickhouse.Database
}

func (c *DatabaseServiceClient) DatabaseIterator(ctx context.Context, req *clickhouse.ListDatabasesRequest, opts ...grpc.CallOption) *DatabaseIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &DatabaseIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *DatabaseIterator) Next() bool {
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

	it.items = response.Databases
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *DatabaseIterator) Take(size int64) ([]*clickhouse.Database, error) {
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

	var result []*clickhouse.Database

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *DatabaseIterator) TakeAll() ([]*clickhouse.Database, error) {
	return it.Take(0)
}

func (it *DatabaseIterator) Value() *clickhouse.Database {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *DatabaseIterator) Error() error {
	return it.err
}
