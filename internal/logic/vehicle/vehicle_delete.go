package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"xiao-yi-transit/internal/dao"
)

// DeleteVehicle 删除车辆信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - vehicleUuid: 车辆的唯一标识符。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sVehicle) DeleteVehicle(ctx context.Context, vehicleUuid string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "DeleteVehicle", "删除车辆: %s", vehicleUuid)

	// 检查车辆是否存在
	vehicle, err := s.GetVehicleById(ctx, vehicleUuid)
	if err != nil {
		return err
	}
	if vehicle == nil {
		return &berror.ErrNotFound
	}

	// 删除车辆信息（软删除）
	vehicleModel := dao.Vehicle.Ctx(ctx)
	_, sqlErr := vehicleModel.Where("vehicle_uuid", vehicleUuid).Delete()
	if sqlErr != nil {
		blog.ServiceError(ctx, "DeleteVehicle", "删除车辆失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 返回结果
	return nil
}
