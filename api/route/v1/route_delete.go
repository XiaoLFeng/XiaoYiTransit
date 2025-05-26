package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"go/types"
)

// DeleteRouteReq defines the request for deleting a route
type DeleteRouteReq struct {
	g.Meta    `path:"/route/{routeUuid}" method:"Delete" tags:"线路管理" sm:"删除线路" dc:"用于删除公交线路信息"`
	RouteUuid string `json:"route_uuid" v:"required#线路UUID不能为空" dc:"线路UUID"`
}

// DeleteRouteRes defines the response for deleting a route
type DeleteRouteRes struct {
	g.Meta                           `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}
