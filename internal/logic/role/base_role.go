package role

import "xiao-yi-transit/internal/service"

type sRole struct {
}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}
