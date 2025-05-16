// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DriverShiftDao is the data access object for the table xf_driver_shift.
type DriverShiftDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DriverShiftColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DriverShiftColumns defines and stores column names for the table xf_driver_shift.
type DriverShiftColumns struct {
	ShiftUuid  string // 排班UUID
	DriverUuid string // 司机UUID
	ShiftDate  string // 排班日期
	ShiftType  string // 班次类型: 1-早班, 2-中班, 3-晚班, 4-全天班
	StartTime  string // 开始时间
	EndTime    string // 结束时间
	Status     string // 状态: 0-取消, 1-待执行, 2-执行中, 3-已完成
	Notes      string // 备注
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// driverShiftColumns holds the columns for the table xf_driver_shift.
var driverShiftColumns = DriverShiftColumns{
	ShiftUuid:  "shift_uuid",
	DriverUuid: "driver_uuid",
	ShiftDate:  "shift_date",
	ShiftType:  "shift_type",
	StartTime:  "start_time",
	EndTime:    "end_time",
	Status:     "status",
	Notes:      "notes",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewDriverShiftDao creates and returns a new DAO object for table data access.
func NewDriverShiftDao(handlers ...gdb.ModelHandler) *DriverShiftDao {
	return &DriverShiftDao{
		group:    "default",
		table:    "xf_driver_shift",
		columns:  driverShiftColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DriverShiftDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DriverShiftDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DriverShiftDao) Columns() DriverShiftColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DriverShiftDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DriverShiftDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DriverShiftDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
