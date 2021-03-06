// Code generated by sdkgen. DO NOT EDIT.

//nolint
package kafka

import (
	"context"

	"google.golang.org/grpc"

	kafka "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/kafka/v1"
)

//revive:disable

// ResourcePresetServiceClient is a kafka.ResourcePresetServiceClient with
// lazy GRPC connection initialization.
type ResourcePresetServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Get implements kafka.ResourcePresetServiceClient
func (c *ResourcePresetServiceClient) Get(ctx context.Context, in *kafka.GetResourcePresetRequest, opts ...grpc.CallOption) (*kafka.ResourcePreset, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewResourcePresetServiceClient(conn).Get(ctx, in, opts...)
}

// List implements kafka.ResourcePresetServiceClient
func (c *ResourcePresetServiceClient) List(ctx context.Context, in *kafka.ListResourcePresetsRequest, opts ...grpc.CallOption) (*kafka.ListResourcePresetsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewResourcePresetServiceClient(conn).List(ctx, in, opts...)
}

type ResourcePresetIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *ResourcePresetServiceClient
	request *kafka.ListResourcePresetsRequest

	items []*kafka.ResourcePreset
}

func (c *ResourcePresetServiceClient) ResourcePresetIterator(ctx context.Context, opts ...grpc.CallOption) *ResourcePresetIterator {
	return &ResourcePresetIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &kafka.ListResourcePresetsRequest{
			PageSize: 1000,
		},
	}
}

func (it *ResourcePresetIterator) Next() bool {
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

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.ResourcePresets
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ResourcePresetIterator) Value() *kafka.ResourcePreset {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ResourcePresetIterator) Error() error {
	return it.err
}
