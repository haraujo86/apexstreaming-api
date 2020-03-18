package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// função principal
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/join", ParticipantJoined).Methods("POST")

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
	var p Participant
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		return
	}
	InsertParticipant(p)

	defer r.Body.Close()
}
