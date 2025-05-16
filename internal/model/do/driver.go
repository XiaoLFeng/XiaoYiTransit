// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Driver is the golang structure of table xf_driver for DAO operations like Where/Data.
type Driver struct {
	g.Meta            `orm:"table:xf_driver, do:true"`
	Id                interface{} // 司机ID
	EmployeeId        interface{} // 工号
	Name              interface{} // 姓名
	Gender            interface{} // 性别: 1-男, 2-女
	IdCard            interface{} // 身份证号
	Phone             interface{} // 联系电话
	EmergencyContact  interface{} // 紧急联系人
	EmergencyPhone    interface{} // 紧急联系电话
	LicenseNumber     interface{} // 驾驶证号
	LicenseType       interface{} // 驾驶证类型
	LicenseExpiryDate *gtime.Time // 驾驶证到期日期
	EntryDate         *gtime.Time // 入职日期
	Status            interface{} // 状态: 0-离职, 1-在职, 2-休假, 3-停职
	Address           interface{} // 住址
	Notes             interface{} // 备注
	CreatedAt         *gtime.Time // 创建时间
	UpdatedAt         *gtime.Time // 更新时间
	DeletedAt         *gtime.Time // 删除时间
}
