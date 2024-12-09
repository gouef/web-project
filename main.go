package main

import (
	"github.com/gouef/diago"
	"github.com/gouef/diago/extensions"
	"github.com/gouef/router"
	extensions2 "github.com/gouef/router/extensions"
	"github.com/gouef/web-project/app"
	"github.com/gouef/web-project/controllers"
	"github.com/gouef/web-project/handlers"
)

func boot() *router.Router {
	r := router.NewRouter()
	n := r.GetNativeRouter()
	d := diago.NewDiago()
	d.AddExtension(extensions.NewDiagoLatencyExtension())
	d.AddExtension(extensions2.NewDiagoRouteExtension(r))

	n.Use(diago.DiagoMiddleware(r, d))

	templateHandler := &handlers.TemplateHandler{Router: r}

	// Inicializace šablon
	templateHandler.Initialize()
	n.SetTrustedProxies([]string{"127.0.0.1"})

	n.LoadHTMLGlob("views/templates/**/*.gohtml")
	n.Static("/static", "./static")
	rl := app.RouterFactory()

	r.SetErrorHandler(404, controllers.Error404)
	r.SetErrorHandler(500, controllers.Error500)
	r.SetDefaultErrorHandler(controllers.Error404)
	r.AddRouteList(rl)

	return r
}

func main() {
	r := boot()
	r.Run(":8080")
}
