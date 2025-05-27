package user

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"xiao-yi-transit/internal/custom/cerror"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

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
func (s *sUser) GetUserById(ctx context.Context, userUUID string) (*entity.User, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetUserById", "获取用户 %s 信息", userUUID)

	userModel := dao.User.Ctx(ctx)

	// 查询用户信息
	var getUser *entity.User
	sqlErr := userModel.Where(&entity.User{UserUuid: userUUID}).Scan(&getUser)
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetUserById", "获取用户 %s 信息失败: %s", userUUID, sqlErr.Error())
		return nil, &berror.ErrDatabaseError
	}
	if getUser == nil {
		blog.ServiceError(ctx, "GetUserById", "获取用户 %s 信息失败: 用户不存在", userUUID)
		return nil, cerror.ErrUserNotExist
	}

	// 返回结果
	return getUser, nil
}

// GetUserByUsername 根据用户名获取用户信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求的生命周期。
//   - username: 用户名，用于查询对应的用户信息。
//
// 返回:
//   - 用户信息的指针，包含详细的用户数据。
//   - 错误码的指针，表示可能的错误类型，例如用户不存在或数据库错误。
func (s *sUser) GetUserByUsername(ctx context.Context, username string) (*entity.User, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetUserByUsername", "获取用户 %s 信息", username)

	userModel := dao.User.Ctx(ctx)

	// 查询用户信息
	var getUser *entity.User
	sqlErr := userModel.Where(&entity.User{Username: username}).Scan(&getUser)
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetUserByUsername", "获取用户 %s 信息失败: %s", username, sqlErr.Error())
		return nil, &berror.ErrDatabaseError
	}
	if getUser == nil {
		return nil, cerror.ErrUserNotExist
	}

	// 返回结果
	return getUser, nil
}
