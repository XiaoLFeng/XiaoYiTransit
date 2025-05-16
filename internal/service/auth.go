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
