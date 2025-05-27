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
	"xiao-yi-transit/utility"

	"xiao-yi-transit/api/auth/v1"
)

// AuthCurrent 获取当前登录用户信息。
//
// 接收客户端请求，通过上下文信息获取当前登录用户的详细信息以及关联的角色和Token数据。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取当前用户信息的请求参数。
//
// 返回:
//   - res: 包含用户信息、角色和Token的响应结果。
//   - err: 执行过程中的错误信息。
func (c *ControllerV1) AuthCurrent(ctx context.Context, req *v1.AuthCurrentReq) (res *v1.AuthCurrentRes, err error) {
	blog.ControllerInfo(ctx, "AuthCurrent", "获取当前登录用户信息")

	// 获取当前登录用户
	userEntity, errorCode := utility.GetUserByCtx(ctx)
	if errorCode != nil {
		blog.ControllerError(ctx, "AuthCurrent", "获取当前用户信息失败: %s", errorCode.Message)
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
	return &v1.AuthCurrentRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "登录成功", backAuthLogin),
	}, nil
}
