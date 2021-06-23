package grpcServer

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"log"
	"os/exec"

	gatewaypb "flynndcs.com/flynndcs/grpc-gateway/proto/service"
	fdb "flynndcs.com/flynndcs/grpc-gateway/src/store/fdb"
)

//	Product service returns a reply directly (this is where service specific logic would occur)
func (s *ProductServer) GetProductStatus(ctx context.Context, in *gatewaypb.ProductStatusRequest) (*gatewaypb.ProductStatusReply, error) {
	return &gatewaypb.ProductStatusReply{Status: "PRODUCT STATUS: NORMAL"}, nil
}

func (s *ProductServer) GetProduct(ctx context.Context, in *gatewaypb.GetProductRequest) (*gatewaypb.GetProductResponse, error) {
	value := fdb.Get(in.ProductName)
	buffer := bytes.NewBuffer(value)
	dec := gob.NewDecoder(buffer)

	var response gatewaypb.GetProductResponse
	dec.Decode(&response)
	return &response, nil
}

func (s *ProductServer) PutProduct(ctx context.Context, in *gatewaypb.PutProductRequest) (*gatewaypb.PutProductResponse, error) {
	log.Default().Println("inside product putproducthandler")
	uuidbytes, err := exec.Command("uuidgen").Output()
	uuidbytes = bytes.Trim(uuidbytes, "\n")
	if err != nil {
		log.Fatalf("Could not create uuid")
		return nil, err
	}
	uuidString := string(uuidbytes)
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(gatewaypb.PutProductResponse{ProductName: in.ProductName, ProductUUID: uuidString})

	if !fdb.Put(in.ProductName, buffer.Bytes()) {
		return nil, errors.New(" could not put product into FDB")
	}
	return &gatewaypb.PutProductResponse{ProductName: in.ProductName, ProductUUID: uuidString}, nil
}
