package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/vehicle/v1"
	"xiao-yi-transit/internal/service"
)

// DeleteVehicle 处理删除车辆信息的逻辑。
//
// 实现车辆信息的删除过程，包括数据验证、信息删除等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 删除车辆的请求参数，包含车辆UUID。
//
// 返回:
//   - res: 删除车辆的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) DeleteVehicle(ctx context.Context, req *v1.DeleteVehicleReq) (res *v1.DeleteVehicleRes, err error) {
	blog.ControllerInfo(ctx, "DeleteVehicle", "删除车辆请求: %s", req.VehicleUuid)

	// 调用服务删除车辆
	iVehicle := service.Vehicle()
	errorCode := iVehicle.DeleteVehicle(ctx, req.VehicleUuid)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.DeleteVehicleRes{
		ResponseDTO: bresult.Success(ctx, "删除车辆成功"),
	}, nil
}
