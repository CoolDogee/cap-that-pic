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
	"sort"
	"strconv"

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

type Index struct {
	x, y int
}

type Pair struct {
	Key   Index
	Value float64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

var (
	computerVisionContext context.Context
	midCaptionLines =  3
	longCaptionLines = 5
	numberOfResultsToReturn = 10
)

func hello(c *gin.Context) {
	c.String(200, "Hello World")
}

//get url of the image and return the caption generated
func getCaption(c *gin.Context) {
	img := c.Request.URL.Query().Get("fileName")
	captionLength, err := strconv.Atoi(c.Request.URL.Query().Get("length"))
	if err != nil {
		log.Println("Get caption length ERROR: ", err)
		c.String(400, "invalid caption size")
		return
	}
	if captionLength<1 || captionLength>3 {
		log.Println("caption length  error: expected 1-3, given ", captionLength)
		c.String(400, "invalid caption length")
		return
	}
	tags, err := GetTagFromRemoteImage(img)
	if err != nil {
		log.Println("Get caption ERROR: ", err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("Get caption ERROR: %s", err))
		return
	}
	log.Println(tags)
	client := db.ConnectToDB()
	captions := db.GetCaptionsUsingTags(client, tags)
	db.CloseConnectionDB(client)

	c.JSON(200, GenerateCaption(&captions, &tags, captionLength))
}

func getTagsFromRemoteImage(c *gin.Context) {
	url := string(c.Query("fileName"))
	res, err := GetTagFromRemoteImage(url)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, res)
}

func GetTagFromRemoteImage(remoteImgUrl string) ([]models.Tag, error) {
	computerVisionKey := os.Getenv("COMPUTER_VISION_KEY")
	endpointURL := os.Getenv("ENDPOINT_URL")
	log.Println(computerVisionKey)

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

//GenerateCaption function generates caption from Caption list and tag list
func GenerateCaption(captions *[]models.Caption, tags *[]models.Tag, captionLength int) []models.Caption {
	var linePoints [][]float64
	linePoints = CalculatePoint(captions, tags, captionLength)
	pl := GetListMaxValue(&linePoints)
	var result []models.Caption
	for i := 0; i < len(pl); i++ {
		var j int
		var candidate string
		var numLinesInCaption int
		if captionLength==1 {
			candidate = (*captions)[pl[i].Key.x].Text[pl[i].Key.y]
		} else {
			numLinesInCaption = midCaptionLines
			if captionLength==3 {
				numLinesInCaption = longCaptionLines
			}
			if pl[i].Value==0.0 {
				continue
			}
			candidate = (*captions)[pl[i].Key.x].Text[pl[i].Key.y]
			for count:=1; count<numLinesInCaption; count++ {
				candidate += ". "
				candidate += (*captions)[pl[i].Key.x].Text[pl[i].Key.y + count]
			}
		}
		alreadyPresent := false
		for j = 0; j < len(result); j++ {
			if candidate == result[j].Text[0] {
				alreadyPresent = true
				break
			}
		}
		if !alreadyPresent {
			candidateCaption := models.Caption{
				Text:          []string{candidate},
				Tags:          (*captions)[pl[i].Key.x].Tags,
				UserGenerated: true,
			}
			result = append(result, candidateCaption)
		}
		if len(result) == numberOfResultsToReturn {
			break
		}
	}
	return result
}

// CalculatePoint function calculates points of every lines for tags
func CalculatePoint(captions *[]models.Caption, tags *[]models.Tag, captionLength int) [][]float64 {
	linePoints := make([][]float64, len(*captions))
	for index, caption := range *captions {
		linePoints[index] = make([]float64, len(caption.Text))
	}
	for _, tag := range *tags {
		for indexX, caption := range *captions {
			for indexY, line := range caption.Text {
				if strings.Contains(line, tag.Name) {
					linePoints[indexX][indexY] += tag.Confidence
				}
			}
		}
	}

	var numLinesInCaption int
	if captionLength==1 {
		return linePoints
	} else if captionLength==2 {
		numLinesInCaption = midCaptionLines
	} else {
		numLinesInCaption = longCaptionLines
	}
	for indexX, caption := range *captions {
		for indexY:=0; indexY<len(caption.Text)-numLinesInCaption+1; indexY++ {
			for nextLine:=1; nextLine<numLinesInCaption; nextLine++ {
				linePoints[indexX][indexY] += linePoints[indexX][indexY + nextLine]
			}
		}
	}
	return linePoints
}

//GetListMaxValue function gets max value and its index from a float64 matrix
func GetListMaxValue(vals *[][]float64) PairList {
	m := make(map[Index]float64)
	for indexX, valsLine := range *vals {
		for indexY, val := range valsLine {
			var index Index = Index{indexX, indexY}
			m[index] = val
		}
	}
	pl := make(PairList, len(m))
	i := 0
	for k, v := range m {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
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
		c.String(http.StatusBadRequest, "Not a valid URL.")
	}
	if uri.Scheme != "http" && uri.Scheme != "https" {
		c.String(http.StatusBadRequest, "Not a valid URL.")
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
