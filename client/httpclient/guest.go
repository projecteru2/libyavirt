package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/projecteru2/libyavirt/types"
)

// CreateGuest .
func (c *HTTPClient) CreateGuest(ctx context.Context, arg types.CreateGuestReq) (reply types.Guest, err error) {
	_, err = c.Post(ctx, "/guests", arg, &reply)
	return
}

// DestroyGuest .
func (c *HTTPClient) DestroyGuest(ctx context.Context, id string, force bool) (types.Msg, error) {
	return c.ctrl(ctx, "/guests/destroy", id, force)
}

// StopGuest .
func (c *HTTPClient) StopGuest(ctx context.Context, id string, force bool) (types.Msg, error) {
	return c.ctrl(ctx, "/guests/stop", id, force)
}

// StartGuest .
func (c *HTTPClient) StartGuest(ctx context.Context, id string) (types.Msg, error) {
	return c.ctrl(ctx, "/guests/start", id, false)
}

func (c *HTTPClient) ctrl(ctx context.Context, path, id string, force bool) (reply types.Msg, err error) {
	var args = types.GuestReq{ID: id, Force: force}
	_, err = c.Post(ctx, path, args, &reply)
	return
}

// GetGuest .
func (c *HTTPClient) GetGuest(ctx context.Context, id string) (reply types.Guest, err error) {
	_, err = c.Get(ctx, fmt.Sprintf("/guests/%s", id), &reply)
	return
}

// GetGuestUUID .
func (c *HTTPClient) GetGuestUUID(ctx context.Context, id string) (uuid string, err error) {
	_, err = c.Get(ctx, fmt.Sprintf("/guests/%s/uuid", id), &uuid)
	return
}

// AttachGuest .
func (c *HTTPClient) AttachGuest(ctx context.Context, id string, cmds []string, flag types.AttachGuestFlags) (io.ReadWriteCloser, error) {
	return nil, fmt.Errorf("AttachGuest not implemented for httpclient")
}

// ResizeConsoleWindow .
func (c *HTTPClient) ResizeConsoleWindow(ctx context.Context, id string, height, width uint) error {
	args := types.ResizeConsoleWindowReq{
		GuestReq: types.GuestReq{
			ID: id,
		},
		Height: height,
		Width:  width,
	}
	_, err := c.Post(ctx, "/guests/resize_window", args, nil)
	return err
}

// Cat .
func (c *HTTPClient) Cat(context.Context, string, string) (io.ReadCloser, error) {
	return nil, fmt.Errorf("Cat not implemented for httpclient")
}

// ExecuteGuest .
func (c *HTTPClient) ExecuteGuest(ctx context.Context, id string, cmds []string) (reply types.ExecuteGuestMessage, err error) {
	var args = types.ExecuteGuestReq{}
	args.ID = id
	args.Commands = cmds
	_, err = c.Post(ctx, "/guests/execute", args, &reply)
	return
}

// ResizeGuest .
func (c *HTTPClient) ResizeGuest(ctx context.Context, args types.ResizeGuestReq) (reply types.Msg, err error) {
	_, err = c.Post(ctx, "/guests/resize", args, &reply)
	return
}

// CaptureGuest .
func (c *HTTPClient) CaptureGuest(ctx context.Context, args types.CaptureGuestReq) (reply types.UserImage, err error) {
	_, err = c.Post(ctx, "/guests/capture", args, &reply)
	return
}

// DisconnectNetwork .
func (c *HTTPClient) DisconnectNetwork(ctx context.Context, args types.DisconnectNetworkReq) (msg string, err error) {
	_, err = c.Post(ctx, "/guests/disconnect", args, &msg)
	return
}

// ConnectNetwork .
func (c *HTTPClient) ConnectNetwork(ctx context.Context, args types.ConnectNetworkReq) (cidr string, err error) {
	_, err = c.Post(ctx, "/guests/connect", args, &cidr)
	return
}

// GetGuestIDList .
func (c *HTTPClient) GetGuestIDList(ctx context.Context, args types.GetGuestIDListReq) (ids []string, err error) {
	params := url.Values{}
	for key, value := range args.Filters {
		params.Set(key, value)
	}
	_, err = c.Get(ctx, fmt.Sprintf("/guests?%s", params.Encode()), &ids)
	return
}

// Events .
func (c *HTTPClient) Events(ctx context.Context, filters map[string]string) (<-chan types.EventMessage, <-chan error) {
	msgChan := make(chan types.EventMessage)
	errChan := make(chan error)

	go func() {
		defer close(errChan)

		params := url.Values{}
		for key, value := range filters {
			params.Set(key, value)
		}

		req, err := http.NewRequestWithContext(ctx, "GET", c.getPath(fmt.Sprintf("/events?%s", params.Encode())), nil)
		if err != nil {
			errChan <- err
			return
		}
		req.URL.Host = c.addr
		req.URL.Scheme = c.scheme

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			errChan <- err
			return
		}
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)

		for {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				var event types.EventMessage
				if err := decoder.Decode(&event); err != nil {
					errChan <- err
					return
				}

				select {
				case msgChan <- event:
				case <-ctx.Done():
					errChan <- ctx.Err()
					return
				}
			}
		}
	}()

	return msgChan, errChan
}
