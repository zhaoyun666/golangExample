package cache

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

func InitCache() redis.Conn {
	Rd, err := redis.Dial("TCP", "127.0.0.1:3306")
	if err != nil {
		log.Fatalf("err:%s", err.Error())
	}

	Rd.Flush()
	return Rd
}
