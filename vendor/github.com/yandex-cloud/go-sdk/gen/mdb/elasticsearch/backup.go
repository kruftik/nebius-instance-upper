// Code generated by sdkgen. DO NOT EDIT.

// nolint
package elasticsearch

import (
	"context"

	"google.golang.org/grpc"

	elasticsearch "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/elasticsearch/v1"
)

//revive:disable

// BackupServiceClient is a elasticsearch.BackupServiceClient with
// lazy GRPC connection initialization.
type BackupServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Get implements elasticsearch.BackupServiceClient
func (c *BackupServiceClient) Get(ctx context.Context, in *elasticsearch.GetBackupRequest, opts ...grpc.CallOption) (*elasticsearch.Backup, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return elasticsearch.NewBackupServiceClient(conn).Get(ctx, in, opts...)
}

// List implements elasticsearch.BackupServiceClient
func (c *BackupServiceClient) List(ctx context.Context, in *elasticsearch.ListBackupsRequest, opts ...grpc.CallOption) (*elasticsearch.ListBackupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return elasticsearch.NewBackupServiceClient(conn).List(ctx, in, opts...)
}

type BackupIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *BackupServiceClient
	request *elasticsearch.ListBackupsRequest

	items []*elasticsearch.Backup
}

func (c *BackupServiceClient) BackupIterator(ctx context.Context, req *elasticsearch.ListBackupsRequest, opts ...grpc.CallOption) *BackupIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &BackupIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *BackupIterator) Next() bool {
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

	it.items = response.Backups
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *BackupIterator) Take(size int64) ([]*elasticsearch.Backup, error) {
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

	var result []*elasticsearch.Backup

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *BackupIterator) TakeAll() ([]*elasticsearch.Backup, error) {
	return it.Take(0)
}

func (it *BackupIterator) Value() *elasticsearch.Backup {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *BackupIterator) Error() error {
	return it.err
}
