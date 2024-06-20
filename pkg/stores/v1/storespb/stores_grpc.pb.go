// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: bobadojo/stores/v1/stores.proto

package storespb

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
	Stores_ListStores_FullMethodName = "/bobadojo.stores.v1.Stores/ListStores"
	Stores_FindStores_FullMethodName = "/bobadojo.stores.v1.Stores/FindStores"
	Stores_GetStore_FullMethodName   = "/bobadojo.stores.v1.Stores/GetStore"
)

// StoresClient is the client API for Stores service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Get stores and related information.
type StoresClient interface {
	// List all stores.
	ListStores(ctx context.Context, in *ListStoresRequest, opts ...grpc.CallOption) (*ListStoresResponse, error)
	// Returns a list of all stores in a specified region.
	FindStores(ctx context.Context, in *FindStoresRequest, opts ...grpc.CallOption) (*FindStoresResponse, error)
	// Returns a specific store.
	GetStore(ctx context.Context, in *GetStoreRequest, opts ...grpc.CallOption) (*Store, error)
}

type storesClient struct {
	cc grpc.ClientConnInterface
}

func NewStoresClient(cc grpc.ClientConnInterface) StoresClient {
	return &storesClient{cc}
}

func (c *storesClient) ListStores(ctx context.Context, in *ListStoresRequest, opts ...grpc.CallOption) (*ListStoresResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListStoresResponse)
	err := c.cc.Invoke(ctx, Stores_ListStores_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesClient) FindStores(ctx context.Context, in *FindStoresRequest, opts ...grpc.CallOption) (*FindStoresResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindStoresResponse)
	err := c.cc.Invoke(ctx, Stores_FindStores_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesClient) GetStore(ctx context.Context, in *GetStoreRequest, opts ...grpc.CallOption) (*Store, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Store)
	err := c.cc.Invoke(ctx, Stores_GetStore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoresServer is the server API for Stores service.
// All implementations must embed UnimplementedStoresServer
// for forward compatibility
//
// Get stores and related information.
type StoresServer interface {
	// List all stores.
	ListStores(context.Context, *ListStoresRequest) (*ListStoresResponse, error)
	// Returns a list of all stores in a specified region.
	FindStores(context.Context, *FindStoresRequest) (*FindStoresResponse, error)
	// Returns a specific store.
	GetStore(context.Context, *GetStoreRequest) (*Store, error)
	mustEmbedUnimplementedStoresServer()
}

// UnimplementedStoresServer must be embedded to have forward compatible implementations.
type UnimplementedStoresServer struct {
}

func (UnimplementedStoresServer) ListStores(context.Context, *ListStoresRequest) (*ListStoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStores not implemented")
}
func (UnimplementedStoresServer) FindStores(context.Context, *FindStoresRequest) (*FindStoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindStores not implemented")
}
func (UnimplementedStoresServer) GetStore(context.Context, *GetStoreRequest) (*Store, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStore not implemented")
}
func (UnimplementedStoresServer) mustEmbedUnimplementedStoresServer() {}

// UnsafeStoresServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoresServer will
// result in compilation errors.
type UnsafeStoresServer interface {
	mustEmbedUnimplementedStoresServer()
}

func RegisterStoresServer(s grpc.ServiceRegistrar, srv StoresServer) {
	s.RegisterService(&Stores_ServiceDesc, srv)
}

func _Stores_ListStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServer).ListStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stores_ListStores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServer).ListStores(ctx, req.(*ListStoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stores_FindStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindStoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServer).FindStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stores_FindStores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServer).FindStores(ctx, req.(*FindStoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stores_GetStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServer).GetStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stores_GetStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServer).GetStore(ctx, req.(*GetStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Stores_ServiceDesc is the grpc.ServiceDesc for Stores service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stores_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bobadojo.stores.v1.Stores",
	HandlerType: (*StoresServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListStores",
			Handler:    _Stores_ListStores_Handler,
		},
		{
			MethodName: "FindStores",
			Handler:    _Stores_FindStores_Handler,
		},
		{
			MethodName: "GetStore",
			Handler:    _Stores_GetStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bobadojo/stores/v1/stores.proto",
}
