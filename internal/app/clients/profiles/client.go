package profiles

import (
	"fmt"

	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/common"
	"google.golang.org/grpc"
)

func NewServiceClient(cfg common.Config) (c ProfileServiceClient, err error) {
	var cc *grpc.ClientConn
	cc, err = common.Dial(cfg)
	if err != nil {
		return nil, fmt.Errorf("dialing profiles service: %w", err)
	}
	return NewProfileServiceClient(cc), nil
}
