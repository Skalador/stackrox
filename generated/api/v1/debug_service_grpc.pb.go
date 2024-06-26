// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v1/debug_service.proto

package v1

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
	DebugService_GetLogLevel_FullMethodName       = "/v1.DebugService/GetLogLevel"
	DebugService_SetLogLevel_FullMethodName       = "/v1.DebugService/SetLogLevel"
	DebugService_StreamAuthzTraces_FullMethodName = "/v1.DebugService/StreamAuthzTraces"
	DebugService_ResetDBStats_FullMethodName      = "/v1.DebugService/ResetDBStats"
)

// DebugServiceClient is the client API for DebugService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DebugServiceClient interface {
	// Get the current logging level for StackRox services.
	GetLogLevel(ctx context.Context, in *GetLogLevelRequest, opts ...grpc.CallOption) (*LogLevelResponse, error)
	// Set logging level for StackRox services.
	SetLogLevel(ctx context.Context, in *LogLevelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Stream authorization traces for all incoming requests.
	StreamAuthzTraces(ctx context.Context, in *Empty, opts ...grpc.CallOption) (DebugService_StreamAuthzTracesClient, error)
	// Reset database debugging statistics.
	ResetDBStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type debugServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDebugServiceClient(cc grpc.ClientConnInterface) DebugServiceClient {
	return &debugServiceClient{cc}
}

func (c *debugServiceClient) GetLogLevel(ctx context.Context, in *GetLogLevelRequest, opts ...grpc.CallOption) (*LogLevelResponse, error) {
	out := new(LogLevelResponse)
	err := c.cc.Invoke(ctx, DebugService_GetLogLevel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugServiceClient) SetLogLevel(ctx context.Context, in *LogLevelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DebugService_SetLogLevel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugServiceClient) StreamAuthzTraces(ctx context.Context, in *Empty, opts ...grpc.CallOption) (DebugService_StreamAuthzTracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &DebugService_ServiceDesc.Streams[0], DebugService_StreamAuthzTraces_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &debugServiceStreamAuthzTracesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DebugService_StreamAuthzTracesClient interface {
	Recv() (*AuthorizationTraceResponse, error)
	grpc.ClientStream
}

type debugServiceStreamAuthzTracesClient struct {
	grpc.ClientStream
}

func (x *debugServiceStreamAuthzTracesClient) Recv() (*AuthorizationTraceResponse, error) {
	m := new(AuthorizationTraceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugServiceClient) ResetDBStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, DebugService_ResetDBStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DebugServiceServer is the server API for DebugService service.
// All implementations should embed UnimplementedDebugServiceServer
// for forward compatibility
type DebugServiceServer interface {
	// Get the current logging level for StackRox services.
	GetLogLevel(context.Context, *GetLogLevelRequest) (*LogLevelResponse, error)
	// Set logging level for StackRox services.
	SetLogLevel(context.Context, *LogLevelRequest) (*emptypb.Empty, error)
	// Stream authorization traces for all incoming requests.
	StreamAuthzTraces(*Empty, DebugService_StreamAuthzTracesServer) error
	// Reset database debugging statistics.
	ResetDBStats(context.Context, *Empty) (*Empty, error)
}

// UnimplementedDebugServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDebugServiceServer struct {
}

func (UnimplementedDebugServiceServer) GetLogLevel(context.Context, *GetLogLevelRequest) (*LogLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogLevel not implemented")
}
func (UnimplementedDebugServiceServer) SetLogLevel(context.Context, *LogLevelRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLogLevel not implemented")
}
func (UnimplementedDebugServiceServer) StreamAuthzTraces(*Empty, DebugService_StreamAuthzTracesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAuthzTraces not implemented")
}
func (UnimplementedDebugServiceServer) ResetDBStats(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetDBStats not implemented")
}

// UnsafeDebugServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DebugServiceServer will
// result in compilation errors.
type UnsafeDebugServiceServer interface {
	mustEmbedUnimplementedDebugServiceServer()
}

func RegisterDebugServiceServer(s grpc.ServiceRegistrar, srv DebugServiceServer) {
	s.RegisterService(&DebugService_ServiceDesc, srv)
}

func _DebugService_GetLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServiceServer).GetLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebugService_GetLogLevel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServiceServer).GetLogLevel(ctx, req.(*GetLogLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebugService_SetLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServiceServer).SetLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebugService_SetLogLevel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServiceServer).SetLogLevel(ctx, req.(*LogLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebugService_StreamAuthzTraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServiceServer).StreamAuthzTraces(m, &debugServiceStreamAuthzTracesServer{stream})
}

type DebugService_StreamAuthzTracesServer interface {
	Send(*AuthorizationTraceResponse) error
	grpc.ServerStream
}

type debugServiceStreamAuthzTracesServer struct {
	grpc.ServerStream
}

func (x *debugServiceStreamAuthzTracesServer) Send(m *AuthorizationTraceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DebugService_ResetDBStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServiceServer).ResetDBStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebugService_ResetDBStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServiceServer).ResetDBStats(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DebugService_ServiceDesc is the grpc.ServiceDesc for DebugService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DebugService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DebugService",
	HandlerType: (*DebugServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLogLevel",
			Handler:    _DebugService_GetLogLevel_Handler,
		},
		{
			MethodName: "SetLogLevel",
			Handler:    _DebugService_SetLogLevel_Handler,
		},
		{
			MethodName: "ResetDBStats",
			Handler:    _DebugService_ResetDBStats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAuthzTraces",
			Handler:       _DebugService_StreamAuthzTraces_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v1/debug_service.proto",
}
