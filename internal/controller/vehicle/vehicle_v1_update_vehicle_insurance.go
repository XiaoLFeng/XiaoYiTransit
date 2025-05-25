package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/vehicle/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

// UpdateVehicleInsurance 处理更新车辆保险记录的逻辑。
//
// 实现车辆保险记录的更新过程，包括数据验证、信息更新等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 更新车辆保险记录的请求参数，包含保险的详细信息。
//
// 返回:
//   - res: 更新车辆保险记录的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) UpdateVehicleInsurance(ctx context.Context, req *v1.UpdateVehicleInsuranceReq) (res *v1.UpdateVehicleInsuranceRes, err error) {
	blog.ControllerInfo(ctx, "UpdateVehicleInsurance", "更新车辆保险记录请求: %s", req.InsuranceUuid)

	// 创建保险记录实体
	insuranceEntity := &entity.VehicleInsurance{
		InsuranceUuid:  req.InsuranceUuid,
		VehicleUuid:    req.VehicleUuid,
		InsuranceType:  req.InsuranceType,
		PolicyNumber:   req.PolicyNumber,
		Insurer:        req.Insurer,
		StartDate:      req.StartDate,
		ExpiryDate:     req.ExpiryDate,
		CoverageAmount: req.CoverageAmount,
		Premium:        req.Premium,
		PaymentDate:    req.PaymentDate,
		PaymentMethod:  req.PaymentMethod,
		ContactPerson:  req.ContactPerson,
		ContactPhone:   req.ContactPhone,
		Notes:          req.Notes,
	}

	// 调用服务更新保险记录
	iVehicle := service.Vehicle()
	errorCode := iVehicle.UpdateVehicleInsurance(ctx, insuranceEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.UpdateVehicleInsuranceRes{
		ResponseDTO: bresult.Success(ctx, "更新车辆保险记录成功"),
	}, nil
}
