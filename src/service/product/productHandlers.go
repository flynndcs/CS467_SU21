package grpcServer

import (
	"bytes"
	"container/list"
	"context"
	"encoding/gob"
	"log"

	"CS467_SU21/proto/service"
	"CS467_SU21/src/store/fdb"
)

func (s *ProductServer) GetProductStatus(ctx context.Context, in *service.ProductStatusRequest) (*service.ProductStatusReply, error) {
	return &service.ProductStatusReply{Status: "PRODUCT STATUS: NORMAL"}, nil
}

func (s *ProductServer) GetProduct(ctx context.Context, in *service.ProductIdentifier) (*service.StoredProduct, error) {
	prefixedIdKey := fdb.ProductSubspace.Bytes()
	prefixedIdKey = append(prefixedIdKey, byte(in.Id))
	value := fdb.Get(prefixedIdKey)
	if value == nil {
		return &service.StoredProduct{}, nil
	}
	buffer := bytes.NewBuffer(value)
	dec := gob.NewDecoder(buffer)

	var response service.StoredProduct
	dec.Decode(&response)
	return &response, nil
}

func (s *ProductServer) GetProducts(ctx context.Context, in *service.GetProductsRequest) (*service.StoredProducts, error) {
	var collectedIdKeys [][]byte
	var resultIdKeys [][]byte

	if in.Origin != "" {
		prefixedOriginKey := fdb.ProductSubspace.Bytes()
		prefixedOriginKey = append(prefixedOriginKey, []byte(in.Origin)...)

		resultIdKeys = fdb.GetRange(prefixedOriginKey)
		if resultIdKeys != nil {
			collectedIdKeys = append(collectedIdKeys, resultIdKeys...)
		}
	}

	if in.Categories != nil {
		prefixedCategoriesKey := fdb.ProductSubspace.Bytes()
		prefixedCategoriesKey = append(prefixedCategoriesKey, fdb.EncodeCategories(in.Categories)...)

		resultIdKeys = fdb.GetRange(prefixedCategoriesKey)
		if resultIdKeys != nil {
			collectedIdKeys = append(collectedIdKeys, resultIdKeys...)
		}

	}

	if in.Tags != nil {
		for _, tag := range in.Tags {
			prefixedTagKey := fdb.ProductSubspace.Bytes()
			prefixedTagKey = append(prefixedTagKey, []byte(tag)...)
			resultIdKeys = fdb.GetRange(prefixedTagKey)
			if resultIdKeys != nil {
				collectedIdKeys = append(collectedIdKeys, resultIdKeys...)
			}
		}
	}

	collectedIdKeys = removeDuplicateKeys(collectedIdKeys)

	var products []*service.StoredProduct

	for _, idKey := range collectedIdKeys {
		product := fdb.Get(idKey)
		if product != nil {
			buffer := bytes.NewBuffer(product)
			dec := gob.NewDecoder(buffer)
			var decodedProduct *service.StoredProduct
			dec.Decode(&decodedProduct)
			products = append(products, decodedProduct)
		}
	}

	return &service.StoredProducts{Products: products}, nil
}

func removeDuplicateKeys(array [][]byte) [][]byte {
	keys := make(map[string]bool)
	list := [][]byte{}

	for _, entry := range array {
		if _, value := keys[string(entry)]; !value {
			keys[string(entry)] = true
			list = append(list, entry)
		}
	}
	return list
}

func (s *ProductServer) PutProduct(ctx context.Context, in *service.PutProductRequest) (*service.StoredProduct, error) {
	var fullProductFamily service.FullProductFamily
	familyQueue := list.New()

	fullProductFamily.Self = in.ProductIdentifier
	fullProductFamily.LocalProductFamilies = append(fullProductFamily.LocalProductFamilies, in.LocalProductFamily)

	if in.LocalProductFamily != nil {
		if in.LocalProductFamily.ParentProducts != nil {
			for _, parent := range in.LocalProductFamily.ParentProducts {
				familyQueue.PushBack(parent)
			}
			for familyQueue.Len() > 0 {
				current := familyQueue.Front().Value.(*service.ProductIdentifier)
				product, productError := s.GetProduct(ctx, current)
				if productError != nil {
					log.Printf("Could not get parents from product %v", current.Id)
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
				product, productError := s.GetProduct(ctx, current)
				if productError != nil {
					log.Printf("Could not get parents from product %v", current.Id)
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
		Name:                     in.Name,
		Categories:               in.Categories,
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

	prefixedIdKey := fdb.ProductSubspace.Bytes()
	prefixedIdKey = append(prefixedIdKey, byte(in.ProductIdentifier.Id))

	if !fdb.Put(prefixedIdKey, buffer.Bytes()) {
		log.Println("Could not put product into FDB")
	}

	prefixedOriginKey := fdb.ProductSubspace.Bytes()
	prefixedOriginKey = append(prefixedOriginKey, []byte(in.Origin)...)
	prefixedOriginKey = append(prefixedOriginKey, byte(in.ProductIdentifier.Id))

	if !fdb.Put(prefixedOriginKey, prefixedIdKey) {
		log.Println("Could not put product into origin index")
	}

	prefixedCategoryKey := fdb.ProductSubspace.Bytes()
	prefixedCategoryKey = append(prefixedCategoryKey, fdb.EncodeCategories(in.Categories)...)
	prefixedCategoryKey = append(prefixedCategoryKey, byte(in.ProductIdentifier.Id))

	if !fdb.Put(prefixedCategoryKey, prefixedIdKey) {
		log.Printf("Could not put product into category index")
	}

	for _, tag := range in.Tags {
		prefixedTagKey := fdb.ProductSubspace.Bytes()
		prefixedTagKey = append(prefixedTagKey, []byte(tag)...)
		prefixedTagKey = append(prefixedTagKey, byte(in.ProductIdentifier.Id))

		if !fdb.Put(prefixedTagKey, prefixedIdKey) {
			log.Printf("Could not put product into tag index")
		}
	}
	return &storedProduct, nil
}

func (s *ProductServer) UpdateProduct(ctx context.Context, in *service.StoredProduct) (*service.StoredProduct, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(&in)

	prefixedIdKey := fdb.ProductSubspace.Bytes()
	prefixedIdKey = append(prefixedIdKey, byte(in.ProductIdentifier.Id))

	if !fdb.Put(prefixedIdKey, buffer.Bytes()) {
		log.Println("Could not put product into FDB")
	}
	return in, nil
}

func (s *ProductServer) ClearProduct(ctx context.Context, in *service.ClearProductMessage) (*service.ClearProductMessage, error) {
	prefixedIdKey := fdb.ProductSubspace.Bytes()
	prefixedIdKey = append(prefixedIdKey, byte(in.Id))
	if !fdb.Clear(prefixedIdKey) {
		log.Println("Could not clear product from FDB")
	}

	prefixedOriginKey := fdb.ProductSubspace.Bytes()
	prefixedOriginKey = append(prefixedOriginKey, []byte(in.Origin)...)
	prefixedOriginKey = append(prefixedOriginKey, byte(in.Id))
	if !fdb.Clear(prefixedOriginKey) {
		log.Println("Could not clear product from origin index")
	}

	prefixedCategoryKey := fdb.ProductSubspace.Bytes()
	prefixedCategoryKey = append(prefixedCategoryKey, fdb.EncodeCategories(in.Categories)...)
	prefixedCategoryKey = append(prefixedCategoryKey, byte(in.Id))
	if !fdb.Clear(prefixedCategoryKey) {
		log.Println("Could not clear product from origin index")
	}

	for _, tag := range in.Tags {
		prefixedTagKey := fdb.ProductSubspace.Bytes()
		prefixedTagKey = append(prefixedTagKey, []byte(tag)...)
		prefixedTagKey = append(prefixedTagKey, byte(in.Id))
		if !fdb.Clear(prefixedTagKey) {
			log.Println("Could not clear product from origin index")
		}
	}

	return &service.ClearProductMessage{Id: in.Id, Origin: in.Origin, Categories: in.Categories, Tags: in.Tags}, nil
}
