// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// LanternClient is the client API for Lantern service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LanternClient interface {
	Illuminate(ctx context.Context, in *IlluminateRequest, opts ...grpc.CallOption) (*IlluminateResponse, error)
	GetVertex(ctx context.Context, in *Keys, opts ...grpc.CallOption) (*Graph, error)
	Put(ctx context.Context, in *Graph, opts ...grpc.CallOption) (*PutResponse, error)
	PutVertex(ctx context.Context, in *Graph, opts ...grpc.CallOption) (*PutResponse, error)
	PutEdge(ctx context.Context, in *Graph, opts ...grpc.CallOption) (*PutResponse, error)
}

type lanternClient struct {
	cc grpc.ClientConnInterface
}

func NewLanternClient(cc grpc.ClientConnInterface) LanternClient {
	return &lanternClient{cc}
}

func (c *lanternClient) Illuminate(ctx context.Context, in *IlluminateRequest, opts ...grpc.CallOption) (*IlluminateResponse, error) {
	out := new(IlluminateResponse)
	err := c.cc.Invoke(ctx, "/Lantern/Illuminate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lanternClient) GetVertex(ctx context.Context, in *Keys, opts ...grpc.CallOption) (*Graph, error) {
	out := new(Graph)
	err := c.cc.Invoke(ctx, "/Lantern/GetVertex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lanternClient) Put(ctx context.Context, in *Graph, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/Lantern/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lanternClient) PutVertex(ctx context.Context, in *Graph, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/Lantern/PutVertex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lanternClient) PutEdge(ctx context.Context, in *Graph, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/Lantern/PutEdge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LanternServer is the server API for Lantern service.
// All implementations must embed UnimplementedLanternServer
// for forward compatibility
type LanternServer interface {
	Illuminate(context.Context, *IlluminateRequest) (*IlluminateResponse, error)
	GetVertex(context.Context, *Keys) (*Graph, error)
	Put(context.Context, *Graph) (*PutResponse, error)
	PutVertex(context.Context, *Graph) (*PutResponse, error)
	PutEdge(context.Context, *Graph) (*PutResponse, error)
	mustEmbedUnimplementedLanternServer()
}

// UnimplementedLanternServer must be embedded to have forward compatible implementations.
type UnimplementedLanternServer struct {
}

func (UnimplementedLanternServer) Illuminate(context.Context, *IlluminateRequest) (*IlluminateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Illuminate not implemented")
}
func (UnimplementedLanternServer) GetVertex(context.Context, *Keys) (*Graph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVertex not implemented")
}
func (UnimplementedLanternServer) Put(context.Context, *Graph) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedLanternServer) PutVertex(context.Context, *Graph) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutVertex not implemented")
}
func (UnimplementedLanternServer) PutEdge(context.Context, *Graph) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutEdge not implemented")
}
func (UnimplementedLanternServer) mustEmbedUnimplementedLanternServer() {}

// UnsafeLanternServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LanternServer will
// result in compilation errors.
type UnsafeLanternServer interface {
	mustEmbedUnimplementedLanternServer()
}

func RegisterLanternServer(s grpc.ServiceRegistrar, srv LanternServer) {
	s.RegisterService(&Lantern_ServiceDesc, srv)
}

func _Lantern_Illuminate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IlluminateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanternServer).Illuminate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Lantern/Illuminate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanternServer).Illuminate(ctx, req.(*IlluminateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lantern_GetVertex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Keys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanternServer).GetVertex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Lantern/GetVertex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanternServer).GetVertex(ctx, req.(*Keys))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lantern_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Graph)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanternServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Lantern/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanternServer).Put(ctx, req.(*Graph))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lantern_PutVertex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Graph)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanternServer).PutVertex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Lantern/PutVertex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanternServer).PutVertex(ctx, req.(*Graph))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lantern_PutEdge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Graph)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanternServer).PutEdge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Lantern/PutEdge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanternServer).PutEdge(ctx, req.(*Graph))
	}
	return interceptor(ctx, in, info, handler)
}

// Lantern_ServiceDesc is the grpc.ServiceDesc for Lantern service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lantern_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Lantern",
	HandlerType: (*LanternServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Illuminate",
			Handler:    _Lantern_Illuminate_Handler,
		},
		{
			MethodName: "GetVertex",
			Handler:    _Lantern_GetVertex_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Lantern_Put_Handler,
		},
		{
			MethodName: "PutVertex",
			Handler:    _Lantern_PutVertex_Handler,
		},
		{
			MethodName: "PutEdge",
			Handler:    _Lantern_PutEdge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}
