package vehicle

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CreateVehicleInsuranceReqDTO defines the request for creating a new vehicle insurance record
type CreateVehicleInsuranceReqDTO struct {
	VehicleUuid     string      `json:"vehicle_uuid" v:"required#车辆UUID不能为空" dc:"车辆UUID"`
	InsuranceType   string      `json:"insurance_type" v:"required#保险类型不能为空" dc:"保险类型"`
	PolicyNumber    string      `json:"policy_number" v:"required#保单号不能为空" dc:"保单号"`
	Insurer         string      `json:"insurer" dc:"保险公司"`
	StartDate       *gtime.Time `json:"start_date" v:"required#保险开始日期不能为空" dc:"保险开始日期"`
	ExpiryDate      *gtime.Time `json:"expiry_date" v:"required#保险到期日期不能为空" dc:"保险到期日期"`
	CoverageAmount  float64     `json:"coverage_amount" dc:"保险金额"`
	Premium         float64     `json:"premium" dc:"保费"`
	PaymentDate     *gtime.Time `json:"payment_date" dc:"缴费日期"`
	PaymentMethod   string      `json:"payment_method" dc:"缴费方式"`
	ContactPerson   string      `json:"contact_person" dc:"联系人"`
	ContactPhone    string      `json:"contact_phone" dc:"联系电话"`
	Notes           string      `json:"notes" dc:"备注"`
}

// UpdateVehicleInsuranceReqDTO defines the request for updating a vehicle insurance record
type UpdateVehicleInsuranceReqDTO struct {
	InsuranceUuid   string      `json:"insurance_uuid" v:"required#保险UUID不能为空" dc:"保险UUID"`
	VehicleUuid     string      `json:"vehicle_uuid" v:"required#车辆UUID不能为空" dc:"车辆UUID"`
	InsuranceType   string      `json:"insurance_type" v:"required#保险类型不能为空" dc:"保险类型"`
	PolicyNumber    string      `json:"policy_number" v:"required#保单号不能为空" dc:"保单号"`
	Insurer         string      `json:"insurer" dc:"保险公司"`
	StartDate       *gtime.Time `json:"start_date" v:"required#保险开始日期不能为空" dc:"保险开始日期"`
	ExpiryDate      *gtime.Time `json:"expiry_date" v:"required#保险到期日期不能为空" dc:"保险到期日期"`
	CoverageAmount  float64     `json:"coverage_amount" dc:"保险金额"`
	Premium         float64     `json:"premium" dc:"保费"`
	PaymentDate     *gtime.Time `json:"payment_date" dc:"缴费日期"`
	PaymentMethod   string      `json:"payment_method" dc:"缴费方式"`
	ContactPerson   string      `json:"contact_person" dc:"联系人"`
	ContactPhone    string      `json:"contact_phone" dc:"联系电话"`
	Notes           string      `json:"notes" dc:"备注"`
}

// GetVehicleInsuranceListReqDTO defines the request for getting a list of vehicle insurance records
type GetVehicleInsuranceListReqDTO struct {
	Page         int         `json:"page" d:"1" v:"min:1#页码最小为1" dc:"页码"`
	Size         int         `json:"size" d:"10" v:"max:100#每页最大100条" dc:"每页数量"`
	VehicleUuid  string      `json:"vehicle_uuid,omitempty" dc:"车辆UUID（可选，用于筛选）"`
	InsuranceType string     `json:"insurance_type,omitempty" dc:"保险类型（可选，用于筛选）"`
	StartDate    *gtime.Time `json:"start_date,omitempty" dc:"开始日期（可选，用于筛选）"`
	EndDate      *gtime.Time `json:"end_date,omitempty" dc:"结束日期（可选，用于筛选）"`
}

// VehicleInsuranceItemDTO defines a single vehicle insurance record in the list
type VehicleInsuranceItemDTO struct {
	InsuranceUuid   string      `json:"insurance_uuid" dc:"保险UUID"`
	VehicleUuid     string      `json:"vehicle_uuid" dc:"车辆UUID"`
	PlateNumber     string      `json:"plate_number" dc:"车牌号"`
	InsuranceType   string      `json:"insurance_type" dc:"保险类型"`
	PolicyNumber    string      `json:"policy_number" dc:"保单号"`
	Insurer         string      `json:"insurer" dc:"保险公司"`
	StartDate       *gtime.Time `json:"start_date" dc:"保险开始日期"`
	ExpiryDate      *gtime.Time `json:"expiry_date" dc:"保险到期日期"`
	CoverageAmount  float64     `json:"coverage_amount" dc:"保险金额"`
	Premium         float64     `json:"premium" dc:"保费"`
	CreatedAt       *gtime.Time `json:"created_at" dc:"创建时间"`
}

// PagedVehicleInsuranceListDTO defines a paged list of vehicle insurance records
type PagedVehicleInsuranceListDTO struct {
	List  []*VehicleInsuranceItemDTO `json:"list" dc:"保险记录列表"`
	Page  int                        `json:"page" dc:"当前页码"`
	Size  int                        `json:"size" dc:"每页数量"`
	Total int                        `json:"total" dc:"总数量"`
}

// VehicleInsuranceDetailDTO defines the detailed vehicle insurance information
type VehicleInsuranceDetailDTO struct {
	InsuranceUuid   string      `json:"insurance_uuid" dc:"保险UUID"`
	VehicleUuid     string      `json:"vehicle_uuid" dc:"车辆UUID"`
	PlateNumber     string      `json:"plate_number" dc:"车牌号"`
	Model           string      `json:"model" dc:"车辆型号"`
	InsuranceType   string      `json:"insurance_type" dc:"保险类型"`
	PolicyNumber    string      `json:"policy_number" dc:"保单号"`
	Insurer         string      `json:"insurer" dc:"保险公司"`
	StartDate       *gtime.Time `json:"start_date" dc:"保险开始日期"`
	ExpiryDate      *gtime.Time `json:"expiry_date" dc:"保险到期日期"`
	CoverageAmount  float64     `json:"coverage_amount" dc:"保险金额"`
	Premium         float64     `json:"premium" dc:"保费"`
	PaymentDate     *gtime.Time `json:"payment_date" dc:"缴费日期"`
	PaymentMethod   string      `json:"payment_method" dc:"缴费方式"`
	ContactPerson   string      `json:"contact_person" dc:"联系人"`
	ContactPhone    string      `json:"contact_phone" dc:"联系电话"`
	Notes           string      `json:"notes" dc:"备注"`
	CreatedAt       *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updated_at" dc:"更新时间"`
}