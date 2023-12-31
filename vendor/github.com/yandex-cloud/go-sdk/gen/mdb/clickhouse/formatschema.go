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

// FormatSchemaServiceClient is a clickhouse.FormatSchemaServiceClient with
// lazy GRPC connection initialization.
type FormatSchemaServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements clickhouse.FormatSchemaServiceClient
func (c *FormatSchemaServiceClient) Create(ctx context.Context, in *clickhouse.CreateFormatSchemaRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewFormatSchemaServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements clickhouse.FormatSchemaServiceClient
func (c *FormatSchemaServiceClient) Delete(ctx context.Context, in *clickhouse.DeleteFormatSchemaRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewFormatSchemaServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements clickhouse.FormatSchemaServiceClient
func (c *FormatSchemaServiceClient) Get(ctx context.Context, in *clickhouse.GetFormatSchemaRequest, opts ...grpc.CallOption) (*clickhouse.FormatSchema, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewFormatSchemaServiceClient(conn).Get(ctx, in, opts...)
}

// List implements clickhouse.FormatSchemaServiceClient
func (c *FormatSchemaServiceClient) List(ctx context.Context, in *clickhouse.ListFormatSchemasRequest, opts ...grpc.CallOption) (*clickhouse.ListFormatSchemasResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewFormatSchemaServiceClient(conn).List(ctx, in, opts...)
}

type FormatSchemaIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *FormatSchemaServiceClient
	request *clickhouse.ListFormatSchemasRequest

	items []*clickhouse.FormatSchema
}

func (c *FormatSchemaServiceClient) FormatSchemaIterator(ctx context.Context, req *clickhouse.ListFormatSchemasRequest, opts ...grpc.CallOption) *FormatSchemaIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &FormatSchemaIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *FormatSchemaIterator) Next() bool {
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

	it.items = response.FormatSchemas
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *FormatSchemaIterator) Take(size int64) ([]*clickhouse.FormatSchema, error) {
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

	var result []*clickhouse.FormatSchema

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *FormatSchemaIterator) TakeAll() ([]*clickhouse.FormatSchema, error) {
	return it.Take(0)
}

func (it *FormatSchemaIterator) Value() *clickhouse.FormatSchema {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *FormatSchemaIterator) Error() error {
	return it.err
}

// Update implements clickhouse.FormatSchemaServiceClient
func (c *FormatSchemaServiceClient) Update(ctx context.Context, in *clickhouse.UpdateFormatSchemaRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewFormatSchemaServiceClient(conn).Update(ctx, in, opts...)
}
