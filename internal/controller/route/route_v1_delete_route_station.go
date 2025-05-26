package route

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/route/v1"
	"xiao-yi-transit/internal/service"
)

// DeleteRouteStation 处理删除线路站点的逻辑。
//
// 实现线路站点的删除过程，包括数据验证、信息删除等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 删除站点的请求参数，包含线路站点UUID。
//
// 返回:
//   - res: 删除站点的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) DeleteRouteStation(ctx context.Context, req *v1.DeleteRouteStationReq) (res *v1.DeleteRouteStationRes, err error) {
	blog.ControllerInfo(ctx, "DeleteRouteStation", "删除线路站点请求: %s", req.RouteStationUuid)

	// 调用服务删除线路站点
	iRoute := service.Route()
	errorCode := iRoute.DeleteRouteStation(ctx, req.RouteStationUuid)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.DeleteRouteStationRes{
		ResponseDTO: bresult.Success(ctx, "删除线路站点成功"),
	}, nil
}
