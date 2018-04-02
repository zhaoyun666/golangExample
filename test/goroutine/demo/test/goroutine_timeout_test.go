package test

import (
	"fmt"
	"testing"
	"time"
)

// 测试goroutine 通道超时
func TestGoroutineTimeout(t *testing.T) {
	timeout := make(chan bool, 1)
	ch := make(chan int)

	// 每秒生产true数据
	go func() {
		time.Sleep(time.Second)
		timeout <- true
	}()

	select {
	case <-ch:
		fmt.Println("ch pop")
	case <-timeout:
		fmt.Println("timeout!")
		break
	}
}

// 测试channel 是否满了
func TestChannelIsFull(t *testing.T) {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	select {
	case ch1 <- 2:
		fmt.Println("channel is not full !")
	default:
		fmt.Println("channel is full !")
	}
}

// 生产者与消费者
func TestProductAndConsumer(t *testing.T) {
	ch := make(chan int, 1)
	timeout := make(chan bool, 1)

	// 生产
	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
		}
	}()

	// 消费
	var val int
	// 设置最大操作数
	for i := 0; i < 10; i++ {
		// 每秒生产一下true数据
		go func() {
			time.Sleep(time.Second)
			timeout <- true
		}()

		select {
		case val = <-ch:
			fmt.Printf("消费了：%d\n", val)
		case <-timeout:
			fmt.Println("timeout!")
		}
	}
}

// select之默认值操作
func TestSelectDefault(t *testing.T) {
	ch1 := make(chan int64, 1)
	ch2 := make(chan int64, 1)

	// 协程操作不仅能同步调用还可以异步调用
	ch1 <- time.Now().Unix()
	//同步
	var val int64
	select {
	case val = <-ch1:
		fmt.Println("ch1 pop one element", val)
	case <-ch2:
		fmt.Println("ch2 pop one element")
	default:
		fmt.Println("default")
	}
}

// 共享变量 start ---------------------------------------------------
func TestSharedVariables(t *testing.T) {
	v := &shared_variables{
		make(chan int),
		make(chan int),
	}
	sharedVariables(v)
	// 读取初始化
	fmt.Println(<-v.reader)
	// 写入一个值
	v.writer <- 1
	// 读取新写入的值
	fmt.Println(<-v.reader)
}

type shared_variables struct {
	reader chan int
	writer chan int
}

func sharedVariables(v *shared_variables) {
	go func() {
		var val int = 0
		for {
			select {
			case val = <-v.writer:
			case v.reader <- val:

			}
		}
	}()
}

func never_leak(ch chan int) {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second)
		timeout <- true
	}()

	select {
	case <-ch:
	case <-timeout:
	}
}

// 共享变量 end -----------------------------------------------------
