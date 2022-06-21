package grpcclient

import (
	"context"

	yavpb "github.com/projecteru2/libyavirt/grpc/gen"
	"github.com/projecteru2/libyavirt/types"
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

func (c *GRPCClient) ListImage(ctx context.Context, filter string) ([]types.SysImage, error) {
	opts := &yavpb.ListImageOptions{Filter: filter}

	msg, err := c.client.ListImage(ctx, opts)
	if err != nil {
		return nil, err
	}

	images := []types.SysImage{}

	for _, image := range msg.Images {
		images = append(images, types.SysImage{
			Name:   image.Name,
			User:   image.User,
			Distro: image.Distro,
			ID:     image.Id,
			Type:   image.Type,
		})
	}

	return images, nil
}

func (c *GRPCClient) PullImage(ctx context.Context, imgName string, all bool) (result string, err error) {
	opts := &yavpb.PullImageOptions{Name: imgName, All: all}

	msg, err := c.client.PullImage(ctx, opts)
	if err != nil {
		return "", err
	}

	return msg.Result, nil
}

func (c *GRPCClient) DigestImage(ctx context.Context, image string, local bool) (digests []string, err error) {
	opts := &yavpb.DigestImageOptions{ImageName: image, Local: local}

	msg, err := c.client.DigestImage(ctx, opts)
	if err != nil {
		return nil, err
	}

	return msg.Digests, nil
}
