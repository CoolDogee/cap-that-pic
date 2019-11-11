package main

import (
	"net/http"

	"github.com/cooldogee/cap-that-pic/server"
)

func main() {
	http.HandleFunc("/api/v1", server.hello)
	http.HandleFunc("/api/v1/getcaption", server.getCaption)
}
