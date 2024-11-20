package main

import (
	pb "common/api"
	"context"
)

type PaymentsService interface {
	CreatePayment(context.Context, *pb.Order) (string, error)
}
