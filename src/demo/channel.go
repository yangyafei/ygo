package demo

import (
	"fmt"
	"time"
)

func Channel()  {
	fmt.Println("Do not communicate by sharing memory, instead, share memory by communicating.")
}

//chan关闭后是否还可以获取数据
func ChannelCheckClose()  {
	ch := make(chan int,5)
	ch <- 18
	close(ch)

	x, ok := <- ch
	if ok{
		fmt.Println("received: ", x)
	}

	x, ok = <- ch
	if !ok{
		fmt.Println("channel close, data invaild.")
	}

	println("chanel关闭后 如果chan有数据还是可以读取出来的")
}

//channel应用：定时
func ChannelApplicationCrond()  {
	go func() {
		ticker := time.Tick(1*time.Second)
		for{
			select {
			case <- ticker:
				fmt.Println("执行 1s定时任务")
			}
		}
	}()
}


//channel应用：解耦生产者和消费者
func ChannelApplicationDecouple()  {
	task := make(chan int, 100)

	//消费者
	go channelApplicationDecoupleCustomer(task)

	//生产者
	for i:=0;i<10;i++{
		task <- i
	}

	select {
	case <-time.After(5*time.Second):
	}
}

func channelApplicationDecoupleCustomer(task chan int)  {
	const N = 5
	for i:=0; i<N;i++{
		go func(id int) {
			ta := <- task
			fmt.Printf("finish task: %d by worker %d\n", ta, id)
		}(i)
	}
}

//channle应用：控制并发
func ChannelApplicationControlConCurrency()  {

	limit := make(chan int,1)
	woker := [5]int{1,3,5,7,9}
	for _,w := range woker{
		go func() {
			limit <- 1 //
			// w() 执行任务
			fmt.Println("执行任务 ", w)
			<- limit
		}()
		time.Sleep(1*time.Second)
	}
}

