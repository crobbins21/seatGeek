package services

import (
	"context"
	"log"
	"reflect"

	mongoDB "github.com/crobbins21/seatGeek.git/client/mongo.go"
	"github.com/crobbins21/seatGeek.git/models"
)

func UpdateDatabase() {

	log.Print("Starting to update database...")

	// Get events from db
	oldEvents := GetEvents(context.TODO())
	oldEventsMap := arrayToMap(oldEvents)
	oldEventsP := *oldEventsMap

	// Get events from SG
	newTopEvents := SeatGeekEvents(10)

	var eventsToAdd []models.Event
	for _, event := range newTopEvents {

		//New event
		if _, val := oldEventsP[event.ID]; !val {
			eventsToAdd = append(eventsToAdd, event)
			log.Print("New event: ", event)

			// Event already in the db
		} else {
			oldEventsP[event.ID] = models.Event{ID: 0}
			log.Print("Event already in the db: ", event)
		}
	}

	var eventsToRemove []int
	for _, oldEvent := range oldEventsP {
		if oldEvent.ID != 0 {
			eventsToRemove = append(eventsToRemove, oldEvent.ID)
			log.Print("Event to remove: ", oldEvent)
		}
	}

	//Delete all old
	mongoDB.DeleteEvents("eventId", eventsToRemove)

	//Add all new
	mongoDB.PostEvents(ToInterface(eventsToAdd))

}

func arrayToMap(events []models.Event) *map[int]models.Event {
	eventsMap := make(map[int]models.Event)
	for _, event := range events {
		eventsMap[event.ID] = event
	}
	return &eventsMap
}

func GetEvents(ctx context.Context) []models.Event {

	events := mongoDB.GetEvents()
	if events != nil {

		var returnResp []models.Event
		events.All(ctx, &returnResp)
		events.Close(ctx)
		return returnResp

	} else {
		return nil
	}
}

func PostEvents(events []models.Event) {

	if len(events) != 0 {
		mongoDB.PostEvents(ToInterface(events))
	}

}

func ToInterface(events []models.Event) []interface{} {

	var retArray []interface{}
	for _, event := range events {
		retArray = append(retArray, reflect.ValueOf(event).Interface())
	}

	return retArray
}
