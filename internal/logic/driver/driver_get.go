package driver

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/do"
	"xiao-yi-transit/internal/model/entity"
)

// GetDriverById 根据司机UUID获取司机信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - driverUuid: 司机的唯一标识符。
//
// 返回:
//   - 司机信息的指针，包含详细司机数据。
//   - 错误码的指针，表示错误类型，如司机不存在或内部错误。
func (s *sDriver) GetDriverById(ctx context.Context, driverUuid string) (*entity.Driver, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetDriverById", "获取司机 %s 信息", driverUuid)

	// 查询司机信息
	var getDriver *entity.Driver
	sqlErr := dao.Driver.Ctx(ctx).Where(&do.Driver{DriverUuid: driverUuid}).Scan(&getDriver)
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetDriverById", "获取司机 %s 信息失败: %s", driverUuid, sqlErr.Error())
		return nil, &berror.ErrDatabaseError
	}
	if getDriver == nil {
		blog.ServiceError(ctx, "GetDriverById", "获取司机 %s 信息失败: 司机不存在", driverUuid)
		return nil, &berror.ErrNotFound
	}

	// 返回结果
	return getDriver, nil
}

// GetDriverList 获取司机列表。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - page: 页码，从1开始。
//   - size: 每页数量。
//   - employeeId: 工号筛选条件，可为空。
//   - name: 姓名筛选条件，可为空。
//   - status: 状态筛选条件，可为0（表示不筛选）。
//
// 返回:
//   - 司机列表。
//   - 总数量。
//   - 错误码的指针，表示可能的错误类型。
func (s *sDriver) GetDriverList(ctx context.Context, page, size int, employeeId, name string, status int) ([]*entity.Driver, int, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetDriverList", "获取司机列表: page=%d, size=%d", page, size)

	// 添加筛选条件
	whereBuilder := dao.Driver.Ctx(ctx).Builder()
	if employeeId != "" {
		whereBuilder = whereBuilder.Where("employee_id = ?", employeeId)
	}
	if name != "" {
		whereBuilder = whereBuilder.Where("name LIKE ?", "%"+name+"%")
	}
	if status != 0 {
		whereBuilder = whereBuilder.Where("status = ?", status)
	}

	// 查询总数
	count, sqlErr := dao.Driver.Ctx(ctx).Count()
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetDriverList", "获取司机总数失败: %s", sqlErr.Error())
		return nil, 0, &berror.ErrDatabaseError
	}

	// 查询列表
	var drivers []*entity.Driver
	sqlErr = dao.Driver.Ctx(ctx).Page(page, size).Order("created_at DESC").Scan(&drivers)
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetDriverList", "获取司机列表失败: %s", sqlErr.Error())
		return nil, 0, &berror.ErrDatabaseError
	}

	// 返回结果
	return drivers, count, nil
}
