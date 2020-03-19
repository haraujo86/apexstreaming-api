package model

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/haraujo86/apexstreaming-api/infrastructure"
)

type Participant struct {
	ID      string `json:"ID"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

var participants []Participant

// InsertParticipant insert a specific participant into a cache room
func InsertParticipant(p Participant) (string, error) {

	var err error
	conn := infrastructure.GetRedis().Conn

	_, err = conn.Do("HMSET", p.ID, "ID", p.ID, "Name", p.Name, "Content", p.Content)

	if err != nil {
		return p.ID, fmt.Errorf("Error setting key %s: %v", p.ID, err)
	}
	return p.ID, nil
}

// GetParticipant return a specific participant
func GetParticipant(key string) (Participant, error) {

	conn := infrastructure.GetRedis().Conn

	var participant Participant
	var data []interface{}

	data, err := redis.Values(conn.Do("HGETALL", key))
	if err != nil {
		return participant, err
	}

	err = redis.ScanStruct(data, &participant)
	if err != nil {
		return participant, err
	}

	return participant, err
}
