package persistence

import (
	"context"
	"fmt"

	"github.com/velibor7/XML/job_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "job_db"
	COLLECTION = "job"
)

type JobMongoDB struct {
	jobs *mongo.Collection
}

func NewJobMongoDB(client *mongo.Client) *JobMongoDB {
	job := client.Database(DATABASE).Collection(COLLECTION)
	return &JobMongoDB{
		jobs: job,
	}
}

func (db *JobMongoDB) Get(id string) (*domain.Job, error) {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return db.filterOne(bson.M{"_id": Id})

}

func (db *JobMongoDB) GetAll(search string) ([]*domain.Job, error) {
	filter := bson.D{{"title", bson.M{"$regex": "^.*" + search + ".*$"}}, {"description", bson.M{"$regex": "^.*" + search + ".*$"}}}
	return db.filter(filter)
}

func (db *JobMongoDB) Create(job *domain.Job) error {
	_, err := db.jobs.InsertOne(context.TODO(), job)
	if err != nil {
		return err
	}
	return nil
}

func (db *JobMongoDB) DeleteAll() error {
	fmt.Print("pre deletemany")
	_, err := db.jobs.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (db *JobMongoDB) filter(filter interface{}) ([]*domain.Job, error) {
	cursor, err := db.jobs.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (db *JobMongoDB) filterOne(filter interface{}) (Job *domain.Job, err error) {
	result := db.jobs.FindOne(context.TODO(), filter)
	err = result.Decode(&Job)
	return
}

func decode(cursor *mongo.Cursor) (Jobs []*domain.Job, err error) {
	for cursor.Next(context.TODO()) {
		var Job domain.Job
		err = cursor.Decode(&Job)
		if err != nil {
			return
		}
		Jobs = append(Jobs, &Job)
	}
	err = cursor.Err()
	return
}
