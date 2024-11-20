package main

import (
	pb "common/api"
	"context"
	"payments/processor"
)

type service struct {
	processor processor.PaymentProcessor
}

func NewService(p processor.PaymentProcessor) *service {
	return &service{p}
}

func (s *service) CreatePayment(ctx context.Context, o *pb.Order) (string, error) {
	link, err := s.processor.CreatePaymentLink(o)
	if err != nil {
		return "", err
	}

	// Update order with the link

	return link, nil
}
