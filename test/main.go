package main

import (
	"fmt"
	"sync"
)

var (
	wg1 sync.WaitGroup
)

func Reload(ch1 chan<-string) {
	fmt.Println("已装弹 ")
	ch1<-"ammunition"
	wg1.Done()
}


func main() {
	ch1:=make(chan string)
	for i:=0;i<10;i++ {
		wg1.Add(1)
		go Reload(ch1)
	}

	go func() {
		wg1.Wait()
		close(ch1)
		//fmt.Println("close ch1")
	}()
	for range ch1 {
		fmt.Println("炮弹已发射")
	}
}