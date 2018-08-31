package test

import (
	"fmt"
	"math/rand"
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

// 多路复用

// 生成随机数
func rand_generator_1() int {
	return rand.Int()
}

// 返回通道 channel
func rand_generator_2() chan int {
	// 创建通道
	out := make(chan int)
	// 创建goroutine
	go func() {
		for {
			out <- rand_generator_1()
		}
	}()
	return out
}

func rand_generator_3() chan int {
	rand_service_handler1 := rand_generator_2()
	rand_service_handler2 := rand_generator_2()

	out := make(chan int)

	go func() {
		for {
			out <- <-rand_service_handler1
			fmt.Println("--------")
		}
	}()

	go func() {
		for {
			out <- <-rand_service_handler2
			fmt.Println("+++++++++")
		}
	}()

	return out
}

func TestMultiplexing(t *testing.T) {
	rand_service_handler := rand_generator_3()
	fmt.Printf("%d\n", <-rand_service_handler)
	fmt.Printf("%d\n", <-rand_service_handler)
	fmt.Printf("%d\n", <-rand_service_handler)
}

// 多路复用结束

func TestLoop(t *testing.T) {
	go func() {
		for {
			fmt.Println(time.Now().Unix())
			//fmt.Println(rand.Int())
		}

	}()
	time.Sleep(3 * time.Second)
}

// 并发循环

func TestConcurrentLoop(t *testing.T) {
	data := make([]int, 0)
	for i := 0; i < 60; i++ {
		for j := 1; j < 4; j++ {
			data = append(data, j)
		}
	}

	N := len(data)
	m := make(chan int, N)
	for i, x := range data {
		go func(i, x int) {
			doSomething(i, x)
			m <- 0
		}(i, x)
	}
	fmt.Println("并发循环结束")
	// 等待结束，查看结果
	for i := 0; i < N; i++ {
		<-m
	}
	fmt.Println("任务完成")
}

func doSomething(i int, x int) {
	fmt.Printf("i=%d, x=%d", i, x)
}
