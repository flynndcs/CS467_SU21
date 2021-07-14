package grpcServer

import (
	"context"
	"log"

	"CS467_SU21/proto/service"
)

func (s *GatewayServer) GetSingleProduct(ctx context.Context, in *service.GetSingleProductRequest) (*service.StoredProduct, error) {
	response, err := productClient.GetSingleProduct(ctx, in)
	if err != nil {
		log.Fatalln("Failed when sending a message with product client:", err)
	}

	return response, nil
}

func (s *GatewayServer) GetProductsInRange(ctx context.Context, in *service.GetProductsInRangeRequest) (*service.StoredProducts, error) {
	response, err := productClient.GetProductsInRange(ctx, in)
	if err != nil {
		log.Fatalln("Failed when sending a message with product client:", err)
	}

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
	response, err := productClient.ClearSingleProduct(ctx, in)
	if err != nil {
		log.Fatalln("Failed to send", err)
	}
	return response, nil
}
