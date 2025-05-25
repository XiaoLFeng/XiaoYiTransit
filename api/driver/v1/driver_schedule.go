package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"xiao-yi-transit/internal/model/dto/driver"
)

// GetDriverScheduleReq defines the request for getting a driver's schedule
type GetDriverScheduleReq struct {
	g.Meta     `path:"/driver/schedule" method:"Get" tags:"司机管理" sm:"获取司机排班" dc:"用于获取司机排班信息"`
	DriverUuid string      `json:"driver_uuid" v:"required#司机UUID不能为空" dc:"司机UUID"`
	StartDate  *gtime.Time `json:"start_date,omitempty" dc:"开始日期（可选，默认为当前日期）"`
	EndDate    *gtime.Time `json:"end_date,omitempty" dc:"结束日期（可选，默认为开始日期后7天）"`
}

// GetDriverScheduleRes defines the response for getting a driver's schedule
type GetDriverScheduleRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*driver.DriverScheduleDTO]
}
