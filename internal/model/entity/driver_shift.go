// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DriverShift is the golang structure for table driver_shift.
type DriverShift struct {
	ShiftUuid  string      `json:"shift_uuid"  orm:"shift_uuid"  description:"排班UUID"`                        // 排班UUID
	DriverUuid string      `json:"driver_uuid" orm:"driver_uuid" description:"司机UUID"`                        // 司机UUID
	ShiftDate  *gtime.Time `json:"shift_date"  orm:"shift_date"  description:"排班日期"`                          // 排班日期
	ShiftType  int         `json:"shift_type"  orm:"shift_type"  description:"班次类型: 1-早班, 2-中班, 3-晚班, 4-全天班"` // 班次类型: 1-早班, 2-中班, 3-晚班, 4-全天班
	StartTime  *gtime.Time `json:"start_time"  orm:"start_time"  description:"开始时间"`                          // 开始时间
	EndTime    *gtime.Time `json:"end_time"    orm:"end_time"    description:"结束时间"`                          // 结束时间
	Status     int         `json:"status"      orm:"status"      description:"状态: 0-取消, 1-待执行, 2-执行中, 3-已完成"` // 状态: 0-取消, 1-待执行, 2-执行中, 3-已完成
	Notes      string      `json:"notes"       orm:"notes"       description:"备注"`                            // 备注
	CreatedAt  *gtime.Time `json:"created_at"  orm:"created_at"  description:"创建时间"`                          // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at"  orm:"updated_at"  description:"更新时间"`                          // 更新时间
}
