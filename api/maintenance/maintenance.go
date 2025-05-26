// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package maintenance

import (
	"context"

	"xiao-yi-transit/api/maintenance/v1"
)

type IMaintenanceV1 interface {
	CreateMaintenance(ctx context.Context, req *v1.CreateMaintenanceReq) (res *v1.CreateMaintenanceRes, err error)
	DeleteMaintenance(ctx context.Context, req *v1.DeleteMaintenanceReq) (res *v1.DeleteMaintenanceRes, err error)
	GetMaintenanceDetail(ctx context.Context, req *v1.GetMaintenanceDetailReq) (res *v1.GetMaintenanceDetailRes, err error)
	GetMaintenanceList(ctx context.Context, req *v1.GetMaintenanceListReq) (res *v1.GetMaintenanceListRes, err error)
	UpdateMaintenance(ctx context.Context, req *v1.UpdateMaintenanceReq) (res *v1.UpdateMaintenanceRes, err error)
}
