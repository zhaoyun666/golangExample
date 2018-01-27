package main

import (
	"time"
	"fmt"
)

var c chan int

func main(){
	c = make(chan int)
	go wait("coffee", 1)
	go wait("orange juice", 2)
	go wait("beer", 3)
	
	i := 0
	L:
	for{
		select {
		case <-c:
			i++
			if i > 2{
				break L
			}
		}
	}
}

func wait(name string, n int){
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println(name, "is ready")
	c<-1
}
