// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: cerbos/svc/v1/svc.proto

package svcv1

import (
	context "context"
	v1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
	v11 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CerbosServiceClient is the client API for CerbosService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CerbosServiceClient interface {
	CheckResourceSet(ctx context.Context, in *v1.CheckResourceSetRequest, opts ...grpc.CallOption) (*v11.CheckResourceSetResponse, error)
	CheckResourceBatch(ctx context.Context, in *v1.CheckResourceBatchRequest, opts ...grpc.CallOption) (*v11.CheckResourceBatchResponse, error)
	ServerInfo(ctx context.Context, in *v1.ServerInfoRequest, opts ...grpc.CallOption) (*v11.ServerInfoResponse, error)
	ResourcesQueryPlan(ctx context.Context, in *v1.ResourcesQueryPlanRequest, opts ...grpc.CallOption) (*v11.ResourcesQueryPlanResponse, error)
}

type cerbosServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCerbosServiceClient(cc grpc.ClientConnInterface) CerbosServiceClient {
	return &cerbosServiceClient{cc}
}

func (c *cerbosServiceClient) CheckResourceSet(ctx context.Context, in *v1.CheckResourceSetRequest, opts ...grpc.CallOption) (*v11.CheckResourceSetResponse, error) {
	out := new(v11.CheckResourceSetResponse)
	err := c.cc.Invoke(ctx, "/cerbos.svc.v1.CerbosService/CheckResourceSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cerbosServiceClient) CheckResourceBatch(ctx context.Context, in *v1.CheckResourceBatchRequest, opts ...grpc.CallOption) (*v11.CheckResourceBatchResponse, error) {
	out := new(v11.CheckResourceBatchResponse)
	err := c.cc.Invoke(ctx, "/cerbos.svc.v1.CerbosService/CheckResourceBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cerbosServiceClient) ServerInfo(ctx context.Context, in *v1.ServerInfoRequest, opts ...grpc.CallOption) (*v11.ServerInfoResponse, error) {
	out := new(v11.ServerInfoResponse)
	err := c.cc.Invoke(ctx, "/cerbos.svc.v1.CerbosService/ServerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cerbosServiceClient) ResourcesQueryPlan(ctx context.Context, in *v1.ResourcesQueryPlanRequest, opts ...grpc.CallOption) (*v11.ResourcesQueryPlanResponse, error) {
	out := new(v11.ResourcesQueryPlanResponse)
	err := c.cc.Invoke(ctx, "/cerbos.svc.v1.CerbosService/ResourcesQueryPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CerbosServiceServer is the server API for CerbosService service.
// All implementations must embed UnimplementedCerbosServiceServer
// for forward compatibility
type CerbosServiceServer interface {
	CheckResourceSet(context.Context, *v1.CheckResourceSetRequest) (*v11.CheckResourceSetResponse, error)
	CheckResourceBatch(context.Context, *v1.CheckResourceBatchRequest) (*v11.CheckResourceBatchResponse, error)
	ServerInfo(context.Context, *v1.ServerInfoRequest) (*v11.ServerInfoResponse, error)
	ResourcesQueryPlan(context.Context, *v1.ResourcesQueryPlanRequest) (*v11.ResourcesQueryPlanResponse, error)
	mustEmbedUnimplementedCerbosServiceServer()
}

// UnimplementedCerbosServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCerbosServiceServer struct {
}

func (UnimplementedCerbosServiceServer) CheckResourceSet(context.Context, *v1.CheckResourceSetRequest) (*v11.CheckResourceSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckResourceSet not implemented")
}
func (UnimplementedCerbosServiceServer) CheckResourceBatch(context.Context, *v1.CheckResourceBatchRequest) (*v11.CheckResourceBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckResourceBatch not implemented")
}
func (UnimplementedCerbosServiceServer) ServerInfo(context.Context, *v1.ServerInfoRequest) (*v11.ServerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerInfo not implemented")
}
func (UnimplementedCerbosServiceServer) ResourcesQueryPlan(context.Context, *v1.ResourcesQueryPlanRequest) (*v11.ResourcesQueryPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResourcesQueryPlan not implemented")
}
func (UnimplementedCerbosServiceServer) mustEmbedUnimplementedCerbosServiceServer() {}

// UnsafeCerbosServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CerbosServiceServer will
// result in compilation errors.
type UnsafeCerbosServiceServer interface {
	mustEmbedUnimplementedCerbosServiceServer()
}

func RegisterCerbosServiceServer(s grpc.ServiceRegistrar, srv CerbosServiceServer) {
	s.RegisterService(&CerbosService_ServiceDesc, srv)
}

func _CerbosService_CheckResourceSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CheckResourceSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CerbosServiceServer).CheckResourceSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerbos.svc.v1.CerbosService/CheckResourceSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CerbosServiceServer).CheckResourceSet(ctx, req.(*v1.CheckResourceSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CerbosService_CheckResourceBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CheckResourceBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CerbosServiceServer).CheckResourceBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerbos.svc.v1.CerbosService/CheckResourceBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CerbosServiceServer).CheckResourceBatch(ctx, req.(*v1.CheckResourceBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CerbosService_ServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ServerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CerbosServiceServer).ServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerbos.svc.v1.CerbosService/ServerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CerbosServiceServer).ServerInfo(ctx, req.(*v1.ServerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CerbosService_ResourcesQueryPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ResourcesQueryPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CerbosServiceServer).ResourcesQueryPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerbos.svc.v1.CerbosService/ResourcesQueryPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CerbosServiceServer).ResourcesQueryPlan(ctx, req.(*v1.ResourcesQueryPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CerbosService_ServiceDesc is the grpc.ServiceDesc for CerbosService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CerbosService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerbos.svc.v1.CerbosService",
	HandlerType: (*CerbosServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckResourceSet",
			Handler:    _CerbosService_CheckResourceSet_Handler,
		},
		{
			MethodName: "CheckResourceBatch",
			Handler:    _CerbosService_CheckResourceBatch_Handler,
		},
		{
			MethodName: "ServerInfo",
			Handler:    _CerbosService_ServerInfo_Handler,
		},
		{
			MethodName: "ResourcesQueryPlan",
			Handler:    _CerbosService_ResourcesQueryPlan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cerbos/svc/v1/svc.proto",
}

// CerbosAdminServiceClient is the client API for CerbosAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CerbosAdminServiceClient interface {
	AddOrUpdatePolicy(ctx context.Context, in *v1.AddOrUpdatePolicyRequest, opts ...grpc.CallOption) (*v11.AddOrUpdatePolicyResponse, error)
	ListPolicies(ctx context.Context, in *v1.ListPoliciesRequest, opts ...grpc.CallOption) (*v11.ListPoliciesResponse, error)
	GetPolicy(ctx context.Context, in *v1.GetPolicyRequest, opts ...grpc.CallOption) (*v11.GetPolicyResponse, error)
	ListAuditLogEntries(ctx context.Context, in *v1.ListAuditLogEntriesRequest, opts ...grpc.CallOption) (CerbosAdminService_ListAuditLogEntriesClient, error)
	AddOrUpdateSchema(ctx context.Context, in *v1.AddOrUpdateSchemaRequest, opts ...grpc.CallOption) (*v11.AddOrUpdateSchemaResponse, error)
	ListSchemas(ctx context.Context, in *v1.ListSchemasRequest, opts ...grpc.CallOption) (*v11.ListSchemasResponse, error)
	GetSchema(ctx context.Context, in *v1.GetSchemaRequest, opts ...grpc.CallOption) (*v11.GetSchemaResponse, error)
	DeleteSchema(ctx context.Context, in *v1.DeleteSchemaRequest, opts ...grpc.CallOption) (*v11.DeleteSchemaResponse, error)
}

type cerbosAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCerbosAdminServiceClient(cc grpc.ClientConnInterface) CerbosAdminServiceClient {
	return &cerbosAdminServiceClient{cc}
}

func (c *cerbosAdminServiceClient) AddOrUpdatePolicy(ctx context.Context, in *v1.AddOrUpdatePolicyRequest, opts ...grpc.CallOption) (*v11.AddOrUpdatePolicyResponse, error) {
	out := new(v11.AddOrUpdatePolicyResponse)
	err := c.cc.Invoke(ctx, "/cerbos.svc.v1.CerbosAdminService/AddOrUpdatePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cerbosAdminServiceClient) ListPolicies(ctx context.Context, in *v1.ListPoliciesRequest, opts ...grpc.CallOption) (*v11.ListPoliciesResponse, error) {
	out := new(v11.ListPoliciesResponse)
	err := c.cc.Invoke(ctx, "/cerbos.svc.v1.CerbosAdminService/ListPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cerbosAdminServiceClient) GetPolicy(ctx context.Context, in *v1.GetPolicyRequest, opts ...grpc.CallOption) (*v11.GetPolicyResponse, error) {
	out := new(v11.GetPolicyResponse)
	err := c.cc.Invoke(ctx, "/cerbos.svc.v1.CerbosAdminService/GetPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cerbosAdminServiceClient) ListAuditLogEntries(ctx context.Context, in *v1.ListAuditLogEntriesRequest, opts ...grpc.CallOption) (CerbosAdminService_ListAuditLogEntriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CerbosAdminService_ServiceDesc.Streams[0], "/cerbos.svc.v1.CerbosAdminService/ListAuditLogEntries", opts...)
	if err != nil {
		return nil, err
	}
	x := &cerbosAdminServiceListAuditLogEntriesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CerbosAdminService_ListAuditLogEntriesClient interface {
	Recv() (*v11.ListAuditLogEntriesResponse, error)
	grpc.ClientStream
}

type cerbosAdminServiceListAuditLogEntriesClient struct {
	grpc.ClientStream
}

func (x *cerbosAdminServiceListAuditLogEntriesClient) Recv() (*v11.ListAuditLogEntriesResponse, error) {
	m := new(v11.ListAuditLogEntriesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cerbosAdminServiceClient) AddOrUpdateSchema(ctx context.Context, in *v1.AddOrUpdateSchemaRequest, opts ...grpc.CallOption) (*v11.AddOrUpdateSchemaResponse, error) {
	out := new(v11.AddOrUpdateSchemaResponse)
	err := c.cc.Invoke(ctx, "/cerbos.svc.v1.CerbosAdminService/AddOrUpdateSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cerbosAdminServiceClient) ListSchemas(ctx context.Context, in *v1.ListSchemasRequest, opts ...grpc.CallOption) (*v11.ListSchemasResponse, error) {
	out := new(v11.ListSchemasResponse)
	err := c.cc.Invoke(ctx, "/cerbos.svc.v1.CerbosAdminService/ListSchemas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cerbosAdminServiceClient) GetSchema(ctx context.Context, in *v1.GetSchemaRequest, opts ...grpc.CallOption) (*v11.GetSchemaResponse, error) {
	out := new(v11.GetSchemaResponse)
	err := c.cc.Invoke(ctx, "/cerbos.svc.v1.CerbosAdminService/GetSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cerbosAdminServiceClient) DeleteSchema(ctx context.Context, in *v1.DeleteSchemaRequest, opts ...grpc.CallOption) (*v11.DeleteSchemaResponse, error) {
	out := new(v11.DeleteSchemaResponse)
	err := c.cc.Invoke(ctx, "/cerbos.svc.v1.CerbosAdminService/DeleteSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CerbosAdminServiceServer is the server API for CerbosAdminService service.
// All implementations must embed UnimplementedCerbosAdminServiceServer
// for forward compatibility
type CerbosAdminServiceServer interface {
	AddOrUpdatePolicy(context.Context, *v1.AddOrUpdatePolicyRequest) (*v11.AddOrUpdatePolicyResponse, error)
	ListPolicies(context.Context, *v1.ListPoliciesRequest) (*v11.ListPoliciesResponse, error)
	GetPolicy(context.Context, *v1.GetPolicyRequest) (*v11.GetPolicyResponse, error)
	ListAuditLogEntries(*v1.ListAuditLogEntriesRequest, CerbosAdminService_ListAuditLogEntriesServer) error
	AddOrUpdateSchema(context.Context, *v1.AddOrUpdateSchemaRequest) (*v11.AddOrUpdateSchemaResponse, error)
	ListSchemas(context.Context, *v1.ListSchemasRequest) (*v11.ListSchemasResponse, error)
	GetSchema(context.Context, *v1.GetSchemaRequest) (*v11.GetSchemaResponse, error)
	DeleteSchema(context.Context, *v1.DeleteSchemaRequest) (*v11.DeleteSchemaResponse, error)
	mustEmbedUnimplementedCerbosAdminServiceServer()
}

// UnimplementedCerbosAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCerbosAdminServiceServer struct {
}

func (UnimplementedCerbosAdminServiceServer) AddOrUpdatePolicy(context.Context, *v1.AddOrUpdatePolicyRequest) (*v11.AddOrUpdatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrUpdatePolicy not implemented")
}
func (UnimplementedCerbosAdminServiceServer) ListPolicies(context.Context, *v1.ListPoliciesRequest) (*v11.ListPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicies not implemented")
}
func (UnimplementedCerbosAdminServiceServer) GetPolicy(context.Context, *v1.GetPolicyRequest) (*v11.GetPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicy not implemented")
}
func (UnimplementedCerbosAdminServiceServer) ListAuditLogEntries(*v1.ListAuditLogEntriesRequest, CerbosAdminService_ListAuditLogEntriesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAuditLogEntries not implemented")
}
func (UnimplementedCerbosAdminServiceServer) AddOrUpdateSchema(context.Context, *v1.AddOrUpdateSchemaRequest) (*v11.AddOrUpdateSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrUpdateSchema not implemented")
}
func (UnimplementedCerbosAdminServiceServer) ListSchemas(context.Context, *v1.ListSchemasRequest) (*v11.ListSchemasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchemas not implemented")
}
func (UnimplementedCerbosAdminServiceServer) GetSchema(context.Context, *v1.GetSchemaRequest) (*v11.GetSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchema not implemented")
}
func (UnimplementedCerbosAdminServiceServer) DeleteSchema(context.Context, *v1.DeleteSchemaRequest) (*v11.DeleteSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSchema not implemented")
}
func (UnimplementedCerbosAdminServiceServer) mustEmbedUnimplementedCerbosAdminServiceServer() {}

// UnsafeCerbosAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CerbosAdminServiceServer will
// result in compilation errors.
type UnsafeCerbosAdminServiceServer interface {
	mustEmbedUnimplementedCerbosAdminServiceServer()
}

func RegisterCerbosAdminServiceServer(s grpc.ServiceRegistrar, srv CerbosAdminServiceServer) {
	s.RegisterService(&CerbosAdminService_ServiceDesc, srv)
}

func _CerbosAdminService_AddOrUpdatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.AddOrUpdatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CerbosAdminServiceServer).AddOrUpdatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerbos.svc.v1.CerbosAdminService/AddOrUpdatePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CerbosAdminServiceServer).AddOrUpdatePolicy(ctx, req.(*v1.AddOrUpdatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CerbosAdminService_ListPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ListPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CerbosAdminServiceServer).ListPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerbos.svc.v1.CerbosAdminService/ListPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CerbosAdminServiceServer).ListPolicies(ctx, req.(*v1.ListPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CerbosAdminService_GetPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CerbosAdminServiceServer).GetPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerbos.svc.v1.CerbosAdminService/GetPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CerbosAdminServiceServer).GetPolicy(ctx, req.(*v1.GetPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CerbosAdminService_ListAuditLogEntries_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(v1.ListAuditLogEntriesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CerbosAdminServiceServer).ListAuditLogEntries(m, &cerbosAdminServiceListAuditLogEntriesServer{stream})
}

type CerbosAdminService_ListAuditLogEntriesServer interface {
	Send(*v11.ListAuditLogEntriesResponse) error
	grpc.ServerStream
}

type cerbosAdminServiceListAuditLogEntriesServer struct {
	grpc.ServerStream
}

func (x *cerbosAdminServiceListAuditLogEntriesServer) Send(m *v11.ListAuditLogEntriesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CerbosAdminService_AddOrUpdateSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.AddOrUpdateSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CerbosAdminServiceServer).AddOrUpdateSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerbos.svc.v1.CerbosAdminService/AddOrUpdateSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CerbosAdminServiceServer).AddOrUpdateSchema(ctx, req.(*v1.AddOrUpdateSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CerbosAdminService_ListSchemas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ListSchemasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CerbosAdminServiceServer).ListSchemas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerbos.svc.v1.CerbosAdminService/ListSchemas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CerbosAdminServiceServer).ListSchemas(ctx, req.(*v1.ListSchemasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CerbosAdminService_GetSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CerbosAdminServiceServer).GetSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerbos.svc.v1.CerbosAdminService/GetSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CerbosAdminServiceServer).GetSchema(ctx, req.(*v1.GetSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CerbosAdminService_DeleteSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.DeleteSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CerbosAdminServiceServer).DeleteSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerbos.svc.v1.CerbosAdminService/DeleteSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CerbosAdminServiceServer).DeleteSchema(ctx, req.(*v1.DeleteSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CerbosAdminService_ServiceDesc is the grpc.ServiceDesc for CerbosAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CerbosAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerbos.svc.v1.CerbosAdminService",
	HandlerType: (*CerbosAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrUpdatePolicy",
			Handler:    _CerbosAdminService_AddOrUpdatePolicy_Handler,
		},
		{
			MethodName: "ListPolicies",
			Handler:    _CerbosAdminService_ListPolicies_Handler,
		},
		{
			MethodName: "GetPolicy",
			Handler:    _CerbosAdminService_GetPolicy_Handler,
		},
		{
			MethodName: "AddOrUpdateSchema",
			Handler:    _CerbosAdminService_AddOrUpdateSchema_Handler,
		},
		{
			MethodName: "ListSchemas",
			Handler:    _CerbosAdminService_ListSchemas_Handler,
		},
		{
			MethodName: "GetSchema",
			Handler:    _CerbosAdminService_GetSchema_Handler,
		},
		{
			MethodName: "DeleteSchema",
			Handler:    _CerbosAdminService_DeleteSchema_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAuditLogEntries",
			Handler:       _CerbosAdminService_ListAuditLogEntries_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cerbos/svc/v1/svc.proto",
}

// CerbosPlaygroundServiceClient is the client API for CerbosPlaygroundService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CerbosPlaygroundServiceClient interface {
	PlaygroundValidate(ctx context.Context, in *v1.PlaygroundValidateRequest, opts ...grpc.CallOption) (*v11.PlaygroundValidateResponse, error)
	PlaygroundEvaluate(ctx context.Context, in *v1.PlaygroundEvaluateRequest, opts ...grpc.CallOption) (*v11.PlaygroundEvaluateResponse, error)
	PlaygroundProxy(ctx context.Context, in *v1.PlaygroundProxyRequest, opts ...grpc.CallOption) (*v11.PlaygroundProxyResponse, error)
}

type cerbosPlaygroundServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCerbosPlaygroundServiceClient(cc grpc.ClientConnInterface) CerbosPlaygroundServiceClient {
	return &cerbosPlaygroundServiceClient{cc}
}

func (c *cerbosPlaygroundServiceClient) PlaygroundValidate(ctx context.Context, in *v1.PlaygroundValidateRequest, opts ...grpc.CallOption) (*v11.PlaygroundValidateResponse, error) {
	out := new(v11.PlaygroundValidateResponse)
	err := c.cc.Invoke(ctx, "/cerbos.svc.v1.CerbosPlaygroundService/PlaygroundValidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cerbosPlaygroundServiceClient) PlaygroundEvaluate(ctx context.Context, in *v1.PlaygroundEvaluateRequest, opts ...grpc.CallOption) (*v11.PlaygroundEvaluateResponse, error) {
	out := new(v11.PlaygroundEvaluateResponse)
	err := c.cc.Invoke(ctx, "/cerbos.svc.v1.CerbosPlaygroundService/PlaygroundEvaluate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cerbosPlaygroundServiceClient) PlaygroundProxy(ctx context.Context, in *v1.PlaygroundProxyRequest, opts ...grpc.CallOption) (*v11.PlaygroundProxyResponse, error) {
	out := new(v11.PlaygroundProxyResponse)
	err := c.cc.Invoke(ctx, "/cerbos.svc.v1.CerbosPlaygroundService/PlaygroundProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CerbosPlaygroundServiceServer is the server API for CerbosPlaygroundService service.
// All implementations must embed UnimplementedCerbosPlaygroundServiceServer
// for forward compatibility
type CerbosPlaygroundServiceServer interface {
	PlaygroundValidate(context.Context, *v1.PlaygroundValidateRequest) (*v11.PlaygroundValidateResponse, error)
	PlaygroundEvaluate(context.Context, *v1.PlaygroundEvaluateRequest) (*v11.PlaygroundEvaluateResponse, error)
	PlaygroundProxy(context.Context, *v1.PlaygroundProxyRequest) (*v11.PlaygroundProxyResponse, error)
	mustEmbedUnimplementedCerbosPlaygroundServiceServer()
}

// UnimplementedCerbosPlaygroundServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCerbosPlaygroundServiceServer struct {
}

func (UnimplementedCerbosPlaygroundServiceServer) PlaygroundValidate(context.Context, *v1.PlaygroundValidateRequest) (*v11.PlaygroundValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaygroundValidate not implemented")
}
func (UnimplementedCerbosPlaygroundServiceServer) PlaygroundEvaluate(context.Context, *v1.PlaygroundEvaluateRequest) (*v11.PlaygroundEvaluateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaygroundEvaluate not implemented")
}
func (UnimplementedCerbosPlaygroundServiceServer) PlaygroundProxy(context.Context, *v1.PlaygroundProxyRequest) (*v11.PlaygroundProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaygroundProxy not implemented")
}
func (UnimplementedCerbosPlaygroundServiceServer) mustEmbedUnimplementedCerbosPlaygroundServiceServer() {
}

// UnsafeCerbosPlaygroundServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CerbosPlaygroundServiceServer will
// result in compilation errors.
type UnsafeCerbosPlaygroundServiceServer interface {
	mustEmbedUnimplementedCerbosPlaygroundServiceServer()
}

func RegisterCerbosPlaygroundServiceServer(s grpc.ServiceRegistrar, srv CerbosPlaygroundServiceServer) {
	s.RegisterService(&CerbosPlaygroundService_ServiceDesc, srv)
}

func _CerbosPlaygroundService_PlaygroundValidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PlaygroundValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CerbosPlaygroundServiceServer).PlaygroundValidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerbos.svc.v1.CerbosPlaygroundService/PlaygroundValidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CerbosPlaygroundServiceServer).PlaygroundValidate(ctx, req.(*v1.PlaygroundValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CerbosPlaygroundService_PlaygroundEvaluate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PlaygroundEvaluateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CerbosPlaygroundServiceServer).PlaygroundEvaluate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerbos.svc.v1.CerbosPlaygroundService/PlaygroundEvaluate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CerbosPlaygroundServiceServer).PlaygroundEvaluate(ctx, req.(*v1.PlaygroundEvaluateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CerbosPlaygroundService_PlaygroundProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PlaygroundProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CerbosPlaygroundServiceServer).PlaygroundProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerbos.svc.v1.CerbosPlaygroundService/PlaygroundProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CerbosPlaygroundServiceServer).PlaygroundProxy(ctx, req.(*v1.PlaygroundProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CerbosPlaygroundService_ServiceDesc is the grpc.ServiceDesc for CerbosPlaygroundService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CerbosPlaygroundService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerbos.svc.v1.CerbosPlaygroundService",
	HandlerType: (*CerbosPlaygroundServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaygroundValidate",
			Handler:    _CerbosPlaygroundService_PlaygroundValidate_Handler,
		},
		{
			MethodName: "PlaygroundEvaluate",
			Handler:    _CerbosPlaygroundService_PlaygroundEvaluate_Handler,
		},
		{
			MethodName: "PlaygroundProxy",
			Handler:    _CerbosPlaygroundService_PlaygroundProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cerbos/svc/v1/svc.proto",
}
