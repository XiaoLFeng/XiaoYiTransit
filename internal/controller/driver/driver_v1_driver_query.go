package driver

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/util/gconv"
	"xiao-yi-transit/api/driver/v1"
	"xiao-yi-transit/internal/service"
)

// GetDriverList 处理获取司机列表的逻辑。
//
// 实现司机列表的查询过程，包括分页、筛选等。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取司机列表的请求参数，包含分页和筛选条件。
//
// 返回:
//   - res: 获取司机列表的响应结果，包含司机列表和分页信息。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) GetDriverList(ctx context.Context, req *v1.GetDriverListReq) (res *v1.GetDriverListRes, err error) {
	blog.ControllerInfo(ctx, "GetDriverList", "获取司机列表请求: page=%d, size=%d", req.Page, req.Size)

	// 获取司机列表
	iDriver := service.Driver()
	drivers, total, errorCode := iDriver.GetDriverList(ctx, req.Page, req.Size, req.EmployeeId, req.Name, req.Status)
	if errorCode != nil {
		return nil, errorCode
	}

	// 转换数据
	var driverList []*v1.DriverListItem
	for _, driver := range drivers {
		var item v1.DriverListItem
		if err := gconv.Struct(driver, &item); err != nil {
			blog.ControllerError(ctx, "GetDriverList", "数据转换失败: %s", err.Error())
			return nil, &berror.ErrInternalServer
		}
		driverList = append(driverList, &item)
	}

	// 构建返回数据
	pagedList := &v1.PagedDriverList{
		List:  driverList,
		Page:  req.Page,
		Size:  req.Size,
		Total: total,
	}

	// 返回结果
	return &v1.GetDriverListRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取司机列表成功", pagedList),
	}, nil
}

// GetDriverDetail 处理获取司机详情的逻辑。
//
// 实现司机详情的查询过程。
//
// 参数:
//   - ctx: 请求的上下文信息。
//   - req: 获取司机详情的请求参数，包含司机UUID。
//
// 返回:
//   - res: 获取司机详情的响应结果，包含司机详细信息。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) GetDriverDetail(ctx context.Context, req *v1.GetDriverDetailReq) (res *v1.GetDriverDetailRes, err error) {
	blog.ControllerInfo(ctx, "GetDriverDetail", "获取司机详情请求: %s", req.DriverUuid)

	// 获取司机详情
	iDriver := service.Driver()
	driver, errorCode := iDriver.GetDriverById(ctx, req.DriverUuid)
	if errorCode != nil {
		return nil, errorCode
	}

	// 转换数据
	var driverDetail v1.DriverDetailItem
	if err := gconv.Struct(driver, &driverDetail); err != nil {
		blog.ControllerError(ctx, "GetDriverDetail", "数据转换失败: %s", err.Error())
		return nil, &berror.ErrInternalServer
	}

	// 返回结果
	return &v1.GetDriverDetailRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取司机详情成功", &driverDetail),
	}, nil
}