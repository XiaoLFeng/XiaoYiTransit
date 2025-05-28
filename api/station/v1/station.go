package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"go/types"
)

// CreateStationReq defines the request for creating a station
type CreateStationReq struct {
	g.Meta     `path:"/station" method:"Post" tags:"站点管理" sm:"创建站点" dc:"用于创建新的公交站点"`
	Name       string  `json:"name" v:"required#站点名称不能为空" dc:"站点名称"`
	Code       string  `json:"code" v:"required#站点编码不能为空" dc:"站点编码"`
	Address    string  `json:"address" v:"required#站点地址不能为空" dc:"站点地址"`
	Longitude  float64 `json:"longitude" v:"required#经度不能为空" dc:"经度"`
	Latitude   float64 `json:"latitude" v:"required#纬度不能为空" dc:"纬度"`
	Status     int     `json:"status" v:"required|in:0,1#状态不能为空|状态只能是0或1" dc:"状态: 0-停用, 1-启用"`
}

// CreateStationRes defines the response for creating a station
type CreateStationRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*StationDTO]
}

// UpdateStationReq defines the request for updating a station
type UpdateStationReq struct {
	g.Meta      `path:"/station/{stationUuid}" method:"Put" tags:"站点管理" sm:"更新站点" dc:"用于更新公交站点信息"`
	StationUuid string  `json:"station_uuid" v:"required#站点UUID不能为空" dc:"站点UUID"`
	Name        string  `json:"name" v:"required#站点名称不能为空" dc:"站点名称"`
	Code        string  `json:"code" v:"required#站点编码不能为空" dc:"站点编码"`
	Address     string  `json:"address" v:"required#站点地址不能为空" dc:"站点地址"`
	Longitude   float64 `json:"longitude" v:"required#经度不能为空" dc:"经度"`
	Latitude    float64 `json:"latitude" v:"required#纬度不能为空" dc:"纬度"`
	Status      int     `json:"status" v:"required|in:0,1#状态不能为空|状态只能是0或1" dc:"状态: 0-停用, 1-启用"`
}

// UpdateStationRes defines the response for updating a station
type UpdateStationRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}

// DeleteStationReq defines the request for deleting a station
type DeleteStationReq struct {
	g.Meta      `path:"/station/{stationUuid}" method:"Delete" tags:"站点管理" sm:"删除站点" dc:"用于删除公交站点"`
	StationUuid string `json:"station_uuid" v:"required#站点UUID不能为空" dc:"站点UUID"`
}

// DeleteStationRes defines the response for deleting a station
type DeleteStationRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}

// GetStationReq defines the request for getting a station
type GetStationReq struct {
	g.Meta      `path:"/station/{stationUuid}" method:"Get" tags:"站点管理" sm:"获取站点" dc:"用于获取公交站点详情"`
	StationUuid string `json:"station_uuid" v:"required#站点UUID不能为空" dc:"站点UUID"`
}

// GetStationRes defines the response for getting a station
type GetStationRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*StationDTO]
}

// ListStationsReq defines the request for listing stations
type ListStationsReq struct {
	g.Meta `path:"/stations" method:"Get" tags:"站点管理" sm:"获取站点列表" dc:"用于获取公交站点列表"`
	Name   string `json:"name" dc:"站点名称，用于模糊查询"`
	Code   string `json:"code" dc:"站点编码，用于模糊查询"`
	Status int    `json:"status" dc:"状态: 0-停用, 1-启用"`
	Page   int    `json:"page" v:"min:1#页码最小为1" dc:"页码，默认为1"`
	Size   int    `json:"size" v:"min:1#每页数量最小为1" dc:"每页数量，默认为10"`
}

// ListStationsRes defines the response for listing stations
type ListStationsRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*StationsListDTO]
}

// StationDTO defines a station
type StationDTO struct {
	StationUuid string  `json:"station_uuid" dc:"站点UUID"`
	Name        string  `json:"name" dc:"站点名称"`
	Code        string  `json:"code" dc:"站点编码"`
	Address     string  `json:"address" dc:"站点地址"`
	Longitude   float64 `json:"longitude" dc:"经度"`
	Latitude    float64 `json:"latitude" dc:"纬度"`
	Status      int     `json:"status" dc:"状态: 0-停用, 1-启用"`
	CreatedAt   string  `json:"created_at" dc:"创建时间"`
	UpdatedAt   string  `json:"updated_at" dc:"更新时间"`
}

// StationsListDTO defines a list of stations
type StationsListDTO struct {
	Total    int          `json:"total" dc:"总数"`
	Page     int          `json:"page" dc:"页码"`
	Size     int          `json:"size" dc:"每页数量"`
	Stations []*StationDTO `json:"stations" dc:"站点列表"`
}

// ListStationsSimpleReq defines the request for listing stations in a simple format (for dropdowns)
type ListStationsSimpleReq struct {
	g.Meta `path:"/stations/simple" method:"Get" tags:"站点管理" sm:"获取站点简单列表" dc:"用于获取公交站点简单列表，适用于下拉框"`
	Status int    `json:"status" dc:"状态: 0-停用, 1-启用，默认为1（启用）"`
}

// ListStationsSimpleRes defines the response for listing stations in a simple format
type ListStationsSimpleRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*StationsSimpleListDTO]
}

// StationSimpleDTO defines a simplified station (for dropdowns)
type StationSimpleDTO struct {
	StationUuid string `json:"station_uuid" dc:"站点UUID"`
	Name        string `json:"name" dc:"站点名称"`
	Code        string `json:"code" dc:"站点编码"`
}

// StationsSimpleListDTO defines a list of simplified stations
type StationsSimpleListDTO struct {
	Stations []*StationSimpleDTO `json:"stations" dc:"站点列表"`
}
