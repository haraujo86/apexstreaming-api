package util

import (
	"github.com/gomodule/redigo/redis"
)

// ScanToMap convert to map a returned list interface object from redis
func ScanToMap(values []interface{}) (map[string]string, error) {
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
			results[key] = ""
		}

	}

	return results, nil
}
