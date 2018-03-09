package cache

import (
    "gopkg.in/redis.v5"
    "time"
)

var Rd *redis.Client

func init() {
    Rd = redis.NewClient(&redis.Options{
        Addr:         "127.0.0.1:6379",
        DialTimeout:  10 * time.Second,
        ReadTimeout:  30 * time.Second,
        WriteTimeout: 30 * time.Second,
        PoolSize:     10,
        PoolTimeout:  30 * time.Second,
        DB: 0,
    })
    Rd.FlushDb()
}
