package server

import (
	"github.com/cooldogee/cap-that-pic/data"
	"github.com/cooldogee/cap-that-pic/db"
	"github.com/cooldogee/cap-that-pic/models"
	"github.com/gin-gonic/gin"

	"strings"
)

func hello(c *gin.Context) {
	c.String(200, "Hello World")
}

func getCaption(c *gin.Context) {
	// songs := data.Song(-1, 0).List
	tags := data.Tag(-1, 0).List
	client := db.ConnectToDB()
	songs := db.GetLyricsUsingTags(client, tags)
	db.CloseConnectionDB(client)
	c.String(200, GenerateCaption(songs, tags))
}

func GenerateCaption(songs []models.Song, tags []models.Tag) string {
	lines := GetLyricsLines(songs)
	linePoints := make([]float64, len(lines))

	for _, tag := range tags {
		for index, line := range lines {
			if strings.Contains(line, tag.Name) {
				linePoints[index] += tag.Confidence
			}
		}
	}
	index, _ := GetListMaxValue(linePoints)
	return lines[index]
}

func GetListMaxValue(vals []float64) (int, float64) {
	var resIndex int
	var resVal float64
	resVal = 0
	for index, val := range vals {
		if val >= resVal {
			resVal = val
			resIndex = index
		}
	}
	return resIndex, resVal
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
