package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/vehicle/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

// CreateVehicle 处理创建车辆信息的逻辑。
//
// 实现车辆信息的创建过程，包括数据验证、信息存储等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 创建车辆的请求参数，包含车辆的详细信息。
//
// 返回:
//   - res: 创建车辆的响应结果，包含创建的车辆UUID。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) CreateVehicle(ctx context.Context, req *v1.CreateVehicleReq) (res *v1.CreateVehicleRes, err error) {
	blog.ControllerInfo(ctx, "CreateVehicle", "创建车辆请求: %s", req.PlateNumber)

	// 创建车辆信息
	vehicleEntity := &entity.Vehicle{
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

	// 调用服务创建车辆
	iVehicle := service.Vehicle()
	vehicleUuid, errorCode := iVehicle.CreateVehicle(ctx, vehicleEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	// Since we can't use SuccessHasData with string, we'll use Success and add the UUID to the message
	return &v1.CreateVehicleRes{
		ResponseDTO: bresult.Success(ctx, "创建车辆成功，UUID: " + vehicleUuid),
	}, nil
}
