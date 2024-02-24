// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: entry_point.proto

package entry_point

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EntryPointClient is the client API for EntryPoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntryPointClient interface {
}

type entryPointClient struct {
	cc grpc.ClientConnInterface
}

func NewEntryPointClient(cc grpc.ClientConnInterface) EntryPointClient {
	return &entryPointClient{cc}
}

// EntryPointServer is the server API for EntryPoint service.
// All implementations must embed UnimplementedEntryPointServer
// for forward compatibility
type EntryPointServer interface {
	mustEmbedUnimplementedEntryPointServer()
}

// UnimplementedEntryPointServer must be embedded to have forward compatible implementations.
type UnimplementedEntryPointServer struct {
}

func (UnimplementedEntryPointServer) mustEmbedUnimplementedEntryPointServer() {}

// UnsafeEntryPointServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntryPointServer will
// result in compilation errors.
type UnsafeEntryPointServer interface {
	mustEmbedUnimplementedEntryPointServer()
}

func RegisterEntryPointServer(s grpc.ServiceRegistrar, srv EntryPointServer) {
	s.RegisterService(&EntryPoint_ServiceDesc, srv)
}

// EntryPoint_ServiceDesc is the grpc.ServiceDesc for EntryPoint service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntryPoint_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entry_point.EntryPoint",
	HandlerType: (*EntryPointServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "entry_point.proto",
}