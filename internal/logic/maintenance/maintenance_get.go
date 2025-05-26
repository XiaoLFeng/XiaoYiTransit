package maintenance

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/os/gtime"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// GetMaintenanceById 根据维修记录UUID获取维修记录信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - maintenanceUuid: 维修记录的唯一标识符。
//
// 返回:
//   - 维修记录信息的指针，包含详细维修数据。
//   - 错误码的指针，表示错误类型，如维修记录不存在或内部错误。
func (s *sMaintenance) GetMaintenanceById(ctx context.Context, maintenanceUuid string) (*entity.Maintenance, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetMaintenanceById", "获取维修记录: %s", maintenanceUuid)

	// 查询维修记录
	maintenanceModel := dao.Maintenance.Ctx(ctx)
	maintenance := &entity.Maintenance{}
	err := maintenanceModel.Where("maintenance_uuid", maintenanceUuid).Scan(maintenance)
	if err != nil {
		blog.ServiceError(ctx, "GetMaintenanceById", "获取维修记录失败: %s", err.Error())
		return nil, &berror.ErrDatabaseError
	}

	// 检查是否存在
	if maintenance.MaintenanceUuid == "" {
		blog.ServiceInfo(ctx, "GetMaintenanceById", "维修记录不存在: %s", maintenanceUuid)
		return nil, &berror.ErrNotFound
	}

	// 返回结果
	return maintenance, nil
}

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
func (s *sMaintenance) GetMaintenanceList(ctx context.Context, page int, size int, vehicleUuid string, maintenanceType int, status int, startDate *gtime.Time, endDate *gtime.Time) ([]*entity.Maintenance, int, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetMaintenanceList", "获取维修记录列表: page=%d, size=%d, vehicleUuid=%s, maintenanceType=%d, status=%d", page, size, vehicleUuid, maintenanceType, status)

	// 构建查询条件
	maintenanceModel := dao.Maintenance.Ctx(ctx)
	model := maintenanceModel.Clone()

	// 添加筛选条件
	if vehicleUuid != "" {
		model = model.Where("vehicle_uuid", vehicleUuid)
	}
	if maintenanceType > 0 {
		model = model.Where("maintenance_type", maintenanceType)
	}
	if status > 0 {
		model = model.Where("status", status)
	}
	if startDate != nil {
		model = model.WhereGTE("maintenance_date", startDate)
	}
	if endDate != nil {
		model = model.WhereLTE("maintenance_date", endDate)
	}

	// 计算总数
	count, err := model.Count()
	if err != nil {
		blog.ServiceError(ctx, "GetMaintenanceList", "获取维修记录总数失败: %s", err.Error())
		return nil, 0, &berror.ErrDatabaseError
	}

	// 分页查询
	var maintenances []*entity.Maintenance
	err = model.Page(page, size).Order("maintenance_date DESC").Scan(&maintenances)
	if err != nil {
		blog.ServiceError(ctx, "GetMaintenanceList", "获取维修记录列表失败: %s", err.Error())
		return nil, 0, &berror.ErrDatabaseError
	}

	// 返回结果
	return maintenances, count, nil
}
