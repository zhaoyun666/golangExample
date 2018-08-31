package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	// 增减操作
	var a int32
	fmt.Println("a : ", a)

	// 函数名以ADD为前缀，加具体类型名， 参数一：是指针类型 参数二： 与参数一类型总是相同
	n_a := atomic.AddInt32(&a, 3)
	fmt.Println("n_a : ", n_a)

	n_a = atomic.AddInt32(&a, -2)
	fmt.Println("n_a : ", n_a)

	// CAS(Compare And Swap)比较并且交换操作, 函数名称以CompareAndSwap为前缀，并具体类型名
	var b int32
	fmt.Println("b : ", b)

	// 函数会先判断参数一指向的值与参数二是否相等，如果相等，则用参数三替换参数一的值。 最后返回是否替换成功
	atomic.CompareAndSwapInt32(&b, 0, 3)
	fmt.Println("b : ", b)

	// 载入操作，当我们对于某个变量进行读取操作时，可能该变量正在被其他操作改变，或许我们读取的是被修改了一半的数据。所以我们通过load这类函数来确保我们正确的读取
	// 函数名以Load为前缀，加具体类型名
	var c int32
	wg := sync.WaitGroup{}
	// 启动100个goroutine
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tmp := atomic.LoadInt32(&c)
			if !atomic.CompareAndSwapInt32(&c, tmp, (tmp + 1)) {
				fmt.Println("c 被修改失败")
			}
		}()
	}
	wg.Wait()
	// c的值有可能不等于100， 频繁修改变量值情况下，CAS皂搓有可能不成功
	fmt.Println("c : ", c)

	// 存储操作，与载入函数相对应，提供原子的存储函数，函数名以Store为前缀，加具体类型名
	var d int32
	fmt.Println("d : ", d)
	// 存储某个值时，任何CPU都不会与该值进行读或写操作，存储操作总会成功，它不关心旧值是什么，与CAS不同
	atomic.StoreInt32(&d, 666)
	fmt.Println("d : ", d)
	atomic.StoreInt32(&d, 66)
	fmt.Println("new d : ", d)

	// 交换操作，直接设置新值，返回旧值，与CAS不同，他不关心旧值。函数名以Swap为前缀，加具体类型名
	var e int32
	wg1 := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			temp := atomic.LoadInt32(&e)
			atomic.SwapInt32(&e, (temp + 1))
		}()
	}
	wg1.Wait()
	fmt.Println("e : ", e)
}
