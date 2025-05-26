// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package route

import (
	"context"

	"xiao-yi-transit/api/route/v1"
)

type IRouteV1 interface {
	CreateRoute(ctx context.Context, req *v1.CreateRouteReq) (res *v1.CreateRouteRes, err error)
	DeleteRoute(ctx context.Context, req *v1.DeleteRouteReq) (res *v1.DeleteRouteRes, err error)
	GetRouteList(ctx context.Context, req *v1.GetRouteListReq) (res *v1.GetRouteListRes, err error)
	GetRouteDetail(ctx context.Context, req *v1.GetRouteDetailReq) (res *v1.GetRouteDetailRes, err error)
	AddRouteStation(ctx context.Context, req *v1.AddRouteStationReq) (res *v1.AddRouteStationRes, err error)
	UpdateRouteStation(ctx context.Context, req *v1.UpdateRouteStationReq) (res *v1.UpdateRouteStationRes, err error)
	DeleteRouteStation(ctx context.Context, req *v1.DeleteRouteStationReq) (res *v1.DeleteRouteStationRes, err error)
	GetRouteStations(ctx context.Context, req *v1.GetRouteStationsReq) (res *v1.GetRouteStationsRes, err error)
	UpdateRoute(ctx context.Context, req *v1.UpdateRouteReq) (res *v1.UpdateRouteRes, err error)
}
