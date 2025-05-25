// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package driver

import (
	"context"

	"xiao-yi-transit/api/driver/v1"
)

type IDriverV1 interface {
	// Driver information entry
	CreateDriver(ctx context.Context, req *v1.CreateDriverReq) (res *v1.CreateDriverRes, err error)
	
	// Driver information query
	GetDriverList(ctx context.Context, req *v1.GetDriverListReq) (res *v1.GetDriverListRes, err error)
	GetDriverDetail(ctx context.Context, req *v1.GetDriverDetailReq) (res *v1.GetDriverDetailRes, err error)
	
	// Driver information modification
	UpdateDriver(ctx context.Context, req *v1.UpdateDriverReq) (res *v1.UpdateDriverRes, err error)
	
	// Driver information deletion
	DeleteDriver(ctx context.Context, req *v1.DeleteDriverReq) (res *v1.DeleteDriverRes, err error)
	
	// Driver scheduling view
	GetDriverSchedule(ctx context.Context, req *v1.GetDriverScheduleReq) (res *v1.GetDriverScheduleRes, err error)
}