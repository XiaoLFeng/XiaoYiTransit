package route

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/google/uuid"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// CreateRoute 创建新的线路信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - route: 线路信息实体，包含需要创建的线路详细信息。
//
// 返回:
//   - 创建成功的线路UUID。
//   - 错误码的指针，表示可能的错误类型。
func (s *sRoute) CreateRoute(ctx context.Context, route *entity.Route) (string, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "CreateRoute", "创建线路: %s", route.RouteNumber)

	// 生成UUID
	route.RouteUuid = uuid.New().String()

	// 创建线路信息
	routeModel := dao.Route.Ctx(ctx)
	_, sqlErr := routeModel.Insert(route)
	if sqlErr != nil {
		blog.ServiceError(ctx, "CreateRoute", "创建线路失败: %s", sqlErr.Error())
		return "", &berror.ErrDatabaseError
	}

	// 返回结果
	return route.RouteUuid, nil
}
