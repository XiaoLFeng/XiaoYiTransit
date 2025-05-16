// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RouteStation is the golang structure of table xf_route_station for DAO operations like Where/Data.
type RouteStation struct {
	g.Meta            `orm:"table:xf_route_station, do:true"`
	Id                interface{} // ID
	RouteId           interface{} // 线路ID
	StationId         interface{} // 站点ID
	Sequence          interface{} // 站点顺序
	DistanceFromStart interface{} // 距起点距离(km)
	EstimatedTime     interface{} // 预计到达时间(分钟)
	CreatedAt         *gtime.Time // 创建时间
	UpdatedAt         *gtime.Time // 更新时间
}
