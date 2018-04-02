package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func main() {

	go wait("A", 1)
	fmt.Println(<-ch)

}

func wait(name string, n int) {
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println(name, "is ready")
	ch <- 1
}
