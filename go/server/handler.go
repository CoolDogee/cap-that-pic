package server

import (
	"github.com/gin-gonic/gin"

	"strings"
)

func hello(c *gin.Context) {
	c.String(200, "Hello World")
}

func generateCaption(lyrics []string, tags []string) string {
	return lyrics[0]
}

func getLyricsLines(lyrics []string) []string {
	var allLines []string
	for _, lyric := range lyrics {
		lines := strings.Split(lyric, "\n")
		allLines = append(allLines, lines...)
	}
	return allLines
}
