package route

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/route/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

// CreateRoute 处理创建线路信息的逻辑。
//
// 实现线路信息的创建过程，包括数据验证、信息存储等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 创建线路的请求参数，包含线路的详细信息。
//
// 返回:
//   - res: 创建线路的响应结果，包含创建的线路UUID。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) CreateRoute(ctx context.Context, req *v1.CreateRouteReq) (res *v1.CreateRouteRes, err error) {
	blog.ControllerInfo(ctx, "CreateRoute", "创建线路请求: %s", req.RouteNumber)

	// 创建线路信息
	routeEntity := &entity.Route{
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

	// 调用服务创建线路
	iRoute := service.Route()
	routeUuid, errorCode := iRoute.CreateRoute(ctx, routeEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.CreateRouteRes{
		ResponseDTO: bresult.Success(ctx, "创建线路成功，UUID: " + routeUuid),
	}, nil
}
