// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Driver is the golang structure for table driver.
type Driver struct {
	Id                int         `json:"id"                  orm:"id"                  description:"司机ID"`                       // 司机ID
	EmployeeId        string      `json:"employee_id"         orm:"employee_id"         description:"工号"`                         // 工号
	Name              string      `json:"name"                orm:"name"                description:"姓名"`                         // 姓名
	Gender            int         `json:"gender"              orm:"gender"              description:"性别: 1-男, 2-女"`               // 性别: 1-男, 2-女
	IdCard            string      `json:"id_card"             orm:"id_card"             description:"身份证号"`                       // 身份证号
	Phone             string      `json:"phone"               orm:"phone"               description:"联系电话"`                       // 联系电话
	EmergencyContact  string      `json:"emergency_contact"   orm:"emergency_contact"   description:"紧急联系人"`                      // 紧急联系人
	EmergencyPhone    string      `json:"emergency_phone"     orm:"emergency_phone"     description:"紧急联系电话"`                     // 紧急联系电话
	LicenseNumber     string      `json:"license_number"      orm:"license_number"      description:"驾驶证号"`                       // 驾驶证号
	LicenseType       string      `json:"license_type"        orm:"license_type"        description:"驾驶证类型"`                      // 驾驶证类型
	LicenseExpiryDate *gtime.Time `json:"license_expiry_date" orm:"license_expiry_date" description:"驾驶证到期日期"`                    // 驾驶证到期日期
	EntryDate         *gtime.Time `json:"entry_date"          orm:"entry_date"          description:"入职日期"`                       // 入职日期
	Status            int         `json:"status"              orm:"status"              description:"状态: 0-离职, 1-在职, 2-休假, 3-停职"` // 状态: 0-离职, 1-在职, 2-休假, 3-停职
	Address           string      `json:"address"             orm:"address"             description:"住址"`                         // 住址
	Notes             string      `json:"notes"               orm:"notes"               description:"备注"`                         // 备注
	CreatedAt         *gtime.Time `json:"created_at"          orm:"created_at"          description:"创建时间"`                       // 创建时间
	UpdatedAt         *gtime.Time `json:"updated_at"          orm:"updated_at"          description:"更新时间"`                       // 更新时间
	DeletedAt         *gtime.Time `json:"deleted_at"          orm:"deleted_at"          description:"删除时间"`                       // 删除时间
}
