// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// CardServiceClient is the client API for CardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardServiceClient interface {
	GetCards(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CardArray, error)
	GetPlayers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlayerArray, error)
	AddCard(ctx context.Context, in *Card, opts ...grpc.CallOption) (*Response, error)
}

type cardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCardServiceClient(cc grpc.ClientConnInterface) CardServiceClient {
	return &cardServiceClient{cc}
}

func (c *cardServiceClient) GetCards(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CardArray, error) {
	out := new(CardArray)
	err := c.cc.Invoke(ctx, "/main.CardService/GetCards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) GetPlayers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlayerArray, error) {
	out := new(PlayerArray)
	err := c.cc.Invoke(ctx, "/main.CardService/GetPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) AddCard(ctx context.Context, in *Card, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/main.CardService/AddCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardServiceServer is the server API for CardService service.
// All implementations must embed UnimplementedCardServiceServer
// for forward compatibility
type CardServiceServer interface {
	GetCards(context.Context, *Empty) (*CardArray, error)
	GetPlayers(context.Context, *Empty) (*PlayerArray, error)
	AddCard(context.Context, *Card) (*Response, error)
	mustEmbedUnimplementedCardServiceServer()
}

// UnimplementedCardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCardServiceServer struct {
}

func (UnimplementedCardServiceServer) GetCards(context.Context, *Empty) (*CardArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCards not implemented")
}
func (UnimplementedCardServiceServer) GetPlayers(context.Context, *Empty) (*PlayerArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayers not implemented")
}
func (UnimplementedCardServiceServer) AddCard(context.Context, *Card) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCard not implemented")
}
func (UnimplementedCardServiceServer) mustEmbedUnimplementedCardServiceServer() {}

// UnsafeCardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardServiceServer will
// result in compilation errors.
type UnsafeCardServiceServer interface {
	mustEmbedUnimplementedCardServiceServer()
}

func RegisterCardServiceServer(s grpc.ServiceRegistrar, srv CardServiceServer) {
	s.RegisterService(&CardService_ServiceDesc, srv)
}

func _CardService_GetCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).GetCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CardService/GetCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).GetCards(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_GetPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).GetPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CardService/GetPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).GetPlayers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_AddCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Card)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).AddCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CardService/AddCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).AddCard(ctx, req.(*Card))
	}
	return interceptor(ctx, in, info, handler)
}

// CardService_ServiceDesc is the grpc.ServiceDesc for CardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.CardService",
	HandlerType: (*CardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCards",
			Handler:    _CardService_GetCards_Handler,
		},
		{
			MethodName: "GetPlayers",
			Handler:    _CardService_GetPlayers_Handler,
		},
		{
			MethodName: "AddCard",
			Handler:    _CardService_AddCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cards.proto",
}
