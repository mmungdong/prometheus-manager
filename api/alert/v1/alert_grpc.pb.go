// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: alert/v1/alert.proto

package v1

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

// AlertClient is the client API for Alert service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertClient interface {
	Webhook(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*WebhookReply, error)
}

type alertClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertClient(cc grpc.ClientConnInterface) AlertClient {
	return &alertClient{cc}
}

func (c *alertClient) Webhook(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*WebhookReply, error) {
	out := new(WebhookReply)
	err := c.cc.Invoke(ctx, "/api.alert.v1.Alert/Webhook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertServer is the server API for Alert service.
// All implementations must embed UnimplementedAlertServer
// for forward compatibility
type AlertServer interface {
	Webhook(context.Context, *WebhookRequest) (*WebhookReply, error)
	mustEmbedUnimplementedAlertServer()
}

// UnimplementedAlertServer must be embedded to have forward compatible implementations.
type UnimplementedAlertServer struct {
}

func (UnimplementedAlertServer) Webhook(context.Context, *WebhookRequest) (*WebhookReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Webhook not implemented")
}
func (UnimplementedAlertServer) mustEmbedUnimplementedAlertServer() {}

// UnsafeAlertServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertServer will
// result in compilation errors.
type UnsafeAlertServer interface {
	mustEmbedUnimplementedAlertServer()
}

func RegisterAlertServer(s grpc.ServiceRegistrar, srv AlertServer) {
	s.RegisterService(&Alert_ServiceDesc, srv)
}

func _Alert_Webhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServer).Webhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.alert.v1.Alert/Webhook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServer).Webhook(ctx, req.(*WebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Alert_ServiceDesc is the grpc.ServiceDesc for Alert service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Alert_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.alert.v1.Alert",
	HandlerType: (*AlertServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Webhook",
			Handler:    _Alert_Webhook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alert/v1/alert.proto",
}
