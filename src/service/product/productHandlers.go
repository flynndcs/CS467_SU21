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

func (s *ProductServer) GetProduct(ctx context.Context, in *service.GetProductRequest) (*service.GetProductResponse, error) {
	value := fdbDriver.Get(in.ProductName)
	buffer := bytes.NewBuffer(value)
	dec := gob.NewDecoder(buffer)

	var response service.GetProductResponse
	dec.Decode(&response)
	return &response, nil
}

func (s *ProductServer) GetProductRange(ctx context.Context, in *service.GetProductRangeRequest) (*service.GetProductRangeResponse, error) {
	value := fdbDriver.GetRange(in.BeginProductName, in.EndProductName)
	buffer := bytes.NewBuffer(value)
	dec := gob.NewDecoder(buffer)

	var products []*service.PutProductResponse //TODO generalize records into a "stored type" independent of GET/PUT/etc

	//ty vm stack overflow https://stackoverflow.com/questions/45603132/im-getting-extra-data-in-buffer-error-when-trying-to-decode-a-gob-in-golang
	var eof error
	for eof != io.EOF {
		var innerResponse service.PutProductResponse
		eof = dec.Decode(&innerResponse)
		if eof != nil {
			continue
		}
		products = append(products, &innerResponse)
	}

	return &service.GetProductRangeResponse{Products: products}, nil
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
