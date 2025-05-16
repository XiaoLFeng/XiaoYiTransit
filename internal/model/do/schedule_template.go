// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ScheduleTemplate is the golang structure of table xf_schedule_template for DAO operations like Where/Data.
type ScheduleTemplate struct {
	g.Meta       `orm:"table:xf_schedule_template, do:true"`
	Id           interface{} // 模板ID
	Name         interface{} // 模板名称
	RouteId      interface{} // 线路ID
	DayOfWeek    interface{} // 星期几: 1-7 代表周一至周日
	StartTime    *gtime.Time // 开始时间
	EndTime      *gtime.Time // 结束时间
	VehicleCount interface{} // 所需车辆数
	Status       interface{} // 状态: 0-禁用, 1-启用
	Notes        interface{} // 备注
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
