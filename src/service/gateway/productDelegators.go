package grpcServer

import (
	"context"
	"log"

	"flynndcs.com/flynndcs/grpc-gateway/proto/service"
	"google.golang.org/grpc"
)

func (s *GatewayServer) GetProduct(ctx context.Context, in *service.GetProductRequest) (*service.GetProductResponse, error) {
	productConn, productErr := grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(), grpc.WithInsecure())
	if productErr != nil {
		log.Fatalln("Failed to dial server when creating product client:", productErr)
	}
	getProductResponse, err := service.NewProductClient(productConn).GetProduct(ctx, &service.GetProductRequest{ProductName: in.ProductName})
	if err != nil {
		log.Fatalln("Failed when sending a message with product client:", err)
	}

	return &service.GetProductResponse{ProductName: getProductResponse.GetProductName(), ProductUUID: getProductResponse.GetProductUUID()}, nil
}

func (s *GatewayServer) PutProduct(ctx context.Context, in *service.PutProductRequest) (*service.PutProductResponse, error) {
	productConn, productErr := grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(), grpc.WithInsecure())
	if productErr != nil {
		log.Fatalln("Failed to dial", productErr)
	}
	putProductResponse, err := service.NewProductClient(productConn).PutProduct(ctx, &service.PutProductRequest{ProductName: in.GetProductName()})
	if err != nil {
		log.Fatalln("Failed to send", err)
	}
	return &service.PutProductResponse{ProductName: putProductResponse.ProductName, ProductUUID: putProductResponse.ProductUUID}, nil
}

func (s *GatewayServer) ClearProduct(ctx context.Context, in *service.ClearProductRequest) (*service.ClearProductResponse, error) {

	productConn, productErr := grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(), grpc.WithInsecure())
	if productErr != nil {
		log.Fatalln("Failed to dial", productErr)
	}
	_, err := service.NewProductClient(productConn).ClearProduct(ctx, &service.ClearProductRequest{ProductName: in.GetProductName()})
	if err != nil {
		log.Fatalln("Failed to send", err)
	}
	return &service.ClearProductResponse{}, nil
}
