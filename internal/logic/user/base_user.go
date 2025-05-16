package user

import "xiao-yi-transit/internal/service"

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}
