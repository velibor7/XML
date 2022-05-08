package persistence

import (
	"context"
	"errors"

	"github.com/velibor7/XML/post_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE   = "post_db"
	COLLECTION = "post"
)

type PostMongoDB struct {
	posts *mongo.Collection
}

func NewPostMongoDB(client *mongo.Client) domain.PostInterface {
	posts := client.Database(DATABASE).Collection(COLLECTION)
	return &PostMongoDB{
		posts: posts,
	}
}

func (store *PostMongoDB) Get(id primitive.ObjectID) (*domain.Post, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *PostMongoDB) GetAll() ([]*domain.Post, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *PostMongoDB) Create(post *domain.Post) (string, error) {
	post.Id = primitive.NewObjectID()
	result, err := store.posts.InsertOne(context.TODO(), post)
	if err != nil {
		return "error", err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)
	return "success", nil
}

func (store *PostMongoDB) Update(post *domain.Post) (string, error) {
	newData := bson.M{"$set": bson.M{
		"text":     post.Text,
		"created":  post.Created,
		"images":   post.Images,
		"links":    post.Links,
		"likes":    post.Likes,
		"dislikes": post.Dislikes,
		"comments": post.Comments,
		"user_id":  post.UserId,
	}}

	opts := options.Update().SetUpsert(true)
	result, err := store.posts.UpdateOne(context.TODO(), bson.M{"_id": post.Id}, newData, opts)

	if err != nil {
		return "error", err
	}
	if result.MatchedCount != 1 {
		return "one document should've been updated", errors.New("one document should've been updated")
	}
	return "success", nil
}

// todo ovo ce trebati da se zakomentarise
// func (store *PostMongoDB) Delete(id primitive.ObjectID) {
// 	store.posts.DeleteOne(context.TODO(), bson.M{"_id": id}, nil)
// }

func (store *PostMongoDB) DeleteAll() {
	store.posts.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *PostMongoDB) GetAllByUser(id string) ([]*domain.Post, error) {
	filter := bson.M{"user_id": id}
	// return store.posts.filter(filter)
	return store.filter(filter)
}

func (store *PostMongoDB) filter(filter interface{}) ([]*domain.Post, error) {
	cursor, err := store.posts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *PostMongoDB) filterOne(filter interface{}) (post *domain.Post, err error) {
	result := store.posts.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func decode(cursor *mongo.Cursor) (posts []*domain.Post, err error) {
	for cursor.Next(context.TODO()) {
		var post domain.Post
		err = cursor.Decode(&post)
		if err != nil {
			return
		}
		posts = append(posts, &post)
	}
	err = cursor.Err()
	return
}
