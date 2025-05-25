package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/util/gconv"
	"xiao-yi-transit/api/vehicle/v1"
	"xiao-yi-transit/internal/model/dto/vehicle"
	"xiao-yi-transit/internal/service"
)

// GetVehicleDetail 处理获取车辆详情的逻辑。
//
// 实现车辆详情的查询过程。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取车辆详情的请求参数，包含车辆UUID。
//
// 返回:
//   - res: 获取车辆详情的响应结果，包含车辆详细信息。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) GetVehicleDetail(ctx context.Context, req *v1.GetVehicleDetailReq) (res *v1.GetVehicleDetailRes, err error) {
	blog.ControllerInfo(ctx, "GetVehicleDetail", "获取车辆详情请求: %s", req.VehicleUuid)

	// 获取车辆详情
	iVehicle := service.Vehicle()
	vehicleEntity, errorCode := iVehicle.GetVehicleById(ctx, req.VehicleUuid)
	if errorCode != nil {
		return nil, errorCode
	}

	// 转换数据
	var vehicleDetail vehicle.VehicleDetailItemDTO
	if err := gconv.Struct(vehicleEntity, &vehicleDetail); err != nil {
		blog.ControllerError(ctx, "GetVehicleDetail", "数据转换失败: %s", err.Error())
		return nil, &berror.ErrInternalServer
	}

	// 返回结果
	return &v1.GetVehicleDetailRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取车辆详情成功", &vehicleDetail),
	}, nil
}
