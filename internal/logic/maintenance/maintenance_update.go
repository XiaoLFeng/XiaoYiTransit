package maintenance

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// UpdateMaintenance 更新维修记录。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - maintenance: 维修记录实体，包含需要更新的维修详细信息。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sMaintenance) UpdateMaintenance(ctx context.Context, maintenance *entity.Maintenance) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "UpdateMaintenance", "更新维修记录: %s", maintenance.MaintenanceUuid)

	// 检查维修记录是否存在
	existingMaintenance, errorCode := s.GetMaintenanceById(ctx, maintenance.MaintenanceUuid)
	if errorCode != nil {
		blog.ServiceError(ctx, "UpdateMaintenance", "维修记录不存在: %s", maintenance.MaintenanceUuid)
		return errorCode
	}

	// 保留不需要更新的字段
	maintenance.CreatedAt = existingMaintenance.CreatedAt
	maintenance.CreatedByUuid = existingMaintenance.CreatedByUuid

	// 更新维修记录
	maintenanceModel := dao.Maintenance.Ctx(ctx)
	_, sqlErr := maintenanceModel.Where("maintenance_uuid", maintenance.MaintenanceUuid).Update(maintenance)
	if sqlErr != nil {
		blog.ServiceError(ctx, "UpdateMaintenance", "更新维修记录失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 返回结果
	return nil
}
