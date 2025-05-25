package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GetDriverScheduleReq defines the request for getting a driver's schedule
type GetDriverScheduleReq struct {
	g.Meta     `path:"/driver/schedule" method:"Get" tags:"司机管理" sm:"获取司机排班" dc:"用于获取司机排班信息"`
	DriverUuid string      `json:"driver_uuid" v:"required#司机UUID不能为空" dc:"司机UUID"`
	StartDate  *gtime.Time `json:"start_date,omitempty" dc:"开始日期（可选，默认为当前日期）"`
	EndDate    *gtime.Time `json:"end_date,omitempty" dc:"结束日期（可选，默认为开始日期后7天）"`
}

// DriverShiftItem defines a single shift item in the schedule
type DriverShiftItem struct {
	ShiftUuid  string      `json:"shift_uuid" dc:"排班UUID"`
	ShiftDate  *gtime.Time `json:"shift_date" dc:"排班日期"`
	ShiftType  int         `json:"shift_type" dc:"班次类型: 1-早班, 2-中班, 3-晚班, 4-全天班"`
	StartTime  *gtime.Time `json:"start_time" dc:"开始时间"`
	EndTime    *gtime.Time `json:"end_time" dc:"结束时间"`
	Status     int         `json:"status" dc:"状态: 0-取消, 1-待执行, 2-执行中, 3-已完成"`
	Notes      string      `json:"notes" dc:"备注"`
}

// DriverSchedule defines the driver's schedule
type DriverSchedule struct {
	DriverUuid string           `json:"driver_uuid" dc:"司机UUID"`
	DriverName string           `json:"driver_name" dc:"司机姓名"`
	StartDate  *gtime.Time      `json:"start_date" dc:"开始日期"`
	EndDate    *gtime.Time      `json:"end_date" dc:"结束日期"`
	Shifts     []*DriverShiftItem `json:"shifts" dc:"排班列表"`
}

// GetDriverScheduleRes defines the response for getting a driver's schedule
type GetDriverScheduleRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*DriverSchedule]
}