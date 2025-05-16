// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id          int         `json:"id"          orm:"id"          description:"角色ID"`           // 角色ID
	Name        string      `json:"name"        orm:"name"        description:"角色名称"`           // 角色名称
	Code        string      `json:"code"        orm:"code"        description:"角色编码"`           // 角色编码
	Description string      `json:"description" orm:"description" description:"角色描述"`           // 角色描述
	Status      int         `json:"status"      orm:"status"      description:"状态: 0-禁用, 1-启用"` // 状态: 0-禁用, 1-启用
	CreatedAt   *gtime.Time `json:"created_at"  orm:"created_at"  description:"创建时间"`           // 创建时间
	UpdatedAt   *gtime.Time `json:"updated_at"  orm:"updated_at"  description:"更新时间"`           // 更新时间
	DeletedAt   *gtime.Time `json:"deleted_at"  orm:"deleted_at"  description:"删除时间"`           // 删除时间
}
