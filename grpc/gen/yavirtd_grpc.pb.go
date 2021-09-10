// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package yavpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// YavirtdRPCClient is the client API for YavirtdRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YavirtdRPCClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingMessage, error)
	GetInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*InfoMessage, error)
	GetGuest(ctx context.Context, in *GetGuestOptions, opts ...grpc.CallOption) (*GetGuestMessage, error)
	GetGuestUUID(ctx context.Context, in *GetGuestOptions, opts ...grpc.CallOption) (*GetGuestUUIDMessage, error)
	CreateGuest(ctx context.Context, in *CreateGuestOptions, opts ...grpc.CallOption) (*CreateGuestMessage, error)
	ControlGuest(ctx context.Context, in *ControlGuestOptions, opts ...grpc.CallOption) (*ControlGuestMessage, error)
	AttachGuest(ctx context.Context, opts ...grpc.CallOption) (YavirtdRPC_AttachGuestClient, error)
	ResizeConsoleWindow(ctx context.Context, in *ResizeWindowOptions, opts ...grpc.CallOption) (*Empty, error)
	ExecuteGuest(ctx context.Context, in *ExecuteGuestOptions, opts ...grpc.CallOption) (*ExecuteGuestMessage, error)
	ResizeGuest(ctx context.Context, in *ResizeGuestOptions, opts ...grpc.CallOption) (*ControlGuestMessage, error)
	CaptureGuest(ctx context.Context, in *CaptureGuestOptions, opts ...grpc.CallOption) (*UserImageMessage, error)
	ConnectNetwork(ctx context.Context, in *ConnectNetworkOptions, opts ...grpc.CallOption) (*ConnectNetworkMessage, error)
	DisconnectNetwork(ctx context.Context, in *DisconnectNetworkOptions, opts ...grpc.CallOption) (*DisconnectNetworkMessage, error)
	Cat(ctx context.Context, in *CatOptions, opts ...grpc.CallOption) (YavirtdRPC_CatClient, error)
	CopyToGuest(ctx context.Context, opts ...grpc.CallOption) (YavirtdRPC_CopyToGuestClient, error)
}

type yavirtdRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewYavirtdRPCClient(cc grpc.ClientConnInterface) YavirtdRPCClient {
	return &yavirtdRPCClient{cc}
}

func (c *yavirtdRPCClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingMessage, error) {
	out := new(PingMessage)
	err := c.cc.Invoke(ctx, "/yavpb.YavirtdRPC/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yavirtdRPCClient) GetInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*InfoMessage, error) {
	out := new(InfoMessage)
	err := c.cc.Invoke(ctx, "/yavpb.YavirtdRPC/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yavirtdRPCClient) GetGuest(ctx context.Context, in *GetGuestOptions, opts ...grpc.CallOption) (*GetGuestMessage, error) {
	out := new(GetGuestMessage)
	err := c.cc.Invoke(ctx, "/yavpb.YavirtdRPC/GetGuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yavirtdRPCClient) GetGuestUUID(ctx context.Context, in *GetGuestOptions, opts ...grpc.CallOption) (*GetGuestUUIDMessage, error) {
	out := new(GetGuestUUIDMessage)
	err := c.cc.Invoke(ctx, "/yavpb.YavirtdRPC/GetGuestUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yavirtdRPCClient) CreateGuest(ctx context.Context, in *CreateGuestOptions, opts ...grpc.CallOption) (*CreateGuestMessage, error) {
	out := new(CreateGuestMessage)
	err := c.cc.Invoke(ctx, "/yavpb.YavirtdRPC/CreateGuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yavirtdRPCClient) ControlGuest(ctx context.Context, in *ControlGuestOptions, opts ...grpc.CallOption) (*ControlGuestMessage, error) {
	out := new(ControlGuestMessage)
	err := c.cc.Invoke(ctx, "/yavpb.YavirtdRPC/ControlGuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yavirtdRPCClient) AttachGuest(ctx context.Context, opts ...grpc.CallOption) (YavirtdRPC_AttachGuestClient, error) {
	stream, err := c.cc.NewStream(ctx, &YavirtdRPC_ServiceDesc.Streams[0], "/yavpb.YavirtdRPC/AttachGuest", opts...)
	if err != nil {
		return nil, err
	}
	x := &yavirtdRPCAttachGuestClient{stream}
	return x, nil
}

type YavirtdRPC_AttachGuestClient interface {
	Send(*AttachGuestOptions) error
	Recv() (*AttachGuestMessage, error)
	grpc.ClientStream
}

type yavirtdRPCAttachGuestClient struct {
	grpc.ClientStream
}

func (x *yavirtdRPCAttachGuestClient) Send(m *AttachGuestOptions) error {
	return x.ClientStream.SendMsg(m)
}

func (x *yavirtdRPCAttachGuestClient) Recv() (*AttachGuestMessage, error) {
	m := new(AttachGuestMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *yavirtdRPCClient) ResizeConsoleWindow(ctx context.Context, in *ResizeWindowOptions, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/yavpb.YavirtdRPC/ResizeConsoleWindow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yavirtdRPCClient) ExecuteGuest(ctx context.Context, in *ExecuteGuestOptions, opts ...grpc.CallOption) (*ExecuteGuestMessage, error) {
	out := new(ExecuteGuestMessage)
	err := c.cc.Invoke(ctx, "/yavpb.YavirtdRPC/ExecuteGuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yavirtdRPCClient) ResizeGuest(ctx context.Context, in *ResizeGuestOptions, opts ...grpc.CallOption) (*ControlGuestMessage, error) {
	out := new(ControlGuestMessage)
	err := c.cc.Invoke(ctx, "/yavpb.YavirtdRPC/ResizeGuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yavirtdRPCClient) CaptureGuest(ctx context.Context, in *CaptureGuestOptions, opts ...grpc.CallOption) (*UserImageMessage, error) {
	out := new(UserImageMessage)
	err := c.cc.Invoke(ctx, "/yavpb.YavirtdRPC/CaptureGuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yavirtdRPCClient) ConnectNetwork(ctx context.Context, in *ConnectNetworkOptions, opts ...grpc.CallOption) (*ConnectNetworkMessage, error) {
	out := new(ConnectNetworkMessage)
	err := c.cc.Invoke(ctx, "/yavpb.YavirtdRPC/ConnectNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yavirtdRPCClient) DisconnectNetwork(ctx context.Context, in *DisconnectNetworkOptions, opts ...grpc.CallOption) (*DisconnectNetworkMessage, error) {
	out := new(DisconnectNetworkMessage)
	err := c.cc.Invoke(ctx, "/yavpb.YavirtdRPC/DisconnectNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yavirtdRPCClient) Cat(ctx context.Context, in *CatOptions, opts ...grpc.CallOption) (YavirtdRPC_CatClient, error) {
	stream, err := c.cc.NewStream(ctx, &YavirtdRPC_ServiceDesc.Streams[1], "/yavpb.YavirtdRPC/Cat", opts...)
	if err != nil {
		return nil, err
	}
	x := &yavirtdRPCCatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type YavirtdRPC_CatClient interface {
	Recv() (*CatMessage, error)
	grpc.ClientStream
}

type yavirtdRPCCatClient struct {
	grpc.ClientStream
}

func (x *yavirtdRPCCatClient) Recv() (*CatMessage, error) {
	m := new(CatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *yavirtdRPCClient) CopyToGuest(ctx context.Context, opts ...grpc.CallOption) (YavirtdRPC_CopyToGuestClient, error) {
	stream, err := c.cc.NewStream(ctx, &YavirtdRPC_ServiceDesc.Streams[2], "/yavpb.YavirtdRPC/CopyToGuest", opts...)
	if err != nil {
		return nil, err
	}
	x := &yavirtdRPCCopyToGuestClient{stream}
	return x, nil
}

type YavirtdRPC_CopyToGuestClient interface {
	Send(*CopyOptions) error
	CloseAndRecv() (*CopyMessage, error)
	grpc.ClientStream
}

type yavirtdRPCCopyToGuestClient struct {
	grpc.ClientStream
}

func (x *yavirtdRPCCopyToGuestClient) Send(m *CopyOptions) error {
	return x.ClientStream.SendMsg(m)
}

func (x *yavirtdRPCCopyToGuestClient) CloseAndRecv() (*CopyMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CopyMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// YavirtdRPCServer is the server API for YavirtdRPC service.
// All implementations must embed UnimplementedYavirtdRPCServer
// for forward compatibility
type YavirtdRPCServer interface {
	Ping(context.Context, *Empty) (*PingMessage, error)
	GetInfo(context.Context, *Empty) (*InfoMessage, error)
	GetGuest(context.Context, *GetGuestOptions) (*GetGuestMessage, error)
	GetGuestUUID(context.Context, *GetGuestOptions) (*GetGuestUUIDMessage, error)
	CreateGuest(context.Context, *CreateGuestOptions) (*CreateGuestMessage, error)
	ControlGuest(context.Context, *ControlGuestOptions) (*ControlGuestMessage, error)
	AttachGuest(YavirtdRPC_AttachGuestServer) error
	ResizeConsoleWindow(context.Context, *ResizeWindowOptions) (*Empty, error)
	ExecuteGuest(context.Context, *ExecuteGuestOptions) (*ExecuteGuestMessage, error)
	ResizeGuest(context.Context, *ResizeGuestOptions) (*ControlGuestMessage, error)
	CaptureGuest(context.Context, *CaptureGuestOptions) (*UserImageMessage, error)
	ConnectNetwork(context.Context, *ConnectNetworkOptions) (*ConnectNetworkMessage, error)
	DisconnectNetwork(context.Context, *DisconnectNetworkOptions) (*DisconnectNetworkMessage, error)
	Cat(*CatOptions, YavirtdRPC_CatServer) error
	CopyToGuest(YavirtdRPC_CopyToGuestServer) error
	mustEmbedUnimplementedYavirtdRPCServer()
}

// UnimplementedYavirtdRPCServer must be embedded to have forward compatible implementations.
type UnimplementedYavirtdRPCServer struct {
}

func (UnimplementedYavirtdRPCServer) Ping(context.Context, *Empty) (*PingMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedYavirtdRPCServer) GetInfo(context.Context, *Empty) (*InfoMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedYavirtdRPCServer) GetGuest(context.Context, *GetGuestOptions) (*GetGuestMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuest not implemented")
}
func (UnimplementedYavirtdRPCServer) GetGuestUUID(context.Context, *GetGuestOptions) (*GetGuestUUIDMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuestUUID not implemented")
}
func (UnimplementedYavirtdRPCServer) CreateGuest(context.Context, *CreateGuestOptions) (*CreateGuestMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGuest not implemented")
}
func (UnimplementedYavirtdRPCServer) ControlGuest(context.Context, *ControlGuestOptions) (*ControlGuestMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ControlGuest not implemented")
}
func (UnimplementedYavirtdRPCServer) AttachGuest(YavirtdRPC_AttachGuestServer) error {
	return status.Errorf(codes.Unimplemented, "method AttachGuest not implemented")
}
func (UnimplementedYavirtdRPCServer) ResizeConsoleWindow(context.Context, *ResizeWindowOptions) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResizeConsoleWindow not implemented")
}
func (UnimplementedYavirtdRPCServer) ExecuteGuest(context.Context, *ExecuteGuestOptions) (*ExecuteGuestMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteGuest not implemented")
}
func (UnimplementedYavirtdRPCServer) ResizeGuest(context.Context, *ResizeGuestOptions) (*ControlGuestMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResizeGuest not implemented")
}
func (UnimplementedYavirtdRPCServer) CaptureGuest(context.Context, *CaptureGuestOptions) (*UserImageMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CaptureGuest not implemented")
}
func (UnimplementedYavirtdRPCServer) ConnectNetwork(context.Context, *ConnectNetworkOptions) (*ConnectNetworkMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectNetwork not implemented")
}
func (UnimplementedYavirtdRPCServer) DisconnectNetwork(context.Context, *DisconnectNetworkOptions) (*DisconnectNetworkMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisconnectNetwork not implemented")
}
func (UnimplementedYavirtdRPCServer) Cat(*CatOptions, YavirtdRPC_CatServer) error {
	return status.Errorf(codes.Unimplemented, "method Cat not implemented")
}
func (UnimplementedYavirtdRPCServer) CopyToGuest(YavirtdRPC_CopyToGuestServer) error {
	return status.Errorf(codes.Unimplemented, "method CopyToGuest not implemented")
}
func (UnimplementedYavirtdRPCServer) mustEmbedUnimplementedYavirtdRPCServer() {}

// UnsafeYavirtdRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YavirtdRPCServer will
// result in compilation errors.
type UnsafeYavirtdRPCServer interface {
	mustEmbedUnimplementedYavirtdRPCServer()
}

func RegisterYavirtdRPCServer(s grpc.ServiceRegistrar, srv YavirtdRPCServer) {
	s.RegisterService(&YavirtdRPC_ServiceDesc, srv)
}

func _YavirtdRPC_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YavirtdRPCServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yavpb.YavirtdRPC/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YavirtdRPCServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _YavirtdRPC_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YavirtdRPCServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yavpb.YavirtdRPC/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YavirtdRPCServer).GetInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _YavirtdRPC_GetGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGuestOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YavirtdRPCServer).GetGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yavpb.YavirtdRPC/GetGuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YavirtdRPCServer).GetGuest(ctx, req.(*GetGuestOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _YavirtdRPC_GetGuestUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGuestOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YavirtdRPCServer).GetGuestUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yavpb.YavirtdRPC/GetGuestUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YavirtdRPCServer).GetGuestUUID(ctx, req.(*GetGuestOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _YavirtdRPC_CreateGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGuestOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YavirtdRPCServer).CreateGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yavpb.YavirtdRPC/CreateGuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YavirtdRPCServer).CreateGuest(ctx, req.(*CreateGuestOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _YavirtdRPC_ControlGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlGuestOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YavirtdRPCServer).ControlGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yavpb.YavirtdRPC/ControlGuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YavirtdRPCServer).ControlGuest(ctx, req.(*ControlGuestOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _YavirtdRPC_AttachGuest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(YavirtdRPCServer).AttachGuest(&yavirtdRPCAttachGuestServer{stream})
}

type YavirtdRPC_AttachGuestServer interface {
	Send(*AttachGuestMessage) error
	Recv() (*AttachGuestOptions, error)
	grpc.ServerStream
}

type yavirtdRPCAttachGuestServer struct {
	grpc.ServerStream
}

func (x *yavirtdRPCAttachGuestServer) Send(m *AttachGuestMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *yavirtdRPCAttachGuestServer) Recv() (*AttachGuestOptions, error) {
	m := new(AttachGuestOptions)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _YavirtdRPC_ResizeConsoleWindow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeWindowOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YavirtdRPCServer).ResizeConsoleWindow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yavpb.YavirtdRPC/ResizeConsoleWindow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YavirtdRPCServer).ResizeConsoleWindow(ctx, req.(*ResizeWindowOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _YavirtdRPC_ExecuteGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteGuestOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YavirtdRPCServer).ExecuteGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yavpb.YavirtdRPC/ExecuteGuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YavirtdRPCServer).ExecuteGuest(ctx, req.(*ExecuteGuestOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _YavirtdRPC_ResizeGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeGuestOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YavirtdRPCServer).ResizeGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yavpb.YavirtdRPC/ResizeGuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YavirtdRPCServer).ResizeGuest(ctx, req.(*ResizeGuestOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _YavirtdRPC_CaptureGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaptureGuestOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YavirtdRPCServer).CaptureGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yavpb.YavirtdRPC/CaptureGuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YavirtdRPCServer).CaptureGuest(ctx, req.(*CaptureGuestOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _YavirtdRPC_ConnectNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectNetworkOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YavirtdRPCServer).ConnectNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yavpb.YavirtdRPC/ConnectNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YavirtdRPCServer).ConnectNetwork(ctx, req.(*ConnectNetworkOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _YavirtdRPC_DisconnectNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectNetworkOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YavirtdRPCServer).DisconnectNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yavpb.YavirtdRPC/DisconnectNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YavirtdRPCServer).DisconnectNetwork(ctx, req.(*DisconnectNetworkOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _YavirtdRPC_Cat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CatOptions)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(YavirtdRPCServer).Cat(m, &yavirtdRPCCatServer{stream})
}

type YavirtdRPC_CatServer interface {
	Send(*CatMessage) error
	grpc.ServerStream
}

type yavirtdRPCCatServer struct {
	grpc.ServerStream
}

func (x *yavirtdRPCCatServer) Send(m *CatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _YavirtdRPC_CopyToGuest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(YavirtdRPCServer).CopyToGuest(&yavirtdRPCCopyToGuestServer{stream})
}

type YavirtdRPC_CopyToGuestServer interface {
	SendAndClose(*CopyMessage) error
	Recv() (*CopyOptions, error)
	grpc.ServerStream
}

type yavirtdRPCCopyToGuestServer struct {
	grpc.ServerStream
}

func (x *yavirtdRPCCopyToGuestServer) SendAndClose(m *CopyMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *yavirtdRPCCopyToGuestServer) Recv() (*CopyOptions, error) {
	m := new(CopyOptions)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// YavirtdRPC_ServiceDesc is the grpc.ServiceDesc for YavirtdRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YavirtdRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yavpb.YavirtdRPC",
	HandlerType: (*YavirtdRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _YavirtdRPC_Ping_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _YavirtdRPC_GetInfo_Handler,
		},
		{
			MethodName: "GetGuest",
			Handler:    _YavirtdRPC_GetGuest_Handler,
		},
		{
			MethodName: "GetGuestUUID",
			Handler:    _YavirtdRPC_GetGuestUUID_Handler,
		},
		{
			MethodName: "CreateGuest",
			Handler:    _YavirtdRPC_CreateGuest_Handler,
		},
		{
			MethodName: "ControlGuest",
			Handler:    _YavirtdRPC_ControlGuest_Handler,
		},
		{
			MethodName: "ResizeConsoleWindow",
			Handler:    _YavirtdRPC_ResizeConsoleWindow_Handler,
		},
		{
			MethodName: "ExecuteGuest",
			Handler:    _YavirtdRPC_ExecuteGuest_Handler,
		},
		{
			MethodName: "ResizeGuest",
			Handler:    _YavirtdRPC_ResizeGuest_Handler,
		},
		{
			MethodName: "CaptureGuest",
			Handler:    _YavirtdRPC_CaptureGuest_Handler,
		},
		{
			MethodName: "ConnectNetwork",
			Handler:    _YavirtdRPC_ConnectNetwork_Handler,
		},
		{
			MethodName: "DisconnectNetwork",
			Handler:    _YavirtdRPC_DisconnectNetwork_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AttachGuest",
			Handler:       _YavirtdRPC_AttachGuest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Cat",
			Handler:       _YavirtdRPC_Cat_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CopyToGuest",
			Handler:       _YavirtdRPC_CopyToGuest_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/gen/yavirtd.proto",
}
