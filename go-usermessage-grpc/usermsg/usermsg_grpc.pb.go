// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_usermessage_grpc

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

// UserMessageClient is the client API for UserMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserMessageClient interface {
	CreateNewMessage(ctx context.Context, in *NewMessage, opts ...grpc.CallOption) (*Message, error)
	GetMessages(ctx context.Context, in *GetMessageParams, opts ...grpc.CallOption) (*MessageList, error)
}

type userMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewUserMessageClient(cc grpc.ClientConnInterface) UserMessageClient {
	return &userMessageClient{cc}
}

func (c *userMessageClient) CreateNewMessage(ctx context.Context, in *NewMessage, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/usermsg.UserMessage/CreateNewMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMessageClient) GetMessages(ctx context.Context, in *GetMessageParams, opts ...grpc.CallOption) (*MessageList, error) {
	out := new(MessageList)
	err := c.cc.Invoke(ctx, "/usermsg.UserMessage/GetMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserMessageServer is the server API for UserMessage service.
// All implementations must embed UnimplementedUserMessageServer
// for forward compatibility
type UserMessageServer interface {
	CreateNewMessage(context.Context, *NewMessage) (*Message, error)
	GetMessages(context.Context, *GetMessageParams) (*MessageList, error)
	mustEmbedUnimplementedUserMessageServer()
}

// UnimplementedUserMessageServer must be embedded to have forward compatible implementations.
type UnimplementedUserMessageServer struct {
}

func (UnimplementedUserMessageServer) CreateNewMessage(context.Context, *NewMessage) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewMessage not implemented")
}
func (UnimplementedUserMessageServer) GetMessages(context.Context, *GetMessageParams) (*MessageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedUserMessageServer) mustEmbedUnimplementedUserMessageServer() {}

// UnsafeUserMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserMessageServer will
// result in compilation errors.
type UnsafeUserMessageServer interface {
	mustEmbedUnimplementedUserMessageServer()
}

func RegisterUserMessageServer(s grpc.ServiceRegistrar, srv UserMessageServer) {
	s.RegisterService(&UserMessage_ServiceDesc, srv)
}

func _UserMessage_CreateNewMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMessageServer).CreateNewMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermsg.UserMessage/CreateNewMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMessageServer).CreateNewMessage(ctx, req.(*NewMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMessage_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMessageServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermsg.UserMessage/GetMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMessageServer).GetMessages(ctx, req.(*GetMessageParams))
	}
	return interceptor(ctx, in, info, handler)
}

// UserMessage_ServiceDesc is the grpc.ServiceDesc for UserMessage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserMessage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usermsg.UserMessage",
	HandlerType: (*UserMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewMessage",
			Handler:    _UserMessage_CreateNewMessage_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _UserMessage_GetMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usermsg/usermsg.proto",
}
