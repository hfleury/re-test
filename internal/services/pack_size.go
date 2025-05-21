package services

import (
	"fmt"
	"slices"
)

type PackSizeService struct {
	PackSizes []int
}

func NewPackSizeService(packSizes []int) *PackSizeService {
	slices.SortFunc(packSizes, func(a, b int) int {
		return b - a
	})
	return &PackSizeService{
		PackSizes: packSizes,
	}
}

func (p *PackSizeService) CalculatePackSizeByOrderAmount(orderItems int) (map[int]int, error) {
	if orderItems <= 0 {
		return nil, fmt.Errorf("order items must be greater than zero")
	}

	packSizeMap := make(map[int]int)
	remainingItems := orderItems

	for _, packSize := range p.PackSizes {
		count := remainingItems / packSize
		if count > 0 {
			packSizeMap[packSize] = count
			remainingItems -= count * packSize
		}
	}

	if remainingItems > 0 {
		smallPack := p.PackSizes[len(p.PackSizes)-1]
		packSizeMap[smallPack] += 1
	}

	return packSizeMap, nil
}
