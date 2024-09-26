// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: microservio/organization/v1/organization.proto

package organization

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

// OrganizationServiceClient is the client API for OrganizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationServiceClient interface {
	ListOrganizations(ctx context.Context, in *ListOrganizationsRequest, opts ...grpc.CallOption) (OrganizationService_ListOrganizationsClient, error)
	GetOrganizationByID(ctx context.Context, in *GetOrganizationByIDRequest, opts ...grpc.CallOption) (*GetOrganizationByIDResponse, error)
	CreateNewOrganization(ctx context.Context, in *CreateNewOrganizationRequest, opts ...grpc.CallOption) (*CreateNewOrganizationResponse, error)
	UpdateOrganization(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*UpdateOrganizationResponse, error)
	AddMemberToOrganization(ctx context.Context, in *AddMemberToOrganizationRequest, opts ...grpc.CallOption) (*AddMemberToOrganizationResponse, error)
	RemoveMemberFromOrganization(ctx context.Context, in *RemoveMemberFromOrganizationRequest, opts ...grpc.CallOption) (*RemoveMemberFromOrganizationResponse, error)
}

type organizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationServiceClient(cc grpc.ClientConnInterface) OrganizationServiceClient {
	return &organizationServiceClient{cc}
}

func (c *organizationServiceClient) ListOrganizations(ctx context.Context, in *ListOrganizationsRequest, opts ...grpc.CallOption) (OrganizationService_ListOrganizationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrganizationService_ServiceDesc.Streams[0], "/microservio.organization.v1.OrganizationService/ListOrganizations", opts...)
	if err != nil {
		return nil, err
	}
	x := &organizationServiceListOrganizationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrganizationService_ListOrganizationsClient interface {
	Recv() (*ListOrganizationsResponse, error)
	grpc.ClientStream
}

type organizationServiceListOrganizationsClient struct {
	grpc.ClientStream
}

func (x *organizationServiceListOrganizationsClient) Recv() (*ListOrganizationsResponse, error) {
	m := new(ListOrganizationsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *organizationServiceClient) GetOrganizationByID(ctx context.Context, in *GetOrganizationByIDRequest, opts ...grpc.CallOption) (*GetOrganizationByIDResponse, error) {
	out := new(GetOrganizationByIDResponse)
	err := c.cc.Invoke(ctx, "/microservio.organization.v1.OrganizationService/GetOrganizationByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) CreateNewOrganization(ctx context.Context, in *CreateNewOrganizationRequest, opts ...grpc.CallOption) (*CreateNewOrganizationResponse, error) {
	out := new(CreateNewOrganizationResponse)
	err := c.cc.Invoke(ctx, "/microservio.organization.v1.OrganizationService/CreateNewOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) UpdateOrganization(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*UpdateOrganizationResponse, error) {
	out := new(UpdateOrganizationResponse)
	err := c.cc.Invoke(ctx, "/microservio.organization.v1.OrganizationService/UpdateOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) AddMemberToOrganization(ctx context.Context, in *AddMemberToOrganizationRequest, opts ...grpc.CallOption) (*AddMemberToOrganizationResponse, error) {
	out := new(AddMemberToOrganizationResponse)
	err := c.cc.Invoke(ctx, "/microservio.organization.v1.OrganizationService/AddMemberToOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) RemoveMemberFromOrganization(ctx context.Context, in *RemoveMemberFromOrganizationRequest, opts ...grpc.CallOption) (*RemoveMemberFromOrganizationResponse, error) {
	out := new(RemoveMemberFromOrganizationResponse)
	err := c.cc.Invoke(ctx, "/microservio.organization.v1.OrganizationService/RemoveMemberFromOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationServiceServer is the server API for OrganizationService service.
// All implementations must embed UnimplementedOrganizationServiceServer
// for forward compatibility
type OrganizationServiceServer interface {
	ListOrganizations(*ListOrganizationsRequest, OrganizationService_ListOrganizationsServer) error
	GetOrganizationByID(context.Context, *GetOrganizationByIDRequest) (*GetOrganizationByIDResponse, error)
	CreateNewOrganization(context.Context, *CreateNewOrganizationRequest) (*CreateNewOrganizationResponse, error)
	UpdateOrganization(context.Context, *UpdateOrganizationRequest) (*UpdateOrganizationResponse, error)
	AddMemberToOrganization(context.Context, *AddMemberToOrganizationRequest) (*AddMemberToOrganizationResponse, error)
	RemoveMemberFromOrganization(context.Context, *RemoveMemberFromOrganizationRequest) (*RemoveMemberFromOrganizationResponse, error)
	mustEmbedUnimplementedOrganizationServiceServer()
}

// UnimplementedOrganizationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrganizationServiceServer struct {
}

func (UnimplementedOrganizationServiceServer) ListOrganizations(*ListOrganizationsRequest, OrganizationService_ListOrganizationsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListOrganizations not implemented")
}
func (UnimplementedOrganizationServiceServer) GetOrganizationByID(context.Context, *GetOrganizationByIDRequest) (*GetOrganizationByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganizationByID not implemented")
}
func (UnimplementedOrganizationServiceServer) CreateNewOrganization(context.Context, *CreateNewOrganizationRequest) (*CreateNewOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewOrganization not implemented")
}
func (UnimplementedOrganizationServiceServer) UpdateOrganization(context.Context, *UpdateOrganizationRequest) (*UpdateOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrganization not implemented")
}
func (UnimplementedOrganizationServiceServer) AddMemberToOrganization(context.Context, *AddMemberToOrganizationRequest) (*AddMemberToOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMemberToOrganization not implemented")
}
func (UnimplementedOrganizationServiceServer) RemoveMemberFromOrganization(context.Context, *RemoveMemberFromOrganizationRequest) (*RemoveMemberFromOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMemberFromOrganization not implemented")
}
func (UnimplementedOrganizationServiceServer) mustEmbedUnimplementedOrganizationServiceServer() {}

// UnsafeOrganizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationServiceServer will
// result in compilation errors.
type UnsafeOrganizationServiceServer interface {
	mustEmbedUnimplementedOrganizationServiceServer()
}

func RegisterOrganizationServiceServer(s grpc.ServiceRegistrar, srv OrganizationServiceServer) {
	s.RegisterService(&OrganizationService_ServiceDesc, srv)
}

func _OrganizationService_ListOrganizations_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListOrganizationsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrganizationServiceServer).ListOrganizations(m, &organizationServiceListOrganizationsServer{stream})
}

type OrganizationService_ListOrganizationsServer interface {
	Send(*ListOrganizationsResponse) error
	grpc.ServerStream
}

type organizationServiceListOrganizationsServer struct {
	grpc.ServerStream
}

func (x *organizationServiceListOrganizationsServer) Send(m *ListOrganizationsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OrganizationService_GetOrganizationByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).GetOrganizationByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservio.organization.v1.OrganizationService/GetOrganizationByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).GetOrganizationByID(ctx, req.(*GetOrganizationByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_CreateNewOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).CreateNewOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservio.organization.v1.OrganizationService/CreateNewOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).CreateNewOrganization(ctx, req.(*CreateNewOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_UpdateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).UpdateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservio.organization.v1.OrganizationService/UpdateOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).UpdateOrganization(ctx, req.(*UpdateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_AddMemberToOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemberToOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).AddMemberToOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservio.organization.v1.OrganizationService/AddMemberToOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).AddMemberToOrganization(ctx, req.(*AddMemberToOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_RemoveMemberFromOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMemberFromOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).RemoveMemberFromOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservio.organization.v1.OrganizationService/RemoveMemberFromOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).RemoveMemberFromOrganization(ctx, req.(*RemoveMemberFromOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizationService_ServiceDesc is the grpc.ServiceDesc for OrganizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microservio.organization.v1.OrganizationService",
	HandlerType: (*OrganizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrganizationByID",
			Handler:    _OrganizationService_GetOrganizationByID_Handler,
		},
		{
			MethodName: "CreateNewOrganization",
			Handler:    _OrganizationService_CreateNewOrganization_Handler,
		},
		{
			MethodName: "UpdateOrganization",
			Handler:    _OrganizationService_UpdateOrganization_Handler,
		},
		{
			MethodName: "AddMemberToOrganization",
			Handler:    _OrganizationService_AddMemberToOrganization_Handler,
		},
		{
			MethodName: "RemoveMemberFromOrganization",
			Handler:    _OrganizationService_RemoveMemberFromOrganization_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListOrganizations",
			Handler:       _OrganizationService_ListOrganizations_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "microservio/organization/v1/organization.proto",
}
