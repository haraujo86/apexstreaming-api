package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haraujo86/apexstreaming-api/model"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/participant", participantJoined).Methods("POST")
	router.HandleFunc("/participant/{id}", participantWhoIs).Methods("GET")

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

func participantJoined(w http.ResponseWriter, r *http.Request) {
	var p model.Participant
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		return
	}
	model.InsertParticipant(p)
	defer r.Body.Close()
}

func participantWhoIs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var id string = vars["id"]
	participant, err := model.GetParticipant(id)
	if err != nil {
		return
	}

	participantJson, err := json.Marshal(participant)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(participantJson)
}
