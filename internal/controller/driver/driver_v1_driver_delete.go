package driver

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/driver/v1"
	"xiao-yi-transit/internal/service"
)

// DeleteDriver 处理删除司机信息的逻辑。
//
// 实现司机信息的删除过程，包括数据验证、信息删除等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 删除司机的请求参数，包含司机UUID。
//
// 返回:
//   - res: 删除司机的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) DeleteDriver(ctx context.Context, req *v1.DeleteDriverReq) (res *v1.DeleteDriverRes, err error) {
	blog.ControllerInfo(ctx, "DeleteDriver", "删除司机请求: %s", req.DriverUuid)

	// 调用服务删除司机
	iDriver := service.Driver()
	errorCode := iDriver.DeleteDriver(ctx, req.DriverUuid)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.DeleteDriverRes{
		ResponseDTO: bresult.Success(ctx, "删除司机成功"),
	}, nil
}