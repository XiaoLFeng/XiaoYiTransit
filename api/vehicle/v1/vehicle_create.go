package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"go/types"
)

// CreateVehicleReq defines the request for creating a new vehicle
type CreateVehicleReq struct {
	g.Meta             `path:"/vehicle" method:"Post" tags:"车辆管理" sm:"创建车辆" dc:"用于创建新的车辆信息"`
	PlateNumber        string      `json:"plate_number" v:"required#车牌号不能为空" dc:"车牌号"`
	Model              string      `json:"model" v:"required#车辆型号不能为空" dc:"车辆型号"`
	PurchaseDate       *gtime.Time `json:"purchase_date" v:"required#购买日期不能为空" dc:"购买日期"`
	Status             int         `json:"status" v:"required|in:1,2,3,4#状态不能为空|状态只能是1(运营),2(维修),3(停运),4(报废)" dc:"状态: 1-运营, 2-维修, 3-停运, 4-报废"`
	Seats              int         `json:"seats" v:"required#座位数不能为空" dc:"座位数"`
	FuelType           string      `json:"fuel_type" dc:"燃料类型"`
	Mileage            float64     `json:"mileage" dc:"行驶里程(km)"`
	LastMaintenanceDate *gtime.Time `json:"last_maintenance_date" dc:"最后维护日期"`
	NextInspectionDate *gtime.Time `json:"next_inspection_date" dc:"下次年检日期"`
	InsuranceExpiryDate *gtime.Time `json:"insurance_expiry_date" dc:"保险到期日期"`
	Notes              string      `json:"notes" dc:"备注"`
}

// CreateVehicleRes defines the response for creating a new vehicle
type CreateVehicleRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}
