package client

import (
	"context"
	"errors"
	"io"
	"net/url"

	"github.com/projecteru2/libyavirt/client/grpcclient"
	"github.com/projecteru2/libyavirt/client/httpclient"
	"github.com/projecteru2/libyavirt/types"
)

// Client .
type Client interface {
	Info(context.Context) (types.HostInfo, error)
	GetGuest(ctx context.Context, ID string) (types.Guest, error)
	GetGuestUUID(ctx context.Context, ID string) (string, error)
	GetGuestIDList(ctx context.Context, args types.GetGuestIDListReq) ([]string, error)
	CreateGuest(ctx context.Context, args types.CreateGuestReq) (types.Guest, error)
	StartGuest(ctx context.Context, ID string) (types.Msg, error)
	StopGuest(ctx context.Context, ID string, force bool) (types.Msg, error)
	DestroyGuest(ctx context.Context, ID string, force bool) (types.Msg, error)
	AttachGuest(ctx context.Context, ID string, cmd []string, flag types.AttachGuestFlags) (io.ReadWriteCloser, error)
	ResizeConsoleWindow(ctx context.Context, ID string, height, width uint) error
	ExecuteGuest(ctx context.Context, ID string, cmd []string) (types.ExecuteGuestMessage, error)
	ExecExitCode(ctx context.Context, ID string, pid int) (exitCode int, err error)
	ResizeGuest(context.Context, types.ResizeGuestReq) (types.Msg, error)
	CaptureGuest(context.Context, types.CaptureGuestReq) (types.UserImage, error)
	ConnectNetwork(context.Context, types.ConnectNetworkReq) (string, error)
	DisconnectNetwork(context.Context, types.DisconnectNetworkReq) (string, error)
	Cat(ctx context.Context, ID, path string) (io.ReadCloser, error)
	Events(context.Context, map[string]string) (<-chan types.EventMessage, <-chan error)
	CopyToGuest(ctx context.Context, ID, dest string, content io.Reader, AllowOverwriteDirWithFile, CopyUIDGID bool) error
	NetworkList(ctx context.Context, drivers []string) ([]*types.Network, error)
	WaitGuest(ctx context.Context, ID string, force bool) (types.WaitResult, error)
}

// New .
func New(yavirtdURI string) (Client, error) {
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
