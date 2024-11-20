package main

import (
	pb "common/api"
	"context"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) CreatePayment(ctx context.Context, p *pb.Order) (string, error) {
	// Connect to payment processor

	return "", nil
}
