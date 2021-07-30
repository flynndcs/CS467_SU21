package grpcServer

import (
	"CS467_SU21/proto/service"
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"log"

	"CS467_SU21/src/store/fdb"
)

func (s *LocationServer) GetLocation(ctx context.Context, in *service.LocationIdentifier) (*service.StoredLocation, error) {
	prefixedNameKey := fdb.LocationSubspace.Bytes()
	prefixedNameKey = append(prefixedNameKey, in.Name...)
	value := fdb.Get(prefixedNameKey)
	if value == nil {
		return &service.StoredLocation{}, nil
	}
	buffer := bytes.NewBuffer(value)
	dec := gob.NewDecoder(buffer)

	var response service.StoredLocation
	dec.Decode(&response)
	return &response, nil
}

func (s *LocationServer) PutLocation(ctx context.Context, in *service.PutLocationMessage) (*service.StoredLocation, error) {
	storedLocation := service.StoredLocation{
		Name:              in.Name,
		Receives:          in.Receives,
		Sends:             in.Sends,
		QuantityByProduct: map[int64]int64{},
		PreviousLocations: in.PreviousLocations,
		NextLocations:     in.NextLocations,
	}

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(&storedLocation)

	prefixedNameKey := fdb.LocationSubspace.Bytes()
	prefixedNameKey = append(prefixedNameKey, in.Name...)

	if !fdb.Put(prefixedNameKey, buffer.Bytes()) {
		log.Println("Could not put location into fdb")
	}

	return &storedLocation, nil
}

func (s *LocationServer) UpdateLocation(ctx context.Context, in *service.StoredLocation) (*service.StoredLocation, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(&in)

	prefixedNameKey := fdb.LocationSubspace.Bytes()
	prefixedNameKey = append(prefixedNameKey, in.Name...)

	if !fdb.Put(prefixedNameKey, buffer.Bytes()) {
		log.Println("Could not put location into fdb")
	}
	return in, nil
}
func (s *LocationServer) SendProduct(ctx context.Context, in *service.MoveProductMessage) (*service.MoveProductMessage, error) {
	location, err := s.GetLocation(ctx, &service.LocationIdentifier{Name: in.SendLocation.Name})
	if err != nil {
		log.Println("Could not get location")
	}
	valid := false
	for _, v := range location.Sends {
		if v == in.ProductName {
			for _, x := range location.Receives {
				if x == in.ProductName {
					for _, y := range location.NextLocations {
						if y == in.ReceiveLocation {
							valid = true
							break
						}
					}
					break
				}
			}
			break
		}
	}
	if !valid {
		return nil, errors.New("product cannot be sent from or received at these locations")
	}
	if location.QuantityByProduct[in.ProductId] > in.Quantity {
		location.QuantityByProduct[in.ProductId] -= in.Quantity
	} else {
		log.Println("Not enough products to send.")
	}
	s.UpdateLocation(ctx, location)

	return in, nil
}

func (s *LocationServer) ReceiveProduct(ctx context.Context, in *service.MoveProductMessage) (*service.MoveProductMessage, error) {
	location, err := s.GetLocation(ctx, &service.LocationIdentifier{Name: in.ReceiveLocation.Name})
	if err != nil {
		log.Println("Could not get location")
	}
	location.GetQuantityByProduct()[in.ProductId] += in.Quantity
	s.UpdateLocation(ctx, location)
	return in, nil
}
