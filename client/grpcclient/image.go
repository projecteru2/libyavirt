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

func (c *GRPCClient) RemoveImage(ctx context.Context, imgName, user string, force, prune bool) (removed []string, err error) {
	opts := &yavpb.RemoveImageOptions{
		Image: imgName,
		User:  user,
		Force: force,
		Prune: prune,
	}

	msg, err := c.client.RemoveImage(ctx, opts)
	if err != nil {
		return nil, err
	}

	return msg.Removed, nil
}
