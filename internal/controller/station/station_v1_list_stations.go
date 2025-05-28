package station

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/station/v1"
	"xiao-yi-transit/internal/service"
)

// ListStations 处理获取站点列表的逻辑。
//
// 实现站点列表的查询过程，包括分页、筛选等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取站点列表的请求参数，包含分页和筛选条件。
//
// 返回:
//   - res: 获取站点列表的响应结果，包含站点列表和分页信息。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) ListStations(ctx context.Context, req *v1.ListStationsReq) (res *v1.ListStationsRes, err error) {
	blog.ControllerInfo(ctx, "ListStations", "获取站点列表请求: page=%d, size=%d", req.Page, req.Size)

	// 设置默认值
	page := req.Page
	if page <= 0 {
		page = 1
	}
	size := req.Size
	if size <= 0 {
		size = 10
	}

	// 获取站点列表
	iStation := service.Station()
	stations, total, errorCode := iStation.GetStationList(ctx, req.Name, req.Code, req.Status, page, size)
	if errorCode != nil {
		return nil, errorCode
	}

	// 转换数据
	stationList := make([]*v1.StationDTO, 0, len(stations))
	for _, stationEntity := range stations {
		stationDTO := &v1.StationDTO{
			StationUuid: stationEntity.StationUuid,
			Name:        stationEntity.Name,
			Code:        stationEntity.Code,
			Address:     stationEntity.Address,
			Longitude:   stationEntity.Longitude,
			Latitude:    stationEntity.Latitude,
			Status:      stationEntity.Status,
			CreatedAt:   stationEntity.CreatedAt.String(),
			UpdatedAt:   stationEntity.UpdatedAt.String(),
		}
		stationList = append(stationList, stationDTO)
	}

	// 构建返回数据
	stationsListDTO := &v1.StationsListDTO{
		Total:    total,
		Page:     page,
		Size:     size,
		Stations: stationList,
	}

	// 返回结果
	return &v1.ListStationsRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取站点列表成功", stationsListDTO),
	}, nil
}
