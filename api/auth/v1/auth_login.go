package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"xiao-yi-transit/internal/model/dto/back"
)

type AuthLoginReq struct {
	g.Meta   `path:"/auth/login" method:"Post" tags:"授权控制器" sm:"用户登录" dc:"用于用户登录"`
	Username string `json:"username" v:"required|size:6,32|regex:^[0-9A-Za-z-\\_]+$#用户名不能为空|用户名长度为6-32个字符|用户名只能包含字母、数字、下划线和短横线" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
}

type AuthLoginRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*back.AuthLoginBackDTO]
}
