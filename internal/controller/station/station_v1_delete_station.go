package station

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/station/v1"
	"xiao-yi-transit/internal/service"
)

// DeleteStation 处理删除站点信息的逻辑。
//
// 实现站点信息的删除过程，包括数据验证、信息删除等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 删除站点的请求参数，包含站点UUID。
//
// 返回:
//   - res: 删除站点的响应结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) DeleteStation(ctx context.Context, req *v1.DeleteStationReq) (res *v1.DeleteStationRes, err error) {
	blog.ControllerInfo(ctx, "DeleteStation", "删除站点请求: %s", req.StationUuid)

	// 调用服务删除站点
	iStation := service.Station()
	errorCode := iStation.DeleteStation(ctx, req.StationUuid)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.DeleteStationRes{
		ResponseDTO: bresult.Success(ctx, "删除站点成功"),
	}, nil
}
