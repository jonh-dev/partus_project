package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBService struct {
	Client *mongo.Client
	DBName string
}

func NewDBService(envGetter *EnvVarGetter) (*DBService, error) {

	inContainer, err := envGetter.Get("IN_CONTAINER")
	if err != nil {
		inContainer = "false"
	}

	var uri string
	if inContainer == "true" {
		uri, err = envGetter.Get("DOCKER_MONGO_URI")
		if err != nil {
			return nil, err
		}
	} else {
		uri, err = envGetter.Get("MONGO_URI")
		if err != nil {
			return nil, err
		}
	}

	dbName, err := envGetter.Get("DB_NAME")
	if err != nil {
		return nil, err
	}

	clientOptions := options.Client().ApplyURI(uri)

	client, err := connectToMongoDB(uri, clientOptions)
	if err != nil {
		return nil, err
	}

	return &DBService{Client: client, DBName: dbName}, nil
}

func connectToMongoDB(uri string, clientOptions *options.ClientOptions) (*mongo.Client, error) {
	var client *mongo.Client
	var err error
	for i := 0; i < 5; i++ {
		client, err = mongo.Connect(context.Background(), clientOptions)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to MongoDB, retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
	return client, err
}
