package utility

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/gogf/gf/v2/net/ghttp"
	"xiao-yi-transit/internal/custom/cerror"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/do"
	"xiao-yi-transit/internal/model/entity"
)

// GetUserByCtx 根据上下文获取用户信息。
//
// 参数:
//   - ctx: 请求上下文对象。
//
// 返回:
//   - 用户实体指针。
//   - 错误代码，如果发生错误。
//
// 错误:
//   - berror.ErrUnauthorized: 授权信息无效或丢失。
//   - cerror.ErrUserTokenNotExist: 用户令牌不存在。
//   - cerror.ErrUserNotExist: 用户信息不存在。
//   - berror.ErrDatabaseError: 数据库操作错误。
func GetUserByCtx(ctx context.Context) (*entity.User, *berror.ErrorCode) {
	req := ghttp.RequestFromCtx(ctx)
	getHeader := req.GetHeader("Authorization")
	if getHeader[:5] == "Bearer" {
		tokenUUID := getHeader[7:]
		if tokenUUID == "" {
			return nil, &berror.ErrUnauthorized
		}
		var userToken *entity.UserToken
		sqlErr := dao.UserToken.Ctx(ctx).Where(&do.UserToken{TokenUuid: tokenUUID}).Scan(&userToken)
		if sqlErr != nil {
			return nil, &berror.ErrDatabaseError
		}
		if userToken == nil {
			return nil, cerror.ErrUserTokenNotExist
		} else {
			// 获取用户信息
			var userEntity *entity.User
			sqlErr = dao.User.Ctx(ctx).Where(&do.User{UserUuid: userToken.UserUuid}).Scan(&userEntity)
			if sqlErr != nil {
				return nil, &berror.ErrDatabaseError
			}
			if userEntity == nil {
				return nil, cerror.ErrUserNotExist
			}
			return userEntity, nil
		}
	} else {
		return nil, &berror.ErrUnauthorized
	}
}
