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

// GetVehicleInsuranceList 处理获取车辆保险记录列表的逻辑。
//
// 实现车辆保险记录列表的查询过程，包括分页、筛选等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取车辆保险记录列表的请求参数，包含分页和筛选条件。
//
// 返回:
//   - res: 获取车辆保险记录列表的响应结果，包含保险记录列表和分页信息。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) GetVehicleInsuranceList(ctx context.Context, req *v1.GetVehicleInsuranceListReq) (res *v1.GetVehicleInsuranceListRes, err error) {
	blog.ControllerInfo(ctx, "GetVehicleInsuranceList", "获取车辆保险记录列表请求: page=%d, size=%d", req.Page, req.Size)

	// 获取车辆保险记录列表
	iVehicle := service.Vehicle()
	insurances, total, errorCode := iVehicle.GetVehicleInsuranceList(ctx, req.Page, req.Size, req.VehicleUuid, req.InsuranceType, req.StartDate, req.EndDate)
	if errorCode != nil {
		return nil, errorCode
	}

	// 转换数据
	var insuranceList []*vehicle.VehicleInsuranceItemDTO
	for _, insuranceEntity := range insurances {
		var item vehicle.VehicleInsuranceItemDTO
		if err := gconv.Struct(insuranceEntity, &item); err != nil {
			blog.ControllerError(ctx, "GetVehicleInsuranceList", "数据转换失败: %s", err.Error())
			return nil, &berror.ErrInternalServer
		}

		// 获取车辆信息以获取车牌号
		vehicleEntity, vehicleErr := iVehicle.GetVehicleById(ctx, insuranceEntity.VehicleUuid)
		if vehicleErr == nil && vehicleEntity != nil {
			item.PlateNumber = vehicleEntity.PlateNumber
		}

		insuranceList = append(insuranceList, &item)
	}

	// 构建返回数据
	pagedList := &vehicle.PagedVehicleInsuranceListDTO{
		List:  insuranceList,
		Page:  req.Page,
		Size:  req.Size,
		Total: total,
	}

	// 返回结果
	return &v1.GetVehicleInsuranceListRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取车辆保险记录列表成功", pagedList),
	}, nil
}
