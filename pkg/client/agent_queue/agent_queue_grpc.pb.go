// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: pkg/api/proto/agent_queue.proto

package agent_queue

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

// AgentQueueClient is the client API for AgentQueue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentQueueClient interface {
	DeseminateConfig(ctx context.Context, in *DeseminateConfigRequest, opts ...grpc.CallOption) (*DeseminateConfigResponse, error)
}

type agentQueueClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentQueueClient(cc grpc.ClientConnInterface) AgentQueueClient {
	return &agentQueueClient{cc}
}

func (c *agentQueueClient) DeseminateConfig(ctx context.Context, in *DeseminateConfigRequest, opts ...grpc.CallOption) (*DeseminateConfigResponse, error) {
	out := new(DeseminateConfigResponse)
	err := c.cc.Invoke(ctx, "/C12S.AgentQueue.AgentQueue/DeseminateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentQueueServer is the server API for AgentQueue service.
// All implementations must embed UnimplementedAgentQueueServer
// for forward compatibility
type AgentQueueServer interface {
	DeseminateConfig(context.Context, *DeseminateConfigRequest) (*DeseminateConfigResponse, error)
	mustEmbedUnimplementedAgentQueueServer()
}

// UnimplementedAgentQueueServer must be embedded to have forward compatible implementations.
type UnimplementedAgentQueueServer struct {
}

func (UnimplementedAgentQueueServer) DeseminateConfig(context.Context, *DeseminateConfigRequest) (*DeseminateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeseminateConfig not implemented")
}
func (UnimplementedAgentQueueServer) mustEmbedUnimplementedAgentQueueServer() {}

// UnsafeAgentQueueServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentQueueServer will
// result in compilation errors.
type UnsafeAgentQueueServer interface {
	mustEmbedUnimplementedAgentQueueServer()
}

func RegisterAgentQueueServer(s grpc.ServiceRegistrar, srv AgentQueueServer) {
	s.RegisterService(&AgentQueue_ServiceDesc, srv)
}

func _AgentQueue_DeseminateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeseminateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentQueueServer).DeseminateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/C12S.AgentQueue.AgentQueue/DeseminateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentQueueServer).DeseminateConfig(ctx, req.(*DeseminateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentQueue_ServiceDesc is the grpc.ServiceDesc for AgentQueue service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentQueue_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "C12S.AgentQueue.AgentQueue",
	HandlerType: (*AgentQueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeseminateConfig",
			Handler:    _AgentQueue_DeseminateConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/proto/agent_queue.proto",
}
