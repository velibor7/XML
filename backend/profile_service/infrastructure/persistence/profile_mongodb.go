package persistence

import (
	"context"
	"errors"
	"time"

	"github.com/velibor7/XML/profile_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE   = "profile_service"
	COLLECTION = "profile"
)

type ProfileMongoDBStore struct {
	profiles *mongo.Collection
}

func NewProfileMongoDB(client *mongo.Client) domain.ProfileInterface {
	profiles := client.Database(DATABASE).Collection(COLLECTION)
	index := mongo.IndexModel{
		Keys:    bson.D{{"fullName", "text"}},
		Options: options.Index().SetUnique(true),
	}
	opts := options.CreateIndexes().SetMaxTime(20 * time.Second)
	_, err := profiles.Indexes().CreateOne(context.TODO(), index, opts)

	if err != nil {
		panic(err)
	}

	return &ProfileMongoDBStore{
		profiles: profiles,
	}
}

func (store *ProfileMongoDBStore) Get(username string) (*domain.Profile, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *ProfileMongoDBStore) GetAll(search string) ([]*domain.Profile, error) {
	filter := bson.D{{"fullName", bson.M{"$regex": "^.*" + search + ".*$"}}}
	return store.filter(filter, search)
}

func (store *ProfileMongoDBStore) Create(profile *domain.Profile) error {
	result, err := store.profiles.InsertOne(context.TODO(), profile)
	if err != nil {
		return err
	}
	profile.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

func (store *ProfileMongoDBStore) Update(username string, profile *domain.Profile) error {
	result, err := store.profiles.ReplaceOne(
		context.TODO(),
		bson.M{"username": username},
		profile,
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New(profile.Id.String())
	}
	return nil
}

func (store *ProfileMongoDBStore) DeleteAll() error {
	_, err := store.profiles.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *ProfileMongoDBStore) filter(filter interface{}, search string) ([]*domain.Profile, error) {
	cursor, err := store.profiles.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, errors.New(search)
	}
	return decode(cursor)
}

func (store *ProfileMongoDBStore) filterOne(filter interface{}) (profile *domain.Profile, err error) {
	result := store.profiles.FindOne(context.TODO(), filter)
	err = result.Decode(&profile)
	return
}

func decode(cursor *mongo.Cursor) (profiles []*domain.Profile, err error) {
	for cursor.Next(context.TODO()) {
		var Profile domain.Profile
		err = cursor.Decode(&Profile)
		if err != nil {
			return
		}
		profiles = append(profiles, &Profile)
	}
	err = cursor.Err()
	return
}