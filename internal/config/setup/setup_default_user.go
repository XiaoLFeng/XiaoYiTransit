package setup

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/butil"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
	"xiao-yi-transit/internal/consts"
	"xiao-yi-transit/internal/dao"
	"xiao-yi-transit/internal/model/do"
	"xiao-yi-transit/internal/model/entity"
)

// DefaultUserSetup 创建系统的默认用户设置。
//
// 执行步骤包括:
//   - 检查默认用户是否已经存在
//   - 若不存在，则创建默认用户及其相关权限
//
// 错误:
//   - 数据库操作失败
//   - 数据格式不符合预期
func (setup *SystemSetup) DefaultUserSetup() {
	// 检查系统表是否存在超级管理员
	var systemEntity *entity.System
	sqlErr := dao.System.Ctx(setup.ctx).Where(&do.System{Key: consts.SystemKeyForSuperUserUUID}).Scan(&systemEntity)
	if sqlErr != nil {
		blog.BambooError(setup.ctx, "DefaultUserSetup", "查询系统表失败: %s", sqlErr.Error())
		panic(sqlErr)
	}
	if systemEntity != nil {
		blog.BambooInfo(setup.ctx, "DefaultUserSetup", "系统表已存在超级管理员: %s", systemEntity.Val)
		// 查找用户是否存在
		var userEntity *entity.User
		sqlErr = dao.User.Ctx(setup.ctx).Where(&do.User{UserUuid: systemEntity.Val}).Scan(&userEntity)
		if sqlErr != nil {
			blog.BambooError(setup.ctx, "DefaultUserSetup", "查询用户表失败: %s", sqlErr.Error())
			panic(sqlErr)
		}
		// 用户存在跳过
		if userEntity != nil {
			return
		}
		blog.BambooInfo(setup.ctx, "DefaultUserSetup", "系统表存在超级管理员，但用户表中不存在，重新创建用户")
	} else {
		blog.BambooInfo(setup.ctx, "DefaultUserSetup", "系统表不存在超级管理员，开始创建")
	}
	createSuperUser(setup.ctx)
}

// createSuperUser 创建一个系统超级管理员用户。
//
// 参数:
//   - ctx: 上下文对象，用于管理请求范围。
//
// 功能:
//   - 在数据库中创建默认的超级管理员用户，确保其唯一性。
//   - 使用预定义的超级管理员 UUID 和默认参数初始化用户信息。
func createSuperUser(ctx context.Context) {
	// 获取管理员角色
	roleModel := dao.Role.Ctx(ctx)
	var roleEntity *entity.Role
	sqlErr := roleModel.Where(&do.Role{RoleUuid: consts.RoleAdminUUID}).Scan(&roleEntity)
	// 创建用户实体
	superUser := &entity.User{
		UserUuid:  uuid.New().String(),
		Username:  "super_admin",
		Password:  butil.PasswordEncode("super_admin"),
		RealName:  "超级管理员",
		RoleUuid:  roleEntity.RoleUuid,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}

	// 插入用户到数据库
	_, sqlErr = dao.User.Ctx(ctx).OmitEmpty().Insert(superUser)
	if sqlErr != nil {
		blog.BambooError(ctx, "createSuperUser", "创建超级管理员用户失败: %s", sqlErr.Error())
		panic(sqlErr)
	}

	_, sqlErr = dao.System.Ctx(ctx).OmitEmpty().Insert(&entity.System{
		SystemUuid: uuid.New().String(),
		Key:        consts.SystemKeyForSuperUserUUID,
		Val:        superUser.UserUuid,
	})
	if sqlErr != nil {
		blog.BambooError(ctx, "createSuperUser", "创建系统表记录失败: %s", sqlErr.Error())
		panic(sqlErr)
	}

	blog.BambooInfo(ctx, "createSuperUser", "系统超级管理员用户创建成功")
}
