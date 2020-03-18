package main

import "fmt"

type Participant struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

var participants []Participant

//InsertParticipant insert a row to identify who is in the rrom
func InsertParticipant(p Participant) {
	fmt.Println(p)
}
