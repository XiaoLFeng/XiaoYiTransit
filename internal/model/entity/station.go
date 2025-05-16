// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Station is the golang structure for table station.
type Station struct {
	Id        int         `json:"id"         orm:"id"         description:"站点ID"`           // 站点ID
	Name      string      `json:"name"       orm:"name"       description:"站点名称"`           // 站点名称
	Code      string      `json:"code"       orm:"code"       description:"站点编码"`           // 站点编码
	Address   string      `json:"address"    orm:"address"    description:"站点地址"`           // 站点地址
	Longitude float64     `json:"longitude"  orm:"longitude"  description:"经度"`             // 经度
	Latitude  float64     `json:"latitude"   orm:"latitude"   description:"纬度"`             // 纬度
	Status    int         `json:"status"     orm:"status"     description:"状态: 0-停用, 1-启用"` // 状态: 0-停用, 1-启用
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`           // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"更新时间"`           // 更新时间
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:"删除时间"`           // 删除时间
}
