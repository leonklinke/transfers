package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leonklinke/transfers/infra/database"
)

type Dependencies struct {
	DB database.DBClient
}

func Serve(db database.DBClient) {
	di := &Dependencies{
		DB: db,
	}

	r := gin.Default()

	r.GET("/keys", di.ListKeys)
	r.POST("/keys", di.CreateKey)

	fmt.Println("Listening on http://localhost:8080")
	r.Run(":8080")
}
