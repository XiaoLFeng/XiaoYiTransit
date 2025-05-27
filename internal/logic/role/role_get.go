package role

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/do"
	"xiao-yi-transit/internal/model/entity"
)

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
func (s *sRole) GetRoleByUUID(ctx context.Context, roleUUID string) (*entity.Role, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetRoleByUUID", "获取角色信息: %s", roleUUID)

	// 查询角色信息
	var roleEntity *entity.Role
	sqlErr := dao.Role.Ctx(ctx).Where(&do.Role{RoleUuid: roleUUID}).Scan(&roleEntity)
	if sqlErr != nil {
		blog.ServiceError(ctx, "GetRoleByUUID", "查询角色信息失败: %s", sqlErr.Error())
		return nil, &berror.ErrDatabaseError
	}
	if roleEntity == nil {
		return nil, &berror.ErrResourceNotFound
	}
	// 返回结果
	return roleEntity, nil
}
