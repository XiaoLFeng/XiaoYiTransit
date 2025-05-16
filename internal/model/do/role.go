// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table xf_role for DAO operations like Where/Data.
type Role struct {
	g.Meta      `orm:"table:xf_role, do:true"`
	RoleUuid    interface{} // 角色UUID
	Name        interface{} // 角色名称
	Code        interface{} // 角色编码
	Description interface{} // 角色描述
	Status      interface{} // 状态: 0-禁用, 1-启用
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
