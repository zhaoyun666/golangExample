package _defer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup
var lock sync.Mutex
var count int

func TestDefer(t1 *testing.T) {
	//defer run()
	//defer walk()
	//defer fly()

	go goroutine_1()

	go goroutine_2()

	wg.Wait()
	time.Sleep(1 * time.Second)
}

func goroutine_1() {
	wg.Add(1)
	defer wg.Done()
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Println("Current count result: ", count)
}

func goroutine_2() {
	wg.Add(1)
	defer wg.Done()
	c := count + 5
	fmt.Println("Goroutine 2 result: ", c)
}

func run() {
	fmt.Println("Run, Quickly!")
}

func walk() {
	fmt.Println("Walk, Slowly!")
}

func fly() {
	fmt.Println("Fly, Quickly!")
}
