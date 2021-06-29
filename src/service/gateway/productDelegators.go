package grpcServer

import (
	"context"
	"log"

	"flynndcs.com/flynndcs/grpc-gateway/proto/service"
)

//this method called by gateway REST proxy via get product endpoint, uses gRPC side ProductClient to call GetProduct
func (s *GatewayServer) GetSingleProduct(ctx context.Context, in *service.GetSingleProductRequest) (*service.StoredProduct, error) {
	//use product client to call GetProduct method defined in handler
	response, err := productClient.GetSingleProduct(ctx, in)
	if err != nil {
		log.Fatalln("Failed when sending a message with product client:", err)
	}

	//return a response message using the response from GetProduct
	return response, nil
}

func (s *GatewayServer) GetProductsInScope(ctx context.Context, in *service.GetProductsInScopeRequest) (*service.StoredProducts, error) {
	//use product client to call GetProduct method defined in handler
	response, err := productClient.GetProductsInScope(ctx, in)
	if err != nil {
		log.Fatalln("Failed when sending a message with product client:", err)
	}

	//return a response message using the response from GetProduct
	return response, nil
}

func (s *GatewayServer) PutSingleProduct(ctx context.Context, in *service.PutSingleProductRequest) (*service.StoredProduct, error) {
	response, err := productClient.PutSingleProduct(ctx, in)
	if err != nil {
		log.Fatalln("Failed to send", err)
	}
	return response, nil
}

func (s *GatewayServer) ClearSingleProduct(ctx context.Context, in *service.ClearSingleProductRequest) (*service.ClearSingleProductResponse, error) {
	_, err := productClient.ClearSingleProduct(ctx, in)
	if err != nil {
		log.Fatalln("Failed to send", err)
	}
	return &service.ClearSingleProductResponse{}, nil
}
