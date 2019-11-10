package server

import (
	"github.com/cooldogee/cap-that-pic/models"
	"github.com/gin-gonic/gin"

	"strings"
)

func hello(c *gin.Context) {
	c.String(200, "Hello World")
}

func GenerateCaption(songs []models.Song, tags []models.Tag) string {
	return " "
}

func GetLyricsLines(songs []models.Song) []string {
	var allLines []string
	var res []string
	for _, song := range songs {
		lines := strings.Split(song.Lyrics, "\n")
		allLines = append(allLines, lines...)
	}
	for _, line := range allLines {
		if len(line) != 0 && line[0] != '[' {
			res = append(res, line)
		}
	}
	return res
}
