// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: outbound.proto

package outbound

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

const (
	OutboundPolicies_Get_FullMethodName   = "/io.linkerd.proxy.outbound.OutboundPolicies/Get"
	OutboundPolicies_Watch_FullMethodName = "/io.linkerd.proxy.outbound.OutboundPolicies/Watch"
)

// OutboundPoliciesClient is the client API for OutboundPolicies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OutboundPoliciesClient interface {
	Get(ctx context.Context, in *TrafficSpec, opts ...grpc.CallOption) (*OutboundPolicy, error)
	Watch(ctx context.Context, in *TrafficSpec, opts ...grpc.CallOption) (OutboundPolicies_WatchClient, error)
}

type outboundPoliciesClient struct {
	cc grpc.ClientConnInterface
}

func NewOutboundPoliciesClient(cc grpc.ClientConnInterface) OutboundPoliciesClient {
	return &outboundPoliciesClient{cc}
}

func (c *outboundPoliciesClient) Get(ctx context.Context, in *TrafficSpec, opts ...grpc.CallOption) (*OutboundPolicy, error) {
	out := new(OutboundPolicy)
	err := c.cc.Invoke(ctx, OutboundPolicies_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboundPoliciesClient) Watch(ctx context.Context, in *TrafficSpec, opts ...grpc.CallOption) (OutboundPolicies_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &OutboundPolicies_ServiceDesc.Streams[0], OutboundPolicies_Watch_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &outboundPoliciesWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OutboundPolicies_WatchClient interface {
	Recv() (*OutboundPolicy, error)
	grpc.ClientStream
}

type outboundPoliciesWatchClient struct {
	grpc.ClientStream
}

func (x *outboundPoliciesWatchClient) Recv() (*OutboundPolicy, error) {
	m := new(OutboundPolicy)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OutboundPoliciesServer is the server API for OutboundPolicies service.
// All implementations must embed UnimplementedOutboundPoliciesServer
// for forward compatibility
type OutboundPoliciesServer interface {
	Get(context.Context, *TrafficSpec) (*OutboundPolicy, error)
	Watch(*TrafficSpec, OutboundPolicies_WatchServer) error
	mustEmbedUnimplementedOutboundPoliciesServer()
}

// UnimplementedOutboundPoliciesServer must be embedded to have forward compatible implementations.
type UnimplementedOutboundPoliciesServer struct {
}

func (UnimplementedOutboundPoliciesServer) Get(context.Context, *TrafficSpec) (*OutboundPolicy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedOutboundPoliciesServer) Watch(*TrafficSpec, OutboundPolicies_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedOutboundPoliciesServer) mustEmbedUnimplementedOutboundPoliciesServer() {}

// UnsafeOutboundPoliciesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OutboundPoliciesServer will
// result in compilation errors.
type UnsafeOutboundPoliciesServer interface {
	mustEmbedUnimplementedOutboundPoliciesServer()
}

func RegisterOutboundPoliciesServer(s grpc.ServiceRegistrar, srv OutboundPoliciesServer) {
	s.RegisterService(&OutboundPolicies_ServiceDesc, srv)
}

func _OutboundPolicies_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrafficSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboundPoliciesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OutboundPolicies_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboundPoliciesServer).Get(ctx, req.(*TrafficSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboundPolicies_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TrafficSpec)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OutboundPoliciesServer).Watch(m, &outboundPoliciesWatchServer{stream})
}

type OutboundPolicies_WatchServer interface {
	Send(*OutboundPolicy) error
	grpc.ServerStream
}

type outboundPoliciesWatchServer struct {
	grpc.ServerStream
}

func (x *outboundPoliciesWatchServer) Send(m *OutboundPolicy) error {
	return x.ServerStream.SendMsg(m)
}

// OutboundPolicies_ServiceDesc is the grpc.ServiceDesc for OutboundPolicies service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OutboundPolicies_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.linkerd.proxy.outbound.OutboundPolicies",
	HandlerType: (*OutboundPoliciesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _OutboundPolicies_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _OutboundPolicies_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "outbound.proto",
}
