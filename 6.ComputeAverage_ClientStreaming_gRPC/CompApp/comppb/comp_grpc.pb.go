// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: comppb/comp.proto

package comppb

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

// CompServiceClient is the client API for CompService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompServiceClient interface {
	// Client Service
	CompAvg(ctx context.Context, opts ...grpc.CallOption) (CompService_CompAvgClient, error)
}

type compServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompServiceClient(cc grpc.ClientConnInterface) CompServiceClient {
	return &compServiceClient{cc}
}

func (c *compServiceClient) CompAvg(ctx context.Context, opts ...grpc.CallOption) (CompService_CompAvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &CompService_ServiceDesc.Streams[0], "/comp.CompService/CompAvg", opts...)
	if err != nil {
		return nil, err
	}
	x := &compServiceCompAvgClient{stream}
	return x, nil
}

type CompService_CompAvgClient interface {
	Send(*CompRequest) error
	CloseAndRecv() (*CompResponse, error)
	grpc.ClientStream
}

type compServiceCompAvgClient struct {
	grpc.ClientStream
}

func (x *compServiceCompAvgClient) Send(m *CompRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *compServiceCompAvgClient) CloseAndRecv() (*CompResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CompResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CompServiceServer is the server API for CompService service.
// All implementations must embed UnimplementedCompServiceServer
// for forward compatibility
type CompServiceServer interface {
	// Client Service
	CompAvg(CompService_CompAvgServer) error
	mustEmbedUnimplementedCompServiceServer()
}

// UnimplementedCompServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompServiceServer struct {
}

func (UnimplementedCompServiceServer) CompAvg(CompService_CompAvgServer) error {
	return status.Errorf(codes.Unimplemented, "method CompAvg not implemented")
}
func (UnimplementedCompServiceServer) mustEmbedUnimplementedCompServiceServer() {}

// UnsafeCompServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompServiceServer will
// result in compilation errors.
type UnsafeCompServiceServer interface {
	mustEmbedUnimplementedCompServiceServer()
}

func RegisterCompServiceServer(s grpc.ServiceRegistrar, srv CompServiceServer) {
	s.RegisterService(&CompService_ServiceDesc, srv)
}

func _CompService_CompAvg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CompServiceServer).CompAvg(&compServiceCompAvgServer{stream})
}

type CompService_CompAvgServer interface {
	SendAndClose(*CompResponse) error
	Recv() (*CompRequest, error)
	grpc.ServerStream
}

type compServiceCompAvgServer struct {
	grpc.ServerStream
}

func (x *compServiceCompAvgServer) SendAndClose(m *CompResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *compServiceCompAvgServer) Recv() (*CompRequest, error) {
	m := new(CompRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CompService_ServiceDesc is the grpc.ServiceDesc for CompService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comp.CompService",
	HandlerType: (*CompServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CompAvg",
			Handler:       _CompService_CompAvg_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "comppb/comp.proto",
}