package route

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GetRouteListReqDTO defines the request for getting a list of routes
type GetRouteListReqDTO struct {
	Page        int    `json:"page" v:"min:1#页码最小为1" d:"1" dc:"页码"`
	Size        int    `json:"size" v:"min:1#每页数量最小为1" d:"10" dc:"每页数量"`
	RouteNumber string `json:"route_number,omitempty" dc:"线路编号（可选，用于筛选）"`
	Name        string `json:"name,omitempty" dc:"线路名称（可选，用于筛选）"`
	Status      int    `json:"status,omitempty" dc:"状态（可选，用于筛选）: 0-停运, 1-运营"`
}

// RouteListItemDTO defines a single route item in the list
type RouteListItemDTO struct {
	RouteUuid      string      `json:"route_uuid" dc:"线路UUID"`
	RouteNumber    string      `json:"route_number" dc:"线路编号"`
	Name           string      `json:"name" dc:"线路名称"`
	StartStation   string      `json:"start_station" dc:"起始站点"`
	EndStation     string      `json:"end_station" dc:"终点站点"`
	Distance       float64     `json:"distance" dc:"线路长度(km)"`
	Fare           float64     `json:"fare" dc:"票价(元)"`
	OperationHours string      `json:"operation_hours" dc:"运营时间"`
	Status         int         `json:"status" dc:"状态: 0-停运, 1-运营"`
	CreatedAt      *gtime.Time `json:"created_at" dc:"创建时间"`
}

// PagedRouteListDTO defines a paged list of routes
type PagedRouteListDTO struct {
	List  []*RouteListItemDTO `json:"list" dc:"线路列表"`
	Page  int                 `json:"page" dc:"当前页码"`
	Size  int                 `json:"size" dc:"每页数量"`
	Total int                 `json:"total" dc:"总数量"`
}

// RouteDetailDTO defines the detailed route information
type RouteDetailDTO struct {
	RouteUuid      string       `json:"route_uuid" dc:"线路UUID"`
	RouteNumber    string       `json:"route_number" dc:"线路编号"`
	Name           string       `json:"name" dc:"线路名称"`
	StartStation   string       `json:"start_station" dc:"起始站点"`
	EndStation     string       `json:"end_station" dc:"终点站点"`
	Stops          []StationDTO `json:"stops" dc:"途经站点"`
	Distance       float64      `json:"distance" dc:"线路长度(km)"`
	Fare           float64      `json:"fare" dc:"票价(元)"`
	OperationHours string       `json:"operation_hours" dc:"运营时间"`
	Frequency      string       `json:"frequency" dc:"发车频率"`
	Status         int          `json:"status" dc:"状态: 0-停运, 1-运营"`
	Notes          string       `json:"notes" dc:"备注"`
	CreatedAt      *gtime.Time  `json:"created_at" dc:"创建时间"`
	UpdatedAt      *gtime.Time  `json:"updated_at" dc:"更新时间"`
}

// StationDTO defines a station in the route
type StationDTO struct {
	StationUuid       string  `json:"station_uuid" dc:"站点UUID"`
	Name              string  `json:"name" dc:"站点名称"`
	Sequence          int     `json:"sequence" dc:"站点顺序"`
	DistanceFromStart float64 `json:"distance_from_start" dc:"距起点距离(km)"`
	EstimatedTime     int     `json:"estimated_time" dc:"预计到达时间(分钟)"`
}
