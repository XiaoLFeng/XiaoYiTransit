package setup

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

var databaseList = []string{
	"xf_role.sql",                 // 角色表
	"xf_station.sql",              // 站点表
	"xf_vehicle.sql",              // 车辆表
	"xf_driver.sql",               // 司机表
	"xf_route.sql",                // 线路表
	"xf_user.sql",                 // 用户表（依赖角色表）
	"xf_user_token.sql",           // token表
	"xf_driver_qualification.sql", // 司机资质表（依赖司机表）
	"xf_driver_shift.sql",         // 司机班次表（依赖司机表）
	"xf_route_station.sql",        // 线路站点关联表（依赖线路表和站点表）
	"xf_maintenance.sql",          // 维护记录表（依赖车辆表）
	"xf_vehicle_fault.sql",        // 车辆故障表（依赖车辆表）
	"xf_maintenance_item.sql",     // 维护项目表（可能依赖维护记录表）
	"xf_schedule_template.sql",    // 调度模板表
	"xf_schedule.sql",             // 调度表（依赖线路、车辆和司机表）
}

// DatabaseSetup 完成数据库的初始化配置。
//
// 执行步骤包括:
//   - 获取数据库连接信息
//   - 检查指定的数据库表是否存在，若不存在则通过执行 SQL 文件来创建表
//   - SQL 文件按需拆分并逐条执行，完成表结构的生成
//
// 错误:
//   - 数据库配置错误或连接失败
//   - SQL 文件解析或执行过程中出现的错误
func (setup *SystemSetup) DatabaseSetup() {
	// 1. 获取数据库连接
	db := g.DB()
	// 获取数据库名字
	getName, configErr := g.Cfg().Get(setup.ctx, "database.default.name")
	if configErr != nil {
		panic(configErr)
	}

	// 2. 遍历数据库列表，执行 SQL 文件
	for _, fileName := range databaseList {
		hasExist, sqlErr := db.Model("information_schema.TABLES").
			Where("TABLE_SCHEMA", getName).
			Where("TABLE_NAME", gstr.Trim(fileName, ".sql")).Exist()
		if sqlErr != nil {
			g.Log().Errorf(setup.ctx, "查询数据库表 %s 失败: %v", fileName, sqlErr)
			return
		}
		if !hasExist {
			getFile := g.Res().Get("template/sql/" + fileName)
			getSql := string(getFile.Content())
			// 拆分为单个 SQL 语句
			sqlStatements := gstr.Split(getSql, ";")
			for _, sql := range sqlStatements {
				sql = gstr.Trim(sql)
				if sql == "" {
					continue
				}
				// 执行 SQL 语句
				_, err := db.Exec(setup.ctx, sql)
				if err != nil {
					g.Log().Errorf(setup.ctx, "执行 SQL 失败: %s, 错误: %v", sql, err)
					return
				}
			}
		}
	}
}
