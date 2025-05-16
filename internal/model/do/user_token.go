// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserToken is the golang structure of table xf_user_token for DAO operations like Where/Data.
type UserToken struct {
	g.Meta    `orm:"table:xf_user_token, do:true"`
	TokenUuid interface{} // Token唯一标识UUID
	UserUuid  interface{} // 用户UUID
	UserAgent interface{} // 用户浏览器和设备信息
	ClientIp  interface{} // 用户客户端IP
	Referer   interface{} // 请求来源页面
	CreatedAt *gtime.Time // 创建时间
	ExpiresAt *gtime.Time // 过期时间
}
