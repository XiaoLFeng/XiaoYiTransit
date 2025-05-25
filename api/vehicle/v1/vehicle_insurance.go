package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"go/types"
	"xiao-yi-transit/internal/model/dto/vehicle"
)

// CreateVehicleInsuranceReq defines the request for creating a new vehicle insurance record
type CreateVehicleInsuranceReq struct {
	g.Meta         `path:"/vehicle/insurance" method:"Post" tags:"车辆管理" sm:"创建保险记录" dc:"用于创建新的车辆保险记录"`
	VehicleUuid    string      `json:"vehicle_uuid" v:"required#车辆UUID不能为空" dc:"车辆UUID"`
	InsuranceType  string      `json:"insurance_type" v:"required#保险类型不能为空" dc:"保险类型"`
	PolicyNumber   string      `json:"policy_number" v:"required#保单号不能为空" dc:"保单号"`
	Insurer        string      `json:"insurer" dc:"保险公司"`
	StartDate      *gtime.Time `json:"start_date" v:"required#保险开始日期不能为空" dc:"保险开始日期"`
	ExpiryDate     *gtime.Time `json:"expiry_date" v:"required#保险到期日期不能为空" dc:"保险到期日期"`
	CoverageAmount float64     `json:"coverage_amount" dc:"保险金额"`
	Premium        float64     `json:"premium" dc:"保费"`
	PaymentDate    *gtime.Time `json:"payment_date" dc:"缴费日期"`
	PaymentMethod  string      `json:"payment_method" dc:"缴费方式"`
	ContactPerson  string      `json:"contact_person" dc:"联系人"`
	ContactPhone   string      `json:"contact_phone" dc:"联系电话"`
	Notes          string      `json:"notes" dc:"备注"`
}

// CreateVehicleInsuranceRes defines the response for creating a new vehicle insurance record
type CreateVehicleInsuranceRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*string]
}

// UpdateVehicleInsuranceReq defines the request for updating a vehicle insurance record
type UpdateVehicleInsuranceReq struct {
	g.Meta         `path:"/vehicle/insurance/{insuranceUuid}" method:"Put" tags:"车辆管理" sm:"更新保险记录" dc:"用于更新车辆保险记录"`
	InsuranceUuid  string      `json:"insurance_uuid" v:"required#保险UUID不能为空" dc:"保险UUID"`
	VehicleUuid    string      `json:"vehicle_uuid" v:"required#车辆UUID不能为空" dc:"车辆UUID"`
	InsuranceType  string      `json:"insurance_type" v:"required#保险类型不能为空" dc:"保险类型"`
	PolicyNumber   string      `json:"policy_number" v:"required#保单号不能为空" dc:"保单号"`
	Insurer        string      `json:"insurer" dc:"保险公司"`
	StartDate      *gtime.Time `json:"start_date" v:"required#保险开始日期不能为空" dc:"保险开始日期"`
	ExpiryDate     *gtime.Time `json:"expiry_date" v:"required#保险到期日期不能为空" dc:"保险到期日期"`
	CoverageAmount float64     `json:"coverage_amount" dc:"保险金额"`
	Premium        float64     `json:"premium" dc:"保费"`
	PaymentDate    *gtime.Time `json:"payment_date" dc:"缴费日期"`
	PaymentMethod  string      `json:"payment_method" dc:"缴费方式"`
	ContactPerson  string      `json:"contact_person" dc:"联系人"`
	ContactPhone   string      `json:"contact_phone" dc:"联系电话"`
	Notes          string      `json:"notes" dc:"备注"`
}

// UpdateVehicleInsuranceRes defines the response for updating a vehicle insurance record
type UpdateVehicleInsuranceRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}

// DeleteVehicleInsuranceReq defines the request for deleting a vehicle insurance record
type DeleteVehicleInsuranceReq struct {
	g.Meta        `path:"/vehicle/insurance/{insuranceUuid}" method:"Delete" tags:"车辆管理" sm:"删除保险记录" dc:"用于删除车辆保险记录"`
	InsuranceUuid string `json:"insurance_uuid" v:"required#保险UUID不能为空" dc:"保险UUID"`
}

// DeleteVehicleInsuranceRes defines the response for deleting a vehicle insurance record
type DeleteVehicleInsuranceRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}

// GetVehicleInsuranceListReq defines the request for getting a list of vehicle insurance records
type GetVehicleInsuranceListReq struct {
	g.Meta        `path:"/vehicle/insurances" method:"Get" tags:"车辆管理" sm:"获取保险记录列表" dc:"用于获取车辆保险记录列表"`
	Page          int         `json:"page" d:"1" v:"min:1#页码最小为1" dc:"页码"`
	Size          int         `json:"size" d:"10" v:"max:100#每页最大100条" dc:"每页数量"`
	VehicleUuid   string      `json:"vehicle_uuid,omitempty" dc:"车辆UUID（可选，用于筛选）"`
	InsuranceType string      `json:"insurance_type,omitempty" dc:"保险类型（可选，用于筛选）"`
	StartDate     *gtime.Time `json:"start_date,omitempty" dc:"开始日期（可选，用于筛选）"`
	EndDate       *gtime.Time `json:"end_date,omitempty" dc:"结束日期（可选，用于筛选）"`
}

// GetVehicleInsuranceListRes defines the response for getting a list of vehicle insurance records
type GetVehicleInsuranceListRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*vehicle.PagedVehicleInsuranceListDTO]
}
