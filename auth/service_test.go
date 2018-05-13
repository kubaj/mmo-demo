package auth

import (
	"github.com/stretchr/testify/suite"
    "testing"
)

type AuthServiceSuite struct {
	suite.Suite

	service *Service
}

func (s *AuthServiceSuite) SetupSuite() {
}

func (s *AuthServiceSuite) SetupTest() {

}

// Test
func (s *AuthServiceSuite) Test() {

}

func TestAuthServiceSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceSuite))
}

func (s *AuthServiceSuite) TearDownSuite() {

}
