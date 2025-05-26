package route

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/route/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

// UpdateRouteStation 处理更新线路站点的逻辑。
//
// 实现线路站点信息的更新过程，包括数据验证、信息更新等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 更新站点的请求参数，包含线路站点UUID和站点顺序等信息。
//
// 返回:
//   - res: 更新站点的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) UpdateRouteStation(ctx context.Context, req *v1.UpdateRouteStationReq) (res *v1.UpdateRouteStationRes, err error) {
	blog.ControllerInfo(ctx, "UpdateRouteStation", "更新线路站点请求: %s", req.RouteStationUuid)

	// 创建线路站点信息实体
	routeStationEntity := &entity.RouteStation{
		RouteStationUuid:  req.RouteStationUuid,
		Sequence:          req.Sequence,
		DistanceFromStart: req.DistanceFromStart,
		EstimatedTime:     req.EstimatedTime,
	}

	// 调用服务更新线路站点
	iRoute := service.Route()
	errorCode := iRoute.UpdateRouteStation(ctx, routeStationEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.UpdateRouteStationRes{
		ResponseDTO: bresult.Success(ctx, "更新线路站点成功"),
	}, nil
}
