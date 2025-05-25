// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VehicleInsuranceDao is the data access object for the table xf_vehicle_insurance.
type VehicleInsuranceDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  VehicleInsuranceColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// VehicleInsuranceColumns defines and stores column names for the table xf_vehicle_insurance.
type VehicleInsuranceColumns struct {
	InsuranceUuid  string // 保险UUID
	VehicleUuid    string // 车辆UUID
	InsuranceType  string // 保险类型
	PolicyNumber   string // 保单号
	Insurer        string // 保险公司
	StartDate      string // 保险开始日期
	ExpiryDate     string // 保险到期日期
	CoverageAmount string // 保险金额
	Premium        string // 保费
	PaymentDate    string // 缴费日期
	PaymentMethod  string // 缴费方式
	ContactPerson  string // 联系人
	ContactPhone   string // 联系电话
	Notes          string // 备注
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 删除时间
}

// vehicleInsuranceColumns holds the columns for the table xf_vehicle_insurance.
var vehicleInsuranceColumns = VehicleInsuranceColumns{
	InsuranceUuid:  "insurance_uuid",
	VehicleUuid:    "vehicle_uuid",
	InsuranceType:  "insurance_type",
	PolicyNumber:   "policy_number",
	Insurer:        "insurer",
	StartDate:      "start_date",
	ExpiryDate:     "expiry_date",
	CoverageAmount: "coverage_amount",
	Premium:        "premium",
	PaymentDate:    "payment_date",
	PaymentMethod:  "payment_method",
	ContactPerson:  "contact_person",
	ContactPhone:   "contact_phone",
	Notes:          "notes",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewVehicleInsuranceDao creates and returns a new DAO object for table data access.
func NewVehicleInsuranceDao(handlers ...gdb.ModelHandler) *VehicleInsuranceDao {
	return &VehicleInsuranceDao{
		group:    "default",
		table:    "xf_vehicle_insurance",
		columns:  vehicleInsuranceColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *VehicleInsuranceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *VehicleInsuranceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *VehicleInsuranceDao) Columns() VehicleInsuranceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *VehicleInsuranceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *VehicleInsuranceDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *VehicleInsuranceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
