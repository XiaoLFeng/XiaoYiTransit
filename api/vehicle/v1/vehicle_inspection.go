package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"go/types"
	"xiao-yi-transit/internal/model/dto/vehicle"
)

// CreateVehicleInspectionReq defines the request for creating a new vehicle inspection record
type CreateVehicleInspectionReq struct {
	g.Meta           `path:"/vehicle/inspection" method:"Post" tags:"车辆管理" sm:"创建年检记录" dc:"用于创建新的车辆年检记录"`
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

// CreateVehicleInspectionRes defines the response for creating a new vehicle inspection record
type CreateVehicleInspectionRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*string]
}

// UpdateVehicleInspectionReq defines the request for updating a vehicle inspection record
type UpdateVehicleInspectionReq struct {
	g.Meta           `path:"/vehicle/inspection/{inspectionUuid}" method:"Put" tags:"车辆管理" sm:"更新年检记录" dc:"用于更新车辆年检记录"`
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

// UpdateVehicleInspectionRes defines the response for updating a vehicle inspection record
type UpdateVehicleInspectionRes struct {
	g.Meta                           `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}

// DeleteVehicleInspectionReq defines the request for deleting a vehicle inspection record
type DeleteVehicleInspectionReq struct {
	g.Meta         `path:"/vehicle/inspection/{inspectionUuid}" method:"Delete" tags:"车辆管理" sm:"删除年检记录" dc:"用于删除车辆年检记录"`
	InspectionUuid string `json:"inspection_uuid" v:"required#年检UUID不能为空" dc:"年检UUID"`
}

// DeleteVehicleInspectionRes defines the response for deleting a vehicle inspection record
type DeleteVehicleInspectionRes struct {
	g.Meta                           `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}

// GetVehicleInspectionListReq defines the request for getting a list of vehicle inspection records
type GetVehicleInspectionListReq struct {
	g.Meta      `path:"/vehicle/inspections" method:"Get" tags:"车辆管理" sm:"获取年检记录列表" dc:"用于获取车辆年检记录列表"`
	Page        int         `json:"page" d:"1" v:"min:1#页码最小为1" dc:"页码"`
	Size        int         `json:"size" d:"10" v:"max:100#每页最大100条" dc:"每页数量"`
	VehicleUuid string      `json:"vehicle_uuid,omitempty" dc:"车辆UUID（可选，用于筛选）"`
	StartDate   *gtime.Time `json:"start_date,omitempty" dc:"开始日期（可选，用于筛选）"`
	EndDate     *gtime.Time `json:"end_date,omitempty" dc:"结束日期（可选，用于筛选）"`
}

// GetVehicleInspectionListRes defines the response for getting a list of vehicle inspection records
type GetVehicleInspectionListRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*vehicle.PagedVehicleInspectionListDTO]
}
