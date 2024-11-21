package gateway

import (
	pb "common/api"
	"context"
)

type OrdersGateway interface {
	CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error)
	GetOrder(context.Context, string, string) (*pb.Order, error)
}
