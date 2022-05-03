package seatGeek

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	seatGeekAPI   = "https://api.seatgeek.com/2"
	eventsEnp     = "/events"
	performersEnp = "/performers"
	venuesEnp     = "/venues"
	fwdSlash      = "/"

	GEEKUSER = "SeatGeekUser"
	GEEKPASS = "SeatGeekPass"
)

func GetTopEvents(numOfEvent int) *http.Response {

	log.Print("Retrieving events...")
	client := &http.Client{Timeout: time.Duration(1) * time.Second}

	req, _ := http.NewRequest("GET", seatGeekAPI+eventsEnp+"?sort=score.desc&score.gt=0.90&per_page="+strconv.Itoa(numOfEvent), nil)
	req.SetBasicAuth(os.Getenv(GEEKUSER), os.Getenv(GEEKPASS))
	resp, doErr := client.Do(req)
	if doErr != nil {
		log.Print("Error to seatGeek: ", doErr)
	}

	return resp
}
