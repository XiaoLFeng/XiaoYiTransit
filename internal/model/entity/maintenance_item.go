// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MaintenanceItem is the golang structure for table maintenance_item.
type MaintenanceItem struct {
	Id              int         `json:"id"               orm:"id"               description:"项目ID"`   // 项目ID
	MaintenanceId   int         `json:"maintenance_id"   orm:"maintenance_id"   description:"维修记录ID"` // 维修记录ID
	ItemName        string      `json:"item_name"        orm:"item_name"        description:"项目名称"`   // 项目名称
	ItemCost        float64     `json:"item_cost"        orm:"item_cost"        description:"项目费用"`   // 项目费用
	ItemDescription string      `json:"item_description" orm:"item_description" description:"项目描述"`   // 项目描述
	CreatedAt       *gtime.Time `json:"created_at"       orm:"created_at"       description:"创建时间"`   // 创建时间
	UpdatedAt       *gtime.Time `json:"updated_at"       orm:"updated_at"       description:"更新时间"`   // 更新时间
}
