package usecase

import "github.com/leonklinke/transfers/domain"

type CreateKeyReq struct {
	UserID string
	Key    string
}

type ListKeyReq struct {
	UserID string
	Key    string
}

func CreateKey(req CreateKeyReq) (*domain.Key, error) {
	return nil, nil
}

func ListKey(req ListKeyReq) ([]*domain.Key, error) {
	return nil, nil
}
