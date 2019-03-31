package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	api.Run(":8080")
}
