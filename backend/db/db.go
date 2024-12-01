package db

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hao123/util"

	"time"
)

// MongoDB defines the database structure
type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
	Config   util.Config
}

func retryConnect(uri, dbUsername, dbPassword, source string, maxRetries int) (*mongo.Client, error) {
	var client *mongo.Client
	var err error
	for i := 0; i < maxRetries; i++ {
		clientOptions := options.Client().ApplyURI(uri).SetMaxPoolSize(100)
		clientOptions.SetAuth(options.Credential{
			Username:   dbUsername,
			Password:   dbPassword,
			AuthSource: source,
		})
		client, err = mongo.Connect(context.Background(), clientOptions)
		if err == nil && client.Ping(context.Background(), nil) == nil {
			return client, nil
		}

		log.Info().
			Err(err).
			Int("retry", i+1).
			Int("maxRetries", maxRetries).
			Msgf("Retrying MongoDB connection (%d/%d)...", i+1, maxRetries)
		time.Sleep(2 * time.Second)
	}

	return nil, err
}

func Connect(uri, dbName, dbUsername, dbPassword, source string, maxRetries int) (*MongoDB, error) {

	client, err := retryConnect(uri, dbUsername, dbPassword, source, maxRetries)
	if err != nil {
		return nil, err
	}

	database := client.Database(dbName)
	fmt.Println("Successfully connect to mongo database name:", database.Name())

	return &MongoDB{
		Client:   client,
		Database: database,
	}, nil
}

// Disconnect from DB
func (mongo *MongoDB) Disconnect(db *MongoDB) error {
	return mongo.Client.Disconnect(context.Background())
}
