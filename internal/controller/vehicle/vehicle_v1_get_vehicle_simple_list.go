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

// GetVehicleSimpleList 处理获取车辆简易列表的逻辑。
//
// 实现车辆简易列表的查询过程，用于前端下拉选项框。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取车辆简易列表的请求参数。
//
// 返回:
//   - res: 获取车辆简易列表的响应结果，包含车辆简易列表。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) GetVehicleSimpleList(ctx context.Context, req *v1.GetVehicleSimpleListReq) (res *v1.GetVehicleSimpleListRes, err error) {
	blog.ControllerInfo(ctx, "GetVehicleSimpleList", "获取车辆简易列表请求")

	// 获取车辆简易列表
	iVehicle := service.Vehicle()
	vehicles, errorCode := iVehicle.GetVehicleSimpleList(ctx)
	if errorCode != nil {
		return nil, errorCode
	}

	// 转换数据
	var vehicleList []*vehicle.SimpleVehicleItemDTO
	for _, vehicleEntity := range vehicles {
		var item vehicle.SimpleVehicleItemDTO
		if err := gconv.Struct(vehicleEntity, &item); err != nil {
			blog.ControllerError(ctx, "GetVehicleSimpleList", "数据转换失败: %s", err.Error())
			return nil, &berror.ErrInternalServer
		}
		vehicleList = append(vehicleList, &item)
	}

	// 构建返回数据
	simpleList := &vehicle.SimpleVehicleListDTO{
		List: vehicleList,
	}

	// 返回结果
	return &v1.GetVehicleSimpleListRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取车辆简易列表成功", simpleList),
	}, nil
}
