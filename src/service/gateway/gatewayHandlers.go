package grpcServer

import (
	"context"
	"log"

	"flynndcs.com/flynndcs/grpc-gateway/proto/service"
	"google.golang.org/grpc"
)

func (s *GatewayServer) GetStatus(ctx context.Context, in *service.StatusRequest) (*service.StatusReply, error) {
	productConn, productErr := grpc.Dial("0.0.0.0:8080", grpc.WithBlock(), grpc.WithInsecure())
	if productErr != nil {
		log.Fatalln("Failed to dial server when creating product client:", productErr)
	}

	productStatusReply, productErr := service.NewProductClient(productConn).GetProductStatus(ctx, &service.ProductStatusRequest{})
	if productErr != nil {
		log.Fatalln("Failed when sending a message with product client:", productErr)
	}
	return &service.StatusReply{Status: "GATEWAY STATUS: NORMAL, " + productStatusReply.Status}, nil
}
