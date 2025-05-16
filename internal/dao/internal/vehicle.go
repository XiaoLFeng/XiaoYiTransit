// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VehicleDao is the data access object for the table xf_vehicle.
type VehicleDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  VehicleColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// VehicleColumns defines and stores column names for the table xf_vehicle.
type VehicleColumns struct {
	Id                  string // 车辆ID
	PlateNumber         string // 车牌号
	Model               string // 车辆型号
	PurchaseDate        string // 购买日期
	Status              string // 状态: 1-运营, 2-维修, 3-停运, 4-报废
	Seats               string // 座位数
	FuelType            string // 燃料类型
	Mileage             string // 行驶里程(km)
	LastMaintenanceDate string // 最后维护日期
	NextInspectionDate  string // 下次年检日期
	InsuranceExpiryDate string // 保险到期日期
	Notes               string // 备注
	CreatedAt           string // 创建时间
	UpdatedAt           string // 更新时间
	DeletedAt           string // 删除时间
}

// vehicleColumns holds the columns for the table xf_vehicle.
var vehicleColumns = VehicleColumns{
	Id:                  "id",
	PlateNumber:         "plate_number",
	Model:               "model",
	PurchaseDate:        "purchase_date",
	Status:              "status",
	Seats:               "seats",
	FuelType:            "fuel_type",
	Mileage:             "mileage",
	LastMaintenanceDate: "last_maintenance_date",
	NextInspectionDate:  "next_inspection_date",
	InsuranceExpiryDate: "insurance_expiry_date",
	Notes:               "notes",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	DeletedAt:           "deleted_at",
}

// NewVehicleDao creates and returns a new DAO object for table data access.
func NewVehicleDao(handlers ...gdb.ModelHandler) *VehicleDao {
	return &VehicleDao{
		group:    "default",
		table:    "xf_vehicle",
		columns:  vehicleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *VehicleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *VehicleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *VehicleDao) Columns() VehicleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *VehicleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *VehicleDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *VehicleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
