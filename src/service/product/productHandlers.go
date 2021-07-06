package grpcServer

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"io"
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

func (s *ProductServer) GetProductsInCategorySequence(ctx context.Context, in *service.GetProductsInCategorySequenceRequest) (*service.StoredProducts, error) {
	value := fdbDriver.GetAllForCategorySequence(in.CategorySequence)
	buffer := bytes.NewBuffer(value)
	dec := gob.NewDecoder(buffer)

	var products []*service.StoredProduct

	// https://stackoverflow.com/questions/45603132/im-getting-extra-data-in-buffer-error-when-trying-to-decode-a-gob-in-golang
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
	}

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(&storedProduct)

	if !fdbDriver.Put(in.Name, in.CategorySequence, buffer.Bytes()) {
		return nil, errors.New(" could not put product into FDB")
	}
	return &storedProduct, nil
}

func (s *ProductServer) ClearSingleProduct(ctx context.Context, in *service.ClearSingleProductRequest) (*service.ClearSingleProductResponse, error) {
	if !fdbDriver.ClearSingle(in.Name, in.CategorySequence) {
		return nil, errors.New(" could not clear product from FDB")
	}
	return &service.ClearSingleProductResponse{DeletedName: in.Name, CategorySequence: in.CategorySequence}, nil
}
