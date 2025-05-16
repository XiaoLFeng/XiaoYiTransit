// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VehicleFault is the golang structure for table vehicle_fault.
type VehicleFault struct {
	VehicleFaultUuid string      `json:"vehicle_fault_uuid" orm:"vehicle_fault_uuid" description:"故障UUID"`                         // 故障UUID
	VehicleUuid      string      `json:"vehicle_uuid"       orm:"vehicle_uuid"       description:"车辆UUID"`                         // 车辆UUID
	UserUuid         string      `json:"user_uuid"          orm:"user_uuid"          description:"报告人UUID"`                        // 报告人UUID
	FaultType        string      `json:"fault_type"         orm:"fault_type"         description:"故障类型"`                           // 故障类型
	FaultDescription string      `json:"fault_description"  orm:"fault_description"  description:"故障描述"`                           // 故障描述
	ReportDate       *gtime.Time `json:"report_date"        orm:"report_date"        description:"报告时间"`                           // 报告时间
	FaultLocation    string      `json:"fault_location"     orm:"fault_location"     description:"故障发生地点"`                         // 故障发生地点
	Severity         int         `json:"severity"           orm:"severity"           description:"严重程度: 1-轻微, 2-一般, 3-严重, 4-危险"`   // 严重程度: 1-轻微, 2-一般, 3-严重, 4-危险
	Status           int         `json:"status"             orm:"status"             description:"状态: 0-已取消, 1-待处理, 2-处理中, 3-已解决"` // 状态: 0-已取消, 1-待处理, 2-处理中, 3-已解决
	MaintenanceUuid  string      `json:"maintenance_uuid"   orm:"maintenance_uuid"   description:"关联的维修记录UUID"`                    // 关联的维修记录UUID
	Notes            string      `json:"notes"              orm:"notes"              description:"备注"`                             // 备注
	CreatedAt        *gtime.Time `json:"created_at"         orm:"created_at"         description:"创建时间"`                           // 创建时间
	UpdatedAt        *gtime.Time `json:"updated_at"         orm:"updated_at"         description:"更新时间"`                           // 更新时间
}
