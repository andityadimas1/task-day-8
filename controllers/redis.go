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
	getData, _ := redis.Bytes(conn.Do("GET", key))
	if getData != nil { // ketika ada datanya di redis
		log.Println("Data Found!")
		log.Println(string(getData))
		return true, getData
	}
	return false, getData
}

// 	// untuk ambil cache di redis
// 	reply, _ := redis.Bytes(conn.Do("GET", "mahasiswa"))
// 	if reply != nil { // ketika ada datanya di redis
// 		log.Println("data dari redis ada, ini datanya")
// 		log.Println(string(reply))
// 	} else { // ketika ngga ada datanya di redis
// 		log.Println("data dari redis ngga ada, coba ambil dari db")
// 		output, _ := json.Marshal(user)

// 		// untuk set cache di redis
// 		conn.Do("SETEX", "mahasiswa", 20, string(output))
// 	}
// }
