// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Maintenance is the golang structure of table xf_maintenance for DAO operations like Where/Data.
type Maintenance struct {
	g.Meta              `orm:"table:xf_maintenance, do:true"`
	MaintenanceUuid     interface{} // 维修UUID
	VehicleUuid         interface{} // 车辆UUID
	MaintenanceType     interface{} // 维修类型: 1-常规保养, 2-故障维修, 3-事故维修, 4-年检维修
	Description         interface{} // 维修描述
	MaintenanceDate     *gtime.Time // 维修日期
	CompletionDate      *gtime.Time // 完成日期
	Cost                interface{} // 维修费用
	Mileage             interface{} // 维修时里程数
	MaintenanceLocation interface{} // 维修地点
	MaintenanceStaff    interface{} // 维修人员
	PartsReplaced       interface{} // 更换的零部件
	Status              interface{} // 状态: 0-已取消, 1-待维修, 2-维修中, 3-已完成
	Notes               interface{} // 备注
	CreatedByUuid       interface{} // 创建人UUID
	CreatedAt           *gtime.Time // 创建时间
	UpdatedAt           *gtime.Time // 更新时间
	DeletedAt           *gtime.Time // 删除时间
}
