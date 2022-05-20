package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Connection struct {
	Id        primitive.ObjectID `bson:"_id"`
	IssuerId  primitive.ObjectID `bson:"issuerId"`
	SubjectId primitive.ObjectID `bson:"subjectId"`
	Date      time.Time          `bson:"time"`
	Approved  bool               `bson:"approved"`
}

type ProfilePrivacy struct {
	Id      primitive.ObjectID `bson:"_id"`
	UserId  primitive.ObjectID `bson:"userId"`
	Private bool               `bson:"private"`
}
