package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sitRyo/twitter-like-app/server/router"
)

func main() {
	// create empty engine
	api := gin.Default()
	// register webApi
	router.Register(api)
	// run localhost port 8080
	api.Run(":8080")
}
