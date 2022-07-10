package persistence

import (
	"context"

	"github.com/velibor7/XML/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE    = "connection_db"
	CONNECTIONS = "connection"
	PRIVACY     = "profilePrivacy"
)

type ConnectionMongoDB struct {
	connections     *mongo.Collection
	profilesPrivacy *mongo.Collection
}

func NewConnectionMongoDB(client *mongo.Client) domain.ConnectionInterface {
	conn := client.Database(DATABASE).Collection(CONNECTIONS)
	profPriv := client.Database(DATABASE).Collection(PRIVACY)
	return &ConnectionMongoDB{
		connections:     conn,
		profilesPrivacy: profPriv,
	}
}

func (db *ConnectionMongoDB) Get(userId string) ([]*domain.Connection, error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"$or": []bson.M{{"subjectId": id}, {"issuerId": id}}}
	return db.filter(filter)
}

func (db *ConnectionMongoDB) Create(conn *domain.Connection) (*domain.Connection, error) {
	insert, err := db.connections.InsertOne(context.TODO(), conn)
	if err != nil {
		return nil, err
	}
	conn.Id = insert.InsertedID.(primitive.ObjectID)
	return conn, nil
}

func (db *ConnectionMongoDB) CreatePrivacy(pPriv *domain.ProfilePrivacy) (*domain.ProfilePrivacy, error) {
	insert, err := db.profilesPrivacy.InsertOne(context.TODO(), pPriv)
	if err != nil {
		return nil, err
	}
	pPriv.Id = insert.InsertedID.(primitive.ObjectID)

	return pPriv, nil
}

func (db *ConnectionMongoDB) DeleteAll() error {
	_, err := db.connections.DeleteMany(context.TODO(), bson.D{})
	_, err = db.profilesPrivacy.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	return nil
}

func (db *ConnectionMongoDB) Delete(id string) error {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = db.connections.DeleteOne(context.TODO(), bson.M{"_id": Id})
	if err != nil {
		return err
	}
	return nil
}

func (db *ConnectionMongoDB) Update(id string) (*domain.Connection, error) {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	conn, err := db.filterOne(bson.M{"_id": Id})
	if err != nil {
		return nil, err
	}
	conn.Approved = !conn.Approved
	_, err = db.connections.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{{"$set", bson.M{"approved": conn.Approved}}})
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (db *ConnectionMongoDB) filter(filter interface{}) ([]*domain.Connection, error) {
	cursor, err := db.connections.Find(context.TODO(), filter)
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
func (db *ConnectionMongoDB) filterOne(filter interface{}) (connection *domain.Connection, err error) {
	result := db.connections.FindOne(context.TODO(), filter)
	err = result.Decode(&connection)
	return
}

func (db *ConnectionMongoDB) filterOnePrivacy(filter interface{}) (privacy *domain.ProfilePrivacy, err error) {
	result := db.profilesPrivacy.FindOne(context.TODO(), filter)
	err = result.Decode(&privacy)
	return
}

func decode(cursor *mongo.Cursor) (connections []*domain.Connection, err error) {
	for cursor.Next(context.TODO()) {
		var Connection domain.Connection
		err = cursor.Decode(&Connection)
		if err != nil {
			return
		}
		connections = append(connections, &Connection)
	}
	err = cursor.Err()
	return
}
