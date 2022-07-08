package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	Id     primitive.ObjectID `bson:"_id"`
	UserId string             `bson:"user_id"`
}

type Dislike struct {
	Id     primitive.ObjectID `bson:"_id"`
	UserId string             `bson:"user_id"`
}

type Post struct {
	Id       primitive.ObjectID `bson:"_id"`
	Text     string             `bson:"text"`
	Images   string             `bson:"images"`
	Links    string             `bson:"links"`
	Created  time.Time          `bson:"created"`
	Likes    []Like             `bson:"likes"`
	Dislikes []Dislike          `bson:"dislikes"`
	UserId   primitive.ObjectID `bson:"user_id"`
}
