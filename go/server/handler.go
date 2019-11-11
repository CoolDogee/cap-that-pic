package server

import (
	"github.com/cooldogee/cap-that-pic/data"
	"github.com/cooldogee/cap-that-pic/models"
	"github.com/gin-gonic/gin"

	"strings"
)

func hello(c *gin.Context) {
	c.String(200, "Hello World")
}

func getCaption(c *gin.Context) {
	songs := data.Song(-1, 0).List
	tags := data.Tag(-1, 0).List
	c.String(200, GenerateCaption(&songs, &tags))
}

//GenerateCaption function generates caption from song list and tag list
func GenerateCaption(songs *[]models.Song, tags *[]models.Tag) string {
	lines := GetLyricsLines(songs)
	linePoints := CalculatePoint(&lines, tags)
	indexWithOneLine, valWithOneLine := GetListMaxValue(&linePoints)
	indexWithTwoLines, valWithTwoLines := GetListMaxValueinTwoLines(&linePoints)
	if indexWithOneLine == 0 {
		return lines[0]
	}
	if valWithTwoLines > valWithOneLine {
		return lines[indexWithTwoLines] + "\n" + lines[indexWithTwoLines+1]
	}
	return lines[indexWithOneLine]
}

//CalculatePoint function calculates points of every lines for tags
func CalculatePoint(lines *[]string, tags *[]models.Tag) []float64 {
	linePoints := make([]float64, len(*lines))
	for _, tag := range *tags {
		for index, line := range *lines {
			if strings.Contains(line, tag.Name) {
				linePoints[index] += tag.Confidence
			}
		}
	}
	return linePoints
}

//GetListMaxValue function gets max value and its index from a float64 list
func GetListMaxValue(vals *[]float64) (int, float64) {
	var resIndex int
	var resVal float64
	resVal = 0
	for index, val := range *vals {
		if val >= resVal {
			resVal = val
			resIndex = index
		}
	}
	return resIndex, resVal
}

//GetListMaxValueinTwoLines function gets max two lines' value and the first line's index from a float64 list
func GetListMaxValueinTwoLines(vals *[]float64) (int, float64) {
	var resIndex int
	var resVal float64
	resVal = 0
	for index, val := range (*vals)[:len(*vals)-1] {
		if val+(*vals)[index+1] >= resVal {
			resVal = val
			resIndex = index
		}
	}
	return resIndex, resVal
}

//GetLyricsLines function gets lines from song list
func GetLyricsLines(songs *[]models.Song) []string {
	var allLines []string
	var res []string
	for _, song := range *songs {
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
