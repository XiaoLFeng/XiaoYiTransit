package station

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"xiao-yi-transit/internal/dao"
)

// DeleteStation 删除站点信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - stationUuid: 站点的唯一标识符。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sStation) DeleteStation(ctx context.Context, stationUuid string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "DeleteStation", "删除站点: %s", stationUuid)

	// 检查站点是否存在
	stationModel := dao.Station.Ctx(ctx)
	count, sqlErr := stationModel.Where(dao.Station.Columns().StationUuid, stationUuid).Count()
	if sqlErr != nil {
		blog.ServiceError(ctx, "DeleteStation", "查询站点失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}
	if count == 0 {
		blog.ServiceError(ctx, "DeleteStation", "站点不存在: %s", stationUuid)
		return &berror.ErrResourceNotFound
	}

	// 检查站点是否被线路使用
	routeStationModel := dao.RouteStation.Ctx(ctx)
	count, sqlErr = routeStationModel.Where(dao.RouteStation.Columns().StationUuid, stationUuid).Count()
	if sqlErr != nil {
		blog.ServiceError(ctx, "DeleteStation", "查询线路站点失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}
	if count > 0 {
		blog.ServiceError(ctx, "DeleteStation", "站点正在被线路使用，无法删除: %s", stationUuid)
		return &berror.ErrInvalidOperation
	}

	// 删除站点信息
	_, sqlErr = stationModel.Where(dao.Station.Columns().StationUuid, stationUuid).Delete()
	if sqlErr != nil {
		blog.ServiceError(ctx, "DeleteStation", "删除站点失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 返回结果
	return nil
}
