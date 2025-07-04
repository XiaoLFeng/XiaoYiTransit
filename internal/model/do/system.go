// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// System is the golang structure of table xf_system for DAO operations like Where/Data.
type System struct {
	g.Meta     `orm:"table:xf_system, do:true"`
	SystemUuid interface{} // 系统UUID
	Key        interface{} // 系统键
	Val        interface{} // 系统值
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
