package grpcServer

import (
	gatewaypb "flynndcs.com/flynndcs/grpc-gateway/proto/service"
)

/*
	PRODUCT SERVICE IMPLEMENTATIONS (?)
*/

type ProductServer struct {
	gatewaypb.UnimplementedProductServer
}

func NewProductServer() *ProductServer {
	return &ProductServer{}
}
