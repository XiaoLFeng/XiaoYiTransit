package maintenance

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/util/gconv"
	"xiao-yi-transit/api/maintenance/v1"
	"xiao-yi-transit/internal/model/dto/maintenance"
	"xiao-yi-transit/internal/service"
)

func (c *ControllerV1) GetMaintenanceList(ctx context.Context, req *v1.GetMaintenanceListReq) (res *v1.GetMaintenanceListRes, err error) {
	blog.ControllerInfo(ctx, "GetMaintenanceList", "获取维修记录列表请求: page=%d, size=%d", req.Page, req.Size)

	// 获取维修记录列表
	maintenances, total, errorCode := service.Maintenance().GetMaintenanceList(ctx, req.Page, req.Size, req.VehicleUuid, req.MaintenanceType, req.Status, req.StartDate, req.EndDate)
	if errorCode != nil {
		return nil, errorCode
	}

	// 转换数据
	var maintenanceList []*maintenance.MaintenanceListItemDTO
	for _, maintenanceEntity := range maintenances {
		var item maintenance.MaintenanceListItemDTO
		if err := gconv.Struct(maintenanceEntity, &item); err != nil {
			blog.ControllerError(ctx, "GetMaintenanceList", "数据转换失败: %s", err.Error())
			return nil, &berror.ErrInternalServer
		}

		// 设置维修类型名称
		switch maintenanceEntity.MaintenanceType {
		case 1:
			item.MaintenanceTypeName = "常规保养"
		case 2:
			item.MaintenanceTypeName = "故障维修"
		case 3:
			item.MaintenanceTypeName = "事故维修"
		case 4:
			item.MaintenanceTypeName = "年检维修"
		}

		// 设置状态名称
		switch maintenanceEntity.Status {
		case 0:
			item.StatusName = "已取消"
		case 1:
			item.StatusName = "待维修"
		case 2:
			item.StatusName = "维修中"
		case 3:
			item.StatusName = "已完成"
		}

		maintenanceList = append(maintenanceList, &item)
	}

	// 构建返回数据
	pagedList := &maintenance.PagedMaintenanceListDTO{
		List:  maintenanceList,
		Page:  req.Page,
		Size:  req.Size,
		Total: total,
	}

	// 返回结果
	return &v1.GetMaintenanceListRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取维修记录列表成功", pagedList),
	}, nil
}
