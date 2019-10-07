package grpcclient

import (
	"context"

	yavpb "github.com/projecteru2/libyavirt/grpc/gen"
	"github.com/projecteru2/libyavirt/types"
	"google.golang.org/grpc"
)

type grpcClient struct {
	client yavpb.YavirtdRPCClient
}

func New(addr string) (*grpcClient, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	client := yavpb.NewYavirtdRPCClient(conn)
	return &grpcClient{client}, nil
}

func (c *grpcClient) Info(ctx context.Context) (info types.HostInfo, err error) {
	msg, err := c.client.GetInfo(ctx, &yavpb.Empty{})
	if err != nil {
		return
	}

	return types.HostInfo{
		ID:      msg.Id,
		Cpu:     int(msg.Cpu),
		Mem:     msg.Memory,
		Storage: msg.Memory,
	}, nil
}

func (c *grpcClient) GetGuest(ctx context.Context, ID string) (guest types.Guest, err error) {
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
		Cpu:       int(msg.Cpu),
		Mem:       msg.Memory,
		Storage:   msg.Storage,
		ImageID:   msg.ImageId,
		ImageName: msg.ImageName,
		Networks:  msg.Networks,
	}, nil
}

func (c *grpcClient) CreateGuest(ctx context.Context, args types.CreateGuestReq) (guest types.Guest, err error) {
	opts := &yavpb.CreateGuestOptions{
		Cpu:       int64(args.Cpu),
		Memory:    args.Mem,
		ImageName: args.ImageName,
		Volumes:   args.Volumes,
	}
	msg, err := c.client.CreateGuest(ctx, opts)
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
		Cpu:       int(msg.Cpu),
		Mem:       msg.Memory,
		Storage:   msg.Storage,
		ImageID:   msg.ImageId,
		ImageName: msg.ImageName,
		Networks:  msg.Networks,
	}, nil
}

func (c *grpcClient) StartGuest(ctx context.Context, ID string) (msg types.Msg, err error) {
	return c.controlGuest(ctx, ID, "start")
}

func (c *grpcClient) StopGuest(ctx context.Context, ID string) (msg types.Msg, err error) {
	return c.controlGuest(ctx, ID, "stop")
}

func (c *grpcClient) DestroyGuest(ctx context.Context, ID string) (msg types.Msg, err error) {
	return c.controlGuest(ctx, ID, "destroy")
}

func (c *grpcClient) controlGuest(ctx context.Context, ID, operation string) (msg types.Msg, err error) {
	opts := &yavpb.ControlGuestOptions{
		Id:        ID,
		Operation: operation,
	}
	m, err := c.client.ControlGuest(ctx, opts)
	if err != nil {
		return
	}

	return types.Msg{Msg: m.Msg}, nil
}

func (c *grpcClient) ExecuteGuest(ctx context.Context, ID string, cmd []string) (msg types.ExecuteGuestMessage, err error) {
	opts := &yavpb.ExecuteGuestOptions{
		Id:       ID,
		Commands: cmd,
	}
	m, err := c.client.ExecuteGuest(ctx, opts)
	if err != nil {
		return
	}
	return types.ExecuteGuestMessage{
		ID:       ID,
		Data:     m.Data,
		ExitCode: int(m.ExitCode),
	}, nil
}
