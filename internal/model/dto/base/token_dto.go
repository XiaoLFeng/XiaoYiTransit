package base

import "github.com/gogf/gf/v2/os/gtime"

type TokenDTO struct {
	TokenUuid string      `json:"token_uuid" orm:"token_uuid" description:"Token唯一标识UUID"`
	UserUuid  string      `json:"user_uuid"  orm:"user_uuid"  description:"用户UUID"`
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`
	ExpiresAt *gtime.Time `json:"expires_at" orm:"expires_at" description:"过期时间"`
}
