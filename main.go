package main

import (
	"github.com/gouef/router"
	bootstrap "github.com/gouef/web-bootstrap"
	"github.com/gouef/web-project/app"
	"github.com/gouef/web-project/controllers"
	"net/http"
	_ "net/http/pprof"
)

func boot() *router.Router {
	b := bootstrap.NewBootstrap()
	b.AddConfig("./config/config.yml")
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
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	r := boot()

	r.Run(":8081")
}
