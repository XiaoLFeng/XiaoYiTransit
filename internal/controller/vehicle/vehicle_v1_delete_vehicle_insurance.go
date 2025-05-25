package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/vehicle/v1"
	"xiao-yi-transit/internal/service"
)

// DeleteVehicleInsurance 处理删除车辆保险记录的逻辑。
//
// 实现车辆保险记录的删除过程，包括数据验证、信息删除等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 删除车辆保险记录的请求参数，包含保险UUID。
//
// 返回:
//   - res: 删除车辆保险记录的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) DeleteVehicleInsurance(ctx context.Context, req *v1.DeleteVehicleInsuranceReq) (res *v1.DeleteVehicleInsuranceRes, err error) {
	blog.ControllerInfo(ctx, "DeleteVehicleInsurance", "删除车辆保险记录请求: %s", req.InsuranceUuid)

	// 调用服务删除保险记录
	iVehicle := service.Vehicle()
	errorCode := iVehicle.DeleteVehicleInsurance(ctx, req.InsuranceUuid)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.DeleteVehicleInsuranceRes{
		ResponseDTO: bresult.Success(ctx, "删除车辆保险记录成功"),
	}, nil
}
