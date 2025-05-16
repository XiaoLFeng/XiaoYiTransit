// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RouteStation is the golang structure for table route_station.
type RouteStation struct {
	Id                int         `json:"id"                  orm:"id"                  description:"ID"`         // ID
	RouteId           int         `json:"route_id"            orm:"route_id"            description:"线路ID"`       // 线路ID
	StationId         int         `json:"station_id"          orm:"station_id"          description:"站点ID"`       // 站点ID
	Sequence          int         `json:"sequence"            orm:"sequence"            description:"站点顺序"`       // 站点顺序
	DistanceFromStart float64     `json:"distance_from_start" orm:"distance_from_start" description:"距起点距离(km)"`  // 距起点距离(km)
	EstimatedTime     int         `json:"estimated_time"      orm:"estimated_time"      description:"预计到达时间(分钟)"` // 预计到达时间(分钟)
	CreatedAt         *gtime.Time `json:"created_at"          orm:"created_at"          description:"创建时间"`       // 创建时间
	UpdatedAt         *gtime.Time `json:"updated_at"          orm:"updated_at"          description:"更新时间"`       // 更新时间
}
