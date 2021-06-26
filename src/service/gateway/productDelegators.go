package grpcServer

import (
	"context"
	"log"

	"flynndcs.com/flynndcs/grpc-gateway/proto/service"
)

//this method called by gateway REST proxy via get product endpoint, uses gRPC side ProductClient to call GetProduct
func (s *GatewayServer) GetProduct(ctx context.Context, in *service.GetProductRequest) (*service.GetProductResponse, error) {
	//use product client to call GetProduct method defined in handler
	getProductResponse, err := productClient.GetProduct(ctx, &service.GetProductRequest{ProductName: in.ProductName})
	if err != nil {
		log.Fatalln("Failed when sending a message with product client:", err)
	}

	//return a response message using the response from GetProduct
	return &service.GetProductResponse{ProductName: getProductResponse.GetProductName(), ProductUUID: getProductResponse.GetProductUUID()}, nil
}

func (s *GatewayServer) PutProduct(ctx context.Context, in *service.PutProductRequest) (*service.PutProductResponse, error) {
	putProductResponse, err := productClient.PutProduct(ctx, &service.PutProductRequest{ProductName: in.GetProductName()})
	if err != nil {
		log.Fatalln("Failed to send", err)
	}
	return &service.PutProductResponse{ProductName: putProductResponse.ProductName, ProductUUID: putProductResponse.ProductUUID}, nil
}

func (s *GatewayServer) ClearProduct(ctx context.Context, in *service.ClearProductRequest) (*service.ClearProductResponse, error) {
	_, err := productClient.ClearProduct(ctx, &service.ClearProductRequest{ProductName: in.GetProductName()})
	if err != nil {
		log.Fatalln("Failed to send", err)
	}
	return &service.ClearProductResponse{}, nil
}
