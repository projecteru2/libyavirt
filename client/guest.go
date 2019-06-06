package client

import (
	"context"
	"fmt"

	"github.com/projecteru2/libyavirt/types"
)

func (c *Client) CreateGuest(ctx context.Context, arg types.CreateGuestReq) (reply types.Guest, err error) {
	_, err = c.Post(ctx, "/guests", arg, &reply)
	return
}

func (c *Client) DestroyGuest(ctx context.Context, id string) (types.Msg, error) {
	return c.ctrl(ctx, "/guests/destroy", id)
}

func (c *Client) StopGuest(ctx context.Context, id string) (types.Msg, error) {
	return c.ctrl(ctx, "/guests/stop", id)
}

func (c *Client) StartGuest(ctx context.Context, id string) (types.Msg, error) {
	return c.ctrl(ctx, "/guests/start", id)
}

func (c *Client) ctrl(ctx context.Context, path, id string) (reply types.Msg, err error) {
	var args = types.GuestReq{ID: id}
	_, err = c.Post(ctx, path, args, &reply)
	return
}

func (c *Client) GetGuest(ctx context.Context, id string) (reply types.Guest, err error) {
	_, err = c.Get(ctx, fmt.Sprintf("/guests/%s", id), &reply)
	return
}
