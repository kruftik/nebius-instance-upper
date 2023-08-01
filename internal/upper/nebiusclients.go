package upper

import (
	"context"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	"google.golang.org/grpc"
)

type InstanceManager interface {
	Get(ctx context.Context, in *compute.GetInstanceRequest, opts ...grpc.CallOption) (*compute.Instance, error)
	List(ctx context.Context, in *compute.ListInstancesRequest, opts ...grpc.CallOption) (*compute.ListInstancesResponse, error)

	Stop(ctx context.Context, in *compute.StopInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	Start(ctx context.Context, in *compute.StartInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error)

	ListOperations(ctx context.Context, in *compute.ListInstanceOperationsRequest, opts ...grpc.CallOption) (*compute.ListInstanceOperationsResponse, error)
}
