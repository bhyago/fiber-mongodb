package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbname = "gotodo"
)

func getConnection() (client *mongo.Client, ctx context.Context, cancel context.CancelFunc, err error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx, cancel, err

}
