package grpcclient

import (
	"context"
	"io"

	yavpb "github.com/projecteru2/libyavirt/grpc/gen"
	"github.com/projecteru2/libyavirt/types"
)

type AttachGuestClient struct {
	ID     string
	client yavpb.YavirtdRPC_AttachGuestClient
}

func (c *AttachGuestClient) Read(p []byte) (n int, err error) {
	msg, err := c.client.Recv()
	if err != nil {
		return
	}
	copy(p, msg.Data)
	return len(msg.Data), nil
}

func (c *AttachGuestClient) Write(p []byte) (n int, err error) {
	msg := &yavpb.AttachGuestOptions{
		Id:      c.ID,
		ReplCmd: p,
	}
	if err = c.client.Send(msg); err != nil {
		return
	}
	return len(p), nil
}

// Close used for WriteCloser only
func (c *AttachGuestClient) Close() error {
	return c.client.CloseSend()
}

func (c *grpcClient) AttachGuest(ctx context.Context, ID string, flags types.AttachGuestFlags) (stream io.ReadWriteCloser, err error) {
	resp, err := c.client.AttachGuest(ctx)
	if err != nil {
		return
	}

	opts := &yavpb.AttachGuestOptions{
		Id:    ID,
		Force: flags.Force,
		Safe:  flags.Safe,
	}
	if err = resp.Send(opts); err != nil {
		return
	}

	consoleClient := &AttachGuestClient{
		ID:     ID,
		client: resp,
	}

	return consoleClient, nil
}
