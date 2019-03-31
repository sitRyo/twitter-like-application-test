package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sitRyo/server/router"
)

func main() {
	api := gin.Default()
	router.Register(api)
	api.Run(":8080")
}
