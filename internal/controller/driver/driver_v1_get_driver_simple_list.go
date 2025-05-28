package driver

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/util/gconv"
	"xiao-yi-transit/api/driver/v1"
	"xiao-yi-transit/internal/model/dto/driver"
	"xiao-yi-transit/internal/service"
)

// GetDriverSimpleList 处理获取司机简易列表的逻辑。
//
// 实现司机简易列表的查询过程，用于前端下拉选项框。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取司机简易列表的请求参数。
//
// 返回:
//   - res: 获取司机简易列表的响应结果，包含司机简易列表。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) GetDriverSimpleList(ctx context.Context, req *v1.GetDriverSimpleListReq) (res *v1.GetDriverSimpleListRes, err error) {
	blog.ControllerInfo(ctx, "GetDriverSimpleList", "获取司机简易列表请求")

	// 获取司机简易列表
	iDriver := service.Driver()
	drivers, errorCode := iDriver.GetDriverSimpleList(ctx)
	if errorCode != nil {
		return nil, errorCode
	}

	// 转换数据
	var driverList []*driver.SimpleDriverItemDTO
	for _, driverEntity := range drivers {
		var item driver.SimpleDriverItemDTO
		if err := gconv.Struct(driverEntity, &item); err != nil {
			blog.ControllerError(ctx, "GetDriverSimpleList", "数据转换失败: %s", err.Error())
			return nil, &berror.ErrInternalServer
		}
		driverList = append(driverList, &item)
	}

	// 构建返回数据
	simpleList := &driver.SimpleDriverListDTO{
		List: driverList,
	}

	// 返回结果
	return &v1.GetDriverSimpleListRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取司机简易列表成功", simpleList),
	}, nil
}
