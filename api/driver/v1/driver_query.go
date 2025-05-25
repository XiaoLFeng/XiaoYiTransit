package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GetDriverListReq defines the request for getting a list of drivers
type GetDriverListReq struct {
	g.Meta     `path:"/driver/list" method:"Get" tags:"司机管理" sm:"获取司机列表" dc:"用于获取司机列表信息"`
	Page       int    `json:"page" d:"1" v:"min:1#页码最小为1" dc:"页码"`
	Size       int    `json:"size" d:"10" v:"max:100#每页最大100条" dc:"每页数量"`
	EmployeeId string `json:"employee_id,omitempty" dc:"工号（可选，用于筛选）"`
	Name       string `json:"name,omitempty" dc:"姓名（可选，用于筛选）"`
	Status     int    `json:"status,omitempty" dc:"状态（可选，用于筛选）: 0-离职, 1-在职, 2-休假, 3-停职"`
}

// DriverListItem defines a single driver item in the list
type DriverListItem struct {
	DriverUuid        string      `json:"driver_uuid" dc:"司机UUID"`
	EmployeeId        string      `json:"employee_id" dc:"工号"`
	Name              string      `json:"name" dc:"姓名"`
	Gender            int         `json:"gender" dc:"性别: 1-男, 2-女"`
	Phone             string      `json:"phone" dc:"联系电话"`
	LicenseNumber     string      `json:"license_number" dc:"驾驶证号"`
	LicenseExpiryDate *gtime.Time `json:"license_expiry_date" dc:"驾驶证到期日期"`
	Status            int         `json:"status" dc:"状态: 0-离职, 1-在职, 2-休假, 3-停职"`
	EntryDate         *gtime.Time `json:"entry_date" dc:"入职日期"`
}

// PagedDriverList defines a paged list of drivers
type PagedDriverList struct {
	List  []*DriverListItem `json:"list" dc:"司机列表"`
	Page  int               `json:"page" dc:"当前页码"`
	Size  int               `json:"size" dc:"每页数量"`
	Total int               `json:"total" dc:"总数量"`
}

// GetDriverListRes defines the response for getting a list of drivers
type GetDriverListRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*PagedDriverList]
}

// GetDriverDetailReq defines the request for getting driver details
type GetDriverDetailReq struct {
	g.Meta     `path:"/driver/detail" method:"Get" tags:"司机管理" sm:"获取司机详情" dc:"用于获取司机详细信息"`
	DriverUuid string `json:"driver_uuid" v:"required#司机UUID不能为空" dc:"司机UUID"`
}

// DriverDetailItem defines the detailed driver information
type DriverDetailItem struct {
	DriverUuid        string      `json:"driver_uuid" dc:"司机UUID"`
	EmployeeId        string      `json:"employee_id" dc:"工号"`
	Name              string      `json:"name" dc:"姓名"`
	Gender            int         `json:"gender" dc:"性别: 1-男, 2-女"`
	IdCard            string      `json:"id_card" dc:"身份证号"`
	Phone             string      `json:"phone" dc:"联系电话"`
	EmergencyContact  string      `json:"emergency_contact" dc:"紧急联系人"`
	EmergencyPhone    string      `json:"emergency_phone" dc:"紧急联系电话"`
	LicenseNumber     string      `json:"license_number" dc:"驾驶证号"`
	LicenseType       string      `json:"license_type" dc:"驾驶证类型"`
	LicenseExpiryDate *gtime.Time `json:"license_expiry_date" dc:"驾驶证到期日期"`
	EntryDate         *gtime.Time `json:"entry_date" dc:"入职日期"`
	Status            int         `json:"status" dc:"状态: 0-离职, 1-在职, 2-休假, 3-停职"`
	Address           string      `json:"address" dc:"住址"`
	Notes             string      `json:"notes" dc:"备注"`
	CreatedAt         *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt         *gtime.Time `json:"updated_at" dc:"更新时间"`
}

// GetDriverDetailRes defines the response for getting driver details
type GetDriverDetailRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*DriverDetailItem]
}
