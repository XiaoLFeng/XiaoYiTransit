package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/vehicle/v1"
	"xiao-yi-transit/internal/service"
)

// DeleteVehicleInspection 处理删除车辆年检记录的逻辑。
//
// 实现车辆年检记录的删除过程，包括数据验证、信息删除等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 删除车辆年检记录的请求参数，包含年检UUID。
//
// 返回:
//   - res: 删除车辆年检记录的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) DeleteVehicleInspection(ctx context.Context, req *v1.DeleteVehicleInspectionReq) (res *v1.DeleteVehicleInspectionRes, err error) {
	blog.ControllerInfo(ctx, "DeleteVehicleInspection", "删除车辆年检记录请求: %s", req.InspectionUuid)

	// 调用服务删除年检记录
	iVehicle := service.Vehicle()
	errorCode := iVehicle.DeleteVehicleInspection(ctx, req.InspectionUuid)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.DeleteVehicleInspectionRes{
		ResponseDTO: bresult.Success(ctx, "删除车辆年检记录成功"),
	}, nil
}
