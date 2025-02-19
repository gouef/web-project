package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gouef/diago"
	"github.com/gouef/diago/extensions"
	"github.com/gouef/router"
	extensions2 "github.com/gouef/router/extensions"
	"path/filepath"
	//web_bootstrap "github.com/gouef/web-bootstrap"
	"github.com/gouef/web-project/app"
	"github.com/gouef/web-project/controllers"
	"github.com/gouef/web-project/handlers"
)

func boot() *router.Router {
	/*b := web_bootstrap.NewBootstrap()
	b.Boot()*/

	r := router.NewRouter()
	n := r.GetNativeRouter() /*
		n.Use(func(c *gin.Context) {
			c.Writer.Header().Del("Purpose")
			c.Writer.Header().Set("Purpose", "prefetch")
			c.Writer.Header().Set("X-DNS-Prefetch-Control", "on")
			c.Next()
		})*/

	if !r.IsRelease() {
		d := diago.NewDiago()
		d.AddExtension(extensions.NewLatencyExtension())
		d.AddExtension(extensions2.NewDiagoRouteExtension(r))

		n.Use(diago.Middleware(r, d))
	}

	templateHandler := &handlers.TemplateHandler{Router: r}

	// Inicializace šablon
	templateHandler.Initialize()
	n.SetTrustedProxies([]string{"127.0.0.1"})
	//n.LoadHTMLGlob("views/templates/**/*")
	n.HTMLRender = loadTemplates("./views/templates", templateHandler)
	n.Static("/static", "./static")
	n.Static("/assets", "./static/assets")
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

func loadTemplates(templatesDir string, templateHandler *handlers.TemplateHandler) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	funcMap := templateHandler.GetFuncMap()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.gohtml")
	if err != nil {
		panic(err.Error())
	}
	rootLayouts, err := filepath.Glob(templatesDir + "/layout.gohtml")
	if err != nil {
		panic(err.Error())
	}
	layouts = append(layouts, rootLayouts...)

	includes, err := filepath.Glob(templatesDir + "/includes/*.gohtml")
	if err != nil {
		panic(err.Error())
	}

	//	includesNamed, err := filepath.Glob(templatesDir + "/**/*.gohtml")
	/*
		if err != nil {
			panic(err.Error())
		}
		includes = append(includes, includesNamed...)
	*/
	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, includes...)
		r.AddFromFilesFuncs(filepath.Base(include), funcMap, files...)
	}
	return r
}
