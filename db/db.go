package db

import (
	"context"
	"github.com/o-z/featuretoggle/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Database struct {
	MongoDB *mongo.Database
}

func Get(config *config.Config) (*Database, error) {
	config.DBHost = "localhost"
	config.DBPort = "27017"
	config.DBName = "deneme"
	clientOptions := options.Client().
		ApplyURI("mongodb://" + config.DBHost + ":" + config.DBPort)
	/*SetAuth(options.Credential{
		AuthSource: config.DBName,
		Username:   config.DBUser,
		Password:   config.DBPass,
	})*/

	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		MongoDB: client.Database(config.DBName),
	}, err
}
