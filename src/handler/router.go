package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Serve() {
	fmt.Println("Listening on http://localhost:8080")

	r := gin.Default()

	r.GET("/keys", ListKeys)
	r.POST("/keys", CreateKey)

	r.Run("localhost:8080")
}
