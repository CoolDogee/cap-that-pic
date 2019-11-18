package db

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"os"
	"time"

	"github.com/cooldogee/cap-that-pic/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToDB makes connection with database
func ConnectToDB() *mongo.Client {
	// Set client options
	// RUNTIME_ENV = "local" or "docker"
	runtimeEnv := os.Getenv("RUNTIME_ENV")
	clientOptions := options.Client().ApplyURI("")
	if runtimeEnv == "LOCAL" {
		// for local mongo db
		log.Println("Environment variable RUNTIME_ENV is LOCAL, use db url localhost:27017")
		clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	} else if runtimeEnv == "DOCKER" || runtimeEnv == "" {
		// for docker mongo db
		if runtimeEnv == "DOCKER" {
			log.Println("Environment variable RUNTIME_ENV is DOCKER, use db url mongo:27017 (docker localhost)")
		} else {
			log.Println("Environment variable RUNTIME_ENV is undefined, use db url mongo:27017 (docker localhost)")
		}
		clientOptions = options.Client().ApplyURI("mongodb://mongo:27017")
	} else {
		log.Fatal("Wrong RUNTIME_ENV")
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Panicln("Fail to connect to MongoDB.")
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println("Fail to connect to MongoDB: ", err)
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	return client
}

// CloseConnectionDB ends the connection with database
func CloseConnectionDB(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Println("Fail to close MongoDB: ", err)
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB closed.")
}

// AddLyricsToDB makes a connection with the NoSQL database
func AddLyricsToDB(client *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	// collection := client.Database("CAP-THAT-PIC").Collection("Lyrics")
	collection := client.Database("CAP-THAT-PIC").Collection("Captions")

	var caption models.Caption
	filter := bson.M{"type": "song"}
	err := collection.FindOne(ctx, filter).Decode(&caption)
	log.Println(caption)
	if err==nil {
		return
	}

	// n, err := collection.DeleteMany(ctx, bson.M{})

	// if err != nil {
	// 	log.Println("DeleteMany ERROR:", err)
	// } else {
	// 	log.Println("Number of documents removed: ", n)
	// }

	// Ref the location of lyrics in the dockerfile
	// byteValues, err := ioutil.ReadFile("../lyrics/lyrics.json")
	byteValues, err := ioutil.ReadFile("./lyrics.json")

	if err != nil {
		// Print any IO errors with the .json file
		log.Println("ioutil.ReadFile ERROR:", err)
	}

	// var docs []models.Song
	// err = json.Unmarshal(byteValues, &docs)

	// Print MongoDB docs object type
	// log.Println("nMongoFields Docs:", reflect.TypeOf(docs), len(docs))

	var songs []map[string]interface{}
	json.Unmarshal([]byte(byteValues), &songs)

	// for i := range docs {
	// 	doc := docs[i]
	// 	result, err := collection.InsertOne(ctx, doc)

	// 	if err != nil {
	// 		log.Println("InsertOne ERROR:", err)
	// 	} else {
	// 		log.Println("InsertOne() API result:", result)
	// 	}
	// }

	for _, song := range songs {
		caption := models.Caption{
			Text:          strings.Split(song["lyrics"].(string), "\n"),
			Src:           song["chartURL"].(string),
			Type:          "song",
			Tags:          []string{song["artist"].(string)},
			UserGenerated: false,
		}
		caption.ID = primitive.NewObjectID()
		result, err := collection.InsertOne(ctx, caption)
		if err != nil {
			log.Println("InsertOne ERROR:", err)
		} else {
			log.Println("InsertOne() API result:", result)
		}
	}
}

// AddPoemsToDB adds poems to mongodb database
func AddPoemsToDB(client *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	// collection := client.Database("CAP-THAT-PIC").Collection("Lyrics")
	collection := client.Database("CAP-THAT-PIC").Collection("Captions")

	var caption models.Caption
	filter := bson.M{"type": "poem"}
	err := collection.FindOne(ctx, filter).Decode(&caption)
	log.Println(caption)
	if err==nil {
		return
	}

	files, err := filepath.Glob("../poem/*.json")

	if err != nil {
		log.Println("filepath.Glob ERROR:", err)
	}
	var poem map[string]interface{}

	for _, f := range files {
		byteValues, err := ioutil.ReadFile(f)
		if err != nil {
			// Print any IO errors with the .json file
			log.Println("ioutil.ReadFile ERROR:", err)
		}
		
		json.Unmarshal([]byte(byteValues), &poem)

		// log.Printf("%#v", poem["text"])

		caption := models.Caption{
			Src:           poem["reference"].(string),
			Type:          "poem",
			UserGenerated: false,
		}
		caption.ID = primitive.NewObjectID()
		for _, line := range poem["text"].([]interface{}) {
			caption.Text = append(caption.Text, line.(string))
		}
		for _, tag := range poem["keywords"].([]interface{}) {
			caption.Tags = append(caption.Tags, tag.(string))
		}
		caption.Tags = append(caption.Tags, poem["author"].(string))
		result, err := collection.InsertOne(ctx, caption)
		if err != nil {
			log.Println("InsertOne ERROR:", err)
		} else {
			log.Println("InsertOne() API result:", result)
		}
	}
}

// AddMovieQuotesToDB adds movie quotes to mongodb
func AddMovieQuotesToDB(client *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	collection := client.Database("CAP-THAT-PIC").Collection("Captions")

	var caption models.Caption
	filter := bson.M{"type": "movie"}
	err := collection.FindOne(ctx, filter).Decode(&caption)
	log.Println(caption)
	if err==nil {
		return
	}

	byteValues, err := ioutil.ReadFile("../movie/movie_quotes.json")

	if err != nil {
		log.Println("ioutil.ReadFile ERROR:", err)
	}

	var movies []map[string]interface{}
	json.Unmarshal([]byte(byteValues), &movies)

	for _, movie := range movies {
		caption := models.Caption{
			Text:          []string{movie["text"].(string)},
			Src:           movie["movie"].(string),
			Type:          "movie",
			UserGenerated: false,
		}
		caption.ID = primitive.NewObjectID()
		result, err := collection.InsertOne(ctx, caption)
		if err != nil {
			log.Println("InsertOne ERROR:", err)
		} else {
			log.Println("InsertOne() API result:", result)
		}
	}
}

// GetLyricsUsingTags returns list of songs which have atleast one tag in their title
func GetLyricsUsingTags(client *mongo.Client, tags []models.Tag) []models.Song {
	var songs []models.Song

	for i := range tags {
		songs = append(songs, GetLyricsUsingTag(client, tags[i].Name)...)
	}

	return songs
}

// GetLyricsUsingTag return list opf songs which have tag in their title
func GetLyricsUsingTag(client *mongo.Client, tag string) []models.Song {
	var songs []models.Song

	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	collection := client.Database("CAP-THAT-PIC").Collection("Captions")

	filter := bson.D{{"lyrics", primitive.Regex{Pattern: tag, Options: ""}}}
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		log.Println("Find ERROR:", err)
		defer cursor.Close(ctx)
	} else {
		log.Println("Find() API result:", cursor)
		for cursor.Next(ctx) {
			var result models.Song
			err = cursor.Decode(&result)

			if err != nil {
				log.Println("cursor.Next() error: ", err)
			} else {
				songs = append(songs, result)
			}
		}
	}

	return songs
}

// SetupDB adds lyrics to DB
func SetupDB() {
	log.Println("Add lyrics to DB...")
	client := ConnectToDB()
	AddLyricsToDB(client)
	log.Println("Added lyrics to DB successfully.")
	AddPoemsToDB(client)
	log.Println("Added poems to DB successfully.")
	AddMovieQuotesToDB(client)
	log.Println("Added movie quotes to DB successfully.")
}

func AddCaptionToDB(client *mongo.Client, caption *models.Caption) error {
	collection := client.Database("CAP-THAT-PIC").Collection("Captions")
	_, err := collection.InsertOne(context.TODO(), *caption)
	return err
}

func AddPostToDB(client *mongo.Client, post *models.Post) error {
	collection := client.Database("CAP-THAT-PIC").Collection("Post")
	_, err := collection.InsertOne(context.TODO(), *post)
	return err
}

func GetCaptionByID(client *mongo.Client, id string) (*models.Caption, error) {
	var result models.Caption
	collection := client.Database("CAP-THAT-PIC").Collection("Captions")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Printf("Found a single document: %+v\n", result)
	return &result, err
}

func GetPostByID(client *mongo.Client, id string) (*models.Post, error) {
	var result models.Post
	collection := client.Database("CAP-THAT-PIC").Collection("Post")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	return &result, err
}
