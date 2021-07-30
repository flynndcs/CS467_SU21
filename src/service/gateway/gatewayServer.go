package grpcServer

import (
	"context"
	"log"

	"CS467_SU21/proto/service"

	"google.golang.org/grpc"
)

var (
	conn           = &grpc.ClientConn{}
	serviceErr     error
	ProductClient  service.ProductClient
	LocationClient service.LocationClient
)

type GatewayServer struct {
	service.UnimplementedGatewayServer
}

func NewGatewayServer() *GatewayServer {
	return &GatewayServer{}
}

func CreateGRPCConnAndClients() {
	conn, serviceErr = grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(), grpc.WithInsecure())
	ProductClient = service.NewProductClient(conn)
	LocationClient = service.NewLocationClient(conn)
	if serviceErr != nil {
		log.Fatalln("Failed to dial server when creating product client:", serviceErr)
	}
}
