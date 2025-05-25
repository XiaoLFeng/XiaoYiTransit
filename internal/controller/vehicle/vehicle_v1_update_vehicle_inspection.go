package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/vehicle/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

// UpdateVehicleInspection 处理更新车辆年检记录的逻辑。
//
// 实现车辆年检记录的更新过程，包括数据验证、信息更新等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 更新车辆年检记录的请求参数，包含年检的详细信息。
//
// 返回:
//   - res: 更新车辆年检记录的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) UpdateVehicleInspection(ctx context.Context, req *v1.UpdateVehicleInspectionReq) (res *v1.UpdateVehicleInspectionRes, err error) {
	blog.ControllerInfo(ctx, "UpdateVehicleInspection", "更新车辆年检记录请求: %s", req.InspectionUuid)

	// 创建年检记录实体
	inspectionEntity := &entity.VehicleInspection{
		InspectionUuid:   req.InspectionUuid,
		VehicleUuid:      req.VehicleUuid,
		InspectionDate:   req.InspectionDate,
		ExpiryDate:       req.ExpiryDate,
		InspectionResult: req.InspectionResult,
		InspectionAgency: req.InspectionAgency,
		Inspector:        req.Inspector,
		CertificateNo:    req.CertificateNo,
		Cost:             req.Cost,
		Notes:            req.Notes,
	}

	// 调用服务更新年检记录
	iVehicle := service.Vehicle()
	errorCode := iVehicle.UpdateVehicleInspection(ctx, inspectionEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.UpdateVehicleInspectionRes{
		ResponseDTO: bresult.Success(ctx, "更新车辆年检记录成功"),
	}, nil
}
