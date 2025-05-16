package cmd

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/bhandler/bhook"
	"github.com/XiaoLFeng/bamboo-utils/bhandler/bmiddle"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"xiao-yi-transit/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 钩子扩展
			s.BindHookHandler("/api/*", ghttp.HookBeforeServe, bhook.BambooHookDefaultCors)
			s.BindHookHandler("/api/*", ghttp.HookBeforeServe, bhook.BambooHookRequestInfo)

			// 后端接口
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(bmiddle.BambooHandlerResponse)

				group.Group("/v1", func(v1 *ghttp.RouterGroup) {
					v1.Bind()
				})
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
				)
			})

			s.Run()
			return nil
		},
	}
)
