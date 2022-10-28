package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg1 sync.WaitGroup
	wg2 sync.WaitGroup
)

const (
	ProducerNum = 10
)


func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	closing := make(chan string)
	prepared := make(chan string)


	go func(){
		prepared<-"开炮"
	}()

	//执行closer
	go closer(ch1,ch2,closing)

	//执行producer
	for i:=0;i<10;i++ {
		wg1.Add(1)
		go producer(ch1,closing,prepared)
	}

	//执行deliver
	for i:=0;i<5;i++ {
		wg2.Add(1)
		go deliver(ch1,ch2)
	}

	go func() {
		//最后main goroutine中会向prepared中多发送一个信息，可是此时所有的producer都已经关闭，没有goroutine来接收，因此必须要创建一个goroutine来接收，否则会报错
		wg1.Wait()
		<-prepared
	}()
	//consumer
	for range ch2 {
		fmt.Println("发射!")
		prepared<-"ok"
		time.Sleep(200*time.Millisecond)
	}
	fmt.Println("done")
}

func closer(ch1 chan string,ch2 chan string,closing chan<-string) {
	//不会引入包，只需要将这个部分的控制逻辑换成键盘输入就可以了.....大概吧.....
	time.Sleep(3*time.Second)
	for i:=0;i<ProducerNum;i++ {
		closing<-"over"
	}
	wg1.Wait()
	close(ch1)
	wg2.Wait()
	close(ch2)
}

func producer(ch1 chan<-string,closing <-chan string,prepared <-chan string) {
	defer wg1.Done()
	for {
		//尝试关闭producer
		select {
		case <-closing:
			return
		default:
		}
		select {
		case <-prepared:
			fmt.Printf("装弹-->")
			select {
			case ch1<-"ok":
			}
		default:
		}
	}
}

func deliver(ch1 <-chan string,ch2 chan<-string) {
	defer wg2.Done()
	for range ch1 {
		fmt.Printf("瞄准-->")
		ch2<-"炮弹"
	}
}
