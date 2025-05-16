// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Route is the golang structure of table xf_route for DAO operations like Where/Data.
type Route struct {
	g.Meta         `orm:"table:xf_route, do:true"`
	RouteUuid      interface{} // 线路UUID
	RouteNumber    interface{} // 线路编号
	Name           interface{} // 线路名称
	StartStation   interface{} // 起始站点
	EndStation     interface{} // 终点站点
	Stops          *gjson.Json // 途经站点(JSON格式)
	Distance       interface{} // 线路长度(km)
	Fare           interface{} // 票价(元)
	OperationHours interface{} // 运营时间
	Frequency      interface{} // 发车频率
	Status         interface{} // 状态: 0-停运, 1-运营
	Notes          interface{} // 备注
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
}
