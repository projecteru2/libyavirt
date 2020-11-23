package grpcclient

import (
	"context"
	"io"

	yavpb "github.com/projecteru2/libyavirt/grpc/gen"
)

// CatReadCloser .
type CatReadCloser struct {
	ID   string
	Path string
	cli  yavpb.YavirtdRPC_CatClient
}

// Read .
func (c *CatReadCloser) Read(p []byte) (int, error) {
	msg, err := c.cli.Recv()
	if err != nil {
		return 0, err
	}

	return copy(p, msg.Data), nil
}

// Close .
func (c *CatReadCloser) Close() error {
	return c.cli.CloseSend()
}

// Cat .
func (c *GRPCClient) Cat(ctx context.Context, id, path string) (io.ReadCloser, error) {
	opts := &yavpb.CatOptions{
		Id:   id,
		Path: path,
	}
	stream, err := c.client.Cat(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &CatReadCloser{
		ID:   id,
		Path: path,
		cli:  stream,
	}, nil
}
