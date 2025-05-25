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
	IUser interface {
		// GetUserById 根据用户ID获取用户信息。
		//
		// 若不产生报错，则 *entity.User 返回用户信息，用户必然存在
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - userUUID: 用户的唯一标识符。
		//
		// 返回:
		//   - 用户信息的指针，包含详细用户数据。
		//   - 错误码的指针，表示错误类型，如用户不存在或内部错误。
		GetUserById(ctx context.Context, userUUID string) (*entity.User, *berror.ErrorCode)
		// GetUserByUsername 根据用户名获取用户信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求的生命周期。
		//   - username: 用户名，用于查询对应的用户信息。
		//
		// 返回:
		//   - 用户信息的指针，包含详细的用户数据。
		//   - 错误码的指针，表示可能的错误类型，例如用户不存在或数据库错误。
		GetUserByUsername(ctx context.Context, username string) (*entity.User, *berror.ErrorCode)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
