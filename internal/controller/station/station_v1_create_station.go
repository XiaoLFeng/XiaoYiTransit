package station

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/station/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

// CreateStation 处理创建站点信息的逻辑。
//
// 实现站点信息的创建过程，包括数据验证、信息存储等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 创建站点的请求参数，包含站点的详细信息。
//
// 返回:
//   - res: 创建站点的响应结果，包含创建的站点UUID。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) CreateStation(ctx context.Context, req *v1.CreateStationReq) (res *v1.CreateStationRes, err error) {
	blog.ControllerInfo(ctx, "CreateStation", "创建站点请求: %s", req.Name)

	// 创建站点信息
	stationEntity := &entity.Station{
		Name:      req.Name,
		Code:      req.Code,
		Address:   req.Address,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
		Status:    req.Status,
	}

	// 调用服务创建站点
	iStation := service.Station()
	stationUuid, errorCode := iStation.CreateStation(ctx, stationEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 查询创建的站点信息
	station, errorCode := iStation.GetStationById(ctx, stationUuid)
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
	return &v1.CreateStationRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "创建站点成功", stationDTO),
	}, nil
}
