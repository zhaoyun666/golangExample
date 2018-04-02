package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var c chan int
var w sync.WaitGroup

func main() {
	c = make(chan int)
	go wait("coffee", 1)
	go wait("orange juice", 2)
	go wait("beer", 3)
	go read(c)
	stop()
}

func wait(name string, n int) {
	w.Add(1)
	fmt.Println(fmt.Sprintf("%s is ready, number %d", name, n))
	c <- n
	w.Done()
}

func read(c chan int) {
	w.Add(1)
	for v := range c {
		fmt.Println(v)
	}
	w.Done()
}
func stop() {
	var sigal = make(chan os.Signal)
	signal.Notify(sigal)
	for {
		s := <-sigal
		switch s {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:

		default:
			continue
		}
		break
	}
}
