// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VehicleFault is the golang structure of table xf_vehicle_fault for DAO operations like Where/Data.
type VehicleFault struct {
	g.Meta           `orm:"table:xf_vehicle_fault, do:true"`
	Id               interface{} // 故障ID
	VehicleId        interface{} // 车辆ID
	ReporterId       interface{} // 报告人ID
	FaultType        interface{} // 故障类型
	FaultDescription interface{} // 故障描述
	ReportDate       *gtime.Time // 报告时间
	FaultLocation    interface{} // 故障发生地点
	Severity         interface{} // 严重程度: 1-轻微, 2-一般, 3-严重, 4-危险
	Status           interface{} // 状态: 0-已取消, 1-待处理, 2-处理中, 3-已解决
	MaintenanceId    interface{} // 关联的维修记录ID
	Notes            interface{} // 备注
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}
