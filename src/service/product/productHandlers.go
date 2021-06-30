package grpcServer

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"io"
	"log"
	"os/exec"
	"time"

	"flynndcs.com/flynndcs/grpc-gateway/proto/service"
	fdbDriver "flynndcs.com/flynndcs/grpc-gateway/src/store/fdb"
)

func (s *ProductServer) GetProductStatus(ctx context.Context, in *service.ProductStatusRequest) (*service.ProductStatusReply, error) {
	return &service.ProductStatusReply{Status: "PRODUCT STATUS: NORMAL"}, nil
}

func (s *ProductServer) GetSingleProduct(ctx context.Context, in *service.GetSingleProductRequest) (*service.StoredProduct, error) {
	value := fdbDriver.GetSingle(in.Name, in.Scope)
	buffer := bytes.NewBuffer(value)
	dec := gob.NewDecoder(buffer)

	var response service.StoredProduct
	dec.Decode(&response)
	return &response, nil
}

func (s *ProductServer) GetProductsInScope(ctx context.Context, in *service.GetProductsInScopeRequest) (*service.StoredProducts, error) {
	value := fdbDriver.GetAllForScope(in.Scope)
	buffer := bytes.NewBuffer(value)
	dec := gob.NewDecoder(buffer)

	var products []*service.StoredProduct //TODO generalize records into a "stored type" independent of GET/PUT/etc

	//ty vm stack overflow https://stackoverflow.com/questions/45603132/im-getting-extra-data-in-buffer-error-when-trying-to-decode-a-gob-in-golang
	var eof error
	for eof != io.EOF {
		var innerResponse service.StoredProduct
		eof = dec.Decode(&innerResponse)
		if eof != nil {
			continue
		}
		products = append(products, &innerResponse)
	}

	return &service.StoredProducts{Products: products}, nil
}

func (s *ProductServer) PutSingleProduct(ctx context.Context, in *service.PutSingleProductRequest) (*service.StoredProduct, error) {
	uuidbytes, err := exec.Command("uuidgen").Output()
	uuidbytes = bytes.Trim(uuidbytes, "\n")
	if err != nil {
		log.Fatalf("Could not create uuid")
		return nil, err
	}
	uuidString := string(uuidbytes)
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	var expiresValue int64
	if in.Expires == nil {
		expiresValue = time.Now().Add(24 * time.Hour).Unix()
	} else {
		expiresValue = *in.Expires
	}

	enc.Encode(service.StoredProduct{Name: in.Name, Scope: in.Scope, Data: uuidString, Expires: expiresValue})

	if !fdbDriver.Put(in.Name, in.Scope, buffer.Bytes()) {
		return nil, errors.New(" could not put product into FDB")
	}
	return &service.StoredProduct{Name: in.Name, Data: uuidString, Scope: in.Scope, Expires: expiresValue}, nil
}

func (s *ProductServer) ClearSingleProduct(ctx context.Context, in *service.ClearSingleProductRequest) (*service.ClearSingleProductResponse, error) {
	if !fdbDriver.ClearSingle(in.Name, in.Scope) {
		return nil, errors.New(" could not clear product from FDB")
	}
	return &service.ClearSingleProductResponse{Deleted: in.Name, Scope: in.Scope}, nil
}
