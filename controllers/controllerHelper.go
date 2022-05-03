package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, statusCode int, body interface{}) {

	w.WriteHeader(statusCode)

	data, err := json.MarshalIndent(body, "", " ")
	if err != nil {
		log.Print("Error marshaling the data to output")
	} else {

		w.Write(data)

	}

}
