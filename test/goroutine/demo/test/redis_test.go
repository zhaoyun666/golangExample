package test

import (
    "testing"
    "learning-golang-process/test/goroutine/demo/cache"
)

func TestRedis(T *testing.T) {
    var i int = 0
    for{
        i++
        cache.Rd.LPush("task1", i)
    }
}
