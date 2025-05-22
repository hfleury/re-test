package services

import (
	"fmt"
	"slices"
)

type PackSizeService struct{}

func NewPackSizeService() *PackSizeService {
	return &PackSizeService{}
}

func (p *PackSizeService) CalculatePackSizeByOrderAmount(orderItems int, packSizes []int) (map[int]int, error) {
	if orderItems <= 0 {
		return nil, fmt.Errorf("order items must be greater than zero")
	}

	slices.SortFunc(packSizes, func(a, b int) int {
		return b - a
	})

	packSizeMap := make(map[int]int)
	remainingItems := orderItems

	for _, packSize := range packSizes {
		count := remainingItems / packSize
		if count > 0 {
			packSizeMap[packSize] = count
			remainingItems -= count * packSize
		}
	}

	if remainingItems > 0 {
		smallPack := packSizes[len(packSizes)-1]
		packSizeMap[smallPack] += 1
	}

	return packSizeMap, nil
}
