package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"xiao-yi-transit/internal/model/dto/maintenance"
)

// MaintenanceDTO defines the data transfer object for maintenance records
type MaintenanceDTO struct {
	MaintenanceUuid     string      `json:"maintenance_uuid" dc:"维修UUID"`
	VehicleUuid         string      `json:"vehicle_uuid" dc:"车辆UUID"`
	PlateNumber         string      `json:"plate_number" dc:"车牌号"`
	MaintenanceType     int         `json:"maintenance_type" dc:"维修类型: 1-常规保养, 2-故障维修, 3-事故维修, 4-年检维修"`
	MaintenanceTypeName string      `json:"maintenance_type_name" dc:"维修类型名称"`
	Description         string      `json:"description" dc:"维修描述"`
	MaintenanceDate     *gtime.Time `json:"maintenance_date" dc:"维修日期"`
	CompletionDate      *gtime.Time `json:"completion_date" dc:"完成日期"`
	Cost                float64     `json:"cost" dc:"维修费用"`
	Mileage             float64     `json:"mileage" dc:"维修时里程数"`
	MaintenanceLocation string      `json:"maintenance_location" dc:"维修地点"`
	MaintenanceStaff    string      `json:"maintenance_staff" dc:"维修人员"`
	PartsReplaced       string      `json:"parts_replaced" dc:"更换的零部件"`
	Status              int         `json:"status" dc:"状态: 0-已取消, 1-待维修, 2-维修中, 3-已完成"`
	StatusName          string      `json:"status_name" dc:"状态名称"`
	Notes               string      `json:"notes" dc:"备注"`
	CreatedAt           *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt           *gtime.Time `json:"updated_at" dc:"更新时间"`
}

// GetMaintenanceDetailReq defines the request for getting a maintenance record detail
type GetMaintenanceDetailReq struct {
	g.Meta          `path:"/maintenance/{maintenance_uuid}" method:"Get" tags:"维修管理" sm:"获取维修记录详情" dc:"用于获取车辆维修记录详情"`
	MaintenanceUuid string `json:"maintenance_uuid" v:"required#维修记录UUID不能为空" dc:"维修记录UUID" in:"path"`
}

// GetMaintenanceDetailRes defines the response for getting a maintenance record detail
type GetMaintenanceDetailRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*MaintenanceDTO]
}

// GetMaintenanceListReq defines the request for getting a list of maintenance records
type GetMaintenanceListReq struct {
	g.Meta          `path:"/maintenance" method:"Get" tags:"维修管理" sm:"获取维修记录列表" dc:"用于获取车辆维修记录列表"`
	Page            int         `json:"page" v:"min:1#页码最小为1" d:"1" dc:"页码，从1开始"`
	Size            int         `json:"size" v:"min:1#每页数量最小为1" d:"10" dc:"每页数量"`
	VehicleUuid     string      `json:"vehicle_uuid" dc:"车辆UUID，可为空"`
	MaintenanceType int         `json:"maintenance_type" dc:"维修类型，可为0（表示不筛选）"`
	Status          int         `json:"status" dc:"状态，可为0（表示不筛选）"`
	StartDate       *gtime.Time `json:"start_date" dc:"开始日期，可为空"`
	EndDate         *gtime.Time `json:"end_date" dc:"结束日期，可为空"`
}

// MaintenanceListItemDTO defines the data transfer object for maintenance list items
type MaintenanceListItemDTO struct {
	MaintenanceUuid     string      `json:"maintenance_uuid" dc:"维修UUID"`
	VehicleUuid         string      `json:"vehicle_uuid" dc:"车辆UUID"`
	PlateNumber         string      `json:"plate_number" dc:"车牌号"`
	MaintenanceType     int         `json:"maintenance_type" dc:"维修类型: 1-常规保养, 2-故障维修, 3-事故维修, 4-年检维修"`
	MaintenanceTypeName string      `json:"maintenance_type_name" dc:"维修类型名称"`
	Description         string      `json:"description" dc:"维修描述"`
	MaintenanceDate     *gtime.Time `json:"maintenance_date" dc:"维修日期"`
	CompletionDate      *gtime.Time `json:"completion_date" dc:"完成日期"`
	Cost                float64     `json:"cost" dc:"维修费用"`
	Status              int         `json:"status" dc:"状态: 0-已取消, 1-待维修, 2-维修中, 3-已完成"`
	StatusName          string      `json:"status_name" dc:"状态名称"`
}

// GetMaintenanceListRes defines the response for getting a list of maintenance records
type GetMaintenanceListRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*maintenance.PagedMaintenanceListDTO]
}
