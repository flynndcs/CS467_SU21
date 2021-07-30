package grpcServer

import (
	"context"
	"log"

	"CS467_SU21/proto/service"
)

func (s *GatewayServer) GetStatus(ctx context.Context, in *service.StatusRequest) (*service.StatusReply, error) {
	productStatusReply, productErr := ProductClient.GetProductStatus(ctx, &service.ProductStatusRequest{})
	if productErr != nil {
		log.Fatalln("Failed when sending a message with product client:", productErr)
	}

	return &service.StatusReply{Status: "GATEWAY STATUS: NORMAL, " + productStatusReply.Status}, nil
}
