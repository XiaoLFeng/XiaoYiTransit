// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserTokenDao is the data access object for the table xf_user_token.
type UserTokenDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserTokenColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserTokenColumns defines and stores column names for the table xf_user_token.
type UserTokenColumns struct {
	TokenUuid string // Token唯一标识UUID
	UserUuid  string // 用户UUID
	UserAgent string // 用户浏览器和设备信息
	ClientIp  string // 用户客户端IP
	Referer   string // 请求来源页面
	CreatedAt string // 创建时间
	ExpiresAt string // 过期时间
}

// userTokenColumns holds the columns for the table xf_user_token.
var userTokenColumns = UserTokenColumns{
	TokenUuid: "token_uuid",
	UserUuid:  "user_uuid",
	UserAgent: "user_agent",
	ClientIp:  "client_ip",
	Referer:   "referer",
	CreatedAt: "created_at",
	ExpiresAt: "expires_at",
}

// NewUserTokenDao creates and returns a new DAO object for table data access.
func NewUserTokenDao(handlers ...gdb.ModelHandler) *UserTokenDao {
	return &UserTokenDao{
		group:    "default",
		table:    "xf_user_token",
		columns:  userTokenColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserTokenDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserTokenDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserTokenDao) Columns() UserTokenColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserTokenDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserTokenDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserTokenDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
