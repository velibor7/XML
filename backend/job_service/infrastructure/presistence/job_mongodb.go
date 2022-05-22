package persistence

import (
	"context"

	"github.com/velibor7/XML/job_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE    = "jobs_db"
	CONNECTIONS = "job"
)

type JobMongoDB struct {
	job *mongo.Collection
}

func NewJobMongoDB(job *mongo.Collection) *JobMongoDB {
	return &JobMongoDB{
		job: job,
	}
}

func (db JobMongoDB) Get(id string) (*domain.Job, error) {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return db.filterOne(bson.M{"_id": Id})

}

func (db JobMongoDB) GetAll(search string) ([]*domain.Job, error) {
	filter := bson.D{{"title", bson.M{"$regex": "^.*" + search + ".*$"}}, {"description", bson.M{"$regex": "^.*" + search + ".*$"}}}
	return db.filter(filter)
}

func (db JobMongoDB) Create(job *domain.Job) (*domain.Job, error) {
	insert, err := db.job.InsertOne(context.TODO(), job)
	if err != nil {
		return nil, err
	}
	job.Id = insert.InsertedID.(primitive.ObjectID)
	return job, nil
}

func (db JobMongoDB) DeleteAll() error {
	_, err := db.job.DeleteMany(context.TODO(), bson.D{})
	return err
}

func (db *JobMongoDB) filter(filter interface{}) ([]*domain.Job, error) {
	cursor, err := db.job.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (db *JobMongoDB) filterOne(filter interface{}) (Job *domain.Job, err error) {
	result := db.job.FindOne(context.TODO(), filter)
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
