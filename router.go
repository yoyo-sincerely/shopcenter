package shopcenter

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/service"
)

func (plug *Shopcenter) initRouter(srv service.List) *context.App {

	app := context.NewApp()
	authRoute := app.Group("/", auth.Middleware(plug.Conn))

	authRoute.GET("/example", plug.guard.Example, plug.handler.Example)

	return app
}
