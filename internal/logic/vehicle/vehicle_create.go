package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/google/uuid"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// CreateVehicle 创建新的车辆信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - vehicle: 车辆信息实体，包含需要创建的车辆详细信息。
//
// 返回:
//   - 创建成功的车辆UUID。
//   - 错误码的指针，表示可能的错误类型。
func (s *sVehicle) CreateVehicle(ctx context.Context, vehicle *entity.Vehicle) (string, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "CreateVehicle", "创建车辆: %s", vehicle.PlateNumber)

	// 生成UUID
	vehicle.VehicleUuid = uuid.New().String()

	// 创建车辆信息
	vehicleModel := dao.Vehicle.Ctx(ctx)
	_, sqlErr := vehicleModel.Insert(vehicle)
	if sqlErr != nil {
		blog.ServiceError(ctx, "CreateVehicle", "创建车辆失败: %s", sqlErr.Error())
		return "", &berror.ErrDatabaseError
	}

	// 返回结果
	return vehicle.VehicleUuid, nil
}
