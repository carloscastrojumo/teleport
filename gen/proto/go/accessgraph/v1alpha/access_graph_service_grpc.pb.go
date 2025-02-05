//
// Teleport
// Copyright (C) 2023  Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: accessgraph/v1alpha/access_graph_service.proto

package accessgraphv1alpha

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
	AccessGraphService_Query_FullMethodName              = "/accessgraph.v1alpha.AccessGraphService/Query"
	AccessGraphService_GetFile_FullMethodName            = "/accessgraph.v1alpha.AccessGraphService/GetFile"
	AccessGraphService_EventsStream_FullMethodName       = "/accessgraph.v1alpha.AccessGraphService/EventsStream"
	AccessGraphService_EventsStreamV2_FullMethodName     = "/accessgraph.v1alpha.AccessGraphService/EventsStreamV2"
	AccessGraphService_Register_FullMethodName           = "/accessgraph.v1alpha.AccessGraphService/Register"
	AccessGraphService_ReplaceCAs_FullMethodName         = "/accessgraph.v1alpha.AccessGraphService/ReplaceCAs"
	AccessGraphService_AWSEventsStream_FullMethodName    = "/accessgraph.v1alpha.AccessGraphService/AWSEventsStream"
	AccessGraphService_GitlabEventsStream_FullMethodName = "/accessgraph.v1alpha.AccessGraphService/GitlabEventsStream"
	AccessGraphService_EntraEventsStream_FullMethodName  = "/accessgraph.v1alpha.AccessGraphService/EntraEventsStream"
	AccessGraphService_NetIQEventsStream_FullMethodName  = "/accessgraph.v1alpha.AccessGraphService/NetIQEventsStream"
	AccessGraphService_AzureEventsStream_FullMethodName  = "/accessgraph.v1alpha.AccessGraphService/AzureEventsStream"
)

// AccessGraphServiceClient is the client API for AccessGraphService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// AccessGraphService is a service for interacting the access graph service.
type AccessGraphServiceClient interface {
	// Query queries the access graph.
	// Currently only used by WebUI.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// GetFile gets a static UI file from the access graph container.
	GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error)
	// EventsStream is a stream of commands to the access graph service.
	// Teleport Auth server creates a stream to the access graph service
	// and pushes all resources and following events to it.
	// This stream is used to sync the access graph with the Teleport database state.
	// Once Teleport finishes syncing the current state, it sends a sync command
	// to the access graph service and resumes sending events.
	EventsStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[EventsStreamRequest, EventsStreamResponse], error)
	// EventsStreamV2 is a stream of commands to the access graph service.
	// This stream works the same way as EventsStream, but it returns a stream of events
	// instead of a single response.
	EventsStreamV2(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[EventsStreamV2Request, EventsStreamV2Response], error)
	// Register submits a new tenant representing this Teleport cluster to the TAG service,
	// identified by its HostCA certificate.
	// The method is idempotent: it succeeds if the tenant has already registered and has the specific CA associated.
	//
	// This method, unlike all others, expects the client to authenticate using a TLS certificate signed by the registration CA,
	// rather than the Teleport cluster's Host CA.
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// ReplaceCAs is a request to completely replace the set of Host CAs that authenticate this tenant with the given set.
	// This accommodates Teleport Host CA rotation. In a transition from certificate authority A to authority B,
	// the client is expected to call the RPC as follows:
	// 1. Authenticate via existing authority A and call ReplaceCAs([A, B]) -- introduce the incoming CA
	// 2.a. If rotation succeeds, authenticate via the new authority B and call ReplaceCAs([B]) -- delete the previous CA
	// 2.b. If rotation is rolled back, authenticate via the old authority A and call ReplaceCAs([A]) -- delete the candidate CA
	ReplaceCAs(ctx context.Context, in *ReplaceCAsRequest, opts ...grpc.CallOption) (*ReplaceCAsResponse, error)
	// AWSEventsStream is a stream of commands to the AWS importer.
	// Teleport Discovery Service creates a stream to the access graph service
	// and pushes all AWS resources and following events to it.
	// This stream is used to sync the access graph with the AWS database state.
	AWSEventsStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[AWSEventsStreamRequest, AWSEventsStreamResponse], error)
	// GitlabEventsStream is a stream of commands to the Gitlab importer.
	GitlabEventsStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[GitlabEventsStreamRequest, GitlabEventsStreamResponse], error)
	// EntraEventsStream is a stream of commands to the Entra ID SSO importer.
	EntraEventsStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[EntraEventsStreamRequest, EntraEventsStreamResponse], error)
	// NetIQEventsStream is a stream of commands to the NetIQ importer.
	NetIQEventsStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[NetIQEventsStreamRequest, NetIQEventsStreamResponse], error)
	// AzureEventsStream is a stream of commands to the Azure importer
	AzureEventsStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[AzureEventsStreamRequest, AzureEventsStreamResponse], error)
}

type accessGraphServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessGraphServiceClient(cc grpc.ClientConnInterface) AccessGraphServiceClient {
	return &accessGraphServiceClient{cc}
}

func (c *accessGraphServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, AccessGraphService_Query_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessGraphServiceClient) GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFileResponse)
	err := c.cc.Invoke(ctx, AccessGraphService_GetFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessGraphServiceClient) EventsStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[EventsStreamRequest, EventsStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AccessGraphService_ServiceDesc.Streams[0], AccessGraphService_EventsStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[EventsStreamRequest, EventsStreamResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessGraphService_EventsStreamClient = grpc.ClientStreamingClient[EventsStreamRequest, EventsStreamResponse]

func (c *accessGraphServiceClient) EventsStreamV2(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[EventsStreamV2Request, EventsStreamV2Response], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AccessGraphService_ServiceDesc.Streams[1], AccessGraphService_EventsStreamV2_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[EventsStreamV2Request, EventsStreamV2Response]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessGraphService_EventsStreamV2Client = grpc.BidiStreamingClient[EventsStreamV2Request, EventsStreamV2Response]

func (c *accessGraphServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, AccessGraphService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessGraphServiceClient) ReplaceCAs(ctx context.Context, in *ReplaceCAsRequest, opts ...grpc.CallOption) (*ReplaceCAsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplaceCAsResponse)
	err := c.cc.Invoke(ctx, AccessGraphService_ReplaceCAs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessGraphServiceClient) AWSEventsStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[AWSEventsStreamRequest, AWSEventsStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AccessGraphService_ServiceDesc.Streams[2], AccessGraphService_AWSEventsStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[AWSEventsStreamRequest, AWSEventsStreamResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessGraphService_AWSEventsStreamClient = grpc.ClientStreamingClient[AWSEventsStreamRequest, AWSEventsStreamResponse]

func (c *accessGraphServiceClient) GitlabEventsStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[GitlabEventsStreamRequest, GitlabEventsStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AccessGraphService_ServiceDesc.Streams[3], AccessGraphService_GitlabEventsStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GitlabEventsStreamRequest, GitlabEventsStreamResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessGraphService_GitlabEventsStreamClient = grpc.BidiStreamingClient[GitlabEventsStreamRequest, GitlabEventsStreamResponse]

func (c *accessGraphServiceClient) EntraEventsStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[EntraEventsStreamRequest, EntraEventsStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AccessGraphService_ServiceDesc.Streams[4], AccessGraphService_EntraEventsStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[EntraEventsStreamRequest, EntraEventsStreamResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessGraphService_EntraEventsStreamClient = grpc.BidiStreamingClient[EntraEventsStreamRequest, EntraEventsStreamResponse]

func (c *accessGraphServiceClient) NetIQEventsStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[NetIQEventsStreamRequest, NetIQEventsStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AccessGraphService_ServiceDesc.Streams[5], AccessGraphService_NetIQEventsStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[NetIQEventsStreamRequest, NetIQEventsStreamResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessGraphService_NetIQEventsStreamClient = grpc.BidiStreamingClient[NetIQEventsStreamRequest, NetIQEventsStreamResponse]

func (c *accessGraphServiceClient) AzureEventsStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[AzureEventsStreamRequest, AzureEventsStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AccessGraphService_ServiceDesc.Streams[6], AccessGraphService_AzureEventsStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[AzureEventsStreamRequest, AzureEventsStreamResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessGraphService_AzureEventsStreamClient = grpc.BidiStreamingClient[AzureEventsStreamRequest, AzureEventsStreamResponse]

// AccessGraphServiceServer is the server API for AccessGraphService service.
// All implementations must embed UnimplementedAccessGraphServiceServer
// for forward compatibility.
//
// AccessGraphService is a service for interacting the access graph service.
type AccessGraphServiceServer interface {
	// Query queries the access graph.
	// Currently only used by WebUI.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	// GetFile gets a static UI file from the access graph container.
	GetFile(context.Context, *GetFileRequest) (*GetFileResponse, error)
	// EventsStream is a stream of commands to the access graph service.
	// Teleport Auth server creates a stream to the access graph service
	// and pushes all resources and following events to it.
	// This stream is used to sync the access graph with the Teleport database state.
	// Once Teleport finishes syncing the current state, it sends a sync command
	// to the access graph service and resumes sending events.
	EventsStream(grpc.ClientStreamingServer[EventsStreamRequest, EventsStreamResponse]) error
	// EventsStreamV2 is a stream of commands to the access graph service.
	// This stream works the same way as EventsStream, but it returns a stream of events
	// instead of a single response.
	EventsStreamV2(grpc.BidiStreamingServer[EventsStreamV2Request, EventsStreamV2Response]) error
	// Register submits a new tenant representing this Teleport cluster to the TAG service,
	// identified by its HostCA certificate.
	// The method is idempotent: it succeeds if the tenant has already registered and has the specific CA associated.
	//
	// This method, unlike all others, expects the client to authenticate using a TLS certificate signed by the registration CA,
	// rather than the Teleport cluster's Host CA.
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// ReplaceCAs is a request to completely replace the set of Host CAs that authenticate this tenant with the given set.
	// This accommodates Teleport Host CA rotation. In a transition from certificate authority A to authority B,
	// the client is expected to call the RPC as follows:
	// 1. Authenticate via existing authority A and call ReplaceCAs([A, B]) -- introduce the incoming CA
	// 2.a. If rotation succeeds, authenticate via the new authority B and call ReplaceCAs([B]) -- delete the previous CA
	// 2.b. If rotation is rolled back, authenticate via the old authority A and call ReplaceCAs([A]) -- delete the candidate CA
	ReplaceCAs(context.Context, *ReplaceCAsRequest) (*ReplaceCAsResponse, error)
	// AWSEventsStream is a stream of commands to the AWS importer.
	// Teleport Discovery Service creates a stream to the access graph service
	// and pushes all AWS resources and following events to it.
	// This stream is used to sync the access graph with the AWS database state.
	AWSEventsStream(grpc.ClientStreamingServer[AWSEventsStreamRequest, AWSEventsStreamResponse]) error
	// GitlabEventsStream is a stream of commands to the Gitlab importer.
	GitlabEventsStream(grpc.BidiStreamingServer[GitlabEventsStreamRequest, GitlabEventsStreamResponse]) error
	// EntraEventsStream is a stream of commands to the Entra ID SSO importer.
	EntraEventsStream(grpc.BidiStreamingServer[EntraEventsStreamRequest, EntraEventsStreamResponse]) error
	// NetIQEventsStream is a stream of commands to the NetIQ importer.
	NetIQEventsStream(grpc.BidiStreamingServer[NetIQEventsStreamRequest, NetIQEventsStreamResponse]) error
	// AzureEventsStream is a stream of commands to the Azure importer
	AzureEventsStream(grpc.BidiStreamingServer[AzureEventsStreamRequest, AzureEventsStreamResponse]) error
	mustEmbedUnimplementedAccessGraphServiceServer()
}

// UnimplementedAccessGraphServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccessGraphServiceServer struct{}

func (UnimplementedAccessGraphServiceServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedAccessGraphServiceServer) GetFile(context.Context, *GetFileRequest) (*GetFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedAccessGraphServiceServer) EventsStream(grpc.ClientStreamingServer[EventsStreamRequest, EventsStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method EventsStream not implemented")
}
func (UnimplementedAccessGraphServiceServer) EventsStreamV2(grpc.BidiStreamingServer[EventsStreamV2Request, EventsStreamV2Response]) error {
	return status.Errorf(codes.Unimplemented, "method EventsStreamV2 not implemented")
}
func (UnimplementedAccessGraphServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAccessGraphServiceServer) ReplaceCAs(context.Context, *ReplaceCAsRequest) (*ReplaceCAsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceCAs not implemented")
}
func (UnimplementedAccessGraphServiceServer) AWSEventsStream(grpc.ClientStreamingServer[AWSEventsStreamRequest, AWSEventsStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method AWSEventsStream not implemented")
}
func (UnimplementedAccessGraphServiceServer) GitlabEventsStream(grpc.BidiStreamingServer[GitlabEventsStreamRequest, GitlabEventsStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GitlabEventsStream not implemented")
}
func (UnimplementedAccessGraphServiceServer) EntraEventsStream(grpc.BidiStreamingServer[EntraEventsStreamRequest, EntraEventsStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method EntraEventsStream not implemented")
}
func (UnimplementedAccessGraphServiceServer) NetIQEventsStream(grpc.BidiStreamingServer[NetIQEventsStreamRequest, NetIQEventsStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method NetIQEventsStream not implemented")
}
func (UnimplementedAccessGraphServiceServer) AzureEventsStream(grpc.BidiStreamingServer[AzureEventsStreamRequest, AzureEventsStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method AzureEventsStream not implemented")
}
func (UnimplementedAccessGraphServiceServer) mustEmbedUnimplementedAccessGraphServiceServer() {}
func (UnimplementedAccessGraphServiceServer) testEmbeddedByValue()                            {}

// UnsafeAccessGraphServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessGraphServiceServer will
// result in compilation errors.
type UnsafeAccessGraphServiceServer interface {
	mustEmbedUnimplementedAccessGraphServiceServer()
}

func RegisterAccessGraphServiceServer(s grpc.ServiceRegistrar, srv AccessGraphServiceServer) {
	// If the following call pancis, it indicates UnimplementedAccessGraphServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AccessGraphService_ServiceDesc, srv)
}

func _AccessGraphService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessGraphServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessGraphService_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessGraphServiceServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessGraphService_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessGraphServiceServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessGraphService_GetFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessGraphServiceServer).GetFile(ctx, req.(*GetFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessGraphService_EventsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccessGraphServiceServer).EventsStream(&grpc.GenericServerStream[EventsStreamRequest, EventsStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessGraphService_EventsStreamServer = grpc.ClientStreamingServer[EventsStreamRequest, EventsStreamResponse]

func _AccessGraphService_EventsStreamV2_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccessGraphServiceServer).EventsStreamV2(&grpc.GenericServerStream[EventsStreamV2Request, EventsStreamV2Response]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessGraphService_EventsStreamV2Server = grpc.BidiStreamingServer[EventsStreamV2Request, EventsStreamV2Response]

func _AccessGraphService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessGraphServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessGraphService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessGraphServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessGraphService_ReplaceCAs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceCAsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessGraphServiceServer).ReplaceCAs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessGraphService_ReplaceCAs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessGraphServiceServer).ReplaceCAs(ctx, req.(*ReplaceCAsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessGraphService_AWSEventsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccessGraphServiceServer).AWSEventsStream(&grpc.GenericServerStream[AWSEventsStreamRequest, AWSEventsStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessGraphService_AWSEventsStreamServer = grpc.ClientStreamingServer[AWSEventsStreamRequest, AWSEventsStreamResponse]

func _AccessGraphService_GitlabEventsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccessGraphServiceServer).GitlabEventsStream(&grpc.GenericServerStream[GitlabEventsStreamRequest, GitlabEventsStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessGraphService_GitlabEventsStreamServer = grpc.BidiStreamingServer[GitlabEventsStreamRequest, GitlabEventsStreamResponse]

func _AccessGraphService_EntraEventsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccessGraphServiceServer).EntraEventsStream(&grpc.GenericServerStream[EntraEventsStreamRequest, EntraEventsStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessGraphService_EntraEventsStreamServer = grpc.BidiStreamingServer[EntraEventsStreamRequest, EntraEventsStreamResponse]

func _AccessGraphService_NetIQEventsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccessGraphServiceServer).NetIQEventsStream(&grpc.GenericServerStream[NetIQEventsStreamRequest, NetIQEventsStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessGraphService_NetIQEventsStreamServer = grpc.BidiStreamingServer[NetIQEventsStreamRequest, NetIQEventsStreamResponse]

func _AccessGraphService_AzureEventsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccessGraphServiceServer).AzureEventsStream(&grpc.GenericServerStream[AzureEventsStreamRequest, AzureEventsStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessGraphService_AzureEventsStreamServer = grpc.BidiStreamingServer[AzureEventsStreamRequest, AzureEventsStreamResponse]

// AccessGraphService_ServiceDesc is the grpc.ServiceDesc for AccessGraphService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessGraphService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accessgraph.v1alpha.AccessGraphService",
	HandlerType: (*AccessGraphServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _AccessGraphService_Query_Handler,
		},
		{
			MethodName: "GetFile",
			Handler:    _AccessGraphService_GetFile_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AccessGraphService_Register_Handler,
		},
		{
			MethodName: "ReplaceCAs",
			Handler:    _AccessGraphService_ReplaceCAs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EventsStream",
			Handler:       _AccessGraphService_EventsStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "EventsStreamV2",
			Handler:       _AccessGraphService_EventsStreamV2_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "AWSEventsStream",
			Handler:       _AccessGraphService_AWSEventsStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GitlabEventsStream",
			Handler:       _AccessGraphService_GitlabEventsStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "EntraEventsStream",
			Handler:       _AccessGraphService_EntraEventsStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "NetIQEventsStream",
			Handler:       _AccessGraphService_NetIQEventsStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "AzureEventsStream",
			Handler:       _AccessGraphService_AzureEventsStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "accessgraph/v1alpha/access_graph_service.proto",
}
