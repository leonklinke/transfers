package domain

type Key struct {
	UserID string `bson:"user_id"`
	Key    string `bson:"key"`
}
