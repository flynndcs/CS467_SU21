package grpcServer

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"io"
	"log"
	"os/exec"

	"flynndcs.com/flynndcs/grpc-gateway/proto/service"
	fdbDriver "flynndcs.com/flynndcs/grpc-gateway/src/store/fdb"
)

func (s *ProductServer) GetProductStatus(ctx context.Context, in *service.ProductStatusRequest) (*service.ProductStatusReply, error) {
	return &service.ProductStatusReply{Status: "PRODUCT STATUS: NORMAL"}, nil
}

func (s *ProductServer) GetSingleProduct(ctx context.Context, in *service.GetSingleProductRequest) (*service.GetSingleProductResponse, error) {
	log.Println(in.Scope)
	value := fdbDriver.GetSingle(in.Scope)
	buffer := bytes.NewBuffer(value)
	dec := gob.NewDecoder(buffer)

	var response service.GetSingleProductResponse
	dec.Decode(&response)
	return &response, nil
}

func (s *ProductServer) GetProductsInScope(ctx context.Context, in *service.GetProductsInScopeRequest) (*service.GetProductsInScopeResponse, error) {
	value := fdbDriver.GetAllForScope(in.Scope)
	buffer := bytes.NewBuffer(value)
	dec := gob.NewDecoder(buffer)

	var products []*service.PutSingleProductResponse //TODO generalize records into a "stored type" independent of GET/PUT/etc

	//ty vm stack overflow https://stackoverflow.com/questions/45603132/im-getting-extra-data-in-buffer-error-when-trying-to-decode-a-gob-in-golang
	var eof error
	for eof != io.EOF {
		var innerResponse service.PutSingleProductResponse
		eof = dec.Decode(&innerResponse)
		if eof != nil {
			continue
		}
		products = append(products, &innerResponse)
	}

	return &service.GetProductsInScopeResponse{Products: products}, nil
}

func (s *ProductServer) PutSingleProduct(ctx context.Context, in *service.PutSingleProductRequest) (*service.PutSingleProductResponse, error) {
	uuidbytes, err := exec.Command("uuidgen").Output()
	uuidbytes = bytes.Trim(uuidbytes, "\n")
	if err != nil {
		log.Fatalf("Could not create uuid")
		return nil, err
	}
	uuidString := string(uuidbytes)
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(service.PutSingleProductResponse{ProductName: in.Scope[0], ProductUUID: uuidString})

	if !fdbDriver.Put(in.Scope, buffer.Bytes()) {
		return nil, errors.New(" could not put product into FDB")
	}
	return &service.PutSingleProductResponse{ProductName: in.Scope[0], ProductUUID: uuidString}, nil
}

func (s *ProductServer) ClearSingleProduct(ctx context.Context, in *service.ClearSingleProductRequest) (*service.ClearSingleProductResponse, error) {
	if !fdbDriver.ClearSingle(in.Scope) {
		return nil, errors.New(" could not clear product from FDB")
	}
	return &service.ClearSingleProductResponse{}, nil
}
