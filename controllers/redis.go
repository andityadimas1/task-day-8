package controllers

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func GetRedis() {
	// bikin pool untuk connection ke redis
	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:5678")
	}, 10)

	pool.MaxActive = 10

	// ambil 1 koneksi
	conn := pool.Get()
	defer conn.Close()

	// coba untuk cache data dari data map
	getData, _ := redis.Bytes(conn.Do("GET"))
	if getData != nil { // ketika ada datanya di redis
		log.Println("Data Found!")
		log.Println(string(getData))
		return true, getData
	}
	return false, getData
}

func SetRedis(key string, value string) {
	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:5678")
	}, 10)
	pool.MaxActive = 10

	// Get connection
	conn := pool.Get()
	defer conn.Close()

	// Finding Data with key
	conn.Do("SETEX", key, 30, string(value))
}
