// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: pb/datafusion.proto

package pb

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
	DataFusionExecutor_CreateSession_FullMethodName = "/df.DataFusionExecutor/CreateSession"
	DataFusionExecutor_ExecuteQuery_FullMethodName  = "/df.DataFusionExecutor/ExecuteQuery"
)

// DataFusionExecutorClient is the client API for DataFusionExecutor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataFusionExecutorClient interface {
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
	ExecuteQuery(ctx context.Context, in *ExecuteQueryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ExecuteQueryResponse], error)
}

type dataFusionExecutorClient struct {
	cc grpc.ClientConnInterface
}

func NewDataFusionExecutorClient(cc grpc.ClientConnInterface) DataFusionExecutorClient {
	return &dataFusionExecutorClient{cc}
}

func (c *dataFusionExecutorClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSessionResponse)
	err := c.cc.Invoke(ctx, DataFusionExecutor_CreateSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataFusionExecutorClient) ExecuteQuery(ctx context.Context, in *ExecuteQueryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ExecuteQueryResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DataFusionExecutor_ServiceDesc.Streams[0], DataFusionExecutor_ExecuteQuery_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ExecuteQueryRequest, ExecuteQueryResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DataFusionExecutor_ExecuteQueryClient = grpc.ServerStreamingClient[ExecuteQueryResponse]

// DataFusionExecutorServer is the server API for DataFusionExecutor service.
// All implementations must embed UnimplementedDataFusionExecutorServer
// for forward compatibility.
type DataFusionExecutorServer interface {
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	ExecuteQuery(*ExecuteQueryRequest, grpc.ServerStreamingServer[ExecuteQueryResponse]) error
	mustEmbedUnimplementedDataFusionExecutorServer()
}

// UnimplementedDataFusionExecutorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDataFusionExecutorServer struct{}

func (UnimplementedDataFusionExecutorServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedDataFusionExecutorServer) ExecuteQuery(*ExecuteQueryRequest, grpc.ServerStreamingServer[ExecuteQueryResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ExecuteQuery not implemented")
}
func (UnimplementedDataFusionExecutorServer) mustEmbedUnimplementedDataFusionExecutorServer() {}
func (UnimplementedDataFusionExecutorServer) testEmbeddedByValue()                            {}

// UnsafeDataFusionExecutorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataFusionExecutorServer will
// result in compilation errors.
type UnsafeDataFusionExecutorServer interface {
	mustEmbedUnimplementedDataFusionExecutorServer()
}

func RegisterDataFusionExecutorServer(s grpc.ServiceRegistrar, srv DataFusionExecutorServer) {
	// If the following call pancis, it indicates UnimplementedDataFusionExecutorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DataFusionExecutor_ServiceDesc, srv)
}

func _DataFusionExecutor_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataFusionExecutorServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataFusionExecutor_CreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataFusionExecutorServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataFusionExecutor_ExecuteQuery_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecuteQueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataFusionExecutorServer).ExecuteQuery(m, &grpc.GenericServerStream[ExecuteQueryRequest, ExecuteQueryResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DataFusionExecutor_ExecuteQueryServer = grpc.ServerStreamingServer[ExecuteQueryResponse]

// DataFusionExecutor_ServiceDesc is the grpc.ServiceDesc for DataFusionExecutor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataFusionExecutor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "df.DataFusionExecutor",
	HandlerType: (*DataFusionExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _DataFusionExecutor_CreateSession_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteQuery",
			Handler:       _DataFusionExecutor_ExecuteQuery_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/datafusion.proto",
}
