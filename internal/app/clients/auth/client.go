package auth

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServiceClient() (AuthServiceClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial("auth:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return NewAuthServiceClient(cc), nil
}
