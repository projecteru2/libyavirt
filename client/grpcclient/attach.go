package grpcclient

import (
	"context"
	"io"

	yavpb "github.com/projecteru2/libyavirt/grpc/gen"
	"github.com/projecteru2/libyavirt/types"
)

// AttachGuestClient .
type AttachGuestClient struct {
	ID     string
	client yavpb.YavirtdRPC_AttachGuestClient
}

func (c *AttachGuestClient) Read(p []byte) (n int, err error) {
	msg, err := c.client.Recv()
	if err != nil {
		return
	}
	return copy(p, msg.Data), nil
}

func (c *AttachGuestClient) Write(p []byte) (n int, err error) {
	msg := &yavpb.AttachGuestOptions{
		Id:      c.ID,
		ReplCmd: p,
	}
	return len(p), c.client.Send(msg)
}

// Close used for WriteCloser only
func (c *AttachGuestClient) Close() error {
	return c.client.CloseSend()
}

// AttachGuest .
func (c *GRPCClient) AttachGuest(ctx context.Context, ID string, flags types.AttachGuestFlags) (stream io.ReadWriteCloser, err error) {
	resp, err := c.client.AttachGuest(ctx)
	if err != nil {
		return
	}

	opts := &yavpb.AttachGuestOptions{
		Id:    ID,
		Force: flags.Force,
		Safe:  flags.Safe,
	}

	return &AttachGuestClient{
		ID:     ID,
		client: resp,
	}, resp.Send(opts)
}
