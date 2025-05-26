package route

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/route/v1"
	"xiao-yi-transit/internal/service"
)

// DeleteRoute 处理删除线路信息的逻辑。
//
// 实现线路信息的删除过程，包括数据验证、信息删除等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 删除线路的请求参数，包含线路UUID。
//
// 返回:
//   - res: 删除线路的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) DeleteRoute(ctx context.Context, req *v1.DeleteRouteReq) (res *v1.DeleteRouteRes, err error) {
	blog.ControllerInfo(ctx, "DeleteRoute", "删除线路请求: %s", req.RouteUuid)

	// 调用服务删除线路
	iRoute := service.Route()
	errorCode := iRoute.DeleteRoute(ctx, req.RouteUuid)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.DeleteRouteRes{
		ResponseDTO: bresult.Success(ctx, "删除线路成功"),
	}, nil
}
