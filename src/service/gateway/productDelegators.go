package grpcServer

import (
	"context"
	"log"

	"CS467_SU21/proto/service"
)

func (s *GatewayServer) GetProduct(ctx context.Context, in *service.ProductIdentifier) (*service.StoredProduct, error) {
	response, err := ProductClient.GetProduct(ctx, in)
	if err != nil {
		log.Fatalln("Failed when sending a message with product client:", err)
	}

	return response, nil
}

func (s *GatewayServer) GetProducts(ctx context.Context, in *service.GetProductsRequest) (*service.StoredProducts, error) {
	response, err := ProductClient.GetProducts(ctx, in)
	if err != nil {
		log.Fatalln("Failed when sending a message with product client:", err)
	}

	return response, nil
}

func (s *GatewayServer) PutProduct(ctx context.Context, in *service.PutProductRequest) (*service.StoredProduct, error) {
	response, err := ProductClient.PutProduct(ctx, in)
	if err != nil {
		log.Fatalln("Failed to send", err)
	}
	return response, nil
}

func (s *GatewayServer) ClearProduct(ctx context.Context, in *service.ClearProductMessage) (*service.ClearProductMessage, error) {
	response, err := ProductClient.ClearProduct(ctx, in)
	if err != nil {
		log.Fatalln("Failed to send", err)
	}
	return response, nil
}
