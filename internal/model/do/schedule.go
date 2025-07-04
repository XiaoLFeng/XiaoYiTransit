// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Schedule is the golang structure of table xf_schedule for DAO operations like Where/Data.
type Schedule struct {
	g.Meta          `orm:"table:xf_schedule, do:true"`
	ScheduleUuid    interface{} // 调度UUID
	RouteUuid       interface{} // 线路UUID
	VehicleUuid     interface{} // 车辆UUID
	DriverUuid      interface{} // 司机UUID
	ScheduleDate    *gtime.Time // 调度日期
	StartTime       *gtime.Time // 开始时间
	EndTime         *gtime.Time // 结束时间
	Status          interface{} // 状态: 0-取消, 1-待执行, 2-执行中, 3-已完成
	ActualStartTime *gtime.Time // 实际开始时间
	ActualEndTime   *gtime.Time // 实际结束时间
	Mileage         interface{} // 行驶里程(km)
	FuelConsumption interface{} // 油耗(L)
	PassengerCount  interface{} // 载客人数
	Notes           interface{} // 备注
	CreatedByUuid   interface{} // 创建人UUID
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
	DeletedAt       *gtime.Time // 删除时间
}
