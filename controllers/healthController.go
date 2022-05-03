package controllers

import "net/http"

func Health(w http.ResponseWriter, req *http.Request) {
	WriteResponse(w, 200, "Health probe")
}
