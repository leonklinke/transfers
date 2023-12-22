package main

import (
	"github.com/leonklinke/transfers/handler"
	"github.com/leonklinke/transfers/infra/database"
)

func main() {
	mongo := &database.Mongo{}
	client := database.NewDB(mongo)
	handler.Serve(client)
}
