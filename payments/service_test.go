package main

import (
	"common/api"
	inmemRegistry "common/discovery/inmem"
	"context"
	"payments/gateway"
	"payments/processor/inmem"
	"testing"
)

func TestService(t *testing.T) {
	processor := inmem.NewInmem()
	registry := inmemRegistry.NewRegistry()
	gateway := gateway.NewGRPCGateway(registry)
	svc := NewService(processor, gateway)

	t.Run("should create a payment link", func(t *testing.T) {
		link, err := svc.CreatePayment(context.Background(), &api.Order{})
		if err != nil {
			t.Errorf("CreatePayment() error = %v, want nil", err)
		}

		if link == "" {
			t.Errorf("CreatePayment() link is empty")
		}
	})
}
