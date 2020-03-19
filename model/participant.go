package model

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/haraujo86/apexstreaming-api/infrastructure"
)

type Participant struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

var participants []Participant

//InsertParticipant insert a specific participant into a cache room
func InsertParticipant(p Participant) (string, error) {

	var err error
	conn := infrastructure.GetRedis().Conn

	_, err = conn.Do("HMSET", p.ID, "ID", p.ID, "name", p.Name, "content", p.Content)

	if err != nil {
		return p.ID, fmt.Errorf("Error setting key %s: %v", p.ID, err)
	}
	return p.ID, err
}

func GetParticipant(key string) Participant {

	conn := infrastructure.GetRedis().Conn

	var part Participant
	var scanValues = make(map[string]string)
	var data []interface{}
	data, err := redis.Values(conn.Do("HGETALL", key))

	if err != nil {
		fmt.Println("Error getting key %s: %v", key, err)
		return part
	}
	scanValues, err = scanMap(data)
	return convertMapToParticipant(scanValues)
}

func convertMapToParticipant(m map[string]string) Participant {
	var participantSplit Participant

	participantSplit.ID = m["ID"]
	participantSplit.Name = m["name"]
	participantSplit.Content = m["content"]

	return participantSplit
}

func scanMap(values []interface{}) (map[string]string, error) {
	results := make(map[string]string)
	var err error

	for len(values) > 0 {
		var key string
		var value string

		values, err = redis.Scan(values, &key)

		if err != nil {
			return nil, err
		}

		if len(values) > 0 {
			values, err = redis.Scan(values, &value)

			if err != nil {
				return nil, err
			}

			results[key] = value
		} else {
			fmt.Println("Unable to find value for %s.", key)
			results[key] = ""
		}

	}

	return results, nil
}
