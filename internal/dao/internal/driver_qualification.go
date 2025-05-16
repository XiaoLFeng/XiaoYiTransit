// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DriverQualificationDao is the data access object for the table xf_driver_qualification.
type DriverQualificationDao struct {
	table    string                     // table is the underlying table name of the DAO.
	group    string                     // group is the database configuration group name of the current DAO.
	columns  DriverQualificationColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler         // handlers for customized model modification.
}

// DriverQualificationColumns defines and stores column names for the table xf_driver_qualification.
type DriverQualificationColumns struct {
	Id                  string // ID
	DriverId            string // 司机ID
	QualificationType   string // 资质类型
	QualificationNumber string // 资质编号
	IssueDate           string // 发证日期
	ExpiryDate          string // 到期日期
	IssuingAuthority    string // 发证机构
	Notes               string // 备注
	CreatedAt           string // 创建时间
	UpdatedAt           string // 更新时间
}

// driverQualificationColumns holds the columns for the table xf_driver_qualification.
var driverQualificationColumns = DriverQualificationColumns{
	Id:                  "id",
	DriverId:            "driver_id",
	QualificationType:   "qualification_type",
	QualificationNumber: "qualification_number",
	IssueDate:           "issue_date",
	ExpiryDate:          "expiry_date",
	IssuingAuthority:    "issuing_authority",
	Notes:               "notes",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
}

// NewDriverQualificationDao creates and returns a new DAO object for table data access.
func NewDriverQualificationDao(handlers ...gdb.ModelHandler) *DriverQualificationDao {
	return &DriverQualificationDao{
		group:    "default",
		table:    "xf_driver_qualification",
		columns:  driverQualificationColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DriverQualificationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DriverQualificationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DriverQualificationDao) Columns() DriverQualificationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DriverQualificationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DriverQualificationDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DriverQualificationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
