package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"go/types"
)

// UpdateDriverReq defines the request for updating a driver
type UpdateDriverReq struct {
	g.Meta            `path:"/driver" method:"Put" tags:"司机管理" sm:"更新司机信息" dc:"用于更新司机信息"`
	DriverUuid        string      `json:"driver_uuid" v:"required#司机UUID不能为空" dc:"司机UUID"`
	EmployeeId        string      `json:"employee_id" v:"required#工号不能为空" dc:"工号"`
	Name              string      `json:"name" v:"required#姓名不能为空" dc:"姓名"`
	Gender            int         `json:"gender" v:"required|in:1,2#性别不能为空|性别只能是1(男)或2(女)" dc:"性别: 1-男, 2-女"`
	IdCard            string      `json:"id_card" v:"required#身份证号不能为空" dc:"身份证号"`
	Phone             string      `json:"phone" v:"required#联系电话不能为空" dc:"联系电话"`
	EmergencyContact  string      `json:"emergency_contact" dc:"紧急联系人"`
	EmergencyPhone    string      `json:"emergency_phone" dc:"紧急联系电话"`
	LicenseNumber     string      `json:"license_number" v:"required#驾驶证号不能为空" dc:"驾驶证号"`
	LicenseType       string      `json:"license_type" v:"required#驾驶证类型不能为空" dc:"驾驶证类型"`
	LicenseExpiryDate *gtime.Time `json:"license_expiry_date" dc:"驾驶证到期日期"`
	EntryDate         *gtime.Time `json:"entry_date" dc:"入职日期"`
	Status            int         `json:"status" v:"required|in:0,1,2,3#状态不能为空|状态只能是0(离职),1(在职),2(休假),3(停职)" dc:"状态: 0-离职, 1-在职, 2-休假, 3-停职"`
	Address           string      `json:"address" dc:"住址"`
	Notes             string      `json:"notes" dc:"备注"`
}

// UpdateDriverRes defines the response for updating a driver
type UpdateDriverRes struct {
	g.Meta                           `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}
