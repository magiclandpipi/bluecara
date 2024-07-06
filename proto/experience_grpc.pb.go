// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: experience.proto

package experiencepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ExperienceService_GetHello_FullMethodName          = "/experience.ExperienceService/GetHello"
	ExperienceService_GetExperienceByID_FullMethodName = "/experience.ExperienceService/GetExperienceByID"
	ExperienceService_ListExperiences_FullMethodName   = "/experience.ExperienceService/ListExperiences"
	ExperienceService_UpsertExperience_FullMethodName  = "/experience.ExperienceService/UpsertExperience"
)

// ExperienceServiceClient is the client API for ExperienceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The experience service definition.
type ExperienceServiceClient interface {
	GetHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// Returns the experience detail based on ID.
	GetExperienceByID(ctx context.Context, in *ExperienceIDRequest, opts ...grpc.CallOption) (*ExperienceDetailResponse, error)
	// Returns a list of experiences.
	ListExperiences(ctx context.Context, in *ListExperiencesRequest, opts ...grpc.CallOption) (*ListExperiencesResponse, error)
	// Upserts an experience.
	UpsertExperience(ctx context.Context, in *UpsertExperienceRequest, opts ...grpc.CallOption) (*UpsertExperienceResponse, error)
}

type experienceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExperienceServiceClient(cc grpc.ClientConnInterface) ExperienceServiceClient {
	return &experienceServiceClient{cc}
}

func (c *experienceServiceClient) GetHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, ExperienceService_GetHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experienceServiceClient) GetExperienceByID(ctx context.Context, in *ExperienceIDRequest, opts ...grpc.CallOption) (*ExperienceDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExperienceDetailResponse)
	err := c.cc.Invoke(ctx, ExperienceService_GetExperienceByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experienceServiceClient) ListExperiences(ctx context.Context, in *ListExperiencesRequest, opts ...grpc.CallOption) (*ListExperiencesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListExperiencesResponse)
	err := c.cc.Invoke(ctx, ExperienceService_ListExperiences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experienceServiceClient) UpsertExperience(ctx context.Context, in *UpsertExperienceRequest, opts ...grpc.CallOption) (*UpsertExperienceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpsertExperienceResponse)
	err := c.cc.Invoke(ctx, ExperienceService_UpsertExperience_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExperienceServiceServer is the server API for ExperienceService service.
// All implementations must embed UnimplementedExperienceServiceServer
// for forward compatibility
//
// The experience service definition.
type ExperienceServiceServer interface {
	GetHello(context.Context, *HelloRequest) (*HelloResponse, error)
	// Returns the experience detail based on ID.
	GetExperienceByID(context.Context, *ExperienceIDRequest) (*ExperienceDetailResponse, error)
	// Returns a list of experiences.
	ListExperiences(context.Context, *ListExperiencesRequest) (*ListExperiencesResponse, error)
	// Upserts an experience.
	UpsertExperience(context.Context, *UpsertExperienceRequest) (*UpsertExperienceResponse, error)
	mustEmbedUnimplementedExperienceServiceServer()
}

// UnimplementedExperienceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExperienceServiceServer struct {
}

func (UnimplementedExperienceServiceServer) GetHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHello not implemented")
}
func (UnimplementedExperienceServiceServer) GetExperienceByID(context.Context, *ExperienceIDRequest) (*ExperienceDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExperienceByID not implemented")
}
func (UnimplementedExperienceServiceServer) ListExperiences(context.Context, *ListExperiencesRequest) (*ListExperiencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExperiences not implemented")
}
func (UnimplementedExperienceServiceServer) UpsertExperience(context.Context, *UpsertExperienceRequest) (*UpsertExperienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertExperience not implemented")
}
func (UnimplementedExperienceServiceServer) mustEmbedUnimplementedExperienceServiceServer() {}

// UnsafeExperienceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExperienceServiceServer will
// result in compilation errors.
type UnsafeExperienceServiceServer interface {
	mustEmbedUnimplementedExperienceServiceServer()
}

func RegisterExperienceServiceServer(s grpc.ServiceRegistrar, srv ExperienceServiceServer) {
	s.RegisterService(&ExperienceService_ServiceDesc, srv)
}

func _ExperienceService_GetHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperienceServiceServer).GetHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperienceService_GetHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperienceServiceServer).GetHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperienceService_GetExperienceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExperienceIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperienceServiceServer).GetExperienceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperienceService_GetExperienceByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperienceServiceServer).GetExperienceByID(ctx, req.(*ExperienceIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperienceService_ListExperiences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExperiencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperienceServiceServer).ListExperiences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperienceService_ListExperiences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperienceServiceServer).ListExperiences(ctx, req.(*ListExperiencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperienceService_UpsertExperience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertExperienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperienceServiceServer).UpsertExperience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperienceService_UpsertExperience_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperienceServiceServer).UpsertExperience(ctx, req.(*UpsertExperienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExperienceService_ServiceDesc is the grpc.ServiceDesc for ExperienceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExperienceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "experience.ExperienceService",
	HandlerType: (*ExperienceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHello",
			Handler:    _ExperienceService_GetHello_Handler,
		},
		{
			MethodName: "GetExperienceByID",
			Handler:    _ExperienceService_GetExperienceByID_Handler,
		},
		{
			MethodName: "ListExperiences",
			Handler:    _ExperienceService_ListExperiences_Handler,
		},
		{
			MethodName: "UpsertExperience",
			Handler:    _ExperienceService_UpsertExperience_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "experience.proto",
}
