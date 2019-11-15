package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Caption struct {
	ID   primitive.ObjectID `bson:"_id, omitempty"`
	Text string
	Src  string
	Mood string
	Tags []string
}
