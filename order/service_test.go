package order

import (
	"github.com/stretchr/testify/suite"
    "testing"
)

type OrderServiceSuite struct {
	suite.Suite

	service *Service
}

func (s *OrderServiceSuite) SetupSuite() {
}

func (s *OrderServiceSuite) SetupTest() {

}

// Test
func (s *OrderServiceSuite) Test() {

}

func TestOrderServiceSuite(t *testing.T) {
	suite.Run(t, new(OrderServiceSuite))
}

func (s *OrderServiceSuite) TearDownSuite() {

}
