package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostEvents(events []interface{}) {
	ctx := context.Background()
	dbConnection := MongoClient()

	dbConnection.Connect(ctx)
	connection := dbConnection.Database("seatGeek").Collection("events")

	_, err := connection.InsertMany(ctx, events)
	if err != nil {
		log.Print("Error posting events to the db")
	} else {
		log.Print("Successfully posted events to the db")
	}
}

func DeleteEvents(identifier string, events []int) {
	ctx := context.Background()
	dbConnection := MongoClient()

	dbConnection.Connect(ctx)
	connection := dbConnection.Database("seatGeek").Collection("events")

	query := bson.M{identifier: bson.M{"$nin": events}}

	_, err := connection.DeleteMany(ctx, query)
	if err != nil {
		log.Print("Error deleting events in the db")
	} else {
		log.Print("Successfully deleted events in the db")
	}
}

func GetEvents() *mongo.Cursor {
	ctx := context.Background()
	dbConnection := MongoClient()

	dbConnection.Connect(ctx)
	connection := dbConnection.Database("seatGeek").Collection("events")

	events, err := connection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Print("Error retreiving events from the db")
		return nil
	} else {
		return events
	}
}
