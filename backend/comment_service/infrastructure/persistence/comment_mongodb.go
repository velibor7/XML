package persistence

import (
	"context"

	"github.com/velibor7/XML/comment_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "comment_db"
	COLLECTION = "comment"
)

type CommentMongoDB struct {
	comments *mongo.Collection
}

func NewCommentMongoDB(client *mongo.Client) domain.CommentInterface {
	comment := client.Database(DATABASE).Collection(COLLECTION)
	return &CommentMongoDB{
		comments: comment,
	}
}

func (db *CommentMongoDB) Create(comment *domain.Comment) error {
	insert, err := db.comments.InsertOne(context.TODO(), comment)
	if err != nil {
		return err
	}
	comment.Id = insert.InsertedID.(primitive.ObjectID)
	return nil
}

func (db *CommentMongoDB) DeleteAll() error {
	_, err := db.comments.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (db *CommentMongoDB) GetForPost(id string) ([]*domain.Comment, error) {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	response, _ := db.filter(bson.M{"post_id": Id})
	return response, nil
}

func (db *CommentMongoDB) UpdateUsername(oldUsername string, newUsername string) error {
	comments, err := db.filter(bson.M{"username": oldUsername})
	if err != nil {
		return nil
	}
	for _, comment := range comments {
		_, err = db.comments.UpdateOne(context.TODO(), bson.M{"_id": comment.Id}, bson.M{"$set": bson.M{"username": newUsername}})
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *CommentMongoDB) filter(filter interface{}) ([]*domain.Comment, error) {
	cursor, err := db.comments.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (db *CommentMongoDB) filterOne(filter interface{}) (Comment *domain.Comment, err error) {
	result := db.comments.FindOne(context.TODO(), filter)
	err = result.Decode(&Comment)
	return
}

func decode(cursor *mongo.Cursor) (Comments []*domain.Comment, err error) {
	for cursor.Next(context.TODO()) {
		var Comment domain.Comment
		err = cursor.Decode(&Comment)
		if err != nil {
			return
		}
		Comments = append(Comments, &Comment)
	}
	err = cursor.Err()
	return
}
