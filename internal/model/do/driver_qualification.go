// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DriverQualification is the golang structure of table xf_driver_qualification for DAO operations like Where/Data.
type DriverQualification struct {
	g.Meta              `orm:"table:xf_driver_qualification, do:true"`
	Id                  interface{} // ID
	DriverId            interface{} // 司机ID
	QualificationType   interface{} // 资质类型
	QualificationNumber interface{} // 资质编号
	IssueDate           *gtime.Time // 发证日期
	ExpiryDate          *gtime.Time // 到期日期
	IssuingAuthority    interface{} // 发证机构
	Notes               interface{} // 备注
	CreatedAt           *gtime.Time // 创建时间
	UpdatedAt           *gtime.Time // 更新时间
}
