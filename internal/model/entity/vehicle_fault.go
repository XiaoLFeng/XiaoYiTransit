// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VehicleFault is the golang structure for table vehicle_fault.
type VehicleFault struct {
	Id               int         `json:"id"                orm:"id"                description:"故障ID"`                           // 故障ID
	VehicleId        int         `json:"vehicle_id"        orm:"vehicle_id"        description:"车辆ID"`                           // 车辆ID
	ReporterId       int         `json:"reporter_id"       orm:"reporter_id"       description:"报告人ID"`                          // 报告人ID
	FaultType        string      `json:"fault_type"        orm:"fault_type"        description:"故障类型"`                           // 故障类型
	FaultDescription string      `json:"fault_description" orm:"fault_description" description:"故障描述"`                           // 故障描述
	ReportDate       *gtime.Time `json:"report_date"       orm:"report_date"       description:"报告时间"`                           // 报告时间
	FaultLocation    string      `json:"fault_location"    orm:"fault_location"    description:"故障发生地点"`                         // 故障发生地点
	Severity         int         `json:"severity"          orm:"severity"          description:"严重程度: 1-轻微, 2-一般, 3-严重, 4-危险"`   // 严重程度: 1-轻微, 2-一般, 3-严重, 4-危险
	Status           int         `json:"status"            orm:"status"            description:"状态: 0-已取消, 1-待处理, 2-处理中, 3-已解决"` // 状态: 0-已取消, 1-待处理, 2-处理中, 3-已解决
	MaintenanceId    int         `json:"maintenance_id"    orm:"maintenance_id"    description:"关联的维修记录ID"`                      // 关联的维修记录ID
	Notes            string      `json:"notes"             orm:"notes"             description:"备注"`                             // 备注
	CreatedAt        *gtime.Time `json:"created_at"        orm:"created_at"        description:"创建时间"`                           // 创建时间
	UpdatedAt        *gtime.Time `json:"updated_at"        orm:"updated_at"        description:"更新时间"`                           // 更新时间
}
