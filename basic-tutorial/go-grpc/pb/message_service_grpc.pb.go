// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/message_service.proto

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

// SendMessageClient is the client API for SendMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendMessageClient interface {
	RequestMessage(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type sendMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewSendMessageClient(cc grpc.ClientConnInterface) SendMessageClient {
	return &sendMessageClient{cc}
}

func (c *sendMessageClient) RequestMessage(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/SendMessage/RequestMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendMessageServer is the server API for SendMessage service.
// All implementations must embed UnimplementedSendMessageServer
// for forward compatibility
type SendMessageServer interface {
	RequestMessage(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedSendMessageServer()
}

// UnimplementedSendMessageServer must be embedded to have forward compatible implementations.
type UnimplementedSendMessageServer struct {
}

func (UnimplementedSendMessageServer) RequestMessage(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestMessage not implemented")
}
func (UnimplementedSendMessageServer) mustEmbedUnimplementedSendMessageServer() {}

// UnsafeSendMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendMessageServer will
// result in compilation errors.
type UnsafeSendMessageServer interface {
	mustEmbedUnimplementedSendMessageServer()
}

func RegisterSendMessageServer(s grpc.ServiceRegistrar, srv SendMessageServer) {
	s.RegisterService(&SendMessage_ServiceDesc, srv)
}

func _SendMessage_RequestMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendMessageServer).RequestMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SendMessage/RequestMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendMessageServer).RequestMessage(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// SendMessage_ServiceDesc is the grpc.ServiceDesc for SendMessage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendMessage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SendMessage",
	HandlerType: (*SendMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestMessage",
			Handler:    _SendMessage_RequestMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/message_service.proto",
}
