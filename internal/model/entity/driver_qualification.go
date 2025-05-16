// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DriverQualification is the golang structure for table driver_qualification.
type DriverQualification struct {
	Id                  int         `json:"id"                   orm:"id"                   description:"ID"`   // ID
	DriverId            int         `json:"driver_id"            orm:"driver_id"            description:"司机ID"` // 司机ID
	QualificationType   string      `json:"qualification_type"   orm:"qualification_type"   description:"资质类型"` // 资质类型
	QualificationNumber string      `json:"qualification_number" orm:"qualification_number" description:"资质编号"` // 资质编号
	IssueDate           *gtime.Time `json:"issue_date"           orm:"issue_date"           description:"发证日期"` // 发证日期
	ExpiryDate          *gtime.Time `json:"expiry_date"          orm:"expiry_date"          description:"到期日期"` // 到期日期
	IssuingAuthority    string      `json:"issuing_authority"    orm:"issuing_authority"    description:"发证机构"` // 发证机构
	Notes               string      `json:"notes"                orm:"notes"                description:"备注"`   // 备注
	CreatedAt           *gtime.Time `json:"created_at"           orm:"created_at"           description:"创建时间"` // 创建时间
	UpdatedAt           *gtime.Time `json:"updated_at"           orm:"updated_at"           description:"更新时间"` // 更新时间
}
