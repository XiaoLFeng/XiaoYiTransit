// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Vehicle is the golang structure for table vehicle.
type Vehicle struct {
	Id                  int         `json:"id"                    orm:"id"                    description:"车辆ID"`                       // 车辆ID
	PlateNumber         string      `json:"plate_number"          orm:"plate_number"          description:"车牌号"`                        // 车牌号
	Model               string      `json:"model"                 orm:"model"                 description:"车辆型号"`                       // 车辆型号
	PurchaseDate        *gtime.Time `json:"purchase_date"         orm:"purchase_date"         description:"购买日期"`                       // 购买日期
	Status              int         `json:"status"                orm:"status"                description:"状态: 1-运营, 2-维修, 3-停运, 4-报废"` // 状态: 1-运营, 2-维修, 3-停运, 4-报废
	Seats               int         `json:"seats"                 orm:"seats"                 description:"座位数"`                        // 座位数
	FuelType            string      `json:"fuel_type"             orm:"fuel_type"             description:"燃料类型"`                       // 燃料类型
	Mileage             float64     `json:"mileage"               orm:"mileage"               description:"行驶里程(km)"`                   // 行驶里程(km)
	LastMaintenanceDate *gtime.Time `json:"last_maintenance_date" orm:"last_maintenance_date" description:"最后维护日期"`                     // 最后维护日期
	NextInspectionDate  *gtime.Time `json:"next_inspection_date"  orm:"next_inspection_date"  description:"下次年检日期"`                     // 下次年检日期
	InsuranceExpiryDate *gtime.Time `json:"insurance_expiry_date" orm:"insurance_expiry_date" description:"保险到期日期"`                     // 保险到期日期
	Notes               string      `json:"notes"                 orm:"notes"                 description:"备注"`                         // 备注
	CreatedAt           *gtime.Time `json:"created_at"            orm:"created_at"            description:"创建时间"`                       // 创建时间
	UpdatedAt           *gtime.Time `json:"updated_at"            orm:"updated_at"            description:"更新时间"`                       // 更新时间
	DeletedAt           *gtime.Time `json:"deleted_at"            orm:"deleted_at"            description:"删除时间"`                       // 删除时间
}
