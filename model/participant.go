package model

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/haraujo86/apexstreaming-api/infrastructure"
	"github.com/haraujo86/apexstreaming-api/util"
)

type Participant struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

var participants []Participant

// InsertParticipant insert a specific participant into a cache room
func InsertParticipant(p Participant) (string, error) {

	var err error
	conn := infrastructure.GetRedis().Conn

	_, err = conn.Do("HMSET", p.ID, "ID", p.ID, "name", p.Name, "content", p.Content)

	if err != nil {
		return p.ID, fmt.Errorf("Error setting key %s: %v", p.ID, err)
	}
	return p.ID, err
}

// GetParticipant return a specific participant
func GetParticipant(key string) (Participant, error) {

	conn := infrastructure.GetRedis().Conn

	var part Participant
	var scanValues = make(map[string]string)
	var data []interface{}
	data, err := redis.Values(conn.Do("HGETALL", key))

	if err != nil {
		return part, err
	}

	scanValues, err = util.ScanToMap(data)

	if err != nil {
		return part, err
	}

	return convertMapToParticipant(scanValues), nil
}

func convertMapToParticipant(m map[string]string) (p Participant) {
	p.ID = m["ID"]
	p.Name = m["name"]
	p.Content = m["content"]

	return p
}
