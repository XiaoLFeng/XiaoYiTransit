// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package vehicle

import (
	"context"

	"xiao-yi-transit/api/vehicle/v1"
)

type IVehicleV1 interface {
	CreateVehicle(ctx context.Context, req *v1.CreateVehicleReq) (res *v1.CreateVehicleRes, err error)
	DeleteVehicle(ctx context.Context, req *v1.DeleteVehicleReq) (res *v1.DeleteVehicleRes, err error)
	CreateVehicleInspection(ctx context.Context, req *v1.CreateVehicleInspectionReq) (res *v1.CreateVehicleInspectionRes, err error)
	UpdateVehicleInspection(ctx context.Context, req *v1.UpdateVehicleInspectionReq) (res *v1.UpdateVehicleInspectionRes, err error)
	DeleteVehicleInspection(ctx context.Context, req *v1.DeleteVehicleInspectionReq) (res *v1.DeleteVehicleInspectionRes, err error)
	GetVehicleInspectionList(ctx context.Context, req *v1.GetVehicleInspectionListReq) (res *v1.GetVehicleInspectionListRes, err error)
	CreateVehicleInsurance(ctx context.Context, req *v1.CreateVehicleInsuranceReq) (res *v1.CreateVehicleInsuranceRes, err error)
	UpdateVehicleInsurance(ctx context.Context, req *v1.UpdateVehicleInsuranceReq) (res *v1.UpdateVehicleInsuranceRes, err error)
	DeleteVehicleInsurance(ctx context.Context, req *v1.DeleteVehicleInsuranceReq) (res *v1.DeleteVehicleInsuranceRes, err error)
	GetVehicleInsuranceList(ctx context.Context, req *v1.GetVehicleInsuranceListReq) (res *v1.GetVehicleInsuranceListRes, err error)
	GetVehicleList(ctx context.Context, req *v1.GetVehicleListReq) (res *v1.GetVehicleListRes, err error)
	GetVehicleSimpleList(ctx context.Context, req *v1.GetVehicleSimpleListReq) (res *v1.GetVehicleSimpleListRes, err error)
	GetVehicleDetail(ctx context.Context, req *v1.GetVehicleDetailReq) (res *v1.GetVehicleDetailRes, err error)
	UpdateVehicle(ctx context.Context, req *v1.UpdateVehicleReq) (res *v1.UpdateVehicleRes, err error)
}
