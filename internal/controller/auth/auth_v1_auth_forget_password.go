package auth

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"xiao-yi-transit/internal/service"
	"xiao-yi-transit/utility"

	"xiao-yi-transit/api/auth/v1"
)

// AuthForgetPassword 用户通过已验证的环境修改密码的逻辑处理。
//
// 接收用户的原密码和新密码，验证原密码后对新密码进行更改。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 忘记密码请求参数，包括原密码、新密码和确认新密码。
//
// 返回:
//   - res: 密码修改的响应结果，包含操作成功消息。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) AuthForgetPassword(ctx context.Context, req *v1.AuthForgetPasswordReq) (res *v1.AuthForgetPasswordRes, err error) {
	blog.ControllerInfo(ctx, "AuthForgetPassword", "用户忘记密码请求")

	// 获取用户信息
	getUser, errorCode := utility.GetUserByCtx(ctx)
	if errorCode != nil {
		blog.ControllerError(ctx, "AuthForgetPassword", "获取用户信息失败: %s", errorCode.Message)
		return nil, errorCode
	}

	// 检查原密码是否正确
	iAuth := service.Auth()
	if _, errorCode := iAuth.VerifyPassword(ctx, getUser, req.OriginPassword); errorCode == nil {
		// 修改密码
		errorCode := iAuth.ChangePassword(ctx, getUser, req.NewPassword, req.NewPasswordConfirm)
		if errorCode != nil {
			blog.ControllerError(ctx, "AuthForgetPassword", "修改密码失败: %s", errorCode.Message)
			return nil, errorCode
		}
	} else {
		blog.ControllerError(ctx, "AuthForgetPassword", "验证原密码失败: %s", errorCode.Message)
		return nil, errorCode
	}

	// 返回结果
	return &v1.AuthForgetPasswordRes{
		ResponseDTO: bresult.Success(ctx, "密码修改成功"),
	}, nil
}
