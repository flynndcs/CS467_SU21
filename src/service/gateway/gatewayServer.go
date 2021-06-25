package grpcServer

import (
	"flynndcs.com/flynndcs/grpc-gateway/proto/service"
)

type GatewayServer struct {
	service.UnimplementedGatewayServer
}

func NewGatewayServer() *GatewayServer {
	return &GatewayServer{}
}
