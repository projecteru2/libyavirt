package grpcclient

import (
	"context"
	"io"

	yavpb "github.com/projecteru2/libyavirt/grpc/gen"
	"github.com/projecteru2/libyavirt/types"
)

type GuestConsoleClient struct {
	ID     string
	client yavpb.YavirtdRPC_ExecuteGuestClient
}

func (c *GuestConsoleClient) Read(p []byte) (n int, err error) {
	msg, err := c.client.Recv()
	if err != nil {
		return
	}
	copy(p, msg.Data)
	return len(msg.Data), nil
}

func (c *GuestConsoleClient) Write(p []byte) (n int, err error) {
	msg := &yavpb.ExecuteGuestOptions{
		Id:      c.ID,
		ReplCmd: p,
	}
	if err = c.client.Send(msg); err != nil {
		return
	}
	return len(p), nil
}

// Close used for WriteCloser only
func (c *GuestConsoleClient) Close() error {
	return c.client.CloseSend()
}

func (c *grpcClient) ExecuteGuest(ctx context.Context, ID string, commands []string, interactive bool, flag types.ExecuteGuestFlags) (stream io.ReadWriteCloser, err error) {
	resp, err := c.client.ExecuteGuest(ctx)
	if err != nil {
		return
	}

	opts := &yavpb.ExecuteGuestOptions{
		Id:          ID,
		Commands:    commands,
		Interactive: interactive,
		Force:       flag.Force,
		Safe:        flag.Safe,
	}
	if err = resp.Send(opts); err != nil {
		return
	}

	consoleClient := &GuestConsoleClient{
		ID:     ID,
		client: resp,
	}

	return consoleClient, nil
}
