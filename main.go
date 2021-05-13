package main

import (
	"Dream/conf"
	"Dream/router"
)

func main() {
	router.Router.Run(":" + conf.Config.Net.Port)
}
