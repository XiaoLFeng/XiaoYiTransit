// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MaintenanceItemDao is the data access object for the table xf_maintenance_item.
type MaintenanceItemDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  MaintenanceItemColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// MaintenanceItemColumns defines and stores column names for the table xf_maintenance_item.
type MaintenanceItemColumns struct {
	MaintenanceItemUuid string // 项目UUID
	MaintenanceUuid     string // 维修记录UUID
	ItemName            string // 项目名称
	ItemCost            string // 项目费用
	ItemDescription     string // 项目描述
	CreatedAt           string // 创建时间
	UpdatedAt           string // 更新时间
}

// maintenanceItemColumns holds the columns for the table xf_maintenance_item.
var maintenanceItemColumns = MaintenanceItemColumns{
	MaintenanceItemUuid: "maintenance_item_uuid",
	MaintenanceUuid:     "maintenance_uuid",
	ItemName:            "item_name",
	ItemCost:            "item_cost",
	ItemDescription:     "item_description",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
}

// NewMaintenanceItemDao creates and returns a new DAO object for table data access.
func NewMaintenanceItemDao(handlers ...gdb.ModelHandler) *MaintenanceItemDao {
	return &MaintenanceItemDao{
		group:    "default",
		table:    "xf_maintenance_item",
		columns:  maintenanceItemColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MaintenanceItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MaintenanceItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MaintenanceItemDao) Columns() MaintenanceItemColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MaintenanceItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MaintenanceItemDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MaintenanceItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
