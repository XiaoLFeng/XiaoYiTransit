package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// GetVehicleById 根据车辆UUID获取车辆信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - vehicleUuid: 车辆的唯一标识符。
//
// 返回:
//   - 车辆信息的指针，包含详细车辆数据。
//   - 错误码的指针，表示错误类型，如车辆不存在或内部错误。
func (s *sVehicle) GetVehicleById(ctx context.Context, vehicleUuid string) (*entity.Vehicle, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetVehicleById", "获取车辆: %s", vehicleUuid)

	// 查询车辆信息
	vehicleModel := dao.Vehicle.Ctx(ctx)
	vehicle := &entity.Vehicle{}
	err := vehicleModel.Where("vehicle_uuid", vehicleUuid).Scan(vehicle)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, &berror.ErrNotFound
		}
		blog.ServiceError(ctx, "GetVehicleById", "获取车辆失败: %s", err.Error())
		return nil, &berror.ErrDatabaseError
	}

	// 返回结果
	return vehicle, nil
}

// GetVehicleList 获取车辆列表。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - page: 页码，从1开始。
//   - size: 每页数量。
//   - plateNumber: 车牌号筛选条件，可为空。
//   - model: 车型筛选条件，可为空。
//   - status: 状态筛选条件，可为0（表示不筛选）。
//
// 返回:
//   - 车辆列表。
//   - 总数量。
//   - 错误码的指针，表示可能的错误类型。
func (s *sVehicle) GetVehicleList(ctx context.Context, page int, size int, plateNumber string, model string, status int) ([]*entity.Vehicle, int, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetVehicleList", "获取车辆列表: page=%d, size=%d, plateNumber=%s, model=%s, status=%d", page, size, plateNumber, model, status)

	// 构建查询条件
	vehicleModel := dao.Vehicle.Ctx(ctx)
	if plateNumber != "" {
		vehicleModel = vehicleModel.Where("plate_number LIKE ?", "%"+plateNumber+"%")
	}
	if model != "" {
		vehicleModel = vehicleModel.Where("model LIKE ?", "%"+model+"%")
	}
	if status > 0 {
		vehicleModel = vehicleModel.Where("status", status)
	}

	// 查询总数
	count, err := vehicleModel.Count()
	if err != nil {
		blog.ServiceError(ctx, "GetVehicleList", "获取车辆总数失败: %s", err.Error())
		return nil, 0, &berror.ErrDatabaseError
	}

	// 查询车辆列表
	var vehicles []*entity.Vehicle
	err = vehicleModel.Page(page, size).Order("created_at DESC").Scan(&vehicles)
	if err != nil {
		blog.ServiceError(ctx, "GetVehicleList", "获取车辆列表失败: %s", err.Error())
		return nil, 0, &berror.ErrDatabaseError
	}

	// 返回结果
	return vehicles, count, nil
}

// GetVehicleSimpleList 获取车辆简易列表。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//
// 返回:
//   - 车辆简易列表。
//   - 错误码的指针，表示可能的错误类型。
func (s *sVehicle) GetVehicleSimpleList(ctx context.Context) ([]*entity.Vehicle, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetVehicleSimpleList", "获取车辆简易列表")

	// 查询车辆列表
	vehicleModel := dao.Vehicle.Ctx(ctx)
	var vehicles []*entity.Vehicle
	err := vehicleModel.Order("plate_number ASC").Scan(&vehicles)
	if err != nil {
		blog.ServiceError(ctx, "GetVehicleSimpleList", "获取车辆简易列表失败: %s", err.Error())
		return nil, &berror.ErrDatabaseError
	}

	// 返回结果
	return vehicles, nil
}
