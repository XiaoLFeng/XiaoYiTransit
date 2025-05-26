package maintenance

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GetMaintenanceListReqDTO defines the request for getting a list of maintenance records
type GetMaintenanceListReqDTO struct {
	Page            int         `json:"page" d:"1" v:"min:1#页码最小为1" dc:"页码"`
	Size            int         `json:"size" d:"10" v:"max:100#每页最大100条" dc:"每页数量"`
	VehicleUuid     string      `json:"vehicle_uuid,omitempty" dc:"车辆UUID（可选，用于筛选）"`
	MaintenanceType int         `json:"maintenance_type,omitempty" dc:"维修类型（可选，用于筛选）: 1-常规保养, 2-故障维修, 3-事故维修, 4-年检维修"`
	Status          int         `json:"status,omitempty" dc:"状态（可选，用于筛选）: 0-已取消, 1-待维修, 2-维修中, 3-已完成"`
	StartDate       *gtime.Time `json:"start_date,omitempty" dc:"开始日期（可选，用于筛选）"`
	EndDate         *gtime.Time `json:"end_date,omitempty" dc:"结束日期（可选，用于筛选）"`
}

// MaintenanceListItemDTO defines a single maintenance item in the list
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

// PagedMaintenanceListDTO defines a paged list of maintenance records
type PagedMaintenanceListDTO struct {
	List  []*MaintenanceListItemDTO `json:"list" dc:"维修记录列表"`
	Page  int                       `json:"page" dc:"当前页码"`
	Size  int                       `json:"size" dc:"每页数量"`
	Total int                       `json:"total" dc:"总数量"`
}

// MaintenanceDetailItemDTO defines the detailed maintenance information
type MaintenanceDetailItemDTO struct {
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