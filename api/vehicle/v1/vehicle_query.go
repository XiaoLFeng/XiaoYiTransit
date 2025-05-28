package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"xiao-yi-transit/internal/model/dto/vehicle"
)

// GetVehicleListReq defines the request for getting a list of vehicles
type GetVehicleListReq struct {
	g.Meta      `path:"/vehicles" method:"Get" tags:"车辆管理" sm:"获取车辆列表" dc:"用于获取车辆列表信息"`
	Page        int    `json:"page" d:"1" v:"min:1#页码最小为1" dc:"页码"`
	Size        int    `json:"size" d:"10" v:"max:100#每页最大100条" dc:"每页数量"`
	PlateNumber string `json:"plate_number,omitempty" dc:"车牌号（可选，用于筛选）"`
	Model       string `json:"model,omitempty" dc:"车辆型号（可选，用于筛选）"`
	Status      int    `json:"status,omitempty" dc:"状态（可选，用于筛选）: 1-运营, 2-维修, 3-停运, 4-报废"`
}

// GetVehicleListRes defines the response for getting a list of vehicles
type GetVehicleListRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*vehicle.PagedVehicleListDTO]
}

type GetVehicleSimpleListReq struct {
	g.Meta `path:"/vehicles/simple" method:"Get" tags:"车辆管理" sm:"获取车辆简易列表" dc:"用于获取车辆简易列表信息"`
}

type GetVehicleSimpleListRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*vehicle.SimpleVehicleListDTO]
}

// GetVehicleDetailReq defines the request for getting vehicle details
type GetVehicleDetailReq struct {
	g.Meta      `path:"/vehicle/{vehicleUuid}" method:"Get" tags:"车辆管理" sm:"获取车辆详情" dc:"用于获取车辆详细信息"`
	VehicleUuid string `json:"vehicle_uuid" v:"required#车辆UUID不能为空" dc:"车辆UUID"`
}

// GetVehicleDetailRes defines the response for getting vehicle details
type GetVehicleDetailRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*vehicle.VehicleDetailItemDTO]
}
