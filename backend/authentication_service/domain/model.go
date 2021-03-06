package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserCredential struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

type JWTToken struct {
	Token string `json:"token"`
}
type JWTTokenFullResponse struct {
	Token    string             `json:"token"`
	UserId   primitive.ObjectID `bson:"userId"`
	Username string             `bson:"username"`
}
