package auth

import "xiao-yi-transit/internal/service"

type sAuth struct {
}

func init() {
	service.RegisterAuth(New())
}

func New() *sAuth {
	return &sAuth{}
}
