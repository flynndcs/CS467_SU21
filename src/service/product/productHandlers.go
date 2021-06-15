package grpcServer

import (
	"context"
	"time"

	gatewaypb "flynndcs.com/flynndcs/grpc-gateway/proto/service"
	fdb "flynndcs.com/flynndcs/grpc-gateway/src/store/fdb"
)

//	Product service returns a reply directly (this is where service specific logic would occur)
func (s *ProductServer) ProductStatus(ctx context.Context, in *gatewaypb.ProductStatusRequest) (*gatewaypb.ProductStatusReply, error) {
	time := time.Now()
	timeString := time.String()
	fdb.Put(timeString, []byte(timeString))
	value := fdb.Get(timeString)
	return &gatewaypb.ProductStatusReply{Status: "PRODUCT STATUS: NORMAL, KEY: " + timeString + ", VALUE: " + string(value)}, nil
}
