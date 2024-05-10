// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: destination.proto

package destination

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
	Destination_Get_FullMethodName        = "/io.linkerd.proxy.destination.Destination/Get"
	Destination_GetProfile_FullMethodName = "/io.linkerd.proxy.destination.Destination/GetProfile"
)

// DestinationClient is the client API for Destination service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DestinationClient interface {
	// Given a destination, return all addresses in that destination as a long-
	// running stream of updates.
	Get(ctx context.Context, in *GetDestination, opts ...grpc.CallOption) (Destination_GetClient, error)
	// Given a destination, return that destination's profile and send an update
	// whenever it changes.
	GetProfile(ctx context.Context, in *GetDestination, opts ...grpc.CallOption) (Destination_GetProfileClient, error)
}

type destinationClient struct {
	cc grpc.ClientConnInterface
}

func NewDestinationClient(cc grpc.ClientConnInterface) DestinationClient {
	return &destinationClient{cc}
}

func (c *destinationClient) Get(ctx context.Context, in *GetDestination, opts ...grpc.CallOption) (Destination_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &Destination_ServiceDesc.Streams[0], Destination_Get_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &destinationGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Destination_GetClient interface {
	Recv() (*Update, error)
	grpc.ClientStream
}

type destinationGetClient struct {
	grpc.ClientStream
}

func (x *destinationGetClient) Recv() (*Update, error) {
	m := new(Update)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *destinationClient) GetProfile(ctx context.Context, in *GetDestination, opts ...grpc.CallOption) (Destination_GetProfileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Destination_ServiceDesc.Streams[1], Destination_GetProfile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &destinationGetProfileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Destination_GetProfileClient interface {
	Recv() (*DestinationProfile, error)
	grpc.ClientStream
}

type destinationGetProfileClient struct {
	grpc.ClientStream
}

func (x *destinationGetProfileClient) Recv() (*DestinationProfile, error) {
	m := new(DestinationProfile)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DestinationServer is the server API for Destination service.
// All implementations must embed UnimplementedDestinationServer
// for forward compatibility
type DestinationServer interface {
	// Given a destination, return all addresses in that destination as a long-
	// running stream of updates.
	Get(*GetDestination, Destination_GetServer) error
	// Given a destination, return that destination's profile and send an update
	// whenever it changes.
	GetProfile(*GetDestination, Destination_GetProfileServer) error
	mustEmbedUnimplementedDestinationServer()
}

// UnimplementedDestinationServer must be embedded to have forward compatible implementations.
type UnimplementedDestinationServer struct {
}

func (UnimplementedDestinationServer) Get(*GetDestination, Destination_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDestinationServer) GetProfile(*GetDestination, Destination_GetProfileServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedDestinationServer) mustEmbedUnimplementedDestinationServer() {}

// UnsafeDestinationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DestinationServer will
// result in compilation errors.
type UnsafeDestinationServer interface {
	mustEmbedUnimplementedDestinationServer()
}

func RegisterDestinationServer(s grpc.ServiceRegistrar, srv DestinationServer) {
	s.RegisterService(&Destination_ServiceDesc, srv)
}

func _Destination_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDestination)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DestinationServer).Get(m, &destinationGetServer{stream})
}

type Destination_GetServer interface {
	Send(*Update) error
	grpc.ServerStream
}

type destinationGetServer struct {
	grpc.ServerStream
}

func (x *destinationGetServer) Send(m *Update) error {
	return x.ServerStream.SendMsg(m)
}

func _Destination_GetProfile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDestination)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DestinationServer).GetProfile(m, &destinationGetProfileServer{stream})
}

type Destination_GetProfileServer interface {
	Send(*DestinationProfile) error
	grpc.ServerStream
}

type destinationGetProfileServer struct {
	grpc.ServerStream
}

func (x *destinationGetProfileServer) Send(m *DestinationProfile) error {
	return x.ServerStream.SendMsg(m)
}

// Destination_ServiceDesc is the grpc.ServiceDesc for Destination service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Destination_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.linkerd.proxy.destination.Destination",
	HandlerType: (*DestinationServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _Destination_Get_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetProfile",
			Handler:       _Destination_GetProfile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "destination.proto",
}
