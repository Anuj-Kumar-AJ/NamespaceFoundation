// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: internal/testdata/service/proto/multidb.proto

package proto

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

// MultiDbListServiceClient is the client API for MultiDbListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MultiDbListServiceClient interface {
	AddRds(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddPostgres(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddMaria(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Merges from all dbs.
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListResponse, error)
}

type multiDbListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMultiDbListServiceClient(cc grpc.ClientConnInterface) MultiDbListServiceClient {
	return &multiDbListServiceClient{cc}
}

func (c *multiDbListServiceClient) AddRds(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/internal.testdata.service.proto.MultiDbListService/AddRds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multiDbListServiceClient) AddPostgres(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/internal.testdata.service.proto.MultiDbListService/AddPostgres", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multiDbListServiceClient) AddMaria(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/internal.testdata.service.proto.MultiDbListService/AddMaria", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multiDbListServiceClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/internal.testdata.service.proto.MultiDbListService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MultiDbListServiceServer is the server API for MultiDbListService service.
// All implementations should embed UnimplementedMultiDbListServiceServer
// for forward compatibility
type MultiDbListServiceServer interface {
	AddRds(context.Context, *AddRequest) (*emptypb.Empty, error)
	AddPostgres(context.Context, *AddRequest) (*emptypb.Empty, error)
	AddMaria(context.Context, *AddRequest) (*emptypb.Empty, error)
	// Merges from all dbs.
	List(context.Context, *emptypb.Empty) (*ListResponse, error)
}

// UnimplementedMultiDbListServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMultiDbListServiceServer struct {
}

func (UnimplementedMultiDbListServiceServer) AddRds(context.Context, *AddRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRds not implemented")
}
func (UnimplementedMultiDbListServiceServer) AddPostgres(context.Context, *AddRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPostgres not implemented")
}
func (UnimplementedMultiDbListServiceServer) AddMaria(context.Context, *AddRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMaria not implemented")
}
func (UnimplementedMultiDbListServiceServer) List(context.Context, *emptypb.Empty) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeMultiDbListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MultiDbListServiceServer will
// result in compilation errors.
type UnsafeMultiDbListServiceServer interface {
	mustEmbedUnimplementedMultiDbListServiceServer()
}

func RegisterMultiDbListServiceServer(s grpc.ServiceRegistrar, srv MultiDbListServiceServer) {
	s.RegisterService(&MultiDbListService_ServiceDesc, srv)
}

func _MultiDbListService_AddRds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiDbListServiceServer).AddRds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.testdata.service.proto.MultiDbListService/AddRds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiDbListServiceServer).AddRds(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultiDbListService_AddPostgres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiDbListServiceServer).AddPostgres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.testdata.service.proto.MultiDbListService/AddPostgres",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiDbListServiceServer).AddPostgres(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultiDbListService_AddMaria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiDbListServiceServer).AddMaria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.testdata.service.proto.MultiDbListService/AddMaria",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiDbListServiceServer).AddMaria(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultiDbListService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiDbListServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.testdata.service.proto.MultiDbListService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiDbListServiceServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MultiDbListService_ServiceDesc is the grpc.ServiceDesc for MultiDbListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MultiDbListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "internal.testdata.service.proto.MultiDbListService",
	HandlerType: (*MultiDbListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRds",
			Handler:    _MultiDbListService_AddRds_Handler,
		},
		{
			MethodName: "AddPostgres",
			Handler:    _MultiDbListService_AddPostgres_Handler,
		},
		{
			MethodName: "AddMaria",
			Handler:    _MultiDbListService_AddMaria_Handler,
		},
		{
			MethodName: "List",
			Handler:    _MultiDbListService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/testdata/service/proto/multidb.proto",
}
