// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: pkg/logchain/logchain.proto

package logchain

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

// LogChainClient is the client API for LogChain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogChainClient interface {
	// Sends a greeting
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
}

type logChainClient struct {
	cc grpc.ClientConnInterface
}

func NewLogChainClient(cc grpc.ClientConnInterface) LogChainClient {
	return &logChainClient{cc}
}

func (c *logChainClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/LogChain/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogChainServer is the server API for LogChain service.
// All implementations must embed UnimplementedLogChainServer
// for forward compatibility
type LogChainServer interface {
	// Sends a greeting
	Log(context.Context, *LogRequest) (*LogResponse, error)
	mustEmbedUnimplementedLogChainServer()
}

// UnimplementedLogChainServer must be embedded to have forward compatible implementations.
type UnimplementedLogChainServer struct {
}

func (UnimplementedLogChainServer) Log(context.Context, *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedLogChainServer) mustEmbedUnimplementedLogChainServer() {}

// UnsafeLogChainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogChainServer will
// result in compilation errors.
type UnsafeLogChainServer interface {
	mustEmbedUnimplementedLogChainServer()
}

func RegisterLogChainServer(s grpc.ServiceRegistrar, srv LogChainServer) {
	s.RegisterService(&LogChain_ServiceDesc, srv)
}

func _LogChain_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogChainServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LogChain/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogChainServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogChain_ServiceDesc is the grpc.ServiceDesc for LogChain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogChain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LogChain",
	HandlerType: (*LogChainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _LogChain_Log_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/logchain/logchain.proto",
}