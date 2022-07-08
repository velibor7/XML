package persistence

import (
	"context"

	"github.com/velibor7/XML/authentication_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	DATABASE   = "authentication_db"
	COLLECTION = "credentials"
)

type AuthMongoDBStore struct {
	userCredentials *mongo.Collection
}

func NewAuthMongoDBStore(client *mongo.Client) domain.AuthStore {
	userCredentials := client.Database(DATABASE).Collection(COLLECTION)
	return &AuthMongoDBStore{
		userCredentials: userCredentials,
	}
}

func (store *AuthMongoDBStore) Login(userCredential *domain.UserCredential) (*domain.UserCredential, error) {
	filter := bson.M{"username": userCredential.Username, "password": userCredential.Password}
	authUser, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return authUser, nil
}

func (store *AuthMongoDBStore) Register(userCredential *domain.UserCredential) (*domain.UserCredential, error) {
	filter := bson.M{"username": userCredential.Username}
	userCredenitalExsist, err := store.filterOne(filter)

	if userCredenitalExsist != nil {
		return nil, status.Error(codes.AlreadyExists, "User already exists with same username")
	}
	userCredential.Id = primitive.NewObjectID()

	result, err := store.userCredentials.InsertOne(context.TODO(), userCredential)
	if err != nil {
		return nil, err
	}
	userCredential.Id = result.InsertedID.(primitive.ObjectID)
	return userCredential, nil
}

func (store *AuthMongoDBStore) GetByUsername(username string) (*domain.UserCredential, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *AuthMongoDBStore) Update(id primitive.ObjectID, username string) (string, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"username": username}}
	_, err := store.userCredentials.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "", err
	}
	return username, nil
}

func (store *AuthMongoDBStore) DeleteAll() {
	store.userCredentials.DeleteMany(context.TODO(), bson.D{})
}

func (store *AuthMongoDBStore) Delete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := store.userCredentials.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *AuthMongoDBStore) filterOne(filter interface{}) (userCredential *domain.UserCredential, err error) {
	result := store.userCredentials.FindOne(context.TODO(), filter)
	err = result.Decode(&userCredential)
	if err != nil {
		return nil, nil
	}
	return
}

func (store *AuthMongoDBStore) checkIfExsist(filter interface{}) (exsist bool, err error) {
	_, err = store.userCredentials.Find(context.TODO(), filter)
	if err != nil {
		return false, nil
	}
	return true, nil
}
