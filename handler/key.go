package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonklinke/transfers/usecase"
)

func (di *Dependencies) CreateKey(ctx *gin.Context) {
	request := &usecase.CreateKeyReq{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := usecase.CreateKey(ctx, di.DB, request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

type ListKeyResp struct {
	UserID string `json:"user_id"`
	Key    string `json:"key"`
}

func (di *Dependencies) ListKeys(ctx *gin.Context) {
	keys, err := usecase.ListKey(ctx, di.DB, usecase.ListKeyReq{
		UserID: ctx.Query("user_id"),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	resp := make([]ListKeyResp, 0)

	for _, key := range keys {
		resp = append(resp, ListKeyResp{
			UserID: key.UserID,
			Key:    key.Key,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"keys": resp})
}
