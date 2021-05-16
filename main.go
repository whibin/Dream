package main

import (
	"Dream/conf"
	"Dream/router"
	"Dream/services"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Minute * 30)
	go func() {
		for range ticker.C {
			services.LikeSave2MySQL()
		}
	}()

	router.Router.Run(":" + conf.Config.Net.Port)
}
