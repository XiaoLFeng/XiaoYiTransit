package route

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/route/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

// AddRouteStation 处理向线路添加站点的逻辑。
//
// 实现向线路添加站点的过程，包括数据验证、信息存储等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 添加站点的请求参数，包含线路UUID、站点UUID和站点顺序等信息。
//
// 返回:
//   - res: 添加站点的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) AddRouteStation(ctx context.Context, req *v1.AddRouteStationReq) (res *v1.AddRouteStationRes, err error) {
	blog.ControllerInfo(ctx, "AddRouteStation", "添加线路站点请求: 线路=%s, 站点=%s", req.RouteUuid, req.StationUuid)

	// 创建线路站点信息实体
	routeStationEntity := &entity.RouteStation{
		RouteUuid:         req.RouteUuid,
		StationUuid:       req.StationUuid,
		Sequence:          req.Sequence,
		DistanceFromStart: req.DistanceFromStart,
		EstimatedTime:     req.EstimatedTime,
	}

	// 调用服务添加线路站点
	iRoute := service.Route()
	routeStationUuid, errorCode := iRoute.AddRouteStation(ctx, routeStationEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.AddRouteStationRes{
		ResponseDTO: bresult.Success(ctx, "添加线路站点成功，UUID: " + routeStationUuid),
	}, nil
}
