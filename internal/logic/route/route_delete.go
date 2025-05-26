package route

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"xiao-yi-transit/internal/dao"
)

// DeleteRoute 删除线路信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - routeUuid: 线路的唯一标识符。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sRoute) DeleteRoute(ctx context.Context, routeUuid string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "DeleteRoute", "删除线路: %s", routeUuid)

	// 检查线路是否存在
	routeModel := dao.Route.Ctx(ctx)
	count, sqlErr := routeModel.Where(dao.Route.Columns().RouteUuid, routeUuid).Count()
	if sqlErr != nil {
		blog.ServiceError(ctx, "DeleteRoute", "查询线路失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}
	if count == 0 {
		blog.ServiceError(ctx, "DeleteRoute", "线路不存在: %s", routeUuid)
		return &berror.ErrResourceNotFound
	}

	// 删除线路站点
	routeStationModel := dao.RouteStation.Ctx(ctx)
	_, sqlErr = routeStationModel.Where(dao.RouteStation.Columns().RouteUuid, routeUuid).Delete()
	if sqlErr != nil {
		blog.ServiceError(ctx, "DeleteRoute", "删除线路站点失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 删除线路信息
	_, sqlErr = routeModel.Where(dao.Route.Columns().RouteUuid, routeUuid).Delete()
	if sqlErr != nil {
		blog.ServiceError(ctx, "DeleteRoute", "删除线路失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 返回结果
	return nil
}
