package persistence

import (
	"context"
	"errors"

	"github.com/velibor7/XML/profile_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "profile_db"
	COLLECTION = "profile"
)

type ProfileMongoDB struct {
	profiles *mongo.Collection
}

func NewProfileMongoDB(client *mongo.Client) domain.ProfileInterface {
	profiles := client.Database(DATABASE).Collection(COLLECTION)
	return &ProfileMongoDB{
		profiles: profiles,
	}
}

func (store *ProfileMongoDB) Get(profileId string) (*domain.Profile, error) {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}
func (store *ProfileMongoDB) GetAll() ([]*domain.Profile, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *ProfileMongoDB) Create(profile *domain.Profile) error {
	result, err := store.profiles.InsertOne(context.TODO(), profile)
	if err != nil {
		return err
	}
	profile.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *ProfileMongoDB) Update(profileId primitive.ObjectID, profile *domain.Profile) error {

	result, err := store.profiles.ReplaceOne(
		context.TODO(),
		bson.M{"_id": profileId},
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

func (store *ProfileMongoDB) DeleteAll() error {
	_, err := store.profiles.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *ProfileMongoDB) filter(filter interface{}) ([]*domain.Profile, error) {
	cursor, err := store.profiles.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *ProfileMongoDB) filterOne(filter interface{}) (profile *domain.Profile, err error) {
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
