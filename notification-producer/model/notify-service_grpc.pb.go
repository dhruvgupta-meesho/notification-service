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
	Notify_SendNotificationInfo_FullMethodName = "/Notify/SendNotificationInfo"
	Notify_AddBlacklisted_FullMethodName       = "/Notify/AddBlacklisted"
	Notify_RemoveBlacklisted_FullMethodName    = "/Notify/RemoveBlacklisted"
	Notify_GetRequestStatus_FullMethodName     = "/Notify/GetRequestStatus"
	Notify_GetLogs_FullMethodName              = "/Notify/GetLogs"
)

// NotifyClient is the client API for Notify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifyClient interface {
	SendNotificationInfo(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	AddBlacklisted(ctx context.Context, in *EmailList, opts ...grpc.CallOption) (*GenericResponse, error)
	RemoveBlacklisted(ctx context.Context, in *EmailList, opts ...grpc.CallOption) (*GenericResponse, error)
	GetRequestStatus(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*RequestStatusResponse, error)
	GetLogs(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogRequestResp, error)
}

type notifyClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifyClient(cc grpc.ClientConnInterface) NotifyClient {
	return &notifyClient{cc}
}

func (c *notifyClient) SendNotificationInfo(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, Notify_SendNotificationInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) AddBlacklisted(ctx context.Context, in *EmailList, opts ...grpc.CallOption) (*GenericResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, Notify_AddBlacklisted_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) RemoveBlacklisted(ctx context.Context, in *EmailList, opts ...grpc.CallOption) (*GenericResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, Notify_RemoveBlacklisted_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) GetRequestStatus(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*RequestStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RequestStatusResponse)
	err := c.cc.Invoke(ctx, Notify_GetRequestStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) GetLogs(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogRequestResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogRequestResp)
	err := c.cc.Invoke(ctx, Notify_GetLogs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifyServer is the server API for Notify service.
// All implementations must embed UnimplementedNotifyServer
// for forward compatibility.
type NotifyServer interface {
	SendNotificationInfo(context.Context, *EmailRequest) (*GenericResponse, error)
	AddBlacklisted(context.Context, *EmailList) (*GenericResponse, error)
	RemoveBlacklisted(context.Context, *EmailList) (*GenericResponse, error)
	GetRequestStatus(context.Context, *RequestID) (*RequestStatusResponse, error)
	GetLogs(context.Context, *LogRequest) (*LogRequestResp, error)
	mustEmbedUnimplementedNotifyServer()
}

// UnimplementedNotifyServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNotifyServer struct{}

func (UnimplementedNotifyServer) SendNotificationInfo(context.Context, *EmailRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotificationInfo not implemented")
}
func (UnimplementedNotifyServer) AddBlacklisted(context.Context, *EmailList) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlacklisted not implemented")
}
func (UnimplementedNotifyServer) RemoveBlacklisted(context.Context, *EmailList) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBlacklisted not implemented")
}
func (UnimplementedNotifyServer) GetRequestStatus(context.Context, *RequestID) (*RequestStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequestStatus not implemented")
}
func (UnimplementedNotifyServer) GetLogs(context.Context, *LogRequest) (*LogRequestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogs not implemented")
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

func _Notify_SendNotificationInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).SendNotificationInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notify_SendNotificationInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).SendNotificationInfo(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_AddBlacklisted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).AddBlacklisted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notify_AddBlacklisted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).AddBlacklisted(ctx, req.(*EmailList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_RemoveBlacklisted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).RemoveBlacklisted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notify_RemoveBlacklisted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).RemoveBlacklisted(ctx, req.(*EmailList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_GetRequestStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).GetRequestStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notify_GetRequestStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).GetRequestStatus(ctx, req.(*RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_GetLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).GetLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notify_GetLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).GetLogs(ctx, req.(*LogRequest))
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
			MethodName: "SendNotificationInfo",
			Handler:    _Notify_SendNotificationInfo_Handler,
		},
		{
			MethodName: "AddBlacklisted",
			Handler:    _Notify_AddBlacklisted_Handler,
		},
		{
			MethodName: "RemoveBlacklisted",
			Handler:    _Notify_RemoveBlacklisted_Handler,
		},
		{
			MethodName: "GetRequestStatus",
			Handler:    _Notify_GetRequestStatus_Handler,
		},
		{
			MethodName: "GetLogs",
			Handler:    _Notify_GetLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model/notify-service.proto",
}
