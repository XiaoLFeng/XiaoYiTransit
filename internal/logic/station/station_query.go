package station

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/util/gconv"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// GetStationById 根据站点UUID获取站点信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - stationUuid: 站点的唯一标识符。
//
// 返回:
//   - 站点信息的指针，包含详细站点数据。
//   - 错误码的指针，表示错误类型，如站点不存在或内部错误。
func (s *sStation) GetStationById(ctx context.Context, stationUuid string) (*entity.Station, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetStationById", "获取站点: %s", stationUuid)

	// 查询站点信息
	stationModel := dao.Station.Ctx(ctx)
	stationResult, sqlErr := stationModel.Where(dao.Station.Columns().StationUuid, stationUuid).One()
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetStationById", "查询站点失败: %s", sqlErr.Error())
		return nil, &berror.ErrDatabaseError
	}
	if stationResult.IsEmpty() {
		blog.ServiceError(ctx, "GetStationById", "站点不存在: %s", stationUuid)
		return nil, &berror.ErrResourceNotFound
	}

	// 转换数据
	var stationEntity *entity.Station
	if err := gconv.Struct(stationResult, &stationEntity); err != nil {
		blog.ServiceError(ctx, "GetStationById", "数据转换失败: %s", err.Error())
		return nil, &berror.ErrInternalServer
	}

	// 返回结果
	return stationEntity, nil
}

// GetStationList 获取站点列表。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - name: 站点名称，用于模糊查询。
//   - code: 站点编码，用于模糊查询。
//   - status: 状态: 0-停用, 1-启用。
//   - page: 页码，默认为1。
//   - size: 每页数量，默认为10。
//
// 返回:
//   - 站点列表。
//   - 总数量。
//   - 错误码的指针，表示可能的错误类型。
func (s *sStation) GetStationList(ctx context.Context, name string, code string, status int, page int, size int) ([]*entity.Station, int, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetStationList", "获取站点列表: 页码=%d, 每页数量=%d", page, size)

	// 构建查询条件
	stationModel := dao.Station.Ctx(ctx)
	if name != "" {
		stationModel = stationModel.WhereLike(dao.Station.Columns().Name, "%"+name+"%")
	}
	if code != "" {
		stationModel = stationModel.WhereLike(dao.Station.Columns().Code, "%"+code+"%")
	}
	if status != 0 {
		stationModel = stationModel.Where(dao.Station.Columns().Status, status)
	}

	// 查询总数
	count, sqlErr := stationModel.Count()
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetStationList", "查询站点总数失败: %s", sqlErr.Error())
		return nil, 0, &berror.ErrDatabaseError
	}

	// 查询列表
	stationResults, sqlErr := stationModel.Page(page, size).Order(dao.Station.Columns().CreatedAt + " DESC").All()
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetStationList", "查询站点列表失败: %s", sqlErr.Error())
		return nil, 0, &berror.ErrDatabaseError
	}

	// 转换数据
	var stationList []*entity.Station
	if err := gconv.Structs(stationResults, &stationList); err != nil {
		blog.ServiceError(ctx, "GetStationList", "数据转换失败: %s", err.Error())
		return nil, 0, &berror.ErrInternalServer
	}

	// 返回结果
	return stationList, count, nil
}
