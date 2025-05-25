package driver

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/driver/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

// UpdateDriver 处理更新司机信息的逻辑。
//
// 实现司机信息的更新过程，包括数据验证、信息更新等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 更新司机的请求参数，包含司机的详细信息。
//
// 返回:
//   - res: 更新司机的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) UpdateDriver(ctx context.Context, req *v1.UpdateDriverReq) (res *v1.UpdateDriverRes, err error) {
	blog.ControllerInfo(ctx, "UpdateDriver", "更新司机请求: %s", req.DriverUuid)

	// 创建司机信息实体
	driverEntity := &entity.Driver{
		DriverUuid:        req.DriverUuid,
		EmployeeId:        req.EmployeeId,
		Name:              req.Name,
		Gender:            req.Gender,
		IdCard:            req.IdCard,
		Phone:             req.Phone,
		EmergencyContact:  req.EmergencyContact,
		EmergencyPhone:    req.EmergencyPhone,
		LicenseNumber:     req.LicenseNumber,
		LicenseType:       req.LicenseType,
		LicenseExpiryDate: req.LicenseExpiryDate,
		EntryDate:         req.EntryDate,
		Status:            req.Status,
		Address:           req.Address,
		Notes:             req.Notes,
	}

	// 调用服务更新司机
	iDriver := service.Driver()
	errorCode := iDriver.UpdateDriver(ctx, driverEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.UpdateDriverRes{
		ResponseDTO: bresult.Success(ctx, "更新司机成功"),
	}, nil
}
