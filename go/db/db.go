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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToDB makes connection with database
func ConnectToDB() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://an_chou:59cCdK$Q5MKSMMu@cluster-lrx2r.mongodb.net/test?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

// CloseConnectionDB ends the connection with database
func CloseConnectionDB(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
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