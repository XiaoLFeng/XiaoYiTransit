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

func (c *ControllerV1) CreateMaintenance(ctx context.Context, req *v1.CreateMaintenanceReq) (res *v1.CreateMaintenanceRes, err error) {
	blog.ControllerInfo(ctx, "CreateMaintenance", "创建维修记录请求: 车辆UUID=%s, 维修类型=%d", req.VehicleUuid, req.MaintenanceType)

	// 创建维修记录实体
	maintenance := &entity.Maintenance{
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
		CreatedByUuid:       "", // 可以从上下文中获取当前用户UUID
		CreatedAt:           gtime.Now(),
		UpdatedAt:           gtime.Now(),
	}

	// 调用服务创建维修记录
	maintenanceUuid, errorCode := service.Maintenance().CreateMaintenance(ctx, maintenance)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.CreateMaintenanceRes{
		ResponseDTO: bresult.Success(ctx, "创建维修记录成功，UUID: "+maintenanceUuid),
	}, nil
}
