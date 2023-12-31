// Code generated by sdkgen. DO NOT EDIT.

// nolint
package devices

import (
	"context"

	"google.golang.org/grpc"

	devices "github.com/yandex-cloud/go-genproto/yandex/cloud/iot/devices/v1"
)

//revive:disable

// DeviceDataServiceClient is a devices.DeviceDataServiceClient with
// lazy GRPC connection initialization.
type DeviceDataServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Publish implements devices.DeviceDataServiceClient
func (c *DeviceDataServiceClient) Publish(ctx context.Context, in *devices.PublishDeviceDataRequest, opts ...grpc.CallOption) (*devices.PublishDeviceDataResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceDataServiceClient(conn).Publish(ctx, in, opts...)
}
