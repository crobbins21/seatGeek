package services

import (
	"encoding/json"

	seatGeek "github.com/crobbins21/seatGeek.git/client/seatGeek.go"
	"github.com/crobbins21/seatGeek.git/models"
)

func SeatGeekEvents(numOfEvents int) []models.Event {

	eventsResp := seatGeek.GetTopEvents(numOfEvents)
	defer eventsResp.Body.Close()

	var eventsRrturn models.EventsReturn
	json.NewDecoder(eventsResp.Body).Decode(&eventsRrturn)

	return eventsRrturn.Events
}
