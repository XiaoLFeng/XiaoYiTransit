// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ScheduleTemplateDao is the data access object for the table xf_schedule_template.
type ScheduleTemplateDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  ScheduleTemplateColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// ScheduleTemplateColumns defines and stores column names for the table xf_schedule_template.
type ScheduleTemplateColumns struct {
	ScheduleTemplateUuid string // 模板UUID
	Name                 string // 模板名称
	RouteUuid            string // 线路UUID
	DayOfWeek            string // 星期几: 1-7 代表周一至周日
	StartTime            string // 开始时间
	EndTime              string // 结束时间
	VehicleCount         string // 所需车辆数
	Status               string // 状态: 0-禁用, 1-启用
	Notes                string // 备注
	CreatedAt            string // 创建时间
	UpdatedAt            string // 更新时间
}

// scheduleTemplateColumns holds the columns for the table xf_schedule_template.
var scheduleTemplateColumns = ScheduleTemplateColumns{
	ScheduleTemplateUuid: "schedule_template_uuid",
	Name:                 "name",
	RouteUuid:            "route_uuid",
	DayOfWeek:            "day_of_week",
	StartTime:            "start_time",
	EndTime:              "end_time",
	VehicleCount:         "vehicle_count",
	Status:               "status",
	Notes:                "notes",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

// NewScheduleTemplateDao creates and returns a new DAO object for table data access.
func NewScheduleTemplateDao(handlers ...gdb.ModelHandler) *ScheduleTemplateDao {
	return &ScheduleTemplateDao{
		group:    "default",
		table:    "xf_schedule_template",
		columns:  scheduleTemplateColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ScheduleTemplateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ScheduleTemplateDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ScheduleTemplateDao) Columns() ScheduleTemplateColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ScheduleTemplateDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ScheduleTemplateDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ScheduleTemplateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
