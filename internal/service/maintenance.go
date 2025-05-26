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
	IMaintenance interface {
		// CreateMaintenance 创建新的维修记录。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - maintenance: 维修记录实体，包含需要创建的维修详细信息。
		//
		// 返回:
		//   - 创建成功的维修记录UUID。
		//   - 错误码的指针，表示可能的错误类型。
		CreateMaintenance(ctx context.Context, maintenance *entity.Maintenance) (string, *berror.ErrorCode)

		// DeleteMaintenance 删除维修记录。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - maintenanceUuid: 维修记录的唯一标识符。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		DeleteMaintenance(ctx context.Context, maintenanceUuid string) *berror.ErrorCode

		// GetMaintenanceById 根据维修记录UUID获取维修记录信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - maintenanceUuid: 维修记录的唯一标识符。
		//
		// 返回:
		//   - 维修记录信息的指针，包含详细维修数据。
		//   - 错误码的指针，表示错误类型，如维修记录不存在或内部错误。
		GetMaintenanceById(ctx context.Context, maintenanceUuid string) (*entity.Maintenance, *berror.ErrorCode)

		// GetMaintenanceList 获取维修记录列表。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - page: 页码，从1开始。
		//   - size: 每页数量。
		//   - vehicleUuid: 车辆UUID筛选条件，可为空。
		//   - maintenanceType: 维修类型筛选条件，可为0（表示不筛选）。
		//   - status: 状态筛选条件，可为0（表示不筛选）。
		//   - startDate: 开始日期筛选条件，可为空。
		//   - endDate: 结束日期筛选条件，可为空。
		//
		// 返回:
		//   - 维修记录列表。
		//   - 总数量。
		//   - 错误码的指针，表示可能的错误类型。
		GetMaintenanceList(ctx context.Context, page int, size int, vehicleUuid string, maintenanceType int, status int, startDate *gtime.Time, endDate *gtime.Time) ([]*entity.Maintenance, int, *berror.ErrorCode)

		// UpdateMaintenance 更新维修记录。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - maintenance: 维修记录实体，包含需要更新的维修详细信息。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		UpdateMaintenance(ctx context.Context, maintenance *entity.Maintenance) *berror.ErrorCode
	}
)

var (
	localMaintenance IMaintenance
)

func Maintenance() IMaintenance {
	if localMaintenance == nil {
		panic("implement not found for interface IMaintenance, forgot register?")
	}
	return localMaintenance
}

func RegisterMaintenance(i IMaintenance) {
	localMaintenance = i
}
