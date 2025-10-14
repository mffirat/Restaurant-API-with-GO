package redis


type FloorRepoInterface interface {
	IncreaseFloorCount(floor int) error
	DecreaseFloorCount(floor int) error
	GetFloorCount(floor int) (int, error)
}