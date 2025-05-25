// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xiao-yi-transit/internal/model/entity"

	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	IDriver interface {
		// CreateDriver 创建新的司机信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - driver: 司机信息实体，包含需要创建的司机详细信息。
		//
		// 返回:
		//   - 创建成功的司机UUID。
		//   - 错误码的指针，表示可能的错误类型。
		CreateDriver(ctx context.Context, driver *entity.Driver) (string, *berror.ErrorCode)
		// DeleteDriver 删除司机信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - driverUuid: 司机的唯一标识符。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		DeleteDriver(ctx context.Context, driverUuid string) *berror.ErrorCode
		// GetDriverById 根据司机UUID获取司机信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - driverUuid: 司机的唯一标识符。
		//
		// 返回:
		//   - 司机信息的指针，包含详细司机数据。
		//   - 错误码的指针，表示错误类型，如司机不存在或内部错误。
		GetDriverById(ctx context.Context, driverUuid string) (*entity.Driver, *berror.ErrorCode)
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
		GetDriverList(ctx context.Context, page int, size int, employeeId string, name string, status int) ([]*entity.Driver, int, *berror.ErrorCode)
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
		GetDriverSchedule(ctx context.Context, driverUuid string, startDate *gtime.Time, endDate *gtime.Time) (*entity.Driver, []*entity.DriverShift, *berror.ErrorCode)
		// UpdateDriver 更新司机信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - driver: 司机信息实体，包含需要更新的司机详细信息。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		UpdateDriver(ctx context.Context, driver *entity.Driver) *berror.ErrorCode
	}
)

var (
	localDriver IDriver
)

func Driver() IDriver {
	if localDriver == nil {
		panic("implement not found for interface IDriver, forgot register?")
	}
	return localDriver
}

func RegisterDriver(i IDriver) {
	localDriver = i
}
