package station

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/station/v1"
	"xiao-yi-transit/internal/service"
)

// ListStationsSimple 处理获取站点简单列表的逻辑。
//
// 实现站点简单列表的查询过程，适用于下拉框。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取站点简单列表的请求参数。
//
// 返回:
//   - res: 获取站点简单列表的响应结果，包含站点UUID、名称和编码。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) ListStationsSimple(ctx context.Context, req *v1.ListStationsSimpleReq) (res *v1.ListStationsSimpleRes, err error) {
	blog.ControllerInfo(ctx, "ListStationsSimple", "获取站点简单列表请求")

	// 设置默认值
	status := req.Status
	if status <= 0 {
		status = 1 // 默认只获取启用的站点
	}

	// 获取站点列表
	iStation := service.Station()
	stations, _, errorCode := iStation.GetStationList(ctx, "", "", status, 1, 1000) // 获取所有启用的站点，不分页
	if errorCode != nil {
		return nil, errorCode
	}

	// 转换数据
	stationList := make([]*v1.StationSimpleDTO, 0, len(stations))
	for _, stationEntity := range stations {
		stationDTO := &v1.StationSimpleDTO{
			StationUuid: stationEntity.StationUuid,
			Name:        stationEntity.Name,
			Code:        stationEntity.Code,
		}
		stationList = append(stationList, stationDTO)
	}

	// 构建返回数据
	stationsSimpleListDTO := &v1.StationsSimpleListDTO{
		Stations: stationList,
	}

	// 返回结果
	return &v1.ListStationsSimpleRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取站点简单列表成功", stationsSimpleListDTO),
	}, nil
}