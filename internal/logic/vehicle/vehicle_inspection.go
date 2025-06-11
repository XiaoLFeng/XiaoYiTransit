package vehicle

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// CreateVehicleInspection 创建新的车辆年检记录。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - inspection: 车辆年检记录实体，包含需要创建的年检详细信息。
//
// 返回:
//   - 创建成功的年检记录UUID。
//   - 错误码的指针，表示可能的错误类型。
func (s *sVehicle) CreateVehicleInspection(ctx context.Context, inspection *entity.VehicleInspection) (string, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "CreateVehicleInspection", "创建车辆年检记录: 车辆UUID=%s", inspection.VehicleUuid)

	// 检查车辆是否存在
	vehicle, err := s.GetVehicleById(ctx, inspection.VehicleUuid)
	if err != nil {
		return "", err
	}
	if vehicle == nil {
		return "", &berror.ErrNotFound
	}

	// 生成UUID
	inspection.InspectionUuid = uuid.New().String()

	// 创建年检记录
	inspectionModel := dao.VehicleInspection.Ctx(ctx)
	_, sqlErr := inspectionModel.Insert(inspection)
	if sqlErr != nil {
		blog.ServiceError(ctx, "CreateVehicleInspection", "创建车辆年检记录失败: %s", sqlErr.Error())
		return "", &berror.ErrDatabaseError
	}

	// 更新车辆的下次年检日期
	vehicle.NextInspectionDate = inspection.ExpiryDate
	updateErr := s.UpdateVehicle(ctx, vehicle)
	if updateErr != nil {
		blog.ServiceError(ctx, "CreateVehicleInspection", "更新车辆下次年检日期失败: %s", updateErr.Error())
	}

	// 返回结果
	return inspection.InspectionUuid, nil
}

// DeleteVehicleInspection 删除车辆年检记录。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - inspectionUuid: 年检记录的唯一标识符。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sVehicle) DeleteVehicleInspection(ctx context.Context, inspectionUuid string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "DeleteVehicleInspection", "删除车辆年检记录: %s", inspectionUuid)

	// 检查年检记录是否存在
	inspection, err := s.GetVehicleInspectionById(ctx, inspectionUuid)
	if err != nil {
		return err
	}
	if inspection == nil {
		return &berror.ErrNotFound
	}

	// 删除年检记录（软删除）
	inspectionModel := dao.VehicleInspection.Ctx(ctx)
	_, sqlErr := inspectionModel.Where("inspection_uuid", inspectionUuid).Delete()
	if sqlErr != nil {
		blog.ServiceError(ctx, "DeleteVehicleInspection", "删除车辆年检记录失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 返回结果
	return nil
}

// GetVehicleInspectionById 根据年检记录UUID获取年检记录信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - inspectionUuid: 年检记录的唯一标识符。
//
// 返回:
//   - 年检记录信息的指针，包含详细年检数据。
//   - 错误码的指针，表示错误类型，如年检记录不存在或内部错误。
func (s *sVehicle) GetVehicleInspectionById(ctx context.Context, inspectionUuid string) (*entity.VehicleInspection, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetVehicleInspectionById", "获取车辆年检记录: %s", inspectionUuid)

	// 查询年检记录信息
	inspectionModel := dao.VehicleInspection.Ctx(ctx)
	inspection := &entity.VehicleInspection{}
	err := inspectionModel.Where("inspection_uuid", inspectionUuid).Scan(inspection)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, &berror.ErrNotFound
		}
		blog.ServiceError(ctx, "GetVehicleInspectionById", "获取车辆年检记录失败: %s", err.Error())
		return nil, &berror.ErrDatabaseError
	}

	// 返回结果
	return inspection, nil
}

// GetVehicleInspectionList 获取车辆年检记录列表。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - page: 页码，从1开始。
//   - size: 每页数量。
//   - vehicleUuid: 车辆UUID筛选条件，可为空。
//   - startDate: 开始日期筛选条件，可为空。
//   - endDate: 结束日期筛选条件，可为空。
//
// 返回:
//   - 年检记录列表。
//   - 总数量。
//   - 错误码的指针，表示可能的错误类型。
func (s *sVehicle) GetVehicleInspectionList(ctx context.Context, page int, size int, vehicleUuid string, startDate *gtime.Time, endDate *gtime.Time) ([]*entity.VehicleInspection, int, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetVehicleInspectionList", "获取车辆年检记录列表: page=%d, size=%d, vehicleUuid=%s", page, size, vehicleUuid)

	// 构建查询条件
	inspectionModel := dao.VehicleInspection.Ctx(ctx)
	if vehicleUuid != "" {
		inspectionModel = inspectionModel.Where("vehicle_uuid", vehicleUuid)
	}
	if startDate != nil {
		inspectionModel = inspectionModel.Where("inspection_date >= ?", startDate)
	}
	if endDate != nil {
		inspectionModel = inspectionModel.Where("inspection_date <= ?", endDate)
	}

	// 查询总数
	count, err := inspectionModel.Count()
	if err != nil {
		blog.ServiceError(ctx, "GetVehicleInspectionList", "获取车辆年检记录总数失败: %s", err.Error())
		return nil, 0, &berror.ErrDatabaseError
	}

	// 查询年检记录列表
	var inspections []*entity.VehicleInspection
	err = inspectionModel.Page(page, size).Order("inspection_date DESC").Scan(&inspections)
	if err != nil {
		blog.ServiceError(ctx, "GetVehicleInspectionList", "获取车辆年检记录列表失败: %s", err.Error())
		return nil, 0, &berror.ErrDatabaseError
	}

	// 返回结果
	return inspections, count, nil
}

// UpdateVehicleInspection 更新车辆年检记录。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - inspection: 年检记录实体，包含需要更新的年检详细信息。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sVehicle) UpdateVehicleInspection(ctx context.Context, inspection *entity.VehicleInspection) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "UpdateVehicleInspection", "更新车辆年检记录: %s", inspection.InspectionUuid)

	// 检查年检记录是否存在
	existingInspection, err := s.GetVehicleInspectionById(ctx, inspection.InspectionUuid)
	if err != nil {
		return err
	}
	if existingInspection == nil {
		return &berror.ErrNotFound
	}

	// 更新年检记录
	inspectionModel := dao.VehicleInspection.Ctx(ctx)
	_, sqlErr := inspectionModel.Where("inspection_uuid", inspection.InspectionUuid).OmitEmpty().Update(inspection)
	if sqlErr != nil {
		blog.ServiceError(ctx, "UpdateVehicleInspection", "更新车辆年检记录失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 如果有效期变更，更新车辆的下次年检日期
	if existingInspection.ExpiryDate.String() != inspection.ExpiryDate.String() {
		vehicle, getErr := s.GetVehicleById(ctx, inspection.VehicleUuid)
		if getErr == nil && vehicle != nil {
			vehicle.NextInspectionDate = inspection.ExpiryDate
			updateErr := s.UpdateVehicle(ctx, vehicle)
			if updateErr != nil {
				blog.ServiceError(ctx, "UpdateVehicleInspection", "更新车辆下次年检日期失败: %s", updateErr.Error())
			}
		}
	}

	// 返回结果
	return nil
}
