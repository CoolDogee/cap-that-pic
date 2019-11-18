package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Caption struct {
	ID   primitive.ObjectID `bson:"_id, omitempty"`
	Text []string // contents of caption
	Src  string // url of the content
	Mood []string // mood extracted from the "text"
	Tags []string // tags for the "text" like proper or common nouns in the text, author, poet, artist, etc.
	Type string // song, movie, poem, book, etc.
	UserGenerated bool // false for original cption objects, true when generated due to post request of user
}
