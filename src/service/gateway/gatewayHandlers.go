package grpcServer

import (
	"context"
	"log"

	gatewaypb "flynndcs.com/flynndcs/grpc-gateway/proto/service"
	"google.golang.org/grpc"
)

/*
	SERVICE LEVEL RPC HANDLERS
*/

//	Gateway Service

//	Connect to product service, send and receive message as a client
func (s *GatewayServer) GetStatus(ctx context.Context, in *gatewaypb.StatusRequest) (*gatewaypb.StatusReply, error) {
	productConn, productErr := grpc.Dial("0.0.0.0:8080", grpc.WithBlock(), grpc.WithInsecure())
	if productErr != nil {
		log.Fatalln("Failed to dial server when creating product client:", productErr)
	}

	productStatusReply, productErr := gatewaypb.NewProductClient(productConn).GetProductStatus(ctx, &gatewaypb.ProductStatusRequest{})
	if productErr != nil {
		log.Fatalln("Failed when sending a message with product client:", productErr)
	}
	return &gatewaypb.StatusReply{Status: "GATEWAY STATUS: NORMAL, " + productStatusReply.Status}, nil
}

func (s *GatewayServer) GetProduct(ctx context.Context, in *gatewaypb.GetProductRequest) (*gatewaypb.GetProductResponse, error) {
	productConn, productErr := grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(), grpc.WithInsecure())
	if productErr != nil {
		log.Fatalln("Failed to dial server when creating product client:", productErr)
	}
	getProductResponse, err := gatewaypb.NewProductClient(productConn).GetProduct(ctx, &gatewaypb.GetProductRequest{ProductName: in.ProductName})
	if err != nil {
		log.Fatalln("Failed when sending a message with product client:", err)
	}

	return &gatewaypb.GetProductResponse{ProductName: getProductResponse.GetProductName(), ProductUUID: getProductResponse.GetProductUUID()}, nil
}

func (s *GatewayServer) PutProduct(ctx context.Context, in *gatewaypb.PutProductRequest) (*gatewaypb.PutProductResponse, error) {
	log.Default().Println("inside gateway putproducthandler")
	productConn, productErr := grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(), grpc.WithInsecure())
	if productErr != nil {
		log.Fatalln("Failed to dial", productErr)
	}
	putProductResponse, err := gatewaypb.NewProductClient(productConn).PutProduct(ctx, &gatewaypb.PutProductRequest{ProductName: in.GetProductName()})
	if err != nil {
		log.Fatalln("Failed to send", err)
	}
	return &gatewaypb.PutProductResponse{ProductName: putProductResponse.ProductName, ProductUUID: putProductResponse.ProductUUID}, nil
}

func (s *GatewayServer) ClearProduct(ctx context.Context, in *gatewaypb.ClearProductRequest) (*gatewaypb.ClearProductResponse, error) {

	productConn, productErr := grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(), grpc.WithInsecure())
	if productErr != nil {
		log.Fatalln("Failed to dial", productErr)
	}
	_, err := gatewaypb.NewProductClient(productConn).ClearProduct(ctx, &gatewaypb.ClearProductRequest{ProductName: in.GetProductName()})
	if err != nil {
		log.Fatalln("Failed to send", err)
	}
	return &gatewaypb.ClearProductResponse{}, nil
}
