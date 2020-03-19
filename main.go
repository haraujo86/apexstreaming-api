package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haraujo86/apexstreaming-api/model"
	"github.com/rs/cors"
)

// função principal
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/join", ParticipantJoined).Methods("POST")
	router.HandleFunc("/participant/{id}", ParticipantWhoIs).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Access-Control-Allow-Origin", "Accept"},
		ExposedHeaders:   []string{"Content-Type", "Access-Control-Allow-Origin", "Accept"},
		AllowedMethods:   []string{"POST"},
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8000", handler))
}

func ParticipantJoined(w http.ResponseWriter, r *http.Request) {
	var p model.Participant
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		return
	}
	model.InsertParticipant(p)
	defer r.Body.Close()
}

func ParticipantWhoIs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(model.GetParticipant(vars["id"]))
}
