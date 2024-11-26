package app

import router2 "github.com/Gouef/router"

func RouterFactory() {
	router := router2.Router{}

	router.Run()
}
