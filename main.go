package main

import (
	"log"
	"net/http"

	"github.com/crobbins21/seatGeek.git/controllers"
	"github.com/crobbins21/seatGeek.git/services"
	"github.com/jasonlvhit/gocron"
)

func main() {

	log.Print("Server started...")
	go executeEventRefresh()

	http.HandleFunc("/", controllers.Health)
	http.HandleFunc("/getEvents", controllers.GetEvents)
	http.ListenAndServe(":3000", nil)

}

func executeEventRefresh() {
	gocron.Every(6).Hour().Do(services.UpdateDatabase)
	<-gocron.Start()
}
