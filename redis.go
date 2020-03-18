package main

import (
	"log"
	"sync"

	"github.com/gomodule/redigo/redis"
)

//RedisConn is a structure to keep a redis connection session
type RedisConn struct {
	conn redis.Conn
}

var rdb *RedisConn
var once sync.Once

//GetRedis returns a instance of redis.Conn
func GetRedis() *RedisConn {
	once.Do(func() {
		connection, err := redis.Dial("tcp", "localhost:6379")
		if err != nil {
			log.Fatal(err)
		}
		rdb = &RedisConn{conn: connection}
	})
	return rdb
}
