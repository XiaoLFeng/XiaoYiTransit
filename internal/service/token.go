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
	IToken interface {
		// GenerateUserToken 生成用户的认证 Token。
		//
		// 参数:
		//   - ctx: 请求上下文。
		//   - user: 用户信息，不能为空。
		//
		// 返回:
		//   - *entity.UserToken: 生成的用户 Token 信息。
		//   - *berror.ErrorCode: 错误码，如果发生错误。
		//
		// 错误:
		//   - 返回 ErrUserNotExist 错误码，如果用户信息为空。
		//   - 返回 ErrDatabaseError 错误码，如果数据库操作失败。
		GenerateUserToken(ctx context.Context, user *entity.User) (*entity.UserToken, *berror.ErrorCode)
	}
)

var (
	localToken IToken
)

func Token() IToken {
	if localToken == nil {
		panic("implement not found for interface IToken, forgot register?")
	}
	return localToken
}

func RegisterToken(i IToken) {
	localToken = i
}
