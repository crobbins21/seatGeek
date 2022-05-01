package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	mongoDB "github.com/crobbins21/seatGeek.git/client"
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

	client := &http.Client{Timeout: time.Duration(1) * time.Second}

	req, _ := http.NewRequest("GET", seatGeekAPI+eventsEnp+"?per_page=100", nil)
	req.SetBasicAuth(os.Getenv(GEEKUSER), os.Getenv(GEEKPASS))
	resp, doErr := client.Do(req)
	if doErr != nil {
		log.Print("Error to seatGeek: ", doErr)
	}

	defer resp.Body.Close()

	var ticketResp models.TicketsResp
	json.NewDecoder(resp.Body).Decode(&ticketResp)

	// for _, event := range ticketResp.Events {
	// 	log.Print(event, "\n")
	// }

	dbConnection := mongoDB.MongoClient()
	log.Print(dbConnection)

	if err := dbConnection.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	} else {
		log.Print("Successfully pinged")
	}
}
