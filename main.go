package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stephenalexbrowne/zoom"
	"github.com/stephenalexbrowne/zoom_example/controllers"
	"github.com/stephenalexbrowne/zoom_example/models"
	"log"
	"net/http"
)

func main() {
	port := "6060"

	err := models.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	defer zoom.Close()

	r := route()

	http.Handle("/", r)
	portName := ":" + port
	fmt.Printf("server listening on port %s...\n", port)
	http.ListenAndServe(portName, nil)
}

func route() *mux.Router {
	r := mux.NewRouter()

	// persons
	personsController := controllers.PersonsController{}
	r.HandleFunc("/persons", personsController.Create).Methods("POST")
	r.HandleFunc("/persons/{Id}", personsController.Update).Methods("PUT")
	r.HandleFunc("/persons/{Id}", personsController.Show).Methods("GET")
	r.HandleFunc("/persons/{Id}", personsController.Delete).Methods("DELETE")

	return r
}
