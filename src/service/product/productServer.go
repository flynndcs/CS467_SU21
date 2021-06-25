package grpcServer

import (
	"flynndcs.com/flynndcs/grpc-gateway/proto/service"
)

type ProductServer struct {
	service.UnimplementedProductServer
}

func NewProductServer() *ProductServer {
	return &ProductServer{}
}
