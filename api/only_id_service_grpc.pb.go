// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: only_id_service.proto

package onlyIdSrv

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OnlyId_GetId_FullMethodName          = "/onlyIdSrv.OnlyId/GetId"
	OnlyId_GetSnowFlakeId_FullMethodName = "/onlyIdSrv.OnlyId/GetSnowFlakeId"
)

// OnlyIdClient is the client API for OnlyId service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OnlyIdClient interface {
	GetId(ctx context.Context, in *ReqId, opts ...grpc.CallOption) (*ResId, error)
	GetSnowFlakeId(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResId, error)
}

type onlyIdClient struct {
	cc grpc.ClientConnInterface
}

func NewOnlyIdClient(cc grpc.ClientConnInterface) OnlyIdClient {
	return &onlyIdClient{cc}
}

func (c *onlyIdClient) GetId(ctx context.Context, in *ReqId, opts ...grpc.CallOption) (*ResId, error) {
	out := new(ResId)
	err := c.cc.Invoke(ctx, OnlyId_GetId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlyIdClient) GetSnowFlakeId(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResId, error) {
	out := new(ResId)
	err := c.cc.Invoke(ctx, OnlyId_GetSnowFlakeId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnlyIdServer is the server API for OnlyId service.
// All implementations should embed UnimplementedOnlyIdServer
// for forward compatibility
type OnlyIdServer interface {
	GetId(context.Context, *ReqId) (*ResId, error)
	GetSnowFlakeId(context.Context, *emptypb.Empty) (*ResId, error)
}

// UnimplementedOnlyIdServer should be embedded to have forward compatible implementations.
type UnimplementedOnlyIdServer struct {
}

func (UnimplementedOnlyIdServer) GetId(context.Context, *ReqId) (*ResId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetId not implemented")
}
func (UnimplementedOnlyIdServer) GetSnowFlakeId(context.Context, *emptypb.Empty) (*ResId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnowFlakeId not implemented")
}

// UnsafeOnlyIdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OnlyIdServer will
// result in compilation errors.
type UnsafeOnlyIdServer interface {
	mustEmbedUnimplementedOnlyIdServer()
}

func RegisterOnlyIdServer(s grpc.ServiceRegistrar, srv OnlyIdServer) {
	s.RegisterService(&OnlyId_ServiceDesc, srv)
}

func _OnlyId_GetId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlyIdServer).GetId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnlyId_GetId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlyIdServer).GetId(ctx, req.(*ReqId))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnlyId_GetSnowFlakeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlyIdServer).GetSnowFlakeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnlyId_GetSnowFlakeId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlyIdServer).GetSnowFlakeId(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// OnlyId_ServiceDesc is the grpc.ServiceDesc for OnlyId service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OnlyId_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "onlyIdSrv.OnlyId",
	HandlerType: (*OnlyIdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetId",
			Handler:    _OnlyId_GetId_Handler,
		},
		{
			MethodName: "GetSnowFlakeId",
			Handler:    _OnlyId_GetSnowFlakeId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "only_id_service.proto",
}