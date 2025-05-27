package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"xiao-yi-transit/internal/model/dto/back"
)

type AuthCurrentReq struct {
	g.Meta `path:"/auth/current" method:"Get" tags:"授权控制器" sm:"获取当前用户信息" dc:"用于获取当前登录用户的信息"`
}

type AuthCurrentRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*back.AuthCurrentBackDTO]
}
