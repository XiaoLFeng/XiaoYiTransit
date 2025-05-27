package token

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
	"time"
	"xiao-yi-transit/internal/custom/cerror"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

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
func (s *sToken) GenerateUserToken(ctx context.Context, user *entity.User) (*entity.UserToken, *berror.ErrorCode) {
	if user == nil {
		blog.ServiceError(ctx, "GenerateUserToken", "生成用户 Token 失败: 用户不存在")
		return nil, cerror.ErrUserNotExist
	} else {
		blog.ServiceInfo(ctx, "GenerateUserToken", "生成用户 %s 的 Token", user.Username)
	}

	// 获取信息
	request := ghttp.RequestFromCtx(ctx)

	// 生成 Token
	token := &entity.UserToken{
		TokenUuid: uuid.New().String(),
		UserUuid:  user.UserUuid,
		UserAgent: request.UserAgent(),
		ClientIp:  request.GetClientIp(),
		Referer:   request.Referer(),
		ExpiresAt: gtime.Now().Add(6 * time.Hour),
	}
	// 插入数据库
	_, sqlErr := dao.UserToken.Ctx(ctx).OmitEmpty().Insert(token)
	if sqlErr != nil {
		blog.ServiceError(ctx, "GenerateUserToken", "生成用户 %s 的 Token 失败: %s", user.Username, sqlErr.Error())
		return nil, &berror.ErrDatabaseError
	}
	// 返回结果
	return token, nil
}
