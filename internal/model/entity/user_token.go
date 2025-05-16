// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserToken is the golang structure for table user_token.
type UserToken struct {
	TokenUuid string      `json:"token_uuid" orm:"token_uuid" description:"Token唯一标识UUID"` // Token唯一标识UUID
	UserUuid  string      `json:"user_uuid"  orm:"user_uuid"  description:"用户UUID"`        // 用户UUID
	UserAgent string      `json:"user_agent" orm:"user_agent" description:"用户浏览器和设备信息"`    // 用户浏览器和设备信息
	ClientIp  string      `json:"client_ip"  orm:"client_ip"  description:"用户客户端IP"`       // 用户客户端IP
	Referer   string      `json:"referer"    orm:"referer"    description:"请求来源页面"`        // 请求来源页面
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`          // 创建时间
	ExpiresAt *gtime.Time `json:"expires_at" orm:"expires_at" description:"过期时间"`          // 过期时间
}
