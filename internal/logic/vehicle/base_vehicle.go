package vehicle

import "xiao-yi-transit/internal/service"

type sVehicle struct {
}

func init() {
	service.RegisterVehicle(New())
}

func New() *sVehicle {
	return &sVehicle{}
}
