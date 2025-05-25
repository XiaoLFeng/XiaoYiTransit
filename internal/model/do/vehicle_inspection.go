// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VehicleInspection is the golang structure of table xf_vehicle_inspection for DAO operations like Where/Data.
type VehicleInspection struct {
	g.Meta           `orm:"table:xf_vehicle_inspection, do:true"`
	InspectionUuid   interface{} // 年检UUID
	VehicleUuid      interface{} // 车辆UUID
	InspectionDate   *gtime.Time // 年检日期
	ExpiryDate       *gtime.Time // 有效期截止日期
	InspectionResult interface{} // 年检结果: 1-通过, 2-不通过
	InspectionAgency interface{} // 检测机构
	Inspector        interface{} // 检测人员
	CertificateNo    interface{} // 检测证书编号
	Cost             interface{} // 年检费用
	Notes            interface{} // 备注
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
	DeletedAt        *gtime.Time // 删除时间
}
