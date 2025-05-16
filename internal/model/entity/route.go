// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// Route is the golang structure for table route.
type Route struct {
	RouteUuid      string      `json:"route_uuid"      orm:"route_uuid"      description:"线路UUID"`         // 线路UUID
	RouteNumber    string      `json:"route_number"    orm:"route_number"    description:"线路编号"`           // 线路编号
	Name           string      `json:"name"            orm:"name"            description:"线路名称"`           // 线路名称
	StartStation   string      `json:"start_station"   orm:"start_station"   description:"起始站点"`           // 起始站点
	EndStation     string      `json:"end_station"     orm:"end_station"     description:"终点站点"`           // 终点站点
	Stops          *gjson.Json `json:"stops"           orm:"stops"           description:"途经站点(JSON格式)"`   // 途经站点(JSON格式)
	Distance       float64     `json:"distance"        orm:"distance"        description:"线路长度(km)"`       // 线路长度(km)
	Fare           float64     `json:"fare"            orm:"fare"            description:"票价(元)"`          // 票价(元)
	OperationHours string      `json:"operation_hours" orm:"operation_hours" description:"运营时间"`           // 运营时间
	Frequency      string      `json:"frequency"       orm:"frequency"       description:"发车频率"`           // 发车频率
	Status         int         `json:"status"          orm:"status"          description:"状态: 0-停运, 1-运营"` // 状态: 0-停运, 1-运营
	Notes          string      `json:"notes"           orm:"notes"           description:"备注"`             // 备注
	CreatedAt      *gtime.Time `json:"created_at"      orm:"created_at"      description:"创建时间"`           // 创建时间
	UpdatedAt      *gtime.Time `json:"updated_at"      orm:"updated_at"      description:"更新时间"`           // 更新时间
	DeletedAt      *gtime.Time `json:"deleted_at"      orm:"deleted_at"      description:"删除时间"`           // 删除时间
}
