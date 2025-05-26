package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"go/types"
)

// UpdateMaintenanceReq defines the request for updating a maintenance record
type UpdateMaintenanceReq struct {
	g.Meta             `path:"/maintenance/{maintenance_uuid}" method:"Put" tags:"维修管理" sm:"更新维修记录" dc:"用于更新车辆维修记录"`
	MaintenanceUuid    string      `json:"maintenance_uuid" v:"required#维修记录UUID不能为空" dc:"维修记录UUID" in:"path"`
	VehicleUuid        string      `json:"vehicle_uuid" v:"required#车辆UUID不能为空" dc:"车辆UUID"`
	MaintenanceType    int         `json:"maintenance_type" v:"required|in:1,2,3,4#维修类型不能为空|维修类型只能是1(常规保养),2(故障维修),3(事故维修),4(年检维修)" dc:"维修类型: 1-常规保养, 2-故障维修, 3-事故维修, 4-年检维修"`
	Description        string      `json:"description" v:"required#维修描述不能为空" dc:"维修描述"`
	MaintenanceDate    *gtime.Time `json:"maintenance_date" v:"required#维修日期不能为空" dc:"维修日期"`
	CompletionDate     *gtime.Time `json:"completion_date" dc:"完成日期"`
	Cost               float64     `json:"cost" dc:"维修费用"`
	Mileage            float64     `json:"mileage" dc:"维修时里程数"`
	MaintenanceLocation string      `json:"maintenance_location" dc:"维修地点"`
	MaintenanceStaff   string      `json:"maintenance_staff" dc:"维修人员"`
	PartsReplaced      string      `json:"parts_replaced" dc:"更换的零部件"`
	Status             int         `json:"status" v:"required|in:0,1,2,3#状态不能为空|状态只能是0(已取消),1(待维修),2(维修中),3(已完成)" dc:"状态: 0-已取消, 1-待维修, 2-维修中, 3-已完成"`
	Notes              string      `json:"notes" dc:"备注"`
}

// UpdateMaintenanceRes defines the response for updating a maintenance record
type UpdateMaintenanceRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}