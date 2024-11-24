package discovery

import (
	"context"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
)

func ServiceConnection(ctx context.Context, serviceName string, registry Registry) (*grpc.ClientConn, error) {
	addrs, err := registry.Discovery(ctx, serviceName)
	if err != nil {
		return nil, err
	}

	log.Printf("Discovered %d instances of %s", len(addrs), serviceName)

	return grpc.NewClient(
		addrs[rand.Intn(len(addrs))],
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// Middlewares
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
}
