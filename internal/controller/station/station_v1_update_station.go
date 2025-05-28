package station

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/station/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

// UpdateStation 处理更新站点信息的逻辑。
//
// 实现站点信息的更新过程，包括数据验证、信息更新等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 更新站点的请求参数，包含站点的详细信息。
//
// 返回:
//   - res: 更新站点的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) UpdateStation(ctx context.Context, req *v1.UpdateStationReq) (res *v1.UpdateStationRes, err error) {
	blog.ControllerInfo(ctx, "UpdateStation", "更新站点请求: %s", req.StationUuid)

	// 创建站点信息实体
	stationEntity := &entity.Station{
		StationUuid: req.StationUuid,
		Name:        req.Name,
		Code:        req.Code,
		Address:     req.Address,
		Longitude:   req.Longitude,
		Latitude:    req.Latitude,
		Status:      req.Status,
	}

	// 调用服务更新站点
	iStation := service.Station()
	errorCode := iStation.UpdateStation(ctx, stationEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.UpdateStationRes{
		ResponseDTO: bresult.Success(ctx, "更新站点成功"),
	}, nil
}
