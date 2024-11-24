package main

import (
	pb "common/api"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"
)

type TelemetryMiddleware struct {
	next OrdersService
}

func NewTelemetryMiddleware(next OrdersService) OrdersService {
	return &TelemetryMiddleware{next}
}

func (t *TelemetryMiddleware) GetOrder(ctx context.Context, p *pb.GetOrderRequest) (*pb.Order, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(fmt.Sprintf("GetOrder: %v", p))

	return t.next.GetOrder(ctx, p)
}

func (t *TelemetryMiddleware) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest, items []*pb.Item) (*pb.Order, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(fmt.Sprintf("CreateOrder: %v", p))

	return t.next.CreateOrder(ctx, p, items)
}

func (t *TelemetryMiddleware) ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) ([]*pb.Item, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(fmt.Sprintf("ValidateOrder: %v", p))

	return t.next.ValidateOrder(ctx, p)
}

func (t *TelemetryMiddleware) UpdateOrder(ctx context.Context, o *pb.Order) (*pb.Order, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(fmt.Sprintf("UpdateOrder: %v", o))

	return t.next.UpdateOrder(ctx, o)
}
