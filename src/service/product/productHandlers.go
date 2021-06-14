package grpcServer

import (
	"context"

	gatewaypb "flynndcs.com/flynndcs/grpc-gateway/proto/service"
)

//	Product service returns a reply directly (this is where service specific logic would occur)
func (s *ProductServer) ProductStatus(ctx context.Context, in *gatewaypb.ProductStatusRequest) (*gatewaypb.ProductStatusReply, error) {
	return &gatewaypb.ProductStatusReply{Status: "PRODUCT STATUS: NORMAL"}, nil
}
