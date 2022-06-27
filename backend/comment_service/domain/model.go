package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	Id       primitive.ObjectID `bson:"_id"`
	PostId   primitive.ObjectID `bson:"post_id"`
	Username string             `bson:"username"`
	Content  string             `bson:"content"`
}
