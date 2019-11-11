package server

import (
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.GET("/", hello)
<<<<<<< Updated upstream
	v1.POST("/getcaption", getCaption)
=======
	v1.GET("/getcaption", getCaption)
	v1.GET("/gettokensfromimage", getTagsFromImage)
>>>>>>> Stashed changes
}
