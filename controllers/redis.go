package controllers

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func GetRedis(key string) (bool, []byte) {
	// bikin pool untuk connection ke redis
	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:5678")
	}, 10)

	pool.MaxActive = 10

	conn := pool.Get()
	defer conn.Close()

	getData, _ := redis.Bytes(conn.Do("GET", key))
	if getData != nil { // ketika ada datanya di redis
		log.Println("Data Found!")
		log.Println(string(getData))
		return true, getData
	}
	return false, getData
}

func SetRedis(key string, value string) {
	newPool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:5678")
	}, 10)
	newPool.MaxActive = 10

	conn := newPool.Get()
	defer conn.Close()

	conn.Do("SETEX", key, 30, string(value))
}
