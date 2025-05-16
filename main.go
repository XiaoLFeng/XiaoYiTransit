package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "xiao-yi-transit/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"xiao-yi-transit/internal/config/cmd"
	"xiao-yi-transit/internal/config/setup"
)

func main() {
	// 生成启动 Context
	ctx := gctx.GetInitCtx()

	// 项目初始化
	su := setup.NewSetup(ctx)
	su.DatabaseSetup()

	// 项目启动
	cmd.Main.Run(ctx)
}
