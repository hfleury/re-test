package domain

type PackSizeService interface {
	CalculatePackSizeByOrderAmount(orderItems int) (map[int]int, error)
}
