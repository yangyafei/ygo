package demo

import (
	"fmt"
	"runtime"
	"time"
)

func fibonacci(n int, ch chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func fibonacci2(ch chan int, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Do_fib1() {
	ch := make(chan int, 10)
	go fibonacci(10, ch)
	for i := range ch {
		fmt.Println(i)
	}
}

func Do_timeout() {
	c := make(chan int)
	d := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				d <- true
				break
			}
		}
	}()
	<-d
}

func Do_fib2() {
	ch := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	fibonacci2(ch, quit)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func Doit_runtime() {
	go say("world") //开一个新的Goroutines执行
	say("hello")    //当前Goroutines执行

	fmt.Println(runtime.NumCPU())       //返回 CPU 核数量
	fmt.Println(runtime.NumGoroutine()) //返回正在执行和排队的任务总数

}
