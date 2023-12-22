package database

import (
	"context"
	"log"
)

type DBDriver interface {
	NewClient() (DBClient, error)
}
type DBClient interface {
	Insert(ctx context.Context, tableName string, model interface{}) (interface{}, error)
}

func NewDB(driver DBDriver) DBClient {
	client, err := driver.NewClient()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return client
}
