package grpcclient

import (
	"context"

	yavpb "github.com/projecteru2/libyavirt/grpc/gen"
)

func (c *GRPCClient) PushImage(ctx context.Context, imgName, user string) (string, error) {
	opts := &yavpb.PushImageOptions{ImgName: imgName, User: user}

	msg, err := c.client.PushImage(ctx, opts)
	if err != nil {
		return "", err
	}

	return msg.Err, nil
}
