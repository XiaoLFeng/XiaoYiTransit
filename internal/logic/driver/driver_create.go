package driver

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/google/uuid"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// CreateDriver 创建新的司机信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - driver: 司机信息实体，包含需要创建的司机详细信息。
//
// 返回:
//   - 创建成功的司机UUID。
//   - 错误码的指针，表示可能的错误类型。
func (s *sDriver) CreateDriver(ctx context.Context, driver *entity.Driver) (string, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "CreateDriver", "创建司机: %s", driver.Name)

	// 生成UUID
	driver.DriverUuid = uuid.New().String()

	// 创建司机信息
	driverModel := dao.Driver.Ctx(ctx)
	_, sqlErr := driverModel.OmitEmpty().Insert(driver)
	if sqlErr != nil {
		blog.ServiceError(ctx, "CreateDriver", "创建司机失败: %s", sqlErr.Error())
		return "", &berror.ErrDatabaseError
	}

	// 返回结果
	return driver.DriverUuid, nil
}
