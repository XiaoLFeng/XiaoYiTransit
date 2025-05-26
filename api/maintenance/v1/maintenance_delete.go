package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"go/types"
)

// DeleteMaintenanceReq defines the request for deleting a maintenance record
type DeleteMaintenanceReq struct {
	g.Meta          `path:"/maintenance/{maintenance_uuid}" method:"Delete" tags:"维修管理" sm:"删除维修记录" dc:"用于删除车辆维修记录"`
	MaintenanceUuid string `json:"maintenance_uuid" v:"required#维修记录UUID不能为空" dc:"维修记录UUID" in:"path"`
}

// DeleteMaintenanceRes defines the response for deleting a maintenance record
type DeleteMaintenanceRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}