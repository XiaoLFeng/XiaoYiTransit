package maintenance

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/api/maintenance/v1"
	"xiao-yi-transit/internal/service"
)

func (c *ControllerV1) DeleteMaintenance(ctx context.Context, req *v1.DeleteMaintenanceReq) (res *v1.DeleteMaintenanceRes, err error) {
	blog.ControllerInfo(ctx, "DeleteMaintenance", "删除维修记录请求: 维修记录UUID=%s", req.MaintenanceUuid)

	// 调用服务删除维修记录
	errorCode := service.Maintenance().DeleteMaintenance(ctx, req.MaintenanceUuid)
	if errorCode != nil {
		return nil, errorCode
	}

	// 返回结果
	return &v1.DeleteMaintenanceRes{
		ResponseDTO: bresult.Success(ctx, "删除维修记录成功"),
	}, nil
}
