package app

import (
	router "github.com/gouef/router"
)

func RouterFactory() *router.RouteList {
	app := App{}
	l := router.NewRouteList()
	l.Add("/ping", app.Ping, router.Get)
	l.Add("/", app.Homepage, router.Get)

	return l
}
