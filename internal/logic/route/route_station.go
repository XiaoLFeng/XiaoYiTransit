package route

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/dto/route"
	"xiao-yi-transit/internal/model/entity"
)

// AddRouteStation 向线路添加站点。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - routeStation: 线路站点信息实体，包含需要添加的站点详细信息。
//
// 返回:
//   - 创建成功的线路站点UUID。
//   - 错误码的指针，表示可能的错误类型。
func (s *sRoute) AddRouteStation(ctx context.Context, routeStation *entity.RouteStation) (string, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "AddRouteStation", "添加线路站点: 线路=%s, 站点=%s", routeStation.RouteUuid, routeStation.StationUuid)

	// 生成UUID
	routeStation.RouteStationUuid = uuid.New().String()

	// 创建线路站点信息
	routeStationModel := dao.RouteStation.Ctx(ctx)
	_, sqlErr := routeStationModel.Insert(routeStation)
	if sqlErr != nil {
		blog.ServiceError(ctx, "AddRouteStation", "添加线路站点失败: %s", sqlErr.Error())
		return "", &berror.ErrDatabaseError
	}

	// 返回结果
	return routeStation.RouteStationUuid, nil
}

// UpdateRouteStation 更新线路站点信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - routeStation: 线路站点信息实体，包含需要更新的站点详细信息。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sRoute) UpdateRouteStation(ctx context.Context, routeStation *entity.RouteStation) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "UpdateRouteStation", "更新线路站点: %s", routeStation.RouteStationUuid)

	// 更新线路站点信息
	routeStationModel := dao.RouteStation.Ctx(ctx)
	_, sqlErr := routeStationModel.
		Where(dao.RouteStation.Columns().RouteStationUuid, routeStation.RouteStationUuid).
		Update(g.Map{
			dao.RouteStation.Columns().Sequence:          routeStation.Sequence,
			dao.RouteStation.Columns().DistanceFromStart: routeStation.DistanceFromStart,
			dao.RouteStation.Columns().EstimatedTime:     routeStation.EstimatedTime,
		})
	if sqlErr != nil {
		blog.ServiceError(ctx, "UpdateRouteStation", "更新线路站点失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 返回结果
	return nil
}

// DeleteRouteStation 删除线路站点。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - routeStationUuid: 线路站点的唯一标识符。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sRoute) DeleteRouteStation(ctx context.Context, routeStationUuid string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "DeleteRouteStation", "删除线路站点: %s", routeStationUuid)

	// 删除线路站点信息
	routeStationModel := dao.RouteStation.Ctx(ctx)
	_, sqlErr := routeStationModel.Where(dao.RouteStation.Columns().RouteStationUuid, routeStationUuid).Delete()
	if sqlErr != nil {
		blog.ServiceError(ctx, "DeleteRouteStation", "删除线路站点失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 返回结果
	return nil
}

// GetRouteStations 获取线路站点列表。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - routeUuid: 线路的唯一标识符。
//
// 返回:
//   - 线路站点DTO，包含线路信息和站点列表。
//   - 错误码的指针，表示可能的错误类型。
func (s *sRoute) GetRouteStations(ctx context.Context, routeUuid string) (*route.RouteStationsDTO, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetRouteStations", "获取线路站点: %s", routeUuid)

	// 查询线路站点信息
	routeStationModel := dao.RouteStation.Ctx(ctx)
	routeStations, sqlErr := routeStationModel.
		Where(dao.RouteStation.Columns().RouteUuid, routeUuid).
		Order(dao.RouteStation.Columns().Sequence + " ASC").
		All()
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetRouteStations", "查询线路站点失败: %s", sqlErr.Error())
		return nil, &berror.ErrDatabaseError
	}

	// 转换路线站点数据
	var routeStationEntities []*entity.RouteStation
	if err := gconv.Structs(routeStations, &routeStationEntities); err != nil {
		blog.ServiceError(ctx, "GetRouteStations", "转换路线站点数据失败: %s", err.Error())
		return nil, &berror.ErrInternalServer
	}

	// 如果没有站点，返回空列表
	if len(routeStationEntities) == 0 {
		return &route.RouteStationsDTO{
			RouteUuid: routeUuid,
			Stations:  []*route.RouteStationItemDTO{},
		}, nil
	}

	// 获取站点UUID列表
	stationUuids := make([]string, len(routeStationEntities))
	for i, rs := range routeStationEntities {
		stationUuids[i] = rs.StationUuid
	}

	// 查询站点信息
	stationModel := dao.Station.Ctx(ctx)
	stationsResult, sqlErr := stationModel.
		WhereIn(dao.Station.Columns().StationUuid, stationUuids).
		All()
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetRouteStations", "查询站点信息失败: %s", sqlErr.Error())
		return nil, &berror.ErrDatabaseError
	}

	// 转换站点数据
	var stationEntities []*entity.Station
	if err := gconv.Structs(stationsResult, &stationEntities); err != nil {
		blog.ServiceError(ctx, "GetRouteStations", "转换站点数据失败: %s", err.Error())
		return nil, &berror.ErrInternalServer
	}

	// 构建站点映射
	stationMap := make(map[string]*entity.Station)
	for _, station := range stationEntities {
		stationMap[station.StationUuid] = station
	}

	// 构建返回结果
	result := &route.RouteStationsDTO{
		RouteUuid: routeUuid,
		Stations:  make([]*route.RouteStationItemDTO, len(routeStationEntities)),
	}

	// 填充站点信息
	for i, rs := range routeStationEntities {
		station := stationMap[rs.StationUuid]
		if station == nil {
			continue
		}

		result.Stations[i] = &route.RouteStationItemDTO{
			RouteStationUuid:  rs.RouteStationUuid,
			StationUuid:       station.StationUuid,
			Name:              station.Name,
			Code:              station.Code,
			Address:           station.Address,
			Longitude:         station.Longitude,
			Latitude:          station.Latitude,
			Sequence:          rs.Sequence,
			DistanceFromStart: rs.DistanceFromStart,
			EstimatedTime:     rs.EstimatedTime,
		}
	}

	// 返回结果
	return result, nil
}
