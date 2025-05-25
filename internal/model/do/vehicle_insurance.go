// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VehicleInsurance is the golang structure of table xf_vehicle_insurance for DAO operations like Where/Data.
type VehicleInsurance struct {
	g.Meta         `orm:"table:xf_vehicle_insurance, do:true"`
	InsuranceUuid  interface{} // 保险UUID
	VehicleUuid    interface{} // 车辆UUID
	InsuranceType  interface{} // 保险类型
	PolicyNumber   interface{} // 保单号
	Insurer        interface{} // 保险公司
	StartDate      *gtime.Time // 保险开始日期
	ExpiryDate     *gtime.Time // 保险到期日期
	CoverageAmount interface{} // 保险金额
	Premium        interface{} // 保费
	PaymentDate    *gtime.Time // 缴费日期
	PaymentMethod  interface{} // 缴费方式
	ContactPerson  interface{} // 联系人
	ContactPhone   interface{} // 联系电话
	Notes          interface{} // 备注
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
}
