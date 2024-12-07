package main

import (
	"github.com/gouef/router"
	"github.com/gouef/web-project/app"
)

func main() {
	r := router.NewRouter()
	rl := app.RouterFactory()

	r.AddRouteList(rl)
	r.Run(":8080")
}
