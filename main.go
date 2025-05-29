package main

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	_ "xiao-yi-transit/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "xiao-yi-transit/internal/logic"

	"xiao-yi-transit/internal/config/cmd"
	"xiao-yi-transit/internal/config/setup"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	// 生成启动 Context
	ctx := gctx.GetInitCtx()

	gres.Dump()

	// 检查程序当前文件夹是否存在配置文件
	if !gfile.Exists("config.yaml") {
		getByte := gres.GetContent("config/config.template.yaml")
		if getByte == nil {
			panic("配置文件不存在，请检查程序包是否完整")
		}
		// 如果不存在则将配置文件从程序包中复制到当前目录
		if err := gfile.PutBytes("config.yaml", getByte); err != nil {
			panic("无法创建配置文件，请检查当前目录权限")
		}
	}

	// 项目初始化
	su := setup.NewSetup(ctx)
	su.DatabaseSetup()
	su.DefaultRoleSetup()
	su.DefaultUserSetup()

	// 项目启动
	cmd.Main.Run(ctx)
}
