package services

import (
	"errors"
	"fmt"
	"math"
	"slices"
)

type PackSizeService struct{}

func NewPackSizeService() *PackSizeService {
	return &PackSizeService{}
}

func (p *PackSizeService) CalculatePackSizeByOrderAmount(orderItems int, packSizes []int) (map[int]int, error) {
	if orderItems <= 0 {
		return nil, errors.New("order items must be greater than zero")
	}

	slices.Sort(packSizes) // ensure ascending order for better pruning

	type state struct {
		total int
		packs map[int]int
	}

	queue := []state{{total: 0, packs: map[int]int{}}}
	visited := make(map[int]bool)

	bestResult := map[int]int{}
	bestOver := math.MaxInt
	minPacks := math.MaxInt

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, size := range packSizes {
			newTotal := curr.total + size
			if newTotal > orderItems+packSizes[len(packSizes)-1] {
				continue // too much overhead
			}

			if visited[newTotal] {
				continue
			}
			visited[newTotal] = true

			newPack := make(map[int]int)
			for k, v := range curr.packs {
				newPack[k] = v
			}
			newPack[size]++

			numPacks := 0
			for _, v := range newPack {
				numPacks += v
			}

			if newTotal >= orderItems {
				if newTotal < bestOver || (newTotal == bestOver && numPacks < minPacks) {
					bestResult = newPack
					bestOver = newTotal
					minPacks = numPacks
				}
				continue
			}

			queue = append(queue, state{total: newTotal, packs: newPack})
		}
	}

	if len(bestResult) == 0 {
		return nil, fmt.Errorf("cannot find a combination of pack sizes to sum to at least %d", orderItems)
	}

	return bestResult, nil
}
