package httpclient

import (
	"context"
	"fmt"
	"io"

	"github.com/projecteru2/libyavirt/types"
)

func (c *httpClient) CreateGuest(ctx context.Context, arg types.CreateGuestReq) (reply types.Guest, err error) {
	_, err = c.Post(ctx, "/guests", arg, &reply)
	return
}

func (c *httpClient) DestroyGuest(ctx context.Context, id string, force bool) (types.Msg, error) {
	return c.ctrl(ctx, "/guests/destroy", id, force)
}

func (c *httpClient) StopGuest(ctx context.Context, id string) (types.Msg, error) {
	return c.ctrl(ctx, "/guests/stop", id, false)
}

func (c *httpClient) StartGuest(ctx context.Context, id string) (types.Msg, error) {
	return c.ctrl(ctx, "/guests/start", id, false)
}

func (c *httpClient) ctrl(ctx context.Context, path, id string, force bool) (reply types.Msg, err error) {
	var args = types.GuestReq{ID: id, Force: force}
	_, err = c.Post(ctx, path, args, &reply)
	return
}

func (c *httpClient) GetGuest(ctx context.Context, id string) (reply types.Guest, err error) {
	_, err = c.Get(ctx, fmt.Sprintf("/guests/%s", id), &reply)
	return
}

func (c *httpClient) GetGuestUUID(ctx context.Context, id string) (uuid string, err error) {
	_, err = c.Get(ctx, fmt.Sprintf("/guests/%s/uuid", id), &uuid)
	return
}

func (c *httpClient) AttachGuest(ctx context.Context, id string, flag types.AttachGuestFlags) (io.ReadWriteCloser, error) {
	return nil, fmt.Errorf("AttachGuest not implemented for httpclient")
}

func (c *httpClient) ExecuteGuest(ctx context.Context, id string, cmds []string) (reply types.ExecuteGuestMessage, err error) {
	var args = types.ExecuteGuestReq{}
	args.ID = id
	args.Commands = cmds
	_, err = c.Post(ctx, "/guests/execute", args, &reply)
	return
}
