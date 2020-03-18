package main

import "fmt"

type Participant struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

var participants []Participant

//Insert partcipants getting into a specfic room
func InsertParticipant(p Participant) {
	fmt.Println(p)
}
