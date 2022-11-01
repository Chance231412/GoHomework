package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	Reminders = make(map[string] *Reminder)
	Scaner = bufio.NewScanner(os.Stdin)
)

var (
	wg sync.WaitGroup
)

func main() {
	Menu()
}

const (
	Repeat = iota + 1
	Disposable
)

type Reminder struct {
	Title string
	Note string
	Time time.Time   //定时时间
	Type int
	Can chan struct{}
}

func Menu() {
	for  {
		content := []string{"单次提醒功能","重复提醒功能","删除提醒","退出"}
		Show(content)
		op := GetOption()
		switch op {
		case 1:
			SetDisposableReminder()
		case 2:
			SetRepeatReminder()
		case 3:
			DeleteReminder()
		case 4:
			//退出前要关闭所有gorotuine
			BeforeExit()
			os.Exit(0)
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
}

func Show(s []string){
	if len(s)==0 {
		fmt.Println("该列表为空!!!")
		return
	}
	for k,v := range s {
		fmt.Printf("%d、%s\n",k+1,v)
	}
}

func getInput() string {
	Scaner.Scan()
	return Scaner.Text()
}

func GetOption()int {
	var t int
	fmt.Scanln(&t)
	return  t
}

func SetRepeatReminder(){
	var title string
	var note string
	var t string
	fmt.Println("请输入记事本名称：")
	title = getInput()
	fmt.Println("请输入提醒的内容：")
	note = getInput()
	fmt.Println("请输入设置的时间(格式“15:04:05”)：")
	t = getInput()
	time,err := time.Parse(`15:04:05`,t)
	if err!=nil {
		fmt.Println(err)
		return
	}
	reminder := NewReminder(title,note,Repeat,time)
	AddReminder(&reminder)
	fmt.Println("设置成功!")
}

func SetDisposableReminder() {
	var title string
	var note string
	var t string
	fmt.Println("请输入记事本名称：")
	title = getInput()
	fmt.Println("请输入提醒的内容：")
	note = getInput()
	fmt.Println("请输入设置的时间(格式为：\"2006-01-02 15:04:05\")：")
	t = getInput()
	time,err := time.Parse(`2006-01-02 15:04:05`,t)

	if err!=nil {
		fmt.Println(err)
		return
	}
	reminder := NewReminder(title,note,Disposable,time)
	AddReminder(&reminder)
	fmt.Println("设置成功!")
}

func AddReminder(r *Reminder) {
	Reminders[r.Title]=r
	wg.Add(1)
	go r.Run()
}

func DeleteReminder() {
	contents := make([]string,0,1)
	count:=0
	for k,_ := range Reminders {
		count++
		contents = append(contents,k)
	}
	Show(contents)
	var op int
	for  {
		fmt.Println("按-1返回")
		op=GetOption()
		if op==-1 {
			return
		}else if op>count {
			fmt.Println("输入不合法，请重新输入: ")
			continue
		} else {
			break
		}
	}

	Reminders[contents[op-1]].Can<- struct{}{}
}

func BeforeExit() {
	defer fmt.Println("退出成功")
	for _,v:=range Reminders {
		v.Cancel()
	}
	wg.Wait()
}

func NewReminder(title string,note string,typ int,time time.Time)Reminder {
	can := make(chan struct{})
	return Reminder{
		Title: title,
		Note: note,
		Type: typ,
		Time: time,
		Can: can,
	}
}

func (r *Reminder)Remind() {
	fmt.Println("*************  温馨提醒  *************")
	fmt.Println(r.Note)
	fmt.Println("************ 一定要记得哦 *************")
}

func (r *Reminder)Run() {
	defer wg.Done()
	for  {
		nextTime := r.nextTime()
		duration := nextTime.Sub(time.Now())
		timer := time.NewTimer(duration)
		select {
		case <-timer.C:
			r.Remind()
			if r.Type==Disposable {
				r.Cancel()
				return
			}
		case <-r.Can:
			if r.Type==Repeat {
				fmt.Printf("\n%s 已取消\n",r.Title)
			}
			return
		}
	}
}

func (r *Reminder)nextTime()time.Time {
	if r.Type==Disposable {
		local,_ := time.LoadLocation("Asia/Shanghai")
		return time.Date(r.Time.Year(),r.Time.Month(),r.Time.Day(),r.Time.Hour(),r.Time.Minute(),r.Time.Second(),0,local)
	}

	now := time.Now()
	nextTime := time.Date(now.Year(),now.Month(),now.Day(),r.Time.Hour(),r.Time.Minute(),r.Time.Second(),0,now.Location())
	if nextTime.After(now) {
		return nextTime
	}
	return nextTime.AddDate(0,0,1)
}

func (r *Reminder)Cancel() {
	delete(Reminders,r.Title)
	//r.Can <- struct {}{}
	close(r.Can)
}
