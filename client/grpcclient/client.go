package grpcclient

import (
	"context"
	"encoding/json"
	"sync"

	"google.golang.org/grpc"

	yavpb "github.com/projecteru2/libyavirt/grpc/gen"
	"github.com/projecteru2/libyavirt/types"
)

var grpcClientCache = sync.Map{}

// GRPCClient .
type GRPCClient struct {
	client yavpb.YavirtdRPCClient
}

// New .
func New(addr string) (*GRPCClient, error) {
	if client, ok := grpcClientCache.Load(addr); ok {
		return client.(*GRPCClient), nil
	}
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	client := yavpb.NewYavirtdRPCClient(conn)
	grpcClient := &GRPCClient{client: client}
	grpcClientCache.Store(addr, grpcClient)
	return grpcClient, nil
}

// Info .
func (c *GRPCClient) Info(ctx context.Context) (info types.HostInfo, err error) {
	msg, err := c.client.GetInfo(ctx, &yavpb.Empty{})
	if err != nil {
		return
	}

	return types.HostInfo{
		ID:      msg.Id,
		CPU:     int(msg.Cpu),
		Mem:     msg.Memory,
		Storage: msg.Memory,
	}, nil
}

// Close .
func (c *GRPCClient) Close() error {
	return nil // grpc will manage the connections, don't have to close it here
}

// GetGuest .
func (c *GRPCClient) GetGuest(ctx context.Context, id string) (guest types.Guest, err error) {
	msg, err := c.client.GetGuest(ctx, &yavpb.GetGuestOptions{Id: id})
	if err != nil {
		return
	}

	return types.Guest{
		Resource: types.Resource{
			ID:            msg.Id,
			Status:        msg.Status,
			TransitStatus: msg.TransitStatus,
			CreateTime:    msg.CreateTime,
			TransitTime:   msg.TransitTime,
			UpdateTime:    msg.UpdateTime,
		},
		CPU:       int(msg.Cpu),
		Mem:       msg.Memory,
		Storage:   msg.Storage,
		ImageID:   msg.ImageId,
		ImageName: msg.ImageName,
		Networks:  msg.Networks,
		Labels:    msg.Labels,
		IPs:       msg.Ips,
		Hostname:  msg.Hostname,
		Running:   msg.Running,
	}, nil
}

// GetGuestUUID .
func (c *GRPCClient) GetGuestUUID(ctx context.Context, id string) (uuid string, err error) {
	msg, err := c.client.GetGuestUUID(ctx, &yavpb.GetGuestOptions{Id: id})
	if err != nil {
		return
	}
	return msg.Uuid, nil
}

func (c *GRPCClient) GetGuestIDList(ctx context.Context, args types.GetGuestIDListReq) ([]string, error) {
	resp, err := c.client.GetGuestIDList(ctx, &yavpb.GetGuestIDListOptions{Filters: args.Filters})
	if err != nil {
		return nil, err
	}
	return resp.Ids, nil
}

// CreateGuest .
func (c *GRPCClient) CreateGuest(ctx context.Context, args types.CreateGuestReq) (guest types.Guest, err error) {
	msg, err := c.client.CreateGuest(ctx, args.GetGrpcOpts())
	if err != nil {
		return
	}

	return types.Guest{
		Resource: types.Resource{
			ID:            msg.Id,
			Status:        msg.Status,
			TransitStatus: msg.TransitStatus,
			CreateTime:    msg.CreateTime,
			TransitTime:   msg.TransitTime,
			UpdateTime:    msg.UpdateTime,
		},
		CPU:       int(msg.Cpu),
		Mem:       msg.Memory,
		Storage:   msg.Storage,
		ImageID:   msg.ImageId,
		ImageName: msg.ImageName,
		ImageUser: msg.ImageUser,
		Networks:  msg.Networks,
		Labels:    msg.Labels,
	}, nil
}

// StartGuest .
func (c *GRPCClient) StartGuest(ctx context.Context, id string) (msg types.Msg, err error) {
	return c.controlGuest(ctx, id, types.OpStart, false)
}

// StopGuest .
func (c *GRPCClient) StopGuest(ctx context.Context, id string, force bool) (msg types.Msg, err error) {
	return c.controlGuest(ctx, id, types.OpStop, force)
}

// WaitGuest .
func (c *GRPCClient) WaitGuest(ctx context.Context, id string, force bool) (msg types.WaitResult, err error) {
	var result *yavpb.WaitGuestMessage
	result, err = c.client.WaitGuest(ctx, &yavpb.WaitGuestOptions{Id: id})
	if err != nil {
		return
	}

	msg.Msg = result.Msg
	msg.Code = result.Code

	return
}

// DestroyGuest .
func (c *GRPCClient) DestroyGuest(ctx context.Context, id string, force bool) (msg types.Msg, err error) {
	return c.controlGuest(ctx, id, types.OpDestroy, force)
}

func (c *GRPCClient) controlGuest(ctx context.Context, id, operation string, force bool) (msg types.Msg, err error) {
	opts := &yavpb.ControlGuestOptions{
		Id:        id,
		Operation: operation,
		Force:     force,
	}
	m, err := c.client.ControlGuest(ctx, opts)
	if err != nil {
		return
	}

	return types.Msg{Msg: m.Msg}, nil
}

// ResizeGuest .
func (c *GRPCClient) ResizeGuest(ctx context.Context, args types.ResizeGuestReq) (msg types.Msg, err error) {
	var m *yavpb.ControlGuestMessage
	if m, err = c.client.ResizeGuest(ctx, args.GetGrpcOpts()); err == nil {
		msg.Msg = m.Msg
	}
	return
}

// ExecuteGuest .
func (c *GRPCClient) ExecuteGuest(ctx context.Context, id string, cmd []string) (msg types.ExecuteGuestMessage, err error) {
	opts := &yavpb.ExecuteGuestOptions{
		Id:       id,
		Commands: cmd,
	}
	m, err := c.client.ExecuteGuest(ctx, opts)
	if err != nil {
		return
	}
	return types.ExecuteGuestMessage{
		Pid:      int(m.Pid),
		Data:     m.Data,
		ExitCode: int(m.ExitCode),
	}, nil
}

// ExecExitCode .
func (c *GRPCClient) ExecExitCode(ctx context.Context, id string, pid int) (exitCode int, err error) {
	opts := &yavpb.ExecExitCodeOptions{
		Id:  id,
		Pid: int64(pid),
	}
	m, err := c.client.ExecExitCode(ctx, opts)
	if err != nil {
		return
	}
	return int(m.ExitCode), err
}

// CaptureGuest .
func (c *GRPCClient) CaptureGuest(ctx context.Context, args types.CaptureGuestReq) (uimg types.UserImage, err error) {
	var msg *yavpb.UserImageMessage
	if msg, err = c.client.CaptureGuest(ctx, args.GetGrpcOpts()); err != nil {
		return
	}

	uimg.ID = msg.Id
	uimg.Name = msg.Name
	uimg.Distro = msg.Distro
	uimg.LatestVersion = msg.LatestVersion

	return
}

// ConnectNetwork .
func (c *GRPCClient) ConnectNetwork(ctx context.Context, args types.ConnectNetworkReq) (cidr string, err error) {
	var msg *yavpb.ConnectNetworkMessage
	if msg, err = c.client.ConnectNetwork(ctx, args.GetGrpcOpts()); err != nil {
		return
	}

	cidr = msg.Cidr

	return
}

// DisconnectNetwork .
func (c *GRPCClient) DisconnectNetwork(ctx context.Context, args types.DisconnectNetworkReq) (message string, err error) {
	var msg *yavpb.DisconnectNetworkMessage
	if msg, err = c.client.DisconnectNetwork(ctx, args.GetGrpcOpts()); err != nil {
		return
	}

	message = msg.Msg

	return
}

func (c *GRPCClient) Events(ctx context.Context, filters map[string]string) (<-chan types.EventMessage, <-chan error) {
	msgChan := make(chan types.EventMessage)
	errChan := make(chan error)

	go func() {
		defer close(errChan)
		defer close(msgChan)

		// makes an eye here, uses the outer context in this goroutine.
		client, err := c.client.Events(ctx, &yavpb.EventsOptions{Filters: filters})
		if err != nil {
			sendEvent(ctx, errChan, err)
			return
		}

		for {
			msg, err := client.Recv()
			if err != nil {
				sendEvent(ctx, errChan, err)
				return
			}

			if err := sendEvent(ctx, msgChan, types.EventMessage{
				ID:       msg.Id,
				Type:     msg.Type,
				Action:   msg.Action,
				TimeNano: msg.TimeNano,
			}); err != nil {
				return
			}
		}
	}()

	return msgChan, errChan
}

func sendEvent[T any](ctx context.Context, ch chan<- T, m T) (err error) {
	select {
	case ch <- m:
	case <-ctx.Done():
		err = ctx.Err()
	}
	return
}

// NetworkList list all networks.
func (c *GRPCClient) NetworkList(ctx context.Context, drivers []string) ([]*types.Network, error) {
	opts := &yavpb.NetworkListOptions{Drivers: drivers}
	msg, err := c.client.NetworkList(ctx, opts)
	if err != nil {
		return nil, err
	}

	var cidr []string
	var networks []*types.Network

	for name, cidrsJSON := range msg.Networks {
		if err := json.Unmarshal(cidrsJSON, &cidr); err != nil {
			return nil, err
		}

		network := &types.Network{Name: name}
		network.Subnets = append(network.Subnets, cidr...)
		networks = append(networks, network)
	}

	return networks, nil
}

// ListSnapshot .
func (c *GRPCClient) ListSnapshot(ctx context.Context, id, volID string) (snaps types.Snapshots, err error) {
	opts := &yavpb.ListSnapshotOptions{
		Id:    id,
		VolId: volID,
	}
	m, err := c.client.ListSnapshot(ctx, opts)
	if err != nil {
		return
	}

	for _, v := range m.Snapshots {
		snaps = append(snaps, &types.Snapshot{
			VolID:       v.VolId,
			VolMountDir: v.VolMountDir,
			SnapID:      v.SnapId,
			CreatedTime: v.CreatedTime,
		})
	}

	return
}

// CreateSnapshot .
func (c *GRPCClient) CreateSnapshot(ctx context.Context, id, volID string) (msg types.Msg, err error) {
	opts := &yavpb.CreateSnapshotOptions{
		Id:    id,
		VolId: volID,
	}
	if m, err := c.client.CreateSnapshot(ctx, opts); err == nil {
		msg.Msg = m.Msg
	}
	return
}

// CommitSnapshot .
func (c *GRPCClient) CommitSnapshot(ctx context.Context, id, volID, snapID string) (msg types.Msg, err error) {
	opts := &yavpb.CommitSnapshotOptions{
		Id:     id,
		VolId:  volID,
		SnapId: snapID,
	}
	if m, err := c.client.CommitSnapshot(ctx, opts); err == nil {
		msg.Msg = m.Msg
	}
	return
}

// RestoreSnapshot .
func (c *GRPCClient) RestoreSnapshot(ctx context.Context, id, volID, snapID string) (msg types.Msg, err error) {
	opts := &yavpb.RestoreSnapshotOptions{
		Id:     id,
		VolId:  volID,
		SnapId: snapID,
	}
	if m, err := c.client.RestoreSnapshot(ctx, opts); err == nil {
		msg.Msg = m.Msg
	}
	return
}
