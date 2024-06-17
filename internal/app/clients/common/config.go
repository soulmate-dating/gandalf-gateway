package common

import (
	"crypto/tls"
	grpcProm "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Address string
	UseSSL  bool
}

func Dial(cfg Config) (*grpc.ClientConn, error) {
	dialOptions := []grpc.DialOption{
		grpc.WithUnaryInterceptor(grpcProm.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpcProm.StreamClientInterceptor),
	}
	if cfg.UseSSL {
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	return grpc.Dial(cfg.Address, dialOptions...)
}
