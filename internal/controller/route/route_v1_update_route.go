package route

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/route/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

// UpdateRoute 处理更新线路信息的逻辑。
//
// 实现线路信息的更新过程，包括数据验证、信息更新等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 更新线路的请求参数，包含线路的详细信息。
//
// 返回:
//   - res: 更新线路的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) UpdateRoute(ctx context.Context, req *v1.UpdateRouteReq) (res *v1.UpdateRouteRes, err error) {
	blog.ControllerInfo(ctx, "UpdateRoute", "更新线路请求: %s", req.RouteUuid)

	// 创建线路信息实体
	routeEntity := &entity.Route{
		RouteUuid:      req.RouteUuid,
		RouteNumber:    req.RouteNumber,
		Name:           req.Name,
		StartStation:   req.StartStation,
		EndStation:     req.EndStation,
		Stops:          req.Stops,
		Distance:       req.Distance,
		Fare:           req.Fare,
		OperationHours: req.OperationHours,
		Frequency:      req.Frequency,
		Status:         req.Status,
		Notes:          req.Notes,
	}

	// 调用服务更新线路
	iRoute := service.Route()
	errorCode := iRoute.UpdateRoute(ctx, routeEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.UpdateRouteRes{
		ResponseDTO: bresult.Success(ctx, "更新线路成功"),
	}, nil
}
