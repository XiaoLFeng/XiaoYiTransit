// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VehicleInsurance is the golang structure for table vehicle_insurance.
type VehicleInsurance struct {
	InsuranceUuid  string      `json:"insurance_uuid"  orm:"insurance_uuid"  description:"保险UUID"` // 保险UUID
	VehicleUuid    string      `json:"vehicle_uuid"    orm:"vehicle_uuid"    description:"车辆UUID"` // 车辆UUID
	InsuranceType  string      `json:"insurance_type"  orm:"insurance_type"  description:"保险类型"`   // 保险类型
	PolicyNumber   string      `json:"policy_number"   orm:"policy_number"   description:"保单号"`    // 保单号
	Insurer        string      `json:"insurer"         orm:"insurer"         description:"保险公司"`   // 保险公司
	StartDate      *gtime.Time `json:"start_date"      orm:"start_date"      description:"保险开始日期"` // 保险开始日期
	ExpiryDate     *gtime.Time `json:"expiry_date"     orm:"expiry_date"     description:"保险到期日期"` // 保险到期日期
	CoverageAmount float64     `json:"coverage_amount" orm:"coverage_amount" description:"保险金额"`   // 保险金额
	Premium        float64     `json:"premium"         orm:"premium"         description:"保费"`     // 保费
	PaymentDate    *gtime.Time `json:"payment_date"    orm:"payment_date"    description:"缴费日期"`   // 缴费日期
	PaymentMethod  string      `json:"payment_method"  orm:"payment_method"  description:"缴费方式"`   // 缴费方式
	ContactPerson  string      `json:"contact_person"  orm:"contact_person"  description:"联系人"`    // 联系人
	ContactPhone   string      `json:"contact_phone"   orm:"contact_phone"   description:"联系电话"`   // 联系电话
	Notes          string      `json:"notes"           orm:"notes"           description:"备注"`     // 备注
	CreatedAt      *gtime.Time `json:"created_at"      orm:"created_at"      description:"创建时间"`   // 创建时间
	UpdatedAt      *gtime.Time `json:"updated_at"      orm:"updated_at"      description:"更新时间"`   // 更新时间
	DeletedAt      *gtime.Time `json:"deleted_at"      orm:"deleted_at"      description:"删除时间"`   // 删除时间
}
