// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RouteStationDao is the data access object for the table xf_route_station.
type RouteStationDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  RouteStationColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// RouteStationColumns defines and stores column names for the table xf_route_station.
type RouteStationColumns struct {
	RouteStationUuid  string // 线路站点UUID
	RouteUuid         string // 线路UUID
	StationUuid       string // 站点UUID
	Sequence          string // 站点顺序
	DistanceFromStart string // 距起点距离(km)
	EstimatedTime     string // 预计到达时间(分钟)
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
}

// routeStationColumns holds the columns for the table xf_route_station.
var routeStationColumns = RouteStationColumns{
	RouteStationUuid:  "route_station_uuid",
	RouteUuid:         "route_uuid",
	StationUuid:       "station_uuid",
	Sequence:          "sequence",
	DistanceFromStart: "distance_from_start",
	EstimatedTime:     "estimated_time",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewRouteStationDao creates and returns a new DAO object for table data access.
func NewRouteStationDao(handlers ...gdb.ModelHandler) *RouteStationDao {
	return &RouteStationDao{
		group:    "default",
		table:    "xf_route_station",
		columns:  routeStationColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RouteStationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RouteStationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RouteStationDao) Columns() RouteStationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RouteStationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RouteStationDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RouteStationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
