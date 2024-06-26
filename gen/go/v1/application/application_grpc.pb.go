// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: v1/application/application.proto

package application

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
	ApplicationService_Create_FullMethodName          = "/v1.application.ApplicationService/Create"
	ApplicationService_GetApplications_FullMethodName = "/v1.application.ApplicationService/GetApplications"
	ApplicationService_GetApplication_FullMethodName  = "/v1.application.ApplicationService/GetApplication"
)

// ApplicationServiceClient is the client API for ApplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApplicationServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	GetApplications(ctx context.Context, in *GetApplicationsRequest, opts ...grpc.CallOption) (*GetApplicationsResponse, error)
	GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error)
}

type applicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationServiceClient(cc grpc.ClientConnInterface) ApplicationServiceClient {
	return &applicationServiceClient{cc}
}

func (c *applicationServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, ApplicationService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) GetApplications(ctx context.Context, in *GetApplicationsRequest, opts ...grpc.CallOption) (*GetApplicationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetApplicationsResponse)
	err := c.cc.Invoke(ctx, ApplicationService_GetApplications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetApplicationResponse)
	err := c.cc.Invoke(ctx, ApplicationService_GetApplication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationServiceServer is the server API for ApplicationService service.
// All implementations must embed UnimplementedApplicationServiceServer
// for forward compatibility
type ApplicationServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	GetApplications(context.Context, *GetApplicationsRequest) (*GetApplicationsResponse, error)
	GetApplication(context.Context, *GetApplicationRequest) (*GetApplicationResponse, error)
	mustEmbedUnimplementedApplicationServiceServer()
}

// UnimplementedApplicationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApplicationServiceServer struct {
}

func (UnimplementedApplicationServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedApplicationServiceServer) GetApplications(context.Context, *GetApplicationsRequest) (*GetApplicationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplications not implemented")
}
func (UnimplementedApplicationServiceServer) GetApplication(context.Context, *GetApplicationRequest) (*GetApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplication not implemented")
}
func (UnimplementedApplicationServiceServer) mustEmbedUnimplementedApplicationServiceServer() {}

// UnsafeApplicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplicationServiceServer will
// result in compilation errors.
type UnsafeApplicationServiceServer interface {
	mustEmbedUnimplementedApplicationServiceServer()
}

func RegisterApplicationServiceServer(s grpc.ServiceRegistrar, srv ApplicationServiceServer) {
	s.RegisterService(&ApplicationService_ServiceDesc, srv)
}

func _ApplicationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_GetApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).GetApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationService_GetApplications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).GetApplications(ctx, req.(*GetApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_GetApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).GetApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationService_GetApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).GetApplication(ctx, req.(*GetApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApplicationService_ServiceDesc is the grpc.ServiceDesc for ApplicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApplicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.application.ApplicationService",
	HandlerType: (*ApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ApplicationService_Create_Handler,
		},
		{
			MethodName: "GetApplications",
			Handler:    _ApplicationService_GetApplications_Handler,
		},
		{
			MethodName: "GetApplication",
			Handler:    _ApplicationService_GetApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/application/application.proto",
}