package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var w sync.WaitGroup

func TestRedis(T *testing.T) {
	a := 12
	switch a {
	case 12:
		fmt.Println("ok")
	default:
		fmt.Println("Default Value is output")
	}
}

func TestSelect(T *testing.T) {
	c := make(chan int)
	go func() {
		for {
			c <- 1
		}

	}()
	for {
		select {
		case <-c:
			x := <-c
			fmt.Println(x)
		default:
			fmt.Println("D")
			break
		}
	}

}

func TestChannel(t *testing.T) {
	ch1 := make(chan int)
	w.Add(2)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
	}()

	go channel(ch1)
	w.Done()
	w.Done()
	w.Wait()
	time.Sleep(6 * time.Second)
}

func channel(ch <-chan int) {
	for c := range ch {
		fmt.Println(c)
	}

}
