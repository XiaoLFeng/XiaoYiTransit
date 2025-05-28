// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package station

import (
	"context"

	"xiao-yi-transit/api/station/v1"
)

type IStationV1 interface {
	CreateStation(ctx context.Context, req *v1.CreateStationReq) (res *v1.CreateStationRes, err error)
	UpdateStation(ctx context.Context, req *v1.UpdateStationReq) (res *v1.UpdateStationRes, err error)
	DeleteStation(ctx context.Context, req *v1.DeleteStationReq) (res *v1.DeleteStationRes, err error)
	GetStation(ctx context.Context, req *v1.GetStationReq) (res *v1.GetStationRes, err error)
	ListStations(ctx context.Context, req *v1.ListStationsReq) (res *v1.ListStationsRes, err error)
	ListStationsSimple(ctx context.Context, req *v1.ListStationsSimpleReq) (res *v1.ListStationsSimpleRes, err error)
}
