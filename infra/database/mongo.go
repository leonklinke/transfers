package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct{}

type MongoClient struct {
	Client *mongo.Client
}

func (m *Mongo) NewClient() (DBClient, error) {
	fmt.Println("Connecting to MongoDB...")

	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017/transfers")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("failed to connect MongoDB: ", err)
		return nil, err
	}

	fmt.Println("Successfully connected to MongoDB")

	return &MongoClient{Client: client}, nil
}

func (mongo *MongoClient) Insert(ctx context.Context, tableName string, model interface{}) (interface{}, error) {
	collection := mongo.Client.Database("transfers").Collection(tableName)
	return collection.InsertOne(ctx, model)
}
