package driver

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/util/gconv"
	"xiao-yi-transit/api/driver/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

// CreateDriver 处理创建司机信息的逻辑。
//
// 实现司机信息的创建过程，包括数据验证、信息存储等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 创建司机的请求参数，包含司机的详细信息。
//
// 返回:
//   - res: 创建司机的响应结果，包含创建的司机UUID。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) CreateDriver(ctx context.Context, req *v1.CreateDriverReq) (res *v1.CreateDriverRes, err error) {
	blog.ControllerInfo(ctx, "CreateDriver", "创建司机请求: %s", req.Name)

	// 创建司机信息
	driverEntity := &entity.Driver{
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

	// 调用服务创建司机
	iDriver := service.Driver()
	driverUuid, errorCode := iDriver.CreateDriver(ctx, driverEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.CreateDriverRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "创建司机成功", driverUuid),
	}, nil
}