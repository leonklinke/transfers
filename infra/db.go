package infra

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBCollection interface {
	Insert(ctx context.Context, name string) DBCollection
}

type DBClient interface {
	Collection(ctx context.Context, name string) DBCollection
}

func Connect() (DBClient, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/transfers")
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB")
	return client, nil
}

func (db *DBClient) Collection(ctx context.Context, name string) DBCollection {

}
