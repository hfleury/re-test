package domain

type PackSizeService interface {
	CalculatePackSizeByOrderAmount(orderItems int, packSizes []int) (map[int]int, error)
}
