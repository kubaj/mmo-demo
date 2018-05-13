package order

import (
	"errors"
	"log"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"github.com/kubaj/mmo-demo/auth"
	"github.com/kubaj/mmo-demo/core"
	"golang.org/x/net/context"
)

// Order represents an implementation of the service interface
type Service struct {
	AuthClient auth.AuthServiceClient
}

// NewOrder creates a new service object
func NewService(authClient auth.AuthServiceClient) *Service {
	return &Service{
		AuthClient: authClient,
	}
}

func (s *Service) GetVersion(ctx context.Context, in *google_protobuf.Empty) (*Version, error) {
	return &Version{
		Name: "1.0.0",
	}, nil
}

func (s *Service) Check(ctx context.Context, in *core.HealthCheckRequest) (*core.HealthCheckResponse, error) {
	return &core.HealthCheckResponse{
		Status: core.HealthCheckResponse_SERVING,
	}, nil
}

func (s *Service) PlaceOrder(ctx context.Context, in *OrderRequest) (*OrderResponse, error) {
	resp, err := s.AuthClient.IsUserLogged(context.Background(), &auth.Token{Token: in.Token})
	if err != nil {
		return &OrderResponse{}, err
	}

	if !resp.Logged {
		return &OrderResponse{}, errors.New("User not authenticated")
	}

	log.Println("Placing order of these products:", in.Products)
	return &OrderResponse{
		Successful: true,
	}, nil
}
