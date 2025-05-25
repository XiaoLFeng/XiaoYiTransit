// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VehicleInspectionDao is the data access object for the table xf_vehicle_inspection.
type VehicleInspectionDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  VehicleInspectionColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// VehicleInspectionColumns defines and stores column names for the table xf_vehicle_inspection.
type VehicleInspectionColumns struct {
	InspectionUuid   string // 年检UUID
	VehicleUuid      string // 车辆UUID
	InspectionDate   string // 年检日期
	ExpiryDate       string // 有效期截止日期
	InspectionResult string // 年检结果: 1-通过, 2-不通过
	InspectionAgency string // 检测机构
	Inspector        string // 检测人员
	CertificateNo    string // 检测证书编号
	Cost             string // 年检费用
	Notes            string // 备注
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
	DeletedAt        string // 删除时间
}

// vehicleInspectionColumns holds the columns for the table xf_vehicle_inspection.
var vehicleInspectionColumns = VehicleInspectionColumns{
	InspectionUuid:   "inspection_uuid",
	VehicleUuid:      "vehicle_uuid",
	InspectionDate:   "inspection_date",
	ExpiryDate:       "expiry_date",
	InspectionResult: "inspection_result",
	InspectionAgency: "inspection_agency",
	Inspector:        "inspector",
	CertificateNo:    "certificate_no",
	Cost:             "cost",
	Notes:            "notes",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

// NewVehicleInspectionDao creates and returns a new DAO object for table data access.
func NewVehicleInspectionDao(handlers ...gdb.ModelHandler) *VehicleInspectionDao {
	return &VehicleInspectionDao{
		group:    "default",
		table:    "xf_vehicle_inspection",
		columns:  vehicleInspectionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *VehicleInspectionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *VehicleInspectionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *VehicleInspectionDao) Columns() VehicleInspectionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *VehicleInspectionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *VehicleInspectionDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *VehicleInspectionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
