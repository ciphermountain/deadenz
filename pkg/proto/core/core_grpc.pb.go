// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: pkg/proto/core/core.proto

package core

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

// DeadenzClient is the client API for Deadenz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeadenzClient interface {
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
	Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*Response, error)
	Assets(ctx context.Context, in *AssetRequest, opts ...grpc.CallOption) (*AssetResponse, error)
}

type deadenzClient struct {
	cc grpc.ClientConnInterface
}

func NewDeadenzClient(cc grpc.ClientConnInterface) DeadenzClient {
	return &deadenzClient{cc}
}

func (c *deadenzClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := c.cc.Invoke(ctx, "/core.Deadenz/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deadenzClient) Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/core.Deadenz/Load", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deadenzClient) Assets(ctx context.Context, in *AssetRequest, opts ...grpc.CallOption) (*AssetResponse, error) {
	out := new(AssetResponse)
	err := c.cc.Invoke(ctx, "/core.Deadenz/Assets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeadenzServer is the server API for Deadenz service.
// All implementations must embed UnimplementedDeadenzServer
// for forward compatibility
type DeadenzServer interface {
	Run(context.Context, *RunRequest) (*RunResponse, error)
	Load(context.Context, *LoadRequest) (*Response, error)
	Assets(context.Context, *AssetRequest) (*AssetResponse, error)
	mustEmbedUnimplementedDeadenzServer()
}

// UnimplementedDeadenzServer must be embedded to have forward compatible implementations.
type UnimplementedDeadenzServer struct {
}

func (UnimplementedDeadenzServer) Run(context.Context, *RunRequest) (*RunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedDeadenzServer) Load(context.Context, *LoadRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (UnimplementedDeadenzServer) Assets(context.Context, *AssetRequest) (*AssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Assets not implemented")
}
func (UnimplementedDeadenzServer) mustEmbedUnimplementedDeadenzServer() {}

// UnsafeDeadenzServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeadenzServer will
// result in compilation errors.
type UnsafeDeadenzServer interface {
	mustEmbedUnimplementedDeadenzServer()
}

func RegisterDeadenzServer(s grpc.ServiceRegistrar, srv DeadenzServer) {
	s.RegisterService(&Deadenz_ServiceDesc, srv)
}

func _Deadenz_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeadenzServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.Deadenz/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeadenzServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deadenz_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeadenzServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.Deadenz/Load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeadenzServer).Load(ctx, req.(*LoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deadenz_Assets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeadenzServer).Assets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.Deadenz/Assets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeadenzServer).Assets(ctx, req.(*AssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Deadenz_ServiceDesc is the grpc.ServiceDesc for Deadenz service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Deadenz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.Deadenz",
	HandlerType: (*DeadenzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _Deadenz_Run_Handler,
		},
		{
			MethodName: "Load",
			Handler:    _Deadenz_Load_Handler,
		},
		{
			MethodName: "Assets",
			Handler:    _Deadenz_Assets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/core/core.proto",
}