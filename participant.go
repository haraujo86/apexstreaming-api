package main

import (
	"fmt"
	"log"
)

type Participant struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

var participants []Participant

//InsertParticipant insert a specific participant into a cache room
func InsertParticipant(p Participant) {
	var err error

	_, err = GetRedis().conn.Do("HMSET", p.ID, "ID", p.ID, "name", p.Name, "content", p.Content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}
