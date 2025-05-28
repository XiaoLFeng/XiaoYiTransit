package station

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/google/uuid"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// CreateStation 创建新的站点信息。
//
// 参数:
//   - ctx: 上下文信息，用于控制请求生命周期。
//   - station: 站点信息实体，包含需要创建的站点详细信息。
//
// 返回:
//   - 创建成功的站点UUID。
//   - 错误码的指针，表示可能的错误类型。
func (s *sStation) CreateStation(ctx context.Context, station *entity.Station) (string, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "CreateStation", "创建站点: %s", station.Name)

	// 生成UUID
	station.StationUuid = uuid.New().String()

	// 创建站点信息
	stationModel := dao.Station.Ctx(ctx)
	_, sqlErr := stationModel.Insert(station)
	if sqlErr != nil {
		blog.ServiceError(ctx, "CreateStation", "创建站点失败: %s", sqlErr.Error())
		return "", &berror.ErrDatabaseError
	}

	// 返回结果
	return station.StationUuid, nil
}
