package station

import "xiao-yi-transit/internal/service"

type sStation struct {
}

func init() {
	service.RegisterStation(New())
}

func New() *sStation {
	return &sStation{}
}
