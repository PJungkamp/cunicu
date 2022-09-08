// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: rpc/signaling.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	proto "github.com/stv0g/cunicu/pkg/proto"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SignalingClient is the client API for Signaling service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignalingClient interface {
	GetSignalingMessage(ctx context.Context, in *GetSignalingMessageParams, opts ...grpc.CallOption) (*GetSignalingMessageResp, error)
	PutSignalingMessage(ctx context.Context, in *PutSignalingMessageParams, opts ...grpc.CallOption) (*proto.Empty, error)
}

type signalingClient struct {
	cc grpc.ClientConnInterface
}

func NewSignalingClient(cc grpc.ClientConnInterface) SignalingClient {
	return &signalingClient{cc}
}

func (c *signalingClient) GetSignalingMessage(ctx context.Context, in *GetSignalingMessageParams, opts ...grpc.CallOption) (*GetSignalingMessageResp, error) {
	out := new(GetSignalingMessageResp)
	err := c.cc.Invoke(ctx, "/wice.rpc.Signaling/GetSignalingMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalingClient) PutSignalingMessage(ctx context.Context, in *PutSignalingMessageParams, opts ...grpc.CallOption) (*proto.Empty, error) {
	out := new(proto.Empty)
	err := c.cc.Invoke(ctx, "/wice.rpc.Signaling/PutSignalingMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignalingServer is the server API for Signaling service.
// All implementations must embed UnimplementedSignalingServer
// for forward compatibility
type SignalingServer interface {
	GetSignalingMessage(context.Context, *GetSignalingMessageParams) (*GetSignalingMessageResp, error)
	PutSignalingMessage(context.Context, *PutSignalingMessageParams) (*proto.Empty, error)
	mustEmbedUnimplementedSignalingServer()
}

// UnimplementedSignalingServer must be embedded to have forward compatible implementations.
type UnimplementedSignalingServer struct {
}

func (UnimplementedSignalingServer) GetSignalingMessage(context.Context, *GetSignalingMessageParams) (*GetSignalingMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignalingMessage not implemented")
}
func (UnimplementedSignalingServer) PutSignalingMessage(context.Context, *PutSignalingMessageParams) (*proto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutSignalingMessage not implemented")
}
func (UnimplementedSignalingServer) mustEmbedUnimplementedSignalingServer() {}

// UnsafeSignalingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignalingServer will
// result in compilation errors.
type UnsafeSignalingServer interface {
	mustEmbedUnimplementedSignalingServer()
}

func RegisterSignalingServer(s grpc.ServiceRegistrar, srv SignalingServer) {
	s.RegisterService(&Signaling_ServiceDesc, srv)
}

func _Signaling_GetSignalingMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignalingMessageParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalingServer).GetSignalingMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.rpc.Signaling/GetSignalingMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalingServer).GetSignalingMessage(ctx, req.(*GetSignalingMessageParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signaling_PutSignalingMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutSignalingMessageParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalingServer).PutSignalingMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.rpc.Signaling/PutSignalingMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalingServer).PutSignalingMessage(ctx, req.(*PutSignalingMessageParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Signaling_ServiceDesc is the grpc.ServiceDesc for Signaling service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Signaling_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wice.rpc.Signaling",
	HandlerType: (*SignalingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSignalingMessage",
			Handler:    _Signaling_GetSignalingMessage_Handler,
		},
		{
			MethodName: "PutSignalingMessage",
			Handler:    _Signaling_PutSignalingMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/signaling.proto",
}
