package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/vehicle/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

// CreateVehicleInsurance 处理创建车辆保险记录的逻辑。
//
// 实现车辆保险记录的创建过程，包括数据验证、信息存储等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 创建车辆保险记录的请求参数，包含保险的详细信息。
//
// 返回:
//   - res: 创建车辆保险记录的响应结果，包含创建的保险UUID。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) CreateVehicleInsurance(ctx context.Context, req *v1.CreateVehicleInsuranceReq) (res *v1.CreateVehicleInsuranceRes, err error) {
	blog.ControllerInfo(ctx, "CreateVehicleInsurance", "创建车辆保险记录请求: 车辆UUID=%s", req.VehicleUuid)

	// 创建保险记录实体
	insuranceEntity := &entity.VehicleInsurance{
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

	// 调用服务创建保险记录
	iVehicle := service.Vehicle()
	insuranceUuid, errorCode := iVehicle.CreateVehicleInsurance(ctx, insuranceEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	insuranceUuidPtr := &insuranceUuid
	return &v1.CreateVehicleInsuranceRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "创建车辆保险记录成功", insuranceUuidPtr),
	}, nil
}
