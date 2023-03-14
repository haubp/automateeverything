package commandStorage

import "go.mongodb.org/mongo-driver/mongo"

type mongoStore struct {
	db *mongo.Client
}

func NewMongoStore(db *mongo.Client) *mongoStore {
	return &mongoStore{db}
}
