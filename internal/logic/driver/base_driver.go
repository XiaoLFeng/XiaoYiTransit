package driver

import "xiao-yi-transit/internal/service"

type sDriver struct {
}

func init() {
	service.RegisterDriver(New())
}

func New() *sDriver {
	return &sDriver{}
}
