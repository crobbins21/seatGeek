package main

import (
	"context"
	"encoding/json"
	"log"

	mongoDB "github.com/crobbins21/seatGeek.git/client/mongo"
	seatGeek "github.com/crobbins21/seatGeek.git/client/seatGeek"

	"github.com/crobbins21/seatGeek.git/models"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	seatGeekAPI   = "https://api.seatgeek.com/2"
	eventsEnp     = "/events"
	performersEnp = "/performers"
	venuesEnp     = "/venues"
	fwdSlash      = "/"
	GEEKUSER      = "SeatGeekUser"
	GEEKPASS      = "SeatGeekPass"
)

func main() {
	log.Print("Getting top 500 tickets...")
	getTickets()
}

func getTickets() {

	ticketsResp := seatGeek.GetEvents(500)

	var ticketResp models.TicketsResp
	json.NewDecoder(resp.Body).Decode(&ticketResp)

	for _, event := range ticketResp.Events {
		log.Print(event, "\n")
	}

	dbConnection := mongoDB.MongoClient()
	log.Print(dbConnection)

	if err := dbConnection.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	} else {
		log.Print("Successfully pinged")
	}
}
