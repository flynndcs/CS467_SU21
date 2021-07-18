package grpcServer

import (
	"bytes"
	"container/list"
	"context"
	"encoding/gob"
	"errors"
	"io"
	"log"

	"CS467_SU21/proto/service"

	fdbDriver "CS467_SU21/src/store/fdb"
)

func (s *ProductServer) GetProductStatus(ctx context.Context, in *service.ProductStatusRequest) (*service.ProductStatusReply, error) {
	return &service.ProductStatusReply{Status: "PRODUCT STATUS: NORMAL"}, nil
}

func (s *ProductServer) GetSingleProduct(ctx context.Context, in *service.ProductIdentifier) (*service.StoredProduct, error) {
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
	var fullProductFamily service.FullProductFamily
	familyQueue := list.New()

	if in.LocalProductFamily != nil {
		fullProductFamily.Self = in.ProductIdentifier
		fullProductFamily.LocalProductFamilies = append(fullProductFamily.LocalProductFamilies, in.LocalProductFamily)
		if in.LocalProductFamily.ParentProducts != nil {
			for _, parent := range in.LocalProductFamily.ParentProducts {
				familyQueue.PushBack(parent)
			}
			for familyQueue.Len() > 0 {
				current := familyQueue.Front().Value.(*service.ProductIdentifier)
				product, productError := s.GetSingleProduct(ctx, current)
				if productError != nil {
					log.Printf("Could not get parents from product %v", current.Name)
				}
				fullProductFamily.LocalProductFamilies = append(fullProductFamily.LocalProductFamilies, product.LocalProductFamily)
				if product.LocalProductFamily.ParentProducts != nil {
					for _, parent := range product.LocalProductFamily.ParentProducts {
						familyQueue.PushBack(parent)
					}
				}
				familyQueue.Remove(familyQueue.Front())
			}
		}
		if in.LocalProductFamily.ChildProducts != nil {
			for _, child := range in.LocalProductFamily.ChildProducts {
				familyQueue.PushBack(child)
			}
			for familyQueue.Len() > 0 {
				current := familyQueue.Front().Value.(*service.ProductIdentifier)
				product, productError := s.GetSingleProduct(ctx, current)
				if productError != nil {
					log.Printf("Could not get parents from product %v", current.Name)
				}
				fullProductFamily.LocalProductFamilies = append(fullProductFamily.LocalProductFamilies, product.LocalProductFamily)
				if product.LocalProductFamily.ChildProducts != nil {
					for _, child := range product.LocalProductFamily.ChildProducts {
						familyQueue.PushBack(child)
					}
				}
				familyQueue.Remove(familyQueue.Front())
			}
		}
	}

	storedProduct := service.StoredProduct{
		ProductIdentifier:        in.ProductIdentifier,
		Tags:                     in.Tags,
		Origin:                   in.Origin,
		IntermediateDestinations: in.IntermediateDestinations,
		EndDestinations:          in.EndDestinations,
		QuantityByLocation:       map[string]int64{in.Origin: in.TotalQuantity},
		TotalQuantity:            in.TotalQuantity,
		QuantityInTransit:        0,
		LocalProductFamily:       in.LocalProductFamily,
		FullProductFamily:        &fullProductFamily,
	}

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(&storedProduct)

	if !fdbDriver.Put(in.ProductIdentifier.Name, in.ProductIdentifier.CategorySequence, buffer.Bytes()) {
		return nil, errors.New(" could not put product into FDB")
	}

	for _, tag := range in.Tags {
		if !fdbDriver.Put(in.ProductIdentifier.Name, []string{tag}, buffer.Bytes()) {
			log.Printf("Could not add record for %v tag to index", tag)
		}
	}
	return &storedProduct, nil
}

func (s *ProductServer) ClearSingleProduct(ctx context.Context, in *service.ClearSingleProductMessage) (*service.ClearSingleProductMessage, error) {
	if !fdbDriver.ClearSingle(in.Name, in.CategorySequence) {
		return nil, errors.New(" could not clear product from FDB")
	}

	for _, tag := range in.Tags {
		if !fdbDriver.ClearSingle(in.Name, []string{tag}) {
			log.Printf("Could not delete record for %v tag from index", tag)
		}
	}

	return &service.ClearSingleProductMessage{Name: in.Name, CategorySequence: in.CategorySequence, Tags: in.Tags}, nil
}
