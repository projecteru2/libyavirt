package grpcclient

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"google.golang.org/grpc"

	yavpb "github.com/projecteru2/libyavirt/grpc/gen"
	"github.com/projecteru2/libyavirt/types"
)

// GRPCClient .
type GRPCClient struct {
	client yavpb.YavirtdRPCClient
}

// New .
func New(addr string) (*GRPCClient, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	client := yavpb.NewYavirtdRPCClient(conn)
	return &GRPCClient{client}, nil
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

// GetGuest .
func (c *GRPCClient) GetGuest(ctx context.Context, ID string) (guest types.Guest, err error) {
	msg, err := c.client.GetGuest(ctx, &yavpb.GetGuestOptions{Id: ID})
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
func (c *GRPCClient) GetGuestUUID(ctx context.Context, ID string) (uuid string, err error) {
	msg, err := c.client.GetGuestUUID(ctx, &yavpb.GetGuestOptions{Id: ID})
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
	}, nil
}

// StartGuest .
func (c *GRPCClient) StartGuest(ctx context.Context, ID string) (msg types.Msg, err error) {
	return c.controlGuest(ctx, ID, "start", false)
}

// StopGuest .
func (c *GRPCClient) StopGuest(ctx context.Context, ID string, force bool) (msg types.Msg, err error) {
	return c.controlGuest(ctx, ID, "stop", force)
}

// DestroyGuest .
func (c *GRPCClient) DestroyGuest(ctx context.Context, ID string, force bool) (msg types.Msg, err error) {
	return c.controlGuest(ctx, ID, "destroy", force)
}

func (c *GRPCClient) controlGuest(ctx context.Context, ID, operation string, force bool) (msg types.Msg, err error) {
	opts := &yavpb.ControlGuestOptions{
		Id:        ID,
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
func (c *GRPCClient) ExecuteGuest(ctx context.Context, ID string, cmd []string) (msg types.ExecuteGuestMessage, err error) {
	opts := &yavpb.ExecuteGuestOptions{
		Id:       ID,
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
func (c *GRPCClient) ExecExitCode(ctx context.Context, ID string, pid int) (exitCode int, err error) {
	opts := &yavpb.ExecExitCodeOptions{
		Id:  ID,
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
	msg := &yavpb.UserImageMessage{}
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
	msg := &yavpb.ConnectNetworkMessage{}
	if msg, err = c.client.ConnectNetwork(ctx, args.GetGrpcOpts()); err != nil {
		return
	}

	cidr = msg.Cidr

	return
}

// DisconnectNetwork .
func (c *GRPCClient) DisconnectNetwork(ctx context.Context, args types.DisconnectNetworkReq) (message string, err error) {
	msg := &yavpb.DisconnectNetworkMessage{}
	if msg, err = c.client.DisconnectNetwork(ctx, args.GetGrpcOpts()); err != nil {
		return
	}

	message = msg.Msg

	return
}

// CopyToGuest .
func (c *GRPCClient) CopyToGuest(ctx context.Context, ID, dest string, content io.Reader, AllowOverwriteDirWithFile, CopyUIDGID bool) error {
	copyClient, err := c.client.CopyToGuest(ctx)
	if err != nil {
		return err
	}

	opts := &yavpb.CopyOptions{
		Id:       ID,
		Dest:     dest,
		Override: AllowOverwriteDirWithFile,
	}

	buf := make([]byte, types.BufferSize)
	for {
		n, err := content.Read(buf)
		if n > 0 {
			opts.Size = int64(n)
			opts.Content = buf[:n]
			if err := copyClient.Send(opts); err != nil {
				return err
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	msg, err := copyClient.CloseAndRecv()
	if err != nil {
		return err
	}
	if msg.Failed {
		return errors.New(msg.Msg)
	}
	return nil
}

func (c *GRPCClient) Events(ctx context.Context, filters map[string]string) (<-chan types.EventMessage, <-chan error) {
	msgChan := make(chan types.EventMessage)
	errChan := make(chan error)
	go func() {
		defer close(errChan)
		defer close(msgChan)

		client, err := c.client.Events(ctx, &yavpb.EventsOptions{Filters: filters})
		if err != nil {
			errChan <- err
			return
		}
		for {
			msg, err := client.Recv()
			if err != nil {
				errChan <- err
				return
			}
			msgChan <- types.EventMessage{
				ID:       msg.Id,
				Type:     msg.Type,
				Action:   msg.Action,
				TimeNano: msg.TimeNano,
			}
		}
	}()

	return msgChan, errChan
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

	for name, cidrsJson := range msg.Networks {
		if err := json.Unmarshal(cidrsJson, &cidr); err != nil {
			return nil, err
		}

		network := &types.Network{Name: name}
		network.Subnets = append(network.Subnets, cidr[:]...)
		networks = append(networks, network)
	}

	return networks, nil
}
