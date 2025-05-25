package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/vehicle/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

// CreateVehicleInspection 处理创建车辆年检记录的逻辑。
//
// 实现车辆年检记录的创建过程，包括数据验证、信息存储等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 创建车辆年检记录的请求参数，包含年检的详细信息。
//
// 返回:
//   - res: 创建车辆年检记录的响应结果，包含创建的年检UUID。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) CreateVehicleInspection(ctx context.Context, req *v1.CreateVehicleInspectionReq) (res *v1.CreateVehicleInspectionRes, err error) {
	blog.ControllerInfo(ctx, "CreateVehicleInspection", "创建车辆年检记录请求: 车辆UUID=%s", req.VehicleUuid)

	// 创建年检记录实体
	inspectionEntity := &entity.VehicleInspection{
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

	// 调用服务创建年检记录
	iVehicle := service.Vehicle()
	inspectionUuid, errorCode := iVehicle.CreateVehicleInspection(ctx, inspectionEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	inspectionUuidPtr := &inspectionUuid
	return &v1.CreateVehicleInspectionRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "创建车辆年检记录成功", inspectionUuidPtr),
	}, nil
}
