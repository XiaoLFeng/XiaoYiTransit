// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package driver

import (
	"context"

	"xiao-yi-transit/api/driver/v1"
)

type IDriverV1 interface {
	CreateDriver(ctx context.Context, req *v1.CreateDriverReq) (res *v1.CreateDriverRes, err error)
	DeleteDriver(ctx context.Context, req *v1.DeleteDriverReq) (res *v1.DeleteDriverRes, err error)
	GetDriverList(ctx context.Context, req *v1.GetDriverListReq) (res *v1.GetDriverListRes, err error)
	GetDriverDetail(ctx context.Context, req *v1.GetDriverDetailReq) (res *v1.GetDriverDetailRes, err error)
	GetDriverSchedule(ctx context.Context, req *v1.GetDriverScheduleReq) (res *v1.GetDriverScheduleRes, err error)
	UpdateDriver(ctx context.Context, req *v1.UpdateDriverReq) (res *v1.UpdateDriverRes, err error)
}
