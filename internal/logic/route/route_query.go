package route

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/util/gconv"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/dto/route"
	"xiao-yi-transit/internal/model/entity"
)

// GetRouteById 根据线路UUID获取线路信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - routeUuid: 线路的唯一标识符。
//
// 返回:
//   - 线路信息的指针，包含详细线路数据。
//   - 错误码的指针，表示错误类型，如线路不存在或内部错误。
func (s *sRoute) GetRouteById(ctx context.Context, routeUuid string) (*entity.Route, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetRouteById", "获取线路: %s", routeUuid)

	// 查询线路信息
	routeModel := dao.Route.Ctx(ctx)
	routeResult, sqlErr := routeModel.Where(dao.Route.Columns().RouteUuid, routeUuid).One()
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetRouteById", "查询线路失败: %s", sqlErr.Error())
		return nil, &berror.ErrDatabaseError
	}
	if routeResult.IsEmpty() {
		blog.ServiceError(ctx, "GetRouteById", "线路不存在: %s", routeUuid)
		return nil, &berror.ErrResourceNotFound
	}

	// 转换数据
	var routeEntity *entity.Route
	if err := gconv.Struct(routeResult, &routeEntity); err != nil {
		blog.ServiceError(ctx, "GetRouteById", "数据转换失败: %s", err.Error())
		return nil, &berror.ErrInternalServer
	}

	// 返回结果
	return routeEntity, nil
}

// GetRouteList 获取线路列表。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - req: 获取线路列表的请求参数，包含分页信息和筛选条件。
//
// 返回:
//   - 线路列表。
//   - 总数量。
//   - 错误码的指针，表示可能的错误类型。
func (s *sRoute) GetRouteList(ctx context.Context, req *route.GetRouteListReqDTO) ([]*entity.Route, int, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetRouteList", "获取线路列表: 页码=%d, 每页数量=%d", req.Page, req.Size)

	// 构建查询条件
	routeModel := dao.Route.Ctx(ctx)
	if req.RouteNumber != "" {
		routeModel = routeModel.WhereLike(dao.Route.Columns().RouteNumber, "%"+req.RouteNumber+"%")
	}
	if req.Name != "" {
		routeModel = routeModel.WhereLike(dao.Route.Columns().Name, "%"+req.Name+"%")
	}
	if req.Status != 0 {
		routeModel = routeModel.Where(dao.Route.Columns().Status, req.Status)
	}

	// 查询总数
	count, sqlErr := routeModel.Count()
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetRouteList", "查询线路总数失败: %s", sqlErr.Error())
		return nil, 0, &berror.ErrDatabaseError
	}

	// 查询列表
	routeResults, sqlErr := routeModel.Page(req.Page, req.Size).Order(dao.Route.Columns().CreatedAt + " DESC").All()
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetRouteList", "查询线路列表失败: %s", sqlErr.Error())
		return nil, 0, &berror.ErrDatabaseError
	}

	// 转换数据
	var routeList []*entity.Route
	if err := gconv.Structs(routeResults, &routeList); err != nil {
		blog.ServiceError(ctx, "GetRouteList", "数据转换失败: %s", err.Error())
		return nil, 0, &berror.ErrInternalServer
	}

	// 返回结果
	return routeList, count, nil
}

// GetRouteDetail 获取线路详情，包括站点信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - routeUuid: 线路的唯一标识符。
//
// 返回:
//   - 线路详情DTO，包含线路信息和站点列表。
//   - 错误码的指针，表示可能的错误类型。
func (s *sRoute) GetRouteDetail(ctx context.Context, routeUuid string) (*route.RouteDetailDTO, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetRouteDetail", "获取线路详情: %s", routeUuid)

	// 查询线路信息
	routeEntity, errCode := s.GetRouteById(ctx, routeUuid)
	if errCode != nil {
		return nil, errCode
	}

	// 查询线路站点
	routeStations, errCode := s.GetRouteStations(ctx, routeUuid)
	if errCode != nil {
		return nil, errCode
	}

	// 构建返回结果
	result := &route.RouteDetailDTO{
		RouteUuid:      routeEntity.RouteUuid,
		RouteNumber:    routeEntity.RouteNumber,
		Name:           routeEntity.Name,
		StartStation:   routeEntity.StartStation,
		EndStation:     routeEntity.EndStation,
		Distance:       routeEntity.Distance,
		Fare:           routeEntity.Fare,
		OperationHours: routeEntity.OperationHours,
		Frequency:      routeEntity.Frequency,
		Status:         routeEntity.Status,
		Notes:          routeEntity.Notes,
		CreatedAt:      routeEntity.CreatedAt,
		UpdatedAt:      routeEntity.UpdatedAt,
	}

	// 如果有站点信息，添加到结果中
	if routeStations != nil && len(routeStations.Stations) > 0 {
		result.Stops = make([]route.StationDTO, len(routeStations.Stations))
		for i, station := range routeStations.Stations {
			result.Stops[i] = route.StationDTO{
				StationUuid:       station.StationUuid,
				Name:              station.Name,
				Sequence:          station.Sequence,
				DistanceFromStart: station.DistanceFromStart,
				EstimatedTime:     station.EstimatedTime,
			}
		}
	}

	// 返回结果
	return result, nil
}
