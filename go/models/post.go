package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	ImgURL    string
	CaptionID string
	Filter    string
	Tags      []string
}

type PostWithCaption struct {
	ID 			primitive.ObjectID `bson:"_id, omitempty"`
	ImgURL		string
	CaptionID	string
	Filter		string
	Tags		[]string
	Caption		Caption
}