// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Station is the golang structure of table xf_station for DAO operations like Where/Data.
type Station struct {
	g.Meta    `orm:"table:xf_station, do:true"`
	Id        interface{} // 站点ID
	Name      interface{} // 站点名称
	Code      interface{} // 站点编码
	Address   interface{} // 站点地址
	Longitude interface{} // 经度
	Latitude  interface{} // 纬度
	Status    interface{} // 状态: 0-停用, 1-启用
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
