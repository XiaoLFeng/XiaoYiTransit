package driver

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/os/gtime"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// DeleteDriver 删除司机信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - driverUuid: 司机的唯一标识符。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sDriver) DeleteDriver(ctx context.Context, driverUuid string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "DeleteDriver", "删除司机 %s 信息", driverUuid)

	// 检查司机是否存在
	_, errorCode := s.GetDriverById(ctx, driverUuid)
	if errorCode != nil {
		return errorCode
	}

	// 软删除司机信息（设置删除时间）
	driverModel := dao.Driver.Ctx(ctx)
	_, sqlErr := driverModel.Where(&entity.Driver{DriverUuid: driverUuid}).Update(&entity.Driver{
		DeletedAt: gtime.Now(),
	})
	if sqlErr != nil {
		blog.ServiceError(ctx, "DeleteDriver", "删除司机 %s 信息失败: %s", driverUuid, sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 返回结果
	return nil
}
