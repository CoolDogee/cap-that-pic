package main

import (
	"github.com/cooldogee/cap-that-pic/db"
	"github.com/cooldogee/cap-that-pic/server"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./web", true)))
	server.SetupRoutes(router)
	db.SetupDB()
	router.Run()
}
