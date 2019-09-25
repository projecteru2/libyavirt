package client

import (
	"context"
	"errors"
	"net/url"

	"github.com/projecteru2/libyavirt/client/grpcclient"
	"github.com/projecteru2/libyavirt/client/httpclient"
	"github.com/projecteru2/libyavirt/types"
)

type Client interface {
	Info(context.Context) (types.HostInfo, error)
	GetGuest(ctx context.Context, ID string) (types.Guest, error)
	CreateGuest(ctx context.Context, args types.CreateGuestReq) (types.Guest, error)
	StartGuest(ctx context.Context, ID string) (types.Msg, error)
	StopGuest(ctx context.Context, ID string) (types.Msg, error)
	DestroyGuest(ctx context.Context, ID string) (types.Msg, error)
}

func NewClient(yavirtdURI string) (Client, error) {
	u, err := url.Parse(yavirtdURI)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "http":
		return httpclient.New(u.Host, u.Path[1:])
	case "grpc":
		return grpcclient.New(u.Host)
	}
	return nil, errors.New("invalid yavirtdURI: " + yavirtdURI)
}
