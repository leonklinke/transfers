package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonklinke/transfers/usecase"
)

func CreateKey(c *gin.Context) {
	key, err := usecase.CreateKey(usecase.CreateKeyReq{
		UserID: c.PostForm("user_id"),
		Key:    c.PostForm("key"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"key": key.Key})
}

type ListKeyResp struct {
	UserID string `json:"user_id"`
	Key    string `json:"key"`
}

func ListKeys(c *gin.Context) {
	keys, err := usecase.ListKey(usecase.ListKeyReq{
		UserID: c.Query("user_id"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	resp := make([]ListKeyResp, 0)

	for _, key := range keys {
		resp = append(resp, ListKeyResp{
			UserID: key.UserID,
			Key:    key.Key,
		})
	}

	c.JSON(http.StatusOK, gin.H{"keys": resp})
}
