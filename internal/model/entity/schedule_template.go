// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ScheduleTemplate is the golang structure for table schedule_template.
type ScheduleTemplate struct {
	ScheduleTemplateUuid string      `json:"schedule_template_uuid" orm:"schedule_template_uuid" description:"模板UUID"`           // 模板UUID
	Name                 string      `json:"name"                   orm:"name"                   description:"模板名称"`             // 模板名称
	RouteUuid            string      `json:"route_uuid"             orm:"route_uuid"             description:"线路UUID"`           // 线路UUID
	DayOfWeek            int         `json:"day_of_week"            orm:"day_of_week"            description:"星期几: 1-7 代表周一至周日"` // 星期几: 1-7 代表周一至周日
	StartTime            *gtime.Time `json:"start_time"             orm:"start_time"             description:"开始时间"`             // 开始时间
	EndTime              *gtime.Time `json:"end_time"               orm:"end_time"               description:"结束时间"`             // 结束时间
	VehicleCount         int         `json:"vehicle_count"          orm:"vehicle_count"          description:"所需车辆数"`            // 所需车辆数
	Status               int         `json:"status"                 orm:"status"                 description:"状态: 0-禁用, 1-启用"`   // 状态: 0-禁用, 1-启用
	Notes                string      `json:"notes"                  orm:"notes"                  description:"备注"`               // 备注
	CreatedAt            *gtime.Time `json:"created_at"             orm:"created_at"             description:"创建时间"`             // 创建时间
	UpdatedAt            *gtime.Time `json:"updated_at"             orm:"updated_at"             description:"更新时间"`             // 更新时间
}
