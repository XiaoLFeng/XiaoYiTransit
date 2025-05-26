// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xiao-yi-transit/internal/model/dto/route"
	"xiao-yi-transit/internal/model/entity"

	"github.com/XiaoLFeng/bamboo-utils/berror"
)

type (
	IRoute interface {
		// CreateRoute 创建新的线路信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - route: 线路信息实体，包含需要创建的线路详细信息。
		//
		// 返回:
		//   - 创建成功的线路UUID。
		//   - 错误码的指针，表示可能的错误类型。
		CreateRoute(ctx context.Context, route *entity.Route) (string, *berror.ErrorCode)
		// DeleteRoute 删除线路信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - routeUuid: 线路的唯一标识符。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		DeleteRoute(ctx context.Context, routeUuid string) *berror.ErrorCode
		// GetRouteById 根据线路UUID获取线路信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - routeUuid: 线路的唯一标识符。
		//
		// 返回:
		//   - 线路信息的指针，包含详细线路数据。
		//   - 错误码的指针，表示错误类型，如线路不存在或内部错误。
		GetRouteById(ctx context.Context, routeUuid string) (*entity.Route, *berror.ErrorCode)
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
		GetRouteList(ctx context.Context, req *route.GetRouteListReqDTO) ([]*entity.Route, int, *berror.ErrorCode)
		// GetRouteDetail 获取线路详情，包括站点信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - routeUuid: 线路的唯一标识符。
		//
		// 返回:
		//   - 线路详情DTO，包含线路信息和站点列表。
		//   - 错误码的指针，表示可能的错误类型。
		GetRouteDetail(ctx context.Context, routeUuid string) (*route.RouteDetailDTO, *berror.ErrorCode)
		// AddRouteStation 向线路添加站点。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - routeStation: 线路站点信息实体，包含需要添加的站点详细信息。
		//
		// 返回:
		//   - 创建成功的线路站点UUID。
		//   - 错误码的指针，表示可能的错误类型。
		AddRouteStation(ctx context.Context, routeStation *entity.RouteStation) (string, *berror.ErrorCode)
		// UpdateRouteStation 更新线路站点信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - routeStation: 线路站点信息实体，包含需要更新的站点详细信息。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		UpdateRouteStation(ctx context.Context, routeStation *entity.RouteStation) *berror.ErrorCode
		// DeleteRouteStation 删除线路站点。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - routeStationUuid: 线路站点的唯一标识符。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		DeleteRouteStation(ctx context.Context, routeStationUuid string) *berror.ErrorCode
		// GetRouteStations 获取线路站点列表。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - routeUuid: 线路的唯一标识符。
		//
		// 返回:
		//   - 线路站点DTO，包含线路信息和站点列表。
		//   - 错误码的指针，表示可能的错误类型。
		GetRouteStations(ctx context.Context, routeUuid string) (*route.RouteStationsDTO, *berror.ErrorCode)
		// UpdateRoute 更新线路信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - route: 线路信息实体，包含需要更新的线路详细信息。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		UpdateRoute(ctx context.Context, route *entity.Route) *berror.ErrorCode
	}
)

var (
	localRoute IRoute
)

func Route() IRoute {
	if localRoute == nil {
		panic("implement not found for interface IRoute, forgot register?")
	}
	return localRoute
}

func RegisterRoute(i IRoute) {
	localRoute = i
}
