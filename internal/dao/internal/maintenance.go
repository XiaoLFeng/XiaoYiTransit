// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MaintenanceDao is the data access object for the table xf_maintenance.
type MaintenanceDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MaintenanceColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MaintenanceColumns defines and stores column names for the table xf_maintenance.
type MaintenanceColumns struct {
	Id                  string // 维修ID
	VehicleId           string // 车辆ID
	MaintenanceType     string // 维修类型: 1-常规保养, 2-故障维修, 3-事故维修, 4-年检维修
	Description         string // 维修描述
	MaintenanceDate     string // 维修日期
	CompletionDate      string // 完成日期
	Cost                string // 维修费用
	Mileage             string // 维修时里程数
	MaintenanceLocation string // 维修地点
	MaintenanceStaff    string // 维修人员
	PartsReplaced       string // 更换的零部件
	Status              string // 状态: 0-已取消, 1-待维修, 2-维修中, 3-已完成
	Notes               string // 备注
	CreatedBy           string // 创建人ID
	CreatedAt           string // 创建时间
	UpdatedAt           string // 更新时间
	DeletedAt           string // 删除时间
}

// maintenanceColumns holds the columns for the table xf_maintenance.
var maintenanceColumns = MaintenanceColumns{
	Id:                  "id",
	VehicleId:           "vehicle_id",
	MaintenanceType:     "maintenance_type",
	Description:         "description",
	MaintenanceDate:     "maintenance_date",
	CompletionDate:      "completion_date",
	Cost:                "cost",
	Mileage:             "mileage",
	MaintenanceLocation: "maintenance_location",
	MaintenanceStaff:    "maintenance_staff",
	PartsReplaced:       "parts_replaced",
	Status:              "status",
	Notes:               "notes",
	CreatedBy:           "created_by",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	DeletedAt:           "deleted_at",
}

// NewMaintenanceDao creates and returns a new DAO object for table data access.
func NewMaintenanceDao(handlers ...gdb.ModelHandler) *MaintenanceDao {
	return &MaintenanceDao{
		group:    "default",
		table:    "xf_maintenance",
		columns:  maintenanceColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MaintenanceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MaintenanceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MaintenanceDao) Columns() MaintenanceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MaintenanceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MaintenanceDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MaintenanceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
