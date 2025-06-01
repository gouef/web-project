package main

import (
	"github.com/gouef/router"
	bootstrap "github.com/gouef/web-bootstrap"
	"github.com/gouef/web-project/app"
	"github.com/gouef/web-project/controllers"
)

func boot() *router.Router {
	b := bootstrap.NewBootstrap()
	b.Boot()

	r := b.GetRouter()
	// r.EnablePrefetch()
	rl := app.RouterFactory()

	r.SetErrorHandler(404, controllers.Error404)
	r.SetErrorHandler(500, controllers.Error500)
	r.SetDefaultErrorHandler(controllers.Error404)
	r.AddRouteList(rl)

	return r
}

func main() {
	r := boot()

	r.Run(":8081")
}
