package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/vehicle/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

// UpdateVehicle 处理更新车辆信息的逻辑。
//
// 实现车辆信息的更新过程，包括数据验证、信息更新等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 更新车辆的请求参数，包含车辆的详细信息。
//
// 返回:
//   - res: 更新车辆的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) UpdateVehicle(ctx context.Context, req *v1.UpdateVehicleReq) (res *v1.UpdateVehicleRes, err error) {
	blog.ControllerInfo(ctx, "UpdateVehicle", "更新车辆请求: %s", req.VehicleUuid)

	// 创建车辆信息实体
	vehicleEntity := &entity.Vehicle{
		VehicleUuid:         req.VehicleUuid,
		PlateNumber:         req.PlateNumber,
		Model:               req.Model,
		PurchaseDate:        req.PurchaseDate,
		Status:              req.Status,
		Seats:               req.Seats,
		FuelType:            req.FuelType,
		Mileage:             req.Mileage,
		LastMaintenanceDate: req.LastMaintenanceDate,
		NextInspectionDate:  req.NextInspectionDate,
		InsuranceExpiryDate: req.InsuranceExpiryDate,
		Notes:               req.Notes,
	}

	// 调用服务更新车辆
	iVehicle := service.Vehicle()
	errorCode := iVehicle.UpdateVehicle(ctx, vehicleEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.UpdateVehicleRes{
		ResponseDTO: bresult.Success(ctx, "更新车辆成功"),
	}, nil
}
