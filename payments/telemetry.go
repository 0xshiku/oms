package main

import (
	pb "common/api"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"
)

type TelemetryMiddleware struct {
	next PaymentsService
}

func NewTelemetryMiddleware(next PaymentsService) PaymentsService {
	return &TelemetryMiddleware{next}
}

func (t *TelemetryMiddleware) CreatePayment(ctx context.Context, o *pb.Order) (string, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(fmt.Sprintf("GetOrder: %v", o))

	return t.next.CreatePayment(ctx, o)
}
