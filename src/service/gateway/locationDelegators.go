package grpcServer

import (
	"context"
	"log"

	"CS467_SU21/proto/service"
)

func (s *GatewayServer) GetLocation(ctx context.Context, in *service.LocationIdentifier) (*service.StoredLocation, error) {
	response, err := LocationClient.GetLocation(ctx, in)
	if err != nil {
		log.Println("Error location client")
	}
	return response, nil
}

func (s *GatewayServer) PutLocation(ctx context.Context, in *service.PutLocationMessage) (*service.StoredLocation, error) {
	response, err := LocationClient.PutLocation(ctx, in)
	if err != nil {
		log.Println("Error location client")
	}
	return response, nil

}
func (s *GatewayServer) SendProduct(ctx context.Context, in *service.MoveProductMessage) (*service.MoveProductMessage, error) {
	//product client here
	response, err := LocationClient.SendProduct(ctx, in)
	if err != nil {
		log.Println("Error location client")
	}

	product, err := ProductClient.GetProduct(ctx, &service.ProductIdentifier{Id: in.ProductId})
	if err != nil {
		log.Println("Could not get product")
	}

	product.QuantityByLocation[in.SendLocation.Name] -= in.Quantity
	product.QuantityByLocation[in.SendLocation.Name+"->"+in.ReceiveLocation.Name] = in.Quantity

	product.QuantityInTransit += in.Quantity

	ProductClient.UpdateProduct(ctx, product)

	return response, nil

}

func (s *GatewayServer) ReceiveProduct(ctx context.Context, in *service.MoveProductMessage) (*service.MoveProductMessage, error) {
	//product client here
	response, err := LocationClient.ReceiveProduct(ctx, in)
	if err != nil {
		log.Println("Error location client")
	}

	product, err := ProductClient.GetProduct(ctx, &service.ProductIdentifier{Id: in.ProductId})
	if err != nil {
		log.Println("Could not get product")
	}

	product.QuantityByLocation[in.ReceiveLocation.Name] += in.Quantity
	if in.SendLocation != nil {
		_, ok := product.QuantityByLocation[in.SendLocation.Name+"->"+in.ReceiveLocation.Name]

		if ok {
			delete(product.QuantityByLocation, in.SendLocation.Name+"->"+in.ReceiveLocation.Name)
		}
	}

	product.QuantityInTransit -= in.Quantity

	ProductClient.UpdateProduct(ctx, product)

	return response, nil

}
