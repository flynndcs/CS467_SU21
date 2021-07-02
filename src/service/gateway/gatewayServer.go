package grpcServer

import (
	"context"
	"log"

	"CS467_SU21/proto/service"

	"google.golang.org/grpc"
)

//global connection and client for entire service lifecycle
var (
	productConn   = &grpc.ClientConn{}
	productErr    error
	productClient service.ProductClient
)

type GatewayServer struct {
	service.UnimplementedGatewayServer
}

func NewGatewayServer() *GatewayServer {
	return &GatewayServer{}
}

//create connection and client globally rather than per API call
func CreateGRPCConnAndClients() {
	productConn, productErr = grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(), grpc.WithInsecure())
	productClient = service.NewProductClient(productConn)
	if productErr != nil {
		if productErr != nil {
			log.Fatalln("Failed to dial server when creating product client:", productErr)
		}
	}
}
