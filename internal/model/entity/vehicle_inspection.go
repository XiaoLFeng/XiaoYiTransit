// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VehicleInspection is the golang structure for table vehicle_inspection.
type VehicleInspection struct {
	InspectionUuid   string      `json:"inspection_uuid"   orm:"inspection_uuid"   description:"年检UUID"`            // 年检UUID
	VehicleUuid      string      `json:"vehicle_uuid"      orm:"vehicle_uuid"      description:"车辆UUID"`            // 车辆UUID
	InspectionDate   *gtime.Time `json:"inspection_date"   orm:"inspection_date"   description:"年检日期"`              // 年检日期
	ExpiryDate       *gtime.Time `json:"expiry_date"       orm:"expiry_date"       description:"有效期截止日期"`           // 有效期截止日期
	InspectionResult int         `json:"inspection_result" orm:"inspection_result" description:"年检结果: 1-通过, 2-不通过"` // 年检结果: 1-通过, 2-不通过
	InspectionAgency string      `json:"inspection_agency" orm:"inspection_agency" description:"检测机构"`              // 检测机构
	Inspector        string      `json:"inspector"         orm:"inspector"         description:"检测人员"`              // 检测人员
	CertificateNo    string      `json:"certificate_no"    orm:"certificate_no"    description:"检测证书编号"`            // 检测证书编号
	Cost             float64     `json:"cost"              orm:"cost"              description:"年检费用"`              // 年检费用
	Notes            string      `json:"notes"             orm:"notes"             description:"备注"`                // 备注
	CreatedAt        *gtime.Time `json:"created_at"        orm:"created_at"        description:"创建时间"`              // 创建时间
	UpdatedAt        *gtime.Time `json:"updated_at"        orm:"updated_at"        description:"更新时间"`              // 更新时间
	DeletedAt        *gtime.Time `json:"deleted_at"        orm:"deleted_at"        description:"删除时间"`              // 删除时间
}
