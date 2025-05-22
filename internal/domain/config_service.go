package domain

type ConfigService interface {
	UpdatePackSizes(sizes []int) error
	GetPackSizes() []int
}
