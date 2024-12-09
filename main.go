package main

import (
	"github.com/gouef/router"
	"github.com/gouef/web-project/app"
	"github.com/gouef/web-project/controllers"
	"github.com/gouef/web-project/handlers"
	"github.com/gouef/web-project/middleware"
)

func boot() *router.Router {
	r := router.NewRouter()
	n := r.GetNativeRouter()
	d := middleware.NewDiago()
	d.AddExtension(middleware.NewDiagoLatencyExtension())
	d.AddExtension(middleware.NewDiagoRouteExtension(r))

	n.Use(middleware.DiagoMiddleware(r, d))

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
