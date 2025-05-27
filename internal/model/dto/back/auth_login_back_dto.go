package back

import "xiao-yi-transit/internal/model/dto/base"

// AuthLoginBackDTO 表示用户登录返回的数据结构。
// 包含用户基本信息和登录凭证。
type AuthLoginBackDTO struct {
	User  *base.UserInfoDTO `json:"user" dc:"用户基本信息"`
	Role  *base.RoleDTO     `json:"role" dc:"用户角色信息"`
	Token *base.TokenDTO    `json:"token" dc:"用户登录凭证"`
}
