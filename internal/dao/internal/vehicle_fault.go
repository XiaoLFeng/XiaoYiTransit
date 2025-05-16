// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VehicleFaultDao is the data access object for the table xf_vehicle_fault.
type VehicleFaultDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  VehicleFaultColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// VehicleFaultColumns defines and stores column names for the table xf_vehicle_fault.
type VehicleFaultColumns struct {
	Id               string // 故障ID
	VehicleId        string // 车辆ID
	ReporterId       string // 报告人ID
	FaultType        string // 故障类型
	FaultDescription string // 故障描述
	ReportDate       string // 报告时间
	FaultLocation    string // 故障发生地点
	Severity         string // 严重程度: 1-轻微, 2-一般, 3-严重, 4-危险
	Status           string // 状态: 0-已取消, 1-待处理, 2-处理中, 3-已解决
	MaintenanceId    string // 关联的维修记录ID
	Notes            string // 备注
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
}

// vehicleFaultColumns holds the columns for the table xf_vehicle_fault.
var vehicleFaultColumns = VehicleFaultColumns{
	Id:               "id",
	VehicleId:        "vehicle_id",
	ReporterId:       "reporter_id",
	FaultType:        "fault_type",
	FaultDescription: "fault_description",
	ReportDate:       "report_date",
	FaultLocation:    "fault_location",
	Severity:         "severity",
	Status:           "status",
	MaintenanceId:    "maintenance_id",
	Notes:            "notes",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewVehicleFaultDao creates and returns a new DAO object for table data access.
func NewVehicleFaultDao(handlers ...gdb.ModelHandler) *VehicleFaultDao {
	return &VehicleFaultDao{
		group:    "default",
		table:    "xf_vehicle_fault",
		columns:  vehicleFaultColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *VehicleFaultDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *VehicleFaultDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *VehicleFaultDao) Columns() VehicleFaultColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *VehicleFaultDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *VehicleFaultDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *VehicleFaultDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
