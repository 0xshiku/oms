package main

import (
	pb "common/api"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	dbName   = "orders"
	CollName = "orders"
)

var orders = make([]*pb.Order, 0)

type store struct {
	db *mongo.Client
}

func NewStore(db *mongo.Client) *store {
	return &store{db}
}

func (s *store) Create(ctx context.Context, p *pb.CreateOrderRequest, items []*pb.Item) (string, error) {
	col := s.db.Database(dbName).Collection(CollName)

	col.InsertOne(ctx)

	id := "42"
	orders = append(orders, &pb.Order{
		ID:          id,
		CustomerID:  p.CustomerID,
		Status:      "pending",
		Items:       items,
		PaymentLink: "",
	})

	return id, nil
}

func (s *store) Get(ctx context.Context, id, customerID string) (*pb.Order, error) {
	for _, o := range orders {
		if o.ID == id && o.CustomerID == customerID {
			return o, nil
		}
	}

	return nil, errors.New("order not found")
}

func (s *store) Update(ctx context.Context, id string, newOrder *pb.Order) error {
	for i, order := range orders {
		if order.ID == id {
			orders[i].Status = newOrder.Status
			orders[i].PaymentLink = newOrder.PaymentLink

			return nil
		}
	}

	return nil
}
