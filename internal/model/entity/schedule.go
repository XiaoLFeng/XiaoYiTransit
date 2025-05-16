// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Schedule is the golang structure for table schedule.
type Schedule struct {
	Id              int         `json:"id"                orm:"id"                description:"调度ID"`                          // 调度ID
	RouteId         int         `json:"route_id"          orm:"route_id"          description:"线路ID"`                          // 线路ID
	VehicleId       int         `json:"vehicle_id"        orm:"vehicle_id"        description:"车辆ID"`                          // 车辆ID
	DriverId        int         `json:"driver_id"         orm:"driver_id"         description:"司机ID"`                          // 司机ID
	ScheduleDate    *gtime.Time `json:"schedule_date"     orm:"schedule_date"     description:"调度日期"`                          // 调度日期
	StartTime       *gtime.Time `json:"start_time"        orm:"start_time"        description:"开始时间"`                          // 开始时间
	EndTime         *gtime.Time `json:"end_time"          orm:"end_time"          description:"结束时间"`                          // 结束时间
	Status          int         `json:"status"            orm:"status"            description:"状态: 0-取消, 1-待执行, 2-执行中, 3-已完成"` // 状态: 0-取消, 1-待执行, 2-执行中, 3-已完成
	ActualStartTime *gtime.Time `json:"actual_start_time" orm:"actual_start_time" description:"实际开始时间"`                        // 实际开始时间
	ActualEndTime   *gtime.Time `json:"actual_end_time"   orm:"actual_end_time"   description:"实际结束时间"`                        // 实际结束时间
	Mileage         float64     `json:"mileage"           orm:"mileage"           description:"行驶里程(km)"`                      // 行驶里程(km)
	FuelConsumption float64     `json:"fuel_consumption"  orm:"fuel_consumption"  description:"油耗(L)"`                         // 油耗(L)
	PassengerCount  int         `json:"passenger_count"   orm:"passenger_count"   description:"载客人数"`                          // 载客人数
	Notes           string      `json:"notes"             orm:"notes"             description:"备注"`                            // 备注
	CreatedBy       int         `json:"created_by"        orm:"created_by"        description:"创建人ID"`                         // 创建人ID
	CreatedAt       *gtime.Time `json:"created_at"        orm:"created_at"        description:"创建时间"`                          // 创建时间
	UpdatedAt       *gtime.Time `json:"updated_at"        orm:"updated_at"        description:"更新时间"`                          // 更新时间
	DeletedAt       *gtime.Time `json:"deleted_at"        orm:"deleted_at"        description:"删除时间"`                          // 删除时间
}
