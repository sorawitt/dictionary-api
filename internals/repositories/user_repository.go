package repositories

import (
	"context"
	"dictionary-api/internals/core/ports"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const MongoClientTimeout = 5

type userRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewUserRepository(conn string) (ports.UserRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return userRepository{
		client:     client,
		database:   client.Database("eternal"),
		collection: client.Database("eternal").Collection("users"),
	}, nil
}

func (r userRepository) Login(email string, password string) error {
	return nil
}

func (r userRepository) Register(email string, password string) error {
	return nil
}
