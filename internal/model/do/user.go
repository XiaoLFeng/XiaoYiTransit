// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table xf_user for DAO operations like Where/Data.
type User struct {
	g.Meta        `orm:"table:xf_user, do:true"`
	Id            interface{} // 用户ID
	Username      interface{} // 用户名
	Password      interface{} // 密码
	RealName      interface{} // 真实姓名
	Email         interface{} // 邮箱
	Phone         interface{} // 手机号
	Avatar        interface{} // 头像
	RoleId        interface{} // 角色ID
	Status        interface{} // 状态: 0-禁用, 1-启用
	LastLoginTime *gtime.Time // 最后登录时间
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
}
