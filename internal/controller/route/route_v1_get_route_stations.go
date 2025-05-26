package route

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/util/gconv"
	"xiao-yi-transit/api/route/v1"
	"xiao-yi-transit/internal/service"
)

// GetRouteStations 处理获取线路站点列表的逻辑。
//
// 实现线路站点列表的查询过程，包括站点基本信息和顺序等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取线路站点列表的请求参数，包含线路UUID。
//
// 返回:
//   - res: 获取线路站点列表的响应结果，包含站点列表。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) GetRouteStations(ctx context.Context, req *v1.GetRouteStationsReq) (res *v1.GetRouteStationsRes, err error) {
	blog.ControllerInfo(ctx, "GetRouteStations", "获取线路站点列表请求: %s", req.RouteUuid)

	// 获取线路站点列表
	iRoute := service.Route()
	routeStations, errorCode := iRoute.GetRouteStations(ctx, req.RouteUuid)
	if errorCode != nil {
		return nil, errorCode
	}

	// 转换为API响应DTO
	var apiRouteStations v1.RouteStationsDTO
	if err := gconv.Struct(routeStations, &apiRouteStations); err != nil {
		blog.ControllerError(ctx, "GetRouteStations", "数据转换失败: %s", err.Error())
		return nil, &berror.ErrInternalServer
	}

	// 转换站点信息
	if len(routeStations.Stations) > 0 {
		apiRouteStations.Stations = make([]*v1.RouteStationItemDTO, len(routeStations.Stations))
		for i, station := range routeStations.Stations {
			var apiStation v1.RouteStationItemDTO
			if err := gconv.Struct(station, &apiStation); err != nil {
				blog.ControllerError(ctx, "GetRouteStations", "站点数据转换失败: %s", err.Error())
				return nil, &berror.ErrInternalServer
			}
			apiRouteStations.Stations[i] = &apiStation
		}
	}

	// 返回结果
	return &v1.GetRouteStationsRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取线路站点列表成功", &apiRouteStations),
	}, nil
}
