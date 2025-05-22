package handlers

import (
	"github.com/stretchr/testify/mock"
)

type MockConfigService struct {
	mock.Mock
}

func (m *MockConfigService) UpdatePackSizes(sizes []int) error {
	return m.Called(sizes).Error(0)
}

func (m *MockConfigService) GetPackSizes() []int {
	return m.Called().Get(0).([]int)
}

type MockPackSizeService struct {
	mock.Mock
}

func (m *MockPackSizeService) CalculatePackSizeByOrderAmount(orderItems int, packSizes []int) (map[int]int, error) {
	args := m.Called(orderItems, packSizes)
	return args.Get(0).(map[int]int), args.Error(1)
}
