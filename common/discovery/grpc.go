package discovery

import (
	"context"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
)

func ServiceConnection(ctx context.Context, serviceName string, registry Registry) (*grpc.ClientConn, error) {
	addrs, err := registry.Discovery(ctx, serviceName)
	if err != nil {
		return nil, err
	}

	return grpc.NewClient(
		addrs[rand.Intn(len(addrs))],
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// Middlewares
		grpc.WithUnaryInterceptor(otelgrpc.NewClientHandler()),
		grpc.WithStreamInterceptor(otelgrpc.NewClientHandler()),
	)
}
