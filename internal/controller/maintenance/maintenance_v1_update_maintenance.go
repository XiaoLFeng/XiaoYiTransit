package maintenance

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/os/gtime"
	"xiao-yi-transit/api/maintenance/v1"
	"xiao-yi-transit/internal/model/entity"
	"xiao-yi-transit/internal/service"
)

func (c *ControllerV1) UpdateMaintenance(ctx context.Context, req *v1.UpdateMaintenanceReq) (res *v1.UpdateMaintenanceRes, err error) {
	blog.ControllerInfo(ctx, "UpdateMaintenance", "更新维修记录请求: 维修记录UUID=%s", req.MaintenanceUuid)

	// 创建维修记录实体
	maintenance := &entity.Maintenance{
		MaintenanceUuid:     req.MaintenanceUuid,
		VehicleUuid:         req.VehicleUuid,
		MaintenanceType:     req.MaintenanceType,
		Description:         req.Description,
		MaintenanceDate:     req.MaintenanceDate,
		CompletionDate:      req.CompletionDate,
		Cost:                req.Cost,
		Mileage:             req.Mileage,
		MaintenanceLocation: req.MaintenanceLocation,
		MaintenanceStaff:    req.MaintenanceStaff,
		PartsReplaced:       req.PartsReplaced,
		Status:              req.Status,
		Notes:               req.Notes,
		UpdatedAt:           gtime.Now(),
	}

	// 调用服务更新维修记录
	errorCode := service.Maintenance().UpdateMaintenance(ctx, maintenance)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.UpdateMaintenanceRes{
		ResponseDTO: bresult.Success(ctx, "更新维修记录成功"),
	}, nil
}
