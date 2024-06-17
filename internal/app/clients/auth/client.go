package auth

import (
	"fmt"

	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/common"
	"google.golang.org/grpc"
)

func NewServiceClient(cfg common.Config) (c AuthServiceClient, err error) {
	var cc *grpc.ClientConn
	cc, err = common.Dial(cfg)
	if err != nil {
		return nil, fmt.Errorf("dialing auth service: %w", err)
	}
	return NewAuthServiceClient(cc), nil
}
