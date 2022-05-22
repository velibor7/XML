package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
	Id          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Skills      []string           `bson:"skills"`
}
