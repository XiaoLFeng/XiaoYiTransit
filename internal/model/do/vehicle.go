// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Vehicle is the golang structure of table xf_vehicle for DAO operations like Where/Data.
type Vehicle struct {
	g.Meta              `orm:"table:xf_vehicle, do:true"`
	Id                  interface{} // 车辆ID
	PlateNumber         interface{} // 车牌号
	Model               interface{} // 车辆型号
	PurchaseDate        *gtime.Time // 购买日期
	Status              interface{} // 状态: 1-运营, 2-维修, 3-停运, 4-报废
	Seats               interface{} // 座位数
	FuelType            interface{} // 燃料类型
	Mileage             interface{} // 行驶里程(km)
	LastMaintenanceDate *gtime.Time // 最后维护日期
	NextInspectionDate  *gtime.Time // 下次年检日期
	InsuranceExpiryDate *gtime.Time // 保险到期日期
	Notes               interface{} // 备注
	CreatedAt           *gtime.Time // 创建时间
	UpdatedAt           *gtime.Time // 更新时间
	DeletedAt           *gtime.Time // 删除时间
}
