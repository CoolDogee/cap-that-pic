package db

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
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
	// username := os.Getenv("MONGODB_USERNAME")
	// password := os.Getenv("MONGODB_PASSWORD")
	// clientOptions := options.Client().ApplyURI("mongodb+srv://" + username + ":" + password + "@cluster-lrx2r.mongodb.net/test?retryWrites=true&w=majority&authMechanism=SCRAM-SHA-1")
	// Use local DB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println("Fail to connect to MongoDB: ", err)
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println("Fail to connect to MongoDB: ", err)
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

// CloseConnectionDB ends the connection with database
func CloseConnectionDB(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		fmt.Println("Fail to close MongoDB: ", err)
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

// AddLyricsToDB makes a connection with the NoSQL database
func AddLyricsToDB(client *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	collection := client.Database("CAP-THAT-PIC").Collection("Lyrics")
	n, err := collection.DeleteMany(ctx, bson.M{})

	if err != nil {
		fmt.Println("DeleteMany ERROR:", err)
	} else {
		fmt.Println("Number of documents removed: ", n)
	}

	byteValues, err := ioutil.ReadFile("../lyrics/lyrics.json")

	if err != nil {
		// Print any IO errors with the .json file
		fmt.Println("ioutil.ReadFile ERROR:", err)
	}

	var docs []models.Song
	err = json.Unmarshal(byteValues, &docs)

	// Print MongoDB docs object type
	fmt.Println("nMongoFields Docs:", reflect.TypeOf(docs), len(docs))

	for i := range docs {
		doc := docs[i]
		result, err := collection.InsertOne(ctx, doc)

		if err != nil {
			fmt.Println("InsertOne ERROR:", err)
		} else {
			fmt.Println("InsertOne() API result:", result)
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
	collection := client.Database("CAP-THAT-PIC").Collection("Lyrics")

	filter := bson.D{{"lyrics", primitive.Regex{Pattern: tag, Options: ""}}}
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		fmt.Println("Find ERROR:", err)
		defer cursor.Close(ctx)
	} else {
		fmt.Println("Find() API result:", cursor)
		for cursor.Next(ctx) {
			var result models.Song
			err = cursor.Decode(&result)

			if err != nil {
				fmt.Println("cursor.Next() error: ", err)
			} else {
				songs = append(songs, result)
			}
		}
	}

	return songs
}
