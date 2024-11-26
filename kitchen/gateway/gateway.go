package gateway

import (
	pb "common/api"
	"context"
)

type KitchenGateway interface {
	UpdateOrder(context.Context, *pb.Order) error
}
