package auth

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/config"
)

func NewServiceClient(cfg config.Config) (c AuthServiceClient, err error) {
	var cc *grpc.ClientConn
	if cfg.UseSSL {
		cc, err = grpc.Dial(cfg.Address, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		cc, err = grpc.Dial(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	if err != nil {
		return nil, fmt.Errorf("dialing media service: %w", err)
	}
	return NewAuthServiceClient(cc), nil
}
