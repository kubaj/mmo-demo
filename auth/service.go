package auth

import (
	"log"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"github.com/kubaj/mmo-demo/core"
	"golang.org/x/net/context"
)

// Auth represents an implementation of the service interface
type Service struct {
}

// NewAuth creates a new service object
func NewService() *Service {
	return &Service{}
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

func (s *Service) IsUserLogged(ctx context.Context, in *Token) (*UserStatusResponse, error) {
	log.Println("Authenticating user with token " + in.Token)
	return &UserStatusResponse{
		Logged: true,
	}, nil
}
