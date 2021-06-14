package grpcServer

import (
	gatewaypb "flynndcs.com/flynndcs/grpc-gateway/proto/gateway"
)

/*
	GATEWAY SERVICE IMPLEMENTATIONS (?)
*/

type GatewayServer struct {
	gatewaypb.UnimplementedGatewayServer
}

func NewGatewayServer() *GatewayServer {
	return &GatewayServer{}
}
