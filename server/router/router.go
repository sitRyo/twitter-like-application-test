package router

import (
	"github.com/gin-gonic/gin"
)

func router(e *gin.Engine) {
	api := e.Group("/api")

	// block
	{
		// /api/upload
		// second param: function
		api.POST("/upload", storeTweet)
	}
}

func storeTweet(e *gin.Context) {

}
