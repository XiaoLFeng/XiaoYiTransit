// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id            int         `json:"id"              orm:"id"              description:"用户ID"`           // 用户ID
	Username      string      `json:"username"        orm:"username"        description:"用户名"`            // 用户名
	Password      string      `json:"password"        orm:"password"        description:"密码"`             // 密码
	RealName      string      `json:"real_name"       orm:"real_name"       description:"真实姓名"`           // 真实姓名
	Email         string      `json:"email"           orm:"email"           description:"邮箱"`             // 邮箱
	Phone         string      `json:"phone"           orm:"phone"           description:"手机号"`            // 手机号
	Avatar        string      `json:"avatar"          orm:"avatar"          description:"头像"`             // 头像
	RoleId        int         `json:"role_id"         orm:"role_id"         description:"角色ID"`           // 角色ID
	Status        int         `json:"status"          orm:"status"          description:"状态: 0-禁用, 1-启用"` // 状态: 0-禁用, 1-启用
	LastLoginTime *gtime.Time `json:"last_login_time" orm:"last_login_time" description:"最后登录时间"`         // 最后登录时间
	CreatedAt     *gtime.Time `json:"created_at"      orm:"created_at"      description:"创建时间"`           // 创建时间
	UpdatedAt     *gtime.Time `json:"updated_at"      orm:"updated_at"      description:"更新时间"`           // 更新时间
	DeletedAt     *gtime.Time `json:"deleted_at"      orm:"deleted_at"      description:"删除时间"`           // 删除时间
}
