// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: model/notify-service.proto

package model

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Notify_GetNotificationInfo_FullMethodName = "/Notify/GetNotificationInfo"
)

// NotifyClient is the client API for Notify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifyClient interface {
	GetNotificationInfo(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*GenericResponse, error)
}

type notifyClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifyClient(cc grpc.ClientConnInterface) NotifyClient {
	return &notifyClient{cc}
}

func (c *notifyClient) GetNotificationInfo(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, Notify_GetNotificationInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifyServer is the server API for Notify service.
// All implementations must embed UnimplementedNotifyServer
// for forward compatibility.
type NotifyServer interface {
	GetNotificationInfo(context.Context, *EmailRequest) (*GenericResponse, error)
	mustEmbedUnimplementedNotifyServer()
}

// UnimplementedNotifyServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNotifyServer struct{}

func (UnimplementedNotifyServer) GetNotificationInfo(context.Context, *EmailRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotificationInfo not implemented")
}
func (UnimplementedNotifyServer) mustEmbedUnimplementedNotifyServer() {}
func (UnimplementedNotifyServer) testEmbeddedByValue()                {}

// UnsafeNotifyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotifyServer will
// result in compilation errors.
type UnsafeNotifyServer interface {
	mustEmbedUnimplementedNotifyServer()
}

func RegisterNotifyServer(s grpc.ServiceRegistrar, srv NotifyServer) {
	// If the following call pancis, it indicates UnimplementedNotifyServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Notify_ServiceDesc, srv)
}

func _Notify_GetNotificationInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).GetNotificationInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notify_GetNotificationInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).GetNotificationInfo(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Notify_ServiceDesc is the grpc.ServiceDesc for Notify service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notify_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Notify",
	HandlerType: (*NotifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotificationInfo",
			Handler:    _Notify_GetNotificationInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model/notify-service.proto",
}
