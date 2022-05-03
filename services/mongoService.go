package services

import (
	"context"
	"log"
	"reflect"
	"time"

	mongoDB "github.com/crobbins21/seatGeek.git/client/mongo.go"
	"github.com/crobbins21/seatGeek.git/models"
)

const (
	dateFormat = " 01-02-2006"
)

func UpdateDatabase() {

	log.Print("Starting to update database...")

	// Get events from db
	oldEvents := GetEvents(context.TODO())
	oldEventsMap := arrayToMap(oldEvents)
	oldEventsP := *oldEventsMap

	newTopEvents := SeatGeekEvents(500)

	var eventsToAdd []models.Event
	for _, event := range newTopEvents {

		//New event
		if _, val := oldEventsP[event.ID]; !val {
			eventsToAdd = append(eventsToAdd, event)
			log.Print("New event: ", event.Title)

			// Event already in the db
		} else {
			oldEventsP[event.ID] = models.Event{ID: 0}
		}
	}

	var eventsToRemove []int
	for _, oldEvent := range oldEventsP {
		if oldEvent.ID != 0 {
			eventsToRemove = append(eventsToRemove, oldEvent.ID)
			log.Print("Event to remove: ", oldEvent.Title)
		}
	}

	if len(eventsToRemove) == 0 && len(eventsToAdd) == 0 {
		now := time.Now().Format(time.Kitchen + dateFormat)
		sixHoursAgo := time.Now().Add(time.Duration(-6) * time.Hour).Format(time.Kitchen + dateFormat)
		log.Printf("No new top events added from %v till %v", sixHoursAgo, now)
	} else {

		mongoDB.DeleteEvents("eventId", eventsToRemove)
		mongoDB.PostEvents(ToInterface(eventsToAdd))
	}

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
