// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MaintenanceItem is the golang structure of table xf_maintenance_item for DAO operations like Where/Data.
type MaintenanceItem struct {
	g.Meta          `orm:"table:xf_maintenance_item, do:true"`
	Id              interface{} // 项目ID
	MaintenanceId   interface{} // 维修记录ID
	ItemName        interface{} // 项目名称
	ItemCost        interface{} // 项目费用
	ItemDescription interface{} // 项目描述
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
