package grpcServer

import (
	"CS467_SU21/proto/service"
)

type ProductServer struct {
	service.UnimplementedProductServer
}

func NewProductServer() *ProductServer {
	return &ProductServer{}
}
