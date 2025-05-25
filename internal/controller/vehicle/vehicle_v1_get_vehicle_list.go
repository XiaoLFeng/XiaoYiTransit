package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/util/gconv"
	"xiao-yi-transit/api/vehicle/v1"
	"xiao-yi-transit/internal/model/dto/vehicle"
	"xiao-yi-transit/internal/service"
)

// GetVehicleList 处理获取车辆列表的逻辑。
//
// 实现车辆列表的查询过程，包括分页、筛选等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取车辆列表的请求参数，包含分页和筛选条件。
//
// 返回:
//   - res: 获取车辆列表的响应结果，包含车辆列表和分页信息。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) GetVehicleList(ctx context.Context, req *v1.GetVehicleListReq) (res *v1.GetVehicleListRes, err error) {
	blog.ControllerInfo(ctx, "GetVehicleList", "获取车辆列表请求: page=%d, size=%d", req.Page, req.Size)

	// 获取车辆列表
	iVehicle := service.Vehicle()
	vehicles, total, errorCode := iVehicle.GetVehicleList(ctx, req.Page, req.Size, req.PlateNumber, req.Model, req.Status)
	if errorCode != nil {
		return nil, errorCode
	}

	// 转换数据
	var vehicleList []*vehicle.VehicleListItemDTO
	for _, vehicleEntity := range vehicles {
		var item vehicle.VehicleListItemDTO
		if err := gconv.Struct(vehicleEntity, &item); err != nil {
			blog.ControllerError(ctx, "GetVehicleList", "数据转换失败: %s", err.Error())
			return nil, &berror.ErrInternalServer
		}
		vehicleList = append(vehicleList, &item)
	}

	// 构建返回数据
	pagedList := &vehicle.PagedVehicleListDTO{
		List:  vehicleList,
		Page:  req.Page,
		Size:  req.Size,
		Total: total,
	}

	// 返回结果
	return &v1.GetVehicleListRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取车辆列表成功", pagedList),
	}, nil
}
