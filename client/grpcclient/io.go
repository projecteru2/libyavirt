package grpcclient

import (
	"context"
	"errors"
	"io"

	yavpb "github.com/projecteru2/libyavirt/grpc/gen"
	"github.com/projecteru2/libyavirt/types"
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

// LogReadCloser .
type LogReadCloser struct {
	ID     string
	N      int
	client yavpb.YavirtdRPC_LogClient
}

// Read .
func (c *LogReadCloser) Read(p []byte) (int, error) {
	msg, err := c.client.Recv()
	if err != nil {
		return 0, err
	}

	return copy(p, msg.Data), nil
}

// Close .
func (c *LogReadCloser) Close() error {
	return c.client.CloseSend()
}

func (c *GRPCClient) Log(ctx context.Context, n int, id string) (io.ReadCloser, error) {
	stream, err := c.client.Log(ctx, &yavpb.LogOptions{N: int64(n), Id: id})
	if err != nil {
		return nil, err
	}

	return &LogReadCloser{
		ID:     id,
		N:      n,
		client: stream,
	}, nil
}

// CopyToGuest .
func (c *GRPCClient) CopyToGuest(ctx context.Context, id, dest string, content io.Reader, allowOverwriteDirWithFile, copyUIDGID bool) error {
	copyClient, err := c.client.CopyToGuest(ctx)
	if err != nil {
		return err
	}

	opts := &yavpb.CopyOptions{
		Id:       id,
		Dest:     dest,
		Override: allowOverwriteDirWithFile,
	}

	buf := make([]byte, types.BufferSize)
	for {
		n, err := content.Read(buf)
		if n > 0 {
			opts.Size = int64(n)
			opts.Content = buf[:n]
			if err := copyClient.Send(opts); err != nil {
				return err
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	msg, err := copyClient.CloseAndRecv()
	if err != nil {
		return err
	}
	if msg.Failed {
		return errors.New(msg.Msg)
	}
	return nil
}
