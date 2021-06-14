package grpcServer

import (
	"context"
	"log"

	gatewaypb "flynndcs.com/flynndcs/grpc-gateway/proto/service"
	"google.golang.org/grpc"
)

/*
	SERVICE LEVEL RPC HANDLERS
*/

//	Gateway Service

//	Connect to product service, send and receive message as a client
func (s *GatewayServer) Status(ctx context.Context, in *gatewaypb.StatusRequest) (*gatewaypb.StatusReply, error) {
	productConn, productErr := grpc.Dial("0.0.0.0:8080", grpc.WithBlock(), grpc.WithInsecure())
	if productErr != nil {
		log.Fatalln("Failed to dial server when creating product client:", productErr)
	}

	productStatusReply, productErr := gatewaypb.NewProductClient(productConn).ProductStatus(ctx, &gatewaypb.ProductStatusRequest{})
	if productErr != nil {
		log.Fatalln("Failed when sending a message with product client:", productErr)
	}
	return &gatewaypb.StatusReply{Status: "GATEWAY STATUS: NORMAL, " + productStatusReply.Status}, nil
}
