package maintenance

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/maintenance/v1"
	"xiao-yi-transit/internal/service"
)

func (c *ControllerV1) GetMaintenanceDetail(ctx context.Context, req *v1.GetMaintenanceDetailReq) (res *v1.GetMaintenanceDetailRes, err error) {
	blog.ControllerInfo(ctx, "GetMaintenanceDetail", "获取维修记录详情请求: 维修记录UUID=%s", req.MaintenanceUuid)

	// 调用服务获取维修记录详情
	maintenance, errorCode := service.Maintenance().GetMaintenanceById(ctx, req.MaintenanceUuid)
	if errorCode != nil {
		return nil, errorCode
	}

	// 构建响应数据
	maintenanceDTO := &v1.MaintenanceDTO{
		MaintenanceUuid:     maintenance.MaintenanceUuid,
		VehicleUuid:         maintenance.VehicleUuid,
		MaintenanceType:     maintenance.MaintenanceType,
		Description:         maintenance.Description,
		MaintenanceDate:     maintenance.MaintenanceDate,
		CompletionDate:      maintenance.CompletionDate,
		Cost:                maintenance.Cost,
		Mileage:             maintenance.Mileage,
		MaintenanceLocation: maintenance.MaintenanceLocation,
		MaintenanceStaff:    maintenance.MaintenanceStaff,
		PartsReplaced:       maintenance.PartsReplaced,
		Status:              maintenance.Status,
		Notes:               maintenance.Notes,
		CreatedAt:           maintenance.CreatedAt,
		UpdatedAt:           maintenance.UpdatedAt,
	}

	// 设置维修类型名称
	switch maintenance.MaintenanceType {
	case 1:
		maintenanceDTO.MaintenanceTypeName = "常规保养"
	case 2:
		maintenanceDTO.MaintenanceTypeName = "故障维修"
	case 3:
		maintenanceDTO.MaintenanceTypeName = "事故维修"
	case 4:
		maintenanceDTO.MaintenanceTypeName = "年检维修"
	}

	// 设置状态名称
	switch maintenance.Status {
	case 0:
		maintenanceDTO.StatusName = "已取消"
	case 1:
		maintenanceDTO.StatusName = "待维修"
	case 2:
		maintenanceDTO.StatusName = "维修中"
	case 3:
		maintenanceDTO.StatusName = "已完成"
	}

	// 返回结果
	return &v1.GetMaintenanceDetailRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取维修记录详情成功", maintenanceDTO),
	}, nil
}
