package driver

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/util/gconv"
	"xiao-yi-transit/api/driver/v1"
	"xiao-yi-transit/internal/model/dto/driver"
	"xiao-yi-transit/internal/service"
)

// GetDriverSchedule 处理获取司机排班信息的逻辑。
//
// 实现司机排班信息的查询过程，包括日期范围筛选等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取司机排班的请求参数，包含司机UUID和日期范围。
//
// 返回:
//   - res: 获取司机排班的响应结果，包含司机排班信息。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) GetDriverSchedule(ctx context.Context, req *v1.GetDriverScheduleReq) (res *v1.GetDriverScheduleRes, err error) {
	blog.ControllerInfo(ctx, "GetDriverSchedule", "获取司机排班请求: %s", req.DriverUuid)

	// 调用服务获取司机排班
	iDriver := service.Driver()
	driverEntity, shifts, errorCode := iDriver.GetDriverSchedule(ctx, req.DriverUuid, req.StartDate, req.EndDate)
	if errorCode != nil {
		return nil, errorCode
	}

	// 转换数据
	var shiftItems []*driver.DriverShiftItemDTO
	for _, shift := range shifts {
		var item driver.DriverShiftItemDTO
		if err := gconv.Struct(shift, &item); err != nil {
			blog.ControllerError(ctx, "GetDriverSchedule", "数据转换失败: %s", err.Error())
			return nil, &berror.ErrInternalServer
		}
		shiftItems = append(shiftItems, &item)
	}

	// 构建返回数据
	schedule := &driver.DriverScheduleDTO{
		DriverUuid: driverEntity.DriverUuid,
		DriverName: driverEntity.Name,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Shifts:     shiftItems,
	}

	// 返回结果
	return &v1.GetDriverScheduleRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取司机排班成功", schedule),
	}, nil
}
