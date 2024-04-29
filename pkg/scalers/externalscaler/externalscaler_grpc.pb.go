// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: externalscaler.proto

package externalscaler

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
	ExternalScaler_IsActive_FullMethodName       = "/externalscaler.ExternalScaler/IsActive"
	ExternalScaler_StreamIsActive_FullMethodName = "/externalscaler.ExternalScaler/StreamIsActive"
	ExternalScaler_GetMetricSpec_FullMethodName  = "/externalscaler.ExternalScaler/GetMetricSpec"
	ExternalScaler_GetMetrics_FullMethodName     = "/externalscaler.ExternalScaler/GetMetrics"
)

// ExternalScalerClient is the client API for ExternalScaler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExternalScalerClient interface {
	IsActive(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (*IsActiveResponse, error)
	StreamIsActive(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (ExternalScaler_StreamIsActiveClient, error)
	GetMetricSpec(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (*GetMetricSpecResponse, error)
	GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
}

type externalScalerClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalScalerClient(cc grpc.ClientConnInterface) ExternalScalerClient {
	return &externalScalerClient{cc}
}

func (c *externalScalerClient) IsActive(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (*IsActiveResponse, error) {
	out := new(IsActiveResponse)
	err := c.cc.Invoke(ctx, ExternalScaler_IsActive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalScalerClient) StreamIsActive(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (ExternalScaler_StreamIsActiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExternalScaler_ServiceDesc.Streams[0], ExternalScaler_StreamIsActive_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &externalScalerStreamIsActiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExternalScaler_StreamIsActiveClient interface {
	Recv() (*IsActiveResponse, error)
	grpc.ClientStream
}

type externalScalerStreamIsActiveClient struct {
	grpc.ClientStream
}

func (x *externalScalerStreamIsActiveClient) Recv() (*IsActiveResponse, error) {
	m := new(IsActiveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *externalScalerClient) GetMetricSpec(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (*GetMetricSpecResponse, error) {
	out := new(GetMetricSpecResponse)
	err := c.cc.Invoke(ctx, ExternalScaler_GetMetricSpec_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalScalerClient) GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, ExternalScaler_GetMetrics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalScalerServer is the server API for ExternalScaler service.
// All implementations must embed UnimplementedExternalScalerServer
// for forward compatibility
type ExternalScalerServer interface {
	IsActive(context.Context, *ScaledObjectRef) (*IsActiveResponse, error)
	StreamIsActive(*ScaledObjectRef, ExternalScaler_StreamIsActiveServer) error
	GetMetricSpec(context.Context, *ScaledObjectRef) (*GetMetricSpecResponse, error)
	GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
	mustEmbedUnimplementedExternalScalerServer()
}

// UnimplementedExternalScalerServer must be embedded to have forward compatible implementations.
type UnimplementedExternalScalerServer struct {
}

func (UnimplementedExternalScalerServer) IsActive(context.Context, *ScaledObjectRef) (*IsActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsActive not implemented")
}
func (UnimplementedExternalScalerServer) StreamIsActive(*ScaledObjectRef, ExternalScaler_StreamIsActiveServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamIsActive not implemented")
}
func (UnimplementedExternalScalerServer) GetMetricSpec(context.Context, *ScaledObjectRef) (*GetMetricSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricSpec not implemented")
}
func (UnimplementedExternalScalerServer) GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}
func (UnimplementedExternalScalerServer) mustEmbedUnimplementedExternalScalerServer() {}

// UnsafeExternalScalerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalScalerServer will
// result in compilation errors.
type UnsafeExternalScalerServer interface {
	mustEmbedUnimplementedExternalScalerServer()
}

func RegisterExternalScalerServer(s grpc.ServiceRegistrar, srv ExternalScalerServer) {
	s.RegisterService(&ExternalScaler_ServiceDesc, srv)
}

func _ExternalScaler_IsActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaledObjectRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalScalerServer).IsActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalScaler_IsActive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalScalerServer).IsActive(ctx, req.(*ScaledObjectRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalScaler_StreamIsActive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ScaledObjectRef)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExternalScalerServer).StreamIsActive(m, &externalScalerStreamIsActiveServer{stream})
}

type ExternalScaler_StreamIsActiveServer interface {
	Send(*IsActiveResponse) error
	grpc.ServerStream
}

type externalScalerStreamIsActiveServer struct {
	grpc.ServerStream
}

func (x *externalScalerStreamIsActiveServer) Send(m *IsActiveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ExternalScaler_GetMetricSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaledObjectRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalScalerServer).GetMetricSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalScaler_GetMetricSpec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalScalerServer).GetMetricSpec(ctx, req.(*ScaledObjectRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalScaler_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalScalerServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalScaler_GetMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalScalerServer).GetMetrics(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExternalScaler_ServiceDesc is the grpc.ServiceDesc for ExternalScaler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExternalScaler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "externalscaler.ExternalScaler",
	HandlerType: (*ExternalScalerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsActive",
			Handler:    _ExternalScaler_IsActive_Handler,
		},
		{
			MethodName: "GetMetricSpec",
			Handler:    _ExternalScaler_GetMetricSpec_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _ExternalScaler_GetMetrics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamIsActive",
			Handler:       _ExternalScaler_StreamIsActive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "externalscaler.proto",
}
