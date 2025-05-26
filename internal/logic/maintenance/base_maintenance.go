package maintenance

import "xiao-yi-transit/internal/service"

type sMaintenance struct {
}

func init() {
	service.RegisterMaintenance(New())
}

func New() *sMaintenance {
	return &sMaintenance{}
}
