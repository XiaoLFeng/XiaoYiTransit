package driver

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// UpdateDriver 更新司机信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - driver: 司机信息实体，包含需要更新的司机详细信息。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sDriver) UpdateDriver(ctx context.Context, driver *entity.Driver) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "UpdateDriver", "更新司机 %s 信息", driver.DriverUuid)

	// 检查司机是否存在
	_, errorCode := s.GetDriverById(ctx, driver.DriverUuid)
	if errorCode != nil {
		return errorCode
	}

	// 更新司机信息
	driverModel := dao.Driver.Ctx(ctx)
	_, sqlErr := driverModel.Where(&entity.Driver{DriverUuid: driver.DriverUuid}).Update(driver)
	if sqlErr != nil {
		blog.ServiceError(ctx, "UpdateDriver", "更新司机 %s 信息失败: %s", driver.DriverUuid, sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 返回结果
	return nil
}
