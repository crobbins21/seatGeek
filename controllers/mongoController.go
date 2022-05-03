package controllers

import (
	"context"
	"net/http"

	"github.com/crobbins21/seatGeek.git/services"
)

const (
	GET         = "GET"
	blankString = ""
)

func GetEvents(w http.ResponseWriter, req *http.Request) {

	if http.MethodGet == GET {
		events := services.GetEvents(context.Background())

		WriteResponse(w, 200, events)
	} else {
		WriteResponse(w, 400, nil)
	}
}
