package profiles

import (
	"crypto/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/config"
)

func NewServiceClient(cfg config.Config) (c ProfileServiceClient, err error) {
	var cc *grpc.ClientConn
	if cfg.UseSSL {
		cc, err = grpc.Dial(cfg.Address, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		cc, err = grpc.Dial(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	if err != nil {
		return nil, err
	}
	return NewProfileServiceClient(cc), nil
}
