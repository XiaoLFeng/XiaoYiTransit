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
	IRole interface {
		// GetRoleByUUID 根据角色UUID获取角色信息。
		//
		// 参数:
		//   - ctx: 上下文信息
		//   - roleUUID: 要查询的角色UUID
		//
		// 返回:
		//   - 角色信息（*entity.Role）
		//   - 错误信息（*berror.ErrorCode）
		//
		// 错误:
		//   - 若数据库查询失败，返回ErrDatabaseError
		//   - 若角色不存在，返回ErrResourceNotFound
		GetRoleByUUID(ctx context.Context, roleUUID string) (*entity.Role, *berror.ErrorCode)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
