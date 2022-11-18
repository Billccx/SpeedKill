package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	template := iris.HTML("./backend/web/views", ".html").
		Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	app.HandleDir("./assets", "./backend/web/assets")
	app.OnAnyErrorCode(
		func(ctx iris.Context) {
			ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面有错"))
			ctx.ViewLayout("")
			ctx.View("shared/error.html")
		},
	)

	//todo 注册控制器
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
