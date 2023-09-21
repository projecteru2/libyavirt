package client

import (
	"context"
	"errors"
	"io"
	"net/url"

	"github.com/projecteru2/libyavirt/client/grpcclient"
	"github.com/projecteru2/libyavirt/types"
)

// Client .
type Client interface {
	Info(context.Context) (types.HostInfo, error)
	Close() error
	GetGuest(ctx context.Context, id string) (types.Guest, error)
	GetGuestUUID(ctx context.Context, id string) (string, error)
	GetGuestIDList(ctx context.Context, args types.GetGuestIDListReq) ([]string, error)
	CreateGuest(ctx context.Context, args types.CreateGuestReq) (types.Guest, error)
	StartGuest(ctx context.Context, id string) (types.Msg, error)
	StopGuest(ctx context.Context, id string, force bool) (types.Msg, error)
	DestroyGuest(ctx context.Context, id string, force bool) (types.Msg, error)
	SuspendGuest(ctx context.Context, id string) (types.Msg, error)
	ResumeGuest(ctx context.Context, id string) (types.Msg, error)
	AttachGuest(ctx context.Context, id string, cmd []string, flag types.AttachGuestFlags) (string, io.ReadWriteCloser, error)
	ResizeConsoleWindow(ctx context.Context, id string, height, width uint) error
	ExecuteGuest(ctx context.Context, id string, cmd []string) (types.ExecuteGuestMessage, error)
	ExecExitCode(ctx context.Context, id string, pid int) (exitCode int, err error)
	ResizeGuest(context.Context, types.ResizeGuestReq) (types.Msg, error)
	CaptureGuest(context.Context, types.CaptureGuestReq) (types.UserImage, error)
	ConnectNetwork(context.Context, types.ConnectNetworkReq) (string, error)
	DisconnectNetwork(context.Context, types.DisconnectNetworkReq) (string, error)
	Cat(ctx context.Context, id, path string) (io.ReadCloser, error)
	Events(context.Context, map[string]string) (<-chan types.EventMessage, <-chan error)
	CopyToGuest(ctx context.Context, id, dest string, content io.Reader, allowOverwriteDirWithFile, copyUIDGID bool) error
	NetworkList(ctx context.Context, drivers []string) ([]*types.Network, error)
	WaitGuest(ctx context.Context, id string, force bool) (types.WaitResult, error)
	Log(ctx context.Context, n int, id string) (io.ReadCloser, error)
	ListSnapshot(ctx context.Context, id, volID string) (reply types.Snapshots, err error)
	CreateSnapshot(ctx context.Context, id, volID string) (reply types.Msg, err error)
	CommitSnapshot(ctx context.Context, id, volID, snapID string) (reply types.Msg, err error)
	RestoreSnapshot(ctx context.Context, id, volID, snapID string) (reply types.Msg, err error)
	PushImage(ctx context.Context, imgName, user string) (msg string, err error)
	RemoveImage(ctx context.Context, imgName, user string, force, prune bool) (removed []string, err error)
	ListImage(ctx context.Context, filter string) (images []types.SysImage, err error)
	PullImage(ctx context.Context, imgName string, all bool) (result string, err error)
	DigestImage(ctx context.Context, image string, local bool) (digests []string, err error)
	RawEngine(ctx context.Context, req types.RawEngineReq) (resp types.RawEngineResp, err error)
}

// New .
func New(cfg *types.Config) (Client, error) {
	u, err := url.Parse(cfg.URI)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "grpc":
		return grpcclient.New(cfg)
	}
	return nil, errors.New("invalid yavirtdURI: " + cfg.URI)
}
