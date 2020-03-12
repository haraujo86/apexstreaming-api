package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// função principal
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/join", ParticipantJoined).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

type Post struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func ParticipantJoined(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//var post Post
	//_ = json.NewDecoder(r.Body).Decode(post)
	//post.ID = strconv.Itoa(rand.Intn(1000000))
	//posts = append(posts, post)
	//json.NewEncoder(w).Encode(&post)
	//fmt.Println(post.Name)

	var p Post
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		return
	}
	defer r.Body.Close()
	fmt.Println(p)
}
