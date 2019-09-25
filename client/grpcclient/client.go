package grpcclient

import (
	"context"

	"github.com/projecteru2/libyavirt/types"
)

type grpcClient struct {
}

func New(addr string) (*grpcClient, error) {
	return &grpcClient{}, nil
}

func (c *grpcClient) Info(ctx context.Context) (info types.HostInfo, err error) {
	return
}

func (c *grpcClient) GetGuest(ctx context.Context, ID string) (guest types.Guest, err error) {
	return
}

func (c *grpcClient) CreateGuest(ctx context.Context, args types.CreateGuestReq) (guest types.Guest, err error) {
	return
}

func (c *grpcClient) StartGuest(ctx context.Context, ID string) (msg types.Msg, err error) {
	return
}

func (c *grpcClient) StopGuest(ctx context.Context, ID string) (msg types.Msg, err error) {
	return
}

func (c *grpcClient) DestroyGuest(ctx context.Context, ID string) (msg types.Msg, err error) {
	return
}
