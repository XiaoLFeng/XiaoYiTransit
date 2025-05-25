package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"xiao-yi-transit/internal/model/dto/driver"
)

// GetDriverListReq defines the request for getting a list of drivers
type GetDriverListReq struct {
	g.Meta     `path:"/driver/list" method:"Get" tags:"司机管理" sm:"获取司机列表" dc:"用于获取司机列表信息"`
	Page       int    `json:"page" d:"1" v:"min:1#页码最小为1" dc:"页码"`
	Size       int    `json:"size" d:"10" v:"max:100#每页最大100条" dc:"每页数量"`
	EmployeeId string `json:"employee_id,omitempty" dc:"工号（可选，用于筛选）"`
	Name       string `json:"name,omitempty" dc:"姓名（可选，用于筛选）"`
	Status     int    `json:"status,omitempty" dc:"状态（可选，用于筛选）: 0-离职, 1-在职, 2-休假, 3-停职"`
}

// GetDriverListRes defines the response for getting a list of drivers
type GetDriverListRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*driver.PagedDriverListDTO]
}

// GetDriverDetailReq defines the request for getting driver details
type GetDriverDetailReq struct {
	g.Meta     `path:"/driver/detail" method:"Get" tags:"司机管理" sm:"获取司机详情" dc:"用于获取司机详细信息"`
	DriverUuid string `json:"driver_uuid" v:"required#司机UUID不能为空" dc:"司机UUID"`
}

// GetDriverDetailRes defines the response for getting driver details
type GetDriverDetailRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*driver.DriverDetailItemDTO]
}
