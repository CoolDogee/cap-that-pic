package main

import (
	"github.com/cooldogee/cap-that-pic/server"
)

func main() {
	router := server.CreateRouter()
	server.StartServer(router)
}
