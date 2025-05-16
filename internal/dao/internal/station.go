// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StationDao is the data access object for the table xf_station.
type StationDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  StationColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// StationColumns defines and stores column names for the table xf_station.
type StationColumns struct {
	StationUuid string // 站点UUID
	Name        string // 站点名称
	Code        string // 站点编码
	Address     string // 站点地址
	Longitude   string // 经度
	Latitude    string // 纬度
	Status      string // 状态: 0-停用, 1-启用
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
}

// stationColumns holds the columns for the table xf_station.
var stationColumns = StationColumns{
	StationUuid: "station_uuid",
	Name:        "name",
	Code:        "code",
	Address:     "address",
	Longitude:   "longitude",
	Latitude:    "latitude",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewStationDao creates and returns a new DAO object for table data access.
func NewStationDao(handlers ...gdb.ModelHandler) *StationDao {
	return &StationDao{
		group:    "default",
		table:    "xf_station",
		columns:  stationColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *StationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *StationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *StationDao) Columns() StationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *StationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *StationDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *StationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
