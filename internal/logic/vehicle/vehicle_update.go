package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// UpdateVehicle 更新车辆信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - vehicle: 车辆信息实体，包含需要更新的车辆详细信息。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sVehicle) UpdateVehicle(ctx context.Context, vehicle *entity.Vehicle) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "UpdateVehicle", "更新车辆: %s", vehicle.VehicleUuid)

	// 检查车辆是否存在
	existingVehicle, err := s.GetVehicleById(ctx, vehicle.VehicleUuid)
	if err != nil {
		return err
	}
	if existingVehicle == nil {
		return &berror.ErrNotFound
	}

	// 更新车辆信息
	vehicleModel := dao.Vehicle.Ctx(ctx)
	_, sqlErr := vehicleModel.Where("vehicle_uuid", vehicle.VehicleUuid).Update(vehicle)
	if sqlErr != nil {
		blog.ServiceError(ctx, "UpdateVehicle", "更新车辆失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 返回结果
	return nil
}
