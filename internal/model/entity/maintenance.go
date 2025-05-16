// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Maintenance is the golang structure for table maintenance.
type Maintenance struct {
	Id                  int         `json:"id"                   orm:"id"                   description:"维修ID"`                                 // 维修ID
	VehicleId           int         `json:"vehicle_id"           orm:"vehicle_id"           description:"车辆ID"`                                 // 车辆ID
	MaintenanceType     int         `json:"maintenance_type"     orm:"maintenance_type"     description:"维修类型: 1-常规保养, 2-故障维修, 3-事故维修, 4-年检维修"` // 维修类型: 1-常规保养, 2-故障维修, 3-事故维修, 4-年检维修
	Description         string      `json:"description"          orm:"description"          description:"维修描述"`                                 // 维修描述
	MaintenanceDate     *gtime.Time `json:"maintenance_date"     orm:"maintenance_date"     description:"维修日期"`                                 // 维修日期
	CompletionDate      *gtime.Time `json:"completion_date"      orm:"completion_date"      description:"完成日期"`                                 // 完成日期
	Cost                float64     `json:"cost"                 orm:"cost"                 description:"维修费用"`                                 // 维修费用
	Mileage             float64     `json:"mileage"              orm:"mileage"              description:"维修时里程数"`                               // 维修时里程数
	MaintenanceLocation string      `json:"maintenance_location" orm:"maintenance_location" description:"维修地点"`                                 // 维修地点
	MaintenanceStaff    string      `json:"maintenance_staff"    orm:"maintenance_staff"    description:"维修人员"`                                 // 维修人员
	PartsReplaced       string      `json:"parts_replaced"       orm:"parts_replaced"       description:"更换的零部件"`                               // 更换的零部件
	Status              int         `json:"status"               orm:"status"               description:"状态: 0-已取消, 1-待维修, 2-维修中, 3-已完成"`       // 状态: 0-已取消, 1-待维修, 2-维修中, 3-已完成
	Notes               string      `json:"notes"                orm:"notes"                description:"备注"`                                   // 备注
	CreatedBy           int         `json:"created_by"           orm:"created_by"           description:"创建人ID"`                                // 创建人ID
	CreatedAt           *gtime.Time `json:"created_at"           orm:"created_at"           description:"创建时间"`                                 // 创建时间
	UpdatedAt           *gtime.Time `json:"updated_at"           orm:"updated_at"           description:"更新时间"`                                 // 更新时间
	DeletedAt           *gtime.Time `json:"deleted_at"           orm:"deleted_at"           description:"删除时间"`                                 // 删除时间
}
