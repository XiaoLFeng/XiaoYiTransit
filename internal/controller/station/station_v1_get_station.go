package station

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/station/v1"
	"xiao-yi-transit/internal/service"
)

// GetStation 处理获取站点详情的逻辑。
//
// 实现站点详情的查询过程，包括站点基本信息。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取站点详情的请求参数，包含站点UUID。
//
// 返回:
//   - res: 获取站点详情的响应结果，包含站点详细信息。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) GetStation(ctx context.Context, req *v1.GetStationReq) (res *v1.GetStationRes, err error) {
	blog.ControllerInfo(ctx, "GetStation", "获取站点详情请求: %s", req.StationUuid)

	// 获取站点详情
	iStation := service.Station()
	station, errorCode := iStation.GetStationById(ctx, req.StationUuid)
	if errorCode != nil {
		return nil, errorCode
	}

	// 构建返回数据
	stationDTO := &v1.StationDTO{
		StationUuid: station.StationUuid,
		Name:        station.Name,
		Code:        station.Code,
		Address:     station.Address,
		Longitude:   station.Longitude,
		Latitude:    station.Latitude,
		Status:      station.Status,
		CreatedAt:   station.CreatedAt.String(),
		UpdatedAt:   station.UpdatedAt.String(),
	}

	// 返回结果
	return &v1.GetStationRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取站点详情成功", stationDTO),
	}, nil
}
