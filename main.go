package main

import (
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

	// 项目初始化
	su := setup.NewSetup(ctx)
	su.DatabaseSetup()

	// 项目启动
	cmd.Main.Run(ctx)
}
