package seatGeek

import (
	"log"
	"net/http"
	"os"
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

func GetEvents(numOfEvent int) *http.Response {
	client := &http.Client{Timeout: time.Duration(1) * time.Second}

	req, _ := http.NewRequest("GET", seatGeekAPI+eventsEnp+"?per_page=100", nil)
	req.SetBasicAuth(os.Getenv(GEEKUSER), os.Getenv(GEEKPASS))
	resp, doErr := client.Do(req)
	if doErr != nil {
		log.Print("Error to seatGeek: ", doErr)
	}

	defer resp.Body.Close()

	return resp
}
