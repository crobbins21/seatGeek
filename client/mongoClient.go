package mongo

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGOUSER = "mongoUser"
	MONGOPASS = "mongoPass"
)

func MongoClient() *mongo.Client {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	mongoUser := os.Getenv(MONGOUSER)
	mongoPass := os.Getenv(MONGOPASS)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://" + mongoUser + ":" + mongoPass + "@cluster0.6ngbw.mongodb.net/myFirstDatabase?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
