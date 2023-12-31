// Code generated by sdkgen. DO NOT EDIT.

// nolint
package cdn

import (
	"context"

	"google.golang.org/grpc"

	cdn "github.com/yandex-cloud/go-genproto/yandex/cloud/cdn/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// OriginServiceClient is a cdn.OriginServiceClient with
// lazy GRPC connection initialization.
type OriginServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements cdn.OriginServiceClient
func (c *OriginServiceClient) Create(ctx context.Context, in *cdn.CreateOriginRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return cdn.NewOriginServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements cdn.OriginServiceClient
func (c *OriginServiceClient) Delete(ctx context.Context, in *cdn.DeleteOriginRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return cdn.NewOriginServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements cdn.OriginServiceClient
func (c *OriginServiceClient) Get(ctx context.Context, in *cdn.GetOriginRequest, opts ...grpc.CallOption) (*cdn.Origin, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return cdn.NewOriginServiceClient(conn).Get(ctx, in, opts...)
}

// List implements cdn.OriginServiceClient
func (c *OriginServiceClient) List(ctx context.Context, in *cdn.ListOriginsRequest, opts ...grpc.CallOption) (*cdn.ListOriginsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return cdn.NewOriginServiceClient(conn).List(ctx, in, opts...)
}

// Update implements cdn.OriginServiceClient
func (c *OriginServiceClient) Update(ctx context.Context, in *cdn.UpdateOriginRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return cdn.NewOriginServiceClient(conn).Update(ctx, in, opts...)
}
