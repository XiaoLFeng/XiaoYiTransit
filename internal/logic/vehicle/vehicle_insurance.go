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

// CreateVehicleInsurance 创建新的车辆保险记录。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - insurance: 车辆保险记录实体，包含需要创建的保险详细信息。
//
// 返回:
//   - 创建成功的保险记录UUID。
//   - 错误码的指针，表示可能的错误类型。
func (s *sVehicle) CreateVehicleInsurance(ctx context.Context, insurance *entity.VehicleInsurance) (string, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "CreateVehicleInsurance", "创建车辆保险记录: 车辆UUID=%s", insurance.VehicleUuid)

	// 检查车辆是否存在
	vehicle, err := s.GetVehicleById(ctx, insurance.VehicleUuid)
	if err != nil {
		return "", err
	}
	if vehicle == nil {
		return "", &berror.ErrNotFound
	}

	// 生成UUID
	insurance.InsuranceUuid = uuid.New().String()

	// 创建保险记录
	insuranceModel := dao.VehicleInsurance.Ctx(ctx)
	_, sqlErr := insuranceModel.Insert(insurance)
	if sqlErr != nil {
		blog.ServiceError(ctx, "CreateVehicleInsurance", "创建车辆保险记录失败: %s", sqlErr.Error())
		return "", &berror.ErrDatabaseError
	}

	// 更新车辆的保险到期日期
	vehicle.InsuranceExpiryDate = insurance.ExpiryDate
	updateErr := s.UpdateVehicle(ctx, vehicle)
	if updateErr != nil {
		blog.ServiceError(ctx, "CreateVehicleInsurance", "更新车辆保险到期日期失败: %s", updateErr.Error())
	}

	// 返回结果
	return insurance.InsuranceUuid, nil
}

// DeleteVehicleInsurance 删除车辆保险记录。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - insuranceUuid: 保险记录的唯一标识符。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sVehicle) DeleteVehicleInsurance(ctx context.Context, insuranceUuid string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "DeleteVehicleInsurance", "删除车辆保险记录: %s", insuranceUuid)

	// 检查保险记录是否存在
	insurance, err := s.GetVehicleInsuranceById(ctx, insuranceUuid)
	if err != nil {
		return err
	}
	if insurance == nil {
		return &berror.ErrNotFound
	}

	// 删除保险记录（软删除）
	insuranceModel := dao.VehicleInsurance.Ctx(ctx)
	_, sqlErr := insuranceModel.Where("insurance_uuid", insuranceUuid).Delete()
	if sqlErr != nil {
		blog.ServiceError(ctx, "DeleteVehicleInsurance", "删除车辆保险记录失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 返回结果
	return nil
}

// GetVehicleInsuranceById 根据保险记录UUID获取保险记录信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - insuranceUuid: 保险记录的唯一标识符。
//
// 返回:
//   - 保险记录信息的指针，包含详细保险数据。
//   - 错误码的指针，表示错误类型，如保险记录不存在或内部错误。
func (s *sVehicle) GetVehicleInsuranceById(ctx context.Context, insuranceUuid string) (*entity.VehicleInsurance, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetVehicleInsuranceById", "获取车辆保险记录: %s", insuranceUuid)

	// 查询保险记录信息
	insuranceModel := dao.VehicleInsurance.Ctx(ctx)
	insurance := &entity.VehicleInsurance{}
	err := insuranceModel.Where("insurance_uuid", insuranceUuid).Scan(insurance)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, &berror.ErrNotFound
		}
		blog.ServiceError(ctx, "GetVehicleInsuranceById", "获取车辆保险记录失败: %s", err.Error())
		return nil, &berror.ErrDatabaseError
	}

	// 返回结果
	return insurance, nil
}

// GetVehicleInsuranceList 获取车辆保险记录列表。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - page: 页码，从1开始。
//   - size: 每页数量。
//   - vehicleUuid: 车辆UUID筛选条件，可为空。
//   - insuranceType: 保险类型筛选条件，可为空。
//   - startDate: 开始日期筛选条件，可为空。
//   - endDate: 结束日期筛选条件，可为空。
//
// 返回:
//   - 保险记录列表。
//   - 总数量。
//   - 错误码的指针，表示可能的错误类型。
func (s *sVehicle) GetVehicleInsuranceList(ctx context.Context, page int, size int, vehicleUuid string, insuranceType string, startDate *gtime.Time, endDate *gtime.Time) ([]*entity.VehicleInsurance, int, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetVehicleInsuranceList", "获取车辆保险记录列表: page=%d, size=%d, vehicleUuid=%s, insuranceType=%s", page, size, vehicleUuid, insuranceType)

	// 构建查询条件
	insuranceModel := dao.VehicleInsurance.Ctx(ctx)
	if vehicleUuid != "" {
		insuranceModel = insuranceModel.Where("vehicle_uuid", vehicleUuid)
	}
	if insuranceType != "" {
		insuranceModel = insuranceModel.Where("insurance_type", insuranceType)
	}
	if startDate != nil {
		insuranceModel = insuranceModel.Where("start_date >= ?", startDate)
	}
	if endDate != nil {
		insuranceModel = insuranceModel.Where("expiry_date <= ?", endDate)
	}

	// 查询总数
	count, err := insuranceModel.Count()
	if err != nil {
		blog.ServiceError(ctx, "GetVehicleInsuranceList", "获取车辆保险记录总数失败: %s", err.Error())
		return nil, 0, &berror.ErrDatabaseError
	}

	// 查询保险记录列表
	var insurances []*entity.VehicleInsurance
	err = insuranceModel.Page(page, size).Order("start_date DESC").Scan(&insurances)
	if err != nil {
		blog.ServiceError(ctx, "GetVehicleInsuranceList", "获取车辆保险记录列表失败: %s", err.Error())
		return nil, 0, &berror.ErrDatabaseError
	}

	// 返回结果
	return insurances, count, nil
}

// UpdateVehicleInsurance 更新车辆保险记录。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - insurance: 保险记录实体，包含需要更新的保险详细信息。
//
// 返回:
//   - 错误码的指针，表示可能的错误类型。
func (s *sVehicle) UpdateVehicleInsurance(ctx context.Context, insurance *entity.VehicleInsurance) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "UpdateVehicleInsurance", "更新车辆保险记录: %s", insurance.InsuranceUuid)

	// 检查保险记录是否存在
	existingInsurance, err := s.GetVehicleInsuranceById(ctx, insurance.InsuranceUuid)
	if err != nil {
		return err
	}
	if existingInsurance == nil {
		return &berror.ErrNotFound
	}

	// 更新保险记录
	insuranceModel := dao.VehicleInsurance.Ctx(ctx)
	_, sqlErr := insuranceModel.Where("insurance_uuid", insurance.InsuranceUuid).OmitEmpty().Update(insurance)
	if sqlErr != nil {
		blog.ServiceError(ctx, "UpdateVehicleInsurance", "更新车辆保险记录失败: %s", sqlErr.Error())
		return &berror.ErrDatabaseError
	}

	// 如果有效期变更，更新车辆的保险到期日期
	if existingInsurance.ExpiryDate.String() != insurance.ExpiryDate.String() {
		vehicle, getErr := s.GetVehicleById(ctx, insurance.VehicleUuid)
		if getErr == nil && vehicle != nil {
			vehicle.InsuranceExpiryDate = insurance.ExpiryDate
			updateErr := s.UpdateVehicle(ctx, vehicle)
			if updateErr != nil {
				blog.ServiceError(ctx, "UpdateVehicleInsurance", "更新车辆保险到期日期失败: %s", updateErr.Error())
			}
		}
	}

	// 返回结果
	return nil
}
