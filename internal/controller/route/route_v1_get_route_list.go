package route

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/util/gconv"
	"xiao-yi-transit/api/route/v1"
	"xiao-yi-transit/internal/model/dto/route"
	"xiao-yi-transit/internal/service"
)

// GetRouteList 处理获取线路列表的逻辑。
//
// 实现线路列表的查询过程，包括分页、筛选等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取线路列表的请求参数，包含分页和筛选条件。
//
// 返回:
//   - res: 获取线路列表的响应结果，包含线路列表和分页信息。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) GetRouteList(ctx context.Context, req *v1.GetRouteListReq) (res *v1.GetRouteListRes, err error) {
	blog.ControllerInfo(ctx, "GetRouteList", "获取线路列表请求: page=%d, size=%d", req.Page, req.Size)

	// 构建请求DTO
	reqDTO := &route.GetRouteListReqDTO{
		Page:        req.Page,
		Size:        req.Size,
		RouteNumber: req.RouteNumber,
		Name:        req.Name,
		Status:      req.Status,
	}

	// 获取线路列表
	iRoute := service.Route()
	routes, total, errorCode := iRoute.GetRouteList(ctx, reqDTO)
	if errorCode != nil {
		return nil, errorCode
	}

	// 转换数据
	var routeList []*v1.RouteListItemDTO
	for _, routeEntity := range routes {
		var item v1.RouteListItemDTO
		if err := gconv.Struct(routeEntity, &item); err != nil {
			blog.ControllerError(ctx, "GetRouteList", "数据转换失败: %s", err.Error())
			return nil, &berror.ErrInternalServer
		}
		routeList = append(routeList, &item)
	}

	// 构建返回数据
	pagedList := &v1.PagedRouteListDTO{
		List:  routeList,
		Page:  req.Page,
		Size:  req.Size,
		Total: total,
	}

	// 返回结果
	return &v1.GetRouteListRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取线路列表成功", pagedList),
	}, nil
}
