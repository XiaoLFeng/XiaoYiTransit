package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"go/types"
)

// DeleteVehicleReq defines the request for deleting a vehicle
type DeleteVehicleReq struct {
	g.Meta      `path:"/vehicle/{vehicleUuid}" method:"Delete" tags:"车辆管理" sm:"删除车辆" dc:"用于删除车辆信息"`
	VehicleUuid string `json:"vehicle_uuid" v:"required#车辆UUID不能为空" dc:"车辆UUID"`
}

// DeleteVehicleRes defines the response for deleting a vehicle
type DeleteVehicleRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}
