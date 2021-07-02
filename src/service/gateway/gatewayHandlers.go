package grpcServer

import (
	"context"
	"log"

	"CS467_SU21/proto/service"
)

func (s *GatewayServer) GetStatus(ctx context.Context, in *service.StatusRequest) (*service.StatusReply, error) {
	//get product status from product service using product client to call GetProductStatus method
	productStatusReply, productErr := productClient.GetProductStatus(ctx, &service.ProductStatusRequest{})
	if productErr != nil {
		log.Fatalln("Failed when sending a message with product client:", productErr)
	}

	//return an instance of the message with an arbitrary status for the gateway and the status from the product service's reply
	return &service.StatusReply{Status: "GATEWAY STATUS: NORMAL, " + productStatusReply.Status}, nil
}
