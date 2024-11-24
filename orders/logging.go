package main

import (
	pb "common/api"
	"context"
	"go.uber.org/zap"
	"time"
)

type LoggingMiddleware struct {
	next OrdersService
}

func NewLoggingMiddleware(next OrdersService) OrdersService {
	return &LoggingMiddleware{next}
}

func (t *LoggingMiddleware) GetOrder(ctx context.Context, p *pb.GetOrderRequest) (*pb.Order, error) {
	start := time.Now()
	defer func() {
		zap.L().Info("UpdateOrder", zap.Duration("took", time.Since(start)))
	}()

	return t.next.GetOrder(ctx, p)
}

func (t *LoggingMiddleware) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest, items []*pb.Item) (*pb.Order, error) {
	start := time.Now()
	defer func() {
		zap.L().Info("UpdateOrder", zap.Duration("took", time.Since(start)))
	}()

	return t.next.CreateOrder(ctx, p, items)
}

func (t *LoggingMiddleware) ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) ([]*pb.Item, error) {
	start := time.Now()
	defer func() {
		zap.L().Info("UpdateOrder", zap.Duration("took", time.Since(start)))
	}()

	return t.next.ValidateOrder(ctx, p)
}

func (t *LoggingMiddleware) UpdateOrder(ctx context.Context, o *pb.Order) (*pb.Order, error) {
	start := time.Now()
	defer func() {
		zap.L().Info("UpdateOrder", zap.Duration("took", time.Since(start)))
	}()

	return t.next.UpdateOrder(ctx, o)
}
