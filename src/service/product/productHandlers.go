package grpcServer

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"log"
	"os/exec"

	"flynndcs.com/flynndcs/grpc-gateway/proto/service"
	fdbDriver "flynndcs.com/flynndcs/grpc-gateway/src/store/fdb"
)

func (s *ProductServer) GetProductStatus(ctx context.Context, in *service.ProductStatusRequest) (*service.ProductStatusReply, error) {
	return &service.ProductStatusReply{Status: "PRODUCT STATUS: NORMAL"}, nil
}

func (s *ProductServer) GetProduct(ctx context.Context, in *service.GetProductRequest) (*service.GetProductResponse, error) {
	value := fdbDriver.Get(in.ProductName)
	buffer := bytes.NewBuffer(value)
	dec := gob.NewDecoder(buffer)

	var response service.GetProductResponse
	dec.Decode(&response)
	return &response, nil
}

func (s *ProductServer) PutProduct(ctx context.Context, in *service.PutProductRequest) (*service.PutProductResponse, error) {
	uuidbytes, err := exec.Command("uuidgen").Output()
	uuidbytes = bytes.Trim(uuidbytes, "\n")
	if err != nil {
		log.Fatalf("Could not create uuid")
		return nil, err
	}
	uuidString := string(uuidbytes)
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(service.PutProductResponse{ProductName: in.ProductName, ProductUUID: uuidString})

	if !fdbDriver.Put(in.ProductName, buffer.Bytes()) {
		return nil, errors.New(" could not put product into FDB")
	}
	return &service.PutProductResponse{ProductName: in.ProductName, ProductUUID: uuidString}, nil
}

func (s *ProductServer) ClearProduct(ctx context.Context, in *service.ClearProductRequest) (*service.ClearProductResponse, error) {
	if !fdbDriver.Clear(in.ProductName) {
		return nil, errors.New(" could not clear product from FDB")
	}
	return &service.ClearProductResponse{}, nil
}
