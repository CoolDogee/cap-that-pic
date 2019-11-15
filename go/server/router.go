package server

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.GET("/", hello)
	v1.POST("/getcaption", getCaption)
	v1.GET("/getcaption", getCaption)
	v1.GET("/getTagsfromImage", getTagsFromImage)
	v1.GET("/getTagsFromRemoteImage", getTagsFromRemoteImage)
}
