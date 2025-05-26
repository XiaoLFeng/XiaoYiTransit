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

// GetRouteDetail 处理获取线路详情的逻辑。
//
// 实现线路详情的查询过程，包括线路基本信息和站点信息。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取线路详情的请求参数，包含线路UUID。
//
// 返回:
//   - res: 获取线路详情的响应结果，包含线路详细信息和站点列表。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) GetRouteDetail(ctx context.Context, req *v1.GetRouteDetailReq) (res *v1.GetRouteDetailRes, err error) {
	blog.ControllerInfo(ctx, "GetRouteDetail", "获取线路详情请求: %s", req.RouteUuid)

	// 获取线路详情
	iRoute := service.Route()
	routeDetail, errorCode := iRoute.GetRouteDetail(ctx, req.RouteUuid)
	if errorCode != nil {
		return nil, errorCode
	}

	// 转换为API响应DTO
	var apiRouteDetail v1.RouteDetailDTO
	if err := gconv.Struct(routeDetail, &apiRouteDetail); err != nil {
		blog.ControllerError(ctx, "GetRouteDetail", "数据转换失败: %s", err.Error())
		return nil, &berror.ErrInternalServer
	}

	// 转换站点信息
	if len(routeDetail.Stops) > 0 {
		apiRouteDetail.Stops = make([]v1.StationDTO, len(routeDetail.Stops))
		for i, stop := range routeDetail.Stops {
			if err := gconv.Struct(stop, &apiRouteDetail.Stops[i]); err != nil {
				blog.ControllerError(ctx, "GetRouteDetail", "站点数据转换失败: %s", err.Error())
				return nil, &berror.ErrInternalServer
			}
		}
	}

	// 返回结果
	return &v1.GetRouteDetailRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取线路详情成功", &apiRouteDetail),
	}, nil
}
