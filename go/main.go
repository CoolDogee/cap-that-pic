package main

import (
	"fmt"

	"github.com/cooldogee/cap-that-pic/db"
)

func main() {
	client := db.ConnectToDB()
	// db.AddLyricsToDB(client)
	var tags []string
	tags = append(tags, "Hello", "World", "Wake", "Sleep")
	fmt.Println(db.GetLyricsUsingTags(client, tags))
	db.CloseConnectionDB(client)
	// router := server.CreateRouter()
	// server.StartServer(router)
}
