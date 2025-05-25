package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"go/types"
)

type AuthForgetPasswordReq struct {
	g.Meta             `path:"/auth/password/change" method:"Put" sm:"修改密码" tags:"授权控制器" dc:"用于用户在已登录的环境下修改自己的密码"`
	OriginPassword     string `json:"origin_password" description:"原密码" v:"required#原密码不能为空"`
	NewPassword        string `json:"new_password" description:"新密码" v:"required#新密码不能为空"`
	NewPasswordConfirm string `json:"new_password_confirm" description:"确认新密码" v:"required#确认新密码不能为空|same:new_password#两次输入的新密码不一致"`
}

type AuthForgetPasswordRes struct {
	g.Meta `mime:"application/json; charset=utf-8" description:"修改密码响应"`
	*bmodels.ResponseDTO[*types.Nil]
}
