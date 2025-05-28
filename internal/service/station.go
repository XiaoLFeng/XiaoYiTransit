// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xiao-yi-transit/internal/model/entity"

	"github.com/XiaoLFeng/bamboo-utils/berror"
)

type (
	IStation interface {
		// CreateStation 创建新的站点信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - station: 站点信息实体，包含需要创建的站点详细信息。
		//
		// 返回:
		//   - 创建成功的站点UUID。
		//   - 错误码的指针，表示可能的错误类型。
		CreateStation(ctx context.Context, station *entity.Station) (string, *berror.ErrorCode)
		// DeleteStation 删除站点信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - stationUuid: 站点的唯一标识符。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		DeleteStation(ctx context.Context, stationUuid string) *berror.ErrorCode
		// UpdateStation 更新站点信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - station: 站点信息实体，包含需要更新的站点详细信息。
		//
		// 返回:
		//   - 错误码的指针，表示可能的错误类型。
		UpdateStation(ctx context.Context, station *entity.Station) *berror.ErrorCode
		// GetStationById 根据站点UUID获取站点信息。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - stationUuid: 站点的唯一标识符。
		//
		// 返回:
		//   - 站点信息的指针，包含详细站点数据。
		//   - 错误码的指针，表示错误类型，如站点不存在或内部错误。
		GetStationById(ctx context.Context, stationUuid string) (*entity.Station, *berror.ErrorCode)
		// GetStationList 获取站点列表。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制请求生命周期。
		//   - name: 站点名称，用于模糊查询。
		//   - code: 站点编码，用于模糊查询。
		//   - status: 状态: 0-停用, 1-启用。
		//   - page: 页码，默认为1。
		//   - size: 每页数量，默认为10。
		//
		// 返回:
		//   - 站点列表。
		//   - 总数量。
		//   - 错误码的指针，表示可能的错误类型。
		GetStationList(ctx context.Context, name string, code string, status int, page int, size int) ([]*entity.Station, int, *berror.ErrorCode)
	}
)

var (
	localStation IStation
)

func Station() IStation {
	if localStation == nil {
		panic("implement not found for interface IStation, forgot register?")
	}
	return localStation
}

func RegisterStation(i IStation) {
	localStation = i
}
