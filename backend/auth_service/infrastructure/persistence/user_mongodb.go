package persistence

import (
	"context"

	"github.com/velibor7/XML/auth_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "auth_service"
	COLLECTION = "user"
)

type UserMongoDB struct {
	users *mongo.Collection
}

func NewUserMongoDBInterface(client *mongo.Client) domain.UserInterface {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDB{
		users: users,
	}
}

func (store *UserMongoDB) Get(username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *UserMongoDB) GetAll() ([]*domain.User, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserMongoDB) Register(User *domain.User) error {
	result, err := store.users.InsertOne(context.TODO(), User)
	if err != nil {
		return err
	}
	User.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserMongoDB) DeleteAll() error {
	_, err := store.users.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDB) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *UserMongoDB) filterOne(filter interface{}) (User *domain.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&User)
	return
}

func decode(cursor *mongo.Cursor) (users []*domain.User, err error) {
	for cursor.Next(context.TODO()) {
		var User domain.User
		err = cursor.Decode(&User)
		if err != nil {
			return
		}
		users = append(users, &User)
	}
	err = cursor.Err()
	return
}
