// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: focus.proto

package proto

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

// FocusServiceClient is the client API for FocusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FocusServiceClient interface {
	GetFocus(ctx context.Context, in *GetFocusRequest, opts ...grpc.CallOption) (*GetFocusResponse, error)
}

type focusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFocusServiceClient(cc grpc.ClientConnInterface) FocusServiceClient {
	return &focusServiceClient{cc}
}

func (c *focusServiceClient) GetFocus(ctx context.Context, in *GetFocusRequest, opts ...grpc.CallOption) (*GetFocusResponse, error) {
	out := new(GetFocusResponse)
	err := c.cc.Invoke(ctx, "/focus.FocusService/GetFocus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FocusServiceServer is the server API for FocusService service.
// All implementations should embed UnimplementedFocusServiceServer
// for forward compatibility
type FocusServiceServer interface {
	GetFocus(context.Context, *GetFocusRequest) (*GetFocusResponse, error)
}

// UnimplementedFocusServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFocusServiceServer struct {
}

func (UnimplementedFocusServiceServer) GetFocus(context.Context, *GetFocusRequest) (*GetFocusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFocus not implemented")
}

// UnsafeFocusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FocusServiceServer will
// result in compilation errors.
type UnsafeFocusServiceServer interface {
	mustEmbedUnimplementedFocusServiceServer()
}

func RegisterFocusServiceServer(s grpc.ServiceRegistrar, srv FocusServiceServer) {
	s.RegisterService(&FocusService_ServiceDesc, srv)
}

func _FocusService_GetFocus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFocusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FocusServiceServer).GetFocus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/focus.FocusService/GetFocus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FocusServiceServer).GetFocus(ctx, req.(*GetFocusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FocusService_ServiceDesc is the grpc.ServiceDesc for FocusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FocusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "focus.FocusService",
	HandlerType: (*FocusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFocus",
			Handler:    _FocusService_GetFocus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "focus.proto",
}
