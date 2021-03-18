package controllers

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func GetRedis(key string) (bool, []byte) {
	// bikin pool untuk connection ke redis
	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:5678")
	}, 60)

	pool.MaxActive = 60

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
	}, 60)
	newPool.MaxActive = 60

	conn := newPool.Get()
	defer conn.Close()

	conn.Do("SETEX", key, 70, string(value))
}
