package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoHandler contains the database handler and client handler
type MongoHandler struct {
	DB     *mongo.Database
	Client *mongo.Client
}

// NewMongoHandler creates a new handle for the database and client pool
func NewMongoHandler() *MongoHandler {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	return &MongoHandler{
		DB:     client.Database(os.Getenv("MONGODB_DATABASE")),
		Client: client,
	}
}

// Disconnect closes sockets to the topology referenced by this Client.
func (m MongoHandler) Disconnect(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}

// Store inserts a single document to the collection
func (m MongoHandler) Store(ctx context.Context, collection string, data interface{}) error {
	if _, err := m.DB.Collection(collection).InsertOne(ctx, data); err != nil {
		return err
	}
	return nil
}

// FindAll decodes the cursor results with the matching filter
func (m MongoHandler) FindAll(ctx context.Context, collection string, filter interface{}, result interface{}) error {
	cursor, err := m.DB.Collection(collection).Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, result); err != nil {
		return err
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return nil
}

// FindOne searches a single document with the matching filter
// then unmarshals the data into the result argument
func (m MongoHandler) FindOne(ctx context.Context, collection string, filter interface{}, result interface{}) error {
	if err := m.DB.Collection(collection).FindOne(ctx, filter).Decode(result); err != nil {
		return err
	}
	return nil
}

// Delete removes a document with the matching filter
func (m MongoHandler) Delete(ctx context.Context, collection string, filter interface{}) (int64, error) {
	count, err := m.DB.Collection(collection).DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count.DeletedCount, nil
}
