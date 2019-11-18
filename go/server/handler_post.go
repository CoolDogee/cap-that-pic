package server

import (
	"net/http"

	"github.com/cooldogee/cap-that-pic/db"
	"github.com/cooldogee/cap-that-pic/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func postCaptionHandler(c *gin.Context) {
	var caption models.Caption
	if err := c.BindJSON(&caption); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "info": ""})
		return
	}
	caption.ID = primitive.NewObjectID()
	client := db.ConnectToDB()
	err := db.AddCaptionToDB(client, &caption)
	db.CloseConnectionDB(client)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "info": ""})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Caption struct created!", "info": caption.ID})
}

func postPostHandler(c *gin.Context) {
	var post models.Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "info": ""})
		return
	}
	post.ID = primitive.NewObjectID()
	client := db.ConnectToDB()
	// check caption id exists before add post to database
	_, err := db.GetCaptionByID(client, post.CaptionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "caption id not found.", "info": ""})
		return
	}
	err = db.AddPostToDB(client, &post)
	db.CloseConnectionDB(client)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "info": ""})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Post struct created!", "info": post.ID})
}

func getPostHandler(c *gin.Context) {
	id := c.Param("id")
	client := db.ConnectToDB()
	post, err := db.GetPostByID(client, id)
	db.CloseConnectionDB(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "info": ""})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Get post successfully!", "info": post})
}

func getCaptionHandler(c *gin.Context) {
	id := c.Param("id")
	client := db.ConnectToDB()
	caption, err := db.GetCaptionByID(client, id)
	db.CloseConnectionDB(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "info": ""})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Get caption successfully!", "info": caption})
}
