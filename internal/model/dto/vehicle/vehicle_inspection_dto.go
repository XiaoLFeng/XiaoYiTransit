package vehicle

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CreateVehicleInspectionReqDTO defines the request for creating a new vehicle inspection record
type CreateVehicleInspectionReqDTO struct {
	VehicleUuid      string      `json:"vehicle_uuid" v:"required#车辆UUID不能为空" dc:"车辆UUID"`
	InspectionDate   *gtime.Time `json:"inspection_date" v:"required#年检日期不能为空" dc:"年检日期"`
	ExpiryDate       *gtime.Time `json:"expiry_date" v:"required#有效期截止日期不能为空" dc:"有效期截止日期"`
	InspectionResult int         `json:"inspection_result" v:"required|in:1,2#年检结果不能为空|年检结果只能是1(通过)或2(不通过)" dc:"年检结果: 1-通过, 2-不通过"`
	InspectionAgency string      `json:"inspection_agency" dc:"检测机构"`
	Inspector        string      `json:"inspector" dc:"检测人员"`
	CertificateNo    string      `json:"certificate_no" dc:"检测证书编号"`
	Cost             float64     `json:"cost" dc:"年检费用"`
	Notes            string      `json:"notes" dc:"备注"`
}

// UpdateVehicleInspectionReqDTO defines the request for updating a vehicle inspection record
type UpdateVehicleInspectionReqDTO struct {
	InspectionUuid   string      `json:"inspection_uuid" v:"required#年检UUID不能为空" dc:"年检UUID"`
	VehicleUuid      string      `json:"vehicle_uuid" v:"required#车辆UUID不能为空" dc:"车辆UUID"`
	InspectionDate   *gtime.Time `json:"inspection_date" v:"required#年检日期不能为空" dc:"年检日期"`
	ExpiryDate       *gtime.Time `json:"expiry_date" v:"required#有效期截止日期不能为空" dc:"有效期截止日期"`
	InspectionResult int         `json:"inspection_result" v:"required|in:1,2#年检结果不能为空|年检结果只能是1(通过)或2(不通过)" dc:"年检结果: 1-通过, 2-不通过"`
	InspectionAgency string      `json:"inspection_agency" dc:"检测机构"`
	Inspector        string      `json:"inspector" dc:"检测人员"`
	CertificateNo    string      `json:"certificate_no" dc:"检测证书编号"`
	Cost             float64     `json:"cost" dc:"年检费用"`
	Notes            string      `json:"notes" dc:"备注"`
}

// GetVehicleInspectionListReqDTO defines the request for getting a list of vehicle inspection records
type GetVehicleInspectionListReqDTO struct {
	Page         int         `json:"page" d:"1" v:"min:1#页码最小为1" dc:"页码"`
	Size         int         `json:"size" d:"10" v:"max:100#每页最大100条" dc:"每页数量"`
	VehicleUuid  string      `json:"vehicle_uuid,omitempty" dc:"车辆UUID（可选，用于筛选）"`
	StartDate    *gtime.Time `json:"start_date,omitempty" dc:"开始日期（可选，用于筛选）"`
	EndDate      *gtime.Time `json:"end_date,omitempty" dc:"结束日期（可选，用于筛选）"`
}

// VehicleInspectionItemDTO defines a single vehicle inspection record in the list
type VehicleInspectionItemDTO struct {
	InspectionUuid   string      `json:"inspection_uuid" dc:"年检UUID"`
	VehicleUuid      string      `json:"vehicle_uuid" dc:"车辆UUID"`
	PlateNumber      string      `json:"plate_number" dc:"车牌号"`
	InspectionDate   *gtime.Time `json:"inspection_date" dc:"年检日期"`
	ExpiryDate       *gtime.Time `json:"expiry_date" dc:"有效期截止日期"`
	InspectionResult int         `json:"inspection_result" dc:"年检结果: 1-通过, 2-不通过"`
	InspectionAgency string      `json:"inspection_agency" dc:"检测机构"`
	CertificateNo    string      `json:"certificate_no" dc:"检测证书编号"`
	Cost             float64     `json:"cost" dc:"年检费用"`
	CreatedAt        *gtime.Time `json:"created_at" dc:"创建时间"`
}

// PagedVehicleInspectionListDTO defines a paged list of vehicle inspection records
type PagedVehicleInspectionListDTO struct {
	List  []*VehicleInspectionItemDTO `json:"list" dc:"年检记录列表"`
	Page  int                         `json:"page" dc:"当前页码"`
	Size  int                         `json:"size" dc:"每页数量"`
	Total int                         `json:"total" dc:"总数量"`
}

// VehicleInspectionDetailDTO defines the detailed vehicle inspection information
type VehicleInspectionDetailDTO struct {
	InspectionUuid   string      `json:"inspection_uuid" dc:"年检UUID"`
	VehicleUuid      string      `json:"vehicle_uuid" dc:"车辆UUID"`
	PlateNumber      string      `json:"plate_number" dc:"车牌号"`
	Model            string      `json:"model" dc:"车辆型号"`
	InspectionDate   *gtime.Time `json:"inspection_date" dc:"年检日期"`
	ExpiryDate       *gtime.Time `json:"expiry_date" dc:"有效期截止日期"`
	InspectionResult int         `json:"inspection_result" dc:"年检结果: 1-通过, 2-不通过"`
	InspectionAgency string      `json:"inspection_agency" dc:"检测机构"`
	Inspector        string      `json:"inspector" dc:"检测人员"`
	CertificateNo    string      `json:"certificate_no" dc:"检测证书编号"`
	Cost             float64     `json:"cost" dc:"年检费用"`
	Notes            string      `json:"notes" dc:"备注"`
	CreatedAt        *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updated_at" dc:"更新时间"`
}