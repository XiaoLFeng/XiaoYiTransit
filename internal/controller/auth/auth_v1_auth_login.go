package auth

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/util/gconv"
	"xiao-yi-transit/internal/model/dto/back"
	"xiao-yi-transit/internal/model/dto/base"
	"xiao-yi-transit/internal/service"

	"xiao-yi-transit/api/auth/v1"
)

// AuthLogin 处理用户登录逻辑。
//
// 实现用户认证过程，包括账号密码验证、Token 生成和登录结果返回。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 用户登录请求参数，包含用户名和密码。
//
// 返回:
//   - res: 用户登录的响应结果，包含用户信息和授权 Token。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) AuthLogin(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error) {
	blog.ControllerInfo(ctx, "AuthLogin", "用户登录请求: %s", req.Username)

	// 获取用户
	iUser := service.User()
	userEntity, errorCode := iUser.GetUserByUsername(ctx, req.Username)
	if errorCode != nil {
		return nil, errorCode
	}

	// 验证密码
	iAuth := service.Auth()
	if ok, errorCode := iAuth.VerifyPassword(ctx, userEntity, req.Password); !ok {
		return nil, errorCode
	}

	// 获取角色
	iRole := service.Role()
	getRole, errorCode := iRole.GetRoleByUUID(ctx, userEntity.RoleUuid)
	if errorCode != nil {
		blog.ServiceError(ctx, "AuthLogin", "获取角色信息失败: %s", errorCode.Error())
		return nil, errorCode
	}

	// 生成token
	iToken := service.Token()
	tokenEntity, errorCode := iToken.GenerateUserToken(ctx, userEntity)
	if errorCode != nil {
		return nil, errorCode
	}

	// 数据转换
	var backUser *base.UserInfoDTO
	operateErr := gconv.Struct(userEntity, &backUser)
	if operateErr != nil {
		blog.ServiceError(ctx, "AuthLogin", "数据转换失败: %s", operateErr.Error())
		return nil, &berror.ErrInternalServer
	}
	var backToken *base.TokenDTO
	operateErr = gconv.Struct(tokenEntity, &backToken)
	if operateErr != nil {
		blog.ServiceError(ctx, "AuthLogin", "数据转换失败: %s", operateErr.Error())
		return nil, &berror.ErrInternalServer
	}
	var backRole *base.RoleDTO
	operateErr = gconv.Struct(getRole, &backRole)
	if operateErr != nil {
		blog.ServiceError(ctx, "AuthLogin", "数据转换失败: %s", operateErr.Error())
		return nil, &berror.ErrInternalServer
	}

	backAuthLogin := &back.AuthLoginBackDTO{
		User:  backUser,
		Token: backToken,
		Role:  backRole,
	}

	// 返回结果
	return &v1.AuthLoginRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "登录成功", backAuthLogin),
	}, nil
}
