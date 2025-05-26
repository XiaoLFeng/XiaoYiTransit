package maintenance

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/google/uuid"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// CreateMaintenance 创建新的维修记录。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - maintenance: 维修记录实体，包含需要创建的维修详细信息。
//
// 返回:
//   - 创建成功的维修记录UUID。
//   - 错误码的指针，表示可能的错误类型。
func (s *sMaintenance) CreateMaintenance(ctx context.Context, maintenance *entity.Maintenance) (string, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "CreateMaintenance", "创建维修记录: 车辆UUID=%s, 维修类型=%d", maintenance.VehicleUuid, maintenance.MaintenanceType)

	// 生成UUID
	maintenance.MaintenanceUuid = uuid.New().String()

	// 创建维修记录
	maintenanceModel := dao.Maintenance.Ctx(ctx)
	_, sqlErr := maintenanceModel.Insert(maintenance)
	if sqlErr != nil {
		blog.ServiceError(ctx, "CreateMaintenance", "创建维修记录失败: %s", sqlErr.Error())
		return "", &berror.ErrDatabaseError
	}

	// 返回结果
	return maintenance.MaintenanceUuid, nil
}
