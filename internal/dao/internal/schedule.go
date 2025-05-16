// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ScheduleDao is the data access object for the table xf_schedule.
type ScheduleDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ScheduleColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ScheduleColumns defines and stores column names for the table xf_schedule.
type ScheduleColumns struct {
	ScheduleUuid    string // 调度UUID
	RouteUuid       string // 线路UUID
	VehicleUuid     string // 车辆UUID
	DriverUuid      string // 司机UUID
	ScheduleDate    string // 调度日期
	StartTime       string // 开始时间
	EndTime         string // 结束时间
	Status          string // 状态: 0-取消, 1-待执行, 2-执行中, 3-已完成
	ActualStartTime string // 实际开始时间
	ActualEndTime   string // 实际结束时间
	Mileage         string // 行驶里程(km)
	FuelConsumption string // 油耗(L)
	PassengerCount  string // 载客人数
	Notes           string // 备注
	CreatedByUuid   string // 创建人UUID
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
	DeletedAt       string // 删除时间
}

// scheduleColumns holds the columns for the table xf_schedule.
var scheduleColumns = ScheduleColumns{
	ScheduleUuid:    "schedule_uuid",
	RouteUuid:       "route_uuid",
	VehicleUuid:     "vehicle_uuid",
	DriverUuid:      "driver_uuid",
	ScheduleDate:    "schedule_date",
	StartTime:       "start_time",
	EndTime:         "end_time",
	Status:          "status",
	ActualStartTime: "actual_start_time",
	ActualEndTime:   "actual_end_time",
	Mileage:         "mileage",
	FuelConsumption: "fuel_consumption",
	PassengerCount:  "passenger_count",
	Notes:           "notes",
	CreatedByUuid:   "created_by_uuid",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}

// NewScheduleDao creates and returns a new DAO object for table data access.
func NewScheduleDao(handlers ...gdb.ModelHandler) *ScheduleDao {
	return &ScheduleDao{
		group:    "default",
		table:    "xf_schedule",
		columns:  scheduleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ScheduleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ScheduleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ScheduleDao) Columns() ScheduleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ScheduleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ScheduleDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ScheduleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
