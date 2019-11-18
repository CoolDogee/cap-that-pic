package server

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/cooldogee/cap-that-pic/data"
	"github.com/cooldogee/cap-that-pic/db"
	"github.com/cooldogee/cap-that-pic/models"
	"github.com/gin-gonic/gin"

	"strings"

	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision"
	"github.com/Azure/go-autorest/autorest"
)

type Image struct {
	URL string
}

type Caption struct {
	Content string
}

var computerVisionContext context.Context

func hello(c *gin.Context) {
	c.String(200, "Hello World")
}

//get url of the image and return the caption generated
func getCaption(c *gin.Context) {
	img := c.Request.URL.Query().Get("fileName")
	tags, err := GetTagFromRemoteImage(img)
	if err != nil {
		log.Println("Get caption ERROR: ", err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("Get caption ERROR: %s", err))
		return
	}
	client := db.ConnectToDB()
	songs := db.GetLyricsUsingTags(client, tags)
	db.CloseConnectionDB(client)

	var caption Caption
	caption.Content = GenerateCaption(&songs, &tags)
	c.JSON(200, caption.Content)
}


func getTagsFromRemoteImage(c *gin.Context) {
	url := string(c.Query("fileName"))
	res, err := GetTagFromRemoteImage(url)
	if err!= nil {
		fmt.Println(err)
	}
	c.JSON(200, res)
}

func GetTagFromRemoteImage(remoteImgUrl string) ([]models.Tag, error) {
	computerVisionKey := os.Getenv("COMPUTER_VISION_KEY")
	endpointURL := os.Getenv("ENDPOINT_URL")

	computerVisionClient := computervision.New(endpointURL)
	computerVisionClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(computerVisionKey)

	computerVisionContext = context.Background()

	return TagRemoteImage(computerVisionClient, remoteImgUrl)
}

func TagRemoteImage(client computervision.BaseClient, remoteImageURL string) ([]models.Tag, error) {
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL
	remoteImageTags, err := client.TagImage(
		computerVisionContext,
		remoteImage,
		"")
	if err != nil {
		return nil, err
	}
	var tags []models.Tag
	for _, caption := range *remoteImageTags.Tags {
		tag := models.Tag{*caption.Name, *caption.Confidence * 100}
		tags = append(tags, tag)
	}
	return tags, nil
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
	if indexY == 0 {
		return lines[indexX][0] + "\n" + lines[indexX][1] + "\n" + lines[indexX][2]
	}
	if indexY == len(lines[indexX])-1 {
		return lines[indexX][indexY-2] + "\n" + lines[indexX][indexY-1] + "\n" + lines[indexX][indexY]
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

// getTagsFromImage for /getTagsfromImage endpoint
func getTagsFromImage(c *gin.Context) {
	computerVisionKey := os.Getenv("COMPUTER_VISION_KEY")
	endpointURL := os.Getenv("ENDPOINT_URL")

	computerVisionClient := computervision.New(endpointURL)
	computerVisionClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(computerVisionKey)

	computerVisionContext = context.Background()

	// Analyze a local image

	baseDir, err := os.Getwd()

	if err != nil {
		log.Println("Get tag from image ERROR: ", err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("Get tag from image ERROR: %s", err))
		return
	}

	imgName := "animals.jpg"

	localImagePath := baseDir + "/../../client/public/uploads/" + imgName

	tags, err := TagLocalImage(computerVisionClient, localImagePath)
	if err != nil {
		log.Println("Tag local image ERROR: ", err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("Tag local image ERROR: %s", err))
		return
	}
	c.JSON(200, tags)
}

// GetTagsFromImage return tags for an image
// Utility function for getCaption
func GetTagsFromImage(img string) ([]models.Tag, error) {
	computerVisionKey := os.Getenv("COMPUTER_VISION_KEY")
	endpointURL := os.Getenv("ENDPOINT_URL")

	computerVisionClient := computervision.New(endpointURL)
	computerVisionClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(computerVisionKey)

	computerVisionContext = context.Background()

	// Analyze a local image

	baseDir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	localImagePath := baseDir + "/../client/public/uploads/" + img
	return TagLocalImage(computerVisionClient, localImagePath)
}

func TagLocalImage(client computervision.BaseClient, localImagePath string) ([]models.Tag, error) {
	var localImage io.ReadCloser
	localImage, err := os.Open(localImagePath)
	if err != nil {
		return nil, err
	}

	localImageTags, err := client.TagImageInStream(
		computerVisionContext,
		localImage,
		"")
	if err != nil {
		return nil, err
	}

	var tags []models.Tag

	for _, caption := range *localImageTags.Tags {
		tag := models.Tag{*caption.Name, *caption.Confidence * 100}
		tags = append(tags, tag)
	}
	return tags, nil
}

func validateImageURL(c *gin.Context) {
	url1 := c.Request.URL.Query().Get("fileName")
	uri, err := url.Parse(url1)
	if err != nil {
		// Error
		c.String(http.StatusBadRequest, "Not a validate URL.")
	}
	if uri.Scheme != "http" && uri.Scheme != "https" {
		c.String(http.StatusBadRequest, "Not a validate URL.")
		return
	}
	r, err := http.Get(url1)
	if err != nil {
		fmt.Println("URL cannot reach: ", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("URL cannot reach: %s", err))
		return
	}
	if r == nil || r.Body == nil {
		log.Println("No body found")
		c.String(http.StatusBadRequest, "No body found")
		return
	}
	if r.StatusCode != http.StatusOK {
		fmt.Println("URL cannot reach: ", r.StatusCode)
		c.String(http.StatusBadRequest, fmt.Sprintf("URL cannot reach: %d", r.StatusCode))
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Cannot read the file: ", err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("Cannot read the file: %s", err.Error()))
		return
	}
	// log.Println(http.DetectContentType(buff)) // do something based on your detection.
	log.Println(http.DetectContentType(body))
	if !strings.HasPrefix(http.DetectContentType(body), "image") {
		c.String(http.StatusBadRequest, "Not an image.")
		return
	}
	c.String(http.StatusOK, "Valid image.")
}
