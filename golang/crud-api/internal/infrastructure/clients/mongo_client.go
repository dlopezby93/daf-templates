package clients

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	mongoClient *mongo.Client
}

var mongoInstance *MongoClient

func NewMongoClient(uri string) *MongoClient {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Failed to connect to MongoDB: %v", err))
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed Mongo connection")
	}

	mongoInstance = &MongoClient{mongoClient: client}

	return mongoInstance
}

func (m *MongoClient) GetClient() *mongo.Client {
	return m.mongoClient
}
