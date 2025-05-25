package driver

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// GetDriverSchedule 获取司机排班信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - driverUuid: 司机的唯一标识符。
//   - startDate: 开始日期。
//   - endDate: 结束日期。
//
// 返回:
//   - 司机信息。
//   - 排班列表。
//   - 错误码的指针，表示可能的错误类型。
func (s *sDriver) GetDriverSchedule(ctx context.Context, driverUuid string, startDate, endDate *gtime.Time) (*entity.Driver, []*entity.DriverShift, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetDriverSchedule", "获取司机 %s 排班信息", driverUuid)

	// 获取司机信息
	driver, errorCode := s.GetDriverById(ctx, driverUuid)
	if errorCode != nil {
		return nil, nil, errorCode
	}

	// 设置默认日期范围
	now := gtime.Now()
	if startDate == nil {
		startDate = now
	}
	if endDate == nil {
		// 默认查询7天
		endDate = now.Add(7 * 24 * time.Hour)
	}

	// 查询排班信息
	shiftModel := dao.DriverShift.Ctx(ctx)
	var shifts []*entity.DriverShift
	sqlErr := shiftModel.Where(&entity.DriverShift{
		DriverUuid: driverUuid,
	}).Where("shift_date BETWEEN ? AND ?", startDate, endDate).
		Order("shift_date ASC, start_time ASC").
		Scan(&shifts)
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetDriverSchedule", "获取司机 %s 排班信息失败: %s", driverUuid, sqlErr.Error())
		return nil, nil, &berror.ErrDatabaseError
	}

	// 返回结果
	return driver, shifts, nil
}
