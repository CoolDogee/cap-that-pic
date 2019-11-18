package main

import (
	"github.com/cooldogee/cap-that-pic/db"
	"github.com/cooldogee/cap-that-pic/server"
	// "github.com/gin-gonic/contrib/static"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags)
	router := gin.Default()
	// router.Use(static.Serve("/", static.LocalFile("./web", true)))
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))
	server.SetupRoutes(router)
	db.SetupDB()
	router.Run()
}
