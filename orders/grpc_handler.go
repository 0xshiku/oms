package main

import (
	pb "common/api"
	"context"
	"google.golang.org/grpc"
	"log"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service OrdersService
}

func newGRPCHandler(grpcServer *grpc.Server, service OrdersService) {
	handler := &grpcHandler{
		service: service,
	}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, orderRequest *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New order received! Order %v", orderRequest)
	o := &pb.Order{
		ID: "42",
	}
	return o, nil
}
