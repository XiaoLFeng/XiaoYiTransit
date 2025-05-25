package setup

import (
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
	"xiao-yi-transit/internal/consts"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/entity"
)

// DefaultRoleSetup 创建系统的默认角色设置。
//
// 执行步骤包括:
//   - 检查默认角色（如管理员、普通用户）是否存在
//   - 若不存在，则创建默认角色
//   - 缓存角色UUID，用于后续系统使用
//
// 错误:
//   - 查询角色表失败
//   - 数据库保存失败
//   - 角色缓存失败
func (setup *SystemSetup) DefaultRoleSetup() {
	// 创建默认角色
	var roles = [][]string{
		{"admin", "A0001", "系统管理员"},
		{"user", "A0002", "普通用户"},
	}
	for _, role := range roles {
		roleModel := dao.Role.Ctx(setup.ctx)
		var roleEntity *entity.Role
		sqlErr := roleModel.Where(&entity.Role{Name: role[0]}).Scan(&roleEntity)
		if sqlErr != nil {
			blog.BambooError(setup.ctx, "DefaultRoleSetup", "查询角色表失败: %s", sqlErr.Error())
			panic(sqlErr)
		}
		// 检查是否存在
		if roleEntity == nil {
			roleEntity = &entity.Role{
				RoleUuid:    uuid.New().String(),
				Name:        role[0],
				Code:        role[1],
				Description: role[2],
				CreatedAt:   gtime.Now(),
				UpdatedAt:   gtime.Now(),
			}
			_, sqlErr := dao.Role.Ctx(setup.ctx).Save(&roleEntity)
			if sqlErr != nil {
				blog.BambooError(setup.ctx, "DefaultRoleSetup", "创建角色 %s 失败: %s", role[0], sqlErr.Error())
				panic(sqlErr)
			}
		}
		// 写入缓存
		switch role[0] {
		case "admin":
			consts.RoleAdminUUID = roleEntity.RoleUuid
		case "user":
			consts.RoleUserUUID = roleEntity.RoleUuid
		default:
			blog.BambooError(setup.ctx, "DefaultRoleSetup", "未知角色 %s，跳过缓存设置", role[0])
		}
	}
}
