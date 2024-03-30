package profiles

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServiceClient() (ProfileServiceClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial("profiles:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return NewProfileServiceClient(cc), nil
}
