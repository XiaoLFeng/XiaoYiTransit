// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DriverShift is the golang structure of table xf_driver_shift for DAO operations like Where/Data.
type DriverShift struct {
	g.Meta     `orm:"table:xf_driver_shift, do:true"`
	ShiftUuid  interface{} // 排班UUID
	DriverUuid interface{} // 司机UUID
	ShiftDate  *gtime.Time // 排班日期
	ShiftType  interface{} // 班次类型: 1-早班, 2-中班, 3-晚班, 4-全天班
	StartTime  *gtime.Time // 开始时间
	EndTime    *gtime.Time // 结束时间
	Status     interface{} // 状态: 0-取消, 1-待执行, 2-执行中, 3-已完成
	Notes      interface{} // 备注
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
