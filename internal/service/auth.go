// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xiao-yi-transit/internal/model/entity"

	"github.com/XiaoLFeng/bamboo-utils/berror"
)

type (
	IAuth interface {
		// VerifyPassword 验证用户密码是否正确。
		//
		// 参数:
		//   - ctx: 上下文对象，用于日志和跟踪。
		//   - entity: 用户实体，包含用户名和加密后的密码信息。
		//   - password: 用户输入的明文密码。
		//
		// 返回:
		//   - bool: 密码是否验证通过。
		//   - *berror.ErrorCode: 错误代码，如果有错误则返回相关错误。
		VerifyPassword(ctx context.Context, entity *entity.User, password string) (bool, *berror.ErrorCode)
		// ChangePassword 更改用户密码。
		//
		// 参数:
		//   - ctx: 上下文对象，用于日志和追踪。
		//   - user: 用户实体，包含用户信息。
		//   - newPassword: 用户设置的新密码。
		//   - confirmPassword: 用户确认的新密码。
		//
		// 返回:
		//   - *berror.ErrorCode: 错误代码，如果操作失败则返回相关错误。
		//
		// 错误:
		//   - 如果用户实体为空，返回 ErrUserNotExist。
		//   - 如果新密码与确认密码不一致，返回 ErrPasswordMismatch。
		//   - 如果发生其他内部错误，返回相应的错误代码。
		ChangePassword(ctx context.Context, user *entity.User, newPassword string, confirmPassword string) *berror.ErrorCode
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
