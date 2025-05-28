package station

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/frame/g"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// UpdateStation 更新站点信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - station: 站点信息实体，包含需要更新的站点详细信息。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sStation) UpdateStation(ctx context.Context, station *entity.Station) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "UpdateStation", "更新站点: %s", station.StationUuid)

	// 检查站点是否存在
	stationModel := dao.Station.Ctx(ctx)
	count, sqlErr := stationModel.Where(dao.Station.Columns().StationUuid, station.StationUuid).Count()
	if sqlErr != nil {
		blog.ServiceError(ctx, "UpdateStation", "查询站点失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}
	if count == 0 {
		blog.ServiceError(ctx, "UpdateStation", "站点不存在: %s", station.StationUuid)
		return &berror.ErrResourceNotFound
	}

	// 更新站点信息
	_, sqlErr = stationModel.
		Where(dao.Station.Columns().StationUuid, station.StationUuid).
		Update(g.Map{
			dao.Station.Columns().Name:      station.Name,
			dao.Station.Columns().Code:      station.Code,
			dao.Station.Columns().Address:   station.Address,
			dao.Station.Columns().Longitude: station.Longitude,
			dao.Station.Columns().Latitude:  station.Latitude,
			dao.Station.Columns().Status:    station.Status,
		})
	if sqlErr != nil {
		blog.ServiceError(ctx, "UpdateStation", "更新站点失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 返回结果
	return nil
}
