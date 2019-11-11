package main

import (
	"github.com/cooldogee/cap-that-pic/server"
)

func main() {
	//	client := db.ConnectToDB()
	//	db.AddLyricsToDB(client)
	//	db.CloseConnectionDB(client)
	router := server.CreateRouter()
	server.StartServer(router)
}
