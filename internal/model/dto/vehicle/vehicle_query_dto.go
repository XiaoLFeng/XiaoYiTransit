package vehicle

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SimpleVehicleItemDTO defines a single vehicle item in the simple list
type SimpleVehicleItemDTO struct {
	VehicleUuid string `json:"vehicle_uuid" dc:"车辆UUID"`
	PlateNumber string `json:"plate_number" dc:"车牌号"`
	Model       string `json:"model" dc:"车辆型号"`
}

// SimpleVehicleListDTO defines a simple list of vehicles
type SimpleVehicleListDTO struct {
	List []*SimpleVehicleItemDTO `json:"list" dc:"车辆列表"`
}

// GetVehicleListReqDTO defines the request for getting a list of vehicles
type GetVehicleListReqDTO struct {
	Page        int    `json:"page" d:"1" v:"min:1#页码最小为1" dc:"页码"`
	Size        int    `json:"size" d:"10" v:"max:100#每页最大100条" dc:"每页数量"`
	PlateNumber string `json:"plate_number,omitempty" dc:"车牌号（可选，用于筛选）"`
	Model       string `json:"model,omitempty" dc:"车辆型号（可选，用于筛选）"`
	Status      int    `json:"status,omitempty" dc:"状态（可选，用于筛选）: 1-运营, 2-维修, 3-停运, 4-报废"`
}

// VehicleListItemDTO defines a single vehicle item in the list
type VehicleListItemDTO struct {
	VehicleUuid         string      `json:"vehicle_uuid" dc:"车辆UUID"`
	PlateNumber         string      `json:"plate_number" dc:"车牌号"`
	Model               string      `json:"model" dc:"车辆型号"`
	PurchaseDate        *gtime.Time `json:"purchase_date" dc:"购买日期"`
	Status              int         `json:"status" dc:"状态: 1-运营, 2-维修, 3-停运, 4-报废"`
	Seats               int         `json:"seats" dc:"座位数"`
	Mileage             float64     `json:"mileage" dc:"行驶里程(km)"`
	NextInspectionDate  *gtime.Time `json:"next_inspection_date" dc:"下次年检日期"`
	InsuranceExpiryDate *gtime.Time `json:"insurance_expiry_date" dc:"保险到期日期"`
}

// PagedVehicleListDTO defines a paged list of vehicles
type PagedVehicleListDTO struct {
	List  []*VehicleListItemDTO `json:"list" dc:"车辆列表"`
	Page  int                   `json:"page" dc:"当前页码"`
	Size  int                   `json:"size" dc:"每页数量"`
	Total int                   `json:"total" dc:"总数量"`
}

// VehicleDetailItemDTO defines the detailed vehicle information
type VehicleDetailItemDTO struct {
	VehicleUuid         string      `json:"vehicle_uuid" dc:"车辆UUID"`
	PlateNumber         string      `json:"plate_number" dc:"车牌号"`
	Model               string      `json:"model" dc:"车辆型号"`
	PurchaseDate        *gtime.Time `json:"purchase_date" dc:"购买日期"`
	Status              int         `json:"status" dc:"状态: 1-运营, 2-维修, 3-停运, 4-报废"`
	Seats               int         `json:"seats" dc:"座位数"`
	FuelType            string      `json:"fuel_type" dc:"燃料类型"`
	Mileage             float64     `json:"mileage" dc:"行驶里程(km)"`
	LastMaintenanceDate *gtime.Time `json:"last_maintenance_date" dc:"最后维护日期"`
	NextInspectionDate  *gtime.Time `json:"next_inspection_date" dc:"下次年检日期"`
	InsuranceExpiryDate *gtime.Time `json:"insurance_expiry_date" dc:"保险到期日期"`
	Notes               string      `json:"notes" dc:"备注"`
	CreatedAt           *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt           *gtime.Time `json:"updated_at" dc:"更新时间"`
}
