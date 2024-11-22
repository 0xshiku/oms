package main

import (
	pb "common/api"
	"context"
)

type OrdersService interface {
	CreateOrder(context.Context, *pb.CreateOrderRequest, []*pb.Item) (*pb.Order, error)
	ValidateOrder(context.Context, *pb.CreateOrderRequest) ([]*pb.Item, error)
	GetOrder(context.Context, *pb.GetOrderRequest) (*pb.Order, error)
	UpdateOrder(context.Context, *pb.Order) (*pb.Order, error)
}

type OrdersStore interface {
	Create(context.Context, *pb.CreateOrderRequest, []*pb.Item) (string, error)
	Get(ctx context.Context, id, customerID string) (*pb.Order, error)
	Update(context.Context, string, *pb.Order) error
}
