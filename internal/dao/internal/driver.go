// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DriverDao is the data access object for the table xf_driver.
type DriverDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DriverColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DriverColumns defines and stores column names for the table xf_driver.
type DriverColumns struct {
	Id                string // 司机ID
	EmployeeId        string // 工号
	Name              string // 姓名
	Gender            string // 性别: 1-男, 2-女
	IdCard            string // 身份证号
	Phone             string // 联系电话
	EmergencyContact  string // 紧急联系人
	EmergencyPhone    string // 紧急联系电话
	LicenseNumber     string // 驾驶证号
	LicenseType       string // 驾驶证类型
	LicenseExpiryDate string // 驾驶证到期日期
	EntryDate         string // 入职日期
	Status            string // 状态: 0-离职, 1-在职, 2-休假, 3-停职
	Address           string // 住址
	Notes             string // 备注
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
	DeletedAt         string // 删除时间
}

// driverColumns holds the columns for the table xf_driver.
var driverColumns = DriverColumns{
	Id:                "id",
	EmployeeId:        "employee_id",
	Name:              "name",
	Gender:            "gender",
	IdCard:            "id_card",
	Phone:             "phone",
	EmergencyContact:  "emergency_contact",
	EmergencyPhone:    "emergency_phone",
	LicenseNumber:     "license_number",
	LicenseType:       "license_type",
	LicenseExpiryDate: "license_expiry_date",
	EntryDate:         "entry_date",
	Status:            "status",
	Address:           "address",
	Notes:             "notes",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	DeletedAt:         "deleted_at",
}

// NewDriverDao creates and returns a new DAO object for table data access.
func NewDriverDao(handlers ...gdb.ModelHandler) *DriverDao {
	return &DriverDao{
		group:    "default",
		table:    "xf_driver",
		columns:  driverColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DriverDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DriverDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DriverDao) Columns() DriverColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DriverDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DriverDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DriverDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
