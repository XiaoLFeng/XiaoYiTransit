package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"go/types"
)

// DeleteDriverReq defines the request for deleting a driver
type DeleteDriverReq struct {
	g.Meta     `path:"/driver/delete" method:"Delete" tags:"司机管理" sm:"删除司机" dc:"用于删除司机信息"`
	DriverUuid string `json:"driver_uuid" v:"required#司机UUID不能为空" dc:"司机UUID"`
}

// DeleteDriverRes defines the response for deleting a driver
type DeleteDriverRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil] // No data returned, just success message
}
