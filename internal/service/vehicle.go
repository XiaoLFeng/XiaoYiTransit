// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xiao-yi-transit/internal/model/entity"

	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	IVehicle interface {
		// CreateVehicle 创建新的车辆信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - vehicle: 车辆信息实体，包含需要创建的车辆详细信息。
		//
		// 返回:
		//   - 创建成功的车辆UUID。
		//   - 错误码的指针，表示可能的错误类型。
		CreateVehicle(ctx context.Context, vehicle *entity.Vehicle) (string, *berror.ErrorCode)
		// DeleteVehicle 删除车辆信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - vehicleUuid: 车辆的唯一标识符。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		DeleteVehicle(ctx context.Context, vehicleUuid string) *berror.ErrorCode
		// GetVehicleById 根据车辆UUID获取车辆信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - vehicleUuid: 车辆的唯一标识符。
		//
		// 返回:
		//   - 车辆信息的指针，包含详细车辆数据。
		//   - 错误码的指针，表示错误类型，如车辆不存在或内部错误。
		GetVehicleById(ctx context.Context, vehicleUuid string) (*entity.Vehicle, *berror.ErrorCode)
		// GetVehicleList 获取车辆列表。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - page: 页码，从1开始。
		//   - size: 每页数量。
		//   - plateNumber: 车牌号筛选条件，可为空。
		//   - model: 车型筛选条件，可为空。
		//   - status: 状态筛选条件，可为0（表示不筛选）。
		//
		// 返回:
		//   - 车辆列表。
		//   - 总数量。
		//   - 错误码的指针，表示可能的错误类型。
		GetVehicleList(ctx context.Context, page int, size int, plateNumber string, model string, status int) ([]*entity.Vehicle, int, *berror.ErrorCode)
		// CreateVehicleInspection 创建新的车辆年检记录。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - inspection: 车辆年检记录实体，包含需要创建的年检详细信息。
		//
		// 返回:
		//   - 创建成功的年检记录UUID。
		//   - 错误码的指针，表示可能的错误类型。
		CreateVehicleInspection(ctx context.Context, inspection *entity.VehicleInspection) (string, *berror.ErrorCode)
		// DeleteVehicleInspection 删除车辆年检记录。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - inspectionUuid: 年检记录的唯一标识符。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		DeleteVehicleInspection(ctx context.Context, inspectionUuid string) *berror.ErrorCode
		// GetVehicleInspectionById 根据年检记录UUID获取年检记录信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - inspectionUuid: 年检记录的唯一标识符。
		//
		// 返回:
		//   - 年检记录信息的指针，包含详细年检数据。
		//   - 错误码的指针，表示错误类型，如年检记录不存在或内部错误。
		GetVehicleInspectionById(ctx context.Context, inspectionUuid string) (*entity.VehicleInspection, *berror.ErrorCode)
		// GetVehicleInspectionList 获取车辆年检记录列表。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - page: 页码，从1开始。
		//   - size: 每页数量。
		//   - vehicleUuid: 车辆UUID筛选条件，可为空。
		//   - startDate: 开始日期筛选条件，可为空。
		//   - endDate: 结束日期筛选条件，可为空。
		//
		// 返回:
		//   - 年检记录列表。
		//   - 总数量。
		//   - 错误码的指针，表示可能的错误类型。
		GetVehicleInspectionList(ctx context.Context, page int, size int, vehicleUuid string, startDate *gtime.Time, endDate *gtime.Time) ([]*entity.VehicleInspection, int, *berror.ErrorCode)
		// UpdateVehicleInspection 更新车辆年检记录。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - inspection: 年检记录实体，包含需要更新的年检详细信息。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		UpdateVehicleInspection(ctx context.Context, inspection *entity.VehicleInspection) *berror.ErrorCode
		// CreateVehicleInsurance 创建新的车辆保险记录。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - insurance: 车辆保险记录实体，包含需要创建的保险详细信息。
		//
		// 返回:
		//   - 创建成功的保险记录UUID。
		//   - 错误码的指针，表示可能的错误类型。
		CreateVehicleInsurance(ctx context.Context, insurance *entity.VehicleInsurance) (string, *berror.ErrorCode)
		// DeleteVehicleInsurance 删除车辆保险记录。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - insuranceUuid: 保险记录的唯一标识符。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		DeleteVehicleInsurance(ctx context.Context, insuranceUuid string) *berror.ErrorCode
		// GetVehicleInsuranceById 根据保险记录UUID获取保险记录信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - insuranceUuid: 保险记录的唯一标识符。
		//
		// 返回:
		//   - 保险记录信息的指针，包含详细保险数据。
		//   - 错误码的指针，表示错误类型，如保险记录不存在或内部错误。
		GetVehicleInsuranceById(ctx context.Context, insuranceUuid string) (*entity.VehicleInsurance, *berror.ErrorCode)
		// GetVehicleInsuranceList 获取车辆保险记录列表。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - page: 页码，从1开始。
		//   - size: 每页数量。
		//   - vehicleUuid: 车辆UUID筛选条件，可为空。
		//   - insuranceType: 保险类型筛选条件，可为空。
		//   - startDate: 开始日期筛选条件，可为空。
		//   - endDate: 结束日期筛选条件，可为空。
		//
		// 返回:
		//   - 保险记录列表。
		//   - 总数量。
		//   - 错误码的指针，表示可能的错误类型。
		GetVehicleInsuranceList(ctx context.Context, page int, size int, vehicleUuid string, insuranceType string, startDate *gtime.Time, endDate *gtime.Time) ([]*entity.VehicleInsurance, int, *berror.ErrorCode)
		// UpdateVehicleInsurance 更新车辆保险记录。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - insurance: 保险记录实体，包含需要更新的保险详细信息。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		UpdateVehicleInsurance(ctx context.Context, insurance *entity.VehicleInsurance) *berror.ErrorCode
		// UpdateVehicle 更新车辆信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - vehicle: 车辆信息实体，包含需要更新的车辆详细信息。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		UpdateVehicle(ctx context.Context, vehicle *entity.Vehicle) *berror.ErrorCode
	}
)

var (
	localVehicle IVehicle
)

func Vehicle() IVehicle {
	if localVehicle == nil {
		panic("implement not found for interface IVehicle, forgot register?")
	}
	return localVehicle
}

func RegisterVehicle(i IVehicle) {
	localVehicle = i
}
