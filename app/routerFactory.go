package app

import router2 "github.com/gouef/router"

func RouterFactory() {
	router := router2.Router{}

	router.Run()
}
