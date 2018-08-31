package timer

import (
	"fmt"
	"time"
)

func Task() {
	done := make(chan struct{}) // 用来等待携程结束

	timer := time.NewTimer(6 * time.Second)

	go func() {
		fmt.Printf("Now is %s\n", <-timer.C)
		done <- struct{}{}
	}()
	fmt.Println("Print in main")
	<-done
	close(done)
}

func Worker() {
	input := make(chan interface{})
	go func() {
		for i := 0; i < 6; i++ {
			input <- i
		}
	}()

	t1 := time.NewTimer(time.Second * 6)
	t2 := time.NewTimer(time.Second * 10)

	for {
		select {
		case msg := <-input:
			fmt.Println(msg)
		case <-t1.C:
			fmt.Println("6s timer")
			t1.Reset(time.Second * 10)
		case <-t2.C:
			fmt.Println("10s timer")
			t2.Stop()
		}
	}
}


