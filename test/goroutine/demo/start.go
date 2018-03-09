package main

import (
	"learning-golang-process/test/goroutine/demo/db"
	"math/rand"
	"os"
	"sync"
	"time"
	"learning-golang-process/test/goroutine/demo/cache"
	"fmt"
)
var mu sync.RWMutex
var ch = make(chan int)
//生成随机int
func getRand(num int) int {
	rand.Seed(time.Now().UnixNano())
	var mu sync.Mutex
	mu.Lock()
	v := rand.Intn(num)
	mu.Unlock()
	return v
}
//生成随机字符串
func GetRandomString(length int64) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
func main() {
	var num int64
	for num = 0; num < 100; num ++ {
		//go write()
		go readRedis(num, num + 1)
	}

	for {
		select {
		case <-ch:
		default:
			goto L
		}
	}
	L:
	os.Exit(1)
}

func write() {
	mu.RLock()
	res, _ := db.MysqlEngine.Exec("insert into actor(name, password) values(?, ?)", GetRandomString(2), GetRandomString(3))
	if res == nil {

	}
	ch <- 1
	mu.RUnlock()
}

func readRedis(i, j int64) {
	r := cache.Rd.LRange("task1", i, j)
	fmt.Println(r.Result())
	cache.Rd.SRem("task1", []int64{i, j})
	ch<-1
}