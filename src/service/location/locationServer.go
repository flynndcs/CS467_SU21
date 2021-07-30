package grpcServer

import (
	"CS467_SU21/proto/service"
)

type LocationServer struct {
	service.UnimplementedLocationServer
}

func NewLocationServer() *LocationServer {
	return &LocationServer{}
}
