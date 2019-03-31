package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Register function
func Register(e *gin.Engine) {
	api := e.Group("/api")

	// block
	{
		// /api/upload
		// second param: function
		api.POST("/upload", storeTweet)
	}
}

func storeTweet(e *gin.Context) {
	e.Request.ParseForm()
	fmt.Println(e.Request.Form)
}
