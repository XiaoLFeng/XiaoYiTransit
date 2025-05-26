package route

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/frame/g"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// UpdateRoute 更新线路信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - route: 线路信息实体，包含需要更新的线路详细信息。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sRoute) UpdateRoute(ctx context.Context, route *entity.Route) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "UpdateRoute", "更新线路: %s", route.RouteUuid)

	// 检查线路是否存在
	routeModel := dao.Route.Ctx(ctx)
	count, sqlErr := routeModel.Where(dao.Route.Columns().RouteUuid, route.RouteUuid).Count()
	if sqlErr != nil {
		blog.ServiceError(ctx, "UpdateRoute", "查询线路失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}
	if count == 0 {
		blog.ServiceError(ctx, "UpdateRoute", "线路不存在: %s", route.RouteUuid)
		return &berror.ErrResourceNotFound
	}

	// 更新线路信息
	_, sqlErr = routeModel.
		Where(dao.Route.Columns().RouteUuid, route.RouteUuid).
		Update(g.Map{
			dao.Route.Columns().RouteNumber:    route.RouteNumber,
			dao.Route.Columns().Name:           route.Name,
			dao.Route.Columns().StartStation:   route.StartStation,
			dao.Route.Columns().EndStation:     route.EndStation,
			dao.Route.Columns().Stops:          route.Stops,
			dao.Route.Columns().Distance:       route.Distance,
			dao.Route.Columns().Fare:           route.Fare,
			dao.Route.Columns().OperationHours: route.OperationHours,
			dao.Route.Columns().Frequency:      route.Frequency,
			dao.Route.Columns().Status:         route.Status,
			dao.Route.Columns().Notes:          route.Notes,
		})
	if sqlErr != nil {
		blog.ServiceError(ctx, "UpdateRoute", "更新线路失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 返回结果
	return nil
}
