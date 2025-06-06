// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.8
// source: kuscia/proto/api/v1alpha1/kusciaapi/domaindata.proto

package kusciaapi

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

const (
	DomainDataService_CreateDomainData_FullMethodName       = "/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataService/CreateDomainData"
	DomainDataService_UpdateDomainData_FullMethodName       = "/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataService/UpdateDomainData"
	DomainDataService_DeleteDomainData_FullMethodName       = "/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataService/DeleteDomainData"
	DomainDataService_DeleteDomainDataAndRaw_FullMethodName = "/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataService/DeleteDomainDataAndRaw"
	DomainDataService_QueryDomainData_FullMethodName        = "/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataService/QueryDomainData"
	DomainDataService_BatchQueryDomainData_FullMethodName   = "/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataService/BatchQueryDomainData"
	DomainDataService_ListDomainData_FullMethodName         = "/kuscia.proto.api.v1alpha1.kusciaapi.DomainDataService/ListDomainData"
)

// DomainDataServiceClient is the client API for DomainDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DomainDataServiceClient interface {
	CreateDomainData(ctx context.Context, in *CreateDomainDataRequest, opts ...grpc.CallOption) (*CreateDomainDataResponse, error)
	UpdateDomainData(ctx context.Context, in *UpdateDomainDataRequest, opts ...grpc.CallOption) (*UpdateDomainDataResponse, error)
	DeleteDomainData(ctx context.Context, in *DeleteDomainDataRequest, opts ...grpc.CallOption) (*DeleteDomainDataResponse, error)
	DeleteDomainDataAndRaw(ctx context.Context, in *DeleteDomainDataRequest, opts ...grpc.CallOption) (*DeleteDomainDataRequest, error)
	QueryDomainData(ctx context.Context, in *QueryDomainDataRequest, opts ...grpc.CallOption) (*QueryDomainDataResponse, error)
	BatchQueryDomainData(ctx context.Context, in *BatchQueryDomainDataRequest, opts ...grpc.CallOption) (*BatchQueryDomainDataResponse, error)
	ListDomainData(ctx context.Context, in *ListDomainDataRequest, opts ...grpc.CallOption) (*ListDomainDataResponse, error)
}

type domainDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDomainDataServiceClient(cc grpc.ClientConnInterface) DomainDataServiceClient {
	return &domainDataServiceClient{cc}
}

func (c *domainDataServiceClient) CreateDomainData(ctx context.Context, in *CreateDomainDataRequest, opts ...grpc.CallOption) (*CreateDomainDataResponse, error) {
	out := new(CreateDomainDataResponse)
	err := c.cc.Invoke(ctx, DomainDataService_CreateDomainData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainDataServiceClient) UpdateDomainData(ctx context.Context, in *UpdateDomainDataRequest, opts ...grpc.CallOption) (*UpdateDomainDataResponse, error) {
	out := new(UpdateDomainDataResponse)
	err := c.cc.Invoke(ctx, DomainDataService_UpdateDomainData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainDataServiceClient) DeleteDomainData(ctx context.Context, in *DeleteDomainDataRequest, opts ...grpc.CallOption) (*DeleteDomainDataResponse, error) {
	out := new(DeleteDomainDataResponse)
	err := c.cc.Invoke(ctx, DomainDataService_DeleteDomainData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainDataServiceClient) DeleteDomainDataAndRaw(ctx context.Context, in *DeleteDomainDataRequest, opts ...grpc.CallOption) (*DeleteDomainDataRequest, error) {
	out := new(DeleteDomainDataRequest)
	err := c.cc.Invoke(ctx, DomainDataService_DeleteDomainDataAndRaw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainDataServiceClient) QueryDomainData(ctx context.Context, in *QueryDomainDataRequest, opts ...grpc.CallOption) (*QueryDomainDataResponse, error) {
	out := new(QueryDomainDataResponse)
	err := c.cc.Invoke(ctx, DomainDataService_QueryDomainData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainDataServiceClient) BatchQueryDomainData(ctx context.Context, in *BatchQueryDomainDataRequest, opts ...grpc.CallOption) (*BatchQueryDomainDataResponse, error) {
	out := new(BatchQueryDomainDataResponse)
	err := c.cc.Invoke(ctx, DomainDataService_BatchQueryDomainData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainDataServiceClient) ListDomainData(ctx context.Context, in *ListDomainDataRequest, opts ...grpc.CallOption) (*ListDomainDataResponse, error) {
	out := new(ListDomainDataResponse)
	err := c.cc.Invoke(ctx, DomainDataService_ListDomainData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainDataServiceServer is the server API for DomainDataService service.
// All implementations must embed UnimplementedDomainDataServiceServer
// for forward compatibility
type DomainDataServiceServer interface {
	CreateDomainData(context.Context, *CreateDomainDataRequest) (*CreateDomainDataResponse, error)
	UpdateDomainData(context.Context, *UpdateDomainDataRequest) (*UpdateDomainDataResponse, error)
	DeleteDomainData(context.Context, *DeleteDomainDataRequest) (*DeleteDomainDataResponse, error)
	DeleteDomainDataAndRaw(context.Context, *DeleteDomainDataRequest) (*DeleteDomainDataRequest, error)
	QueryDomainData(context.Context, *QueryDomainDataRequest) (*QueryDomainDataResponse, error)
	BatchQueryDomainData(context.Context, *BatchQueryDomainDataRequest) (*BatchQueryDomainDataResponse, error)
	ListDomainData(context.Context, *ListDomainDataRequest) (*ListDomainDataResponse, error)
	mustEmbedUnimplementedDomainDataServiceServer()
}

// UnimplementedDomainDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDomainDataServiceServer struct {
}

func (UnimplementedDomainDataServiceServer) CreateDomainData(context.Context, *CreateDomainDataRequest) (*CreateDomainDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDomainData not implemented")
}
func (UnimplementedDomainDataServiceServer) UpdateDomainData(context.Context, *UpdateDomainDataRequest) (*UpdateDomainDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDomainData not implemented")
}
func (UnimplementedDomainDataServiceServer) DeleteDomainData(context.Context, *DeleteDomainDataRequest) (*DeleteDomainDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDomainData not implemented")
}
func (UnimplementedDomainDataServiceServer) DeleteDomainDataAndRaw(context.Context, *DeleteDomainDataRequest) (*DeleteDomainDataRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDomainDataAndRaw not implemented")
}
func (UnimplementedDomainDataServiceServer) QueryDomainData(context.Context, *QueryDomainDataRequest) (*QueryDomainDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDomainData not implemented")
}
func (UnimplementedDomainDataServiceServer) BatchQueryDomainData(context.Context, *BatchQueryDomainDataRequest) (*BatchQueryDomainDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchQueryDomainData not implemented")
}
func (UnimplementedDomainDataServiceServer) ListDomainData(context.Context, *ListDomainDataRequest) (*ListDomainDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDomainData not implemented")
}
func (UnimplementedDomainDataServiceServer) mustEmbedUnimplementedDomainDataServiceServer() {}

// UnsafeDomainDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DomainDataServiceServer will
// result in compilation errors.
type UnsafeDomainDataServiceServer interface {
	mustEmbedUnimplementedDomainDataServiceServer()
}

func RegisterDomainDataServiceServer(s grpc.ServiceRegistrar, srv DomainDataServiceServer) {
	s.RegisterService(&DomainDataService_ServiceDesc, srv)
}

func _DomainDataService_CreateDomainData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDomainDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainDataServiceServer).CreateDomainData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainDataService_CreateDomainData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainDataServiceServer).CreateDomainData(ctx, req.(*CreateDomainDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainDataService_UpdateDomainData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDomainDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainDataServiceServer).UpdateDomainData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainDataService_UpdateDomainData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainDataServiceServer).UpdateDomainData(ctx, req.(*UpdateDomainDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainDataService_DeleteDomainData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDomainDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainDataServiceServer).DeleteDomainData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainDataService_DeleteDomainData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainDataServiceServer).DeleteDomainData(ctx, req.(*DeleteDomainDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainDataService_DeleteDomainDataAndRaw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDomainDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainDataServiceServer).DeleteDomainDataAndRaw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainDataService_DeleteDomainDataAndRaw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainDataServiceServer).DeleteDomainDataAndRaw(ctx, req.(*DeleteDomainDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainDataService_QueryDomainData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDomainDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainDataServiceServer).QueryDomainData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainDataService_QueryDomainData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainDataServiceServer).QueryDomainData(ctx, req.(*QueryDomainDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainDataService_BatchQueryDomainData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchQueryDomainDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainDataServiceServer).BatchQueryDomainData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainDataService_BatchQueryDomainData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainDataServiceServer).BatchQueryDomainData(ctx, req.(*BatchQueryDomainDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainDataService_ListDomainData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDomainDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainDataServiceServer).ListDomainData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainDataService_ListDomainData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainDataServiceServer).ListDomainData(ctx, req.(*ListDomainDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DomainDataService_ServiceDesc is the grpc.ServiceDesc for DomainDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DomainDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kuscia.proto.api.v1alpha1.kusciaapi.DomainDataService",
	HandlerType: (*DomainDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDomainData",
			Handler:    _DomainDataService_CreateDomainData_Handler,
		},
		{
			MethodName: "UpdateDomainData",
			Handler:    _DomainDataService_UpdateDomainData_Handler,
		},
		{
			MethodName: "DeleteDomainData",
			Handler:    _DomainDataService_DeleteDomainData_Handler,
		},
		{
			MethodName: "DeleteDomainDataAndRaw",
			Handler:    _DomainDataService_DeleteDomainDataAndRaw_Handler,
		},
		{
			MethodName: "QueryDomainData",
			Handler:    _DomainDataService_QueryDomainData_Handler,
		},
		{
			MethodName: "BatchQueryDomainData",
			Handler:    _DomainDataService_BatchQueryDomainData_Handler,
		},
		{
			MethodName: "ListDomainData",
			Handler:    _DomainDataService_ListDomainData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kuscia/proto/api/v1alpha1/kusciaapi/domaindata.proto",
}
