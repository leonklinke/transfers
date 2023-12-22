package usecase

import (
	"context"

	"github.com/leonklinke/transfers/domain"
	"github.com/leonklinke/transfers/infra/database"
)

type CreateKeyReq struct {
	UserID string `json:"user_id"`
	Key    string `json:"key"`
}

type ListKeyReq struct {
	UserID string `json:"user_id"`
	Key    string `json:"key"`
}

func CreateKey(ctx context.Context, db database.DBClient, req *CreateKeyReq) error {
	insertKey := domain.Key{
		UserID: req.UserID,
		Key:    req.Key,
	}

	if _, err := db.Insert(ctx, "keys", insertKey); err != nil {
		return err
	}

	return nil
}

func ListKey(ctx context.Context, db database.DBClient, req ListKeyReq) ([]*domain.Key, error) {
	return nil, nil
}
