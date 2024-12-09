package app

import (
	router "github.com/gouef/router"
	"github.com/gouef/web-project/controllers"
)

func RouterFactory() *router.RouteList {
	app := controllers.DefaultController{}
	l := router.NewRouteList()
	l.Add("home", "/", app.Index, router.Get)
	l.Add("ping", "/ping", app.Ping, router.Get)
	l.Add("users:detail", "/users/:id", app.Index, router.Get)

	return l
}
