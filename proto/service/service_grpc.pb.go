// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
	GetSingleProduct(ctx context.Context, in *GetSingleProductRequest, opts ...grpc.CallOption) (*StoredProduct, error)
	GetProductsInScope(ctx context.Context, in *GetProductsInScopeRequest, opts ...grpc.CallOption) (*StoredProducts, error)
	PutSingleProduct(ctx context.Context, in *PutSingleProductRequest, opts ...grpc.CallOption) (*StoredProduct, error)
	ClearSingleProduct(ctx context.Context, in *ClearSingleProductRequest, opts ...grpc.CallOption) (*ClearSingleProductResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/service.Gateway/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetSingleProduct(ctx context.Context, in *GetSingleProductRequest, opts ...grpc.CallOption) (*StoredProduct, error) {
	out := new(StoredProduct)
	err := c.cc.Invoke(ctx, "/service.Gateway/GetSingleProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetProductsInScope(ctx context.Context, in *GetProductsInScopeRequest, opts ...grpc.CallOption) (*StoredProducts, error) {
	out := new(StoredProducts)
	err := c.cc.Invoke(ctx, "/service.Gateway/GetProductsInScope", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) PutSingleProduct(ctx context.Context, in *PutSingleProductRequest, opts ...grpc.CallOption) (*StoredProduct, error) {
	out := new(StoredProduct)
	err := c.cc.Invoke(ctx, "/service.Gateway/PutSingleProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) ClearSingleProduct(ctx context.Context, in *ClearSingleProductRequest, opts ...grpc.CallOption) (*ClearSingleProductResponse, error) {
	out := new(ClearSingleProductResponse)
	err := c.cc.Invoke(ctx, "/service.Gateway/ClearSingleProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetStatus(context.Context, *StatusRequest) (*StatusReply, error)
	GetSingleProduct(context.Context, *GetSingleProductRequest) (*StoredProduct, error)
	GetProductsInScope(context.Context, *GetProductsInScopeRequest) (*StoredProducts, error)
	PutSingleProduct(context.Context, *PutSingleProductRequest) (*StoredProduct, error)
	ClearSingleProduct(context.Context, *ClearSingleProductRequest) (*ClearSingleProductResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetStatus(context.Context, *StatusRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedGatewayServer) GetSingleProduct(context.Context, *GetSingleProductRequest) (*StoredProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleProduct not implemented")
}
func (UnimplementedGatewayServer) GetProductsInScope(context.Context, *GetProductsInScopeRequest) (*StoredProducts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsInScope not implemented")
}
func (UnimplementedGatewayServer) PutSingleProduct(context.Context, *PutSingleProductRequest) (*StoredProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutSingleProduct not implemented")
}
func (UnimplementedGatewayServer) ClearSingleProduct(context.Context, *ClearSingleProductRequest) (*ClearSingleProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearSingleProduct not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Gateway/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetSingleProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSingleProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetSingleProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Gateway/GetSingleProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetSingleProduct(ctx, req.(*GetSingleProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetProductsInScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsInScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetProductsInScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Gateway/GetProductsInScope",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetProductsInScope(ctx, req.(*GetProductsInScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_PutSingleProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutSingleProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).PutSingleProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Gateway/PutSingleProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).PutSingleProduct(ctx, req.(*PutSingleProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_ClearSingleProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearSingleProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).ClearSingleProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Gateway/ClearSingleProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).ClearSingleProduct(ctx, req.(*ClearSingleProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _Gateway_GetStatus_Handler,
		},
		{
			MethodName: "GetSingleProduct",
			Handler:    _Gateway_GetSingleProduct_Handler,
		},
		{
			MethodName: "GetProductsInScope",
			Handler:    _Gateway_GetProductsInScope_Handler,
		},
		{
			MethodName: "PutSingleProduct",
			Handler:    _Gateway_PutSingleProduct_Handler,
		},
		{
			MethodName: "ClearSingleProduct",
			Handler:    _Gateway_ClearSingleProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/service.proto",
}

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	GetProductStatus(ctx context.Context, in *ProductStatusRequest, opts ...grpc.CallOption) (*ProductStatusReply, error)
	GetSingleProduct(ctx context.Context, in *GetSingleProductRequest, opts ...grpc.CallOption) (*StoredProduct, error)
	GetProductsInScope(ctx context.Context, in *GetProductsInScopeRequest, opts ...grpc.CallOption) (*StoredProducts, error)
	PutSingleProduct(ctx context.Context, in *PutSingleProductRequest, opts ...grpc.CallOption) (*StoredProduct, error)
	ClearSingleProduct(ctx context.Context, in *ClearSingleProductRequest, opts ...grpc.CallOption) (*ClearSingleProductResponse, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) GetProductStatus(ctx context.Context, in *ProductStatusRequest, opts ...grpc.CallOption) (*ProductStatusReply, error) {
	out := new(ProductStatusReply)
	err := c.cc.Invoke(ctx, "/service.Product/GetProductStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetSingleProduct(ctx context.Context, in *GetSingleProductRequest, opts ...grpc.CallOption) (*StoredProduct, error) {
	out := new(StoredProduct)
	err := c.cc.Invoke(ctx, "/service.Product/GetSingleProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetProductsInScope(ctx context.Context, in *GetProductsInScopeRequest, opts ...grpc.CallOption) (*StoredProducts, error) {
	out := new(StoredProducts)
	err := c.cc.Invoke(ctx, "/service.Product/GetProductsInScope", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) PutSingleProduct(ctx context.Context, in *PutSingleProductRequest, opts ...grpc.CallOption) (*StoredProduct, error) {
	out := new(StoredProduct)
	err := c.cc.Invoke(ctx, "/service.Product/PutSingleProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ClearSingleProduct(ctx context.Context, in *ClearSingleProductRequest, opts ...grpc.CallOption) (*ClearSingleProductResponse, error) {
	out := new(ClearSingleProductResponse)
	err := c.cc.Invoke(ctx, "/service.Product/ClearSingleProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility
type ProductServer interface {
	GetProductStatus(context.Context, *ProductStatusRequest) (*ProductStatusReply, error)
	GetSingleProduct(context.Context, *GetSingleProductRequest) (*StoredProduct, error)
	GetProductsInScope(context.Context, *GetProductsInScopeRequest) (*StoredProducts, error)
	PutSingleProduct(context.Context, *PutSingleProductRequest) (*StoredProduct, error)
	ClearSingleProduct(context.Context, *ClearSingleProductRequest) (*ClearSingleProductResponse, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (UnimplementedProductServer) GetProductStatus(context.Context, *ProductStatusRequest) (*ProductStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductStatus not implemented")
}
func (UnimplementedProductServer) GetSingleProduct(context.Context, *GetSingleProductRequest) (*StoredProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleProduct not implemented")
}
func (UnimplementedProductServer) GetProductsInScope(context.Context, *GetProductsInScopeRequest) (*StoredProducts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsInScope not implemented")
}
func (UnimplementedProductServer) PutSingleProduct(context.Context, *PutSingleProductRequest) (*StoredProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutSingleProduct not implemented")
}
func (UnimplementedProductServer) ClearSingleProduct(context.Context, *ClearSingleProductRequest) (*ClearSingleProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearSingleProduct not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_GetProductStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetProductStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Product/GetProductStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetProductStatus(ctx, req.(*ProductStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetSingleProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSingleProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetSingleProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Product/GetSingleProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetSingleProduct(ctx, req.(*GetSingleProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetProductsInScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsInScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetProductsInScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Product/GetProductsInScope",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetProductsInScope(ctx, req.(*GetProductsInScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_PutSingleProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutSingleProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).PutSingleProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Product/PutSingleProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).PutSingleProduct(ctx, req.(*PutSingleProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ClearSingleProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearSingleProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ClearSingleProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Product/ClearSingleProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ClearSingleProduct(ctx, req.(*ClearSingleProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductStatus",
			Handler:    _Product_GetProductStatus_Handler,
		},
		{
			MethodName: "GetSingleProduct",
			Handler:    _Product_GetSingleProduct_Handler,
		},
		{
			MethodName: "GetProductsInScope",
			Handler:    _Product_GetProductsInScope_Handler,
		},
		{
			MethodName: "PutSingleProduct",
			Handler:    _Product_PutSingleProduct_Handler,
		},
		{
			MethodName: "ClearSingleProduct",
			Handler:    _Product_ClearSingleProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/service.proto",
}