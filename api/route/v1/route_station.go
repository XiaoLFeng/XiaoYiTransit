package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"go/types"
)

// AddRouteStationReq defines the request for adding a station to a route
type AddRouteStationReq struct {
	g.Meta            `path:"/route/{routeUuid}/station" method:"Post" tags:"线路管理" sm:"添加站点" dc:"用于向公交线路添加站点"`
	RouteUuid         string  `json:"route_uuid" v:"required#线路UUID不能为空" dc:"线路UUID"`
	StationUuid       string  `json:"station_uuid" v:"required#站点UUID不能为空" dc:"站点UUID"`
	Sequence          int     `json:"sequence" v:"required#站点顺序不能为空" dc:"站点顺序"`
	DistanceFromStart float64 `json:"distance_from_start" v:"required#距起点距离不能为空" dc:"距起点距离(km)"`
	EstimatedTime     int     `json:"estimated_time" v:"required#预计到达时间不能为空" dc:"预计到达时间(分钟)"`
}

// AddRouteStationRes defines the response for adding a station to a route
type AddRouteStationRes struct {
	g.Meta                           `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}

// UpdateRouteStationReq defines the request for updating a station in a route
type UpdateRouteStationReq struct {
	g.Meta            `path:"/route/station/{routeStationUuid}" method:"Put" tags:"线路管理" sm:"更新站点" dc:"用于更新公交线路站点信息"`
	RouteStationUuid  string  `json:"route_station_uuid" v:"required#线路站点UUID不能为空" dc:"线路站点UUID"`
	Sequence          int     `json:"sequence" v:"required#站点顺序不能为空" dc:"站点顺序"`
	DistanceFromStart float64 `json:"distance_from_start" v:"required#距起点距离不能为空" dc:"距起点距离(km)"`
	EstimatedTime     int     `json:"estimated_time" v:"required#预计到达时间不能为空" dc:"预计到达时间(分钟)"`
}

// UpdateRouteStationRes defines the response for updating a station in a route
type UpdateRouteStationRes struct {
	g.Meta                           `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}

// DeleteRouteStationReq defines the request for deleting a station from a route
type DeleteRouteStationReq struct {
	g.Meta           `path:"/route/station/{routeStationUuid}" method:"Delete" tags:"线路管理" sm:"删除站点" dc:"用于从公交线路删除站点"`
	RouteStationUuid string `json:"route_station_uuid" v:"required#线路站点UUID不能为空" dc:"线路站点UUID"`
}

// DeleteRouteStationRes defines the response for deleting a station from a route
type DeleteRouteStationRes struct {
	g.Meta                           `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}

// GetRouteStationsReq defines the request for getting stations in a route
type GetRouteStationsReq struct {
	g.Meta    `path:"/route/{routeUuid}/stations" method:"Get" tags:"线路管理" sm:"获取线路站点" dc:"用于获取公交线路站点列表"`
	RouteUuid string `json:"route_uuid" v:"required#线路UUID不能为空" dc:"线路UUID"`
}

// GetRouteStationsRes defines the response for getting stations in a route
type GetRouteStationsRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*RouteStationsDTO]
}

// RouteStationsDTO defines a list of stations in a route
type RouteStationsDTO struct {
	RouteUuid string                 `json:"route_uuid" dc:"线路UUID"`
	Stations  []*RouteStationItemDTO `json:"stations" dc:"站点列表"`
}

// RouteStationItemDTO defines a single station in a route
type RouteStationItemDTO struct {
	RouteStationUuid  string  `json:"route_station_uuid" dc:"线路站点UUID"`
	StationUuid       string  `json:"station_uuid" dc:"站点UUID"`
	Name              string  `json:"name" dc:"站点名称"`
	Code              string  `json:"code" dc:"站点编码"`
	Address           string  `json:"address" dc:"站点地址"`
	Longitude         float64 `json:"longitude" dc:"经度"`
	Latitude          float64 `json:"latitude" dc:"纬度"`
	Sequence          int     `json:"sequence" dc:"站点顺序"`
	DistanceFromStart float64 `json:"distance_from_start" dc:"距起点距离(km)"`
	EstimatedTime     int     `json:"estimated_time" dc:"预计到达时间(分钟)"`
}
