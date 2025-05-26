package maintenance

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"xiao-yi-transit/internal/dao"
)

// DeleteMaintenance 删除维修记录。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - maintenanceUuid: 维修记录的唯一标识符。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sMaintenance) DeleteMaintenance(ctx context.Context, maintenanceUuid string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "DeleteMaintenance", "删除维修记录: %s", maintenanceUuid)

	// 删除维修记录
	maintenanceModel := dao.Maintenance.Ctx(ctx)
	_, sqlErr := maintenanceModel.Where("maintenance_uuid", maintenanceUuid).Delete()
	if sqlErr != nil {
		blog.ServiceError(ctx, "DeleteMaintenance", "删除维修记录失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 返回结果
	return nil
}
