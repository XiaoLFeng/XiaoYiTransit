package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"go/types"
)

// UpdateRouteReq defines the request for updating a route
type UpdateRouteReq struct {
	g.Meta         `path:"/route/{routeUuid}" method:"Put" tags:"线路管理" sm:"更新线路" dc:"用于更新公交线路信息"`
	RouteUuid      string      `json:"route_uuid" v:"required#线路UUID不能为空" dc:"线路UUID"`
	RouteNumber    string      `json:"route_number" v:"required#线路编号不能为空" dc:"线路编号"`
	Name           string      `json:"name" v:"required#线路名称不能为空" dc:"线路名称"`
	StartStation   string      `json:"start_station" v:"required#起始站点不能为空" dc:"起始站点"`
	EndStation     string      `json:"end_station" v:"required#终点站点不能为空" dc:"终点站点"`
	Stops          *gjson.Json `json:"stops" dc:"途经站点(JSON格式)"`
	Distance       float64     `json:"distance" dc:"线路长度(km)"`
	Fare           float64     `json:"fare" v:"required#票价不能为空" dc:"票价(元)"`
	OperationHours string      `json:"operation_hours" dc:"运营时间"`
	Frequency      string      `json:"frequency" dc:"发车频率"`
	Status         int         `json:"status" v:"required|in:0,1#状态不能为空|状态只能是0(停运)或1(运营)" dc:"状态: 0-停运, 1-运营"`
	Notes          string      `json:"notes" dc:"备注"`
}

// UpdateRouteRes defines the response for updating a route
type UpdateRouteRes struct {
	g.Meta                           `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}
