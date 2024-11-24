package main

import (
	"common"
	"common/broker"
	"common/discovery"
	"common/discovery/consul"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var (
	serviceName = "orders"
	grpcAddr    = common.EnvString("GRPC_ADDR", "localhost:2000")
	consulAddr  = common.EnvString("CONSUL_ADDR", "localhost:8500")
	amqpUser    = common.EnvString("RABBITMQ_USER", "guest")
	amqpPass    = common.EnvString("RABBITMQ_PASS", "guest")
	amqpHost    = common.EnvString("RABBITMQ_HOST", "localhost")
	amqpPort    = common.EnvString("RABBITMQ_PORT", "5672")
	jaegerAddr  = common.EnvString("JAEGER_PORT", "localhost:4318")
)

func main() {
	if err := common.SetGlobalTracer(context.TODO(), serviceName, jaegerAddr); err != nil {
		log.Fatal("Failed to set global tracer")
	}

	registry, err := consul.NewRegistry(consulAddr, serviceName)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, grpcAddr); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				log.Fatal("Failed to health check")
			}
			time.Sleep(time.Second * 1)
		}
	}()

	defer registry.DeRegister(ctx, instanceID, serviceName)

	ch, close := broker.Connect(amqpUser, amqpPass, amqpHost, amqpPort)
	defer func() {
		close()
		ch.Close()
	}()

	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)
	// Decorator Pattern
	// Note: NewTelemetryMiddleware needs to implement the same interface as svc
	// New Telemetry middleware just adds the telemetry functionality on the service
	// Notice that each method on svcWithTelemetry just returns the svc method with .next
	// Before the return we add a new functionality
	svcWithTelemetry := NewTelemetryMiddleware(svc)

	newGRPCHandler(grpcServer, svcWithTelemetry, ch)

	consumer := NewConsumer(svcWithTelemetry)
	go consumer.Listen(ch)

	log.Println("GRPC Server started at ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
