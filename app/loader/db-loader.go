package loader

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoDB *mongo.Database

func DBLoader() {
	dbName := os.Getenv("DB_NAME")
	dbURL := os.Getenv("DB_URL")
	dbOptions := options.Client().ApplyURI(dbURL)
	mongoClient, err := mongo.Connect(context.TODO(), dbOptions)
	onFailingDB(err)
	err = mongoClient.Ping(context.TODO(), readpref.Primary())
	onFailingDB(err)
	MongoDB = mongoClient.Database(dbName)
	fmt.Println("DB loaded successfully ...")
}

func onFailingDB(err error) {
	if err != nil {
		panic("The database has failed to connect ...")
	}
}
