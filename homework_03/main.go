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

func Reload(ch1 chan<-string) {
	defer wg1.Done()
	ch1<-"ammunition"
	fmt.Println("已装弹 ")
	//time.Sleep(time.Second)
}

func Aim(ch1 <-chan string,ch2 chan<-string) {
	defer wg2.Done()
	for t:=range ch1 {
		ch2<-t
		fmt.Println("已瞄准")
		//time.Sleep(time.Second)
	}
}

func Launch(ch2 chan string) {
	for range ch2 {
		fmt.Println("炮弹已发送")
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go Launch(ch2)

	for i:=0;i<10;i++ {
		wg1.Add(1)
		go Reload(ch1)
	}

	for i:=0;i<5;i++ {
		wg2.Add(1)
		Aim(ch1,ch2)
	}

	go func() {
		wg1.Wait()
		close(ch1)
		fmt.Println("close ch1")
	}()

	go func() {
		wg2.Wait()
		close(ch2)
		fmt.Println("close ch2")
	}()
	time.Sleep(6*time.Second)
	fmt.Println("done")
}