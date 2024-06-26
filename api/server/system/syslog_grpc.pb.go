// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: server/system/syslog.proto

package system

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

// SyslogClient is the client API for Syslog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyslogClient interface {
	// 日志列表接口
	ListSyslog(ctx context.Context, in *ListSyslogRequest, opts ...grpc.CallOption) (*ListSyslogReply, error)
}

type syslogClient struct {
	cc grpc.ClientConnInterface
}

func NewSyslogClient(cc grpc.ClientConnInterface) SyslogClient {
	return &syslogClient{cc}
}

func (c *syslogClient) ListSyslog(ctx context.Context, in *ListSyslogRequest, opts ...grpc.CallOption) (*ListSyslogReply, error) {
	out := new(ListSyslogReply)
	err := c.cc.Invoke(ctx, "/api.system.Syslog/ListSyslog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyslogServer is the server API for Syslog service.
// All implementations must embed UnimplementedSyslogServer
// for forward compatibility
type SyslogServer interface {
	// 日志列表接口
	ListSyslog(context.Context, *ListSyslogRequest) (*ListSyslogReply, error)
	mustEmbedUnimplementedSyslogServer()
}

// UnimplementedSyslogServer must be embedded to have forward compatible implementations.
type UnimplementedSyslogServer struct {
}

func (UnimplementedSyslogServer) ListSyslog(context.Context, *ListSyslogRequest) (*ListSyslogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSyslog not implemented")
}
func (UnimplementedSyslogServer) mustEmbedUnimplementedSyslogServer() {}

// UnsafeSyslogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyslogServer will
// result in compilation errors.
type UnsafeSyslogServer interface {
	mustEmbedUnimplementedSyslogServer()
}

func RegisterSyslogServer(s grpc.ServiceRegistrar, srv SyslogServer) {
	s.RegisterService(&Syslog_ServiceDesc, srv)
}

func _Syslog_ListSyslog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSyslogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogServer).ListSyslog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.system.Syslog/ListSyslog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogServer).ListSyslog(ctx, req.(*ListSyslogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Syslog_ServiceDesc is the grpc.ServiceDesc for Syslog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Syslog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.system.Syslog",
	HandlerType: (*SyslogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSyslog",
			Handler:    _Syslog_ListSyslog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/system/syslog.proto",
}
