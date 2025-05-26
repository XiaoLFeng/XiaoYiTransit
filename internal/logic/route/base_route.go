package route

import "xiao-yi-transit/internal/service"

type sRoute struct {
}

func init() {
	service.RegisterRoute(New())
}

func New() *sRoute {
	return &sRoute{}
}
