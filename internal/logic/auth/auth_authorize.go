package auth

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/butil"
	"xiao-yi-transit/internal/custom/cerror"
	"xiao-yi-transit/internal/model/entity"
)

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
func (s *sAuth) VerifyPassword(ctx context.Context, entity *entity.User, password string) (bool, *berror.ErrorCode) {
	if entity == nil {
		blog.ServiceError(ctx, "VerifyPassword", "验证用户密码失败: 用户不存在")
		return false, cerror.ErrUserNotExist
	} else {
		blog.ServiceInfo(ctx, "VerifyPassword", "验证用户 %s 密码", entity.Username)
	}

	// 验证密码
	if butil.PasswordVerify(password, entity.Password) {
		return true, nil
	} else {
		blog.ServiceError(ctx, "VerifyPassword", "验证用户 %s 密码失败: 密码错误", entity.Username)
		return false, cerror.ErrPasswordIncorrect
	}

}
