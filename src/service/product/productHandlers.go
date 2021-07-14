package grpcServer

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"io"
	"log"
	"time"

	"CS467_SU21/proto/service"

	fdbDriver "CS467_SU21/src/store/fdb"
)

func (s *ProductServer) GetProductStatus(ctx context.Context, in *service.ProductStatusRequest) (*service.ProductStatusReply, error) {
	return &service.ProductStatusReply{Status: "PRODUCT STATUS: NORMAL"}, nil
}

func (s *ProductServer) GetSingleProduct(ctx context.Context, in *service.GetSingleProductRequest) (*service.StoredProduct, error) {
	value := fdbDriver.GetSingle(in.Name, in.CategorySequence)
	buffer := bytes.NewBuffer(value)
	dec := gob.NewDecoder(buffer)

	var response service.StoredProduct
	dec.Decode(&response)
	return &response, nil
}

func (s *ProductServer) GetProductsInRange(ctx context.Context, in *service.GetProductsInRangeRequest) (*service.StoredProducts, error) {
	value := fdbDriver.GetAllForRange(in.Range)
	buffer := bytes.NewBuffer(value)
	dec := gob.NewDecoder(buffer)

	var products []*service.StoredProduct

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
	var expiresValue int64
	if in.Expires == nil {
		expiresValue = time.Now().Add(24 * time.Hour).Unix()
	} else {
		expiresValue = *in.Expires
	}

	storedProduct := service.StoredProduct{
		Name:                     in.Name,
		CategorySequence:         in.CategorySequence,
		Expires:                  expiresValue,
		Tags:                     in.Tags,
		Origin:                   in.Origin,
		IntermediateDestinations: in.IntermediateDestinations,
		EndDestinations:          in.EndDestinations,
		QuantityByLocation:       map[string]int64{in.Origin: in.TotalQuantity},
		TotalQuantity:            in.TotalQuantity,
		QuantityInTransit:        0,
		ParentProducts:           in.ParentProducts,
		ChildProducts:            in.ChildProducts,
	}

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(&storedProduct)

	if !fdbDriver.Put(in.Name, in.CategorySequence, buffer.Bytes()) {
		return nil, errors.New(" could not put product into FDB")
	}

	for _, tag := range in.Tags {
		if !fdbDriver.Put(in.Name, []string{tag}, buffer.Bytes()) {
			log.Printf("Could not add record for %v tag to index", tag)
		}
	}
	return &storedProduct, nil
}

func (s *ProductServer) ClearSingleProduct(ctx context.Context, in *service.ClearSingleProductRequest) (*service.ClearSingleProductResponse, error) {
	if !fdbDriver.ClearSingle(in.Name, in.CategorySequence) {
		return nil, errors.New(" could not clear product from FDB")
	}

	for _, tag := range in.Tags {
		if !fdbDriver.ClearSingle(in.Name, []string{tag}) {
			log.Printf("Could not delete record for %v tag from index", tag)
		}
	}

	return &service.ClearSingleProductResponse{DeletedName: in.Name, CategorySequence: in.CategorySequence, Tags: in.Tags}, nil
}
