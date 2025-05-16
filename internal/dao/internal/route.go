// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RouteDao is the data access object for the table xf_route.
type RouteDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  RouteColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// RouteColumns defines and stores column names for the table xf_route.
type RouteColumns struct {
	RouteUuid      string // 线路UUID
	RouteNumber    string // 线路编号
	Name           string // 线路名称
	StartStation   string // 起始站点
	EndStation     string // 终点站点
	Stops          string // 途经站点(JSON格式)
	Distance       string // 线路长度(km)
	Fare           string // 票价(元)
	OperationHours string // 运营时间
	Frequency      string // 发车频率
	Status         string // 状态: 0-停运, 1-运营
	Notes          string // 备注
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 删除时间
}

// routeColumns holds the columns for the table xf_route.
var routeColumns = RouteColumns{
	RouteUuid:      "route_uuid",
	RouteNumber:    "route_number",
	Name:           "name",
	StartStation:   "start_station",
	EndStation:     "end_station",
	Stops:          "stops",
	Distance:       "distance",
	Fare:           "fare",
	OperationHours: "operation_hours",
	Frequency:      "frequency",
	Status:         "status",
	Notes:          "notes",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewRouteDao creates and returns a new DAO object for table data access.
func NewRouteDao(handlers ...gdb.ModelHandler) *RouteDao {
	return &RouteDao{
		group:    "default",
		table:    "xf_route",
		columns:  routeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RouteDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RouteDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RouteDao) Columns() RouteColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RouteDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RouteDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RouteDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
