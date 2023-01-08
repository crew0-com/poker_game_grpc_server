// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: service_poker.proto

package pb

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

// PokerServiceClient is the client API for PokerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PokerServiceClient interface {
	CreateGameRoom(ctx context.Context, in *CreateGameRoomRequest, opts ...grpc.CallOption) (*CreateGameRoomResponse, error)
}

type pokerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPokerServiceClient(cc grpc.ClientConnInterface) PokerServiceClient {
	return &pokerServiceClient{cc}
}

func (c *pokerServiceClient) CreateGameRoom(ctx context.Context, in *CreateGameRoomRequest, opts ...grpc.CallOption) (*CreateGameRoomResponse, error) {
	out := new(CreateGameRoomResponse)
	err := c.cc.Invoke(ctx, "/pb.PokerService/CreateGameRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokerServiceServer is the server API for PokerService service.
// All implementations must embed UnimplementedPokerServiceServer
// for forward compatibility
type PokerServiceServer interface {
	CreateGameRoom(context.Context, *CreateGameRoomRequest) (*CreateGameRoomResponse, error)
	mustEmbedUnimplementedPokerServiceServer()
}

// UnimplementedPokerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPokerServiceServer struct {
}

func (UnimplementedPokerServiceServer) CreateGameRoom(context.Context, *CreateGameRoomRequest) (*CreateGameRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGameRoom not implemented")
}
func (UnimplementedPokerServiceServer) mustEmbedUnimplementedPokerServiceServer() {}

// UnsafePokerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PokerServiceServer will
// result in compilation errors.
type UnsafePokerServiceServer interface {
	mustEmbedUnimplementedPokerServiceServer()
}

func RegisterPokerServiceServer(s grpc.ServiceRegistrar, srv PokerServiceServer) {
	s.RegisterService(&PokerService_ServiceDesc, srv)
}

func _PokerService_CreateGameRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokerServiceServer).CreateGameRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PokerService/CreateGameRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokerServiceServer).CreateGameRoom(ctx, req.(*CreateGameRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PokerService_ServiceDesc is the grpc.ServiceDesc for PokerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PokerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PokerService",
	HandlerType: (*PokerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGameRoom",
			Handler:    _PokerService_CreateGameRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_poker.proto",
}
