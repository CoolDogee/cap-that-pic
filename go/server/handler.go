package server

import (
	"github.com/cooldogee/cap-that-pic/data"
	"github.com/cooldogee/cap-that-pic/models"
	"github.com/gin-gonic/gin"

	"strings"
)

type Image struct {
	URL string
}

type Caption struct {
	Content string
}

func hello(c *gin.Context) {
	c.String(200, "Hello World")
}

//get url of the image and return the caption generated
func getCaption(c *gin.Context) {
	var img Image
	c.ShouldBindJSON(&img)
	tags := getTagsOfImage(img.URL)
	songs := getLyricsFromTags(&tags)
	var caption = Caption{Content: GenerateCaption(&songs, &tags)}
	c.JSON(200, caption)
}

//********** Use API get tags of an image Here ***************
func getTagsOfImage(url string) []models.Tag {
	tags := data.Tag(-1, 0).List
	return tags
}

//********** Use DB get lyrics from tags Here ***************
func getLyricsFromTags(tags *[]models.Tag) []models.Song {
	songs := data.Song(-1, 0).List
	return songs
}

//GenerateCaption function generates caption from song list and tag list
func GenerateCaption(songs *[]models.Song, tags *[]models.Tag) string {
	lines := GetLyricsLines(songs)
	linePoints := CalculatePoint(&lines, tags)
	indexX, indexY, _ := GetListMaxValue(&linePoints)
	//	indexWithTwoLines, valWithTwoLines := GetListMaxValueinTwoLines(&linePoints)
	if indexX == 0 && indexY == 0 {
		return lines[0][0] + "\n" + lines[0][1] + "\n" + lines[0][2]
	}
	return lines[indexX][indexY-1] + "\n" + lines[indexX][indexY] + "\n" + lines[indexX][indexY+1]
}

//CalculatePoint function calculates points of every lines for tags
func CalculatePoint(lines *[][]string, tags *[]models.Tag) [][]float64 {
	linePoints := make([][]float64, len(*lines))
	for index, linesInSong := range *lines {
		linePoints[index] = make([]float64, len(linesInSong))
	}
	for _, tag := range *tags {
		for indexX, linesInSong := range *lines {
			for indexY, line := range linesInSong {
				if strings.Contains(line, tag.Name) {
					linePoints[indexX][indexY] += tag.Confidence
				}
			}
		}
	}
	return linePoints
}

//GetListMaxValue function gets max value and its index from a float64 matrix
func GetListMaxValue(vals *[][]float64) (int, int, float64) {
	var resIndexX, resIndexY int
	var resVal float64
	resVal = 0
	for indexX, valsLine := range *vals {
		for indexY, val := range valsLine {
			if val >= resVal {
				resVal = val
				resIndexX = indexX
				resIndexY = indexY
			}
		}
	}
	return resIndexX, resIndexY, resVal
}

// //GetListMaxValueinTwoLines function gets max two lines' value and the first line's index from a float64 list
// func GetListMaxValueinTwoLines(vals *[]float64) (int, float64) {
// 	var resIndex int
// 	var resVal float64
// 	resVal = 0
// 	for index, val := range (*vals)[:len(*vals)-1] {
// 		if val+(*vals)[index+1] >= resVal {
// 			resVal = val
// 			resIndex = index
// 		}
// 	}
// 	return resIndex, resVal
// }

//GetLyricsLines function gets lines from song list
func GetLyricsLines(songs *[]models.Song) [][]string {
	allLines := make([][]string, len(*songs))
	var res [][]string
	for index, song := range *songs {
		lines := strings.Split(song.Lyrics, "\n")
		allLines[index] = lines
	}
	for _, lines := range allLines {
		var tmp []string
		for _, line := range lines {
			if len(line) != 0 && line[0] != '[' {
				tmp = append(tmp, line)
			}
		}
		if len(tmp) != 0 {
			res = append(res, tmp)
		}
	}

	return res
}
