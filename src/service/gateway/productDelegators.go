package grpcServer

import (
	"context"
	"log"

	"flynndcs.com/flynndcs/grpc-gateway/proto/service"
)

//this method called by gateway REST proxy via get product endpoint, uses gRPC side ProductClient to call GetProduct
func (s *GatewayServer) GetSingleProduct(ctx context.Context, in *service.GetSingleProductRequest) (*service.GetSingleProductResponse, error) {
	//use product client to call GetProduct method defined in handler
	log.Println(in.Scope)
	GetSingleProductResponse, err := productClient.GetSingleProduct(ctx, &service.GetSingleProductRequest{Scope: in.Scope})
	if err != nil {
		log.Fatalln("Failed when sending a message with product client:", err)
	}

	//return a response message using the response from GetProduct
	return &service.GetSingleProductResponse{ProductName: GetSingleProductResponse.GetProductName(), ProductUUID: GetSingleProductResponse.GetProductUUID()}, nil
}

func (s *GatewayServer) GetProductsInScope(ctx context.Context, in *service.GetProductsInScopeRequest) (*service.GetProductsInScopeResponse, error) {
	//use product client to call GetProduct method defined in handler
	GetProductsInScopeResponse, err := productClient.GetProductsInScope(ctx, &service.GetProductsInScopeRequest{Scope: in.Scope})
	if err != nil {
		log.Fatalln("Failed when sending a message with product client:", err)
	}

	//return a response message using the response from GetProduct
	return &service.GetProductsInScopeResponse{Products: GetProductsInScopeResponse.GetProducts()}, nil
}

func (s *GatewayServer) PutSingleProduct(ctx context.Context, in *service.PutSingleProductRequest) (*service.PutSingleProductResponse, error) {
	PutSingleProductResponse, err := productClient.PutSingleProduct(ctx, &service.PutSingleProductRequest{Scope: in.Scope})
	if err != nil {
		log.Fatalln("Failed to send", err)
	}
	return &service.PutSingleProductResponse{ProductName: PutSingleProductResponse.ProductName, ProductUUID: PutSingleProductResponse.ProductUUID}, nil
}

func (s *GatewayServer) ClearSingleProduct(ctx context.Context, in *service.ClearSingleProductRequest) (*service.ClearSingleProductResponse, error) {
	_, err := productClient.ClearSingleProduct(ctx, &service.ClearSingleProductRequest{Scope: in.Scope})
	if err != nil {
		log.Fatalln("Failed to send", err)
	}
	return &service.ClearSingleProductResponse{}, nil
}
